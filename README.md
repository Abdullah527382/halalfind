### Setting everything up

```
# First, install WSL https://learn.microsoft.com/en-us/windows/wsl/install

# Install docker desktop https://www.docker.com/products/docker-desktop/
# You will also install compose: https://docs.docker.com/compose/install/

# Build
docker compose build
docker compose up

# You should be able to access frontend and backend at
# localhost:3000 and localhost:8080 respectively in the browser

# Access database, password=postgres
psql -h localhost -p 5432 -U postgres -d halalfind

# optional:
Install the SDK for openfoodfacts
go get github.com/openfoodfacts/openfoodfacts-go


```

### Resources:

Follow this tutorial:
https://go.dev/doc/tutorial/web-service-gin

### Database and API

https://world.openfoodfacts.org/

https://github.com/openfoodfacts

NOTE: I have also sent a request to data.gov.au querying for a dataset which contains the nutritional value of foods in AU, surely there is one that also contains ingredients.

Request a slack account if you wish to ask for help:
https://slack.openfoodfacts.org/

Download datasets here:
https://world.openfoodfacts.org/data

### Application structure:

https://www.calhoun.io/using-mvc-to-structure-go-web-applications/

### Frontend (follow this guide)

https://betterprogramming.pub/how-to-serve-a-single-page-application-using-go-4b9a38d92987
