package entity

type Config struct {
	Database DBConfig `yaml:"database"`
}

type DBConfig struct {
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	DatabaseName string `yaml:"database_name"`
}
