package route

import (
	"github.com/julienschmidt/httprouter"
)

type RouteInfo struct {
	BelongToService string
	route           *httprouter.Router
}

func NewRoute(serviceName string) *RouteInfo {
	newRoute := &RouteInfo{
		BelongToService: serviceName,
		route:           httprouter.New(),
	}
	return newRoute
}
