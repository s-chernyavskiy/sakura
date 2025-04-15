package config

type AppConfig struct {
	Port       int    `yaml:"port"`
	Host       string `yaml:"host"`
	MaxClients int    `yaml:"max-clients"`
	MaxTimeout int    `yaml:"max-timeout"`
	Verbose    bool   `yaml:"verbose"`
	Logging    bool   `yaml:"logging"`
	Logfile string `yaml:"logfile"`
}
