// Code generated by mockery v1.0.0. DO NOT EDIT.
package pg

import mock "github.com/stretchr/testify/mock"

// MockDB is an autogenerated mock type for the DB type
type MockDB struct {
	mock.Mock
}

// AlterDatabaseOwner provides a mock function with given fields: _a0, _a1
func (_m *MockDB) AlterDatabaseOwner(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BigintQuery provides a mock function with given fields: _a0, _a1
func (_m *MockDB) BigintQuery(_a0 string, _a1 ...interface{}) (int64, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, ...interface{}) int64); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BoolQuery provides a mock function with given fields: _a0, _a1
func (_m *MockDB) BoolQuery(_a0 string, _a1 ...interface{}) (bool, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, ...interface{}) bool); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *MockDB) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConnectedUserIsSuperuser provides a mock function with given fields:
func (_m *MockDB) ConnectedUserIsSuperuser() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDatabase provides a mock function with given fields: _a0
func (_m *MockDB) CreateDatabase(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateDatabaseWithOwner provides a mock function with given fields: _a0, _a1
func (_m *MockDB) CreateDatabaseWithOwner(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateExtension provides a mock function with given fields: _a0
func (_m *MockDB) CreateExtension(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateRole provides a mock function with given fields: _a0
func (_m *MockDB) CreateRole(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DatabaseExists provides a mock function with given fields: _a0
func (_m *MockDB) DatabaseExists(_a0 string) (bool, error) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DropDatabase provides a mock function with given fields: _a0
func (_m *MockDB) DropDatabase(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExecStatement provides a mock function with given fields: _a0, _a1
func (_m *MockDB) ExecStatement(_a0 string, _a1 ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...interface{}) error); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GrantAll provides a mock function with given fields: _a0, _a1
func (_m *MockDB) GrantAll(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ping provides a mock function with given fields:
func (_m *MockDB) Ping() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemovePassword provides a mock function with given fields: _a0
func (_m *MockDB) RemovePassword(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RenameDatabase provides a mock function with given fields: _a0, _a1
func (_m *MockDB) RenameDatabase(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StringQuery provides a mock function with given fields: _a0, _a1
func (_m *MockDB) StringQuery(_a0 string, _a1 ...interface{}) (string, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...interface{}) string); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
