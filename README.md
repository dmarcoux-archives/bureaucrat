```
@dmarcoux/bureaucrat

Corporate Directory for Bureaucr.at (A fake company...)
```

# Introduction

*Bureaucr.at* is a typical hierarchical organization. Claire, its CEO has a
hierarchy of employees reporting to her. An employee can have a list of other
employees reporting to him/her. An employee with at least one reporting is
called a Manager.

The corporate directory for *Bureaucr.at* provides an interface to find the
closest common Manager (i.e. farthest from the CEO) between two employees. All
employees eventually report up to the CEO.

Requirements for this corporate directory:

- It must be using an in-memory structure
- The Manager's node has links to its employees and not the other way around

# Usage

There are 2 approaches to use this application:

*Docker (RECOMMENDED)*

Install Docker and Docker-Compose. This application has been tested on Linux
with Docker version `1.12.3` and Docker-Compose version `1.8.1`.

*Native*

Install Go with your package manager or directly at https://golang.org/dl/. This
application has been tested on Linux with Go version `1.7.3`.

## Application

*With Docker:*

Run the application with `docker-compose up development`

*Without Docker:*

Run the application with `go run main.go`

## Tests

*With Docker:*

Run the tests with `docker-compose up test`

*Without Docker:*

Run the tests with `go test -v ./...`
