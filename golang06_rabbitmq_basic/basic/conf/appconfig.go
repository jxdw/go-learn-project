package conf

type AppConfig struct {
	RabbitmqAddress string `yaml:"rabbitmqAddress"`
	RabbitmqQueue string `yaml:"rabbitmqQueue"`
}
