
# farmsvc-go

Service to manage Farm & Pond

## Features

- Create Farm
- Update Farm
- Get Farm by ID
- Get All Farm
- Delete Farm
- Create Pond
- Update Pond
- Get Pond by ID
- Get All Pond
- Delete Pond
- Get API Statistic




## Tech Stack


- Golang 1.17
- Redis 6.2
- MySQL 8.0
- Docker


## Run Locally Docker

Clone the project

```bash
  git clone https://github.com/alvinatthariq/farmsvc-go
```

Go to the project directory

```bash
  cd farmsvc-go
```

Install dependencies

- Docker https://docs.docker.com/desktop/install/mac-install/

Run docker compose

```bash
  docker-compose up -d
```


## Run Locally

Clone the project

```bash
  git clone https://github.com/alvinatthariq/farmsvc-go
```

Go to the project directory

```bash
  cd farmsvc-go
```

Install dependencies

- MySQL 8.0
- Redis 6.2




Start the server

```bash
  make run
```


## Running Tests

To run tests, run the following command

```bash
  make run-test
```

