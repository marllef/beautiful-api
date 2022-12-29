package configs

type configs struct {
	Port        string `mapstructure:"PORT"`
	DatabaseUrl string `mapstructure:"DATABASE_URL"`
	GPTApiKey string `mapstructure:"GPT_API_KEY"`
}


