// db.createUser({
//     user: 'root',
//     pwd: 'example',
//     roles: [
//         {
//             role: 'readWrite',
//             db: 'university',
//         },
//     ],
// });

// db = new Mongo().getDB("university");
// db.createCollection('courses', {capped: false});

db.courses.insert(
    [{
        "_id": "60cc77a34accca62b6194235",
        "title": "title 1",
        "department_id": "60cc77a34accca62b6194238",
        "description": "description 1"
    },
        {
            "_id": "60cc77a34accca62b6194236",
            "title": "title 2",
            "department_id": "60cc77a34accca62b6194238",
            "description": "description 2"
        },
        {
            "_id": "60cc77a34accca62b6194237",
            "title": "title 3",
            "department_id": "60cc77a34accca62b6194238",
            "description": "description 3"
        }
    ]
);
