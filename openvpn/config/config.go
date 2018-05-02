package config

import (
	"strconv"
)

func NewConfig() *Config {
	return &Config{
		options: make([]configOption, 0),
	}
}

type Config struct {
	options []configOption
}

type configOption interface {
	getName() string
}

func (c *Config) AddOptions(options ...configOption) {
	c.options = append(c.options, options...)
}

func (c *Config) SetParam(name, value string) {
	c.AddOptions(
		OptionParam(name, value),
	)
}

func (c *Config) SetFlag(name string) {
	c.AddOptions(
		OptionFlag(name),
	)
}

func (c *Config) SetManagementSocket(socketAddress string) {
	c.SetParam("management", socketAddress+" unix")
	c.SetFlag("management-client")
}

func (c *Config) SetPort(port int) {
	c.SetParam("port", strconv.Itoa(port))
}

func (c *Config) SetDevice(deviceName string) {
	c.SetParam("dev", deviceName)
}

func (c *Config) SetTLSCACertificate(caFile string) {
	c.AddOptions(OptionFile("ca", caFile))
}

func (c *Config) SetTLSPrivatePubKeys(certFile string, certKeyFile string) {
	c.AddOptions(OptionFile("cert", certFile))
	c.AddOptions(OptionFile("key", certKeyFile))
}

func (c *Config) SetTLSCrypt(cryptFile string) {
	c.AddOptions(OptionFile("tls-crypt", cryptFile))
}

// RestrictReconnects describes conditions which enforces client to close a session in case of failed authentication
func (c *Config) RestrictReconnects() {
	c.SetParam("connect-retry-max", "2")
	c.SetParam("remap-usr1", "SIGTERM")
	c.SetFlag("single-session")
	c.SetFlag("tls-exit")
}

func (c *Config) SetKeepAlive(interval, timeout int) {
	c.SetParam("keepalive", strconv.Itoa(interval)+" "+strconv.Itoa(timeout))
}

func (c *Config) SetPingTimerRemote() {
	c.SetFlag("ping-timer-rem")
}

func (c *Config) SetPersistTun() {
	c.SetFlag("persist-tun")
}

func (c *Config) SetPersistKey() {
	c.SetFlag("persist-key")
}
