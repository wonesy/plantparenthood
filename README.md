# PlantParenthood - Backend

Keep track of the plants you own.
* Watering schedules
* Sun / Soil preferences
* Repotting schedules

Features TBD

## Development

Requires `go` and `docker` to be installed.

Some other tools you may need to install:
```
go get github.com/99designs/gqlgen
go get github.com/golang-migrate/migrate
```

1. start the database
   ```
   docker-compose up
   ```
2. Execute database migrations
   ```
   ./scripts/migrate.bash
   ```
2. build the go binary
   ```
   go build .
   ```
3. run the binary
   ```
   ./plantparenthood
   ```

## GQL schema changes
If any changes are made to the graphql schema, you need to regenerate the code

```bash
./scripts/regen.bash
```