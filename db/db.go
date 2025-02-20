package db

var databaseConnection DatabaseInterface

type DatabaseInterface interface {
	Where(query interface{}, args ...interface{}) DatabaseInterface
	First(dest interface{}, where ...interface{}) DatabaseInterface
	Create(value interface{}) DatabaseInterface
	Find(dest interface{}, where ...interface{}) DatabaseInterface
	GetError() error
}

func getConnection() DatabaseInterface {
	if databaseConnection != nil {
		return databaseConnection
	}
	databaseConnection = newDatabase()
	return databaseConnection
}

func ResetConnection() {
	databaseConnection = nil
}

var GetConnection = getConnection
