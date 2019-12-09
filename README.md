# go-products
Aggregate product information from multiple datasources.

## Dependencies
- GoLang version 1.13
- Locally running docker
- Administrative access on your system

## Running Locally
- Clone the repo to your local go/src folder
- Start up a locally running mongoDB in Docker by running `sudo docker run -d -p 27017:27017 -v ~/data:/data/db mongo`
- Within your local go-products repository run `go build` to pull down all dependencies and build the project
- Once the build is successful run ` ./go-products`
- Service will start on port 8080

## Endpoints
| Name         | Request Type | Endpoint                                   | Request Body                                                  |
|--------------|--------------|--------------------------------------------|---------------------------------------------------------------|
| Health       | GET          | http://localhost:8080/products/health      |                                                               |
| Get Product  | GET          | http://localhost:8080/products/v1/13860423 |                                                               |
| Update Price | PUT          | http://localhost:8080/products/v1/         | {"id": "13860423","current_price": {"value": 5.00,"currency_code": "USD"} } |


