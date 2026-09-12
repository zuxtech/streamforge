# Infrastructure

Infrastructure configuration and deployment resources for StreamForge.

This directory contains the infrastructure required to run, develop, deploy, and operate StreamForge.

## Structure

```text
infrastructure/
├── docker/
│   ├── compose.yml
│   ├── dev.yml
│   ├── .env
│   ├── .env.example
│   ├── config/
│   │   ├── kratos/
│   │   │   └── kratos.yml
│   │   └── streamforge/
│   │       └── config.yml
│   └── volumes/
│       └── postgres/
│           └── init/
│               └── 01-schemas.sql
│
├── kubernetes/
│   └── ...
│
└── terraform/
    └── ...
```

## Docker

The `docker/` directory contains the local Docker Compose environment.

It is primarily intended for:

* Local development
* Running infrastructure dependencies
* Testing infrastructure integrations
* Running StreamForge in simple self-hosted environments

### Services

The development environment currently includes:

```text
PostgreSQL
    │
    ├── StreamForge database
    │
    └── Kratos schema

ORY Kratos
    │
    └── Authentication

Kratos Self-Service UI
    │
    └── Development authentication UI
```

Additional services such as Redis, object storage, and video processing workers can be added as the platform evolves.

## Configuration

Docker configuration is separated into three categories.

### Environment

```text
docker/.env
```

Contains environment-specific values and secrets.

Example:

```env
POSTGRES_USER=streamforge
POSTGRES_PASSWORD=change-me
POSTGRES_DB=streamforge

KRATOS_UI_COOKIE_SECRET=change-me
KRATOS_UI_CSRF_COOKIE_SECRET=change-me
```

Do not commit `.env`.

Use `.env.example` as the template for local configuration.

### Service configuration

Static service configuration belongs under:

```text
docker/config/
```

For example:

```text
config/
├── kratos/
│   └── kratos.yml
└── streamforge/
    └── config.yml
```

Kratos configuration is kept separate from StreamForge configuration because Kratos is an external infrastructure component.

### Initialization

PostgreSQL initialization scripts are stored under:

```text
docker/volumes/postgres/init/
```

These scripts run automatically when PostgreSQL initializes a new database.

The current initialization creates the application schemas:

```text
streamforge
kratos
```

The `streamforge` schema is owned by StreamForge.

The `kratos` schema is owned by ORY Kratos.

StreamForge must not access Kratos tables directly. Communication with Kratos happens through its API.

## Local Development

The recommended local development workflow is:

```text
Docker Compose
     │
     ├── PostgreSQL
     ├── ORY Kratos
     └── Kratos UI
     
Local machine
     │
     └── StreamForge Go API
```

Start infrastructure:

```bash
make docker-dev
```

Run the StreamForge API:

```bash
make dev
```

Stop infrastructure:

```bash
make docker-down
```

View logs:

```bash
make docker-logs
```

Alternatively, Docker Compose can be invoked directly:

```bash
docker compose \
  --env-file infrastructure/docker/.env \
  -f infrastructure/docker/compose.yml \
  -f infrastructure/docker/dev.yml \
  up -d
```

## Docker Compose Files

### `compose.yml`

Contains the base infrastructure definition.

This should contain configuration that is common across environments.

### `dev.yml`

Contains development-specific configuration.

Examples include:

* Development ports
* Development-only services
* Local volume mounts
* Development environment overrides
* Debug configuration

Keeping development overrides separate makes it possible to reuse the base Compose configuration for other environments.

## PostgreSQL

StreamForge currently uses PostgreSQL as its primary relational database.

The development environment uses a single PostgreSQL database with separate schemas:

```text
PostgreSQL
└── streamforge
    ├── streamforge
    │   └── StreamForge tables
    │
    └── kratos
        └── Kratos tables
```

This provides logical separation while keeping local development simple.

The architecture can later move Kratos to a separate database without requiring StreamForge to access Kratos's internal tables.

## ORY Kratos

Kratos is used for authentication and identity management.

