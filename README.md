**UniBook**

UniBook is a compact social-platform prototype (REST API + SPA) implemented as a Go monorepo backend and a React + TypeScript frontend. This repo contains the backend services, routing and middleware, authentication, migrations, and a Vite-powered frontend.

**Quick Pitch**: Built and maintained — implemented core backend services, JWT auth, DB integration, and a polished TypeScript React frontend.

**Highlights**
- **Language & Frameworks:** Go (net/http style services), React + TypeScript, Vite
- **Data store:** MySQL (migration SQL included)
- **Auth:** JWT-based authentication and token management
- **Architecture:** Clean separation of handlers, services, repositories and models
- **Dev ergonomics:** Vite dev server, ESLint + TypeScript checks, modular Go packages

**Repository Layout (key files)**
- **Backend entry:** [monorepo/wb_api/main.go](monorepo/wb_api/main.go#L1)
- **Routing & routes:** [monorepo/wb_router/routes/routes.go](monorepo/wb_router/routes/routes.go#L1)
- **Auth utilities:** [monorepo/util/authentication](monorepo/util/authentication)
- **Migrations:** [migrations/database.sql](migrations/database.sql)
- **Frontend:** [wb-front/src](wb-front/src)

**Features**
- User registration, login, JWT token issuance and validation
- Community and post models with RESTful CRUD endpoints
- Middleware for CORS, authentication and structured responses
- Frontend SPA with pages for feed, communities, login and register

**Getting Started**
Prerequisites: `Go` (1.20+ recommended), `Node` (16+), `npm` or `pnpm`, and a running MySQL instance.

1. Database
   - Create a MySQL database (name and credentials are set via `.env`).
   - Import the schema from [migrations/database.sql](migrations/database.sql).

2. Environment
   - Copy the repo root `.env` and adjust values (`DB_NAME`, `DB_USER`, `DB_PASSWORD`, `API_PORT`, `SECRET_KEY`).

3. Backend (development)
   - Change to the backend module: `cd monorepo`
   - Fetch deps: `go mod tidy`
   - Run: `go run ./wb_api`

4. Frontend (development)
   - Change to the UI: `cd wb-front`
   - Install: `npm install`
   - Start: `npm run dev`

**Testing & Linting**
- Frontend lint: `cd wb-front && npm run lint`
- Add unit tests and integration tests as needed (currently not included).
