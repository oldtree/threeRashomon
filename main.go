package main

import (
	"os"
	"os/signal"
	"syscall"

	_ "github.com/oldtree/threeRashomon/config"
	_ "github.com/oldtree/threeRashomon/description"
	_ "github.com/oldtree/threeRashomon/discover"
	_ "github.com/oldtree/threeRashomon/plugin"
	_ "github.com/oldtree/threeRashomon/server"
	_ "github.com/oldtree/threeRashomon/stub"
	_ "github.com/oldtree/threeRashomon/swagger"
)

func main() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-sc
}
