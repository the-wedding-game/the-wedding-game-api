package db

var databaseConnection DatabaseInterface

type DatabaseInterface interface {
	GetSession() DatabaseInterface
	Where(query interface{}, args ...interface{}) DatabaseInterface
	First(dest interface{}, where ...interface{}) DatabaseInterface
	Create(value interface{}) DatabaseInterface
	Find(dest interface{}, where ...interface{}) DatabaseInterface
	GetError() error
}

func getConnection() DatabaseInterface {
	if databaseConnection != nil {
		return databaseConnection.GetSession()
	}
	databaseConnection = newDatabase()
	return databaseConnection.GetSession()
}

func ResetConnection() {
	databaseConnection = nil
}

var GetConnection = getConnection
