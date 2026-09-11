# StreamForge Architecture

> Video infrastructure for developers.

StreamForge is an API-first video infrastructure platform that provides developers with the primitives required to upload, process, store, deliver, and manage video without having to build and operate the underlying video pipeline themselves.

StreamForge separates its **control plane**, **data plane**, and **delivery plane** so that infrastructure providers can change without forcing changes to the product's core business logic.

---

## 1. Architecture Goals

StreamForge is designed around the following principles:

1. **API-first**

   * Everything important should be accessible through APIs.
   * SDKs, dashboards, CLIs, and other interfaces consume the API rather than becoming separate sources of business logic.

2. **Infrastructure independence**

   * Storage, compute, queues, authentication, and CDN providers should be replaceable.
   * Core application logic must not depend directly on infrastructure vendors.

3. **Direct media transfer**

   * StreamForge should not proxy large video uploads or playback traffic through the API.
   * Uploads go directly to object storage.
   * Playback is delivered through CDN infrastructure.

4. **Control plane ownership**

   * StreamForge owns the metadata, authorization, orchestration, lifecycle, usage, and developer experience around video.

5. **Replaceable data plane**

   * Video processing infrastructure should be replaceable.
   * The same control plane should support StreamForge-managed infrastructure and customer-owned infrastructure.

6. **Asynchronous processing**

   * Video processing is asynchronous.
   * Long-running work must not block API requests.

7. **PostgreSQL as the source of truth**

   * Business state belongs in PostgreSQL.
   * Queues and caches are not authoritative sources of application state.

8. **Capabilities over vendors**

   * Application code depends on capabilities such as `Storage`, `Queue`, `Transcoder`, and `Authenticator`.
   * Adapters implement those capabilities.

9. **Start simple**

   * StreamForge begins as a modular distributed monolith with separate API and worker processes.
   * Services should only be split when independent scaling, reliability, ownership, or deployment requirements justify it.

10. **Multi-tenant by design**

    * Organizations, projects, users, and resources are isolated through explicit ownership and authorization boundaries.

---

# 2. High-Level Architecture

```text
                         INTERNET
                            │
             ┌──────────────┴──────────────┐
             │                             │
       Customer Backend              Customer Frontend
             │                             │
             │                         SDK / Player
             │                             │
             ▼                             ▼
      ┌─────────────────────────────────────────┐
      │              StreamForge API            │
      │                  Go                     │
      └───────────────────┬─────────────────────┘
                          │
             ┌────────────┼────────────┐
             │            │            │
             ▼            ▼            ▼
        PostgreSQL      Redis      Object Storage
             │            │            │
             │            ▼            │
             │       Job / Events      │
             │            │            │
             │      ┌─────┴─────┐      │
             │      │           │      │
             │      ▼           ▼      │
             │   Video       Webhook   │
             │   Worker       Worker   │
             │      │                  │
             │      ▼                  │
             │    FFmpeg               │
             │      │                  │
             └──────┼──────────────────┘
                    │
                    ▼
                   HLS
                    │
                    ▼
                   CDN
                    │
                    ▼
                End User
```

The API coordinates the system but does not carry the video bytes through the control plane.

---

# 3. Three-Plane Architecture

StreamForge is divided conceptually into three planes.

## 3.1 Control Plane

The control plane contains StreamForge's product logic.

Responsibilities include:

* Authentication
* Users
* Organizations
* Organization membership
* Roles and permissions
* Projects
* API keys
* Video metadata
* Upload sessions
* Processing jobs
* Processing state
* Playback authorization
* Webhooks
* Usage tracking
* Billing metadata
* Analytics metadata
* Configuration
* Developer experience

The control plane is the part StreamForge fundamentally owns.

---

## 3.2 Data Plane

The data plane performs operations on media.

Responsibilities include:

* Video upload
* Object storage
* Media probing
* Transcoding
* HLS packaging
* Thumbnail generation
* Caption processing
* Media transformations
* Other long-running media operations

The data plane should be replaceable.

For example:

```text
StreamForge Control Plane
          │
          ▼
       Queue
          │
          ▼
     Video Worker
          │
          ▼
       FFmpeg
          │
          ▼
   Object Storage
```

In another deployment:

