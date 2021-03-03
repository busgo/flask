package balancer

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type RandomLoadBalancer struct {
}

func init() {

	RegisterLoadBalancer("random_balancer", &RandomLoadBalancer{})
}

func (lb *RandomLoadBalancer) Balance(instances []*Instance) (*Instance, error) {

	if len(instances) == 0 {
		return nil, errors.New("has not instance")
	}

	rand.Seed(time.Now().UnixNano())
	pos := rand.Intn(len(instances))
	fmt.Println("pos", pos)
	if pos >= len(instances) {
		pos = 0
	}

	return instances[pos], nil
}
