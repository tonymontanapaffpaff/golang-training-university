CREATE TABLE students
(
    id         int PRIMARY KEY,
    first_name varchar(20) NOT NULL,
    last_name  varchar(20) NOT NULL,
    is_active  boolean     NOT NULL
);

CREATE TABLE courses
(
    id          int PRIMARY KEY,
    title       varchar(100) NOT NULL,
    description varchar(255) NOT NULL,
    fee         varchar(255) NOT NULL
);

CREATE TABLE payments
(
    id         SERIAL PRIMARY KEY,
    student_id int     NOT NULL,
    course_id  int     NOT NULL,
    date       date    NOT NULL,
    passed     boolean NOT NULL
);

INSERT INTO students
VALUES (20174201, 'John', 'Lennon', true),
       (20174202, 'Paul', 'McCartney', false),
       (20174203, 'George', 'Harrison', false),
       (20174204, 'Ringo', 'Starr', false),
       (20174205, 'Keith', 'Richards', false);

INSERT INTO courses
VALUES (207, 'Mobile Application Development',
        'Mobile Application Development course description...',
        '100'),
       (208, 'Java Web Development',
        'Java Web Development course description...',
        '120'),
       (209, 'Architecture Operating Systems',
        'Architecture Operating Systems course description...',
        '130'),
       (202, 'General Physics',
        'General Physics course description...',
        '90'),
       (203, 'Discrete Math',
        'Discrete Math course description...',
        '90');