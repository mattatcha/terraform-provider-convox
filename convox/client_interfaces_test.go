package convox_test

import (
	"io"

	"github.com/convox/rack/pkg/structs"
)

// MockClient is a test mock for the Client interface
type MockClient struct {
	CreateAppFunc      func(name string, generation string) (*structs.App, error)
	GetResourceFunc    func(name string) (*structs.Resource, error)
	CreateResourceFunc func(kind string, options map[string]string) (*structs.Resource, error)
	UpdateResourceFunc func(name string, options map[string]string) (*structs.Resource, error)
	DeleteResourceFunc func(name string) (*structs.Resource, error)

	CreateLinkFunc func(app string, name string) (*structs.Resource, error)
	DeleteLinkFunc func(app string, name string) (*structs.Resource, error)
}

func (m *MockClient) ResetNoop() {
	m.CreateAppFunc = func(name string, generation string) (*structs.App, error) {
		return nil, nil
	}

	m.GetResourceFunc = func(name string) (*structs.Resource, error) {
		return nil, nil
	}

	m.CreateResourceFunc = func(kind string, options map[string]string) (*structs.Resource, error) {
		return nil, nil
	}

	m.UpdateResourceFunc = func(name string, options map[string]string) (*structs.Resource, error) {
		return nil, nil
	}

	m.DeleteResourceFunc = func(name string) (*structs.Resource, error) {
		return nil, nil
	}

	m.CreateLinkFunc = func(name string, app string) (*structs.Resource, error) {
		return nil, nil
	}

	m.DeleteLinkFunc = func(name string, app string) (*structs.Resource, error) {
		return nil, nil
	}
}

func (m *MockClient) CreateApp(name string, generation string) (*structs.App, error) {
	return m.CreateAppFunc(name, generation)
}

func (m *MockClient) GetApp(name string) (*structs.App, error) {
	panic("not implemented")
}

func (m *MockClient) DeleteApp(name string) (*structs.App, error) {
	panic("not implemented")
}

func (m *MockClient) ListFormation(app string) (structs.Formation, error) {
	panic("not implemented")
}

func (m *MockClient) ListParameters(app string) (structs.Parameters, error) {
	panic("not implemented")
}

func (m *MockClient) SetParameters(app string, params map[string]string) error {
	panic("not implemented")
}

func (m *MockClient) GetEnvironment(app string) (structs.Environment, error) {
	panic("not implemented")
}

func (m *MockClient) SetEnvironment(app string, body io.Reader) (structs.Environment, string, error) {
	panic("not implemented")
}

func (m *MockClient) GetResource(name string) (*structs.Resource, error) {
	return m.GetResourceFunc(name)
}

func (m *MockClient) CreateResource(kind string, options map[string]string) (*structs.Resource, error) {
	return m.CreateResourceFunc(kind, options)
}

func (m *MockClient) UpdateResource(name string, options map[string]string) (*structs.Resource, error) {
	return m.UpdateResourceFunc(name, options)
}

func (m *MockClient) DeleteResource(name string) (*structs.Resource, error) {
	return m.DeleteResourceFunc(name)
}

func (m *MockClient) CreateLink(app string, name string) (*structs.Resource, error) {
	return m.CreateLinkFunc(app, name)
}

func (m *MockClient) DeleteLink(app string, name string) (*structs.Resource, error) {
	return m.DeleteLinkFunc(app, name)
}
