package config

type appOptions struct {
	Env  string
	Port string
}

type dbOptions struct {
	Driver   string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}
