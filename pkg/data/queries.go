package data

const (
	addCourseQuery          = `INSERT INTO "courses" ("code","title","department_code","description") VALUES ($1,$2,$3,$4)`
	readCourseQuery         = `SELECT * FROM "courses" WHERE "courses"."code" = $1`
	readAllCoursesQuery     = `SELECT * FROM "courses"`
	changeDescriptionQuery  = `UPDATE "courses" SET "description"=$1 WHERE code = $2`
	deleteCourseQuery       = `DELETE FROM "courses" WHERE "courses"."code" = $1`
	getDepartmentNameQuery  = `SELECT departments.name FROM "courses" join departments on department_code = departments.code WHERE courses.code = $1`
	addDepartmentQuery      = `INSERT INTO "departments" ("code","name") VALUES ($1,$2)`
	readDepartmentQuery     = `SELECT * FROM "departments" WHERE "departments"."code" = $1`
	readAllDepartmentsQuery = `SELECT * FROM "departments"`
	changeNameQuery         = `UPDATE "departments" SET "name"=$1 WHERE "code = " = $2`
	deleteDepartmentQuery   = `DELETE FROM "departments" WHERE "departments"."code" = $1`
	readLecturerQuery       = `SELECT * FROM "lecturers" WHERE "lecturers"."id" = $1`
	readAllLecturersQuery   = `SELECT * FROM "lecturers"`
	changeFullNameQuery     = `UPDATE "lecturers" SET "first_name"=$1,"last_name"=$2 WHERE id = $3`
	deleteLecturerQuery     = `DELETE FROM "lecturers" WHERE "lecturers"."id" = $1`
	readStudentQuery        = `SELECT * FROM "students" WHERE "students"."id" = $1`
	readAllStudentsQuery    = `SELECT * FROM "students"`
	changeStatusQuery       = `UPDATE students SET is_active = NOT is_active WHERE id = $1`
	deleteStudentQuery      = `DELETE FROM "students" WHERE "students"."id" = $1`
	getCurrentRateQuery     = `SELECT AVG(enrollments.average_grade) FROM "students" join enrollments on id = enrollments.student_id WHERE students.id = $1`
	getCoursesListQuery     = `SELECT courses.code, courses.title, courses.department_code, courses.description FROM "students" join enrollments on id = enrollments.student_id join lessons on enrollments.lesson_id = lessons.id join courses on lessons.course_code = courses.code WHERE students.id = $1`
)
