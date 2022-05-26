package constants

import "net"

const (
	Host = "smtp.mail.yahoo.com"
	Port = "465"
)

var HostAndPort = net.JoinHostPort(Host, Port)