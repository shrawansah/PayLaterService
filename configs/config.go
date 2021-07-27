package configs

type configs struct {
	Environment string

	DBHost 		string
	DBPassword  string
	DBUser 		string
	DBPort 		string
	DBSchema 	string
	DBDriver	string

	UserAtCreditLimitThreshold int64
}

var Configs = configs {
	Environment: "development",
	
	DBHost: "127.0.0.1",
	DBPassword: "YOUR_PASSWORD",
	DBUser: "YOUR_USER",
	DBSchema: "YOUR_SCHEMA",
	DBPort: "3306",
	DBDriver: "mysql",

	UserAtCreditLimitThreshold: 0,
}