package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	dbName     string
	testDBHost string
	testDBName string
}

func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.dbUser, "dbuser", os.Getenv("POSTGRES_USER"), "DB user name")
	flag.StringVar(&conf.dbPassword, "dbpswd", os.Getenv("POSTGRES_PASSWORD"), "DB pass")
	flag.StringVar(&conf.dbPort, "dbport", os.Getenv("POSTGRES_PORT"), "DB port")
	flag.StringVar(&conf.dbHost, "dbhost", os.Getenv("POSTGRES_HOST"), "DB host")
	flag.StringVar(&conf.dbName, "dbname", os.Getenv("POSTGRES_DB"), "DB name")
	flag.StringVar(&conf.testDBHost, "testdbhost", os.Getenv("TEST_DB_HOST"), "test database host")
	flag.StringVar(&conf.testDBName, "testdbname", os.Getenv("TEST_DB_NAME"), "test database name")

	flag.Parse()

	return conf
}

func (c *Config) GetDBConnectionString() string {
	return c.getDBConnectionString(c.dbHost, c.dbName)
}

func (c *Config) GetTestDBConnectionString() string {
	return c.getDBConnectionString(c.testDBHost, c.testDBName)
}

func (c *Config) getDBConnectionString(dbHost, dbName string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.dbUser,
		c.dbPassword,
		dbHost,
		c.dbPort,
		dbName,
	)
}
