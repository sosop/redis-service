package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"redis-service/config"
	"redis-service/models"
	"redis-service/redis"
	"strconv"
	"strings"
	"time"
)

func main() {
	host := config.GetString("host", true, "127.0.0.1")
	port := config.GetString("port", true, "8088")
	s := NewServer(fmt.Sprint(host, ":", port))
	http.HandleFunc("/do", do)
	s.ListenAndServe()
}

func do(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	cmd := r.FormValue("cmd")
	if strings.TrimSpace(cmd) == "" {
		outJson(w, models.NewParamRet())
		return
	}
	db := r.FormValue("db")
	dbInt := 0
	var err error
	if strings.TrimSpace(db) != "" {
		dbInt, err = strconv.Atoi(db)
		if err != nil {
			outJson(w, models.NewServerRet(err.Error()))
			return
		}
	}
	args := r.FormValue("args")
	argsSpice := strings.Split(args, ",")
	length := len(argsSpice)
	interArgs := make([]interface{}, length, length)
	for i, v := range argsSpice {
		interArgs[i] = v
	}
	reply, err := redis.Exec(dbInt, cmd, interArgs...)
	if err != nil {
		outJson(w, models.NewServerRet(err.Error()))
		return
	}
	outJson(w, models.NewSucRet(reply))
}

func NewServer(addr string) *http.Server {
	return &http.Server{
		Addr:           addr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func outJson(w http.ResponseWriter, ret models.Ret) {
	data, err := json.Marshal(ret)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(data)
}