```text
StreamForge Control Plane
          │
          ▼
 Customer Worker
          │
          ▼
 Customer Compute
          │
          ▼
 Customer Storage
```

The control plane should not care which implementation performs the work.

---

## 3.3 Delivery Plane

The delivery plane delivers processed video to end users.

Responsibilities include:

* HLS delivery
* CDN integration
* Signed playback URLs
* Playback authorization
* Domain restrictions
* Token validation
* Edge caching
* Playback optimization

The StreamForge API should authorize playback but should not become the media delivery path.

```text
Application
     │
     ▼
StreamForge API
     │
     │ authorize
     ▼
Signed Playback URL
     │
     ▼
CDN
     │
     ▼
HLS
     │
     ▼
Viewer
```

---

# 4. Repository Architecture

```text
streamforge/
├── apps/
│   ├── dashboard/
│   ├── docs/
│   └── website/
│
├── packages/
│   ├── javascript/
│   ├── react/
│   ├── react-native/
│   ├── player/
│   ├── types/
│   └── config/
│
├── cmd/
│   ├── api/
│   └── worker/
│
├── internal/
│   ├── user/
│   ├── organization/
│   ├── project/
│   ├── video/
│   ├── processing/
│   ├── playback/
│   ├── webhook/
│   ├── analytics/
│   ├── billing/
│   ├── usage/
│   │
│   └── platform/
│       ├── auth/
│       ├── config/
│       ├── errors/
│       ├── http/
│       ├── logging/
│       └── observability/
│
├── adapters/
│   ├── auth/
│   │   └── kratos/
│   ├── storage/
│   │   └── s3/
│   ├── transcoding/
│   │   └── ffmpeg/
│   ├── queue/
│   │   └── redis/
│   └── cdn/
│
├── db/
│   ├── migrations/
│   └── seeds/
│
├── api/
│   └── openapi.yaml
│
├── infrastructure/
│   ├── docker/
│   ├── terraform/
│   └── kubernetes/
│
├── tests/
│   ├── integration/
│   ├── e2e/
│   └── load/
│
├── ARCHITECTURE.md
├── README.md
├── Makefile
├── package.json
├── pnpm-workspace.yaml
├── turbo.json
├── go.mod
└── go.sum
```

---

# 5. Directory Responsibilities

## `cmd/`

Contains executable Go applications.

```text
cmd/api/
cmd/worker/
```

`cmd/api` starts the HTTP API.

`cmd/worker` starts background processing workers.

`cmd/` should contain application composition and startup code, not business logic.

---

## `internal/`

Contains StreamForge's core application and domain logic.

The code should be organized primarily by **business capability**, not by technical layer.

Examples:

```text
internal/video/
internal/project/
internal/organization/
internal/processing/
internal/playback/
internal/webhook/
```

A feature should ideally own its:

* domain models
* business rules
* application services
* repository contracts
* capability interfaces
* feature-specific logic

---

## `internal/platform/`

Contains shared technical capabilities.

Examples:

```text
internal/platform/auth/
internal/platform/config/
internal/platform/errors/
internal/platform/http/
internal/platform/logging/
internal/platform/observability/
```

Platform code should provide reusable infrastructure primitives without containing StreamForge-specific business rules.

---

## `adapters/`

Contains concrete implementations of infrastructure capabilities.

Examples:

```text
adapters/storage/s3/
adapters/transcoding/ffmpeg/
adapters/queue/redis/
adapters/auth/kratos/
```

Adapters know about external technologies.

The core domain should not.

---

## `db/`

Contains database-specific resources.

```text
db/migrations/
db/seeds/
```

PostgreSQL is the initial primary database.

---

## `api/`

Contains the public API contract.

The initial API contract will be defined using OpenAPI.

```text
api/openapi.yaml
```

SDKs and API documentation should ultimately derive from or remain consistent with this contract.

---

## `infrastructure/`

Contains deployment and infrastructure configuration.

Examples:

```text
infrastructure/docker/
infrastructure/terraform/
infrastructure/kubernetes/
```

Infrastructure configuration must remain separate from application business logic.

---

## `packages/`

Contains reusable JavaScript/TypeScript packages.

Potential packages include:

```text
@streamforge/javascript
@streamforge/react
@streamforge/react-native
@streamforge/player
@streamforge/types
@streamforge/config
```

