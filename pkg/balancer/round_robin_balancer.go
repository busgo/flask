package balancer

import "errors"

// round robin balancer
type RoundRobinBalancer struct {
	current int
}

func init() {
	RegisterLoadBalancer("round_robin_balancer", &RoundRobinBalancer{})
}

func (lb *RoundRobinBalancer) Balance(instances []*Instance) (*Instance, error) {
	if len(instances) == 0 {
		return nil, errors.New("has no instance")
	}

	pos := lb.current
	if pos >= len(instances) {
		pos = 0
	}

	instance := instances[pos]
	lb.current = pos + 1
	return instance, nil

}
