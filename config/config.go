package config

import (
	"encoding/json"
)

func init() {

}

type superConfig interface {
	Format() []byte
}

type Config struct {
	Version string
	AppName string

	DiscoverConfig []superConfig
}

type EtcdConfig struct {
	Address string
	RootKey string
}

func (e *EtcdConfig) Format() []byte {
	data, _ := json.Marshal(e)
	return data
}

type KubenetesConfig struct {
	ApiAddress string
	InCluster  bool
}

func (k *KubenetesConfig) Format() []byte {
	data, _ := json.Marshal(k)
	return data
}

type BoltdbConfig struct {
	Address string
}

func (b *BoltdbConfig) Format() []byte {
	data, _ := json.Marshal(b)
	return data
}
