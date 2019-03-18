package main

import (
	"os/exec"
	"fmt"
	"io/ioutil"
)

func main(){
	ipCmd := exec.Command("whoami")

	ipOut,err := ipCmd.Output()
	if err!=nil{
		panic(err)
	}

	fmt.Println("> whoami")
	fmt.Println(string(ipOut))

	grepCmd := exec.Command("findstr","package")
	grepIn,_:=grepCmd.StdinPipe()
	grepOut,_:= grepCmd.StdoutPipe()

	grepCmd.Start()
	grepIn.Write([]byte("hellp grep"))
	grepIn.Close()

	grepBytes,_ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep package")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c","ls -a -l -h")
	lsOut,err := lsCmd.Output()
	if err!=nil{
		panic(err)
	}

	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
