# PostgreSQL Database Tester

A simple Go application for testing PostgreSQL database connectivity and performing basic queries.

## Description

This tool connects to a PostgreSQL database and executes a simple query to retrieve the current time from the database server. It's useful for verifying database connectivity, testing connection parameters, and as a minimal example of using GORM with PostgreSQL.

## Features

- Connect to PostgreSQL using connection string or individual parameters
- Support for environment variable configuration
- Simple health check by querying database time
- Built with GORM ORM framework

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

### Using Full Connection String

```bash
export DATABASE_URL="postgresql://user:password@localhost:5432/dbname?sslmode=disable"
./pg-db-tester
```

### Using Individual Environment Variables

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

| Variable | Description | Default |
|----------|-------------|---------|
| `DATABASE_URL` | Full PostgreSQL connection string | - |
| `DB_USER` | Database username | `postgres` |
| `DB_PASSWORD` | Database password | `password` |
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `5432` |
| `DB_NAME` | Database name | `postgres` |
| `DB_SSLMODE` | SSL mode (disable, require, verify-ca, verify-full) | `disable` |

**Note:** If `DATABASE_URL` is set, it takes precedence over individual variables.

## Example Output

```
Current time from DB: 2025-10-10 12:42:27.123456 +0000 UTC
```

## Dependencies

- [GORM](https://gorm.io/) - The fantastic ORM library for Golang
- [PostgreSQL Driver](https://github.com/jackc/pgx) - PostgreSQL driver and toolkit for Go

## License

This project is open source and available under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
