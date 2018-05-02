package openvpn

import "github.com/mysterium/node/openvpn/config"

// ServerConfig represents openvpn as server configuration
type ServerConfig struct {
	*config.Config
}

// SetServerMode sets required configuration parameters to act as server
func (c *ServerConfig) SetServerMode(port int, network, netmask string) {
	c.SetPort(port)
	c.SetParam("server", network+" "+netmask)
	c.SetParam("topology", "subnet")
}

// SetTLSServer sets required configuration params for tls-server
func (c *ServerConfig) SetTLSServer() {
	c.SetFlag("tls-server")
	c.AddOptions(config.OptionFile("dh", "none"))
}

// SetProtocol sets specified protocol options
func (c *ServerConfig) SetProtocol(protocol string) {
	if protocol == "tcp" {
		c.SetParam("proto", "tcp-server")
	}
}
