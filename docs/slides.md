---
marp: true
theme: default
paginate: true
size: 16:9
title: CI/CD with GitHub Actions, Docker & Automated Deployment
description: Workshop 1 · Hacktoberfest Hack Day Buea 2026
style: |
  /* Theme from the Hacktoberfest 2026 Hack Day template (docs/Hacktoberfest 2026 Gemma Opening Ceremony.pptx). */
  @import url('https://fonts.googleapis.com/css2?family=Roboto+Mono:wght@300;400;600&family=Roboto+Condensed:wght@400;700&family=Inter:wght@400;600;700&display=swap');
  :root {
    --cream: #E4E5DA;
    --green: #203C36;
    --yellow: #F5B726;
    --red: #E53927;
    --coral: #EB7C77;
    --blue: #4787E2;
    --muted: #4B5E59;
    --panel: #EEEFE6;
  }
  section {
    background-color: var(--cream);
    background-image: url('assets/mlh-logo.png'), url('assets/hacktoberfest-bg.png');
    background-size: 6% auto, cover;
    background-position: right 1.2% bottom 2.4%, center;
    background-repeat: no-repeat;
    color: var(--green);
    font-family: 'Roboto Mono', ui-monospace, monospace;
    font-weight: 300;
    font-size: 27px;
    line-height: 1.45;
    padding: 110px 110px 56px 128px;
    justify-content: flex-start;
  }
  h1, h2 { font-family: 'Roboto Mono', monospace; font-weight: 300; color: var(--green); letter-spacing: -0.01em; }
  h1 { font-size: 1.75em; line-height: 1.2; margin: 0 0 0.45em; }
  h2 { font-size: 1.15em; margin: 0 0 0.3em; }
  h3 { font-family: 'Roboto Condensed', sans-serif; font-weight: 700; text-transform: uppercase; letter-spacing: 0.05em; color: var(--red); font-size: 0.72em; margin: 0 0 0.2em; }
  mark { background: var(--yellow); color: var(--green); padding: 0 0.12em; }
  strong { font-weight: 600; color: var(--green); }
  a { color: var(--blue); }
  ul { padding-left: 1.1em; }
  li::marker { color: var(--red); }
  blockquote { border-left: 8px solid var(--yellow); color: var(--muted); margin: 0.4em 0; padding: 0.1em 0.8em; }
  code { font-family: 'Roboto Mono', monospace; font-weight: 400; background: rgba(32, 60, 54, 0.09); color: var(--green); border-radius: 4px; padding: 0.03em 0.28em; }
  pre { background: var(--green); border-radius: 0; border-left: 8px solid var(--yellow); font-size: 0.6em; line-height: 1.4; padding: 0.8em 1em; }
  pre code { background: transparent; color: #F4F4EC; padding: 0; }
  pre .hljs-keyword, pre .hljs-built_in, pre .hljs-attr, pre .hljs-name { color: #F9C878; }
  pre .hljs-string { color: #8FE3A8; }
  pre .hljs-comment { color: #9DB3AC; font-style: italic; }
  pre .hljs-number, pre .hljs-literal, pre .hljs-variable { color: #F3A39F; }
  pre .hljs-title, pre .hljs-function, pre .hljs-tag, pre .hljs-type { color: #9CC2F5; }
  table { font-size: 0.7em; border-collapse: collapse; background: transparent; }
  table th { background: var(--green); color: var(--cream); font-weight: 600; border: 0; }
  table td { background: transparent; border: 0; border-bottom: 1px solid rgba(32, 60, 54, 0.25); }
  table tr:nth-child(2n) td { background: rgba(32, 60, 54, 0.04); }
  table td:first-child { font-weight: 600; }
  section::after { color: var(--muted); font-size: 0.5em; right: auto; left: 128px; bottom: 26px; }

  /* Section divider: big banner background, red bracketed title (template "[ Organizers ]"). */
  section.divider {
    background-image: url('assets/mlh-logo.png'), url('assets/hacktoberfest-section.png');
    justify-content: center;
    padding: 0 6% 0 43%;
  }
  section.divider h1 { font-family: 'Roboto Condensed', sans-serif; font-weight: 400; color: var(--red); font-size: 2.9em; line-height: 1.1; margin: 0; }
  section.divider p { color: var(--green); font-size: 0.9em; }
  section.divider::after { display: none; }

  /* Title slide: template "[ Hack Day Name ] [ Location ] [ Date ]". */
  section.title h1 { font-size: 2.3em; line-height: 1.25; }

  section.code { padding-top: 88px; padding-bottom: 40px; }
  section.code h1 { font-size: 1.3em; margin-bottom: 0.35em; }
  section.code pre { font-size: 0.46em; }

  .cols { display: grid; grid-template-columns: 1fr 1fr; gap: 36px; }
  .cols3 { display: grid; grid-template-columns: 1fr 1fr 1fr; gap: 26px; }
  .card { background: var(--panel); border-top: 8px solid var(--yellow); padding: 16px 22px; font-size: 0.74em; }
  .card h3 { margin-top: 0; }
  .card h2 { font-size: 1.2em; }
  .big { font-family: 'Roboto Condensed', sans-serif; font-size: 2.4em; font-weight: 700; color: var(--green); line-height: 1.05; }
  .muted { color: var(--muted); }
  .small { font-size: 0.68em; }
  .schedule { font-size: 0.9em; line-height: 1.9; }
  .schedule p { margin: 0; }
  .schedule b { font-weight: 600; background: var(--yellow); padding: 0 0.2em; margin-right: 0.4em; }
  .hey { display: grid; grid-template-columns: 30% 1fr; gap: 48px; align-items: start; margin-top: 30px; }
  .hey img { width: 100%; border: 8px solid var(--panel); }
  .hey h1 { margin-bottom: 0.6em; }
---

<!-- _class: divider -->
<!-- _paginate: false -->

# CI/CD Workshop
# Buea, Cameroon
# Hack Day 2026 

---

<!-- _class: title -->

# <mark>CI/CD</mark> with GitHub Actions, Docker & Automated Deployment



Norbel Ambanumben · Hacktoberfest Hack Day Buea 2026

---

<div class="hey">
<img src="https://github.com/aanorbel.png" alt="Norbel Ambanumben">
<div>

# Hey! I'm <mark>Norbel</mark>.

- Software engineer, 10+ years building software systems
- Active open-source contributor and community mentor
- **Today's goal:** ship an app to Render **or** Cloud Run automatically, on every push

</div>
</div>

---

# <mark>Roadmap</mark>

<div class="cols schedule">
<div>

<b>01</b> Foundations of CI/CD

<b>02</b> Containerfile vs. Dockerfile

<b>03</b> Go 1.23 + HTMX stack

<b>04</b> Multi-stage, multi-arch image

</div>
<div>

<b>05</b> GitHub Actions → ghcr.io

<b>06</b> Deploy hook to Render

<b>07</b> Same image to Cloud Run

<b>08</b> Lessons from the first run

</div>
</div>

<p class="small muted">Sprint goal: live on <code>*.onrender.com</code> and <code>*.a.run.app</code> before 15:25.</p>

---

<!-- _class: divider -->

# Foundations 

---

# Why modern workflows need <mark>both</mark>

> "If it hurts, do it more frequently — automation turns deploy day into just another push."

- **Problem:** manual deploys are rare, painful and error-prone
- **Containers:** identical runtime from laptop to cloud, so there are no surprises
- **CI/CD:** every push runs the exact same test and deploy steps, making deploys routine

---

# The two pillars of <mark>consistency</mark>

Painful manual releases only get better when the same image and the same steps run on every push.

<div class="cols">
<div class="card">

### Pillar 1 · Containerized runtime
## Container consistency

Freezes the binary, OS certificates and templates into one immutable OCI image.
The same image runs on laptops, GitHub Actions, Cloud Run and Render.

</div>
<div class="card">

### Pillar 2 · Automated verification
## CI/CD consistency

Every push runs: checkout → vet → build → test → publish → deploy.
No manual credentials, forgotten flags or human error.

</div>
</div>

---

# Why <mark>Containerfile</mark> instead of Dockerfile?

- Vendor-neutral **Open Container Initiative (OCI)** standard
- Identical syntax: `FROM`, `RUN`, `COPY`, `ENTRYPOINT`
- Works across Podman, Buildah and Docker:

```sh
podman build -f Containerfile -t go-htmx .
docker build -f Containerfile -t go-htmx .
```

---

# Containerfile <mark>vs.</mark> Dockerfile

| | |
|---|---|
| **01 Specification** | Vendor-neutral OCI format (Linux Foundation) vs. Docker Inc. origin |
| **02 Engines** | Podman, Buildah, Docker Buildx, Kaniko and GitHub Actions all read it |
| **03 Security** | Built for daemonless, rootless builds on shared CI runners |
| **04 Ecosystem** | Standard filename across CNCF and Red Hat OpenShift tooling |

Any tool that takes `-f` or `file:` accepts it: `docker/build-push-action` uses `file: ./Containerfile`.

---

<!-- _class: divider -->

# The Stack

---

# The practical stack: <mark>Go + HTMX</mark>

<div class="cols3">
<div class="card">

### Go
`net/http` + `//go:embed`

One static binary, no dependencies. The standard library only.

</div>
<div class="card">

### HTMX
`hx-post`, `hx-trigger`, `hx-swap`

Server-rendered HTML fragments. No JavaScript build step.

</div>
<div class="card">

### Container
distroless, non-root

**~10 MB** image for amd64 and arm64. Fast cold starts on Cloud Run and Render.

</div>
</div>

---

# Go + HTMX <mark>benchmarks</mark>

<div class="cols3">
<div class="card">

### Image size
<div class="big">10.3 MB</div>

Measured with `podman images`. About 98% smaller than typical 400 MB–1 GB Node.js or Python images.

</div>
<div class="card">

### Cold start
<div class="big">&lt;100 ms</div>

Target for app start on Cloud Run. Check yours in the **Container startup latency** metric.

</div>
<div class="card">

### Base RAM
<div class="big">&lt;30 MB</div>

Target, well inside the 256Mi tier. Check yours in the **Memory utilization** metric.

</div>
</div>

<p class="small muted">The page itself shows server time per request: "Rendered by Go in 69.2µs".</p>

---

<!-- _class: code -->

# Step 1a: Go + HTMX server (`main.go`)

```go
//go:embed templates/*
var templatesFS embed.FS

var version = "dev" // set at build time: -ldflags "-X main.version=<sha>"

func main() {
	port := os.Getenv("PORT") // Cloud Run: 8080, Render: 10000
	if port == "" {
		port = "8080"
	}
	tmpl := template.Must(template.ParseFS(templatesFS, "templates/*.html"))

	srv := &http.Server{Addr: ":" + port, Handler: logRequests(newServer(tmpl)),
		ReadHeaderTimeout: 5 * time.Second}

	// Cloud Run sends SIGTERM before stopping an instance: finish in-flight requests.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go srv.ListenAndServe()
	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Shutdown(shutdownCtx)
}
```

---

<!-- _class: code -->

# Step 1b: routes with Go 1.22+ method patterns

```go
func newServer(tmpl *template.Template) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		platform, instance := detectPlatform() // K_SERVICE → Cloud Run, RENDER → Render
		render(w, tmpl, "index.html", page{Results: analyze(""),
			Platform: platform, Instance: instance, Version: version})
	})
	mux.HandleFunc("POST /analyze", func(w http.ResponseWriter, r *http.Request) {
		handleHTMXFragment(w, r, tmpl) // returns only the #results fragment
	})
	return mux
}
```

`GET /analyze` returns **405** automatically. Unknown paths return **404**.

---

# Step 1c: operational <mark>probes</mark> & endpoints

| | |
|---|---|
| **Port binding** `:${PORT}` | Reads `PORT` from the platform. Defaults to 8080 locally |
| **Health probe** `GET /healthz` → `200 ok` | Used by Cloud Run startup checks and the Render health check |
| **Embedded assets** `//go:embed` | HTML templates ship inside the binary. No files to copy |
| **Request log** to stdout | `POST /analyze 200 2400B 412µs htmx=true` |
| **Graceful shutdown** on `SIGTERM` | Drains in-flight requests for up to 10 s |
| **Footer version** | The short commit SHA, so you always know which build is live |

---

<!-- _class: code -->

# Step 1d: the HTMX side (`templates/index.html`)

```html
<form hx-post="/analyze"
      hx-target="#results"
      hx-swap="innerHTML"
      hx-trigger="submit, input changed delay:400ms from:#text,
                  keydown[(metaKey||ctrlKey)&&key=='Enter'] from:#text"
      hx-sync="this:replace"
      hx-indicator="#analyze-btn">
  <textarea id="text" name="text"></textarea>
  <button id="analyze-btn" type="submit">Analyze</button>
</form>

<div id="results">{{template "results" .Results}}</div>
```

- Live analysis 400 ms after the last keystroke, or `Ctrl`/`⌘` + `Enter`
- `hx-sync="this:replace"` cancels stale requests, so old answers never win
- Styling: Tailwind v4 browser build. No Node, no npm

---

<!-- _class: divider -->

# Package

---

<!-- _class: code -->

# Step 2a: multi-stage, multi-arch `Containerfile`

```dockerfile
# Stage 1: the builder runs on the build machine's own arch; Go cross-compiles.
FROM --platform=$BUILDPLATFORM golang:1.23-alpine AS builder
WORKDIR /src
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
ARG VERSION=dev
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -ldflags="-s -w -X main.version=${VERSION}" -o /bin/app .

# Stage 2: minimal rootless distroless runtime
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /bin/app /app
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/app"]
```

Build in Alpine (~800 MB), run on distroless (**~10 MB**).

---

# Step 2b: <mark>security</mark> & speed wins

| | |
|---|---|
| **01 Module caching** | `COPY go.mod go.sum* ./` + `go mod download` reuses cached layers. The `*` matters: a stdlib-only module has no `go.sum` |
| **02 Stripped binary** | `-ldflags="-s -w"` removes debug symbols, about 30% smaller |
| **03 No OS shell** | distroless has no `/bin/sh`, `curl` or package manager |
| **04 Non-root** | `USER nonroot:nonroot` runs as unprivileged UID 65532 |
| **05 Multi-arch** | `$BUILDPLATFORM` + `$TARGETARCH`: amd64 and arm64 with no emulation |
| **06 Build version** | `ARG VERSION`: CI passes the short SHA into the binary |

---

<!-- _class: divider -->

# Pipeline

---

# Automated <mark>pipeline</mark> flow

```
push / PR ──► build ──► test ──► (main only) artifact-registry ──┬──► deploy-render
              go vet    go test      Containerfile → ghcr.io     │     deploy hook → Render (primary)
              go build               :sha + :latest, multi-arch  │
                                                                 └──► deploy-cloudrun
                                                                       OIDC → Cloud Run
```

| Job | Needs | Permissions |
|---|---|---|
| `build`, `test` | nothing | `contents: read` |
| `artifact-registry` | `test` | `+ packages: write` |
| `deploy-render` | `artifact-registry` | `contents: read` |
| `deploy-cloudrun` | `artifact-registry` | `+ id-token: write` |

Pull requests stop after `test`. Each deploy job owns its platform, so one never blocks the other.

---

<!-- _class: code -->

# Step 3a: build and test stages

```yaml
on:
  push: { branches: [main] }
  pull_request:

permissions: { contents: read }   # least privilege by default

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with: { go-version: '1.23', cache: true }
      - run: go vet ./...
      - run: CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/app .

  test:
    needs: build
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with: { go-version: '1.23', cache: true }
      - run: go test -v ./...
```

---

<!-- _class: code -->

# Step 3b: publish to ghcr.io (`artifact-registry`)

```yaml
  artifact-registry:
    needs: test
    if: github.event_name == 'push'
    permissions: { contents: read, packages: write }
    outputs: { image: '${{ steps.names.outputs.image }}' }
    steps:
      - uses: actions/checkout@v4
      - uses: docker/setup-buildx-action@v3      # needed for the type=gha cache
      - id: names
        run: |
          echo "image=ghcr.io/${GITHUB_REPOSITORY,,}" >> "$GITHUB_OUTPUT"   # lowercase!
          echo "version=${GITHUB_SHA::7}" >> "$GITHUB_OUTPUT"
      - uses: docker/login-action@v3
        with: { registry: ghcr.io, username: '${{ github.actor }}', password: '${{ secrets.GITHUB_TOKEN }}' }
      - uses: docker/build-push-action@v6
        with:
          file: ./Containerfile
          platforms: linux/amd64,linux/arm64
          push: true
          tags: ${{ steps.names.outputs.image }}:${{ github.sha }},${{ steps.names.outputs.image }}:latest
          cache-from: type=gha
          cache-to: type=gha,mode=max
          build-args: VERSION=${{ steps.names.outputs.version }}
```

---

# Step 3c: <mark>zero-secret</mark> auth & BuildKit cache

| | |
|---|---|
| **01 Native registry auth** `GITHUB_TOKEN` | `packages: write` pushes to ghcr.io with no PATs or passwords |
| **02 BuildKit layer cache** `type=gha` | `cache-from` / `cache-to: mode=max` make repeat builds much faster |
| **03 Immutable tags** `:${{ github.sha }}` | Every image maps to one commit, for rollback and traceability |
| **04 Multi-arch manifest** | amd64 for Cloud Run, Render and CI; arm64 for Apple Silicon |
| **05 Registry-only job** | No cloud steps here, so every deploy target reuses the same image |

---

<!-- _class: divider -->

# Deploy

---

<!-- _class: code -->

# Step 4: deploy to Render (`deploy-render`) — primary

```yaml
  deploy-render:
    needs: artifact-registry
    steps:
      - env:
          HOOK: ${{ secrets.RENDER_DEPLOY_HOOK_URL }}
        run: |
          if [ -z "$HOOK" ]; then
            echo "::notice::RENDER_DEPLOY_HOOK_URL is not set; skipping."; exit 0
          fi
          IMAGE="${{ needs.artifact-registry.outputs.image }}:${GITHUB_SHA}"
          curl -fsS -X POST "${HOOK}&imgURL=$(jq -rn --arg v "$IMAGE" '$v|@uri')"
```

- Render service type: **Existing Image** from `ghcr.io/owner/repo`
- `imgURL` pins each deploy to the commit's `:sha`, not `:latest`
- Render sets `PORT=10000`. The app already reads `PORT`

Full setup: `docs/render-setup.md`

---

# Bonus: Cloud Run <mark>can't pull</mark> from ghcr.io

Cloud Run pulls only from **Artifact Registry**, `gcr.io` and **Docker Hub**. Fix: a **remote repository** that proxies ghcr.io.

```sh
gcloud artifacts repositories create ghcr \
  --repository-format=docker --location=europe-west1 \
  --mode=remote-repository --remote-docker-repo=https://ghcr.io

# ghcr.io/owner/repo:sha  ➜  europe-west1-docker.pkg.dev/PROJECT/ghcr/owner/repo:sha
```

The ghcr package must be **public**: the proxy pulls without credentials.

---

<!-- _class: code -->

# Step 5a: OIDC sign-in & image path

```yaml
  deploy-cloudrun:
    needs: artifact-registry
    permissions: { contents: read, id-token: write }
    steps:
      - uses: actions/checkout@v4          # gives auth a workspace for its credentials file
      - id: auth
        uses: google-github-actions/auth@v2
        with:
          workload_identity_provider: ${{ secrets.GCP_WIF_PROVIDER }}
          service_account: ${{ secrets.GCP_SERVICE_ACCOUNT }}
      - id: names
        env:
          GHCR_IMAGE: ${{ needs.artifact-registry.outputs.image }}
          PROJECT_ID: ${{ vars.GCP_PROJECT_ID || steps.auth.outputs.project_id }}
        run: |
          RUN_IMAGE="europe-west1-docker.pkg.dev/${PROJECT_ID}/ghcr/${GHCR_IMAGE#ghcr.io/}"
          echo "project_id=${PROJECT_ID}" >> "$GITHUB_OUTPUT"
          echo "run_image=${RUN_IMAGE}" >> "$GITHUB_OUTPUT"
```

The project ID falls back to the service account's project, so a missing variable can't break the image path.

---

<!-- _class: code -->

# Step 5b: pull through the proxy, then deploy

```yaml
      - uses: google-github-actions/setup-gcloud@v2
      - name: Pull image through Artifact Registry proxy   # fails fast if ghcr is private
        run: |
          gcloud auth configure-docker europe-west1-docker.pkg.dev --quiet
          docker pull ${{ steps.names.outputs.run_image }}:${{ github.sha }}
      - id: deploy
        uses: google-github-actions/deploy-cloudrun@v2
        with:
          service: go-htmx-hackday
          region: europe-west1
          image: ${{ steps.names.outputs.run_image }}:${{ github.sha }}
          flags: >-
            --allow-unauthenticated --port=8080 --memory=256Mi
            --service-account=go-htmx-runtime@${{ steps.names.outputs.project_id }}.iam.gserviceaccount.com
      - run: echo "Live at ${{ steps.deploy.outputs.url }}"
```

`go-htmx-runtime` is the identity the container runs as. It starts with no roles.

---

# Step 5c: <mark>keyless</mark> OIDC & rollouts

| | |
|---|---|
| **Workload Identity Federation** `id-token: write` | GitHub's OIDC token is exchanged for a ~1 h Google token. No JSON keys anywhere |
| **Locked to one repo** | `--attribute-condition="assertion.repository=='owner/repo'"` |
| **Two service accounts** | `gh-deployer` deploys. `go-htmx-runtime` runs the app with no roles |
| **Zero-downtime rollout** | Traffic moves to the new revision only after it starts and passes its checks |
| **Scale to zero** | $0 when idle, scales up for judging demos (`--memory=256Mi`) |

Full setup: `docs/cloud-run-setup.md`

---

<!-- _class: divider -->

# Lessons

---

# Lessons from the <mark>real</mark> first run

| Symptom | Cause | Fix |
|---|---|---|
| Proxy pull: `403` / `not found` | New ghcr packages are **private** | Package settings → visibility **Public**, re-run |
| `invalid reference format` | `GCP_PROJECT_ID` saved as a secret, so `vars.` was empty: `pkg.dev//ghcr/...` | Read the project ID from the `auth` step's output |
| `no image found ... arm64` on a Mac | CI built only `linux/amd64` | `platforms: linux/amd64,linux/arm64` |
| `GitHub workspace is empty` warning | The deploy job had no checkout | Add `actions/checkout@v4` |
| Badge says "Running locally" on Render | Only checked Cloud Run's `K_REVISION` | Also check `RENDER` |

---

# Hackathon <mark>sprint</mark> tips (12:25 – 15:25)

- Commit `Containerfile` and `.github/workflows/deploy.yml` in **minute 15**
- Make the ghcr package **public** right after the first push
- Inject `GEMINI_API_KEY` with Render environment variables or Cloud Run `--set-secrets`. **Never in Git!**
- Auto-deploy every merge before the **15:25** judges deadline
- Check the footer version: if it shows your latest SHA, your latest code is live

---

# <mark>Live</mark> now

### Same image, same commit, two platforms

**Render** — primary
https://go-htmx-hackday-26.onrender.com

**Google Cloud Run** — same image
https://go-htmx-hackday-762003939747.europe-west1.run.app

---

# <mark>Q&A</mark> & resources


- [GitHub Actions: publishing Docker images](https://docs.github.com/en/actions/use-cases-and-examples/publishing-packages/publishing-docker-images)
- [google-github-actions/deploy-cloudrun](https://github.com/google-github-actions/deploy-cloudrun)
- [Artifact Registry remote repositories](https://cloud.google.com/artifact-registry/docs/repositories/remote-overview)
- [Render deploy hooks](https://render.com/docs/deploy-hooks)
- [HTMX docs](https://htmx.org/docs/) · [OCI image spec](https://github.com/opencontainers/image-spec)

Engr. Norbel Ambanumben · Workshop 1

---

<!-- _class: divider -->
<!-- _paginate: false -->

# Happy Hacktoberfest!
