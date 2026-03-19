# MMO3D New Dream API (Go)

Este repositório é **apenas uma exibição técnica** do projeto real [Age of Genesis](https://ageofgenesis.online/), criado exclusivamente para recrutadores avaliarem minha organização de código, decisões de arquitetura e qualidade de implementação.

## Objetivo

A API implementa um recorte de back-end focado em:

- autenticação de usuário (login por login ou e-mail);
- cobrança com cartão de crédito via integração com gateway de pagamento;
- persistência de dados essenciais para fluxo de autenticação e pagamento.

## Stack

- Go
- Gin (HTTP)
- MySQL
- JWT
- Integração com PagSeguro

## Arquitetura e padrões

O projeto segue separação por camadas, com foco em princípios de DDD tático e SOLID:

- `cmd/api`: ponto de entrada da aplicação;
- `internal/domain`: entidades e contratos do domínio;
- `internal/application`: casos de uso e regras de aplicação;
- `internal/infrastructure`: implementações técnicas (DB, JWT, gateway externo);
- `internal/interfaces/http`: handlers e configuração de rotas HTTP.

### Decisões de design

- **DIP (Dependency Inversion)**: casos de uso dependem de interfaces (`Gateway`, `UserRepository`, `TokenService`, `TransactionRepository`).
- **SRP (Single Responsibility)**: modelos de pagamento foram separados em arquivos distintos para facilitar evolução e manutenção.
- **Strategy**: mapeamento de status de gateway foi isolado em `StatusMapper`, evitando regra embutida diretamente no fluxo principal.
- **Testabilidade**: serviços de aplicação testados com stubs, sem acoplamento direto à infraestrutura.

## Funcionalidades atuais

### 1) Login

Fluxo:
1. Busca usuário por login/e-mail.
2. Valida senha com bcrypt.
3. Bloqueia usuário banido.
4. Atualiza último IP.
5. Gera JWT.

### 2) Cobrança cartão de crédito

Fluxo:
1. Recebe dados da cobrança.
2. Monta request de cobrança para gateway.
3. Envia para PagSeguro.
4. Trata retorno e status.
5. Persiste transação por `reference_id` (upsert).

## Pré-requisitos

- Go 1.22+ (ou versão compatível)
- MySQL 8+

## Configuração

1. Copie o `.env.example` para `.env`.
2. Defina as variáveis:
   - `APP_ENV` (opcional, padrão: `local`)
   - `APP_PORT` (opcional, padrão: `8080`)
   - `MYSQL_DSN` (**obrigatória**)
   - `JWT_SECRET` (**obrigatória**)
   - `PAGSEGURO_BASE_URL` (opcional, padrão sandbox)
   - `PAGSEGURO_TOKEN` (**obrigatória**)

## Banco de dados

Execute a migration inicial:

```bash
# ajuste conforme seu ambiente
# arquivo: migrations/001_init.sql
```

## Como rodar

```bash
go mod tidy
go run ./cmd/api
```

Servidor padrão: `http://localhost:8080`

## Endpoints

- `GET /health`
- `POST /api/login`
- `POST /api/payments/credit-card`

### Exemplo: `POST /api/login`

```json
{
  "login": "usuario_ou_email",
  "password": "123456"
}
```

### Exemplo: `POST /api/payments/credit-card`

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

## Limitações deste recorte

- Não representa toda a complexidade do projeto de produção.
- Não inclui todos os módulos do jogo/plataforma.
- Foi desenhado para demonstrar boas práticas de engenharia em um escopo enxuto.

## Nota para recrutadores

Se quiserem, posso detalhar em entrevista:

- trade-offs de arquitetura;
- decisões de modelagem de domínio;
- estratégia de testes e evolução incremental.
