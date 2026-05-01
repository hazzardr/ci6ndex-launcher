# Backend Server — TODO

This document tracks infrastructure, security, and deployment tasks for the containerized Go backend server. The owner will provide specific instructions for each item.

---

## HTTPS / TLS

- [ ] **TLS Certificate**: Obtain a certificate (Let's Encrypt, Cloudflare origin cert, or custom CA).
- [ ] **HTTPS Termination**: Decide on termination strategy:
  - **Option A**: Terminate TLS in the Go app (embed cert/key, use `autocert` for Let's Encrypt).
  - **Option B**: Terminate TLS at reverse proxy (Caddy, Nginx, Traefik) and run the container on localhost/internal network.
- [ ] **HSTS Headers**: Add `Strict-Transport-Security` header if terminating HTTPS.
- [ ] **Redirect HTTP → HTTPS**: Ensure all HTTP traffic is redirected to HTTPS.

## Networking & Ingress

- [ ] **Reverse Proxy**: Configure Caddy, Nginx, or Traefik in front of the container.
- [ ] **Firewall Rules**: Restrict container port exposure (only 443/80 public; app port internal).
- [ ] **Rate Limiting**: Configure per-IP rate limits at the reverse proxy or application level.
- [ ] **DDoS Protection**: Consider Cloudflare or similar if sharing codes become public-facing.

## Secrets & Configuration

- [ ] **Environment Variables**: Document all required env vars (`PORT`, `DATABASE_URL`, `API_KEY`, etc.).
- [ ] **Secrets Management**: Use Docker secrets, Kubernetes secrets, or a vault (HashiCorp Vault, Doppler, etc.) rather than plain env vars in production.
- [ ] **Database Credentials**: Rotate credentials; do not commit them to the repo.

## Container & Orchestration

- [ ] **Distroless Image**: Use a minimal base image (`gcr.io/distroless/static`, `chainguard/static`, or Alpine) for the final stage.
- [ ] **Non-root User**: Run the container as an unprivileged user.
- [ ] **Read-only Root FS**: Mount the root filesystem as read-only where possible.
- [ ] **Health Checks**: Configure Docker/Kubernetes health checks using the `GET /health` endpoint.
- [ ] **Graceful Shutdown**: Ensure the app handles `SIGTERM` gracefully (close DB connections, finish requests).
- [ ] **Resource Limits**: Set CPU and memory limits/requests.

## Observability

- [ ] **Structured Logging**: Use JSON logging format for ingestion into Loki, CloudWatch, or similar.
- [ ] **Metrics**: Expose Prometheus metrics endpoint (`/metrics`) or integrate with OpenTelemetry.
- [ ] **Error Tracking**: Integrate Sentry, Rollbar, or similar for panic/crash reporting.
- [ ] **Uptime Monitoring**: Configure Pingdom, UptimeRobot, or similar external health checks.

## Database

- [ ] **Migrations**: Set up schema migration tool (golang-migrate, atlas, or goose).
- [ ] **Backup Strategy**: Automated backups for the profile database.
- [ ] **Connection Pooling**: Configure `sql.DB` max open/idle connections appropriately.

## Security Hardening

- [ ] **CORS Policy**: Restrict `Access-Control-Allow-Origin` to the desktop app origin(s) only.
- [ ] **Content Security Policy**: Add CSP headers if serving any frontend assets.
- [ ] **Input Validation**: Validate all incoming JSON payloads (max size, field lengths, charset).
- [ ] **Share Code Entropy**: Ensure generated share codes are sufficiently random and non-guessable.
- [ ] **API Key / Auth (optional)**: If adding user accounts or admin endpoints, implement API key or JWT auth.

---

## Owner Instructions Needed

Please provide guidance on the following:

1. **Hosting platform**: VPS (DigitalOcean, Hetzner, AWS EC2), container platform (Fly.io, Railway, Render), or Kubernetes?
2. **Domain name**: What domain (or subdomain) will the API be served from?
3. **TLS preference**: Do you want TLS termination in the Go app or at a reverse proxy?
4. **Database preference**: SQLite (simple, single-node), PostgreSQL (robust, scalable), or other?
5. **Observability stack**: Do you have an existing monitoring setup (Grafana, Datadog, etc.)?
6. **CI/CD pipeline**: GitHub Actions, GitLab CI, or manual deployments?
