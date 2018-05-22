package config

type Connections struct {
	DATABASE_IP       string
	DATABASE_PORT     string
	DATABASE_USERNAME string
	DATABASE_PASSWORD string
	DATABASE_NAME     string
}

func GetCons() map[string]*Connections {

	//return map[string]*Connections{
	//	"official_account": &Connections{
	//		DATABASE_USERNAME: "11",
	//		DATABASE_IP:       "11",
	//		DATABASE_PASSWORD: "11",
	//		DATABASE_NAME:     "11",
	//		DATABASE_PORT:     "3306",
	//	},
	//}

	return map[string]*Connections{}
}
