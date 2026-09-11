# StreamForge

**Composable video infrastructure for developers.**

StreamForge is a pluggable video infrastructure platform that gives developers the building blocks to **ingest, process, encode, store, deliver, and manage video** without being locked into a single cloud provider, storage backend, codec, or deployment model.

It is designed to work as a **managed service, BYO infrastructure platform, BYO storage platform, or fully self-hosted video stack**.

---

## Why StreamForge?

Building video infrastructure from scratch requires solving a large number of problems:

* Video ingestion
* Transcoding and encoding
* Codec management
* Media processing
* Storage
* CDN and delivery
* Authentication and authorization
* Webhooks and events
* Job orchestration
* Monitoring
* Infrastructure provisioning
* Multi-tenancy
* API and SDK design

StreamForge provides these capabilities behind a consistent developer-facing platform while keeping the underlying infrastructure **replaceable and extensible**.

The goal is not to create another closed video platform.

The goal is to provide a **video infrastructure layer developers can own and control**.

---

## Core Principles

### Pluggable by design

Infrastructure components should be replaceable.

Storage, encoders, codecs, queues, databases, authentication providers, and delivery systems should not be permanently coupled to the core platform.

### Developer-first

Everything should be accessible through APIs, SDKs, CLI tooling, and infrastructure configuration.

### Infrastructure agnostic

StreamForge should work across:

* AWS
* GCP
* Azure
* Hetzner
* DigitalOcean
* Kubernetes
* Bare metal
* Private infrastructure

### Data ownership

Users should be able to control where their video data lives.

### Deployment freedom

The same platform should support managed and self-hosted deployments without requiring application-level rewrites.

### Open architecture

Core interfaces should remain open and implementation-independent.

---

# Product Modes

StreamForge supports multiple infrastructure models.

| Mode                   | Infrastructure | Storage                 | Best for                   |
| ---------------------- | -------------- | ----------------------- | -------------------------- |
| **Managed**            | StreamForge    | StreamForge             | Fastest way to get started |
| **BYO Storage**        | StreamForge    | Customer                | Data ownership             |
| **BYO Infrastructure** | Customer       | Customer / configurable | Infrastructure control     |
| **Self-hosted**        | Customer       | Customer                | Full ownership             |

### Managed

StreamForge operates the infrastructure.

```text
Application
    |
    v
StreamForge API
    |
    +--> Ingest
    +--> Processing
    +--> Storage
    +--> Delivery
```

This is the simplest way to integrate StreamForge.

---

### BYO Storage

StreamForge manages the video infrastructure while the customer controls the storage layer.

Examples:

* Amazon S3
* Cloudflare R2
* Google Cloud Storage
* Azure Blob Storage
* MinIO
* S3-compatible storage

```text
Application
    |
    v
StreamForge
    |
    +--> Processing
    |
    +--> Customer Storage
```

---

### BYO Infrastructure

The customer provides the underlying compute and infrastructure while StreamForge provides the control and video platform layer.

This allows organizations to maintain control over:

* Compute
* Networking
* Kubernetes
* GPUs
* Storage
* Databases
* Queues

---

### Self-hosted

The entire StreamForge stack can run inside the customer's infrastructure.

```text
Customer Infrastructure

+-------------------------------+
|          StreamForge          |
|                               |
| API                           |
| Control Plane                 |
| Processing                    |
| Encoding                      |
| Delivery                      |
| Storage adapters              |
+-------------------------------+
```

---

# Architecture

StreamForge is organized around three major planes:

```text
                    +----------------+
                    |   Application  |
                    +-------+--------+
                            |
                            v
                    +---------------+
                    |  Control Plane|
                    +-------+-------+
                            |
             +--------------+--------------+
             |                             |
             v                             v
      +-------------+               +-------------+
      |  Data Plane |               |Delivery     |
      |             |               |Plane        |
      +-------------+               +-------------+
             |                             |
             v                             v
        Processing                    CDN / Playback
             |
             v
          Storage
```

## Control Plane

