// 代码生成时间: 2025-08-28 11:50:49
package main

import (
    "net/http"
    "strings"
    "log"
    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"
)

// Store 是用于会话管理的存储对象
var store = sessions.NewCookieStore([]byte("secret-key"))

// User 结构体用于表示用户信息
type User struct {
    Username string
    Password string
}

// AuthHandler 是身份认证的处理函数
func AuthHandler(w http.ResponseWriter, r *http.Request) {
    var user User

    // 从表单中获取用户名和密码
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Error parsing form", http.StatusInternalServerError)
        return
    }

    username := r.Form.Get("username")
    password := r.Form.Get("password")

    // 检查用户名和密码是否为空
    if username == "" || password == "" {
        http.Error(w, "Username or password cannot be empty", http.StatusUnauthorized)
        return
    }

    // 这里应该添加对用户名和密码的验证逻辑，例如查询数据库
    // 为了简化示例，我们假设所有用户都是有效的
    user.Username = username
    user.Password = password

    // 创建一个新的会话，并设置用户名
    session, err := store.Get(r, "auth-session")
    if err != nil {
        http.Error(w, "Error creating session", http.StatusInternalServerError)
        return
    }
    session.Values["user"] = user
    err = session.Save(r, w)
    if err != nil {
        http.Error(w, "Error saving session", http.StatusInternalServerError)
        return
    }

    // 重定向到主页
    http.Redirect(w, r, "/", http.StatusFound)
}

// LogoutHandler 是用户注销的处理函数
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    // 从会话中删除用户信息
    session, err := store.Get(r, "auth-session")
    if err != nil {
        http.Error(w, "Error getting session", http.StatusInternalServerError)
        return
    }
    delete(session.Values, "user")
    err = session.Save(r, w)
    if err != nil {
        http.Error(w, "Error saving session", http.StatusInternalServerError)
        return
    }

    // 重定向到登录页面
    http.Redirect(w, r, "/login", http.StatusFound)
}

// main 函数是程序的入口点
func main() {
    r := mux.NewRouter()

    // 设置静态文件服务
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // 设置身份认证和注销的路由
    r.HandleFunc("/login", AuthHandler).Methods("POST")
    r.HandleFunc("/logout", LogoutHandler).Methods("GET")

    // 设置主页路由
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        session, err := store.Get(r, "auth-session")
        if err != nil {
            http.Error(w, "Error getting session", http.StatusInternalServerError)
            return
        }

        if session.IsNew {
            http.Redirect(w, r, "/login", http.StatusFound)
            return
        }

        user, ok := session.Values["user"].(User)
        if !ok {
            http.Redirect(w, r, "/login", http.StatusFound)
            return
        }

        // 显示主页
        w.Write([]byte("Welcome, " + user.Username + "!"))
    })

    // 启动HTTP服务器
    log.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", r)
    if err != nil {
        log.Fatal("Error starting server: ", err)
    }
}