These packages consume the StreamForge API.

They should not duplicate backend business logic.

---

# 6. Dependency Direction

The most important dependency rule is:

```text
cmd/
  │
  ▼
internal/
  │
  │ capability interfaces
  ▼
adapters/
  │
  ▼
external infrastructure
```

The dependency direction should generally flow inward.

For example:

```text
internal/video
      │
      │ Storage interface
      ▼
adapters/storage/s3
      │
      ▼
Amazon S3
```

The video domain must not import an S3 SDK simply because S3 is currently being used.

Similarly:

```text
internal/processing
      │
      │ Transcoder interface
      ▼
adapters/transcoding/ffmpeg
      │
      ▼
FFmpeg
```

And:

```text
internal/platform/auth
      │
      │ Authenticator interface
      ▼
adapters/auth/kratos
      │
      ▼
ORY Kratos
```

This allows the infrastructure implementation to change without rewriting business logic.

---

# 7. Capability Interfaces

Interfaces should represent capabilities required by the core application.

They should live close to the consumer that needs them.

For example:

```go
// internal/processing/transcoder.go

type Transcoder interface {
    Probe(ctx context.Context, input Input) (MediaInfo, error)
    Transcode(ctx context.Context, job Job) (Result, error)
}
```

An FFmpeg adapter implements this interface:

```text
internal/processing
       │
       │ Transcoder
       ▼
adapters/transcoding/ffmpeg
```

Do not create a giant global interface package containing every possible abstraction.

Prefer small capability interfaces.

Good:

```text
Storage
Queue
Transcoder
Authenticator
CDN
```

Avoid:

```text
UniversalInfrastructureProvider
```

The objective is **pluggability without premature abstraction**.

---

# 8. Authentication Architecture

Authentication is treated as a pluggable capability.

Conceptually:

```text
                         StreamForge API
                              │
                              ▼
                    ┌──────────────────┐
                    │   Authenticator  │
                    │     interface    │
                    └────────┬─────────┘
                             │
              ┌──────────────┼──────────────┐
              ▼              ▼              ▼
         ORY Kratos       OIDC/Auth0     Self-hosted
           adapter          adapter        adapter
```

The initial implementation may use ORY Kratos.

Kratos owns:

* Authentication
* Identity management
* Sessions
* Password flows
* Login
* Registration
* Recovery

StreamForge owns:

* User records
* Organizations
* Organization membership
* Roles
* Projects
* API keys
* Business authorization

The relationship is:

```text
Kratos Identity
      │
      │ identity ID
      ▼
StreamForge User
      │
      ├── Organizations
      │
      ├── Memberships
      │
      └── Projects
```

StreamForge must not make its business model dependent on Kratos-specific concepts.

---

# 9. API Authentication

StreamForge has two major authentication paths.

## Dashboard

```text
Browser
   │
   ▼
Kratos
   │
   ▼
Session
   │
   ▼
StreamForge API
```

## Customer Backend

```text
Customer Server
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

API keys belong to projects rather than users.

Keys should:

* be scoped to a project
* have test/live environments
* be stored hashed
* be displayed only once
* support revocation
* support rotation

Example formats:

```text
sf_test_...
sf_live_...
```

---

# 10. Multi-Tenancy

StreamForge is multi-tenant.

The primary hierarchy is:

```text
User
 │
 ▼
Organization
 │
 ├── Members
 │
 └── Projects
       │
       ├── Videos
       ├── API Keys
       ├── Processing Jobs
       ├── Webhooks
       └── Usage
```

Most resources should belong to an organization directly or through a project.

For example:

```text
Video
  └── Project
        └── Organization
```

The initial architecture uses a shared PostgreSQL database with tenant-aware records.

Separate databases per customer are not required for the initial architecture.

---

# 11. Public Resource IDs

Public API resources should use opaque IDs rather than exposing database implementation details.

Examples:

```text
org_...
proj_...
vid_...
job_...
key_...
wh_...
```

Internal database identifiers may still use UUIDs.

Public identifiers provide:

* safer API contracts
* easier resource identification
* reduced coupling to database implementation
* room for future identifier changes

---

# 12. Video Lifecycle

A video moves through a state machine.

Initial lifecycle:

```text
CREATED
   │
   ▼
