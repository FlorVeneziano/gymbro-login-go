package envs

import "sync"

type env struct {
	PORT  string
	LOCAL bool

	PEPPER string
	ENV    string

	MONGO_HOST     string
	MONGO_DATABASE string
}

var envs *env
var once sync.Once
