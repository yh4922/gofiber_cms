package setup

import (
	"fmt"
	"time"

	g "gopress/internal/global"
	"gopress/internal/models"

	"github.com/gookit/config/v2"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDatabase() {
	driver := config.String("Database.Driver", "none")
	if driver == "none" {
		panic("The database driver is not configured.")
	} else if driver == "mysql" {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s",
			config.String("Database.Username", "root"),
			config.String("Database.Password", ""),
			config.String("Database.Host", "127.0.0.1"),
			config.Int("Database.Port", 3306),
			config.String("Database.Database", "gopress"),
		)

		fmt.Printf("Initializing database connection to %s:%d...\n", config.String("Database.Host", "127.0.0.1"), config.Int("Database.Port", 3306))
		var db *gorm.DB
		var err error
		maxRetries := 3
		for i := 0; i < maxRetries; i++ {
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err == nil {
				break
			}
			fmt.Printf("Failed to connect to database (attempt %d/%d): %v\n", i+1, maxRetries, err)
			if i < maxRetries-1 {
				time.Sleep(time.Second * 3)
			}
		}
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to database after %d attempts: %v", maxRetries, err))
		}
		g.SetDatabase(db)
		fmt.Println("Database connection established successfully")
	} else if driver == "sqlite" {
		db, err := gorm.Open(sqlite.Open(config.String("Database.Dbpath", "./runtime/database.db")), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		g.SetDatabase(db)
	} else {
		panic("The database driver is not supported.")
	}

	models.InitModel()
}
