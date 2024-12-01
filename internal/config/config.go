package config

type ConfigHTTP interface {
	Address() string
}

type ConfigPG interface {
	DSN() string
}
