package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"

	"github.com/sevlyar/go-daemon"
)

// 要终止守护进程，使用:
// kill `cat sample.pid`。
func main() {

	context := &daemon.Context{
		PidFileName: "sample.pid",
		PidFilePerm: 0644,
		LogFileName: "sample.log",
		LogFilePerm: 0640,
		WorkDir:     "./",                        //创建子进程时会切到该目录
		Umask:       027,                         //If Umask is non-zero, the daemon-process call Umask() func with given value.
		Args:        []string{"go-daemon守护进程服务"}, //传递给子进程的参数 是nil就用os.Args
	}

	d, err := context.Reborn() //拷贝上下文创建子进程 在父进程返回*os.Process 在子进程程返回的是nil 其他情况返回错误
	if err != nil {
		log.Fatal("守护进程应该已经存在了, 创建守护进程失败, err:", err.Error())
	}
	if d != nil {
		log.Printf("这是在父进程的标志")
		return
	}

	defer func(context *daemon.Context) {
		err := context.Release()
		if err != nil {
			log.Printf("释放失败:%s", err.Error())
		}
		log.Printf("释放成功!!!")
	}(context)

	log.Print("守护进程启动!!!!!")
	serveHTTP()
	log.Print("子进程服务退出了~~~")
}

func serveHTTP() {
	log.Printf("参数:%v", os.Args)
	http.HandleFunc("/", httpHandler)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Printf("启动服务失败, err:%s", err.Error())
		return
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("request from %s: %s %q", r.RemoteAddr, r.Method, r.URL)
	_, err := fmt.Fprintf(w, "hello 你好啊!: %q", html.EscapeString(r.URL.Path))
	if err != nil {
		log.Printf("返回信息失败, err:%s", err)
		return
	}
}
