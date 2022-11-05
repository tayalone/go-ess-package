package config

import (
	"fmt"
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

/*Config is Data Structure for config HTTP Router*/
type Config struct {
	Port uint
}

/*Read Config from env file */
func Read() Config {
	env := make(map[string]string)

	env, err := godotenv.Read()
	if err != nil {
		log.Panicln(".env not found")
	}

	fmt.Println("env", env)

	c := Config{
		Port: 3000,
	}
	// / Set Port ------------------------------------------------------
	if env["ROUTER_PORT"] != "" {
		u32, err := strconv.ParseUint(env["ROUTER_PORT"], 10, 32)
		if err == nil {
			c.Port = uint(u32)
		}
	}
	// // ---------------------------------------------------------------
	return c
}
