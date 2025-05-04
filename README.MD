# DYEL (Do You Even Lift)

A simple Go + Gin + GORM microservice for managing workout programs, backed by PostgreSQL and Docker Compose.

## Features

- **Create** a program with days and exercises
- **Retrieve** a single program by ID
- **List** all programs
- **Update** (full replace) an existing program
- **Delete** a program
- **Persistent storage** in PostgreSQL via GORM
- **Containerized** with Docker & Docker Compose

---

## 📦 Getting Started

### Prerequisites

- [Docker & Docker Compose](https://docs.docker.com/compose/install/)

### Clone & Configure

```bash
git clone https://github.com/iraunchy/dyel.git
cd dyel
cp .env.example .env
# Edit .env to set your db credentials, e.g.:
# POSTGRES_USER=
# POSTGRES_PASSWORD=
```

### Build and Run

```bash
  docker-compose build
  docker-compose up -d
```
* The API will be available at http://localhost/
* PostgreSQL is running in the db container

### Sample payload

```json
{
  "name": "Full Body Blast",
  "shared_by": "jane.doe@example.com",
  "days": [
    {
      "name": "Day 1 – Push",
      "exercises": [
        { "name": "Bench Press", "sets": 4, "reps": "8-10", "rest": "90s" },
        { "name": "Overhead Press", "sets": 3, "reps": "6-8", "rest": "90s" }
      ]
    }
  ]
}
```