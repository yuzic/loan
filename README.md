# Decentralized Loan Service

A decentralized loan service written in Go, using the **Echo framework** and **OpenAPI** for handling requests. It provides functionality for creating and managing cryptocurrency-backed loans.

## Features

- Create loans with a specified amount and collateral.
- Repay loans and handle repayments securely.
- Implements **Clean Architecture** principles for maintainability.
- Uses `zap` for structured logging.
- Interacts with blockchain (e.g., Ethereum) via `go-ethereum`.

---

## Project Structure

The project follows the **Clean Architecture** pattern, separating concerns into different layers:

- **`internal/entity`**: Contains core domain models like `Loan`.
- **`internal/service`**: Business logic for handling loans.
- **`internal/handler`**: API handlers for processing HTTP requests.
- **`internal/api`**: OpenAPI-generated code for request/response types and server interfaces.
- **`cmd`**: Entry point of the application.

---

## Prerequisites

- Go 1.23 or later
- Docker (optional, for running external services like a database or blockchain nodes)
- `golangci-lint` for linting

---

## Getting Started

### Clone the Repository

```bash
git clone git@github.com:yuzic/loan.git
cd loan
