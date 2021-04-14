INSERT INTO student
VALUES ('PMS201701', 'John', 'Lennon', true),
       ('PMS201702', 'Paul', 'McCartney', true),
       ('PMS201703', 'George', 'Harrison', true),
       ('PMS201704', 'Ringo', 'Starr', true),
       ('PMS201705', 'Keith', 'Richards', true);

INSERT INTO location
VALUES ('GSUBY', 'Francysk Skoryna Gomel State University', 'BY');

INSERT INTO department
VALUES ('PHYS', 'Faculty of Physics and IT');

INSERT INTO course
VALUES ('PHYS-02-08', 'Mobile Application Development',
        (SELECT department_code FROM department WHERE department_code = 'PHYS'),
        'Mobile Application Development course description...'),
       ('PHYS-02-07', 'Java Web Development',
        (SELECT department_code FROM department WHERE department_code = 'PHYS'),
        'Java Web Development course description...'),
       ('PHYS-02-09', 'Architecture Operating Systems',
        (SELECT department_code FROM department WHERE department_code = 'PHYS'),
        'Architecture Operating Systems course description...'),
       ('PHYS-02-02', 'General Physics',
        (SELECT department_code FROM department WHERE department_code = 'PHYS'),
        'General Physics course description...'),
       ('PHYS-02-03', 'Discrete Math',
        (SELECT department_code FROM department WHERE department_code = 'PHYS'),
        'Discrete Math course description...');

INSERT INTO prerequisite
VALUES ((SELECT course_code FROM course WHERE course_code = 'PHYS-02-07'),
        (SELECT course_code FROM course WHERE course_code = 'PHYS-02-02')),
       ((SELECT course_code FROM course WHERE course_code = 'PHYS-02-07'),
        (SELECT course_code FROM course WHERE course_code = 'PHYS-02-03')),
       ((SELECT course_code FROM course WHERE course_code = 'PHYS-02-08'),
        (SELECT course_code FROM course WHERE course_code = 'PHYS-02-02')),
       ((SELECT course_code FROM course WHERE course_code = 'PHYS-02-08'),
        (SELECT course_code FROM course WHERE course_code = 'PHYS-02-03')),
       ((SELECT course_code FROM course WHERE course_code = 'PHYS-02-09'),
        (SELECT course_code FROM course WHERE course_code = 'PHYS-02-02')),
       ((SELECT course_code FROM course WHERE course_code = 'PHYS-02-09'),
        (SELECT course_code FROM course WHERE course_code = 'PHYS-02-03'));

INSERT INTO lecturer
VALUES ('PHY200123', 'Ella', 'Fitzgerald', (SELECT department_code FROM department WHERE department_code = 'PHYS')),
       ('PHY203127', 'Frank', 'Sinatra', (SELECT department_code FROM department WHERE department_code = 'PHYS')),
       ('PHY199355', 'Nat', 'Cole', (SELECT department_code FROM department WHERE department_code = 'PHYS')),
       ('PHY200699', 'Billie', 'Holiday', (SELECT department_code FROM department WHERE department_code = 'PHYS')),
       ('PHY199988', 'Sarah', 'Vaughan', (SELECT department_code FROM department WHERE department_code = 'PHYS'));

INSERT INTO qualified
VALUES ((SELECT lecturer_id FROM lecturer WHERE lecturer_id = 'PHY199355'),
        (SELECT course_code FROM course WHERE course_code = 'PHYS-02-02')),
       ((SELECT lecturer_id FROM lecturer WHERE lecturer_id = 'PHY199988'),
        (SELECT course_code FROM course WHERE course_code = 'PHYS-02-03')),
       ((SELECT lecturer_id FROM lecturer WHERE lecturer_id = 'PHY200123'),
        (SELECT course_code FROM course WHERE course_code = 'PHYS-02-07')),
       ((SELECT lecturer_id FROM lecturer WHERE lecturer_id = 'PHY200699'),
        (SELECT course_code FROM course WHERE course_code = 'PHYS-02-08')),
       ((SELECT lecturer_id FROM lecturer WHERE lecturer_id = 'PHY203127'),
        (SELECT course_code FROM course WHERE course_code = 'PHYS-02-09'));

INSERT INTO section (course_code,
                     teacher_id,
                     start_date,
                     end_date,
                     section_building,
                     section_room,
                     section_time,
                     location_code)
