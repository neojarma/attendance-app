package connection

import (
	"errors"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func SQLServerConn() (*gorm.DB, error) {
	dsn := os.Getenv("DB_URL")

	maxCounter := 10

	for i := 0; i < maxCounter; i++ {
		db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: "dbo.",
			},
		})
		if err == nil {
			log.Println("SQL Server Ready")
			return db, nil
		}

		time.Sleep(30 * time.Second)
		log.Println("SQL Server not yet ready, try again in 30 second")
	}

	return nil, errors.New("failed to connect to SQL Server")
}
