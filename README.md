# Subscription Microservice for Sportiq project

This microservice is responsible for consuming messages from queue and processing subscription operations. Microservice is splitted into two part: HTTP-server and message consumer.

![Microservice Architecture](assets/diagram-dark.png#gh-dark-mode-only)
![Microservice Architecture](assets/diagram.png#gh-light-mode-only)

## Related Sportiq services

- [API Gateway](https://github.com/Str1kez/SportiqAPIGateway)
- [User Service](https://github.com/Str1kez/SportiqUserService)
- [Event Service](https://github.com/Str1kez/SportiqEventService)
- [Frontend App](https://github.com/Str1kez/SportiqReactApp)

## Documentation

OpenAPI - https://str1kez.github.io/SportiqSubscriptionService

## Message handler

Message handler is based on Rabbit MQ, has `N` parallel workers on goroutines. Variable can be set in `.env`. \
History is stored in PostgreSQL, subscription data is stored in ReJSON module with RediSearch support. This allows to improve the speed of working on frequently changing data.

### Preparing to start

You can execute only one command to startup the microservice:

```commandline
docker compose up -d --remove-orphans --build
```

Or do it manually:

1. Create `.env` file and fill it:
   ```commandline
   make env
   ```
2. Make [migrations](#migrations)
3. Build Docker-image:
   ```commandline
   make build-image-handler
   ```
4. Check [Start on HTTP-server](#start)

## HTTP-server

The server allows the user to interact with the microservice. It runs on gin framework with parallel workers.

### Start

You can execute only one command to startup the microservice:

```commandline
docker compose up -d --remove-orphans --build
```

Or do it manually:

1. Create `.env` file and fill it (if didn't do it on message handler step):
   ```commandline
   make env
   ```
2. Make [migrations](#migrations) (if didn't do it on message handler step)
3. Build Docker-image:
   ```commandline
   make build-image-server
   ```
4. Start the microservice:
   ```commandline
   make up
   ```

## Migrations

Need to install [migrate](https://github.com/golang-migrate/migrate)

```commandline
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

After that you must run migrations with fullfilled `.env.dev` file

## TODO

- [ ] Create a reserve queue for unprocessed messages
- [ ] Improve code architecture (more separated layers)
