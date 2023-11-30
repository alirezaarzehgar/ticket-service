# Ticketing Service
This is a pilot project for Arttizens company. I tried to write clean and standard code
and use some cool tools to make my codebase more professional.

# Tools
- Docker Compose
- Swagger
- Github Actions for CI
- MySQL
- GNU Make

## Technical Structure
Before coding I design my idea on [this](https://github.com/alirezaarzehgar/ticketservice/issues/1) issue.
You can check.
## Go libraries
- Echo web framework
- GORM
- Gocron
- slog
- smtp

## Log system
This project implemented a log system using slog and gocron.
Everyday logs of application will wrote on ./log/date.log
You can silent logs with `DEBUG=false` on .env

# Run project
Firt of all you should config `.env`. Copy `.env.example` and change some fields like smtp configurations.

```
cp .env.example .env
```

Then you can deploy application using `make` commands.
```
$ make help
help                           Show this help.
prod                           Deploy project
dev                            Run database on docker and code on local machine
stop                           Stop and remove all containers
test-compile                   Compile project
test                           Run tests and return status code of unit tests
test-dev                       Run database on container and tests on local machine
swag                           Install swagger and format/init comments
```

## Swagger documentions
After running project you can access to documentions using following path:

localhost:8000/swagger/index.html

## Contributing
We welcome and appreciate all contributions, no matter how small!
