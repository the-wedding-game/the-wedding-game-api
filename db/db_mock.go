package db

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	Error error
	items []interface{}
}

func (m MockDB) Where(query interface{}, args ...interface{}) DatabaseInterface {
	return m
}

func (m MockDB) First(dest interface{}, where ...interface{}) DatabaseInterface {
	if len(m.items) == 0 {
		m.Error = gorm.ErrRecordNotFound
		return m
	}

	dest = m.items[0]
	m.items = m.items[1:]
	return m
}

func (m MockDB) Create(value interface{}) DatabaseInterface {
	m.items = append(m.items, value)
	return m
}

func (m MockDB) Find(dest interface{}, where ...interface{}) DatabaseInterface {
	return nil
}

func (m MockDB) GetError() error {
	return nil
}

type DatabaseMocker struct {
	mock.Mock
}

func (m *DatabaseMocker) GetConnection() DatabaseInterface {
	ret := m.Called()
	return ret.Get(0).(DatabaseInterface)
}
