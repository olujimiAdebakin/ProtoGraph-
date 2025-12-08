Here is the comprehensive README file for your GitHub project:

```markdown
# ProtoGraph GraphQL API ⚡️

## Overview
ProtoGraph is a high-performance GraphQL API gateway built with Go, leveraging `gqlgen` to orchestrate a suite of core microservices including Account, Catalog, and Order. This architecture provides a unified API surface, simplifying data access and enabling efficient application development by abstracting the underlying microservice complexities.

## Features
-   **GraphQL API**: Provides a single, powerful endpoint for querying and mutating data across multiple services using a flexible and strongly typed schema.
-   **Microservices Architecture**: Designed with independent Account, Catalog, and Order services for improved scalability, maintainability, and fault isolation.
-   **Go (`Golang`)**: Utilizes Go for its excellent concurrency features, strong type system, and superior performance characteristics, making it ideal for backend services.
-   **`gqlgen`**: Employs `gqlgen` for robust and type-safe GraphQL server implementation, automatically generating boilerplate code from the GraphQL schema, ensuring consistency and reducing manual effort.
-   **Docker**: Containerized deployment using `docker-compose` ensures consistent environments across development and production, simplifying setup and deployment.
-   **Environment Configuration**: Integrates `kelseyhightower/envconfig` for flexible, secure, and declarative environment variable management, making configuration straightforward.
-   **Dedicated Databases**: Each microservice is designed to interact with its own dedicated database (implied by `docker-compose.yaml`), promoting data independence and service autonomy.

## Getting Started
To get the ProtoGraph GraphQL API up and running locally, follow these steps.

### Installation
1.  **Clone the Repository**:
    ```bash
    git clone https://github.com/olujimiAdebakin/ProtoGraph-.git
    cd ProtoGraph
    ```

2.  **Install Go Modules**:
    Navigate to the GraphQL service directory and fetch its dependencies.
    ```bash
    cd graphql
    go mod tidy
    cd ..
    ```

3.  **Build and Run with Docker Compose**:
    The project uses Docker Compose to manage its microservices and databases.
    ```bash
    docker-compose up --build
    ```
    This command will build the Docker images for your GraphQL gateway and all configured microservices (Account, Catalog, Order, and their respective databases), then start all services.

### Environment Variables
The GraphQL gateway requires the following environment variables to connect to its downstream microservices. These are typically managed by `docker-compose.yaml` in a production-like setup or can be set directly for local Go execution.

*   `ACCOUNT_SERVICE_URL`: URL for the Account microservice.
    *   Example: `http://localhost:8081`
*   `CATALOG_SERVICE_URL`: URL for the Catalog microservice.
    *   Example: `http://localhost:8082`
*   `ORDER_SERVICE_URL`: URL for the Order microservice.
    *   Example: `http://localhost:8083`

## API Documentation

### Base URL
The GraphQL API is accessible via a single endpoint.
`http://localhost:8080/graphql`

You can explore the API using the GraphQL Playground at `http://localhost:8080/playground`.

### Endpoints

All GraphQL interactions (queries and mutations) are performed via `POST` requests to the Base URL.

#### Query: `getAccount(id: String!): Account`
Retrieves a single account by its unique identifier.

**Request**:
```graphql
query GetAccountDetails($id: String!) {
  getAccount(id: $id) {
    id
    name
    createdAt
    updatedAt
    orders {
      id
      totalPrice
      products {
        name
        quantity
      }
    }
  }
}
```
**Variables**:
```json
{
  "id": "some-account-id-123"
}
```

**Response**:
```json
{
  "data": {
    "getAccount": {
      "id": "some-account-id-123",
      "name": "Adebakin Olujimi",
      "createdAt": "2023-10-26T10:00:00Z",
      "updatedAt": "2023-10-26T10:00:00Z",
      "orders": [
        {
          "id": "order-abc-456",
          "totalPrice": 150.75,
          "products": [
            {
              "name": "Product A",
              "quantity": 2
            },
            {
              "name": "Product B",
              "quantity": 1
            }
          ]
        }
      ]
    }
  }
}
```

**Errors**:
-   `200 OK` (with `errors` array in body): Account not found or internal service error.
    Example `errors` array:
    ```json
    {
      "errors": [
        {
          "message": "Account with ID 'invalid-id' not found",
          "extensions": {
            "code": "NOT_FOUND"
          }
        }
      ],
      "data": null
    }
    ```
