package server

import (
	"net/http"
	"os"

	"crypto/tls"
	"net"
	"sync"

	"github.com/oldtree/threeRashomon/stub"
)

func init() {

}

var _ = stub.RouteInfoMessage{}

type Server struct {
	HttpServer *http.Server
	Listener   net.Listener

	StopChan chan struct{}
	Signal   chan os.Signal

	UseTls    bool
	TlsConfig *tls.Config

	//GroutinePool
	//Metricinfo
	domainMap *sync.Map
}

func (s *Server) AddNewDomain(domain string, enbaleRoute bool) *Frontend {
	newFrontend := &Frontend{
		EnableRoute: enbaleRoute,
		Domain:      domain,
	}
	s.domainMap.Store(domain, newFrontend)
	return newFrontend
}

func (s *Server) RemoveDomain(domain string) *Frontend {
	needRemoveFrontend, ok := s.domainMap.Load(domain)
	if !ok {
		return nil
	}
	s.domainMap.Delete(domain)
	return needRemoveFrontend.(*Frontend)
}

func (s *Server) UpdateDomain(domain string, enbaleRoute bool) *Frontend {
	s.domainMap.Delete(domain)
	return s.AddNewDomain(domain, enbaleRoute)
}

func createListener(address string) (net.Listener, error) {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	return listen, nil
}

func (s *Server) Start() error {
	err := s.HttpServer.Serve(s.Listener)
	return err
}

func (s *Server) Close() error {
	err := s.HttpServer.Close()
	return err
}

func (s *Server) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	domainName := request.URL.Hostname()
	if value, ok := s.domainMap.Load(domainName); ok && value != nil {
		value.(*Frontend).ServeHTTP(response, request)
		return
	}
	http.NotFound(response, request)
	return
}
