INSERT INTO students
VALUES (20174201, 'John', 'Lennon', true),
       (20174202, 'Paul', 'McCartney', true),
       (20174203, 'George', 'Harrison', true),
       (20174204, 'Ringo', 'Starr', true),
       (20174205, 'Keith', 'Richards', true);

INSERT INTO locations
VALUES (1, 'Francysk Skoryna Gomel State University', 'BY');

INSERT INTO departments
VALUES (5, 'Faculty of Physics and IT');

INSERT INTO courses
VALUES (207, 'Mobile Application Development',
        (SELECT code FROM departments WHERE code = 5),
        'Mobile Application Development course description...'),
       (208, 'Java Web Development',
        (SELECT code FROM departments WHERE code = 5),
        'Java Web Development course description...'),
       (209, 'Architecture Operating Systems',
        (SELECT code FROM departments WHERE code = 5),
        'Architecture Operating Systems course description...'),
       (202, 'General Physics',
        (SELECT code FROM departments WHERE code = 5),
        'General Physics course description...'),
       (203, 'Discrete Math',
        (SELECT code FROM departments WHERE code = 5),
        'Discrete Math course description...');

INSERT INTO prerequisites
VALUES ((SELECT code FROM courses WHERE code = 207),
        (SELECT code FROM courses WHERE code = 202)),
       ((SELECT code FROM courses WHERE code = 207),
        (SELECT code FROM courses WHERE code = 203)),
       ((SELECT code FROM courses WHERE code = 208),
        (SELECT code FROM courses WHERE code = 202)),
       ((SELECT code FROM courses WHERE code = 208),
        (SELECT code FROM courses WHERE code = 203)),
       ((SELECT code FROM courses WHERE code = 209),
        (SELECT code FROM courses WHERE code = 202)),
       ((SELECT code FROM courses WHERE code = 209),
        (SELECT code FROM courses WHERE code = 203));

INSERT INTO lecturers
VALUES (200123, 'Ella', 'Fitzgerald', (SELECT code FROM departments WHERE code = 5)),
       (203127, 'Frank', 'Sinatra', (SELECT code FROM departments WHERE code = 5)),
       (199355, 'Nat', 'Cole', (SELECT code FROM departments WHERE code = 5)),
       (200699, 'Billie', 'Holiday', (SELECT code FROM departments WHERE code = 5)),
       (199988, 'Sarah', 'Vaughan', (SELECT code FROM departments WHERE code = 5));

INSERT INTO qualifications
VALUES ((SELECT id FROM lecturers WHERE id = 199355),
        (SELECT code FROM courses WHERE code = 202)),
       ((SELECT id FROM lecturers WHERE id = 199988),
        (SELECT code FROM courses WHERE code = 203)),
       ((SELECT id FROM lecturers WHERE id = 200123),
        (SELECT code FROM courses WHERE code = 207)),
       ((SELECT id FROM lecturers WHERE id = 200699),
        (SELECT code FROM courses WHERE code = 208)),
       ((SELECT id FROM lecturers WHERE id = 203127),
        (SELECT code FROM courses WHERE code = 209));

INSERT INTO lessons (course_code,
                     lecturer_id,
                     start_date,
                     end_date,
                     building,
                     room,
                     start_time,
                     location_code)
VALUES ((SELECT code FROM courses WHERE code = 202),
        (SELECT id FROM lecturers WHERE id = 199355),
        '2020-09-03', '2020-12-25', 5, '4-28', '9:00',
        (SELECT code FROM locations WHERE code = 1)),
       ((SELECT code FROM courses WHERE code = 203),
        (SELECT id FROM lecturers WHERE id = 199988),
        '2020-09-03', '2020-12-25', 5, '3-4', '10:55',
        (SELECT code FROM locations WHERE code = 1)),
       ((SELECT code FROM courses WHERE code = 207),
        (SELECT id FROM lecturers WHERE id = 200123),
        '2020-09-03', '2020-12-27', 5, '4-8', '9:00',
        (SELECT code FROM locations WHERE code = 1)),
       ((SELECT code FROM courses WHERE code = 208),
        (SELECT id FROM lecturers WHERE id = 200699),
        '2020-09-03', '2020-12-20', 5, '2-9', '12:25',
        (SELECT code FROM locations WHERE code = 1)),
       ((SELECT code FROM courses WHERE code = 209),
        (SELECT id FROM lecturers WHERE id = 203127),
        '2020-09-03', '2020-12-30', 5, '4-1', '10:55',
        (SELECT code FROM locations WHERE code = 1));

