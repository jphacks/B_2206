package controllers

import (
	"github.com/jphacks/B_2206/backend/domain"
	"github.com/jphacks/B_2206/backend/interfaces/database"
	"github.com/jphacks/B_2206/backend/usecase"

	"strconv"

	"github.com/labstack/echo/v4"
)

type TagController struct {
	Interactor usecase.TagInteractor
}

func NewTagController(sqlHandler database.SqlHandler) *TagController {
	return &TagController{
		Interactor: usecase.TagInteractor{
			TagRepository: &database.TagRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// endpoint POST /tags
func (controller *TagController) CreateTag(c echo.Context, name string) (err error) {
	u := domain.Tag{
		Name: name,
	}
	c.Bind(&u)
	tag, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, tag)
	return
}

// endpoint GET /tags/:id/
func (controller *TagController) GetTag(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	tag, err := controller.Interactor.TagById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, tag)
	return
}

// endpoint GET /tags/
func (controller *TagController) GetTags(c echo.Context) (err error) {
	tags, err := controller.Interactor.Tags()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, tags)
	return
}

// endpoint UPDATE /tags/:id/
func (controller *TagController) UpdateTag(c echo.Context, name string) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))

	// idをTag構造体のIDフィールドに格納
	tag := domain.Tag{
		ID:   id,
		Name: name,
	}
	// tagをUpdate()に代入
	tag, err = controller.Interactor.Update(tag)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, tag)
	return
}

// endpoint DELETE /tags/:id/
func (controller *TagController) DeleteTag(c echo.Context) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))
	// idをTag構造体のIDフィールドに格納
	tag := domain.Tag{ID: id}
	// tagをDeleteByTag()に代入
	err = controller.Interactor.DeleteByTag(tag)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, tag)
	return
}
