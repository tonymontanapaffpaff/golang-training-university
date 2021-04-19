package data

const addCourseQuery = `INSERT INTO "courses" ("code","title","department_code","description") VALUES ($1,$2,$3,$4)`
const readCourseQuery = `SELECT * FROM "courses" WHERE "courses"."code" = $1`
const readAllCoursesQuery = `SELECT * FROM "courses"`
const changeDescriptionQuery = `UPDATE "courses" SET "description"=$1 WHERE code = $2`
const deleteCourseQuery = `DELETE FROM "courses" WHERE "courses"."code" = $1`
const getDepartmentNameQuery = `SELECT departments.name FROM "courses" join departments on department_code = departments.code WHERE courses.code = $1`

const addDepartmentQuery = `INSERT INTO "departments" ("code","name") VALUES ($1,$2)`
const readDepartmentQuery = `SELECT * FROM "departments" WHERE "departments"."code" = $1`
const readAllDepartmentsQuery = `SELECT * FROM "departments"`
const changeNameQuery = `UPDATE "departments" SET "name"=$1 WHERE "code = " = $2`
const deleteDepartmentQuery = `DELETE FROM "departments" WHERE "departments"."code" = $1`

const readLecturerQuery = `SELECT * FROM "lecturers" WHERE "lecturers"."id" = $1`
const readAllLecturersQuery = `SELECT * FROM "lecturers"`
const changeFullNameQuery = `UPDATE "lecturers" SET "first_name"=$1,"last_name"=$2 WHERE id = $3`
const deleteLecturerQuery = `DELETE FROM "lecturers" WHERE "lecturers"."id" = $1`

const readStudentQuery = `SELECT * FROM "students" WHERE "students"."id" = $1`
const readAllStudentsQuery = `SELECT * FROM "students"`
const changeStatusQuery = `UPDATE students SET is_active = NOT is_active, WHERE id = $1`
const deleteStudentQuery = `DELETE FROM "students" WHERE "students"."id" = $1`
const getCurrentRateQuery = `SELECT AVG(enrollments.average_grade) FROM "students" join enrollments on id = enrollments.student_id WHERE students.id = $1`
const getCoursesListQuery = `SELECT courses.code, courses.title, courses.department_code, courses.description FROM "students" join enrollments on id = enrollments.student_id join lessons on enrollments.lesson_id = lessons.id join courses on lessons.course_code = courses.code WHERE students.id = $1`
