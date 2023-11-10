package model

type Config struct {
	Name    string `yaml:"name" json:"name"`
	Version string `yaml:"version" json:"version"`
	Mode    string `yaml:"mode" json:"mode"`
	Host    string `yaml:"host" json:"host"`
	Port    int    `yaml:"port" json:"port"`

	Redis Redis `yaml:"redis" json:"redis"`
	Mysql Mysql `yaml:"mysql" json:"mysql"`
}

type Redis struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	DB       int    `yaml:"db" json:"db"`
}

type Mysql struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	DB       string `yaml:"db" json:"db"`
}
