# Without docker

sudo apt update
sudo apt-get install postgresql
sudo service postgresql start

# With docker

docker compose build
docker compose up -d

docker ps
docker exec -it <HASH> bash

# POST FORWARDING (expose to external)

docker run --name postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres

# Connect to the database

psql -h localhost -p 5432 -U postgres -d halalfind

# Our app should be similar to this:

https://github.com/qingwave/weave/commit/15e2ce08ce85776372e1be102e9a14633de00fe3

# Use nginx to serve the app in deployment

https://stackoverflow.com/questions/67456810/change-default-port-for-vue-js-app-on-docker
