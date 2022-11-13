package config

import (
	"log"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

const (
	mb = 1 << 20
)

/*Config is Data Structure for config HTTP Router*/
type Config struct {
	Port            uint
	Mode            string
	ReadTimeOutSec  time.Duration
	WriteTimeOutSec time.Duration
	IdealTimeOutSec time.Duration
	MaxHeadBytes    int
}

func (c *Config) setUint(key string, value string) {
	if value != "" {
		u32, err := strconv.ParseUint(value, 10, 32)
		if err == nil {
			switch key {
			case "ROUTER_PORT":
				c.Port = uint(u32)
			case "ROUTER_READTIMEOUT_SEC":
				c.ReadTimeOutSec = time.Duration(u32) * time.Second
			case "ROUTER_WRITETIMEOUT_SEC":
				c.WriteTimeOutSec = time.Duration(u32) * time.Second
			case "ROUTER_IDELTIMEOUT_SEC":
				c.IdealTimeOutSec = time.Duration(u32) * time.Second
			default:
			}
		}
	}
}

// func strToUint(key string) (uint error) {
// }

/*Read Config from env file */
func Read() Config {
	env := make(map[string]string)

	env, err := godotenv.Read()
	if err != nil {
		log.Println(".env not found")
	}

	c := Config{
		Port:            3000,
		Mode:            "PRODUCTION",
		ReadTimeOutSec:  10 * time.Second,
		WriteTimeOutSec: 10 * time.Second,
		IdealTimeOutSec: 10 * time.Second,
		MaxHeadBytes:    mb,
	}
	// / Set Port ------------------------------------------------------
	c.setUint("ROUTER_PORT", env["ROUTER_PORT"])
	// / ---------------------------------------------------------------
	// / Set Mode ------------------------------------------------------
	if env["ROUTER_MODE"] == "DEBUG" {
		c.Mode = "DEBUG"
	}
	// / ---------------------------------------------------------------
	c.setUint("ROUTER_READTIMEOUT_MS", env["ROUTER_READTIMEOUT_MS"])
	c.setUint("ROUTER_WRITETIMEOUT_MS", env["ROUTER_WRITETIMEOUT_MS"])
	c.setUint("ROUTER_IDELTIMEOUT_MS", env["ROUTER_IDELTIMEOUT_MS"])

	// // -------- max header bytes ----------------------------------
	if env["ROUTER_MAXHEADERBYTES_MB"] != "" {
		if s, err := strconv.ParseFloat(env["ROUTER_MAXHEADERBYTES_MB"], 32); err == nil {
			// not error && s > 0.01mb && s <= 5mb
			if err == nil && s > 0.01 && s <= 5 {
				c.MaxHeadBytes = int(s * mb)
			}
		}
	}

	// // ------------------------------------------------------------

	return c
}

/*RespMsg is Map of Default resp msg in app*/
var RespMsg = map[string]string{
	"NOT_FOUND": "Not Found",
}
