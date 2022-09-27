# Auth Proxy 2
This is proxy that requires user to enter a password to proxy website. It's little bit fancier than basic auth. It is written in [Go](https://go.dev/). This is my second try, first one was written in Deno (link [here](https://github.com/czM1K3/auth-proxy/)) but it had weird issue with proxying Express.

## Usage
- Main usage is with [Docker](https://www.docker.com/).
- Command to run is:
	```bash
	docker run -d -p 4000:4000 -e PASSWORD=yourpassword -e SERVICE_ADDRESS=http://127.0.0.1:8080 -e LOGIN_TIME=30 ghcr.io/czm1k3/auth-proxy-2
	```
- You should change environment variables and maybe port.
	- You can change left side of port to match port you like the most.
	- **PASSWORD** is here in plain text. Quotes and backslahes may be problematic.
	- **SERVICE_ADDRESS** is address of your service. In docker-compose you can use name of service instead of ip address.
	- **LOGIN_TIME** is time in minutes for token to expire.

## Development
### Requirements
- [Go](https://go.dev/)
### Running dev
```bash
go run .
```
### Building
```bash
go build .
```
