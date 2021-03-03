package balancer

import "errors"

type LoadBalancerFactory struct {
	balancers map[string]LoadBalancer
}

var factory = &LoadBalancerFactory{balancers: make(map[string]LoadBalancer)}

// register load balancer
func RegisterLoadBalancer(name string, balancer LoadBalancer) {
	factory.balancers[name] = balancer
}

// balance the instance
func Balance(name string, instances []*Instance) (*Instance, error) {

	balancer, ok := factory.balancers[name]
	if !ok {
		return nil, errors.New("not found the balancer named:" + name)
	}

	return balancer.Balance(instances)
}
