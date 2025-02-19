package db

type DatabaseInterface interface {
	Where(query interface{}, args ...interface{}) DatabaseInterface
	First(dest interface{}, where ...interface{}) DatabaseInterface
	Create(value interface{}) DatabaseInterface
	Find(dest interface{}, where ...interface{}) DatabaseInterface
	GetError() error
}

func getConnection() DatabaseInterface {
	return newDatabase()
}

var GetConnection = getConnection
