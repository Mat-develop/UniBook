**UniBook**

UniBook is a compact social-platform prototype (REST API + SPA) implemented as a Go monorepo backend and a React + TypeScript frontend. This repo contains the backend services, routing and middleware, authentication, migrations, and a Vite-powered frontend.

**Quick Pitch**: Built and maintained as a trainee project — implemented core backend services, JWT auth, DB integration, and a polished TypeScript React frontend. Great examples of full-stack engineering fundamentals to present in interviews.

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

**Notes for Hiring Managers / Recruiters**
- The codebase shows experience with: API design, token-based auth, DB schema/migrations, separation of concerns (handlers/services/repositories), and a modern TypeScript React UI.
- Suggested talking points for interviews: JWT auth flow, how middleware protects routes, the role of the repository layer for DB isolation, and how front-to-back integration is achieved.

**How You Can Use This in Job Applications**
- Resume bullet examples:
  - Implemented a RESTful Go backend with JWT authentication and MySQL persistence for a social feed and communities feature set.
  - Built a TypeScript React SPA using Vite with routing, stateful components and API integration.
- Elevator pitch: "Full-stack engineer who implemented secure API endpoints, DB migrations, and a responsive React frontend for a social-feed prototype. Comfortable owning features end-to-end."

**Next Steps & Suggestions**
- Add automated tests (unit + integration), CI pipeline and sample seeded data for easy demos.
- Improve error handling with structured API errors and request validation.
- Add Docker Compose for local reproducible development (db + backend + frontend).

**License & Contact**
- This repo currently has no license file — add one if you want to share publicly.
- If you'd like, I can create a job-tailored one-page summary or a short README variant that highlights specific skills and contributions for applications.

---
_If you'd like, I can also generate a one-page summary tailored for interviews or create a short cover-letter bullet list you can reuse._
