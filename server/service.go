package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Service struct {
	Name    string
	Backend []*Backend
	router  *httprouter.Router
}

//real service incoming,after domain dispatch to here
func (s *Service) ServeHTTP(response http.ResponseWriter, requets http.Request) {

}
