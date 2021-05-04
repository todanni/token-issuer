package database

import (
	"github.com/token-issuer/internal/config"
	"gorm.io/gorm"
)

func Open(cfg config.Config) (*gorm.DB, error) {
	//// Make connection string
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	//
	//db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
	//	Logger: logger.Default.LogMode(logger.Info),
	//})
	return nil, nil
}
