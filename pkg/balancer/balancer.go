package balancer

//  load balancer
type LoadBalancer interface {
	Balance(instances []*Instance) (*Instance, error)
}
