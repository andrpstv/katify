# Report Parser

Сервис для сбора аналитики в CRM-системах с возможностью интеграции ИИ-решений для бизнеса. В настоящее время ориентирован на отделы продаж, но планируется расширение до полноценной SaaS-платформы.

## Описание проекта

Report Parser - это backend-сервис, разработанный на Go, который предоставляет API для интеграции с популярными CRM-системами (AmoCRM, GetCourse, Bitrix). Сервис позволяет собирать аналитику, управлять аккаунтами и проектами, а также предоставляет инфраструктуру для будущих ИИ-интеграций.

## Основные возможности

### Текущие возможности
- **Аутентификация пользователей**: Регистрация и авторизация с JWT-токенами
- **Интеграция с AmoCRM**: Получение аккаунтов и проектов пользователя
- **Многопользовательская архитектура**: Поддержка нескольких пользователей и аккаунтов
- **Управление проектами**: Создание и управление проектами на основе CRM-аккаунтов
- **Telegram-бот**: Интеграция для уведомлений и взаимодействия
- **REST API**: Полноценное API для всех операций

## Технологический стек

### Backend
- **Go 1.24.3** - Основной язык программирования
- **Gin** - HTTP-фреймворк для REST API
- **PostgreSQL** - Основная база данных
- **JWT** - Аутентификация и авторизация
- **Zerolog** - Логирование
- **Goose** - Миграции базы данных
- **sqlc** - Генерация Go-кода из SQL

### Внешние интеграции
- **AmoCRM API** - Интеграция с CRM-системой
- **Telegram Bot API** - Бот для уведомлений
- **GetCourse API** - Интеграция с образовательной платформой
- **Bitrix24 API** - Интеграция с CRM

### Архитектура
Проект построен на принципах чистой архитектуры (Clean Architecture):
- **Adapters**: Внешние интерфейсы (HTTP, DB, внешние API)
- **Application**: Бизнес-логика и use cases
- **Domain**: Доменные модели и бизнес-правила
- **Ports**: Интерфейсы для зависимостей

## Установка и запуск

### Предварительные требования
- Go 1.24+
- PostgreSQL 12+
- Git

### 1. Клонирование репозитория
```bash
git clone <repository-url>
cd report_parser
```

### 2. Установка зависимостей
```bash
go mod download
```

### 3. Настройка переменных окружения
Создайте файл `.env` в корне проекта:

```env
# База данных
DEV_DB_HOST=localhost
DEV_DB_PORT=5432
DEV_DB_USER=your_db_user
DEV_DB_PASSWORD=your_db_password
DEV_DB_NAME=report_parser

# Сервер
HTTP_PORT=:8080
API_VER=/api/v1
LOG_LEVEL=info
APP_ENV=dev

# AmoCRM
AMOCRM_URL=https://www.amocrm.ru
AMOCRM_BASEURL=https://www.amocrm.ru/oauth
AMOCRM_ACCOUNTS=https://www.amocrm.ru/api/v4/accounts
AMO_LOGIN=your_amo_login
AMO_PASSWORD=your_amo_password

# Telegram
TELEGRAM_API_TOKEN=your_telegram_bot_token
```

### 6. Запуск сервера
```bash
go run cmd/bot/main.go
```

Сервер будет доступен по адресу `http://localhost:8080`

## API Документация

### Аутентификация

#### Регистрация пользователя
```http
POST /auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "username": "username",
  "fullName": "Full Name",
  "password": "password123"
}
```

#### Авторизация пользователя
```http
POST /auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

### Ответ при успешной авторизации
```json
{
  "id": "uuid",
  "email": "user@example.com",
  "name": "Full Name",
  "accessToken": "jwt_access_token",
  "refreshToken": "jwt_refresh_token",
  "expiresAt": "2024-01-01T12:00:00Z",
  "cookies": {}
}
```

## Database

### Основные таблицы

#### users
Пользователи системы
```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    full_name TEXT,
    mfa_enabled BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

#### accounts
Внешние аккаунты пользователей (AmoCRM, GetCourse, etc.)
```sql
CREATE TABLE accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    provider_id INT NOT NULL REFERENCES external_providers(id),
    provider_user_id TEXT NOT NULL,
    display_name TEXT,
    email TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

#### workspaces
Рабочие пространства (аккаунты в CRM)
```sql
CREATE TABLE workspaces (
    id BIGSERIAL PRIMARY KEY,
    account_id UUID NOT NULL REFERENCES accounts(id),
    uuid UUID DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    subdomain TEXT,
    shard_type INT,
    version INT DEFAULT 0,
    is_kommo BOOLEAN DEFAULT FALSE,
    is_trial BOOLEAN DEFAULT FALSE,
    trial_ended BOOLEAN DEFAULT FALSE,
    is_payed BOOLEAN DEFAULT FALSE,
    payed_ended BOOLEAN DEFAULT FALSE,
    mfa_enabled BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

#### projects
Проекты пользователей
```sql
CREATE TABLE projects (
    id BIGSERIAL PRIMARY KEY,
    amo_workspace_id BIGINT REFERENCES workspaces(id),
    getcourse_workspace_id BIGINT REFERENCES workspaces(id),
    name TEXT NOT NULL,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW()
);
```

## Разработка

### Структура проекта
```
report_parser/
├── cmd/bot/              # Точка входа приложения
├── internal/
│   ├── adapters/         # Внешние интерфейсы
│   │   ├── api/          # Клиенты внешних API
│   │   ├── app/          # Репозитории и сервисы приложения
│   │   ├── db/           # Работа с БД
│   │   └── http/         # HTTP-обработчики
│   ├── application/      # Бизнес-логика
│   ├── domain/           # Доменные модели
│   ├── dto/              # Data Transfer Objects
│   └── ports/            # Интерфейсы
├── migrations/           # Миграции БД
├── sqlc/                 # Генерируемый код для работы с БД
├── pkg/                  # Общие пакеты
└── config/               # Конфигурация
```

### Генерация кода БД
```bash
sqlc generate
```

### Запуск тестов
```bash
go test ./...
```

### Линтинг
```bash
go vet ./...
gofmt -d .
```

## Развертывание

### Docker (планируется)
```bash
docker build -t report-parser .
docker run -p 8080:8080 report-parser
```

### Production настройки
Для production окружения установите переменные:
```env
APP_ENV=prod
PROD_DB_HOST=your_prod_db_host
PROD_DB_USER=your_prod_db_user
PROD_DB_PASSWORD=your_prod_db_password
PROD_DB_NAME=your_prod_db_name
```

## Мониторинг и логирование

- **Zerolog**: Структурированное логирование
- **Уровни логирования**: debug, info, warn, error, fatal
- **Формат**: JSON для production, human-readable для development

## Безопасность

- **JWT-токены**: Access tokens (15 мин), Refresh tokens (7 дней)
- **Пароль hashing**: bcrypt
- **CORS**: Настроен для frontend-приложения
- **MFA**: Поддержка многофакторной аутентификации (планируется)

## TODO()

### shortly
- [ ] Полная интеграция с AmoCRM API
- [ ] Веб-интерфейс (React/Vue)
- [ ] Сбор базовой аналитики
- [ ] Telegram-бот для уведомлений

### longly
- [ ] ИИ-анализ данных CRM
- [ ] Автоматизация бизнес-процессов
- [ ] SaaS-платформа с биллингом
- [ ] Интеграция с дополнительными CRM
- [ ] Мобильное приложение
- [ ] API для партнеров
