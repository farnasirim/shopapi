# Shopapi

Simple shop api for Shopify internship assignment.

## Docs
[here](docs/README.md)

## How to use
To quickly test things out go to [https://farnasirim.ir/debug/](https://farnasirim.ir/debug/)
The docker image is running on kubernetes on [http://35.229.101.218/debug](http://35.229.101.218/debug) and another server is proxying the `https` requests to the kubernetes cluster.

This project is written in `go`. Make sure `$GOPATH` is set and this project is in the correct path, which is `$GOPATH/src/github.com/farnasirim/shopapi`. 

To build the project:
```
make dependencies
make generate 
make shopapi
```

In case you don't have mongodb listening anywhere:
```
docker run --name shopapi-mongo -d -p 127.0.0.1:27017:27017 mongo
```
And to Run it:

```
./shopapi serve --address=localhost:8080 --initdb --mongodb-uri mongodb://localhost:27017 --dbname shopapidb
```

Or use the lightweight docker image:
```bash
docker run -p 8080:8080 --name my-shopapi colonelmo/shopapi serve --address=0.0.0.0:8080 --initdb --mongodb-uri mongodb://$MONGO_HOST:27017 --dbname shopapidb
```
where `MONGO_HOST` is the IP of the machine where mongodb is running. Don't forget to expose mongodb to local ip (instead of `127.0.0.1`) if you're using this method alongside a dockerized mongodb.

The api will be at `/api`

There is very little error handling done in the db layer, so by issuing queries that point to an object that is not there, the server will crash. This is unlikely to happen in a normal usage but a temporary solution is here if this bothers you.

```bash
while true; do
./shopapi serve --address=localhost:8080 --initdb --mongodb-uri mongodb://localhost:27017 --dbname shopapidb
done
```


To query to api using the graphiql ui (and to be able to auto complete using ctrl+space) go to [localhost:8080/debug](localhost:8080/debug).

Here is a couple of samples:


A query:
```graphql
query {
  shopByName(shopName:"apple") {
    products{
      name
      id
      price{
        display
      }
    }
    orders{
      id
      lines{
        product {
          id
        }
        quantity
        price{
          display
        }
      }
      price{
        display
      }
    }
  }
}
```
result:
```json
{
  "data": {
    "shopByName": {
      "products": [
        {
          "name": "iphone X",
          "id": "5ba85b61ed2f5fe9455a91b2",
          "price": {
            "display": "$999.99"
          }
        },
        {
          "name": "ipad",
          "id": "5ba85b61ed2f5fe9455a91b3",
          "price": {
            "display": "$665.50"
          }
        }
      ],
      "orders": [
        {
          "id": "5ba85b61ed2f5fe9455a91b7",
          "lines": [
            {
              "product": {
                "id": "5ba85b61ed2f5fe9455a91b2"
              },
              "quantity": 2,
              "price": {
                "display": "$1999.98"
              }
            },
            {
              "product": {
                "id": "5ba85b61ed2f5fe9455a91b3"
              },
              "quantity": 2,
              "price": {
                "display": "$1331.0"
              }
            }
          ],
          "price": {
            "display": "$3330.98"
          }
        }
      ]
    }
  }
}
```

And a mutation:
```graphql
mutation{
  addProductToOrder(productID: "5ba85b61ed2f5fe9455a91b2",
    orderID: "5ba85b61ed2f5fe9455a91b7",
    howMany: 1) {
    id
    product{
      name
      price{
        display
      }
    }
    quantity
    price{
      display
    }
  }
}

```
result:
```json
{
  "data": {
    "addProductToOrder": {
      "id": "5ba860e5772cc6b3ba8fab7d",
      "product": {
        "name": "iphone X",
        "price": {
          "display": "$999.99"
        }
      },
      "quantity": 1,
      "price": {
        "display": "$999.99"
      }
    }
  }
}
```


To run the tests:
```
make test-dependencies
make test
```

## License
MIT
