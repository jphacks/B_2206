package controller

import (
	"github.com/jphacks/B_2206/api/internals/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type companyInfoController struct {
	u usecase.CompanyInfoUseCase
}

type CompanyInfoController interface {
	IndexCompanyInfo(echo.Context) error
	ShowCompanyInfo(echo.Context) error
	CreateCompanyInfo(echo.Context) error
	UpdateCompanyInfo(echo.Context) error
	DestroyCompanyInfo(echo.Context) error
}

func NewCompanyInfoController(u usecase.CompanyInfoUseCase) CompanyInfoController{
	return &companyInfoController{u}
}

func (b *companyInfoController) IndexCompanyInfo(c echo.Context) error {
	companyInfos ,err := b.u.GetCompanyInfos(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, companyInfos)
}

func (b *companyInfoController) ShowCompanyInfo(c echo.Context) error{
	id := c.Param("id")
	companyInfo ,err :=b.u.GetCompanyInfoByID(c.Request().Context(),id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK,companyInfo)
}

func (b *companyInfoController) CreateCompanyInfo(c echo.Context) error{
	name :=c.QueryParam("name")
	err := b.u.CreateCompanyInfo(c.Request().Context(),name)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK,"Created CompanyInfo")
} 

func (b *companyInfoController) UpdateCompanyInfo(c echo.Context) error{
	id := c.Param("id")
	name := c.QueryParam("name")
	err := b.u.UpdateCompanyInfo(c.Request().Context(),id,name)

	if err != nil {
		return err
	}
	return c.String(http.StatusOK,"Update CompanyInfo")
}

func (b *companyInfoController) DestroyCompanyInfo(c echo.Context) error{
	id  := c.Param("id")
	err := b.u.DestroyCompanyInfo(c.Request().Context(),id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Destroy CompanyInfo")
}
