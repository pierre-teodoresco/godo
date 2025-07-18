# Godo

A web-based todo application built with Go and Vue.js.

---

## Shoutout

The project setup was streamlined using [go-blueprint](https://github.com/Melkeydev/go-blueprint), an excellent CLI tool for spinning up Go projects with popular frameworks.

---

## Features

- **Add** new tasks with a modern web interface
- **List** and filter tasks (all, pending, completed)
- **Mark** tasks as completed or undo completion
- **Delete** tasks with confirmation
- **Real-time updates** with a responsive Vue.js frontend
- **RESTful API** built with Go's standard net/http
- **Type-safe database queries** with SQLC
- **Database migrations** for easy schema management

---

## Tech Stack

### Backend

- **Go** with standard `net/http` library
- **PostgreSQL** database
- **SQLC** for type-safe SQL code generation
- **Repository pattern** for clean architecture
- **golang-migrate** for database migrations
- **godotenv** for environment variable management

### Frontend

- **Vue.js 3** with Composition API
- **TypeScript** for type safety
- **Vite** for fast development and building
- **Tailwind CSS** for modern styling

### DevOps

- **Make** for build automation

---

## Installation

### Prerequisites

- Go 1.21+
- Node.js 18+
- PostgreSQL 13+
- Docker and Docker Compose (optional)

### 1. Clone the Repository

```sh
git clone https://github.com/pierre-teodoresco/godo.git
cd godo
```

### 2. Setup

#### Backend Setup

```sh
# Set up environment variables
cp .env.example .env
# Edit .env with your database configuration

# Install Go dependencies
go mod download

# Run database migrations
make migrate-up

# Start the API server
make run
```

#### Frontend Setup

```sh
# Navigate to frontend directory
cd frontend

# Install dependencies
npm install
# or
pnpm install

# Start development server
npm run dev
# or
pnpm dev
```

---

## Environment Configuration

Create a `.env` file in the root directory:

```env
# Database
PORT=8080
APP_ENV=local
BLUEPRINT_DB_HOST=localhost
BLUEPRINT_DB_PORT=5432
BLUEPRINT_DB_DATABASE=your_db_name
BLUEPRINT_DB_USERNAME=your_db_username
BLUEPRINT_DB_PASSWORD=your_db_password
BLUEPRINT_DB_SCHEMA=public
```

---

## Database Setup

### Using Make Commands

```sh
# Create a new migration
make migrate-create name=your_migration_name

# Run migrations
make migrate-up

# Rollback migrations
make migrate-down

# Generate SQLC code
make sqlc-generate
```

---

## API Endpoints

| Method | Endpoint     | Description       |
| ------ | ------------ | ----------------- |
| GET    | `/api/tasks` | Get all tasks     |
| POST   | `/api/task`  | Create a new task |
| PUT    | `/api/task`  | Update a task     |
| DELETE | `/api/task`  | Delete a task     |

### Example API Usage

```sh
# Get all tasks
curl http://localhost:8080/api/tasks

# Create a new task
curl -X POST http://localhost:8080/api/task \
  -H "Content-Type: application/json" \
  -d '{"title": "Buy groceries"}'

# Mark task as completed
curl -X PUT http://localhost:8080/api/task \
  -H "Content-Type: application/json" \
  -d '{"id": <uuid>, "title": "Buy groceries", "completed": true}'
```

---

## Development

### Available Make Commands

```sh
make build           # Build the application
make run             # Run the application
make migrate-up      # Run database migrations
make migrate-down    # Rollback migrations
make sqlc-generate   # Generate SQLC code
```

### Frontend Development

```sh
cd frontend

# Start development server
npm run dev

# Build for production
npm run build

# Run linting
npm run lint
```

---

## Project Structure

```
godo/
├── cmd/api/                # Application entry point
├── internal/
│   ├── database/           # Database connection
│   ├── repository/         # Data access layer
│   └── server/             # HTTP handlers and routing
├── frontend/
│   ├── src/
│   │   ├── components/     # Vue components
│   │   ├── api/            # API client
│   │   ├── types/          # TypeScript types
│   │   ├── utils/          # Utility functions
│   │   └── assets/         # Static assets
│   └── public/             # Favicon
├── migrations/             # Database migrations
└── Makefile                # Build automation
```

---

## Credits

**Backend Framework:** Built with Go's standard library and enhanced with:

- [SQLC](https://sqlc.dev/) for type-safe SQL code generation
- [golang-migrate](https://github.com/golang-migrate/migrate) for database migrations
- [godotenv](https://github.com/joho/godotenv) for environment variable management

**Frontend Framework:** Built with:

- [Vue.js 3](https://vuejs.org/) for the reactive user interface
- [Vite](https://vitejs.dev/) for fast development and building
- [Tailwind CSS](https://tailwindcss.com/) for modern styling
- [TypeScript](https://www.typescriptlang.org/) for type safety

**Project Setup:** Bootstrapped with [go-blueprint](https://github.com/Melkeydev/go-blueprint) by @Melkeydev

**Development Tools:**

- Code written with assistance from VaultAI (GPT-4.1) for Tailwind CSS implementation
- README co-written by [Cursor](https://www.cursor.com/) (Claude-4 Sonnet)

---

## Author

- **Pierre Teodoresco**

---

_Happy productivity with Godo!_ 🚀
