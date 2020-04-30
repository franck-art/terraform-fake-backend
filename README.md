# Fake Backend

## Project description

This project is a fake backend used for educational purpose.
The goal of this app is to serve a static website and connect to a MySQL database as a simple backend app would do.

_Please this projet is for private education purpose, it means not production ready and only used at SUPINFO International University_

## Configuration

This app will run on a container. To configure it you can use the following environments variables:

- STATIC_PATH: The static website path to serve (sould only contain static assets: html, css and js files, default: /etc/backend/static)
- DATABASE_HOST: The mysql host (default: "db")
- DATABASE_PORT: The mysql port (default: 3306)
- DATABASE_USER: The mysql username (default: "user")
- DATABASE_PASSWORD: The mysql password (default: "pass")
- DATABASE_NAME: The mysql database (default: "supinfo")

## Tips

1. This server is logging as many useful debug information it can for you. Don't forget to check the container logs.
2. This project exposes a `/health` route. It returns `200 - OK` if the databse connection is working, otherwise it returns `500 - KO`.
3. This project is working well with the `mysql:5.7` docker image.
4. This container exposes the port 3000

## Example usage

```sh
docker run --name fake-backend \
  --link test-mysql:db \
  -v ~/Downloads/battleboat-gh-pages:/etc/backend/static \
  -p 80:3000 \
  -d alexandrevilain/fake-backend
```
