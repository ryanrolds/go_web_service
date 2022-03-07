# Go Web Template

## Setup

Requires:
  * Docker or Docker Desktop - https://docs.docker.com/get-docker/

> The project makes use of Task, which is an alternative to Make. The project will run without [Task](https://taskfile.dev/) installed locally when using Docker. To run without Docker, you will need Go 1.17 or later and Task (or run the commands in Taskfile.yml manually). Not using Task or Docker will require that the env vars in `.env` be set.

## Running with Docker

Start Postgres, run the migrator, and start the remaining services
```
docker-compose up -d postgres
docker-compose up migrator
docker-compose up -d
```

At this point the backend service should be available at `http://localhost:8080/`.