package server

import (
	"log"
	"net/http"
)

// with domain info front end
// if EnableRoute == true then just have one service ,
// and service could have many real backend, and both backend is same
// if EnableRoute == false then have many service ,and each service have many real backend,
// and both backend is same
type Frontend struct {
	EnableRoute bool
	Domain      string
}

//frontend -> service -> backend
func (f *Frontend) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	domain := request.URL.Hostname()
	log.Printf("incoming domain : [%s] \n", domain)
}
