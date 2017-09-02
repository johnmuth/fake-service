package main

type AppConfig struct {
	Port           int    `default:"8001"`
	Env            string `envconfig:"ENV_NAME" required:"true"`
}

func (c *AppConfig) IsLocal() bool {
	return c.Env == "local"
}
