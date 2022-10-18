package controllers

import (
	"github.com/jphacks/B_2206/backend/domain"
	"github.com/jphacks/B_2206/backend/interfaces/database"
	"github.com/jphacks/B_2206/backend/usecase"

	"github.com/labstack/echo/v4"
	"strconv"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// endpoint POST /users
func (controller *UserController) CreateUser(c echo.Context, name string, email string, password string, personalInfoId int, companyInfoId int) (err error) {
	u := domain.User{
		Name:           name,
		Email:          email,
		Password:       password,
		PersonalInfoId: personalInfoId,
		CompanyInfoId:  companyInfoId,
	}
	c.Bind(&u)
	user, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
	return
}

// endpoint GET /users/:id/
func (controller *UserController) GetUser(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}

// endpoint GET /users/
func (controller *UserController) GetUsers(c echo.Context) (err error) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
	return
}

// endpoint UPDATE /users/:id/
func (controller *UserController) UpdateUser(c echo.Context, name string, email string, password string, personalInfoId int, companyInfoId int) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))
	// idをUser構造体のIDフィールドに格納
	user := domain.User{
		ID:             id,
		Name:           name,
		Email:          email,
		Password:       password,
		PersonalInfoId: personalInfoId,
		CompanyInfoId:  companyInfoId,
	}
	// userをUpdate()に代入
	user, err = controller.Interactor.Update(user)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
	return
}

// endpoint DELETE /users/:id/
func (controller *UserController) DeleteUser(c echo.Context) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))
	// idをUser構造体のIDフィールドに格納
	user := domain.User{ID: id}
	// userをDeleteByUser()に代入
	err = controller.Interactor.DeleteByUser(user)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}