UPLOADING
   │
   ▼
UPLOADED
   │
   ▼
PROCESSING
   │
   ├──────────────┐
   ▼              ▼
 READY          FAILED
                  │
                  │ retry
                  ▼
               PROCESSING
```

A future state such as `PUBLISHED` may be introduced if publishing becomes distinct from processing.

The video record represents the authoritative state of the video.

Processing jobs represent work performed against that state.

---

# 13. Commands, Jobs, and Events

These concepts are intentionally separate.

## Commands

Commands express requested actions.

Examples:

```text
CreateVideo
CompleteUpload
StartProcessing
GenerateThumbnail
DeliverWebhook
DeleteVideo
```

## Jobs

Jobs represent asynchronous work.

Examples:

```text
TranscodeVideoJob
GenerateThumbnailJob
DeliverWebhookJob
RecordAnalyticsJob
```

## Events

Events describe something that happened.

Examples:

```text
video.created
video.uploaded
video.processing
video.ready
video.failed
video.deleted
```

The mental model is:

```text
COMMAND
   │
   ▼
JOB
   │
   ▼
WORKER
   │
   ▼
STATE CHANGE
   │
   ▼
EVENT
```

---

# 14. Transactional Outbox

Important state changes should use the transactional outbox pattern.

Example:

```sql
CREATE TABLE outbox_events (
    id UUID PRIMARY KEY,
    aggregate_type TEXT NOT NULL,
    aggregate_id UUID NOT NULL,
    event_type TEXT NOT NULL,
    payload JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    published_at TIMESTAMPTZ
);
```

The state change and outbox insertion happen in the same PostgreSQL transaction.

For example:

```text
PostgreSQL Transaction
       │
       ├── update video
       │
       └── insert outbox event
```

A publisher can then deliver the event asynchronously.

This prevents the following failure:

```text
Update database
     │
     ▼
Application crashes
     │
     ▼
Event never published
```

The outbox makes the database the durable source of truth.

---

# 15. Queue Architecture

Redis/Valkey is initially used as a job/delivery mechanism.

Example queues:

```text
video.probe
video.transcode
video.thumbnail
webhook.delivery
analytics
usage
```

PostgreSQL remains authoritative.

Redis/Valkey is not the primary source of business state.

A formal event streaming platform such as Kafka, Pulsar, or NATS should only be introduced if actual scale or architecture requirements justify it.

---

# 16. Upload Architecture

Video uploads must avoid passing large media through the API.

The intended flow is:

```text
Customer
   │
   │ POST /v1/videos
   ▼
StreamForge API
   │
   │ create upload session
   ▼
Signed Upload URL
   │
   ▼
Customer
   │
   │ direct upload
   ▼
Object Storage
   │
   │ completion
   ▼
StreamForge API
   │
   ▼
Processing Job
```

The API handles authorization and metadata.

Object storage handles the actual media transfer.

This architecture improves:

* scalability
* upload performance
* API reliability
* bandwidth costs
* horizontal scaling

---

# 17. Video Processing

Once an upload is completed:

```text
video.uploaded
      │
      ▼
Processing Queue
      │
      ▼
Video Worker
      │
      ├── Probe
      │
      ├── Transcode
      │
      ├── Package HLS
      │
      └── Generate thumbnails
      │
      ▼
Object Storage
      │
      ▼
video.ready
```

The initial transcoding implementation uses FFmpeg.

Typical initial output profiles may include:

```text
1080p
720p
480p
```

The exact encoding ladder should remain configurable.

---

# 18. HLS Packaging

StreamForge uses HLS as the initial delivery format.

A processed video produces assets such as:

```text
master.m3u8
1080p/
720p/
480p/
thumbnails/
```

Example:

```text
Object Storage
└── videos/
    └── vid_123/
        ├── master.m3u8
        ├── 1080p/
        ├── 720p/
        ├── 480p/
        └── thumbnails/
```

The exact storage layout is an implementation detail and should not become part of the public API contract unless necessary.

---

# 19. Playback Architecture

Playback is authorized by StreamForge and delivered through the CDN.

```text
Customer Application
        │
        ▼
StreamForge API
        │
        │ authorize
        ▼
Signed Playback URL / Token
        │
        ▼
CDN
        │
        ▼
HLS
        │
        ▼
