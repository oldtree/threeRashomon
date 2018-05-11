package serviceGate

import (
	"fmt"

	"sync"

	"github.com/oldtree/threeRashomon/pkg/backend"
	"github.com/oldtree/threeRashomon/pkg/balancer"
)

const (
	SERVICE_STATUS_INIT = iota
	SERVICE_STATUS_NOT_USABLE
	SERVICE_STATUS_USABLE
)

type ServiceGate struct {
	ServiceStatus int
	ServiceName   string

	ProtactServiceBackend sync.RWMutex
	ServiceBackend        []backend.Backend

	ProtactServiceBalance sync.RWMutex
	ServiceBalanceJudge   balancer.Balancer
}

func NewServiceGate(serivceName string, bla balancer.Balancer, backends []backend.Backend) *ServiceGate {
	if bla == nil {
		bla = nil //use default balance
	}
	newService := &ServiceGate{
		ServiceName:         serivceName,
		ServiceBackend:      backends,
		ServiceBalanceJudge: bla,
	}
	if len(backends) <= 0 {
		newService.ServiceStatus = SERVICE_STATUS_NOT_USABLE
	}
	return newService
}

func (service *ServiceGate) AddBackend(backends []backend.Backend) error {
	if len(backends) <= 0 {
		return fmt.Errorf("service add backend failed : %s", "backends is empty")
	}
	service.ProtactServiceBackend.Lock()
	defer service.ProtactServiceBackend.Unlock()
	service.ServiceBackend = append(service.ServiceBackend, backends...)
	return nil
}

func (service *ServiceGate) ModifyBalance(bla balancer.Balancer) error {
	if bla == nil {
		return fmt.Errorf("service modity balance failed : ", "balance params is nil")
	}
	service.ProtactServiceBalance.Lock()
	defer service.ProtactServiceBalance.Unlock()
	service.ServiceBalanceJudge = bla
	return nil
}
