package common

import (
	seerIface "github.com/taubyte/go-interfaces/services/seer"
	http "github.com/taubyte/http"
	"github.com/taubyte/p2p/peer"
)

type GenericConfig struct {
	Shape        string
	Node         peer.Node
	Http         http.Service
	ClientNode   peer.Node
	DVPrivateKey []byte
	DVPublicKey  []byte

	Root string

	Protocols []string `yaml:"protocols"`

	PrivateKey []byte `yaml:"privatekey"`
	SwarmKey   []byte `yaml:"swarmkey"`

	Ports map[string]int

	P2PListen   []string `yaml:"p2p-listen"`
	P2PAnnounce []string `yaml:"p2p-announce"`

	Location *seerIface.Location `yaml:"location"`

	Bootstrap bool
	Peers     []string `yaml:"peers"`
	DevMode   bool     `yaml:"devmode"`

	HttpListen string `yaml:"http-listen"`
	HttpSecure bool   `yaml:"http-secure"`
	NetworkUrl string `yaml:"network-url"`
	DnsPort    int    `yaml:"dns"`

	Verbose bool   `yaml:"verbose"`
	Branch  string `yaml:"branch"`

	TLS     TLSConfig      `yaml:"tls"`
	Domains HttpDomainInfo `yaml:"domains"`

	Identity Identity
	Plugins  []string
}

type Identity struct {
	Client GithubOauth
}

type GithubOauth struct {
	Id     string
	Secret string
}
type TLSConfig struct {
	Certificate string `yaml:"certificate"`
	Key         string `yaml:"key"`
}

type HttpDomainInfo struct {
	Key         DVKey              `yaml:"key"`
	Whitelisted WhiteListedDomains `yaml:"whitelist"`
	Services    string             `yaml:"services"`
	Generated   string             `yaml:"generated"`
}

// TODO: combine with the structure inside mycelium
type WhiteListedDomains struct {
	Postfix []string
	Regex   []string
}

type DVKey struct {
	Private string `yaml:"private"`
	Public  string `yaml:"public"`
}