Viewer
```

The API must not proxy video bytes.

Future playback security capabilities may include:

* Expiring URLs
* Signed tokens
* Domain restrictions
* Hotlink protection
* Geo restrictions
* IP restrictions
* Token binding

DRM is intentionally not an MVP requirement.

---

# 20. Webhook Architecture

Webhooks are asynchronous.

Example:

```text
Video State Change
       │
       ▼
Outbox Event
       │
       ▼
Webhook Queue
       │
       ▼
Webhook Worker
       │
       ▼
Customer Endpoint
```

Webhook delivery should support:

* retries
* exponential backoff
* timeouts
* signatures
* idempotency
* delivery status
* dead-letter handling

Example event:

```text
video.ready
```

The customer should be able to safely process the same event more than once.

---

# 21. Analytics Architecture

Player and platform analytics should be asynchronous.

```text
Player
   │
   ▼
Analytics API
   │
   ▼
Queue
   │
   ▼
Analytics Worker
   │
   ▼
PostgreSQL
```

PostgreSQL is sufficient for the initial implementation.

If analytics volume eventually exceeds PostgreSQL's practical workload, the analytics storage layer can move to a specialized analytical database such as ClickHouse without changing the core product model.

---

# 22. Reliability Architecture

Long-running and external operations must be designed for failure.

Important mechanisms include:

* Idempotency
* Retries
* Exponential backoff
* Dead-letter queues
* Job leases
* Timeouts
* Heartbeats
* Graceful shutdown
* Context cancellation
* Concurrency limits

Workers must assume that:

```text
jobs can be duplicated
workers can crash
networks can fail
external services can timeout
```

Therefore operations should be safe to retry wherever practical.

---

# 23. Observability

StreamForge should provide three primary observability layers.

## Logs

Structured application logs should include identifiers such as:

```text
request_id
organization_id
project_id
video_id
job_id
```

Sensitive information must not be logged.

## Metrics

Important metrics include:

```text
API latency
API error rate
queue depth
job duration
processing duration
FFmpeg failures
storage failures
webhook success rate
playback errors
upload-to-ready duration
```

## Traces

Distributed tracing should eventually connect:

```text
API request
   │
   ▼
database operation
   │
   ▼
queue
   │
   ▼
worker
   │
   ▼
storage
```

This is particularly important for diagnosing long video-processing workflows.

---

# 24. Product Deployment Modes

StreamForge supports multiple infrastructure ownership models.

## Managed

StreamForge operates the infrastructure.

```text
Customer
   │
   ▼
StreamForge Control Plane
   │
   ├── StreamForge Queue
   ├── StreamForge Worker
   ├── StreamForge Storage
   └── StreamForge CDN
```

This is the simplest customer experience.

---

## BYO Storage

The customer owns the object storage.

```text
Customer
   │
   ▼
StreamForge Control Plane
   │
   ▼
StreamForge Worker
   │
   ▼
Customer Storage
```

The customer may provide an S3-compatible bucket or equivalent storage configuration.

---

## BYO Infrastructure

The customer operates the processing infrastructure.

```text
Customer
   │
   ▼
StreamForge Control Plane
   │
   ▼
Customer Worker
   │
   ├── Customer Compute
   │
   └── FFmpeg
   │
   ▼
Customer Storage
```

The StreamForge control plane continues to provide:

* API
* authorization
* metadata
* orchestration
* job management
* webhooks
* usage
* developer experience

The customer provides the processing infrastructure.

---

## Self-hosted / Enterprise

The customer operates StreamForge infrastructure within its own environment.

Potential deployment:

```text
Customer Infrastructure
│
├── StreamForge API
├── StreamForge Workers
├── PostgreSQL
├── Redis/Valkey
├── Object Storage
├── CDN
└── Authentication Provider
```

Self-hosted deployment should reuse the same application architecture rather than becoming a separate product codebase.

---

# 25. Worker Architecture

Workers are responsible for executing asynchronous jobs.

The initial worker may support:

```text
video.probe
video.transcode
video.thumbnail
```

Later, customer-operated workers can register their capabilities.

Conceptual worker registration:

```json
{
  "workerId": "wrk_123",
  "version": "1.4.0",
  "capabilities": {
    "h264": true,
    "hevc": true,
    "av1": false,
    "maxConcurrentJobs": 4
  }
}
```

This allows the control plane to select an appropriate worker for a job.

A future CLI may look like:

```bash
streamforge-worker connect \
  --project proj_123 \
  --token ...
