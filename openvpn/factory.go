package openvpn

import (
	"github.com/mysterium/node/openvpn/config"
	"github.com/mysterium/node/openvpn/primitives"
	"io/ioutil"
)

// NewServerConfig creates openvpn configuration options preset for openvpn as server
func NewServerConfig(
	network, netmask string,
	secPrimitives *primitives.SecurityPrimitives,
	protocol string,
) *ServerConfig {
	serverConfig := ServerConfig{config.NewConfig()}
	serverConfig.SetServerMode(1194, network, netmask)
	serverConfig.SetTLSServer()
	serverConfig.SetProtocol(protocol)
	serverConfig.SetTLSCACertificate(secPrimitives.CACertPath)
	serverConfig.SetTLSPrivatePubKeys(secPrimitives.ServerCertPath, secPrimitives.ServerKeyPath)
	serverConfig.SetTLSCrypt(secPrimitives.TLSCryptKeyPath)

	serverConfig.SetDevice("tun")
	serverConfig.SetParam("cipher", "AES-256-GCM")
	serverConfig.SetParam("verb", "3")
	serverConfig.SetParam("tls-version-min", "1.2")
	serverConfig.SetFlag("management-client-auth")
	serverConfig.SetParam("verify-client-cert", "none")
	serverConfig.SetParam("tls-cipher", "TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384")
	serverConfig.SetParam("reneg-sec", "60")
	serverConfig.SetKeepAlive(10, 60)
	serverConfig.SetPingTimerRemote()
	serverConfig.SetPersistTun()
	serverConfig.SetPersistKey()
	serverConfig.SetFlag("explicit-exit-notify")

	return &serverConfig
}

// NewClientConfig creates openvpn configuration options preset for openvpn as client
func NewClientConfig(
	remote string,
	caCertPath, tlsCryptKeyPath string,
	protocol string,
) *ClientConfig {
	clientConfig := ClientConfig{config.NewConfig()}
	clientConfig.SetClientMode(remote, 1194)
	clientConfig.SetProtocol(protocol)
	clientConfig.SetTLSCACertificate(caCertPath)
	clientConfig.SetTLSCrypt(tlsCryptKeyPath)
	clientConfig.RestrictReconnects()

	clientConfig.SetDevice("tun")
	clientConfig.SetParam("cipher", "AES-256-GCM")
	clientConfig.SetParam("verb", "3")
	clientConfig.SetParam("tls-cipher", "TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384")
	clientConfig.SetKeepAlive(10, 60)
	clientConfig.SetPingTimerRemote()
	clientConfig.SetPersistTun()
	clientConfig.SetPersistKey()

	clientConfig.SetParam("reneg-sec", "60")
	clientConfig.SetParam("resolv-retry", "infinite")
	clientConfig.SetParam("redirect-gateway", "def1 bypass-dhcp")
	clientConfig.SetParam("dhcp-option", "DNS 208.67.222.222")
	clientConfig.SetParam("dhcp-option", "DNS 208.67.220.220")
	clientConfig.SetFlag("explicit-exit-notify")

	return &clientConfig
}

// NewClientConfigFromString parses given string and creates configuration options for openvpn as client
func NewClientConfigFromString(
	configString, configFile string,
	scriptUp, scriptDown string,
) (*ClientConfig, error) {
	err := ioutil.WriteFile(configFile, []byte(configString), 0600)
	if err != nil {
		return nil, err
	}

	clientConfig := ClientConfig{config.NewConfig()}
	clientConfig.AddOptions(config.OptionParam("clientConfig", configFile))

	clientConfig.SetParam("up", scriptUp)
	clientConfig.SetParam("down", scriptDown)

	return &clientConfig, nil
}
