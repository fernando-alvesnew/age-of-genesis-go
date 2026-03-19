# MMO3D Age of Genesis (Go)

This repository is **only a technical showcase** of the real project [Age of Genesis](https://ageofgenesis.online/), created exclusively for recruiters to evaluate my code organization, architectural decisions, and implementation quality.

## Objective

The API implements a backend slice focused on:

- user authentication (login via username or email);
- credit card payments through integration with a payment gateway;
- persistence of essential data for authentication and payment flows.

## Stack

- Go
- Gin (HTTP)
- MySQL
- JWT
- PagSeguro integration
- Docker / Docker Compose

## Architecture and Patterns

The project follows a layered architecture, focusing on tactical DDD and SOLID principles:

- `cmd/api`: application entry point;
- `internal/domain`: domain entities and contracts;
- `internal/application`: use cases and application rules;
- `internal/infrastructure`: technical implementations (DB, JWT, external gateway);
- `internal/interfaces/http`: HTTP handlers and route configuration.

### Design Decisions

- **DIP (Dependency Inversion)**: use cases depend on interfaces (`Gateway`, `UserRepository`, `TokenService`, `TransactionRepository`).
- **SRP (Single Responsibility)**: payment models were split into separate files to ease evolution and maintenance.
- **Strategy**: gateway status mapping was isolated in `StatusMapper`, avoiding embedded logic in the main flow.
- **Testability**: application services are tested using stubs, without direct coupling to infrastructure.

## Current Features

### 1) Login

Flow:
1. Find user by username/email.
2. Validate password using bcrypt.
3. Block banned users.
4. Update last IP.
5. Generate JWT.

### 2) Credit Card Payment

Flow:
1. Receive payment data.
2. Build payment request for the gateway.
3. Send to PagSeguro.
4. Handle response and status.
5. Persist transaction by `reference_id` (upsert).

## Prerequisites

- Go 1.22+ (or compatible version)
- MySQL 8+

## Configuration

1. Copy `.env.example` to `.env`.
2. Set the variables:
   - `APP_ENV` (optional, default: `local`)
   - `APP_PORT` (optional, default: `8080`)
   - `MYSQL_DSN` (**required**)
   - `JWT_SECRET` (**required**)
   - `PAGSEGURO_BASE_URL` (optional, default: sandbox)
   - `PAGSEGURO_TOKEN` (**required**)

## Database

Run the initial migration:

```bash
# adjust according to your environment
# file: migrations/001_init.sql
```

## How to Run

```bash
go mod tidy
go run ./cmd/api
```

Default server: `http://localhost:8080`

## Running with Docker Compose

This project includes full containerization for local execution with API + MySQL.

### Included files

- `Dockerfile`
- `docker-compose.yml`
- `.dockerignore`

### Start services

```bash
docker compose up -d --build
```

### View logs

```bash
docker compose logs -f api
docker compose logs -f mysql
```

### Stop services

```bash
docker compose down
```

### Reset database volume (clean state)

```bash
docker compose down -v
```

Notes:

- API runs at `http://localhost:8080`.
- MySQL runs at `localhost:3306`.
- Migration script in `migrations/001_init.sql` is mounted into `/docker-entrypoint-initdb.d` and executes automatically on first database initialization.
- The API container currently uses environment variables defined directly in `docker-compose.yml` for a reproducible recruiter setup.

## Endpoints

- `GET /health`
- `POST /api/login`
- `POST /api/payments/credit-card`

### Example: `POST /api/login`

```json
{
  "login": "usuario_ou_email",
  "password": "123456"
}
```

### Example: `POST /api/payments/credit-card`

```json
{
  "user_id": 1,
  "store_carts_id": 12,
  "credit_card_holder": "Nome Sobrenome",
  "cpf_for_card": "12345678910",
  "encrypted_card": "CARD_ENCRYPTED_DATA",
  "amount": 1500,
  "description": "Pedido #12",
  "customer_email": "player@email.com",
  "notification_url": "https://seu-dominio.com/api/payment-notification",
  "items": [
    {
      "reference_id": "item_1",
      "name": "Item A",
      "quantity": 1,
      "unit_amount": 1500
    }
  ]
}
```

## Testes

```bash
go test ./...
```

## Limitations of This Scope

- Does not represent the full complexity of the production project.
- Does not include all game/platform modules.
- Designed to demonstrate good engineering practices within a focused scope.

## Note for Recruiters

If needed, I can elaborate in an interview on:

- architectural trade-offs;
- domain modeling decisions;
- testing strategy and incremental evolution.
