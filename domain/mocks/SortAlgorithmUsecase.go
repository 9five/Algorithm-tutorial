// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SortAlgorithmUsecase is an autogenerated mock type for the SortAlgorithmUsecase type
type SortAlgorithmUsecase struct {
	mock.Mock
}

// PrintAry provides a mock function with given fields:
func (_m *SortAlgorithmUsecase) PrintAry() []int {
	ret := _m.Called()

	var r0 []int
	if rf, ok := ret.Get(0).(func() []int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	return r0
}

// Sort provides a mock function with given fields:
func (_m *SortAlgorithmUsecase) Sort() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewSortAlgorithmUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewSortAlgorithmUsecase creates a new instance of SortAlgorithmUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSortAlgorithmUsecase(t mockConstructorTestingTNewSortAlgorithmUsecase) *SortAlgorithmUsecase {
	mock := &SortAlgorithmUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
