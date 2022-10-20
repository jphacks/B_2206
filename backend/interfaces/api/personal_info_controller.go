package controllers

import (
	"github.com/jphacks/B_2206/backend/domain"
	"github.com/jphacks/B_2206/backend/interfaces/database"
	"github.com/jphacks/B_2206/backend/usecase"

	"strconv"

	"github.com/labstack/echo/v4"
)

type PersonalInfoController struct {
	Interactor usecase.PersonalInfoInteractor
}

func NewPersonalInfoController(sqlHandler database.SqlHandler) *PersonalInfoController {
	return &PersonalInfoController{
		Interactor: usecase.PersonalInfoInteractor{
			PersonalInfoRepository: &database.PersonalInfoRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// endpoint POST /personalInfos
func (controller *PersonalInfoController) CreatePersonalInfo(c echo.Context, familyName string, firstName string, familyNameKana string, firstNameKana string, birthday string, phoneNumber string, userID int) (err error) {
	p := domain.PersonalInfo{
		FamilyName:     familyName,
		FirstName:      firstName,
		FamilyNameKana: familyNameKana,
		FirstNameKana:  firstNameKana,
		Birthday:       birthday,
		PhoneNumber:    phoneNumber,
		UserID:         userID,
	}
	c.Bind(&p)
	personalInfo, err := controller.Interactor.Add(p)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, personalInfo)
	return
}

// endpoint GET /personalInfos
func (controller *PersonalInfoController) GetPersonalInfos(c echo.Context) (err error) {
	personalInfo, err := controller.Interactor.PersonalInfos()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, personalInfo)
	return
}

// endpoint GET /personalInfos/:id
func (controller *PersonalInfoController) GetPersonalInfo(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	personalInfos, err := controller.Interactor.PersonalInfoById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, personalInfos)
	return
}

// endpoint UPDATE /personalInfos/:id/
func (controller *PersonalInfoController) UpdatePersonalInfo(c echo.Context, familyName string, firstName string, familyNameKana string, firstNameKana string, birthday string, phoneNumber string) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))
	// idをPersonalInfo構造体のIDフィールドに格納
	personalInfo := domain.PersonalInfo{
		ID:             id,
		FamilyName:     familyName,
		FirstName:      firstName,
		FamilyNameKana: familyNameKana,
		FirstNameKana:  firstNameKana,
		Birthday:       birthday,
		PhoneNumber:    phoneNumber,
	}
	// personalInfoをUpdate()に代入
	personalInfo, err = controller.Interactor.Update(personalInfo)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, personalInfo)
	return
}

// endpoint DELETE /personalInfos/:id/
func (controller *PersonalInfoController) DeletePersonalInfo(c echo.Context) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))
	// idをPersonalInfo構造体のIDフィールドに格納
	personalInfo := domain.PersonalInfo{ID: id}
	// personalInfoをDeleteByPersonalInfo()に代入
	err = controller.Interactor.DeleteByPersonalInfo(personalInfo)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, personalInfo)
	return
}
