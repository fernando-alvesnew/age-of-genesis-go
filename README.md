# MMO3D New Dream API (Go)

Projeto em Go para avaliação técnica, com recorte dos módulos:

- Login (email ou login).
- Pagamento via PagSeguro (cartao de credito).

## Arquitetura

- `internal/domain`: entidades e contratos.
- `internal/application`: casos de uso.
- `internal/infrastructure`: MySQL, JWT e cliente PagSeguro.
- `internal/interfaces/http`: handlers e rotas HTTP.

Foi aplicada separacao por camadas para facilitar manutencao, testes e evolucao (SOLID/DDD).

## Requisitos

- Go 1.22+
- MySQL 8+

## Configuracao

1. Copie o `.env.example` para `.env` (ja existe um modelo no projeto).
2. Ajuste as variaveis:
   - `MYSQL_DSN`
   - `JWT_SECRET`
   - `PAGSEGURO_TOKEN`

## Banco de dados

Execute o SQL em `migrations/001_init.sql`.

## Rodar localmente

```bash
go mod tidy
go run ./cmd/api
```

Servidor por padrao em `http://localhost:8080`.

## Endpoints

- `GET /health`
- `POST /api/login`
- `POST /api/payments/credit-card`

### Exemplo de login

```json
{
  "login": "usuario_ou_email",
  "password": "123456"
}
```

### Exemplo de pagamento cartao

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
