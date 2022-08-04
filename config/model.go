package config

type Config struct {
	Database Database `yaml:"Database"`
}
type Database struct {
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
}
