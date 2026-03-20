# Database Tester

A simple Go application for testing PostgreSQL and MongoDB database connectivity.

## Description

This tool connects to a PostgreSQL or MongoDB database and verifies connectivity. The database type is auto-detected from the connection URI — no extra configuration needed.

## Features

- Auto-detect database type from connection URI (`mongodb://`, `mongodb+srv://`, or PostgreSQL)
- Connect to PostgreSQL using connection string or individual parameters
- Connect to MongoDB using a URI (standalone, replica set, or Atlas)
- Support for environment variable configuration
- Built with GORM (PostgreSQL) and the official MongoDB Go driver v2 (MongoDB)

## Installation

### Download Pre-built Binaries

Pre-built binaries for various platforms are available on the [Releases](https://github.com/dxas90/pg-database-tester/releases) page.

Download the appropriate binary for your platform:
- Linux (amd64, arm64)
- macOS (amd64, arm64)
- Windows (amd64)

### Build from Source

Requires Go 1.24 or later.

```bash
git clone https://github.com/dxas90/pg-database-tester.git
cd pg-database-tester
go build -o pg-db-tester
```

## Usage

### MongoDB (URI detection is automatic)

```bash
# Atlas (SRV)
export DATABASE_URL="mongodb+srv://user:password@cluster.mongodb.net/dbname"
./pg-db-tester

# Standalone / replica set
export DATABASE_URL="mongodb://user:password@localhost:27017/dbname"
./pg-db-tester
```

### PostgreSQL — Full Connection String

```bash
export DATABASE_URL="postgresql://user:password@localhost:5432/dbname?sslmode=disable"
./pg-db-tester
```

### PostgreSQL — Individual Environment Variables

```bash
export DB_USER="postgres"
export DB_PASSWORD="password"
export DB_HOST="localhost"
export DB_PORT="5432"
export DB_NAME="postgres"
export DB_SSLMODE="disable"
./pg-db-tester
```

## Environment Variables

### Shared

| Variable | Description | Default |
|----------|-------------|---------|
| `DATABASE_URL` | Full connection URI (PostgreSQL or MongoDB) | - |

> When `DATABASE_URL` starts with `mongodb://` or `mongodb+srv://`, MongoDB mode is used automatically.
> When `DATABASE_URL` is not set, the PostgreSQL individual variables below are used.

### PostgreSQL (individual variables)

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_USER` | Database username | `postgres` |
| `DB_PASSWORD` | Database password | `password` |
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `5432` |
| `DB_NAME` | Database name | `postgres` |
| `DB_SSLMODE` | SSL mode (`disable`, `require`, `verify-ca`, `verify-full`) | `disable` |

## Example Output

```
# PostgreSQL
Current time from DB: 2025-10-10 12:42:27.123456 +0000 UTC

# MongoDB
Successfully connected to MongoDB
```

## Dependencies

- [GORM](https://gorm.io/) — ORM for PostgreSQL
- [pgx](https://github.com/jackc/pgx) — PostgreSQL driver
- [MongoDB Go Driver v2](https://github.com/mongodb/mongo-go-driver) — Official MongoDB driver

## License

This project is open source and available under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
