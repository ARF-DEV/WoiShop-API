# WOISHOP API

## Clone Repo
```bash
git clone https://github.com/ARF-DEV/bluesky-API.git
```
## Run via Docker Compose (Recommended)
1. Make sure you have installed [docker and docker compose](https://docs.docker.com/compose/install). I recommended you to install docker desktop because it installed docker CLI and docker compose.
2. Run this command:
```bash
docker compose up --build
```
3. To stop the program you can press `ctrl+c`
4. After that run this command:
```bash
docker compose down --remove-orphans --volumes
```

## Run manually
1. Make sure you have golang installed in your device. If not you can follow instruction [here](https://go.dev/doc/install)
2. You can install postgres in your device or you can install docker to run the postgres container **(recommended)**
3. We're gonna use docker for this one
4. Run this command to run postgreSQL:
```bash
docker run --name pgsql-dev -rm -h localhost -e POSTGRES_PASSWORD=test -dp 5432:5432 postgres
```
5. use `exit` to exit the psql and the bash.
6. download modules using this command:
```bash
go mod download
```
7. run the program using this command:
```bash
go run main.go
```
8. To stop the program you can press `ctrl+c`
9. To stop the database you can run this command:
```bash
docker rm -f pgsql-dev
```

## Test
1. Make sure you run the database first (with docker compose or manually).
2. Run this command:
```bash
go test -v ./...
```

## How to open API Documentation
1. Install NPM
2. Open `documentation` directory on terminal
3. Run This Command to serve the API documentation :
```bash
npx serve
```
4. Open the link to access the API documnetation

## ERD Design
![Latest ERD Design](https://gitlab.com/arf-dev-azura-intern/study-case-1/-/raw/main/ERD%20Design.png)
