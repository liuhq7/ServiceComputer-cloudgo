# ServiceComputer-cloudgo
## 1.关于cloudgo
cloudgo小程序使用了三个框架mux，negroni，render。并且可以通过命令行参数-p来指定监听的端口号。</br>
具体可以参考老师的博客：https://blog.csdn.net/pmlpml/article/details/78539261</br>
## 2.代码实现
### 2.1.服务器端代码
```go
package service

import (
	"net/http"
	"github.com/mux"
	"github.com/negroni"
	"github.com/render"
)

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/welcome/{id}", testHandler(formatter)).Methods("GET")
}

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

func testHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Welcome " + id})
	}
}
```
### 2.2.客户端代码
```go
package main

import (
	"os"
	"cloudgo/service"
	flag "github.com/pflag"
)

const (
	PORT string = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	server := service.NewServer()
	server.Run(":" + port)
}
```
## 3.curl工具
### 3.1.用crul工具访问web

