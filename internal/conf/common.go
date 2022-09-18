package conf

type PgSQLConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
	Port     int    `json:"port"`
}
