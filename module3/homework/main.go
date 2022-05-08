package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
)

const (
	contextPath string = "/"
	serverPort  string = ":80"
	VERSION     string = "VERSION"
	EMPTY       string = ""
	Version1000 string = "v1.0.0"
)

func main() {
	// init and add handler
	dispatcher := http.NewServeMux()
	dispatcher.HandleFunc("/", rootHandler)
	dispatcher.HandleFunc("/healthz", healthzHandler)
	// for pprof
	dispatcher.HandleFunc("/debug/pprof/", pprof.Index)
	dispatcher.HandleFunc("/debug/pprof/profile", pprof.Profile)
	dispatcher.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	dispatcher.HandleFunc("/debug/pprof/trace", pprof.Trace)
	// mock shutdown
	dispatcher.HandleFunc("/shutdown", shutdownHandler)

	// start http server
	err := http.ListenAndServe(serverPort, dispatcher)
	if err != nil {
		log.Fatal("starting http server is failed", err)
	}
}

// 作业1. 接收客户端 request，并将 request 中带的 header 写入 response header
// 作业2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
func rootHandler(w http.ResponseWriter, r *http.Request) {
	err := os.Setenv(VERSION, Version1000)
	if err != nil {
		log.Printf("Setting OS Env failed: %s\n", err.Error())
	}

	w.Header().Set("version", os.Getenv(VERSION)) // Header key自动驼峰处理
	// header from request
	header := r.Header
	for key, v := range header {
		for _, value := range v {
			// write to response header
			w.Header().Set(key, value)
		}
	}
	// write response body
	_, err = w.Write([]byte("hello world"))
	if err != nil {
		w.WriteHeader(500)
		io.WriteString(w, "server error\n")
	}
	printLog(w, r)
}

// 作业3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
func printLog(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	remote := getClientIp(r)
	method := r.Method
	// do logs
	log.Printf("[INFO] receive request from: %s, method: %s, url: %s\n",
		remote, method, url)
}

// 作业4. 当访问 localhost/healthz 时，应返回 200
func healthzHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK\n")
	printLog(w, r)
}

//TODO 优雅停止，实现此功能
func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(502)
	io.WriteString(w, "out of service\n")
}

// SUPPORT
// 解析可能存在的反向代理，比如nginx, haproxy
func getClientIp(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip != EMPTY {
		return ip
	}

	r.Header.Get("X-Real-Ip")
	if ip != EMPTY {
		return ip
	}

	if host, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		return host
	}

	return "failed to obtain host ip address!"
}
