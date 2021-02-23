package conf

type Yaml struct {
	Mysql struct {
		User         string `yaml:"user"`
		Host         string `yaml:"host"`
		Server       string `yaml:server`
		Password     string `yaml:"password"`
		Port         int    `yaml:"port"`
		Name         string `yaml:"name"`
		MaxIdleConns int    `yaml:"maxIdleConns"`
		MaxOpenConns int    `yaml:"maxOpenConns"`
	}

	Redis struct {
		Addr         string `yaml:"Addr"`
		Password     string `yaml:"Password"`
		MaxRetries   int    `yaml:"MaxRetries"`
		DialTimeout  string `yaml:"DialTimeout"`
		WriteTimeout int    `yaml:"WriteTimeout"`
		PoolSize     int    `yaml:"PoolSize"`
		PoolTimeout  int    `yaml:"PoolTimeout"`
		IdleTimeout  int    `yaml:"IdleTimeout"`
		DB           int    `yaml:"DB"`
	}
}
