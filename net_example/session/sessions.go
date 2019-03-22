package session

import (
	"sync"
	"time"
	"io"
	"crypto/rand"
	"strconv"
	"encoding/base64"
	"net/http"
	"net/url"
)



/*session 会话管理*/
type SessionMgr struct {
	CookieName string //cookie名称
	Lock sync.Mutex  //保证goroutine安全
	MaxLifeTime int64  //最长生命周期
	Sessions map[string]*Session //保存session的指针，[sessionId]=session
}

/*session 会话*/
type Session struct {
	SessionID string
	LastTimeAccessed time.Time  //最后访问时间
	Values map[interface{}]interface{}  //其他，比如用户的权限
}

//初始化一个会话管理器
func NewSessionMgr(cookieName string, maxLifeTime int64) *SessionMgr{
	mgr := &SessionMgr{
		CookieName:cookieName,
		MaxLifeTime:maxLifeTime,
		Sessions:make(map[string]*Session),
	}

	//启动一个用实时监测session回收的函数
	go mgr.GC()
	return mgr
}

func (mgr *SessionMgr)GC(){
	mgr.Lock.Lock()
	defer mgr.Lock.Unlock()

	for sessionID, session:=range mgr.Sessions{
		//删除超过时限的session
		if session.LastTimeAccessed.Unix()+mgr.MaxLifeTime < time.Now().Unix(){
			delete(mgr.Sessions,sessionID)
		}
	}

	//定时器，循环执行该函数
	time.AfterFunc(time.Duration(mgr.MaxLifeTime)*time.Second, func() {mgr.GC()})
}

/*创建一个唯一的sessionID*/
func (mgr *SessionMgr)NewSessionID() string{
	b := make([]byte,64)
	if _,err := io.ReadFull(rand.Reader, b);err!=nil{
		nano := time.Now().UnixNano()
		return strconv.FormatInt(nano, 10)
	}
	return base64.URLEncoding.EncodeToString(b)
}

//session开始，只要是登录界面或者访问界面
func (mgr *SessionMgr)StartSession(w http.ResponseWriter,r *http.Request) string{
	mgr.Lock.Lock()
	defer mgr.Lock.Unlock()

	//无论原来是否存在session，都创建一个新的session
	newSessionID := url.QueryEscape(mgr.NewSessionID())

	session := &Session{
		SessionID:newSessionID,
		LastTimeAccessed:time.Now(),
		Values:make(map[interface{}]interface{}),
	}
	mgr.Sessions[newSessionID] = session

	//cookie设置过期时间
	cookie := http.Cookie{
		Name:mgr.CookieName,
		Value:newSessionID,
		Path:"/",
		HttpOnly:true,
		MaxAge:int(mgr.MaxLifeTime),
	}
	http.SetCookie(w, &cookie)
	return newSessionID

}

//会话结束
func (mgr *SessionMgr)EndSession(w http.ResponseWriter, r http.Request){
	cookie,err := r.Cookie(mgr.CookieName)
	if err!=nil||cookie.Value == ""{
		return
	} else{
		mgr.Lock.Lock()
		defer mgr.Lock.Unlock()
		delete(mgr.Sessions, cookie.Value)
		//最后让浏览器的cookie立即过期
		expire := time.Now()
		cookie := &http.Cookie{
			Name:mgr.CookieName,
			Path:"/",
			HttpOnly:true,
			Expires:expire,
			MaxAge:-1,
		}
		http.SetCookie(w, cookie)
	}
}

//用于删除重复的会话.sessionId 相同的
func (mgr *SessionMgr)EndSessionBy(sessionID string){
	mgr.Lock.Lock()
	defer mgr.Lock.Unlock()

	delete(mgr.Sessions, sessionID)
}

//获取session中的值
func (mgr *SessionMgr)GetSessionVal(sessionID string, key interface{})(interface{},bool){
	mgr.Lock.Lock()
	defer mgr.Lock.Unlock()

	if session, ok:= mgr.Sessions[sessionID];ok{
		if val,ok := session.Values[key];ok{
			return val,ok
		}
	}
	return nil,false
}

//设置session的值
func(mgr *SessionMgr)SetSessionVal(sessionID string, key interface{} ,value interface{}) bool{
	mgr.Lock.Lock()
	defer mgr.Lock.Unlock()

	if session,ok:= mgr.Sessions[sessionID];ok{
		session.Values[key] = value
		return true
	}
	return false
}

func (mgr *SessionMgr)GetSessionIDList() []string{
	mgr.Lock.Lock()
	defer mgr.Lock.Unlock()

	sessionIDList := make([]string,0)

	for k ,_:=range mgr.Sessions{
		sessionIDList = append(sessionIDList, k)
	}

	return sessionIDList
}

//session合法性判断
func (mgr *SessionMgr)CheckCookie(w http.ResponseWriter, r *http.Request)string{
	var cookie,err = r.Cookie(mgr.CookieName)
	if cookie==nil || err!=nil{
		return ""
	}

	mgr.Lock.Lock()
	defer mgr.Lock.Unlock()

	sessionID := cookie.Value
	if session,ok := mgr.Sessions[sessionID];ok{
		session.LastTimeAccessed = time.Now() //更新最后访问时间
		return sessionID
	}
	return ""
}

//获取最新的访问时间
func (mgr *SessionMgr)GetLastAccessTime(sessionID string) time.Time{
	mgr.Lock.Lock()
	defer mgr.Lock.Unlock()

	if session,ok := mgr.Sessions[sessionID];ok{
		return session.LastTimeAccessed
	}
	return time.Now()
}



