package config

type Config struct {
	Port              string `env:"PORT"                          envDefault:"8080"`
	FeatureFlagAPIURL string `env:"FEATURE_FLAG_API_URL,required"`
}
