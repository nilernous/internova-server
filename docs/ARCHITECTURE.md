```
project-root/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── app/
│   │   ├── bootstrap/
│   │   │   └── app.go
│   │   └── container/
│   │       └── di.go
│   │
│   ├── config/
│   │   └── config.go
│   │
│   ├── domain/
│   │   ├── user/
│   │   │   ├── entity.go
│   │   │   ├── repository.go
│   │   │   └── service.go
│   │   └── order/
│   │       └── ...
│   │
│   ├── application/
│   │   ├── user/
│   │   │   ├── create_user.go
│   │   │   └── get_user.go
│   │   └── order/
│   │       └── ...
│   │
│   ├── delivery/
│   │   ├── http/
│   │   │   ├── handler/
│   │   │   │   └── user_handler.go
│   │   │   ├── middleware/
│   │   │   │   ├── auth.go
│   │   │   │   └── logging.go
│   │   │   └── router.go
│   │   │
│   │   └── grpc/
│   │       └── ...
│   │
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   ├── postgres/
│   │   │   │   └── user_repository.go
│   │   │   └── redis/
│   │   ├── external/
│   │   │   └── email/
│   │   └── logger/
│   │
│   └── shared/
│       ├── errors/
│       ├── utils/
│       └── constants/
│
├── migrations/
│   └── 001_init.sql
│
├── pkg/              (optional public packages)
│
├── tests/
│   ├── integration/
│   └── e2e/
│
├── .env
├── go.mod
└── README.md
```