INSERT INTO enrollments
VALUES ((SELECT id FROM students WHERE id = 20174201),
        (SELECT id FROM lessons WHERE id = 1),
        7.6, false),
       ((SELECT id FROM students WHERE id = 20174201),
        (SELECT id FROM lessons WHERE id = 2),
        8.5, false),
       ((SELECT id FROM students WHERE id = 20174201),
        (SELECT id FROM lessons WHERE id = 3),
        9.2, false),
       ((SELECT id FROM students WHERE id = 20174201),
        (SELECT id FROM lessons WHERE id = 4),
        9.5, false),
       ((SELECT id FROM students WHERE id = 20174201),
        (SELECT id FROM lessons WHERE id = 5),
        9, false);

INSERT INTO enrollments
VALUES ((SELECT id FROM students WHERE id = 20174202),
        (SELECT id FROM lessons WHERE id = 1),
        6.5, false),
       ((SELECT id FROM students WHERE id = 20174202),
        (SELECT id FROM lessons WHERE id = 2),
        7.5, false),
       ((SELECT id FROM students WHERE id = 20174202),
        (SELECT id FROM lessons WHERE id = 3),
        8.5, false),
       ((SELECT id FROM students WHERE id = 20174202),
        (SELECT id FROM lessons WHERE id = 4),
        9.5, false),
       ((SELECT id FROM students WHERE id = 20174202),
        (SELECT id FROM lessons WHERE id = 5),
        6, false);

INSERT INTO enrollments
VALUES ((SELECT id FROM students WHERE id = 20174203),
        (SELECT id FROM lessons WHERE id = 1),
        4.5, false),
       ((SELECT id FROM students WHERE id = 20174203),
        (SELECT id FROM lessons WHERE id = 2),
        6.5, false),
       ((SELECT id FROM students WHERE id = 20174203),
        (SELECT id FROM lessons WHERE id = 3),
        8.3, false),
       ((SELECT id FROM students WHERE id = 20174203),
        (SELECT id FROM lessons WHERE id = 4),
        8.4, false),
       ((SELECT id FROM students WHERE id = 20174203),
        (SELECT id FROM lessons WHERE id = 5),
        7.2, false);

INSERT INTO enrollments
VALUES ((SELECT id FROM students WHERE id = 20174204),
        (SELECT id FROM lessons WHERE id = 1),
        7.2, false),
       ((SELECT id FROM students WHERE id = 20174204),
        (SELECT id FROM lessons WHERE id = 2),
        9.7, false),
       ((SELECT id FROM students WHERE id = 20174204),
        (SELECT id FROM lessons WHERE id = 3),
        6.4, false),
       ((SELECT id FROM students WHERE id = 20174204),
        (SELECT id FROM lessons WHERE id = 4),
        6.7, false),
       ((SELECT id FROM students WHERE id = 20174204),
        (SELECT id FROM lessons WHERE id = 5),
        8.8, false);

INSERT INTO enrollments
VALUES ((SELECT id FROM students WHERE id = 20174205),
        (SELECT id FROM lessons WHERE id = 1),
        7.1, false),
       ((SELECT id FROM students WHERE id = 20174205),
        (SELECT id FROM lessons WHERE id = 2),
        8.3, false),
       ((SELECT id FROM students WHERE id = 20174205),
        (SELECT id FROM lessons WHERE id = 3),
        9.2, false),
       ((SELECT id FROM students WHERE id = 20174205),
        (SELECT id FROM lessons WHERE id = 4),
        9.3, false);