// Code generated by mockery v1.0.0. DO NOT EDIT.
package automock

import fileheader "github.com/kyma-project/kyma/components/asset-metadata-service/pkg/fileheader"
import mock "github.com/stretchr/testify/mock"

// Extractor is an autogenerated mock type for the Extractor type
type Extractor struct {
	mock.Mock
}

// ReadMetadata provides a mock function with given fields: fileHeader
func (_m *Extractor) ReadMetadata(fileHeader fileheader.FileHeader) (map[string]interface{}, error) {
	ret := _m.Called(fileHeader)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(fileheader.FileHeader) map[string]interface{}); ok {
		r0 = rf(fileHeader)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(fileheader.FileHeader) error); ok {
		r1 = rf(fileHeader)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
