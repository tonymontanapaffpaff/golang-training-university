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

type StudentSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository StudentData
	course     *Student
}

func (s *StudentSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(postgres.New(postgres.Config{Conn: db}))
	require.NoError(s.T(), err)

	s.DB.Logger.LogMode(4)

	s.repository = StudentData{s.DB}
}

func (s *StudentSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInitStudent(t *testing.T) {
	suite.Run(t, new(StudentSuite))
}

var testStudent = &Student{
	Id:        1,
	FirstName: "Andre",
	LastName:  "Drummond",
	IsActive:  true,
}

func (s *StudentSuite) TestRead() {
	assertions := assert.New(s.T())
	rows := s.mock.NewRows([]string{"id", "first_name", "last_name", "is_active"}).
		AddRow(testStudent.Id, testStudent.FirstName, testStudent.LastName, testStudent.IsActive)
	s.mock.ExpectQuery(regexp.QuoteMeta(readStudentQuery)).WithArgs(testStudent.Id).WillReturnRows(rows)

	student, err := s.repository.Read(testStudent.Id)
	assertions.NoError(err)
	assertions.NotEmpty(student)
	assertions.Equal(*testStudent, student)
}

func (s *StudentSuite) TestReadErr() {
	assertions := assert.New(s.T())
	s.mock.ExpectQuery(regexp.QuoteMeta(readStudentQuery)).WillReturnError(errors.New("error"))

	student, err := s.repository.Read(testStudent.Id)
	assertions.Error(err)
	assertions.Empty(student)
}

func (s *StudentSuite) TestReadAll() {
	assertions := assert.New(s.T())
	rows := s.mock.NewRows([]string{"id", "first_name", "last_name", "is_active"}).
		AddRow(testStudent.Id, testStudent.FirstName, testStudent.LastName, testStudent.IsActive)
	s.mock.ExpectQuery(regexp.QuoteMeta(readAllStudentsQuery)).WillReturnRows(rows)

	students, err := s.repository.ReadAll()
	assertions.NoError(err)
	assertions.NotEmpty(students)
	assertions.Equal(*testStudent, students[0])
	assertions.Len(students, 1)
}

func (s *StudentSuite) TestReadAllErr() {
	assertions := assert.New(s.T())
	s.mock.ExpectQuery(regexp.QuoteMeta(readAllStudentsQuery)).WillReturnError(errors.New("error"))

	students, err := s.repository.ReadAll()
	assertions.Error(err)
	assertions.Empty(students)
}

func (s *StudentSuite) TestChangeStatus() {
	assertions := assert.New(s.T())

	s.mock.ExpectExec(regexp.QuoteMeta(changeStatusQuery)).
		WithArgs(testStudent.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	studentId, err := s.repository.ChangeStatus(testStudent.Id)
	assertions.NoError(err)
	assertions.NotEmpty(studentId)
	assertions.Equal(testStudent.Id, studentId)
}

func (s *StudentSuite) TestChangeStatusErr() {
	assertions := assert.New(s.T())

	s.mock.ExpectExec(regexp.QuoteMeta(changeStatusQuery)).WillReturnError(errors.New("error"))

	studentId, err := s.repository.ChangeStatus(testStudent.Id)
	assertions.Error(err)
	assertions.Equal(-1, studentId)
}

func (s *StudentSuite) TestDelete() {
	assertions := assert.New(s.T())
	s.mock.ExpectExec(regexp.QuoteMeta(deleteStudentQuery)).
		WithArgs(testStudent.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.Delete(testStudent.Id)
	assertions.NoError(err)
}

func (s *StudentSuite) TestDeleteErr() {
	assertions := assert.New(s.T())
	s.mock.ExpectExec(regexp.QuoteMeta(deleteStudentQuery)).WillReturnError(errors.New("error"))

	err := s.repository.Delete(testStudent.Id)
	assertions.Error(err)
}

func (s *StudentSuite) TestCurrentRate() {
	assertions := assert.New(s.T())

	rows := s.mock.NewRows([]string{"avg"}).AddRow(1)
	s.mock.ExpectQuery(regexp.QuoteMeta(getCurrentRateQuery)).WithArgs(testStudent.Id).WillReturnRows(rows)

	averageMark, err := s.repository.GetCurrentRate(testStudent.Id)
	assertions.NoError(err)
	assertions.NotEmpty(averageMark)
}

func (s *StudentSuite) TestCurrentRateErr() {
	assertions := assert.New(s.T())
	s.mock.ExpectQuery(regexp.QuoteMeta(getCurrentRateQuery)).WillReturnError(errors.New("error"))

	averageMark, err := s.repository.GetCurrentRate(testStudent.Id)
	assertions.Error(err)
	assertions.Equal(float64(-1), averageMark)
}

func (s *StudentSuite) TestGetCoursesList() {
	assertions := assert.New(s.T())

	rows := s.mock.NewRows([]string{"courses.code, courses.title, courses.department_code, courses.description"}).AddRow(1)
	s.mock.ExpectQuery(regexp.QuoteMeta(getCoursesListQuery)).WithArgs(testStudent.Id).WillReturnRows(rows)

	coursesList, err := s.repository.GetCoursesList(testStudent.Id)
	assertions.NoError(err)
	assertions.NotEmpty(coursesList)
	assertions.Len(coursesList, 1)
}

func (s *StudentSuite) TestGetCoursesListErr() {
	assertions := assert.New(s.T())
	s.mock.ExpectQuery(regexp.QuoteMeta(getCoursesListQuery)).WillReturnError(errors.New("error"))

	coursesList, err := s.repository.GetCoursesList(testStudent.Id)
	assertions.Error(err)
	assertions.Empty(coursesList)
}
