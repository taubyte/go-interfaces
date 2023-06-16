package network

type Shape struct {
	Protocols []string `yaml:",omitempty"`
	Ports     Ports    `yaml:",omitempty"`
}

type Ports struct {
	P2P map[string]int `yaml:"p2p"`
}
