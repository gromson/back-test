# Back Test Challenge

## Description

The project contains two main packages, one for public part
and one for private part.

The project is deployed on heroku and can be accessed via 
the following addresses:
- [https://romson-back-public-api.herokuapp.com/](https://romson-back-public-api.herokuapp.com/) - public api
- [https://romson-back-private-api.herokuapp.com/](https://romson-back-private-api.herokuapp.com/) - private api

Access to the private API is protected with a basic authentication.
Credentials `admin:back-challenge`.

The project uses https://travis-ci.org as a CI/CD tool.

NOTE: The first call to heroku app can be slow because free nodes fall asleep
after 30 minutes idle.

## Run locally

To start the project locally run `docker-compose up`.
Private and public API will be started as well as mongodb
and mongo-express (web interface for mongodb).

Public API can be accessed via [http://localhost:8080](http://localhost:8080)

Private API can be accessed via [http://localhost:8081](http://localhost:8081)

Mongo-Express can be accessed via [http://localhost:8082](http://localhost:8082)

NOTE: docker-compose always starts private server with `--import` flag which will lead
to import of messages to the database on every start therefore it's recommended to
run `docker-compose down` after stopping containers to remove them or modify an entrypoint
for private service to `["./wait-for-it.sh", "mongo:27017", "--", "private", "--addr=:8081"]`.  

## Documentation

For private and public api the following documentations provided:
- [https://app.swaggerhub.com/apis-docs/gromson/back-test-private/1.0.0](https://app.swaggerhub.com/apis-docs/gromson/back-test-private/1.0.0) - Private API 
- [https://app.swaggerhub.com/apis-docs/gromson/back-test-public/1.0.0](https://app.swaggerhub.com/apis-docs/gromson/back-test-public/1.0.0) - Public API

## Test

Tests provided as an example for `back-api/internal/authentication` package and for 
public server handler.

Postman collection for testing endpoints: `/test/back.postman_collection.json`

## Trade-offs 

### Security

Each endpoint in private API is protected with basic authentication however, in the real life
application it would make sense to have a dedicated authentication service that provides
authorization token. Based on that token access to particular endpoints could be granted.

### CI/CD

In this challenge task quite simple CI/CD pipeline has been set up whereas, in a real world
application it should be more customized.

### Tests

Tests in this aplication have been provided as examples but in a real world application
all handlers had to be covered by tests and most of the method in other services.

### Messages import

To start private server instantly, import of the big amount of messages can be considered to
run in the goroutine.

### Database

MongoDB was chosen only for challenge task purpose, not sure I'd use it for the real
life application.