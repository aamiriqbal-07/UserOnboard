package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"userOnboard/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockUserService struct {
	CreateUserFunc func(user models.User) error
	GetUserByIDFunc func(id string) (*models.User, error)
	ListUsersFunc   func() ([]models.User, error)
}

func (m *MockUserService) CreateUser(user models.User) error {
	return m.CreateUserFunc(user)
}

func (m *MockUserService) GetUserByID(id string) (*models.User, error) {
	return m.GetUserByIDFunc(id)
}

func (m *MockUserService) ListUsers() ([]models.User, error) {
	return m.ListUsersFunc()
}

func setupRouter(ctrl *UserController) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/users", ctrl.CreateUser)
		api.GET("/users/:id", ctrl.GetUserByID)
		api.GET("/users", ctrl.ListUsers)
	}
	return r
}

func TestCreateUserController(t *testing.T) {
	mockService := &MockUserService{
		CreateUserFunc: func(user models.User) error {
			return nil
		},
	}
	ctrl := NewUserController(mockService)
	r := setupRouter(ctrl)

	t.Run("Success", func(t *testing.T) {
		body := `{"id": "1", "name": "John Doe", "signupTime": 1700000000000}`
		req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestGetUserByIDController(t *testing.T) {
	mockService := &MockUserService{
		GetUserByIDFunc: func(id string) (*models.User, error) {
			return &models.User{ID: "1", Name: "John Doe", SignupTime: 1700000000000}, nil
		},
	}
	ctrl := NewUserController(mockService)
	r := setupRouter(ctrl)

	t.Run("Success", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/users/1", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
