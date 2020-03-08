package handler

import (
	"fmt"
	mydb "goLangStudy/cloudStore/codes/db"
	"goLangStudy/cloudStore/codes/util"
	"io/ioutil"
	"net/http"
)

const (
	pwd_salt = "8-*0"
)

// SignupHandler
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
	} else if r.Method == "POST" {
		r.ParseForm()
		username := r.Form.Get("username")
		passwd := r.Form.Get("password")
		if len(username) <= 3 || len(passwd) < 5 {
			w.Write([]byte("Invalid param!!!"))
			return
		}
		enc_passwd := util.Sha1([]byte(pwd_salt + passwd))
		ret := mydb.UserSignup(username, enc_passwd)
		if ret != true {
			fmt.Printf("注册失败 %s \n", username)
			return
		}
	}
}
func GenToken(username) string {
	//40位： username + current + token_salt
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + "_tokensalt"))
	return tokenPrefix + ts[:8]
}
func SiginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	encPasswd := util.Sha1([]byte(password + pwd_salt))
	//1.校验用户名及密码
	pwdChecked := mydb.UserSignin(username, encPasswd)
	if !pwdChecked {
		w.Write([]byte("login failed"))
		return
	}
	//2.生成访问凭证
	token := GenToken(username)
	upRet := mydb.UpdateToken(username, token)
	if !upRet {
		w.Write([]byte("update token faiked"))
		return
	}
	//3.重定向到首页
	w.Write([]byte("http://" + r.Host + "/static/view/home.html"))
}
