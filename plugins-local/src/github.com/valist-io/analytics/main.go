package analytics

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type Config struct{}

func CreateConfig() *Config {
	return &Config{}
}

type Analytics struct {
	next http.Handler
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Analytics{
		next: next,
	}, nil
}

func (u *Analytics) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	cid := strings.Split(fmt.Sprintf("%s", req.URL), "/")
	fmt.Println(cid[2])
	rw.Write([]byte("Hello\n"))
}
