package conf

import "fmt"

type ConfigServer struct {
	Host string
	Port uint
}

func (conf ConfigServer) Addr() string {
	return fmt.Sprintf("%s:%d", conf.Host, conf.Port)
}
