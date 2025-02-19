package test

import (
	"errors"
	"reflect"
	"the-wedding-game-api/db"
)

type MockDB struct {
	items []interface{}
	Error error
}

func (m *MockDB) Where(_ interface{}, _ ...interface{}) db.DatabaseInterface {
	return m
}

func (m *MockDB) First(dest interface{}, _ ...interface{}) db.DatabaseInterface {
	if len(m.items) == 0 {
		destValue := reflect.ValueOf(dest)
		m.Error = errors.New("record not found: " + destValue.Type().String())
		return m
	}

	destValue := reflect.ValueOf(dest)
	if destValue.Kind() != reflect.Ptr || destValue.IsNil() {
		m.Error = errors.New("dest must be a non-nil pointer")
		return m
	}

	itemValue := reflect.ValueOf(m.items[0])
	if itemValue.Kind() == reflect.Ptr {
		itemValue = itemValue.Elem()
	}

	destValue.Elem().Set(itemValue)
	m.items = m.items[1:]
	return m
}

func (m *MockDB) Create(value interface{}) db.DatabaseInterface {
	m.items = append(m.items, value)
	return m
}

func (m *MockDB) Find(_ interface{}, _ ...interface{}) db.DatabaseInterface {
	return nil
}

func (m *MockDB) GetError() error {
	return m.Error
}
