# Cloud-Native Microservices E-Commerce

> Event-driven microservices architecture with Go, gRPC, Kafka, PostgreSQL, Redis, and Docker — implementing CQRS and Saga patterns.

## 🚀 Overview

A production-grade e-commerce platform decomposed into independently deployable microservices communicating via gRPC and Kafka events. Implements the Saga pattern for distributed transactions (order → payment → inventory), CQRS for read/write separation, and includes full Docker Compose orchestration.

## ✨ Features

| Feature | Description |
|---------|-------------|
| 🏗️ Microservices | Product, Order, Payment, User services |
| 📡 gRPC Communication | Type-safe inter-service RPC |
| 📨 Kafka Events | Event-driven async communication |
| 🔄 Saga Pattern | Distributed transaction orchestration |
| 🚪 API Gateway | Single entry point with routing |
| 🐳 Docker Compose | One-command full-stack startup |
| 💾 PostgreSQL + Redis | Persistence + caching |
| ☸️ K8s Manifests | Production deployment configs |

## 🛠️ Tech Stack

| Component | Technology |
|-----------|-----------|
| Services | Go |
| RPC | gRPC + Protocol Buffers |
| Events | Apache Kafka |
| Database | PostgreSQL |
| Cache | Redis |
| Container | Docker + Docker Compose |
| Orchestration | Kubernetes + Istio |

## ⚡ Quick Start

```bash
# Start all services with Docker
docker-compose up -d

# Or run individually
cd services/product && go run main.go
cd gateway && go run main.go
```

### Services

| Service | Port | Responsibility |
|---------|------|---------------|
| Product | :50051 | Catalog, inventory |
| Order | :50052 | Order lifecycle, saga |
| Payment | :50053 | Payment processing |
| User | :50054 | Auth, profiles |
| Gateway | :8080 | HTTP → gRPC routing |

## 📄 License

MIT
