package conf

type Config struct {
	Mysql struct {
		User string `yaml:"user"`
		Host string `yaml:"host"`
		Password string `yaml:"password"`
		Port string `yaml:"port"`
		Name string `yaml:"name"`
	}
	Cache struct {
		Enable bool `yaml:"enable"`
		List []string `yaml:"list,flow"`
	}
}