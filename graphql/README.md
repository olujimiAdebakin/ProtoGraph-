# GraphQL Microservice Gateway

This project serves as a GraphQL gateway for managing accounts, products, and orders, integrating with various backend microservices. It provides a unified GraphQL API layer, simplifying data fetching and manipulation across your services.

## Features

*   **Comprehensive GraphQL API**: Exposes a rich GraphQL API for all core functionalities, including:
    *   **Account Management**: Create, update, delete, retrieve, and list user accounts.
    *   **Product Management**: Create, update, delete, retrieve, and list products.
    *   **Order Management**: Create, update, delete orders, including associating products and quantities.
*   **Microservice Integration**: Seamlessly connects to separate Account, Catalog, and Order microservices, abstracting their complexities behind a single GraphQL endpoint.
*   **Auto-generated Resolvers**: Leverages `gqlgen` for efficient GraphQL schema and resolver generation, reducing boilerplate code.
*   **GraphQL Playground**: Includes an interactive GraphQL Playground for easy API exploration, testing queries, and mutations.
*   **Configuration Management**: Uses `envconfig` for robust environment variable-based configuration, making deployment flexible.
*   **Go-based Performance**: Built with Go, offering high performance and concurrency suitable for API gateways.

## Stacks / Technologies

| Technology          | Description                                                    | Link                                            |
| :------------------ | :------------------------------------------------------------- | :---------------------------------------------- |
| **Go**              | Backend programming language                                   | [https://golang.org/](https://golang.org/)      |
| **gqlgen**          | GraphQL server library for Go, generating resolvers and types. | [https://gqlgen.com/](https://gqlgen.com/)      |
| **envconfig**       | Processes environment variables into Go structs.               | [https://github.com/kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) |
| **net/http**        | Standard Go library for HTTP server functionality.             | [https://pkg.go.dev/net/http](https://pkg.go.dev/net/http) |
| **GraphQL Schema**  | Defines the structure of the API data and operations.          | [https://graphql.org/](https://graphql.org/)    |

## Installation

To get this project up and running locally, follow these steps:

1.  **Prerequisites**:
    *   [Go](https://golang.org/doc/install) (version 1.16 or higher recommended)
    *   Access to the respective Account, Catalog, and Order microservices (or mock versions running on default ports).

2.  **Clone the repository**:
    ```bash
    git clone <repository_url>
    cd <project_directory>
    ```

3.  **Install Dependencies**:
    The project uses Go modules. Ensure you have the necessary dependencies by running:
    ```bash
    go mod tidy
    ```

4.  **Configure Environment Variables**:
    This gateway relies on environment variables to connect to the downstream microservices. Create a `.env` file or set these variables in your environment:
    ```
    ACCOUNT_SERVICE_URL="http://localhost:8081"
    CATALOG_SERVICE_URL="http://localhost:8082"
    ORDER_SERVICE_URL="http://localhost:8083"
    ```
    Adjust the URLs as per your microservice deployments.

5.  **Run the application**:
    ```bash
    go run main.go graph.go account_resolver.go mutation_resolver.go query_resolver.go generated.go models_gen.go models.go
    ```
    Alternatively, build and run the executable:
    ```bash
    go build -o graphql-gateway .
    ./graphql-gateway
    ```

## Usage

Once the server is running, you can access the GraphQL API and the GraphQL Playground.

*   **GraphQL Endpoint**: `http://localhost:8080/graphql`
*   **GraphQL Playground**: `http://localhost:8080/playground`

Open your web browser to `http://localhost:8080/playground` to interact with the API.

### Example Query

To list all accounts:

```graphql
query {
  listAccounts {
    id
    name
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

### Example Mutation

To create a new account:

```graphql
mutation {
  createAccount(input: {
    name: "John Doe",
    email: "john.doe@example.com",
    password: "securepassword123"
  }) {
    id
    name
    createdAt
  }
}
```

## Contributing

We welcome contributions! If you'd like to contribute, please follow these steps:

1.  **Fork the repository**.
2.  **Create a new branch** for your feature or bug fix: `git checkout -b feature/your-feature-name` or `git checkout -b bugfix/issue-description`.
3.  **Make your changes** and ensure they adhere to the project's coding standards.
4.  **Write and run tests** to cover your changes.
5.  **Commit your changes** with a clear and descriptive commit message.
6.  **Push your branch** to your forked repository.
7.  **Open a Pull Request** to the main repository, describing your changes and their benefits.

[![Readme was generated by Readmit](https://img.shields.io/badge/Readme%20was%20generated%20by-Readmit-brightred)](https://readmit.vercel.app)