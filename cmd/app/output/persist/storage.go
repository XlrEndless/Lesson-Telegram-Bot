package persist

import (
	entity2 "TgBot/cmd/app/output/persist/entity"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB(cfg *DatabaseConfig) *gorm.DB {
	if cfg == nil {
		log.Fatal("config is nil")
	}
	db, err := gorm.Open(postgres.Open(buildDSN(cfg)), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(entity2.Offset{})
	err = db.AutoMigrate(entity2.Lesson{})
	err = db.AutoMigrate(entity2.Teacher{})
	err = db.AutoMigrate(entity2.Student{})
	err = db.AutoMigrate(entity2.Placeholder{})
	return db
}

func buildDSN(cfg *DatabaseConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName)
}
