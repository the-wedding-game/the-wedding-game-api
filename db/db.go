package db

import "fmt"

type DatabaseInterface interface {
	Where(query interface{}, args ...interface{}) DatabaseInterface
	First(dest interface{}, where ...interface{}) DatabaseInterface
	Create(value interface{}) DatabaseInterface
	Find(dest interface{}, where ...interface{}) DatabaseInterface
	GetError() error
}

func getConnection() DatabaseInterface {
	fmt.Println("Connecting to database")
	return newDatabase()
}

var GetConnection = getConnection
