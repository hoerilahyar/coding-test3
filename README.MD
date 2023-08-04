## **Required**

1. golang
2. MySQL

## **Documentation**
[Postman](https://www.postman.com/warped-meteor-617085/workspace/transvision/collection/25344978-dbf91800-5eed-41c4-951e-a9df8d55c078?action=share&creator=25344978)

## **Installation**

1. git clone `https://github.com/hoerilahyar/coding-test3.git`


## **Run the Application**

1. run `go run main.py`

## **Route**
| Http Method | Resource URI  | Description | Role |
|--|--|--|--|
| POST | /api/login | Login for user | Any |
| POST | /api/register| Create a new user | Any |
| GET | /api/user | Show all users | Any |
| GET| /api/user/:id | Show a user | Admin |
| POST | /api/user | Create a new user | Admin |
| PUT| /api/user/:id | Update a user | Admin |
| DELETE| /api/user/:id | Delete a user | Admin |
| GET | /api/product | Show all products | Any |
| GET| /api/product/:id | Show a product | Any |
| POST | /api/product| Create a new product | Admin |
| PUT| /api/product/:id | Update a product | Admin |
| DELETE| /api/product/:id | Delete a product | Admin |
| POST | /api/transaction/add-to-cart | Add the products to transaction | Any |
| POST | /api/transaction/checkout/:id | Checkout a transaction | Any |
| GET | /api/transaction | Show all transactions | Any |
| GET | /api/transaction/:id | Show a transaction | Any |