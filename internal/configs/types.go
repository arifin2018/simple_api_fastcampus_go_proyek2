package configs

type (
	Config struct {
		Service  Service  `mapstructure:"service"`
		Database Database `mapstructure:"database"`
	}

	Service struct {
		Port      string `mapstructure:"port"`
		SecretJwt string `mapstructure:"secretJWT"`
	}

	Database struct {
		Drivername string `mapstructure:"drivername"`
		Host       string `mapstructure:"host"`
		Port       string `mapstructure:"port"`
		Username   string `mapstructure:"username"`
		Password   string `mapstructure:"password"`
		DB         string `mapstructure:"db"`
	}
)
