// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import gqlschema "github.com/kyma-project/kyma/components/ui-api-layer/internal/gqlschema"

import mock "github.com/stretchr/testify/mock"
import v1alpha1 "github.com/kyma-project/kyma/components/remote-environment-broker/pkg/apis/applicationconnector/v1alpha1"

// reConverter is an autogenerated mock type for the reConverter type
type reConverter struct {
	mock.Mock
}

// ToGQL provides a mock function with given fields: in
func (_m *reConverter) ToGQL(in *v1alpha1.RemoteEnvironment) gqlschema.RemoteEnvironment {
	ret := _m.Called(in)

	var r0 gqlschema.RemoteEnvironment
	if rf, ok := ret.Get(0).(func(*v1alpha1.RemoteEnvironment) gqlschema.RemoteEnvironment); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(gqlschema.RemoteEnvironment)
	}

	return r0
}
