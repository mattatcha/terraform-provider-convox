package convox_test

import (
	"io"

	"github.com/convox/rack/client"
)

// MockClient is a test mock for the Client interface
type MockClient struct {
	CreateAppFunc      func(name string) (*client.App, error)
	GetResourceFunc    func(name string) (*client.Resource, error)
	CreateResourceFunc func(kind string, options map[string]string) (*client.Resource, error)
	UpdateResourceFunc func(name string, options map[string]string) (*client.Resource, error)
	DeleteResourceFunc func(name string) (*client.Resource, error)
}

func (m *MockClient) ResetNoop() {
	m.CreateAppFunc = func(name string) (*client.App, error) {
		return nil, nil
	}

	m.GetResourceFunc = func(name string) (*client.Resource, error) {
		return nil, nil
	}

	m.CreateResourceFunc = func(kind string, options map[string]string) (*client.Resource, error) {
		return nil, nil
	}

	m.UpdateResourceFunc = func(name string, options map[string]string) (*client.Resource, error) {
		return nil, nil
	}

	m.DeleteResourceFunc = func(name string) (*client.Resource, error) {
		return nil, nil
	}
}

func (m *MockClient) CreateApp(name string) (*client.App, error) {
	return m.CreateAppFunc(name)
}

func (m *MockClient) GetApp(name string) (*client.App, error) {
	panic("not implemented")
}

func (m *MockClient) DeleteApp(name string) (*client.App, error) {
	panic("not implemented")
}

func (m *MockClient) ListFormation(app string) (client.Formation, error) {
	panic("not implemented")
}

func (m *MockClient) ListParameters(app string) (client.Parameters, error) {
	panic("not implemented")
}

func (m *MockClient) SetParameters(app string, params map[string]string) error {
	panic("not implemented")
}

func (m *MockClient) GetEnvironment(app string) (client.Environment, error) {
	panic("not implemented")
}

func (m *MockClient) SetEnvironment(app string, body io.Reader) (client.Environment, string, error) {
	panic("not implemented")
}

func (m *MockClient) GetResource(name string) (*client.Resource, error) {
	return m.GetResourceFunc(name)
}

func (m *MockClient) CreateResource(kind string, options map[string]string) (*client.Resource, error) {
	return m.CreateResourceFunc(kind, options)
}

func (m *MockClient) UpdateResource(name string, options map[string]string) (*client.Resource, error) {
	return m.UpdateResourceFunc(name, options)
}

func (m *MockClient) DeleteResource(name string) (*client.Resource, error) {
	return m.DeleteResourceFunc(name)
}
