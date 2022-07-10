package main

type Server struct {
	*router
	middlewares 	[]Middleware
	startHandler	HandlerFunc
}

func NewServer() *Server {
	r := &router{make(map[string]map[string]HandlerFunc)}
	s := &Server{router: r}
	s.middlewares = []Middleware{
		logHandler,
		recoverHandler,
		staticHandler,
		parseFormHandler,
		parseJsonBodyHandler,
	}
	// TODO : 포인터로 받아야하지 않나??
	return s
}

// Middleware chain 생성