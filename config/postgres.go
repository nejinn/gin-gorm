package config

type PostgresYaml struct {
	Postgres PostgresConf `yaml:"postgres"`
}

type PostgresConf struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	DataBase     string `yaml:"database"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Sslmode      string `yaml:"sslmode"`
	MaxIdleConns int    `yaml:"maxidleconns"`
	MaxOpenConns int    `yaml:"maxopenconns"`
}
