# Job Board Service

REST API has following features

- Return all boards
- Return specific board(s)
- Add new job in certain category
- Move job to new category
- Delete Job From category
- Edit Job Details
- Create,edit,delete Tasks

## Data schema

Json in the follwoing format:

```json
[
  {
    "id": "5ded7445b671314cb6a04fe5",
    "title": "Title",
    "Categories": [
      {
        "id": 1,
        "title": "Wishlist",
        "Jobs": [
          {
            "id": 0,
            "JobDetails": {
              "company": "Trello",
              "title": "Software Engineer",
              "location": "Dublin",
              "category": "SWE",
              "post_date": "2000-01-01T00:00:00Z",
              "description": "A job description blah blah blah",
              "experience": "Junior",
              "url": "www.career.com/swe",
              "date_added": "2000-01-01T00:00:00Z",
              "salary": 100000,
              "Tasks": [
                {
                  "id": "",
                  "name": "Accept Offer",
                  "due_date": "2000-01-01T00:00:00Z"
                }
              ]
            }
          }
        ]
      },
      {
        "id": 2,
        "title": "Applied",
        "Jobs": []
      },
      {
        "id": 3,
        "title": "Interview",
        "Jobs": []
      },
      {
        "id": 4,
        "title": "Offer",
        "Jobs": [
          {
            "id": 0,
            "JobDetails": {
              "company": "Trello",
              "title": "Software Engineer",
              "location": "Dublin",
              "category": "SWE",
              "post_date": "2000-01-01T00:00:00Z",
              "description": "A job description blah blah blah",
              "experience": "Junior",
              "url": "www.career.com/swe",
              "date_added": "2000-01-01T00:00:00Z",
              "salary": 100000,
              "Tasks": [
                {
                  "id": "",
                  "name": "Accept Offer",
                  "due_date": "2000-01-01T00:00:00Z"
                }
              ]
            }
          }
        ]
      },
      {
        "id": 5,
        "title": "Rejected",
        "Jobs": []
      }
    ]
  }
]
```

## Usage

### Get All Job Board Data

'GET /boards'

**Responses**

- '200 OK' on success

```json
[
  {
    "id": "5ded7445b671314cb6a04fe5",
    "title": "Title",
    "Categories": [
      {
        "id": 1,
        "title": "Wishlist",
        "Jobs": [
          {
            "id": 0,
            "JobDetails": {
              "company": "Trello",
              "title": "Software Engineer",
              "location": "Dublin",
              "category": "SWE",
              "post_date": "2000-01-01T00:00:00Z",
              "description": "A job description blah blah blah",
              "experience": "Junior",
              "url": "www.career.com/swe",
              "date_added": "2000-01-01T00:00:00Z",
              "salary": 100000,
              "Tasks": [
                {
                  "id": "",
                  "name": "Accept Offer",
                  "due_date": "2000-01-01T00:00:00Z"
                }
              ]
            }
          }
        ]
      },
      {
        "id": 2,
        "title": "Applied",
        "Jobs": []
      },
      {
        "id": 3,
        "title": "Interview",
        "Jobs": []
      },
      {
        "id": 4,
        "title": "Offer",
        "Jobs": [
          {
            "id": 0,
            "JobDetails": {
              "company": "Trello",
              "title": "Software Engineer",
              "location": "Dublin",
              "category": "SWE",
              "post_date": "2000-01-01T00:00:00Z",
              "description": "A job description blah blah blah",
              "experience": "Junior",
              "url": "www.career.com/swe",
              "date_added": "2000-01-01T00:00:00Z",
              "salary": 100000,
              "Tasks": [
                {
                  "id": "",
                  "name": "Accept Offer",
                  "due_date": "2000-01-01T00:00:00Z"
                }
              ]
            }
          }
        ]
      },
      {
        "id": 5,
        "title": "Rejected",
        "Jobs": []
      }
    ]
  }
]
```

### Get A Job Board By Its ID

'GET /boards/{id}'

**Responses**

- '200 OK' on success

```json
[
  {
    "id": "5ded7445b671314cb6a04fe5",
    "title": "Title",
    "Categories": [
      {
        "id": 1,
        "title": "Wishlist",
        "Jobs": [
          {
            "id": 0,
            "JobDetails": {
              "company": "Trello",
              "title": "Software Engineer",
              "location": "Dublin",
              "category": "SWE",
              "post_date": "2000-01-01T00:00:00Z",
              "description": "A job description blah blah blah",
              "experience": "Junior",
              "url": "www.career.com/swe",
              "date_added": "2000-01-01T00:00:00Z",
              "salary": 100000,
              "Tasks": [
                {
                  "id": "",
                  "name": "Accept Offer",
                  "due_date": "2000-01-01T00:00:00Z"
                }
              ]
            }
          }
        ]
      },
      {
        "id": 2,
        "title": "Applied",
        "Jobs": []
      },
      {
        "id": 3,
        "title": "Interview",
        "Jobs": []
      },
      {
        "id": 4,
        "title": "Offer",
        "Jobs": [
          {
            "id": 0,
            "JobDetails": {
              "company": "Trello",
              "title": "Software Engineer",
              "location": "Dublin",
              "category": "SWE",
              "post_date": "2000-01-01T00:00:00Z",
              "description": "A job description blah blah blah",
              "experience": "Junior",
              "url": "www.career.com/swe",
              "date_added": "2000-01-01T00:00:00Z",
              "salary": 100000,
              "Tasks": [
                {
                  "id": "",
                  "name": "Accept Offer",
                  "due_date": "2000-01-01T00:00:00Z"
                }
              ]
            }
          }
        ]
      },
      {
        "id": 5,
        "title": "Rejected",
        "Jobs": []
      }
    ]
  }
]
```

### Create a Board

'POST /boards

**Responses**

- '200 OK' on success

```json
{
  "id": "000000000000000000000000",
  "title": "Title",
  "Categories": [
    {
      "id": 1,
      "title": "Wishlist",
      "Jobs": [
        {
          "id": 0,
          "JobDetails": {
            "company": "Trello",
            "title": "Software Engineer",
            "location": "Dublin",
            "category": "SWE",
            "post_date": "2000-01-01T00:00:00Z",
            "description": "A job description blah blah blah",
            "experience": "Junior",
            "url": "www.career.com/swe",
            "date_added": "2000-01-01T00:00:00Z",
            "salary": 100000,
            "Tasks": [
              {
                "id": "",
                "name": "Accept Offer",
                "due_date": "2000-01-01T00:00:00Z"
              }
            ]
          }
        }
      ]
    },
    {
      "id": 2,
      "title": "Applied",
      "Jobs": []
    },
    {
      "id": 3,
      "title": "Interview",
      "Jobs": []
    },
    {
      "id": 4,
      "title": "Offer",
      "Jobs": []
    },
    {
      "id": 5,
      "title": "Rejected",
      "Jobs": []
    }
  ]
}
```

### Delete a Board

'DELETE /board/{id}

**Responses**

- '200 OK' on success

"board deleted: 5ded742db671314cb6a04fe4"

### Update a Board

'UPDATE /board/{id}

**Responses**

- '200 OK' on success

"board updated: 5ded742db671314cb6a04fe4"

## Dev

Add this to /etc/hosts

127.0.0.1 boards.local
