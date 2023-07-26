package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/nilu2039/tokiio-server-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var err error

// type CamelCaseNamingStrategy struct{}

// func (strategy CamelCaseNamingStrategy) TableNamer(name string) string {
// 	return toCamelCase(name)
// }

// func (strategy CamelCaseNamingStrategy) ColumnNamer(table, name string) string {
// 	return toCamelCase(name)
// }

// func (strategy CamelCaseNamingStrategy) CheckerName(tableName, columnName string) string {
// 	return ""
// }

// func toCamelCase(name string) string {
// 	words := strings.Fields(strings.ToLower(name))
// 	for i, word := range words {
// 		if i > 0 {
// 			words[i] = strings.Title(word)
// 		}
// 	}
// 	return strings.Join(words, "")
// }

func ConnectToDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=require", os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(&models.History{}, &models.HistoryPerEpisode{})
}
