# golang-simple-blog-api
This is simple blog api with golang and gin web framework. The API have endpoints for user authentication (using JWT), creating, reading,
updating, and deleting blog posts. I have use Gin for handling HTTP requests and responses. Also, I have use GORM as the ORM to interact with the PostgreSQL database. 
Along with that I have added pagination to the API so that only a certain number of blog posts are returned in each request.


## Endpoints
| Type  | URL | Use case |
| ------------- | ------------- | ------------- |
| POST  | api/v1/signup | Signup as a new user     |
| POST  | api/v1/login  | Login        |
| GET  | api/v1/logout  | Logout        |
| GET  | api/v1/users  | list all the Users |
| GET  | api/v1/users/:id  | find the user by ID         |
| PUT  | api/v1/users/:id  | Update the user         |
| DELETE  api/v1| /users/:id  | Delete the user         |
| GET  | api/v1/posts/specific/:user  | list all the posts by Author         |
| GET  | api/v1/posts | list all the post  with defult limit (10) of pagination       |
| GET  | api/v1/posts?page=1&limit=3 | list all the post with requsted page and limit of  pagination      |
| GET  | api/v1/posts/:id | Get a post by ID         |
| POST  | api/v1/posts  | Add a new post         |
| PUT  | api/v1/posts/:id  | Update the post         |
| DELETE  api/v1| /posts/:id  | Delete the post        |


## Sample API Request and Response for User and Post
### User
#### RequestBody
```
{
  "Name": "MS1",
  "Email": "msuvariya1@gmail.com",
  "Password": "admin"
}
```

#### ResponseBody
```
{
    "id": 4,
    "name": "MS1",
    "email": "msuvariya2@gmail.com"
}
```

### Post 
#### RequestBody
```
{
    "Title":"The majestic mountain",
    "Description":"The majestic mountain peak rose high above the lush green valley, casting a magnificent shadow across the landscape. Sunlight filtered through the tall pine trees, creating a mesmerizing dance of light and shadows on the forest floor. ",
    "UserID":1
}
```

#### ResponseBody
```
{
    "CreatedAt": "2023-05-26T14:10:31.31943+05:30",
    "UpdatedAt": "2023-05-26T14:10:31.31943+05:30",
    "DeletedAt": null,
    "ID": 31,
    "title": "The majestic mountain",
    "Description": "The majestic mountain peak rose high above the lush green valley, casting a magnificent shadow across the landscape. Sunlight filtered through the tall pine trees, creating a mesmerizing dance of light and shadows on the forest floor. ",
    "UserID": 4,
    "User": {
        "id": 4,
        "name": "MS1",
        "email": "msuvariya2@gmail.com"
    }
}
```
