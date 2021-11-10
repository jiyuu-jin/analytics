package main

import (
	"context"
	"net/http"
)

type Config struct{}

func CreateConfig() *Config {
	return &Config{}
}

type Analytics struct {
	next http.Handler
	name string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Analytics{
		next: next,
		name: name,
	}, nil
}

func (u *Analytics) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Hello\n"))
}