-   `400 Bad Request`: Malformed GraphQL query.
-   `500 Internal Server Error`: Unexpected server-side issue.

#### Query: `listAccounts(pagination: PaginationInput): [Account!]!`
Retrieves a list of accounts, with optional pagination.

**Request**:
```graphql
query ListAllAccounts($offset: Int, $limit: Int) {
  listAccounts(pagination: { offset: $offset, limit: $limit }) {
    id
    name
  }
}
```
**Variables**:
```json
{
  "offset": 0,
  "limit": 10
}
```

**Response**:
```json
{
  "data": {
    "listAccounts": [
      {
        "id": "acc-123",
        "name": "John Doe"
      },
      {
        "id": "acc-456",
        "name": "Alice Smith"
      }
    ]
  }
}
```

**Errors**:
-   `200 OK` (with `errors` array in body): Internal service error or invalid pagination input.
-   `400 Bad Request`: Malformed GraphQL query.
-   `500 Internal Server Error`: Unexpected server-side issue.

#### Query: `getProduct(id: String!): Product`
Retrieves a single product by its unique identifier.

**Request**:
```graphql
query GetProductDetails($id: String!) {
  getProduct(id: $id) {
    id
    name
    description
    price
    createdAt
  }
}
```
**Variables**:
```json
{
  "id": "some-product-id-789"
}
```

**Response**:
```json
{
  "data": {
    "getProduct": {
      "id": "some-product-id-789",
      "name": "Example Product",
      "description": "A detailed description of the example product.",
      "price": 29.99,
      "createdAt": "2023-10-26T14:00:00Z"
    }
  }
}
```

#### Query: `listProducts(pagination: PaginationInput): [Product!]!`
Retrieves a list of products, with optional pagination.

**Request**:
```graphql
query ListAllProducts($offset: Int, $limit: Int) {
  listProducts(pagination: { offset: $offset, limit: $limit }) {
    id
    name
    price
  }
}
```
**Variables**:
```json
{
  "offset": 0,
  "limit": 5
}
```

**Response**:
```json
{
  "data": {
    "listProducts": [
      {
        "id": "prod-1",
        "name": "Laptop",
        "price": 1200.00
      },
      {
        "id": "prod-2",
        "name": "Mouse",
        "price": 25.00
      }
    ]
  }
}
```

#### Mutation: `createAccount(input: AccountInput!): Account!`
Creates a new account.

**Request**:
```graphql
mutation CreateNewAccount($input: AccountInput!) {
  createAccount(input: $input) {
    id
    name
    email
    createdAt
  }
}
```
**Variables**:
```json
{
  "input": {
    "name": "New User Account",
    "email": "new.user@example.com",
    "password": "StrongSecurePassword123"
  }
}
```

**Response**:
```json
{
  "data": {
    "createAccount": {
      "id": "generated-acc-id-789",
      "name": "New User Account",
      "email": "new.user@example.com",
      "createdAt": "2023-10-26T10:30:00Z"
    }
  }
}
```

**Errors**:
-   `200 OK` (with `errors` array in body): Invalid input (e.g., duplicate email), or internal service error.
    Example `errors` array:
    ```json
    {
      "errors": [
        {
          "message": "Email 'new.user@example.com' already exists",
          "extensions": {
            "code": "CONFLICT"
          }
        }
      ],
      "data": null
    }
    ```
-   `400 Bad Request`: Malformed GraphQL query or invalid `input` structure.
-   `500 Internal Server Error`: Unexpected server-side issue.

#### Mutation: `updateAccount(id: String!, input: AccountInput!): Account!`
Updates an existing account by its unique identifier.

**Request**:
```graphql
mutation UpdateExistingAccount($id: String!, $input: AccountInput!) {
  updateAccount(id: $id, input: $input) {
    id
    name
    email
    updatedAt
  }
}
```
**Variables**:
```json
{
  "id": "generated-acc-id-789",
  "input": {
    "name": "Adebakin Olujimi",
    "email": "updated.email@example.com"
    // Password can also be updated here
  }
}
```

**Response**:
```json
{
  "data": {
    "updateAccount": {
      "id": "generated-acc-id-789",
      "name": "Adebakin Olujimi",
      "email": "updated.email@example.com",
      "updatedAt": "2023-10-26T11:00:00Z"
    }
  }
}
```

