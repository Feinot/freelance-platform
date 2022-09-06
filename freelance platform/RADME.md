# rest-api

# user-service

# REST API   

GET /users -- list of users --200, 400, 404
GET /users -- user by id -- 200, 400, 404
POST /users/:id -- create a new user -- 204, 4xx, Header Location: url
PUT /users/:id -- fully apdate user -- 204/200, 404, 400, 500
PATCH /users/:id -- partially updated user -- 204/200, 404, 400, 500
DELETe /user/:id -- delete user by id -- 204, 4044 ,400