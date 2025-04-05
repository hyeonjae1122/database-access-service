package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", defaultHandler)

	fmt.Println("서버가 http://localhost:8080 에서 시작되었습니다.")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("서버 시작 오류:", err)
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "GET 메서드만 허용됩니다", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head><title>자기소개</title></head>
		<body>
			<h1>안녕하세요 👋</h1>
			<p>저는 Go 언어로 작성된 간단한 웹 애플리케이션입니다.</p>
		</body>
		</html>
	`)
}
