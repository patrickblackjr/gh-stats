package repos

import "github.com/stretchr/testify/mock"

type GitHubAPI interface {
	FetchData(query string) (any, error)
}

type RealGitHubAPI struct{}

func (api *RealGitHubAPI) FetchData(query string) (any, error) {
	return nil, nil
}

type MockGitHubAPI struct {
	mock.Mock
}

func (m *MockGitHubAPI) FetchData(query string) (any, error) {
	args := m.Called(query)
	return args.Get(0), args.Error(1)
}
