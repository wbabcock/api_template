package dbcontroller

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"gitlab.com/vimiew/api_template/configuration"
)

// DefaultModelFields :  This is a template used for ORM
type DefaultModelFields struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// Database - Connect to Database
func Database(name string) *gorm.DB {
	// Connection String to database
	connString := ""
	database := configuration.GetDatabaseConfig(name)

	if database.Type == "mssql" {
		connString = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
			database.Host,
			database.Username,
			database.Password,
			database.Port,
			database.Database)
		// fmt.Printf("\nDatabase: %s\nConnection String: %s\n\n", database.Type, connString)
	} else if database.Type == "mysql" {
		connString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			database.Username,
			database.Password,
			database.Host,
			database.Port,
			database.Database)
		// fmt.Printf("\nDatabase: %s\nConnection String: %s\n\n", database.Type, connString)
	} else {
		fmt.Printf("You may have passed the wrong database name. Please verify it is correct. Passed database name: %s", name)
		panic("Failed to get Database name!!!")
	}

	// Open the db connection
	db, err := gorm.Open(database.Type, connString)
	if err != nil {
		fmt.Println("Failed to connect to database")
		panic(err)
	}
	return db
}
