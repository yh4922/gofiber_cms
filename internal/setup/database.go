package setup

import (
	"fmt"

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
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.String("Database.Username", "root"),
			config.String("Database.Password", "123456"),
			config.String("Database.Host", "127.0.0.1"),
			config.Int("Database.Port", 3306),
			config.String("Database.Database", "gopress_db"),
		)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		g.SetDatabase(db)
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