VALUES ((SELECT course_code FROM course WHERE course_code = 'PHYS-02-02'),
        (SELECT lecturer_id FROM lecturer WHERE lecturer_id = 'PHY199355'),
        '2020-09-03', '2020-12-25', 5, '4-28', '9:00',
        (SELECT location_code FROM location WHERE location_code = 'GSUBY')),
       ((SELECT course_code FROM course WHERE course_code = 'PHYS-02-03'),
        (SELECT lecturer_id FROM lecturer WHERE lecturer_id = 'PHY199988'),
        '2020-09-03', '2020-12-25', 5, '3-4', '10:55',
        (SELECT location_code FROM location WHERE location_code = 'GSUBY')),
       ((SELECT course_code FROM course WHERE course_code = 'PHYS-02-07'),
        (SELECT lecturer_id FROM lecturer WHERE lecturer_id = 'PHY200123'),
        '2020-09-03', '2020-12-27', 5, '4-8', '9:00',
        (SELECT location_code FROM location WHERE location_code = 'GSUBY')),
       ((SELECT course_code FROM course WHERE course_code = 'PHYS-02-08'),
        (SELECT lecturer_id FROM lecturer WHERE lecturer_id = 'PHY200699'),
        '2020-09-03', '2020-12-20', 5, '2-9', '12:25',
        (SELECT location_code FROM location WHERE location_code = 'GSUBY')),
       ((SELECT course_code FROM course WHERE course_code = 'PHYS-02-09'),
        (SELECT lecturer_id FROM lecturer WHERE lecturer_id = 'PHY203127'),
        '2020-09-03', '2020-12-30', 5, '4-1', '10:55',
        (SELECT location_code FROM location WHERE location_code = 'GSUBY'));

INSERT INTO enrollment
VALUES ((SELECT student_id FROM student WHERE student_id = 'PMS201701'),
        (SELECT section_id FROM section WHERE section_id = 1),
        7.6, false),
       ((SELECT student_id FROM student WHERE student_id = 'PMS201701'),
        (SELECT section_id FROM section WHERE section_id = 2),
        8.5, false),
       ((SELECT student_id FROM student WHERE student_id = 'PMS201701'),
        (SELECT section_id FROM section WHERE section_id = 3),
        9.2, false),
       ((SELECT student_id FROM student WHERE student_id = 'PMS201701'),
        (SELECT section_id FROM section WHERE section_id = 4),
        9.5, false),
       ((SELECT student_id FROM student WHERE student_id = 'PMS201701'),
        (SELECT section_id FROM section WHERE section_id = 5),
        9, false);

INSERT INTO enrollment
VALUES ((SELECT student_id FROM student WHERE student_id = 'PMS201702'),
        (SELECT section_id FROM section WHERE section_id = 1),
        6.5, false),
       ((SELECT student_id FROM student WHERE student_id = 'PMS201702'),
        (SELECT section_id FROM section WHERE section_id = 2),
        7.5, false),
       ((SELECT student_id FROM student WHERE student_id = 'PMS201702'),
        (SELECT section_id FROM section WHERE section_id = 3),
        8.5, false),
       ((SELECT student_id FROM student WHERE student_id = 'PMS201702'),
        (SELECT section_id FROM section WHERE section_id = 4),
        9.5, false),
       ((SELECT student_id FROM student WHERE student_id = 'PMS201702'),
        (SELECT section_id FROM section WHERE section_id = 5),
        6, false);

INSERT INTO enrollment
VALUES ((SELECT student_id FROM student WHERE student_id = 'PMS201703'),
        (SELECT section_id FROM section WHERE section_id = 1),
        4.5, false),
       ((SELECT student_id FROM student WHERE student_id = 'PMS201703'),
        (SELECT section_id FROM section WHERE section_id = 2),
        6.5, false),
       ((SELECT student_id FROM student WHERE student_id = 'PMS201703'),
        (SELECT section_id FROM section WHERE section_id = 3),
        8.3, false),
       ((SELECT student_id FROM student WHERE student_id = 'PMS201703'),
        (SELECT section_id FROM section WHERE section_id = 4),
        8.4, false),
       ((SELECT student_id FROM student WHERE student_id = 'PMS201703'),
        (SELECT section_id FROM section WHERE section_id = 5),
        7.2, false);

INSERT INTO enrollment
VALUES ((SELECT student_id FROM student WHERE student_id = 'PMS201704'),
        (SELECT section_id FROM section WHERE section_id = 1),
        7.2, false),
       ((SELECT student_id FROM student WHERE student_id = 'PMS201704'),
        (SELECT section_id FROM section WHERE section_id = 2),
        9.7, false),
       ((SELECT student_id FROM student WHERE student_id = 'PMS201704'),
        (SELECT section_id FROM section WHERE section_id = 3),
        6.4, false),
       ((SELECT student_id FROM student WHERE student_id = 'PMS201704'),
        (SELECT section_id FROM section WHERE section_id = 4),
        6.7, false),
       ((SELECT student_id FROM student WHERE student_id = 'PMS201704'),
        (SELECT section_id FROM section WHERE section_id = 5),
        8.8, false);

INSERT INTO enrollment
VALUES ((SELECT student_id FROM student WHERE student_id = 'PMS201705'),
        (SELECT section_id FROM section WHERE section_id = 1),
        7.1, false),
       ((SELECT student_id FROM student WHERE student_id = 'PMS201705'),
        (SELECT section_id FROM section WHERE section_id = 2),
        8.3, false),
       ((SELECT student_id FROM student WHERE student_id = 'PMS201705'),
        (SELECT section_id FROM section WHERE section_id = 3),
        9.2, false),
       ((SELECT student_id FROM student WHERE student_id = 'PMS201705'),
        (SELECT section_id FROM section WHERE section_id = 4),
        9.3, false);