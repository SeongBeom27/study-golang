package main

import "net/http"

type Server struct {
	*router
	middlewares  []Middleware
	startHandler HandlerFunc
}

func NewServer() *Server {
	r := &router{make(map[string]map[string]HandlerFunc)}
	s := &Server{router: r}
	s.middlewares = []Middleware{
		logHandler,
		recoverHandler,
		staticHandler,
		parseFormHandler,
		parseJsonBodyHandler}

	return s
}

func (s *Server) Run(addr string) {
	// staartHandler를 라우터 핸들러 함수로 지정
	s.startHandler = s.router.handler()

	// 등록된 미들웨어를 라우터 핸들러 앞에 하나씩 추가
	for i := len(s.middlewares) - 1; i >= 0; i-- {
		s.startHandler = s.middlewares[i](s.startHandler)
	}

	// 웹 서버 시작
	if err := http.ListenAndServe(addr, s); err != nil {
		panic(err)
	}
}

// Server가 웹 요청을 받아 처리할 수 있도록 ServeHTTP 메서드 정의
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Context 생성
	c := &Context{
		Params:         make(map[string]interface{}),
		ResponseWriter: w,
		Request:        r,
	}

	for k, v := range r.URL.Query() {
		c.Params[k] = v[0]
	}

	s.startHandler(c)
}
