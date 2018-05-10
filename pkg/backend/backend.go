package backend

import "net/http"

type Backend interface {
	Do(req *http.Request) (http.Response, error)
	Alive() (bool, error)
}
