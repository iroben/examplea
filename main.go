package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

const (
	// 从配置服务中取
	PROJECT_B_DOMAIN = "http://exampleb:8000"
)

func main() {
	http.HandleFunc("/a/api", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		resp, _ := http.Get(PROJECT_B_DOMAIN + "/b/api")
		defer resp.Body.Close()
		result, _ := ioutil.ReadAll(resp.Body)
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Write([]byte("debug project a<br/>"))
		writer.Write(result)
	})
	http.ListenAndServe(":8000", nil)
}
