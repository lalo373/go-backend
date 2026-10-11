package config

type Config struct {
	Port string
	ENV  string
}

func MustLoad()