**Errors**:
-   `200 OK` (with `errors` array in body): Account not found, invalid input, or internal service error.
-   `400 Bad Request`: Malformed GraphQL query or invalid `input` structure.
-   `500 Internal Server Error`: Unexpected server-side issue.

#### Mutation: `deleteAccount(id: String!): Boolean!`
Deletes an account by its unique identifier.

**Request**:
```graphql
mutation DeleteAnAccount($id: String!) {
  deleteAccount(id: $id)
}
```
**Variables**:
```json
{
  "id": "generated-acc-id-789"
}
```

**Response**:
```json
{
  "data": {
    "deleteAccount": true
  }
}
```

**Errors**:
-   `200 OK` (with `errors` array in body): Account not found, or internal service error.
-   `400 Bad Request`: Malformed GraphQL query.
-   `500 Internal Server Error`: Unexpected server-side issue.

#### Mutation: `createProduct(input: ProductInput!): Product!`
Creates a new product.

**Request**:
```graphql
mutation CreateNewProduct($input: ProductInput!) {
  createProduct(input: $input) {
    id
    name
    description
    price
  }
}
```
**Variables**:
```json
{
  "input": {
    "name": "New Gadget",
    "description": "A fantastic new gadget for everyday use.",
    "price": 99.99
  }
}
```

**Response**:
```json
{
  "data": {
    "createProduct": {
      "id": "generated-prod-id-101",
      "name": "New Gadget",
      "description": "A fantastic new gadget for everyday use.",
      "price": 99.99
    }
  }
}
```

#### Mutation: `updateProduct(id: String!, input: ProductInput!): Product!`
Updates an existing product by its unique identifier.

**Request**:
```graphql
mutation UpdateExistingProduct($id: String!, $input: ProductInput!) {
  updateProduct(id: $id, input: $input) {
    id
    name
    price
    updatedAt
  }
}
```
**Variables**:
```json
{
  "id": "generated-prod-id-101",
  "input": {
    "name": "Updated Gadget Name",
    "price": 109.99
  }
}
```

**Response**:
```json
{
  "data": {
    "updateProduct": {
      "id": "generated-prod-id-101",
      "name": "Updated Gadget Name",
      "price": 109.99,
      "updatedAt": "2023-10-26T15:30:00Z"
    }
  }
}
```

#### Mutation: `deleteProduct(id: String!): Boolean!`
Deletes a product by its unique identifier.

**Request**:
```graphql
mutation DeleteAProduct($id: String!) {
  deleteProduct(id: $id)
}
```
**Variables**:
```json
{
  "id": "generated-prod-id-101"
}
```

**Response**:
```json
{
  "data": {
    "deleteProduct": true
  }
}
```

#### Mutation: `createOrder(input: OrderInput!): Order!`
Creates a new order.

**Request**:
```graphql
mutation CreateNewOrder($input: OrderInput!) {
  createOrder(input: $input) {
    id
    accountId
    quantity
    totalPrice
    createdAt
  }
}
```
**Variables**:
```json
{
  "input": {
    "accountId": "some-account-id-123",
    "products": [
      {
        "id": "some-product-id-789",
        "quantity": 2
      },
      {
        "id": "another-product-id-456",
        "quantity": 1
      }
    ]
  }
}
```

**Response**:
```json
{
  "data": {
    "createOrder": {
      "id": "new-order-id-555",
      "accountId": "some-account-id-123",
      "quantity": 3,
      "totalPrice": 129.97,
      "createdAt": "2023-10-26T16:00:00Z"
    }
  }
}
```

#### Mutation: `updateOrder(id: String!, input: OrderInput!): Order!`
Updates an existing order by its unique identifier.

**Request**:
```graphql
mutation UpdateExistingOrder($id: String!, $input: OrderInput!) {
  updateOrder(id: $id, input: $input) {
    id
    quantity
    totalPrice
    updatedaAt
  }
}
```
**Variables**:
```json
{
  "id": "new-order-id-555",
  "input": {
    "accountId": "some-account-id-123",
    "products": [
      {
        "id": "some-product-id-789",
        "quantity": 3
      }
    ]
  }
}
```

**Response**:
```json
{
  "data": {
    "updateOrder": {
      "id": "new-order-id-555",
      "quantity": 3,
      "totalPrice": 89.97,
      "updatedaAt": "2023-10-26T16:45:00Z"
    }
  }
}
```

