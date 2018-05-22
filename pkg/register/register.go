package register

import (
	"log"
)

type Register interface {
	Register()
	UnRegister()
}

type EtcdRegister struct {
	Name        string
	EtcdAddress string
	MaxRetry    int
}

func (e *EtcdRegister) Register() {
	log.Printf("register service %s start \n", e.Name)
}

func (e *EtcdRegister) UnRegister() {
	log.Println("UnRegister service %s end \n", e.Name)
}
