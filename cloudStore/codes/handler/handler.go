package handler

import (
	"encoding/json"
	"fmt"
	"goLangStudy/cloudStore/codes/meta"
	"goLangStudy/cloudStore/codes/util"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//UploadHandler 处理文件上传
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//返回上传html页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "Internal error")
			return
		}
		io.WriteString(w, string(data))
	} else {
		//接受文件流以及存储到本地目录
		file, head, err := r.FormFile("file") // 返回三个句柄：
		defer file.Close()
		if err != nil {
			fmt.Printf("fail to get data: %v \n", err)
			return
		}
		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: "./tmp/" + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		newFile, err := os.Create(fileMeta.Location)
		defer newFile.Close()
		if err != nil {
			fmt.Printf("fail to create file: %v \n", err)
			return
		}
		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("fail to copy file: %v \n", err)
			return
		}
		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		//保存meta信息到map中
		// meta.UpdateFileMeta(fileMeta)
		//将元信息保存到数据库
		meta.UpdateFileMetaDB(fileMeta)

		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

//上传完成
func UploadSuccessHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "upload success")
}

//获取单个文件元信息
func GetFileQueryHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filehash := r.Form.Get("filehash")
	//从缓存中获取对应文件的数据
	// fmeta := meta.GetFileMeta(filehash)
	fmeta := meta.GetFileMetaDB(filehash)
	data, err := json.Marshal(fmeta)
	if err != nil {
		fmt.Printf("failed to Marshal: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// 下载文件
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filehash := r.Form.Get("filehash")
	//获取对应文件信息
	fm := meta.GetFileMeta(filehash)
	fp, err := os.Open(fm.Location)
	defer fp.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := ioutil.ReadAll(fp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//修改header 使客户端能够识别下载文件
	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("Content-Descrption", "attachment/filename=\""+fm.FileName+"\"")
	w.Write(data)
}

//修改文件元信息（重命名）
func FileUpdateMetaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	file_hash := r.Form.Get("filehash")
	new_filename := r.Form.Get("filename")
	//文件重命名操作
	opType := r.Form.Get("op")
	if opType != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	curFileMeta := meta.GetFileMeta(file_hash)
	if err := os.Rename(curFileMeta.Location, "./tmp/"+new_filename); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	curFileMeta.FileName = new_filename
	curFileMeta.Location = "./tmp/" + new_filename
	// meta.UpdateFileMeta(curFileMeta)
	meta.UpdateFileMetaDB(curFileMeta)

	data, err := json.Marshal(curFileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

//文件删除接口
func FileDeleteHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	file_hash := r.Form.Get("filehash")
	//删除文件
	fm := meta.GetFileMeta(file_hash)
	os.Remove(fm.Location)
	//删除文件元信息
	meta.RemoveFileMeta(file_hash)

	w.WriteHeader(http.StatusOK)
}
