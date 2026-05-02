# PullVid Monorepo

This repository is split into two deployable apps:

- `frontend/`: Next.js UI intended for Vercel
- `backend/`: Go API intended for a VM or container host

## Local development

Start the backend:

```powershell
cd backend
go run ./cmd/api-server
```

Start the frontend in a second terminal:

```powershell
cd frontend
bun run dev
```

In local development, the frontend proxies `/video-info` and `/download` to the Go server via `BACKEND_URL`.

## Deployment shape

- `pullvid.com` -> Vercel project from `frontend/`
- `api.pullvid.com` -> Go server from `backend/`