```

The exact protocol is intentionally not fixed at this stage.

---

# 26. Configuration and Adapter Selection

Infrastructure providers should be selected through configuration rather than scattered conditionals.

Example:

```yaml
auth:
  provider: kratos

storage:
  provider: s3

queue:
  provider: redis

transcoding:
  provider: ffmpeg
```

Application startup creates the required adapters.

Conceptually:

```text
Configuration
      │
      ▼
Composition Root
      │
      ├── Authenticator
      ├── Storage
      ├── Queue
      ├── Transcoder
      └── CDN
      │
      ▼
Application
```

Business logic should not repeatedly ask:

```go
if provider == "s3" {
    ...
}
```

Provider-specific decisions belong at the composition boundary.

---

# 27. Database Architecture

PostgreSQL is the initial primary database.

It stores:

* Users
* Organizations
* Memberships
* Projects
* API keys
* Videos
* Upload sessions
* Processing jobs
* Webhooks
* Usage
* Analytics metadata
* Outbox events

The database is the source of truth for business state.

Redis/Valkey should not become a second source of truth.

---

# 28. Caching

Caching should be introduced only where there is a demonstrated performance need.

Potential cache targets include:

* frequently accessed project configuration
* authorization metadata
* API rate-limit state
* temporary processing state
* expensive read models

Cached data must always have a clear invalidation or expiration strategy.

The system must remain correct if the cache is empty.

---

# 29. Security Principles

Security is part of the architecture rather than an afterthought.

Important principles:

* Authenticate every protected request.
* Authorize every tenant-scoped resource.
* Never trust project IDs supplied by clients without checking access.
* Hash API secrets.
* Show secrets only once.
* Support key rotation and revocation.
* Use signed URLs for protected playback.
* Use short-lived upload credentials.
* Avoid logging credentials or tokens.
* Validate uploaded media.
* Limit processing resources.
* Apply request and upload rate limits.
* Isolate worker execution where necessary.
* Treat customer-provided infrastructure credentials as sensitive.

---

# 30. API Design Principles

The API should be:

* REST-oriented
* versioned
* predictable
* idempotent where appropriate
* resource-oriented
* documented with OpenAPI

Initial version:

```text
/v1/...
```

Example resources:

```text
/v1/organizations
/v1/projects
/v1/videos
/v1/uploads
/v1/webhooks
/v1/playback
```

API responses should use stable public identifiers.

Breaking API changes require an explicit versioning strategy.

---

# 31. Idempotency

Operations that can be retried by clients should support idempotency where appropriate.

For example:

```http
Idempotency-Key: 01J...
```

This is particularly important for:

* creating videos
* completing uploads
* starting processing
* billing-related operations
* webhook handling

The objective is to prevent network retries from accidentally creating duplicate resources or jobs.

---

# 32. Scaling Strategy

StreamForge should scale incrementally.

Initial architecture:

```text
                 Load Balancer
                       │
               ┌───────┴───────┐
               ▼               ▼
             API 1           API 2
               │               │
               └───────┬───────┘
                       │
                 PostgreSQL
                       │
                    Redis
                       │
              ┌────────┴────────┐
              ▼                 ▼
          Worker 1           Worker 2
```

API servers are horizontally scalable.

Workers can scale independently based on processing workload.

Different worker types can eventually have separate scaling policies:

```text
Transcoding Workers
Thumbnail Workers
Webhook Workers
Analytics Workers
```

Services should not be split prematurely.

---

# 33. When to Introduce Microservices

The initial architecture is a modular distributed monolith.

A component should become a separate service only when there is a concrete reason such as:

* independent scaling requirements
* independent deployment requirements
* strong fault isolation requirements
* separate ownership/team boundaries
* materially different runtime requirements
* infrastructure isolation requirements

For example, transcoding workers naturally have different resource requirements from API servers.

That justifies a separate worker process early.

Kafka, multiple databases, and dozens of microservices do not.

---

# 34. Architectural Boundaries

The following boundaries should remain explicit:

```text
Business Logic
      │
      ▼
