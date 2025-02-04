package service

import (
	"errors"
	"testing"
	"userOnboard/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByID(id string) (*models.User, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) ListUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := NewUserService(mockRepo)

	user := models.User{ID: "1", Name: "John Doe", SignupTime: 1700000000000}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("CreateUser", user).Return(nil)
		err := svc.CreateUser(user)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetUserByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := NewUserService(mockRepo)

	user := models.User{ID: "1", Name: "John Doe", SignupTime: 1700000000000}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("GetUserByID", "1").Return(&user, nil)
		result, err := svc.GetUserByID("1")
		assert.NoError(t, err)
		assert.Equal(t, "1", result.ID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		mockRepo.On("GetUserByID", "2").Return(nil, errors.New("user not found"))
		_, err := svc.GetUserByID("2")
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestListUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := NewUserService(mockRepo)

	users := []models.User{
		{ID: "1", Name: "John Doe", SignupTime: 1700000000000},
		{ID: "2", Name: "Jane Doe", SignupTime: 1700000001000},
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("ListUsers").Return(users, nil)
		result, err := svc.ListUsers()
		assert.NoError(t, err)
		assert.Len(t, result, 2)
		mockRepo.AssertExpectations(t)
	})

	t.Run("EmptyList", func(t *testing.T) {
		mockRepo.On("ListUsers").Return([]models.User{}, nil)
		result, err := svc.ListUsers()
		assert.NoError(t, err)
		assert.Len(t, result, 2)
		mockRepo.AssertExpectations(t)
	})
}
