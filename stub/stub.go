package stub

import (
	"log"
	"time"

	"github.com/julienschmidt/httprouter"
)

func init() {

}

type Handler = httprouter.Handle

type RouteinfoNotifyRegister func()
type RouteinfoListenRegister func()

//route info from servivce discover client side
type RouteInfoMessage struct {
}

//可以注册多个被通知的服务，通知他们变更路由信息，是否进行变更在这里进行判断
//注册多个用于服务发现的通知函数，这样，可以做到多来源，保证服务发现的高可用性
//比如同时兼容ectd,kubenetes,consul....etc
type NightWatcher struct {
	TickCircle time.Ticker
}

func (n *NightWatcher) RegistListener(rnr RouteinfoNotifyRegister) {

}

func (n *NightWatcher) RegisterNotifier(rlr RouteinfoListenRegister) {

}

func (n *NightWatcher) RunInCircle() {
	for {
		select {
		case <-n.TickCircle.C:
			log.Printf("update tick timestamp : [%s] \n", time.Now())
		}
	}
}
