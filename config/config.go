package config

type Config struct {
	Version      string `json:"version,omitempty"`
	ProtocolFile string `json:"protocol_file,omitempty"`
	HttpPort     string `json:"http_port,omitempty"`
	RpcPort      string `json:"rpc_port,omitempty"`
	Debug        bool   `json:"debug,omitempty"`
	PprofPort    string `json:"pprof_port,omitempty"`
}