The **control plane** manages the system.

Responsibilities include:

* Projects
* Organizations
* Users
* API keys
* Authentication
* Authorization
* Video metadata
* Assets
* Processing jobs
* Workflows
* Configuration
* Webhooks
* Billing
* Provider configuration

The control plane answers:

> **What should happen?**

---

## Data Plane

The **data plane** performs the actual video work.

Responsibilities include:

* Uploading
* Ingestion
* Demuxing
* Transcoding
* Encoding
* Packaging
* Thumbnail generation
* Metadata extraction
* Media analysis
* Quality processing

The data plane answers:

> **How is the media processed?**

---

## Delivery Plane

The **delivery plane** makes processed video available to end users.

Responsibilities include:

* HLS
* DASH
* Progressive download
* CDN integration
* Signed URLs
* Access control
* Playback manifests
* Edge caching
* Origin management

The delivery plane answers:

> **How does the video reach the viewer?**

---

# Pluggable Architecture

StreamForge is built around interfaces rather than infrastructure implementations.

For example:

```text
Storage
   |
   +-- S3
   +-- GCS
   +-- Azure Blob
   +-- R2
   +-- MinIO
   +-- Local filesystem

Encoder
   |
   +-- FFmpeg
   +-- Hardware encoder
   +-- Custom encoder

Codec
   |
   +-- H.264
   +-- H.265
   +-- AV1
   +-- Custom codec

Queue
   |
   +-- Redis
   +-- NATS
   +-- Kafka
   +-- SQS

Database
   |
   +-- PostgreSQL
   +-- MySQL
   +-- Other adapters
```

The core platform should depend on **contracts**, not implementations.

---

# Video Pipeline

A typical StreamForge workflow looks like:

```text
                Upload
                  |
                  v
              Ingestion
                  |
                  v
             Validation
                  |
                  v
          Media Inspection
                  |
                  v
             Processing
                  |
                  v
              Encoding
                  |
                  v
              Packaging
                  |
                  v
               Storage
                  |
                  v
               Delivery
                  |
                  v
               Playback
```

Each stage can be independently extended or replaced.

---

# Processing

StreamForge supports programmable video workflows.

Example:

```yaml
workflow:
  input:
    type: video

  processing:
    - probe
    - thumbnail

  encoding:
    - codec: h264
      resolution: 1080p
      bitrate: 5000k

    - codec: h264
      resolution: 720p
      bitrate: 3000k

    - codec: h264
      resolution: 480p
      bitrate: 1500k

  packaging:
    - hls

  delivery:
    - cdn
```

Future workflows can support:

* AV1
* H.265
* HDR
* Audio normalization
* Subtitle extraction
* Scene detection
* AI processing
* Watermarking
* Content moderation
* Custom processing stages

---

# Custom Video Codecs

StreamForge is designed to allow codec implementations to be introduced without changing the core platform.

A codec implementation should conform to a common interface:

```text
Codec
├── name()
├── capabilities()
├── validate()
├── encode()
└── metadata()
```

This makes it possible to support:

* Software codecs
* Hardware codecs
* GPU encoders
* Cloud encoding services
* Custom codecs

without coupling the application layer to a specific encoder.

---

# Storage

Storage is accessed through an abstraction layer.

```text
StorageAdapter

    put()
    get()
    delete()
    exists()
    signedURL()
    metadata()
```

Implementations may include:

* S3
* Cloudflare R2
* Google Cloud Storage
* Azure Blob Storage
* MinIO
* Local filesystem

This allows customers to change storage providers without changing their application integration.

---

# Authentication & Authorization

StreamForge separates identity from authorization.

Authentication can be provided by an identity system such as **Ory Kratos**, while authorization can be handled independently.

The architecture can support:

```text
Identity
   |
   v
Authentication
   |
   v
Authorization
   |
   v
Resource Access
```

This separation allows StreamForge to integrate with different identity providers and authorization systems without coupling the application to a single authentication implementation.

---

# Multi-tenancy

StreamForge is designed around organizations, projects, and resources.

