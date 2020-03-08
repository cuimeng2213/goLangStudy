package main

import (
	"fmt"
	"goLangStudy/cloudStore/codes/handler"
	"net/http"
)

func main() {
	//文件操作相关api
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSuccessHandler)
	http.HandleFunc("/file/meta", handler.GetFileQueryHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/update", handler.FileUpdateMetaHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)
	//
	http.HandleFunc("/user/signup", handler.SignupHandler)
	http.HandleFunc("/user/signin", handler.SigninHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("fail to start server: ", err)
		return
	}
}
