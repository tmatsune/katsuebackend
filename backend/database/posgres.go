package database

import (
  "gorm.io/gorm"
  "fmt"
  "gorm.io/driver/postgres"
)

type Config struct{
	Host string
	User string
	Password string
	DBname string
	Port string
	Sslmode string
}
var config Config = Config{
	Host: "localhost",
	User: "postgres",
	Password: "vgislife22",
	DBname: "katsuedb",
	Port: "5432",
	Sslmode: "disable",
}

func Connect(config *Config) (*gorm.DB, error){
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s slmode=%s TimeZone=los angeles",
	config.Host, config.User, config.Password, config.DBname, config.Port, config.Sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("adsdsf")
	}
	fmt.Print(db);
	return db, nil;
}