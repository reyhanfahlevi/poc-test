## How To Test & Sampel Query

Open the graphql playground http://localhost:8080/ 

### Create User
```graphql
mutation {
  createUser(input: {
    name: "Reyhan",
    email: "reyhan.drive@gmail.com",
    password: "tokopedia",
  })
}
```
```shell script
curl 'http://localhost:8080/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' --data-binary '{"query":"# Write your query or mutation here\nmutation {\n  createUser(input: {\n    name: \"Reyhan\",\n    email: \"reyhan.drive@gmail.com\",\n    password: \"tokopedia\",\n  })\n}"}' --compressed
```

### Login (to get AuthToken)

```graphql
mutation {
  login(input: {
    email: "reyhan.drive@gmail.com",
    password: "tokopedia"
  })
}
```
```shell script
curl 'http://localhost:8080/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' --data-binary '{"query":"mutation {\n  login(input: {\n    email: \"reyhan.drive@gmail.com\",\n    password: \"tokopedia\"\n  })\n}"}' --compressed
```

### Open Shop

```graphql
mutation {
  openShop(input: {
    name: "Zone Comp",
    address: "DKI Jakarta"
  })
}
```
```shell script
curl 'http://localhost:8080/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' -H 'AuthToken: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MywibmFtZSI6IlJleWhhbiIsImVtYWlsIjoicmV5aGFuLmRyaXZlQGdtYWlsLmNvbSIsInN0YXR1cyI6MSwicmVzZXRfdG9rZW4iOiIiLCJjcmVhdGVkX3RpbWUiOiIyMDIwLTExLTE0VDAxOjI5OjU0LjcxMDI1M1oiLCJjcmVhdGVkX2J5IjowLCJ1cGRhdGVkX3RpbWUiOiIyMDIwLTExLTE0VDAxOjI5OjU0LjcxMDI1M1oiLCJ1cGRhdGVkX2J5IjowLCJleHAiOjE2MDU0Mjg4Mzl9.UB7ZxCI3klXY5hkgecq8C-ME6lRdUU8NF8s_1xOxxwg' --data-binary '{"query":"mutation {\n  openShop(input: {\n    name: \"Zone Comp\",\n    address: \"DKI Jakarta\"\n  })\n}"}' --compressed
```

### Get My Shop Info
```graphql
query {
  getMyShopInfo {
    id
    name
    address
    status
  }
}
```
```shell script
curl 'http://localhost:8080/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' -H 'AuthToken: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MywibmFtZSI6IlJleWhhbiIsImVtYWlsIjoicmV5aGFuLmRyaXZlQGdtYWlsLmNvbSIsInN0YXR1cyI6MSwicmVzZXRfdG9rZW4iOiIiLCJjcmVhdGVkX3RpbWUiOiIyMDIwLTExLTE0VDAxOjI5OjU0LjcxMDI1M1oiLCJjcmVhdGVkX2J5IjowLCJ1cGRhdGVkX3RpbWUiOiIyMDIwLTExLTE0VDAxOjI5OjU0LjcxMDI1M1oiLCJ1cGRhdGVkX2J5IjowLCJleHAiOjE2MDU0Mjg4Mzl9.UB7ZxCI3klXY5hkgecq8C-ME6lRdUU8NF8s_1xOxxwg' --data-binary '{"query":"query {\n  getMyShopInfo {\n    id\n    name\n    address\n    status\n  }\n}"}' --compressed
```

### Add Product
```graphql
mutation {
  addProduct(input: {
    name: "Test Baju 3",
    description: "Test",
    price: 100000,
    images: ["https://via.placeholder.com/300", "https://via.placeholder.com/400"]
  })
}
```
```shell script
curl 'http://localhost:8080/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' -H 'AuthToken: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MywibmFtZSI6IlJleWhhbiIsImVtYWlsIjoicmV5aGFuLmRyaXZlQGdtYWlsLmNvbSIsInN0YXR1cyI6MSwicmVzZXRfdG9rZW4iOiIiLCJjcmVhdGVkX3RpbWUiOiIyMDIwLTExLTE0VDAxOjI5OjU0LjcxMDI1M1oiLCJjcmVhdGVkX2J5IjowLCJ1cGRhdGVkX3RpbWUiOiIyMDIwLTExLTE0VDAxOjI5OjU0LjcxMDI1M1oiLCJ1cGRhdGVkX2J5IjowLCJleHAiOjE2MDU0Mjg4Mzl9.UB7ZxCI3klXY5hkgecq8C-ME6lRdUU8NF8s_1xOxxwg' --data-binary '{"query":"mutation {\n  addProduct(input: {\n    name: \"Test Baju 3\",\n    description: \"Test\",\n    price: 100000,\n    images: [\"https://via.placeholder.com/300\", \"https://via.placeholder.com/400\"]\n  })\n}"}' --compressed
```

### Get Product

