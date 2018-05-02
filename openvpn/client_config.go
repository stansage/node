package openvpn

import "github.com/mysterium/node/openvpn/config"

type ClientConfig struct {
	*config.Config
}

func (c *ClientConfig) SetClientMode(serverIP string, serverPort int) {
	c.SetFlag("client")
	c.SetParam("script-security", "2")
	c.SetFlag("auth-nocache")
	c.SetParam("remote", serverIP)
	c.SetPort(serverPort)
	c.SetFlag("nobind")
	c.SetParam("remote-cert-tls", "server")
	c.SetFlag("auth-user-pass")
	c.SetFlag("management-query-passwords")
}

func (c *ClientConfig) SetProtocol(protocol string) {
	if protocol == "tcp" {
		c.SetParam("proto", "tcp-client")
	}
}
