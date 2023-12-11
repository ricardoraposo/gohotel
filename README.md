# Hotel reservation backend

# Project outline
- users -> book room from an hotel 
- admins -> going to check reservation/bookings 
- Authentication and authorization -> JWT tokens
- Hotels -> CRUD API -> JSON
- Rooms -> CRUD API -> JSON
- Scripts -> database management -> seeding, migration

## Resources
### Mongodb driver 
Documentation

https://mongodb.com/docs/drivers/go/current/quick-start

Installing mongodb client
```bash
go get go.mongodb.org/mongo-driver/mongo
```

### gofiber 
Documentation
https://gofiber.io

Installing gofiber
```bash
go get github.com/gofiber/fiber/v2
```

## Docker
### Installing mongodb as a Docker container
```bash
docker run --name mongodb -d mongo:latest -p 27017:27017
