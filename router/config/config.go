package config

import (
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

/*Config is Data Structure for config HTTP Router*/
type Config struct {
	Port uint
	Mode string
}

/*Read Config from env file */
func Read() Config {
	env := make(map[string]string)

	env, err := godotenv.Read()
	if err != nil {
		log.Println(".env not found")
	}

	c := Config{
		Port: 3000,
		Mode: "PRODUCTION",
	}
	// / Set Port ------------------------------------------------------
	if env["ROUTER_PORT"] != "" {
		u32, err := strconv.ParseUint(env["ROUTER_PORT"], 10, 32)
		if err == nil {
			c.Port = uint(u32)
		}
	}
	// / ---------------------------------------------------------------
	// / Set Mode ------------------------------------------------------
	if env["ROUTER_MODE"] == "DEBUG" {
		c.Mode = "DEBUG"
	}

	// / ---------------------------------------------------------------

	return c
}

/*RespMsg is Map of Default resp msg in app*/
var RespMsg = map[string]string{
	"NOT_FOUND": "Not Found",
}
