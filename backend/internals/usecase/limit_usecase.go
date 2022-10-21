package usecase

import (
	"context"
	rep "github.com/jphacks/B_2206/api/externals/repository"
	"github.com/jphacks/B_2206/api/internals/domain"
)

type limitUseCase struct {
	rep rep.LimitRepository
}

type LimitUseCase interface {
	GetLimits(context.Context) ([]domain.Limit, error)
	GetLimitByID(context.Context, string) (domain.Limit, error)
	CreateLimit(context.Context, string, string, string, string, string, string) error
	UpdateLimit(context.Context, string, string, string, string, string, string, string) error
	DestroyLimit(context.Context, string) error
	GetFundInforWithUserAndTeach(context.Context) ([]domain.FundInforWithUserAndTeacher, error)
	GetFundInforWithUserAndTeachByID(context.Context, string) (domain.FundInforWithUserAndTeacher,error)
}

func NewLimitUseCase(rep rep.LimitRepository) LimitUseCase {
	return &limitUseCase{rep}
}

//Limitsの取得(Get)
func (f *limitUseCase) GetLimits(c context.Context) ([]domain.Limit, error) {
	limit := domain.Limit{}
	var limits []domain.Limit

	rows, err := f.rep.All(c)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(
			&limit.ID,
			&limit.UserID,
			&limit.TeacherID,
			&limit.Price,
			&limit.Remark,
			&limit.IsFirstCheck,
			&limit.IsLastCheck,
			&limit.CreatedAt,
			&limit.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		limits = append(limits, limit)
	}
	return limits, nil
}

//FundInfomationの取得(Get)
func (f *limitUseCase) GetLimitByID(c context.Context, id string) (domain.Limit, error) {
	limit := domain.Limit{}

	row, err := f.rep.Find(c, id)
	err = row.Scan(
		&limit.ID,
		&limit.UserID,
		&limit.TeacherID,
		&limit.Price,
		&limit.Remark,
		&limit.IsFirstCheck,
		&limit.IsLastCheck,
		&limit.CreatedAt,
		&limit.UpdatedAt,
	)
	if err != nil {
		return limit, err
	}
	return limit, nil
}

//FundInfomationの作成(Create)
func (f *limitUseCase) CreateLimit(
	c context.Context,
	userID string,
	teacherID string,
	price string,
	remark string,
	isFirstCheck string,
	isLastCheck string,
) error {
	err := f.rep.Create(c, userID, teacherID, price, remark, isFirstCheck, isLastCheck)
	return err
}

//Limitの修正(Update)
func (f *limitUseCase) UpdateLimit(
	c context.Context,
	id string,
	userID string,
	teacherID string,
	price string,
	remark string,
	isFirstCheck string,
	isLastCheck string,
) error {
	err := f.rep.Update(c, id, userID, teacherID, price, remark, isFirstCheck, isLastCheck)
	return err
}

//funcInformationの削除(delete)
func (f *limitUseCase) DestroyLimit(c context.Context, id string) error {
	err := f.rep.Delete(c, id)
	return err
}

//fund_informations-api(GETS)
func (f *limitUseCase) GetFundInforWithUserAndTeach(c context.Context) ([]domain.FundInforWithUserAndTeacher, error) {
	fundinforuserandteacher := domain.FundInforWithUserAndTeacher{}
	var fundinforuserandteachers []domain.FundInforWithUserAndTeacher

	rows ,err := f.rep.AllWithUAndT(c)
	if err != nil {
		return nil, err
	}
	for rows.Next(){
		err := rows.Scan(
			&fundinforuserandteacher.Limit.ID,
			&fundinforuserandteacher.Limit.UserID,
			&fundinforuserandteacher.Limit.TeacherID,
			&fundinforuserandteacher.Limit.Price,
			&fundinforuserandteacher.Limit.Remark,
			&fundinforuserandteacher.Limit.IsFirstCheck,
			&fundinforuserandteacher.Limit.IsLastCheck,
			&fundinforuserandteacher.Limit.CreatedAt,
			&fundinforuserandteacher.Limit.UpdatedAt,
			&fundinforuserandteacher.User.ID,
			&fundinforuserandteacher.User.Name,
			&fundinforuserandteacher.User.BureauID,
			&fundinforuserandteacher.User.RoleID,
			&fundinforuserandteacher.User.CreatedAt,
			&fundinforuserandteacher.User.UpdatedAt,
			&fundinforuserandteacher.Teacher.ID,
			&fundinforuserandteacher.Teacher.Name,
			&fundinforuserandteacher.Teacher.Position,
			&fundinforuserandteacher.Teacher.DepartmentID,
			&fundinforuserandteacher.Teacher.Room,
			&fundinforuserandteacher.Teacher.IsBlack,
			&fundinforuserandteacher.Teacher.Remark,
			&fundinforuserandteacher.Teacher.CreatedAt,
			&fundinforuserandteacher.Teacher.UpdatedAt,
		) 
		if err != nil {
			return nil,err
		}
		fundinforuserandteachers = append(fundinforuserandteachers, fundinforuserandteacher)
	}
	return fundinforuserandteachers, nil
}

//fund_information-api(GET)
func (f *limitUseCase) GetFundInforWithUserAndTeachByID(c context.Context, id string) (domain.FundInforWithUserAndTeacher,error) {
	var fundinforuserandteacher domain.FundInforWithUserAndTeacher 

	row ,err:= f.rep.FindWithUAndT(c,id)
	
	err  = row.Scan(
		&fundinforuserandteacher.Limit.ID,
		&fundinforuserandteacher.Limit.UserID,
		&fundinforuserandteacher.Limit.TeacherID,
		&fundinforuserandteacher.Limit.Price,
		&fundinforuserandteacher.Limit.Remark,
		&fundinforuserandteacher.Limit.IsFirstCheck,
		&fundinforuserandteacher.Limit.IsLastCheck,
		&fundinforuserandteacher.Limit.CreatedAt,
		&fundinforuserandteacher.Limit.UpdatedAt,
		&fundinforuserandteacher.User.ID,
		&fundinforuserandteacher.User.Name,
		&fundinforuserandteacher.User.BureauID,
		&fundinforuserandteacher.User.RoleID,
		&fundinforuserandteacher.User.CreatedAt,
		&fundinforuserandteacher.User.UpdatedAt,
		&fundinforuserandteacher.Teacher.ID,
		&fundinforuserandteacher.Teacher.Name,
		&fundinforuserandteacher.Teacher.Position,
		&fundinforuserandteacher.Teacher.DepartmentID,
		&fundinforuserandteacher.Teacher.Room,
		&fundinforuserandteacher.Teacher.IsBlack,
		&fundinforuserandteacher.Teacher.Remark,
		&fundinforuserandteacher.Teacher.CreatedAt,
		&fundinforuserandteacher.Teacher.UpdatedAt,
	)
	if err != nil {
		return fundinforuserandteacher, err
	}
	return fundinforuserandteacher,nil
}