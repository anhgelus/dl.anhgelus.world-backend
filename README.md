# dl.anhgelus.world - backend

Backend of the website [dl.anhgelus.world](https://dl.anhgelus.world/).

Frontend is available [here](https://github.com/anhgelus/websites).

This application is a REST API.

## Usage

`/{path:.*}` -> get the list of files stored in `/data/{path}`

Response type:
```json
{
  "status": 200,
  "message": "Valid",
  "data": {}
}
```

- `status` is the status code
- `message` is the message
- `data` is the data returned by the API

## Technologies

- go 1.20
- gorilla/mux
