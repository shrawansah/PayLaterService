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
	DBPassword: "iuseNOTE4@OHAAHAI",
	DBUser: "root",
	DBSchema: "simpl_paylater",
	DBPort: "3306",
	DBDriver: "mysql",

	UserAtCreditLimitThreshold: 0,
}