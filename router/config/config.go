package config

/*Config is Data Structure for config HTTP Router*/
type Config struct {
	Port uint
}

/*Read Config from env file */
func Read() Config {
	c := Config{
		Port: 3000,
	}
	return c
}
