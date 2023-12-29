package database

import (
	"scsystem/pkg/utils"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connection *gorm.DB = nil
var lock *sync.Mutex = &sync.Mutex{}

func Connection() (*gorm.DB, error) {
	if connection == nil {
		lock.Lock()
		defer lock.Unlock()
		if connection == nil {
			dsn, _ := utils.ConnectionURLBuilder("mysql")
			conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				return nil, err
			}
			connection = conn
		}
	}
	return connection, nil
}
