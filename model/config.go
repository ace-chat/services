package model

type Config struct {
	Name    string `yaml:"name" json:"name"`
	Version string `yaml:"version" json:"version"`
	Mode    string `yaml:"mode" json:"mode"`
	Host    string `yaml:"host" json:"host"`
	Port    int    `yaml:"port" json:"port"`

	Redis   Redis   `yaml:"redis" json:"redis"`
	Mysql   Mysql   `yaml:"mysql" json:"mysql"`
	Request Request `yaml:"request" json:"request"`
	Logger  Logger  `yaml:"logger" json:"logger"`
	Upload  Upload  `yaml:"upload" json:"upload"`
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

type Request struct {
	ContentGeneration string `yaml:"contentGeneration" json:"contentGeneration"`
	Analytics         string `yaml:"analytics" json:"analytics"`
}

type Logger struct {
	Path  string `yaml:"path" json:"path"`
	Level string `yaml:"level" json:"level"`
}

type Upload struct {
	Path string `yaml:"path" json:"path"`
}
