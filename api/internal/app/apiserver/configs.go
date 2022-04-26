package apiserver

type Api struct {
	DatabaseURL string `toml:"database_url"`
	BindAdress 	string `toml:"Bind_adr"`
	LogLev 		string `toml:"log_lev"`
	SessionKey  string `toml:"session_key"`
}

func NewConfigs() *Api{
	return &Api{
		BindAdress: ":8080",
		LogLev: 	"debug",
	}
}