#### Mutation: `deleteOrder(id: String!): Boolean!`
Deletes an order by its unique identifier.

**Request**:
```graphql
mutation DeleteAnOrder($id: String!) {
  deleteOrder(id: $id)
}
```
**Variables**:
```json
{
  "id": "new-order-id-555"
}
```

**Response**:
```json
{
  "data": {
    "deleteOrder": true
  }
}
```

## Usage
Once the services are running via `docker-compose up`, you can interact with the GraphQL API:

1.  **Access the GraphQL Playground**: Open your web browser and navigate to `http://localhost:8080/playground`. This interactive environment allows you to write and test GraphQL queries and mutations, providing features like schema introspection and auto-completion.

2.  **Execute Queries and Mutations**: Use the Playground to send the example queries and mutations provided in the [API Documentation](#api-documentation) section. For instance, to create an account, paste the `createAccount` mutation and its variables, then execute it to see the API in action.

3.  **Integrate with Clients**: For programmatic access, you can use any standard GraphQL client library in your preferred language (e.g., Apollo Client for JavaScript, `graphql-go/client` for Go, etc.) to send `POST` requests to `http://localhost:8080/graphql` with your GraphQL payload.

## Technologies Used

| Technology                                                                                                  | Purpose                                                                                | Link                                                                        |
| :---------------------------------------------------------------------------------------------------------- | :------------------------------------------------------------------------------------- | :-------------------------------------------------------------------------- |
| ![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)                             | Primary language for backend development and microservices.                            | [https://golang.org/](https://golang.org/)                                  |
| ![GraphQL](https://img.shields.io/badge/GraphQL-E10098?style=flat&logo=graphql&logoColor=white)             | Query language for APIs and runtime for fulfilling queries.                            | [https://graphql.org/](https://graphql.org/)                                |
| ![gqlgen](https://img.shields.io/badge/gqlgen-blue?style=flat)                                             | GraphQL server library for Go, generating type-safe resolvers from schema.             | [https://gqlgen.com/](https://gqlgen.com/)                                  |
| ![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)                 | Containerization of services for consistent and isolated deployment.                   | [https://www.docker.com/](https://www.docker.com/)                          |
| ![Docker Compose](https://img.shields.io/badge/Docker--Compose-2496ED?style=flat&logo=docker&logoColor=white) | Orchestration tool for defining and running multi-container Docker applications.       | [https://docs.docker.com/compose/](https://docs.docker.com/compose/)       |
| ![Envconfig](https://img.shields.io/badge/Envconfig-informational?style=flat)                               | Manages environment variables for configuration, simplifying setup.                     | [https://github.com/kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) |
| ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=flat&logo=postgresql&logoColor=white)     | Relational database used by microservices (inferred from `_db` services in Docker Compose). | [https://www.postgresql.org/](https://www.postgresql.org/)                 |
| Microservices Architecture                                                                                  | Decoupled, independently deployable services for scalability and resilience.           | -                                                                           |

## Contributing
We welcome contributions to the ProtoGraph project! If you're interested in improving this API gateway or its associated services, please follow these guidelines:

*   🐛 **Report Bugs**: If you find a bug, open an issue detailing the problem, steps to reproduce, and expected behavior.
*   💡 **Suggest Features**: Have an idea for a new feature or improvement? Open an issue to discuss it.
*   🛠️ **Submit Pull Requests**:
    *   Fork the repository.
    *   Create a new branch for your feature or bug fix (e.g., `feat/new-feature` or `bugfix/issue-123`).
    *   Ensure your code adheres to the project's coding standards.
    *   Write clear, concise commit messages.
    *   Open a pull request with a detailed description of your changes and reference any related issues.

## License
Refer to any existing LICENSE file in the repository for details.

## Author Info
Developed by a passionate software engineer with a focus on robust and scalable systems.

*   **Name**: olujimiAdebakin
*   **Email**: omoladebu231@gmail.com
---

[![Go Version](https://img.shields.io/badge/Go-1.25.3-blue.svg)](https://golang.org/)
[![GraphQL API](https://img.shields.io/badge/API-GraphQL-e10098.svg)](https://graphql.org/)
[![Docker Compose](https://img.shields.io/badge/Docker--Compose-up-informational.svg)](https://docs.docker.com/compose/)
[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen)](https://example.com/build-status)
```