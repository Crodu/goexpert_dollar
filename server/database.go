package server

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Usdbrl struct
type Usdbrl struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

// DatabaseConnection creates a connection to the database.
func DatabaseConnection(timeout ...time.Duration) (*gorm.DB, context.CancelFunc, error) {

	var millis time.Duration

	if len(timeout) > 0 {
		millis = timeout[0]
	} else {
		millis = 10
	}

	ctx, cancel := context.WithTimeout(context.Background(), millis*time.Millisecond)

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		return nil, cancel, err
	}
	return db.WithContext(ctx), cancel, nil
}

// SetupDatabase sets up the database.
func SetupDatabase() (*gorm.DB, error) {

	db, cancel, err := DatabaseConnection(1000)
	defer cancel()
	if err != nil {
		return nil, err
	}

	// Auto Migrate the Exchange struct
	err = db.AutoMigrate(&Usdbrl{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Auto migration successful.")

	return db, nil
}

// InsertData inserts data into the database.
func InsertData(db *gorm.DB, data Usdbrl) error {
	err := db.Create(&data).Error
	if err != nil {
		return err
	}
	fmt.Println("Data inserted successfully.")
	return nil
}
