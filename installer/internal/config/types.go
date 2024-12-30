package config

import (
	"net"
)

type JumpHost struct {
	IP       net.IP `yaml:"ip"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
