# Welcome to Duly Noted, the Speech-to-text note taking application's API.

A REST API that handles utilizes the GORM package to allow Object-remational mapping styled programming to communicate with the database to execute CRUD operations for users & notes database tables.

## Getting Started

To run this program locally on your machine, Fork this repository and you can start the server by running "$ go run main.go" in your terminal.

## Examples

### Creating a new user

In an API development platform like Postman write a POST request to http://localhost:3001/api/v1/users/create with this in the request body:

```json
{
  "email": (insert email here),
  "password": (insert password here)
}
```