Capability Interfaces
      │
      ▼
Infrastructure Adapters
      │
      ▼
External Systems
```

Examples:

```text
Video Domain
    │
    ▼
Storage
    │
    ▼
S3-compatible storage
```

```text
Processing Domain
    │
    ▼
Transcoder
    │
    ▼
FFmpeg
```

```text
Authentication
    │
    ▼
Authenticator
    │
    ▼
ORY Kratos
```

These boundaries allow StreamForge to support multiple deployment models without maintaining completely different application architectures.

---

# 35. MVP Vertical Slice

The first complete vertical slice should be:

```text
Create Project
      │
      ▼
Create Video
      │
      ▼
Create Upload Session
      │
      ▼
Direct Upload
      │
      ▼
Complete Upload
      │
      ▼
Queue Processing
      │
      ▼
FFmpeg
      │
      ▼
Generate HLS
      │
      ▼
Store Assets
      │
      ▼
Video READY
      │
      ▼
Generate Playback URL
      │
      ▼
CDN
      │
      ▼
External Website
      │
      ▼
Play Video
      │
      ▼
video.ready Webhook
```

This vertical slice should be considered the core proof that the architecture works.

---

# 36. Initial Technology Choices

The initial implementation uses:

| Concern         | Initial Technology           |
| --------------- | ---------------------------- |
| Backend         | Go                           |
| API             | HTTP / REST                  |
| Database        | PostgreSQL                   |
| Queue           | Redis / Valkey               |
| Object Storage  | S3-compatible                |
| Transcoding     | FFmpeg                       |
| Authentication  | ORY Kratos adapter           |
| Video Delivery  | HLS                          |
| CDN             | Pluggable                    |
| API Contract    | OpenAPI                      |
| Frontend        | TypeScript / React           |
| Package Manager | pnpm                         |
| Monorepo        | pnpm + Turborepo             |
| Deployment      | Docker initially             |
| Infrastructure  | Terraform / Kubernetes later |

These are implementation choices, not permanent architectural dependencies.

---

# 37. Explicit Non-Goals

The following are intentionally not part of the initial architecture:

* Kafka-first architecture
* Large microservice fleet
* Database-per-tenant
* Custom video codec implementation
* Custom CDN
* Custom identity provider
* DRM as an MVP requirement
* Complex workflow engine
* Distributed event sourcing
* Multiple database engines
* Premature multi-region infrastructure

These may become necessary later, but should be introduced because of concrete requirements rather than architectural fashion.

---

# 38. Architectural Decision Summary

StreamForge follows these core rules:

```text
1. API owns control.
2. Storage owns bytes.
3. Workers own long-running processing.
4. CDN owns delivery.
5. PostgreSQL owns business state.
6. Redis/Valkey handles asynchronous work.
7. Interfaces represent capabilities.
8. Adapters implement infrastructure.
9. Configuration selects adapters.
10. Business logic never depends directly on vendors.
11. API never proxies large media.
12. Everything is tenant-aware.
13. Jobs are retryable and idempotent.
14. Events are asynchronous.
15. Start modular; split services when justified.
```

The architecture should evolve as StreamForge scales, but these principles should remain stable.

---

# 39. Architectural North Star

The long-term architecture is:

```text
                         STREAMFORGE
                              │
               ┌──────────────┴──────────────┐
               │                             │
          CONTROL PLANE                 CUSTOMER APPS
               │                             │
       ┌───────┼────────┐                    │
       │       │        │                    │
      API   Metadata   Auth                  │
       │       │        │                    │
       └───────┼────────┘                    │
               │                             │
         Orchestration                       │
               │                             │
        ┌──────┴──────┐                      │
        │             │                      │
     Managed       Customer                  │
     Workers       Workers                   │
        │             │                      │
        └──────┬──────┘                      │
               │                             │
          Processing                         │
               │                             │
          Object Storage                     │
               │                             │
              HLS                            │
               │                             │
              CDN ◄──────────────────────────┘
               │
               ▼
             Viewer
```

StreamForge's strategic advantage is not owning every piece of infrastructure.

It is owning the **control plane that makes video infrastructure programmable** while allowing the underlying infrastructure to evolve from managed resources to customer-owned storage, customer-owned compute, and eventually fully self-hosted deployments.