```graphql
query {
  getProductInfo(productID: 12) {
    productID
    shopID
    name
    description
    price
  	images {
      ImageURL
    }
  }
}
```
```shell script
curl 'http://localhost:8080/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' -H 'AuthToken: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MywibmFtZSI6IlJleWhhbiIsImVtYWlsIjoicmV5aGFuLmRyaXZlQGdtYWlsLmNvbSIsInN0YXR1cyI6MSwicmVzZXRfdG9rZW4iOiIiLCJjcmVhdGVkX3RpbWUiOiIyMDIwLTExLTE0VDAxOjI5OjU0LjcxMDI1M1oiLCJjcmVhdGVkX2J5IjowLCJ1cGRhdGVkX3RpbWUiOiIyMDIwLTExLTE0VDAxOjI5OjU0LjcxMDI1M1oiLCJ1cGRhdGVkX2J5IjowLCJleHAiOjE2MDU0Mjg4Mzl9.UB7ZxCI3klXY5hkgecq8C-ME6lRdUU8NF8s_1xOxxwg' --data-binary '{"query":"query {\n  getProductInfo(productID: 12) {\n    productID\n    shopID\n    name\n    description\n    price\n  \timages {\n      ImageURL\n    }\n  }\n}"}' --compressed
```

### Get Product List (For Seller)

```graphql
query {
  getProductList(limit: 10, offset: 0) {
    productID
    shopID
    name
    description
    images {
      productID
      ImageURL
    }
  }
}
```
```shell script
curl 'http://localhost:8080/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' -H 'AuthToken: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MywibmFtZSI6IlJleWhhbiIsImVtYWlsIjoicmV5aGFuLmRyaXZlQGdtYWlsLmNvbSIsInN0YXR1cyI6MSwicmVzZXRfdG9rZW4iOiIiLCJjcmVhdGVkX3RpbWUiOiIyMDIwLTExLTE0VDAxOjI5OjU0LjcxMDI1M1oiLCJjcmVhdGVkX2J5IjowLCJ1cGRhdGVkX3RpbWUiOiIyMDIwLTExLTE0VDAxOjI5OjU0LjcxMDI1M1oiLCJ1cGRhdGVkX2J5IjowLCJleHAiOjE2MDU0Mjg4Mzl9.UB7ZxCI3klXY5hkgecq8C-ME6lRdUU8NF8s_1xOxxwg' --data-binary '{"query":"query {\n  getProductList(limit: 10, offset: 0) {\n    productID\n    shopID\n    name\n    description\n    images {\n      productID\n      ImageURL\n    }\n  }\n}"}' --compressed
```

### Edit Product

```graphql
mutation {
  editProduct(input: {
    productID: 12,
  	name: "Test Baju 21",
    description: "Test",
    price: 100000,
    images: ["https://via.placeholder.com/300", "https://via.placeholder.com/400"]
  })
}
```
```shell script
curl 'http://localhost:8080/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' -H 'AuthToken: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MywibmFtZSI6IlJleWhhbiIsImVtYWlsIjoicmV5aGFuLmRyaXZlQGdtYWlsLmNvbSIsInN0YXR1cyI6MSwicmVzZXRfdG9rZW4iOiIiLCJjcmVhdGVkX3RpbWUiOiIyMDIwLTExLTE0VDAxOjI5OjU0LjcxMDI1M1oiLCJjcmVhdGVkX2J5IjowLCJ1cGRhdGVkX3RpbWUiOiIyMDIwLTExLTE0VDAxOjI5OjU0LjcxMDI1M1oiLCJ1cGRhdGVkX2J5IjowLCJleHAiOjE2MDU0Mjg4Mzl9.UB7ZxCI3klXY5hkgecq8C-ME6lRdUU8NF8s_1xOxxwg' --data-binary '{"query":"mutation {\n  editProduct(input: {\n    productID: 12,\n  \tname: \"Test Baju 21\",\n    description: \"Test\",\n    price: 100000,\n    images: [\"https://via.placeholder.com/300\", \"https://via.placeholder.com/400\"]\n  })\n}"}' --compressed
```

### Delete Product
```graphql
mutation {
  deleteProduct(productID: 11) 
}
```
```shell script
curl 'http://localhost:8080/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' -H 'AuthToken: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MywibmFtZSI6IlJleWhhbiIsImVtYWlsIjoicmV5aGFuLmRyaXZlQGdtYWlsLmNvbSIsInN0YXR1cyI6MSwicmVzZXRfdG9rZW4iOiIiLCJjcmVhdGVkX3RpbWUiOiIyMDIwLTExLTE0VDAxOjI5OjU0LjcxMDI1M1oiLCJjcmVhdGVkX2J5IjowLCJ1cGRhdGVkX3RpbWUiOiIyMDIwLTExLTE0VDAxOjI5OjU0LjcxMDI1M1oiLCJ1cGRhdGVkX2J5IjowLCJleHAiOjE2MDU0Mjg4Mzl9.UB7ZxCI3klXY5hkgecq8C-ME6lRdUU8NF8s_1xOxxwg' --data-binary '{"query":"mutation {\n  deleteProduct(productID: 11) \n}"}' --compressed
```

There is more query scheme there, you can check the docs at http://localhost:8080/ 

Thank you!