### About
Service **Find Your Team** (FYT) helps find teammates for your IT projects.

### Installation
```shell
git clone
# To create db check PostgreSQL credentials first on file `dev.yaml`
make createdb
# Create empty tables
make migrateup
# build
make
# Start server (port 8080 by default)
make run
```



