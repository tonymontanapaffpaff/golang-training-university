package data

import (
	"database/sql"
	"errors"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DepartmentSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository DepartmentData
	course     *Department
}

func (s *DepartmentSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(postgres.New(postgres.Config{Conn: db}))
	require.NoError(s.T(), err)

	s.DB.Logger.LogMode(4)

	s.repository = DepartmentData{s.DB}
}

func (s *DepartmentSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInitDepartment(t *testing.T) {
	suite.Run(t, new(DepartmentSuite))
}

var testDepartment = &Department{
	Code: 1,
	Name: "Department 1",
}

func (s *DepartmentSuite) TestAdd() {
	assertions := assert.New(s.T())
	s.mock.ExpectExec(regexp.QuoteMeta(addDepartmentQuery)).
		WithArgs(testDepartment.Code, testDepartment.Name).
		WillReturnResult(sqlmock.NewResult(1, 1))

	departmentCode, err := s.repository.Add(*testDepartment)
	assertions.NoError(err)
	assertions.NotEmpty(departmentCode)
	assertions.Equal(testDepartment.Code, departmentCode)
}

func (s *DepartmentSuite) TestAddErr() {
	assertions := assert.New(s.T())
	s.mock.ExpectExec(regexp.QuoteMeta(addDepartmentQuery)).WillReturnError(errors.New("error"))

	departmentCode, err := s.repository.Add(*testDepartment)
	assertions.Error(err)
	assertions.Equal(-1, departmentCode)
}

func (s *DepartmentSuite) TestRead() {
	assertions := assert.New(s.T())
	rows := s.mock.NewRows([]string{"code", "name"}).
		AddRow(testDepartment.Code, testDepartment.Name)
	s.mock.ExpectQuery(regexp.QuoteMeta(readDepartmentQuery)).WithArgs(testDepartment.Code).WillReturnRows(rows)

	department, err := s.repository.Read(testDepartment.Code)
	assertions.NoError(err)
	assertions.NotEmpty(department)
	assertions.Equal(*testDepartment, department)
}

func (s *DepartmentSuite) TestReadErr() {
	assertions := assert.New(s.T())
	s.mock.ExpectQuery(regexp.QuoteMeta(readDepartmentQuery)).WillReturnError(errors.New("error"))

	department, err := s.repository.Read(testDepartment.Code)
	assertions.Error(err)
	assertions.Empty(department)
}

func (s *DepartmentSuite) TestReadAll() {
	assertions := assert.New(s.T())
	rows := s.mock.NewRows([]string{"code", "name"}).
		AddRow(testDepartment.Code, testDepartment.Name)
	s.mock.ExpectQuery(regexp.QuoteMeta(readAllDepartmentsQuery)).WillReturnRows(rows)

	departments, err := s.repository.ReadAll()
	assertions.NoError(err)
	assertions.NotEmpty(departments)
	assertions.Equal(*testDepartment, departments[0])
	assertions.Len(departments, 1)
}

func (s *DepartmentSuite) TestReadAllErr() {
	assertions := assert.New(s.T())
	s.mock.ExpectQuery(regexp.QuoteMeta(readAllDepartmentsQuery)).WillReturnError(errors.New("error"))

	departments, err := s.repository.ReadAll()
	assertions.Error(err)
	assertions.Empty(departments)
}

func (s *DepartmentSuite) TestChangeName() {
	assertions := assert.New(s.T())

	name := "another name"
	s.mock.ExpectExec(regexp.QuoteMeta(changeNameQuery)).
		WithArgs(name, testDepartment.Code).
		WillReturnResult(sqlmock.NewResult(1, 1))

	departmentCode, err := s.repository.ChangeName(testDepartment.Code, name)
	assertions.NoError(err)
	assertions.NotEmpty(departmentCode)
	assertions.Equal(testDepartment.Code, departmentCode)
}

func (s *DepartmentSuite) TestChangeNameErr() {
	assertions := assert.New(s.T())

	name := "another name"
	s.mock.ExpectExec(regexp.QuoteMeta(changeNameQuery)).WillReturnError(errors.New("error"))

	departmentCode, err := s.repository.ChangeName(testDepartment.Code, name)
	assertions.Error(err)
	assertions.Equal(-1, departmentCode)
}

func (s *DepartmentSuite) TestDelete() {
	assertions := assert.New(s.T())
	s.mock.ExpectExec(regexp.QuoteMeta(deleteDepartmentQuery)).
		WithArgs(testDepartment.Code).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.Delete(testDepartment.Code)
	assertions.NoError(err)
}

func (s *DepartmentSuite) TestDeleteErr() {
	assertions := assert.New(s.T())
	s.mock.ExpectExec(regexp.QuoteMeta(deleteDepartmentQuery)).WillReturnError(errors.New("error"))

	err := s.repository.Delete(testDepartment.Code)
	assertions.Error(err)
}
