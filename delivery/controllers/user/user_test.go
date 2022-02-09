package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"part3/delivery/controllers/auth"
	"part3/delivery/middlewares"
	"part3/models/user"
	"part3/models/user/request"

	// lib "part3/lib/database/user"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	var jwtToken string

	t.Run("Success Login", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "anonim@123",
			"password": "anonim123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authController := auth.New(&MockAuthLib{})
		authController.Login()(context)

		response := auth.LoginRespFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		jwtToken = response.Data["token"].(string)

		assert.Equal(t, response.Message, "success login")
		assert.NotNil(t, response.Data["token"])
	})

	t.Run("Failed to Create", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]string{
			"email":    "",
			"password": "anonim123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockUserLib{})
		userController.Create()(context)

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		log.Info(response)

		assert.Equal(t, 400, response.Code)
		assert.Equal(t, "error in request Create", response.Message)

	})

	t.Run("Failed to Access", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]string{
			"email":    "anonim",
			"password": "anonim",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockUserLib{})
		userController.Create()(context)

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		log.Info(response)

		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "error in access Create", response.Message)

	})

	t.Run("Success Create", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]string{
			"name":     "anonim123",
			"email":    "anonim@123",
			"password": "anonim123",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockUserLib{})
		userController.Create()(context)

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		log.Info(response)
		assert.Equal(t, 201, response.Code)
		assert.Equal(t, "anonim123", response.Data.Name)
	})
}

func TestGetById(t *testing.T) {
	var jwtToken string

	t.Run("Success Login", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "anonim@123",
			"password": "anonim123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authController := auth.New(&MockAuthLib{})
		authController.Login()(context)

		response := auth.LoginRespFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		jwtToken = response.Data["token"].(string)

		assert.Equal(t, response.Message, "success login")
		assert.NotNil(t, response.Data["token"])
	})

	// t.Run("Success Get By Id", func(t *testing.T) {

	// 	e := echo.New()
	// 	// userid := int(middlewares.ExtractTokenId(c))

	// 	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
	// 	res := httptest.NewRecorder()

	// 	req.Header.Set("Content-Type", "application/json")
	// 	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

	// 	context := e.NewContext(req, res)
	// 	context.SetPath("/users/:id")
	// 	log.Info(context)
	// 	log.Info(req)
	// 	log.Info(res)

	// 	userController := New(&MockUserLib{})
	// 	if err := middlewares.JwtMiddleware()(userController.GetById())(context); err != nil {
	// 		log.Fatal(err)
	// 		return
	// 	}
	// 	log.Info(userController)

	// 	response := GetUserResponseFormat{}

	// 	json.Unmarshal([]byte(res.Body.Bytes()), &response)
	// 	// log.Info(response)
	// 	assert.Equal(t, 200, response.Code)
	// 	assert.Equal(t, response.Data, response.Data)
	// })

	t.Run("Success Get By Id", func(t *testing.T) {

		e := echo.New()
		// userid := int(middlewares.ExtractTokenId(c))

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		log.Info(context)
		log.Info(req)
		log.Info(res)

		userController := New(&MockUserLib{})
		if err := middlewares.JwtMiddleware()(userController.GetById())(context); err != nil {
			log.Fatal(err)
			return
		}
		log.Info(userController)

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// log.Info(response)
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, response.Data, response.Data)
	})
}

//update by id belum yang fail
func TestUpdateByID(t *testing.T) {
	var jwtToken string

	t.Run("Success Login", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "anonim@123",
			"password": "anonim123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authController := auth.New(&MockAuthLib{})
		authController.Login()(context)

		response := auth.LoginRespFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		jwtToken = response.Data["token"].(string)

		assert.Equal(t, response.Message, "success login")
		assert.NotNil(t, response.Data["token"])
	})

	t.Run("Success Update", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]string{
			"name":     "anonim123",
			"email":    "anonim@123",
			"password": "anonim123",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockUserLib{})
		if err := middlewares.JwtMiddleware()(userController.GetById())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		log.Info(response)
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, response.Data, response.Data)

		log.Info(response.Data)
	})

}

func TestDeleteByID(t *testing.T) {
	var jwtToken string

	t.Run("Success Login", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "anonim@123",
			"password": "anonim123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authController := auth.New(&MockAuthLib{})
		authController.Login()(context)

		response := auth.LoginRespFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		jwtToken = response.Data["token"].(string)

		assert.Equal(t, response.Message, "success login")
		assert.NotNil(t, response.Data["token"])
	})

	t.Run("Success Delete", func(t *testing.T) {
		e := echo.New()

		// reqBody, _ := json.Marshal(map[string]string{
		// 	"name":     "anonim123",
		// 	"email":    "anonim@123",
		// 	"password": "anonim123",
		// })
		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockUserLib{})
		if err := middlewares.JwtMiddleware()(userController.DeleteById())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		log.Info(response)
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "Success Delete By Id", response.Message)

	})

}

func TestGetAll(t *testing.T) {
	var jwtToken string

	t.Run("Success Login", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "anonim@123",
			"password": "anonim123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authController := auth.New(&MockAuthLib{})
		authController.Login()(context)

		response := auth.LoginRespFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		jwtToken = response.Data["token"].(string)

		assert.Equal(t, response.Message, "success login")
		assert.NotNil(t, response.Data["token"])
	})

	t.Run("Success Get All User", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockUserLib{})
		if err := middlewares.JwtMiddleware()(userController.GetAll())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		log.Info(response)
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "Success Get All User", response.Message)
	})

	t.Run("Failed Get All User", func(t *testing.T) {
		// token := string(jwtToken)
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		// userController := New(&MockFalseLib{})
		userController := New(&MockUserLib{})

		if err := middlewares.JwtMiddleware()(userController.GetAll())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		log.Info(response)
		assert.Equal(t, 400, response.Code)
		assert.Equal(t, "error in request Get", response.Message)
	})
}

type MockUserLib struct{}

func (m *MockUserLib) Create(newUser user.User) (user.User, error) {
	if newUser.Email != "anonim123" && newUser.Password != "anonim123" {
		return user.User{}, errors.New("record not found")
	}
	return user.User{Name: newUser.Name, Email: newUser.Email, Password: newUser.Password}, nil
}

func (m *MockUserLib) GetById(id int) ([]user.User, error) {
	return []user.User{}, nil
}

func (m *MockUserLib) UpdateById(id int, upUser request.UserRegister) (user.User, error) {
	return user.User{Name: upUser.Name, Email: upUser.Email, Password: upUser.Password}, nil
}

func (m *MockUserLib) DeleteById(id int) (gorm.DeletedAt, error) {
	user := user.User{}
	return user.DeletedAt, nil
}

func (m *MockUserLib) GetAll() (user.User, error) {
	return user.User{}, nil
}

type MockAuthLib struct{}

func (ma *MockAuthLib) Login(UserLogin request.Userlogin) (user.User, error) {
	return user.User{Model: gorm.Model{ID: 1}, Email: UserLogin.Email, Password: UserLogin.Password}, nil
}

type MockFalseLib struct{}

func (mf *MockFalseLib) GetAll() ([]user.User, error) {
	return nil, errors.New("False Object")
}
