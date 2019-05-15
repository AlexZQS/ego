package user

import (
	"encoding/json"
	"net/http"
)

/**
 * @description
 * @time 2019/5/15 22:17
 * @version
 */

func UserHandler() {
	http.HandleFunc("/login", loginController)
}

//登录
func loginController(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	result := LoginService(username, password)

	//把结构体转换为json数据
	bytes, _ := json.Marshal(result)
	//设置响应内容为json
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(bytes)
}
