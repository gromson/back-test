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

## Run locally

To start the project locally run `docker-compose up`.
Private and public API will be started as well as mongodb
and mongo-express (web interface for mongodb).

Public API can be accessed via [http://localhost:8080](http://localhost:8080)

Private API can be accessed via [http://localhost:8081](http://localhost:8081)

Mongo-Express can be accessed via [http://localhost:8082](http://localhost:8082)

# Documentation

For private and public api the following documentations provided:
- [https://app.swaggerhub.com/apis-docs/gromson/back-test-private/1.0.0](https://app.swaggerhub.com/apis-docs/gromson/back-test-private/1.0.0) - Private API 
- [https://app.swaggerhub.com/apis-docs/gromson/back-test-public/1.0.0](https://app.swaggerhub.com/apis-docs/gromson/back-test-public/1.0.0) - Public API