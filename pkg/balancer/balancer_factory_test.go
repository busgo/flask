package balancer

import (
	"fmt"
	"testing"
)

func TestRandomLoadBalancer(t *testing.T) {

	instances := make([]*Instance, 0)

	port := 8080
	for i := 0; i < 10; i++ {

		ip := fmt.Sprintf("192.168.1.%d", i)
		instances = append(instances, &Instance{Ip: ip, Port: port})
	}

	for i := 0; i < 1000; i++ {

		instance, err := Balance("random_balancer", instances)
		if err != nil {
			t.Error(err)
		}

		t.Log(instance.String())
	}

}

func TestRoundRobinLoadBalancer(t *testing.T) {

	instances := make([]*Instance, 0)

	port := 8080
	for i := 0; i < 10; i++ {

		ip := fmt.Sprintf("192.168.1.%d", i)
		instances = append(instances, &Instance{Ip: ip, Port: port})
	}

	for i := 0; i < 1000; i++ {

		instance, err := Balance("round_robin_balancer", instances)
		if err != nil {
			t.Error(err)
		}

		t.Log(instance.String())
	}

}