A typical hierarchy is:

```text
Organization
    |
    +-- Project
          |
          +-- Videos
          +-- Assets
          +-- Encodings
          +-- Webhooks
          +-- API Keys
          +-- Storage
          +-- Processing Jobs
```

This provides isolation while allowing organizations to manage multiple applications or environments.

---

# API

The StreamForge API is the primary interface for applications.

Example:

```http
POST /v1/videos
```

```json
{
  "filename": "example.mp4",
  "workflow": "standard"
}
```

Response:

```json
{
  "id": "vid_123",
  "status": "processing"
}
```

The API should remain independent of the underlying infrastructure.

---

# Events

StreamForge uses events to communicate asynchronous state changes.

Example:

```text
video.created
video.uploaded
video.processing.started
video.processing.completed
video.processing.failed
video.encoding.started
video.encoding.completed
video.ready
video.deleted
```

Events can be delivered through:

* Webhooks
* Message queues
* Event streams
* Internal event buses

Example:

```json
{
  "event": "video.ready",
  "id": "evt_123",
  "data": {
    "video_id": "vid_123"
  }
}
```

---

# Webhooks

Applications can subscribe to lifecycle events.

```text
StreamForge
     |
     v
Webhook
     |
     v
Customer Application
```

Webhooks should support:

* Signing
* Retries
* Idempotency
* Delivery history
* Failure handling
* Replay

---

# Project Structure

The repository is organized around platform boundaries rather than infrastructure vendors.

```text
streamforge/
│
├── api/
│   ├── openapi/
│   └── proto/
│
├── cmd/
│   ├── api/
│   ├── worker/
│   └── cli/
│
├── internal/
│   ├── controlplane/
│   ├── dataplane/
│   ├── delivery/
│   ├── auth/
│   ├── authorization/
│   ├── media/
│   ├── workflow/
│   ├── events/
│   └── tenancy/
│
├── pkg/
│   ├── storage/
│   ├── encoder/
│   ├── codec/
│   ├── queue/
│   └── provider/
│
├── deployments/
│   ├── docker/
│   ├── kubernetes/
│   └── helm/
│
├── docs/
│   ├── ARCHITECTURE.md
│   ├── DEVELOPMENT.md
│   └── CONTRIBUTING.md
│
├── examples/
│
├── scripts/
│
├── tests/
│
├── LICENSE
├── NOTICE
└── README.md
```

The exact structure may evolve as the implementation matures.

---

# Deployment

StreamForge can be deployed using containers or Kubernetes.

### Docker

```bash
docker compose up
```

### Kubernetes

```bash
helm install streamforge ./deployments/helm/streamforge
```

Production deployments can separate:

```text
API
Workers
Queues
Database
Storage
Delivery
```

allowing each component to scale independently.

---

# Local Development

Clone the repository:

```bash
git clone https://github.com/streamforge/streamforge.git

cd streamforge
```

Start the development environment:

```bash
docker compose up
```

Run tests:

```bash
make test
```

Build:

```bash
make build
```

Run the API:

```bash
make dev
```

---

# Configuration

StreamForge should be configurable through environment variables and configuration files.

Example:

```env
STREAMFORGE_ENV=development

DATABASE_URL=postgres://localhost/streamforge

STORAGE_DRIVER=s3

S3_ENDPOINT=
S3_BUCKET=streamforge

QUEUE_DRIVER=redis

REDIS_URL=redis://localhost:6379

AUTH_DRIVER=ory
```

Provider-specific configuration should remain isolated from the core application.

---

# Observability

StreamForge should provide first-class observability.

Supported capabilities include:

* Structured logging
* Metrics
* Distributed tracing
* Processing job metrics
* Queue metrics
* Storage metrics
* Encoding metrics
* Delivery metrics

Recommended standards:

* OpenTelemetry
* Prometheus
* OpenTelemetry-compatible tracing

---

# Security

Security is a core architectural concern.

StreamForge should support:

