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

This project has been set to run under docker.

To run the application the `Makefile` has been set to:

- start: to start the application in docker
- stop: stop and drop the containers
- restart: stop and start
- logs: tail logs from the containers

So, the `start` action will do:

- gradle build the project to generate the jar
- docker build the image, based on the jar generated on the previous step
- start database and the application, which a dependency from application to the database start, to wait until the database is up and ready for the application

The `stop` action will do:

- stop the application and database
- destroy the containers for the application and the database

### Out-of-scope

It was got out of scope any software to create the database structure, out from mysql setup.

This is to keep this app simple, to not get deeper on other solutions outside of the endpoint exposures.
