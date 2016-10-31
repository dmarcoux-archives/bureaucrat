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

Using Docker and Docker-Compose simplifies the setup of the project. It's highly
suggested to use this approach.

## Application

*With Docker:*

Run the application with `docker-compose up development`

*Without Docker:*

TODO

## Tests

*With Docker:*

Run the tests with `docker-compose up test`

*Without Docker:*

TODO
