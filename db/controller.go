package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// DefaultModelFields :  This is a template used for ORM
type DefaultModelFields struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// Database - Connect to Database
func Database() *gorm.DB {
	// SQL Server Connect String
	connection := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		"localhost",
		"sa",
		"P@ssw0rd",
		1433,
		"database_name")
	
	// MySQL/MariaDB Connection String
	// connesction := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	"localhost",
	//	"root",
	//	"P@ssw0rd",
	//	3306,
	//	"database_name")


	// Open the db connection
	db, err := gorm.Open("mssql", connection)
	if err != nil {
		fmt.Println("Failed to connect to database")
		panic(err)
	}
	return db
}
