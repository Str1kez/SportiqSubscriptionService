# SportiqSubscriptionService
Subscription Service for Sportiq project

## Migrations
Need to install [migrate](https://github.com/golang-migrate/migrate)

```commandline
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

After that you must run migrations with fullfilled `.env.dev` file
