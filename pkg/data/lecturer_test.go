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

type LecturerSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository LecturerData
	course     *Lecturer
}

func (s *LecturerSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(postgres.New(postgres.Config{Conn: db}))
	require.NoError(s.T(), err)

	s.DB.Logger.LogMode(4)

	s.repository = LecturerData{s.DB}
}

func (s *LecturerSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInitLecturer(t *testing.T) {
	suite.Run(t, new(LecturerSuite))
}

var (
	testLecturer = &Lecturer{
		Id:             1,
		FirstName:      "Alex",
		LastName:       "Caruso",
		DepartmentCode: 1,
	}
	FirstName = "The"
	lastName  = "GOAT"
)

func (s *LecturerSuite) TestRead() {
	assertions := assert.New(s.T())
	rows := s.mock.NewRows([]string{"id", "first_name", "last_name", "department_code"}).
		AddRow(testLecturer.Id, testLecturer.FirstName, testLecturer.LastName, testLecturer.DepartmentCode)
	s.mock.ExpectQuery(regexp.QuoteMeta(readLecturerQuery)).WithArgs(testLecturer.Id).WillReturnRows(rows)

	lecturer, err := s.repository.Read(testLecturer.Id)
	assertions.NoError(err)
	assertions.NotEmpty(lecturer)
	assertions.Equal(*testLecturer, lecturer)
}

func (s *LecturerSuite) TestReadErr() {
	assertions := assert.New(s.T())
	s.mock.ExpectQuery(regexp.QuoteMeta(readLecturerQuery)).WillReturnError(errors.New("error"))

	lecturer, err := s.repository.Read(testLecturer.Id)
	assertions.Error(err)
	assertions.Empty(lecturer)
}

func (s *LecturerSuite) TestReadAll() {
	assertions := assert.New(s.T())
	rows := s.mock.NewRows([]string{"id", "first_name", "last_name", "department_code"}).
		AddRow(testLecturer.Id, testLecturer.FirstName, testLecturer.LastName, testLecturer.DepartmentCode)
	s.mock.ExpectQuery(regexp.QuoteMeta(readAllLecturersQuery)).WillReturnRows(rows)

	lecturers, err := s.repository.ReadAll()
	assertions.NoError(err)
	assertions.NotEmpty(lecturers)
	assertions.Equal(*testLecturer, lecturers[0])
	assertions.Len(lecturers, 1)
}

func (s *LecturerSuite) TestReadAllErr() {
	assertions := assert.New(s.T())
	s.mock.ExpectQuery(regexp.QuoteMeta(readAllLecturersQuery)).WillReturnError(errors.New("error"))

	lecturers, err := s.repository.ReadAll()
	assertions.Error(err)
	assertions.Empty(lecturers)
}

func (s *LecturerSuite) TestChangeFullName() {
	assertions := assert.New(s.T())

	s.mock.ExpectExec(regexp.QuoteMeta(changeFullNameQuery)).
		WithArgs(FirstName, lastName, testLecturer.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	lecturerId, err := s.repository.ChangeFullName(testLecturer.Id, FirstName, lastName)
	assertions.NoError(err)
	assertions.NotEmpty(lecturerId)
	assertions.Equal(testLecturer.Id, lecturerId)
}

func (s *LecturerSuite) TestChangeFullNameErr() {
	assertions := assert.New(s.T())

	s.mock.ExpectExec(regexp.QuoteMeta(changeFullNameQuery)).WillReturnError(errors.New("error"))

	departmentCode, err := s.repository.ChangeFullName(testLecturer.Id, FirstName, lastName)
	assertions.Error(err)
	assertions.Equal(-1, departmentCode)
}

func (s *LecturerSuite) TestDelete() {
	assertions := assert.New(s.T())
	s.mock.ExpectExec(regexp.QuoteMeta(deleteLecturerQuery)).
		WithArgs(testLecturer.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.Delete(testLecturer.Id)
	assertions.NoError(err)
}

func (s *LecturerSuite) TestDeleteErr() {
	assertions := assert.New(s.T())
	s.mock.ExpectExec(regexp.QuoteMeta(deleteLecturerQuery)).WillReturnError(errors.New("error"))

	err := s.repository.Delete(testLecturer.Id)
	assertions.Error(err)
}
