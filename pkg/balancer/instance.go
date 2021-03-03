package balancer

import "fmt"

type Instance struct {
	Ip   string
	Port int
}

func NewInstance(ip string, port int) *Instance {

	return &Instance{}
}

func (s Instance) String() string {

	return fmt.Sprintf("ip:%s,port:%d", s.Ip, s.Port)
}
