package models

func SetupMockDb() *MockDB {
	mockDB := &MockDB{}
	GetConnection = func() DatabaseInterface {
		return mockDB
	}
	return mockDB
}
