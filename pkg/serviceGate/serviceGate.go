package serviceGate

import "github.com/oldtree/threeRashomon/pkg/backend"
import "github.com/oldtree/threeRashomon/pkg/balancer"

type ServiceGate struct {
	ServiceName         string
	ServiceBackend      []backend.Backend
	ServiceBalanceJudge balancer.Balancer
}
