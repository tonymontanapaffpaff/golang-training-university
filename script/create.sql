CREATE TABLE Student
(
    student_id      char(9) PRIMARY KEY,
    student_name    varchar(20) NOT NULL,
    student_surname varchar(20) NOT NULL,
    student_status  boolean     NOT NULL
);

CREATE TABLE Department
(
    department_code char(4) PRIMARY KEY,
    department_name varchar(40) NOT NULL UNIQUE
);

CREATE TABLE Lecturer
(
    lecturer_id      char(9) PRIMARY KEY,
    lecturer_name    varchar(20) NOT NULL,
    lecturer_surname varchar(20) NOT NULL,
    department_code  char(4)     NOT NULL REFERENCES Department (department_code)
);

CREATE TABLE Location
(
    location_code    char(5) PRIMARY KEY,
    location_name    varchar(40) NOT NULL,
    location_country char(2)     NOT NULL
);

CREATE TABLE Course
(
    course_code        char(10) PRIMARY KEY,
    course_title       varchar(100) NOT NULL,
    department_code    char(4)      NOT NULL REFERENCES Department (department_code),
    course_description varchar(255) NOT NULL
);

CREATE TABLE Section
(
    section_id       SERIAL PRIMARY KEY,
    course_code      char(10) NOT NULL REFERENCES Course (course_code),
    teacher_id       char(9) REFERENCES Lecturer (lecturer_id),
    start_date       date,
    end_date         date,
    section_building int,
    section_room     varchar(4),
    section_time     time,
    location_code    char(5)  NOT NULL REFERENCES Location (location_code)
);

CREATE TABLE Enrollment
(
    student_id    char(9) REFERENCES Student (student_id),
    section_id    int REFERENCES Section (section_id),
    average_grade float,
    terminated    boolean NOT NULL,
    PRIMARY KEY (student_id, section_id)
);

CREATE TABLE Prerequisite
(
    course_code     char(10) REFERENCES Course (course_code),
    course_requires char(10) REFERENCES Course (course_code),
    PRIMARY KEY (course_code, course_requires)
);

CREATE TABLE Qualified
(
    teacher_id  char(9) REFERENCES Lecturer (lecturer_id),
    course_code char(10) REFERENCES Course (course_code),
    PRIMARY KEY (teacher_id, course_code)
);