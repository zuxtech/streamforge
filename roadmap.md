Absolutely. Here’s the **StreamForge roadmap** we’ve been building toward, in implementation order.

## StreamForge Roadmap

### Phase 1 — Foundation

* [x] Go project setup
* [x] Repository structure
* [x] Apache-2.0 license
* [x] `ARCHITECTURE.md`
* [x] Configuration system
* [x] HTTP server
* [x] Request ID middleware
* [x] Health endpoints
* [x] Platform package boundaries
* [x] Adapter architecture

---

### Phase 2 — Authentication

**Goal: establish identity and tenancy.**

* [x] ORY Kratos integration
* [x] `Authenticator` interface
* [x] `Registrar` interface
* [x] Kratos client
* [x] Registration flow
* [x] Registration completion
* [x] Kratos error translation
* [ ] Login flow
* [ ] Session handling
* [ ] Logout
* [ ] Password recovery
* [ ] Email verification

Current position:

```text
Registration
     ↓
Kratos Identity
     ↓
[ NEXT ]
StreamForge User
```

---

### Phase 3 — User & Tenancy

**Goal: establish the StreamForge ownership model.**

* [ ] User model
* [ ] Create StreamForge User after registration
* [ ] Link `kratos_identity_id`
* [ ] Organization model
* [ ] Organization creation
* [ ] Membership model
* [ ] Roles
* [ ] Project model
* [ ] Project creation
* [ ] Organization/project authorization

Target:

```text
User
 │
 └── Organization
       │
       ├── Memberships
       │
       └── Projects
             │
             ├── Videos
             ├── API Keys
             ├── Webhooks
             └── Processing Jobs
```

---

### Phase 4 — API Keys

**Goal: allow customer backends to communicate with StreamForge.**

* [ ] API key model
* [ ] `sf_test_...` keys
* [ ] `sf_live_...` keys
* [ ] Key hashing
* [ ] Key creation
* [ ] Key revocation
* [ ] Key rotation
* [ ] API-key authentication middleware
* [ ] Project-scoped authorization

Then:

```text
Customer Backend
      │
      │ Authorization: Bearer sf_live_...
      ▼
StreamForge API
      │
      ▼
API Key
      │
      ▼
Project
```

---

# Phase 5 — Video Assets

**This is where StreamForge becomes a video platform.**

* [ ] Video model
* [ ] Video IDs (`vid_...`)
* [ ] Create video API
* [ ] Video metadata
* [ ] Video lifecycle
* [ ] Video authorization
* [ ] Delete video

Lifecycle:

```text
CREATED
   ↓
UPLOADING
   ↓
UPLOADED
   ↓
PROCESSING
   ↓
READY
```

Failure:

```text
PROCESSING
    ↓
  FAILED
    ↓
  RETRY
```

---

# Phase 6 — Uploads

**Goal: upload video without sending video bytes through the API.**

* [ ] Upload session
* [ ] Signed upload URL
* [ ] S3 storage adapter
* [ ] Direct upload
* [ ] Complete upload
* [ ] Upload validation
* [ ] Object metadata

Flow:

```text
Customer
   │
   │ POST /videos
   ▼
StreamForge
   │
   │ signed URL
   ▼
Customer
   │
   │ video bytes
   ▼
Object Storage
```

The API never handles the actual video bytes.

---

# Phase 7 — Processing

**Goal: turn uploaded files into playable video.**

* [ ] Job model
* [ ] Queue abstraction
* [ ] Redis/Valkey adapter
* [ ] Worker service
* [ ] Job lifecycle
* [ ] FFmpeg adapter
* [ ] Media probing
* [ ] Transcoding
* [ ] Thumbnail generation
* [ ] Error handling
* [ ] Retry system

Architecture:

```text
video.uploaded
      ↓
    Queue
      ↓
    Worker
      ↓
   FFmpeg
      ↓
Processed Assets
```

---

# Phase 8 — HLS & Playback

**Goal: make the video playable.**

* [ ] HLS packaging
* [ ] `.m3u8` manifests
* [ ] `.ts` / fMP4 segments
* [ ] Multi-bitrate encoding
* [ ] Playback metadata
* [ ] Playback authorization
* [ ] Signed playback URLs
* [ ] CDN integration

Example:

```text
1080p ──┐
720p  ──┼──> HLS
480p  ──┘
           ↓
          CDN
           ↓
        Player
```

---

# Phase 9 — Webhooks & Events

**Goal: let customer applications react to video events.**

* [ ] Event model
* [ ] Transactional outbox
* [ ] Webhook configuration
* [ ] Webhook signing
* [ ] Delivery worker
* [ ] Retries
* [ ] Idempotency
* [ ] Delivery history
* [ ] Replay
* [ ] Dead-letter handling

Events:

```text
video.created
video.uploaded
video.processing
video.ready
video.failed
video.deleted
```

---

# Phase 10 — Observability

* [ ] Structured logging
* [ ] Metrics
* [ ] OpenTelemetry
* [ ] Distributed tracing
* [ ] Queue metrics
* [ ] Processing metrics
* [ ] Storage metrics
* [ ] Webhook metrics
* [ ] Upload-to-ready metrics

---

# Phase 11 — Infrastructure

**Goal: make StreamForge deployable anywhere.**

* [ ] Docker
* [ ] Docker Compose
* [ ] Kubernetes
* [ ] Helm chart
* [ ] Terraform
* [ ] BYO storage
* [ ] BYO infrastructure
* [ ] Self-hosted deployment
* [ ] Configuration-driven providers

Provider model:

```text
StreamForge
    │
    ├── Auth
    │    └── Kratos
    │
    ├── Storage
    │    └── S3
    │
    ├── Queue
    │    └── Redis
    │
    └── Transcoder
         └── FFmpeg
```

Later these become replaceable.

---

# Phase 12 — Advanced Video

* [ ] H.265 / HEVC
* [ ] AV1
* [ ] GPU encoding
* [ ] HDR
* [ ] Subtitles
* [ ] Audio normalization
* [ ] Watermarking
* [ ] Scene detection
* [ ] Content analysis
* [ ] AI processing
* [ ] Custom codec plugins

---

# Phase 13 — Platform / Business

* [ ] Usage metering
* [ ] Storage usage
* [ ] Processing usage
* [ ] Bandwidth usage
* [ ] Billing integration
* [ ] Plans
* [ ] Limits
* [ ] Quotas
* [ ] Audit logs
* [ ] Admin dashboard

---

# The Immediate Roadmap

Don't jump ahead to FFmpeg yet.

We're currently here:

```text
Foundation
    ✓
    │
    ▼
Authentication
    ✓ Registration
    ✓ Kratos integration
    ✓ Registration errors
    │
    ▼
User & Tenancy          ← NEXT
    │
    ├── User
    ├── Organization
    ├── Membership
    └── Project
    │
    ▼
API Keys
    │
    ▼
Videos
    │
    ▼
Uploads
    │
    ▼
Processing
    │
    ▼
HLS
    │
    ▼
Playback
    │
    ▼
Webhooks
    │
    ▼
Observability
    │
    ▼
BYO / Self-hosted
```

### The next concrete implementation

I recommend we do this next:

**`Registration → Kratos Identity → StreamForge User`**

Then:

**`User → Organization → Membership → Project`**

That gives us the complete ownership/authorization foundation before we start building the video pipeline.
