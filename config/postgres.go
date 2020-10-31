package config

// 构造 postgres 配置结构体， 配置文件为 config.yaml
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
