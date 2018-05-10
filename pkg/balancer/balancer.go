package balancer

import "github.com/oldtree/threeRashomon/pkg/backend"

type Balancer interface {
	Get([]backend.Backend) backend.Backend
}
