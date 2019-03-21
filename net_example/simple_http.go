package main

import (
	"time"
	"sync"
	"math/rand"
	"strconv"
	"net/http"
	"io"
)

type Session struct {
	Username string
	Uuid string
	Expire time.Time
	Lock sync.Mutex //保护session
}

type Profile struct {
	Password string
	CreateTime time.Time

}

type UserCreate struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

var WebSessions=make(map[string]Session)
var ProfileMap = make(map[string]Profile)

func GenSession(username string) bool{

	var s Session
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.Username = username
	rand.Seed(time.Now().Unix())
	s.Uuid = strconv.Itoa(rand.Intn(1000))
	s.Expire = time.Now().Add(30*time.Second)
	//将session存入临时数据
	WebSessions[username] = s
	return true
}

func ValidSession(_uuid string,username string) bool{
	session,ok := WebSessions[username]
	if !ok{
		return false
	}
	uuid := session.Uuid
	expire := session.Expire
	if uuid != _uuid{
		return false
	}
	if expire.After(time.Now()){
		return true
	}else{
		return false
	}
}

func DestroySession(username string) bool{
	if _,exist:=WebSessions[username];exist==false{
		return false
	}
	delete(WebSessions,username)
	return true
}


func InsertSql(username,password string) bool{
	if _,err:=ProfileMap[username]; err== true{
		return false
	}
	var profile Profile
	profile.Password = password
	profile.CreateTime = time.Now()
	ProfileMap[username] = profile
	return true
}

func DeleteSql(username string) bool{
	if _,err:= ProfileMap[username];err==false{
		return false
	}
	delete(ProfileMap, username)
	return true
}



func CreateUser(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		//body,err := ioutil.ReadAll(r.Body)

		err := r.ParseForm() //r.Body与其不能同时存在
		if err!=nil{
			panic(err)
		}
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")
		if _,is:=ProfileMap[username];is!=false{
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, string("username exists"))
		}
		var userProfile Profile
		userProfile.Password=password
		userProfile.CreateTime=time.Now()
		ProfileMap[username] = userProfile
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, string("create success"))
		return
	}else if r.Method=="GET"{
		w.WriteHeader(http.StatusOK)
		io.WriteString(w,"hello world")
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	return

}
//
func Login(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		w.WriteHeader(http.StatusOK)
		io.WriteString(w,"Please login!!!")
		return
	}
	if r.Method == "POST"{
		err := r.ParseForm()
		if err!=nil{
			panic(err)
		}
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")
		if _,is:= ProfileMap[username];is==false{
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, string("username error"))
			return
		}
		if password != ProfileMap[username].Password{
			w.WriteHeader(http.StatusOK)
			io.WriteString(w,"password error")
			return
		}
		GenSession(username)

		cookie:= http.Cookie{
			Name:"sessionId",
			Value: WebSessions[username].Uuid,
			}
		http.SetCookie(w,&cookie)

		cookie2 := http.Cookie{
			Name:"username",
			Value:username,
		}
		http.SetCookie(w,&cookie2)
		w.WriteHeader(http.StatusOK)
		io.WriteString(w,"login success")
		return
	}


}

func Logout(w http.ResponseWriter, r *http.Request){
	cookies := r.Cookies()
	var name string
	for _,cookie:= range cookies{
		if cookie.Name=="username"{
			name = cookie.Value
			break
		}
	}
	is := DestroySession(name)
	if is{
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "logout success!!!")
		return
	}else{
		http.Redirect(w,r,"/login",http.StatusMovedPermanently)
		return
	}

}

func Info(w http.ResponseWriter, r *http.Request){
	cookies := r.Cookies()
	var username,sessionId string
	username = ""
	sessionId = ""
	for _,cookie :=range cookies{
		switch cookie.Name {
		case "username":
			username = cookie.Value
		case "sessionId":
			sessionId = cookie.Value
		default:
			continue
		}
	}
	//验证session
	is := ValidSession(sessionId, username)
	if is{
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "welcome to china!!!")
		return
	}else{
		http.Redirect(w,r,"/login",http.StatusMovedPermanently)
		return
	}

}

func main(){

	//s := InsertSql("chow","123456")
	//fmt.Println(ProfileMap,s)
	//DeleteSql("chow")
	//
	//fmt.Println(ProfileMap)
	http.HandleFunc("/create", CreateUser)
	http.HandleFunc("/login",Login)
	http.HandleFunc("/info",Info)
	http.HandleFunc("/logout",Logout)
	http.ListenAndServe(":8080",nil)

}