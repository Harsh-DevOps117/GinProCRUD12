package db

import (
	"log"
	"os"

	"github.com/harshdevops117/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func DataBaseInit() (*gorm.DB,error){

	err := godotenv.Load("../.env")
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	db_url:=os.Getenv("DB_URL")
	db,err:=gorm.Open(postgres.Open(db_url),&gorm.Config{})
	if err!=nil{
		return nil,err
	}
	return db,nil
}

func DBAutoMigrate(db *gorm.DB){
	db.AutoMigrate(&models.User{},&models.Notes{})
}
