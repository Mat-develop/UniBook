# UniBook — Project Knowledge

## What is this

UniBook is a university-focused social platform inspired by Reddit. Users can join communities, post content, and follow each other. Built as an incremental personal/student project.

## Architecture

```
UniBook/
├── monorepo/          # Go backend (API + router + domain logic)
│   ├── wb_api/        # main.go — server entry point, DI wiring
│   ├── wb_router/     # HTTP routing (Gorilla Mux)
│   │   └── routes/    # One file per domain: user, post, community, login
│   ├── v1/handlers/   # HTTP handler layer (thin, delegates to service)
│   ├── users/         # User domain: model, service, repository
│   ├── post/          # Post domain: model, service, repository
│   ├── community/     # Community domain: model, service, repository
│   └── util/          # JWT auth, middleware, response helpers
└── wb-front/          # React + TypeScript frontend
    └── src/
        ├── pages/     # Login, Register, Index (redirect)
        ├── components/ # Feed, PostContainer, Communities, CommunityContainer, Layout, Header, SideMenu
        └── utils/     # auth.ts (JWT storage + axios setup)
```

## Tech Stack

**Backend**
- Go
- Gorilla Mux (routing)
- JWT (authentication via `util/authentication`)
- MySQL

**Frontend**
- React 18 + TypeScript
- React Router v6
- Axios (HTTP, token set as default header)
- Ant Design (UI components)
- SCSS (styling)
- React Toastify (notifications)
- Vite (build tool)

## Auth Flow

1. `POST /login` → returns JWT
2. Frontend stores token in `localStorage` (remember me) or `sessionStorage`
3. `setAuthToken()` in `utils/auth.ts` sets `axios.defaults.headers.common['Authorization']`
4. All protected routes use `m.IsAuth()` middleware that validates the JWT
5. A 401 response triggers the Axios interceptor in `App.tsx` → clears token → redirects to `/login`
6. `PrivateRoute` component guards all authenticated frontend routes

## Backend Conventions

- Handlers are **interfaces** (`UserHandler`, `PostHandler`, `CommunityHandler`) — concrete types are unexported structs wired via `New*Handler()` constructors
- All responses go through `util/response`: `response.JSON(w, status, data)` and `response.Erro(w, status, err)`
- `authentication.ExtractUserId(r)` reads the userId from the JWT inside any handler
- Routes are declared as `[]Route` slices, assembled in `routes.Config()`, auth routes wrapped with `m.IsAuth()`

## Frontend Conventions

- API base URL comes from `import.meta.env.VITE_API_URL`
- Token helpers are all in `wb-front/src/utils/auth.ts`
- Components export from `wb-front/src/components/index.tsx`
- Pages live under `wb-front/src/pages/` and are registered in `App.tsx`

## What Is and Isn't Done

### Working end-to-end
- Login / Register
- List all communities
- View posts from a community (hardcoded to community ID 1 in Feed.tsx — needs to be dynamic)

### Backend ready, frontend missing
- Profile page (`/profile`) — `GET /users?user=` + `PUT /users/{userId}`
- Change password page (`/change-password`) — `PUT /users/{userId}/update-password`
- Create post — `POST /post` (no form in UI)
- View user posts — `GET /post/{userId}` (belongs on profile page)
- Search users — `GET /users?user=` (Header search bar is not wired)
- Follow / Unfollow user — `POST /users/{userId}/follow|unfollow`
- View followers — `GET /users/{userId}/followers`

### Backend handler exists but route not registered yet
- `Following()` — needs a GET route added to `user_routes.go`
- `CreateCommunity()` — needs a POST route added to `community_routes.go`

### Both backend and frontend TODO (handler stubs)
- Update post / Delete post
- Search posts by title
- Get community by name / ID
- Delete community
- Follow a community
- Get community members
- Feed/timeline (posts from followed communities/users)
- Popular posts
- New/recent posts

### Dead UI elements to wire up
- PostContainer: Like, Comment, Share buttons
- CommunityContainer: Join button + member count
- Header search bar
