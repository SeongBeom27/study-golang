package main

import "net/http"

// 컨텍스트 타입 정의, 핸들러 내부에는 Context 값이 전달되게함
type Context struct {
	Params map[string]interface{} // 라우터에서 해석한 URL 매개변수를 담음

	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

// 핸들러 타입 정의
type HandlerFunc func(*Context)
