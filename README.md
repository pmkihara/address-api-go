# Implementation of a 1-route app in Go

## Description
This is an exercise to get a glimpse of Go usage by creating a simple REST API using Gin Gonic.

The goal is to have a simple GET endpoint `/address/{id}` that makes an HTTP request to `https://jsonplaceholder.typicode.com/users` to get all the users, parse the data, find the user that matches the given id and return a json with the following information:

- `id`: the id given in the HTTP request
- `address`: A concatenated string with the following information separated by spaces:
  - `address.city`
  - `address.zipcode`
  - `address.geo.lat` and `address.geo.lng` wrapped in parenthesis and comma-space separated

Example:
```json
{
  "id": 1,
  "address": "Gwenborough 92998-3874 (-37.3159, 81.1496)"
}
``````

## How to run this code
To run the server, run the following code in the console:
```bash
go run main.go
```
The endpoint will be available on `localhost:8080/address/{id}`

## Improvements to be made
- Refactor the logic into separate files and functions
