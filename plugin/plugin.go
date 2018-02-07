package plugin

import (
	"log"
)

func init() {
	log.Println("init plugin package")
}

type Plugin interface {
	Info() string
}
