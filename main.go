package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	. "github.com/jangozw/simple-curd/database"
)

//---------------------go 原生http包结合gorm 进行API接口开发的简单示范----------------------------------
//
//

// user  表模型
type User struct {
	Id   int64
	Name string
}

// api 返回的json格式
type Response struct {
	Code int
	Msg  string
	Data interface{}
}

// 入口
func main() {
	defer Db.Close()
	fmt.Println("------------------------simple curd----------------------------------")
	fmt.Println("127.0.0.1:8008/user/list")
	fmt.Println("127.0.0.1:8008/user/add")
	fmt.Println("--------------")

	// 定义路由
	http.HandleFunc("/user/list", UserList)
	http.HandleFunc("/user/add", UserAdd)

	// 监听服务
	err := http.ListenAndServe("127.0.0.1:8008", nil)
	if err != nil {
		log.Fatal(err)
	}

	// 现在就可以请求了
	fmt.Println("http server started")
}

// 用户列表接口
func UserList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var users []User
	// 查询所有id>=1的用户
	Db.Where("id>=?", 1).Find(&users)
	resp := Response{
		Code: 0,
		Msg:  "success",
		Data: users,
	}
	respBytes, _ := json.Marshal(resp)
	// 返回json
	w.Write(respBytes)
}

// 添加用户 用户名取url 参数 name
func UserAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	args := r.URL.Query()
	fmt.Println(args)
	name := args.Get("name")
	insert := User{
		Name: name + "_test",
	}
	// 插入一条记录
	err := Db.Create(&insert).Error

	// 返回json
	if err != nil {
		resp := Response{
			Code: 400,
			Msg:  "fail",
			Data: "add user fail",
		}
		respBytes, _ := json.Marshal(resp)
		w.Write(respBytes)
	} else {
		resp := Response{
			Code: 0,
			Msg:  "success",
			Data: "add user succ",
		}
		respBytes, _ := json.Marshal(resp)
		w.Write(respBytes)
	}
}
