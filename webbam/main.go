package main

import (
	"fmt"
)

func main() {
	// 서버 생성
	s := NewServer()

	s.HandleFunc("GET", "/", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "welcome!")
	})

	s.HandleFunc("GET", "/about", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "about")
	})

	s.HandleFunc("GET", "/users/:id", func(c *Context) {
		if c.Params["id"] == "0" {
			panic("id is zero")
		}
		fmt.Fprintf(c.ResponseWriter, "retrieve user %v\n", c.Params["id"])
	})

	s.HandleFunc("GET", "/users/:user_id/addresses/:address_id", func(c *Context) {
		fmt.Fprintf(c.ResponseWriter, "retrieve user's address %v's address %v\n", c.Params["user_id"], c.Params["address_id"])
	})

	s.HandleFunc("POST", "/users", (func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "create user")
	}))

	s.HandleFunc("POST", "/users/:user_id/addresses", func(c *Context) {
		fmt.Fprintf(c.ResponseWriter, "create user %v's address\n", c.Params["user_id"])
	})

	s.Run(":8080")
}
