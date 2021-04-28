package data

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CourseSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository CourseData
	course     *Course
}

func (s *CourseSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(postgres.New(postgres.Config{Conn: db}))
	require.NoError(s.T(), err)

	s.DB.Logger.LogMode(4)

	s.repository = CourseData{s.DB}
}

func (s *CourseSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInitCourse(t *testing.T) {
	suite.Run(t, new(CourseSuite))
}

var testCourse = &Course{
	Code:           1,
	Title:          "Course name",
	DepartmentCode: 1,
	Description:    "Course description",
}

func (s *CourseSuite) TestAdd() {
	assertions := assert.New(s.T())
	s.mock.ExpectExec(regexp.QuoteMeta(addCourseQuery)).
		WithArgs(testCourse.Code, testCourse.Title, testCourse.DepartmentCode, testCourse.Description).
		WillReturnResult(sqlmock.NewResult(1, 1))

	courseCode, err := s.repository.Add(*testCourse)
	assertions.NoError(err)
	assertions.NotEmpty(courseCode)
	assertions.Equal(testCourse.Code, courseCode)
}

func (s *CourseSuite) TestAddErr() {
	assertions := assert.New(s.T())
	s.mock.ExpectExec(regexp.QuoteMeta(addCourseQuery)).WillReturnError(errors.New("error"))

	courseCode, err := s.repository.Add(*testCourse)
	assertions.Error(err)
	assertions.Equal(-1, courseCode)
}

func (s *CourseSuite) TestRead() {
	assertions := assert.New(s.T())
	rows := s.mock.NewRows([]string{"code", "title", "department_code", "description"}).
		AddRow(testCourse.Code, testCourse.Title, testCourse.DepartmentCode, testCourse.Description)
	s.mock.ExpectQuery(regexp.QuoteMeta(readCourseQuery)).WithArgs(testCourse.Code).WillReturnRows(rows)

	course, err := s.repository.Read(testCourse.Code)
	assertions.NoError(err)
	assertions.NotEmpty(course)
	assertions.Equal(*testCourse, course)
}

func (s *CourseSuite) TestReadErr() {
	assertions := assert.New(s.T())
	s.mock.ExpectQuery(regexp.QuoteMeta(readCourseQuery)).WillReturnError(errors.New("error"))

	course, err := s.repository.Read(testCourse.Code)
	assertions.Error(err)
	assertions.Empty(course)
}

func (s *CourseSuite) TestReadAll() {
	assertions := assert.New(s.T())
	rows := s.mock.NewRows([]string{"code", "title", "department_code", "description"}).
		AddRow(testCourse.Code, testCourse.Title, testCourse.DepartmentCode, testCourse.Description)
	s.mock.ExpectQuery(regexp.QuoteMeta(readAllCoursesQuery)).WillReturnRows(rows)

	courses, err := s.repository.ReadAll()
	assertions.NoError(err)
	assertions.NotEmpty(courses)
	assertions.Equal(*testCourse, courses[0])
	assertions.Len(courses, 1)
}

func (s *CourseSuite) TestReadAllErr() {
	assertions := assert.New(s.T())
	s.mock.ExpectQuery(regexp.QuoteMeta(readAllCoursesQuery)).WillReturnError(errors.New("error"))

	courses, err := s.repository.ReadAll()
	assertions.Error(err)
	assertions.Empty(courses)
}

func (s *CourseSuite) TestChangeDescription() {
	assertions := assert.New(s.T())

	description := "another description"
	s.mock.ExpectExec(regexp.QuoteMeta(changeDescriptionQuery)).
		WithArgs(description, testCourse.Code).
		WillReturnResult(sqlmock.NewResult(1, 1))

	courseCode, err := s.repository.ChangeDescription(testCourse.Code, description)
	assertions.NoError(err)
	assertions.NotEmpty(courseCode)
	assertions.Equal(testCourse.Code, courseCode)
}

func (s *CourseSuite) TestChangeDescriptionErr() {
	assertions := assert.New(s.T())

	description := "another description"
	s.mock.ExpectExec(regexp.QuoteMeta(changeDescriptionQuery)).WillReturnError(errors.New("error"))

	courseCode, err := s.repository.ChangeDescription(testCourse.Code, description)
	assertions.Error(err)
	assertions.Equal(-1, courseCode)
}

func (s *CourseSuite) TestDelete() {
	assertions := assert.New(s.T())
	s.mock.ExpectExec(regexp.QuoteMeta(deleteCourseQuery)).
		WithArgs(testCourse.Code).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.Delete(testCourse.Code)
	assertions.NoError(err)
}

func (s *CourseSuite) TestDeleteErr() {
	assertions := assert.New(s.T())
	s.mock.ExpectExec(regexp.QuoteMeta(deleteCourseQuery)).WillReturnError(errors.New("error"))

	err := s.repository.Delete(testCourse.Code)
	assertions.Error(err)
}

func (s *CourseSuite) TestGetDepartmentName() {
	assertions := assert.New(s.T())
	expectedName := "Department 1"

	rows := s.mock.NewRows([]string{"actualName"}).AddRow(expectedName)
	s.mock.ExpectQuery(regexp.QuoteMeta(getDepartmentNameQuery)).WithArgs(testCourse.Code).WillReturnRows(rows)

	actualName, err := s.repository.GetDepartmentName(testCourse.Code)
	assertions.NoError(err)
	assertions.NotEmpty(actualName)
	assertions.Equal(expectedName, actualName)
}

func (s *CourseSuite) TestGetDepartmentNameErr() {
	assertions := assert.New(s.T())
	s.mock.ExpectQuery(regexp.QuoteMeta(getDepartmentNameQuery)).WillReturnError(errors.New("error"))

	name, err := s.repository.GetDepartmentName(testCourse.Code)
	assertions.Error(err)
	assertions.Empty(name)
}