The infrastructure configuration is located at:

```text
docker/config/kratos/kratos.yml
```

The development services expose:

| Service                |   Port | Purpose                       |
| ---------------------- | -----: | ----------------------------- |
| Kratos Public API      | `4433` | Browser/API-facing Kratos API |
| Kratos Admin API       | `4434` | Internal/admin API            |
| Kratos Self-Service UI | `4455` | Development authentication UI |

The StreamForge API communicates with Kratos through its public API rather than directly accessing the Kratos database.

## StreamForge Configuration

Application configuration is kept separate from infrastructure configuration.

Docker:

```text
docker/config/streamforge/config.yml
```

Go application:

```text
internal/platform/config/
```

The Docker environment provides infrastructure-specific values, while the StreamForge configuration defines application behavior.

## Volumes

Persistent Docker data is stored using named volumes where appropriate.

For example:

```yaml
volumes:
  postgres_data:
    name: sf_postgres
```

PostgreSQL data is mounted at the configured PostgreSQL data directory.

Initialization scripts are mounted read-only:

```yaml
- ./volumes/postgres/init:/docker-entrypoint-initdb.d:ro
```

Do not manually modify database files inside Docker volumes.

## Resetting Local Infrastructure

To stop the environment:

```bash
make docker-down
```

To remove containers and persistent volumes:

```bash
docker compose \
  --env-file infrastructure/docker/.env \
  -f infrastructure/docker/compose.yml \
  -f infrastructure/docker/dev.yml \
  down -v
```

**Warning:** removing volumes deletes local PostgreSQL data.

After removing the database volume, starting the environment again will execute the PostgreSQL initialization scripts.

## Secrets

Secrets should never be committed to the repository.

Use:

```text
docker/.env
```

for local development secrets.

Commit only:

```text
docker/.env.example
```

Example:

```env
POSTGRES_USER=streamforge
POSTGRES_PASSWORD=
POSTGRES_DB=streamforge

KRATOS_UI_COOKIE_SECRET=
KRATOS_UI_CSRF_COOKIE_SECRET=
```

Production deployments should use the secret-management facilities of the target platform rather than committing secrets to configuration files.

## Production

Production infrastructure will support multiple deployment models.

### StreamForge Cloud

StreamForge-managed infrastructure.

```text
Customer
   │
   ▼
StreamForge
   │
   ├── API
   ├── Workers
   ├── PostgreSQL
   ├── Object Storage
   ├── Queue
   └── CDN
```

### Self-hosted

Customers can deploy StreamForge into their own infrastructure.

For smaller installations, Docker Compose may be sufficient.

For larger installations:

```text
Kubernetes
    │
    ├── StreamForge API
    ├── StreamForge Workers
    ├── PostgreSQL
    ├── Queue
    ├── Object Storage
    └── Delivery
```

Kubernetes and Terraform configurations will live under their respective directories.

## Design Principles

Infrastructure configuration should follow these principles:

### Infrastructure is replaceable

StreamForge should not depend on a single cloud provider or infrastructure vendor.

### Configuration is separated from secrets

Static configuration belongs in configuration files.

Secrets belong in environment variables or a secret manager.

### Development should be reproducible

A new developer should be able to start the required infrastructure with a small number of commands.

### Production should not depend on development configuration

Development Compose files are not production deployment definitions.

Production deployments should use the appropriate Kubernetes, Terraform, or managed infrastructure configuration.

### Application code should remain infrastructure-agnostic

The application should depend on interfaces and capabilities rather than directly depending on PostgreSQL, Redis, S3, FFmpeg, or a particular cloud provider.

## Future Infrastructure

As StreamForge develops, this directory is expected to contain configuration for:

* Redis / Valkey
* Object storage
* FFmpeg workers
* CDN
* Kubernetes
* Helm
* Terraform
* Observability
* Production secrets
* Cloud providers
* Self-hosted deployments

The infrastructure layer should evolve without coupling the StreamForge application to a specific deployment environment.
