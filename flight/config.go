package flight

type Config struct {
	filename string
}

func (c *Config) Filename() string {
	return c.filename
}

func NewConfig() *Config {
	return &Config{
		filename: "flight_data.json",
	}
}