* API key authentication
* OAuth/OIDC
* Signed URLs
* Resource-level authorization
* Project isolation
* Organization isolation
* Webhook signatures
* Encryption in transit
* Encryption at rest
* Secret management
* Audit logs

Self-hosted deployments should allow organizations to integrate their existing identity and security infrastructure.

---

# Extensibility

StreamForge is intended to become an ecosystem rather than a monolithic implementation.

Potential extension points include:

```text
Storage Providers
Encoder Providers
Codec Providers
Queue Providers
Database Providers
Auth Providers
Authorization Providers
CDN Providers
Processing Plugins
AI Providers
Webhook Providers
```

The preferred extension model is:

```text
Core Interface
      |
      +---- Provider A
      +---- Provider B
      +---- Provider C
```

rather than:

```text
Core
 |
 +---- Provider A-specific logic
 +---- Provider B-specific logic
 +---- Provider C-specific logic
```

---

# Roadmap

## Foundation

* [ ] Core API
* [ ] Project management
* [ ] Video assets
* [ ] Storage abstraction
* [ ] Processing jobs
* [ ] Worker architecture
* [ ] PostgreSQL support
* [ ] S3-compatible storage
* [ ] Docker development environment

## Video

* [ ] FFmpeg integration
* [ ] H.264 encoding
* [ ] HLS packaging
* [ ] Thumbnail generation
* [ ] Media probing
* [ ] Multi-bitrate encoding
* [ ] Hardware acceleration

## Platform

* [ ] Authentication
* [ ] Authorization
* [ ] Organizations
* [ ] API keys
* [ ] Webhooks
* [ ] Usage metering
* [ ] Audit logs

## Infrastructure

* [ ] Kubernetes deployment
* [ ] Helm chart
* [ ] BYO storage
* [ ] BYO infrastructure
* [ ] Self-hosted distribution

## Advanced Video

* [ ] H.265
* [ ] AV1
* [ ] HDR
* [ ] Subtitles
* [ ] Watermarking
* [ ] Content analysis
* [ ] Custom codec plugins
* [ ] GPU encoding

---

# Design Goals

StreamForge is ultimately trying to achieve five things:

### 1. Make video infrastructure programmable

Developers should be able to define video workflows through APIs and configuration.

### 2. Make infrastructure replaceable

A customer should be able to replace storage, encoding, queues, or infrastructure without rewriting their application.

### 3. Make deployment portable

The same architecture should work in a managed environment, cloud infrastructure, private infrastructure, or a customer's own Kubernetes cluster.

### 4. Give customers control

Customers should be able to control their data, infrastructure, and providers when required.

### 5. Keep the developer experience simple

Complex infrastructure should be hidden behind a clean API.

---

# Contributing

Contributions are welcome.

Before submitting a pull request:

1. Read `ARCHITECTURE.md`.
2. Understand the relevant platform boundary.
3. Keep provider-specific logic behind interfaces.
4. Add tests for new functionality.
5. Update documentation when behavior or architecture changes.

See `CONTRIBUTING.md` for development guidelines.

---

# License

StreamForge is distributed under the license specified in [`LICENSE`](LICENSE).

See [`NOTICE`](NOTICE) for third-party attribution and licensing information.

---

# Status

> **StreamForge is under active development.**

APIs, interfaces, deployment configuration, and internal architecture may change before the first stable release.

The project currently prioritizes **architecture, extensibility, and infrastructure portability** over backward compatibility.

---

## Vision

StreamForge aims to become the infrastructure layer developers use when they need video without wanting to build and operate an entire video platform themselves.

```text
                 Your Application
                        |
                        v
                +---------------+
                |   StreamForge |
                +-------+-------+
                        |
        +---------------+---------------+
        |               |               |
        v               v               v
     Ingest          Process         Deliver
        |               |               |
        +---------------+---------------+
                        |
                 Pluggable Providers
                        |
       +--------+-------+-------+--------+
       |        |       |       |        |
    Storage  Encoder   Queue    CDN     Auth
```

**Build your video product. Own your video infrastructure.**
