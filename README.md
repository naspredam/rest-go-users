# Simple REST api in Go (Golang)

This project has as objective to have a rest api build on:

- Go 1.14
- MySQL (8.0.20)

To be comparable with jara standard applicaiton, it was used gorm as ORM.

## What the rest api will stand for?

The rest will have the resource:

```
/users
```

Where the endpoints exposed are:

| Method | Endpoint | Description  |
| ---    |:------- |:-----|
|GET| /users | Get all the users |
|POST| /users | Create a new user |
|GET| /users/{user_id} | Get specific user data |
|DELETE| /users/{user_id} | Delete specific user data |

## Run application

