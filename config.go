package main

type AppConfig struct {
	Port             int    `default:"8001"`
	Env              string `envconfig:"ENV_NAME" required:"true"`
	MaxRandomDelayMS int `envconfig:"MAX_RANDOM_DELAY_MS" required:"true"`
}

func (c *AppConfig) IsLocal() bool {
	return c.Env == "local"
}
