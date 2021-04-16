CREATE TABLE Students
(
    id         int PRIMARY KEY,
    first_name varchar(20) NOT NULL,
    last_name  varchar(20) NOT NULL,
    is_active  boolean     NOT NULL
);

CREATE TABLE Departments
(
    code int PRIMARY KEY,
    name varchar(40) NOT NULL UNIQUE
);

CREATE TABLE Lecturers
(
    id              int PRIMARY KEY,
    first_name      varchar(20) NOT NULL,
    last_name       varchar(20) NOT NULL,
    department_code int         NOT NULL REFERENCES Departments (code)
);

CREATE TABLE Locations
(
    code    int PRIMARY KEY,
    name    varchar(40) NOT NULL,
    country char(2)     NOT NULL
);

CREATE TABLE Courses
(
    code            int PRIMARY KEY,
    title           varchar(100) NOT NULL,
    department_code int          NOT NULL REFERENCES Departments (code),
    description     varchar(255) NOT NULL
);

CREATE TABLE Lessons
(
    id            SERIAL PRIMARY KEY,
    course_code   int NOT NULL REFERENCES Courses (code),
    lecturer_id   int REFERENCES Lecturers (id),
    start_date    date,
    end_date      date,
    building      int,
    room          varchar(4),
    start_time    time,
    location_code int NOT NULL REFERENCES Locations (code)
);

CREATE TABLE Enrollments
(
    student_id    int REFERENCES Students (id),
    lesson_id     int REFERENCES Lessons (id),
    average_grade float,
    is_terminated boolean NOT NULL,
    PRIMARY KEY (student_id, lesson_id)
);

CREATE TABLE Prerequisites
(
    course_code     int REFERENCES Courses (code),
    course_requires int REFERENCES Courses (code),
    PRIMARY KEY (course_code, course_requires)
);

CREATE TABLE Qualifications
(
    teacher_id  int REFERENCES Lecturers (id),
    course_code int REFERENCES Courses (code),
    PRIMARY KEY (teacher_id, course_code)
);