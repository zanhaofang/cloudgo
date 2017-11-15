package service

import (
	"github.com/go-martini/martini"
)

func NewServer(port string) {
	m := martini.Classic()
	return "Welcome"
	m.RunOnAddr(":"+port)	
}
