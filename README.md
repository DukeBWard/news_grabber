RSS News Feed

Running this program will scan the internet for RSS feeds and put them all in one place.

Using Golang, Goose, sqlc, Go Routines, Postgres

# Helpful notes:

* `sqlc generate` is really cool, will generate code in internal package based off of .sql files in sql directory.
* `goose up` and `goose down` migrations
* `goose postgres "postgres://postgres:password@localhost:5432/news_grabber" up`
