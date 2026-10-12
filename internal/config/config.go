package config

import (
	"errors"
	"flag"
)

type Config struct {
	Endpoint string
	DSN      string
	Debug    bool
}

func Load() (Config, error) {
	var c Config
	flag.StringVar(&c.Endpoint, "endpoint", "opc.tcp://192.168.0.1:4840", "OPC UA Endpoint URL")
	flag.StringVar(&c.DSN, "dsn", "postgres://postgres@localhost:5432/plc200_data?sslmode=disable", "строка подключения к PostgreSQL.")
	flag.BoolVar(&c.Debug, "debug", false, "enable debug logging")
	flag.Parse()

	if c.DSN == "" {
		return c, errors.New("не задана строка подключения к БД: флаг -dsn или переменная PLC_DB_DSN")
	}
	return c, nil
}
