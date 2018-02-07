package cmd

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var configpath = flag.String("-c", "cfg.json", "config file")

func main() {
	log.Printf("config path : [%s] \n", *configpath)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-sc
}
