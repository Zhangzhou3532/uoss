package db

type Config struct {
	Host     string
	Port     uint32
	Username string
	Password string
}

var MysqlConfig = &Config{
	Host:     "",
	Port:     3306,
	Username: "root",
	Password: "mihoyo@321",
}
