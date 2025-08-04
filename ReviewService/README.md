# Review Service

A microservice for managing hotel reviews in the Airbnb-like platform. This service handles the creation, retrieval, and management of user reviews for hotels and bookings.

## 🏗️ Architecture

The Review Service follows a clean architecture pattern with the following layers:

- **Controllers**: Handle HTTP requests and responses
- **Services**: Business logic and validation
- **Repositories**: Data access layer
- **Models**: Data structures
- **DTOs**: Data transfer objects for API communication
- **Clients**: External service communication (Booking Service)
- **Middlewares**: Request validation and processing

## 🚀 Features

- ✅ Create hotel reviews with validation
- ✅ Retrieve reviews by ID
- ✅ Get all reviews
- ✅ Booking validation through Booking Service integration
- ✅ Input validation and error handling
- ✅ Database migrations with Goose
- ✅ RESTful API design

## 📋 Prerequisites

- Go 1.24.4 or higher
- MySQL 8.0 or higher
- Booking Service (for booking validation)


### 3. Environment Configuration

Create a `.env` file in the root directory:

```env
# Server Configuration
PORT=:3004

# Database Configuration
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=auth_dev

# External Services
BOOKING_SERVICE_URL=http://localhost:3000
```

### 4. Database Setup

#### Run Migrations

```bash
# Apply all migrations
make migrate-up

# Check migration status
make migrate-status

# Rollback migrations if needed
make migrate-down
```

#### Available Migration Commands

```bash
# Create a new migration
make migrate-create name="create_new_table"

# Apply migrations
make migrate-up

# Rollback migrations
make migrate-down

# Reset database (rollback all)
make migrate-reset

# Check migration status
make migrate-status

# Redo last migration
make migrate-redo

# Migrate to specific version
make migrate-to version=20250727060514

# Rollback to specific version
make migrate-down-to version=20250727060514

# Force specific migration version
make migrate-force version=20250727060514
```

### 5. Run the Service

```bash
go run main.go
```

The service will start on the configured port (default: `:3004`).

## 📊 Database Schema

### Review Table

```sql
CREATE TABLE review (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    hotel_id INT NOT NULL,
    booking_id INT NOT NULL,
    comment TEXT,
    rating INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    is_synced BOOLEAN NOT NULL DEFAULT FALSE
);
```

## 🔌 API Endpoints

### Base URL
```
http://localhost:3004/api/v1
```

### 1. Create Review

**POST** `/reviews/`

Creates a new review for a hotel booking.

**Request Body:**
```json
{
    "user_id": 123,
    "hotel_id": 456,
    "booking_id": 789,
    "comment": "Great hotel with excellent service!",
    "rating": 5
}
```

**Response:**
```json
{
    "status": "success",
    "message": "Review created successfully",
    "data": {
        "id": 1,
        "user_id": 123,
        "hotel_id": 456,
        "booking_id": 789,
        "comment": "Great hotel with excellent service!",
        "rating": 5,
        "created_at": "2024-01-15T10:30:00Z",
        "updated_at": "2024-01-15T10:30:00Z",
        "is_synced": false
    }
}
```

**Validation Rules:**
- `user_id`: Required, must be a positive integer
- `hotel_id`: Required, must be a positive integer
- `booking_id`: Required, must be a positive integer
- `comment`: Optional, string
- `rating`: Required, must be between 1 and 5

### 2. Get All Reviews

**GET** `/reviews/`

Retrieves all reviews.

**Response:**
```json
{
    "status": "success",
    "message": "Reviews fetched successfully",
    "data": [
        {
            "id": 1,
            "user_id": 123,
            "hotel_id": 456,
            "booking_id": 789,
            "comment": "Great hotel with excellent service!",
            "rating": 5,
            "created_at": "2024-01-15T10:30:00Z",
            "updated_at": "2024-01-15T10:30:00Z",
            "is_synced": false
        }
    ]
}
```

### 3. Get Review by ID

**GET** `/reviews/{id}`

Retrieves a specific review by ID.

**Response:**
```json
{
    "status": "success",
    "message": "Review fetched successfully",
    "data": {
        "id": 1,
        "user_id": 123,
        "hotel_id": 456,
        "booking_id": 789,
        "comment": "Great hotel with excellent service!",
        "rating": 5,
        "created_at": "2024-01-15T10:30:00Z",
        "updated_at": "2024-01-15T10:30:00Z",
        "is_synced": false
    }
}
```

## 🔗 External Service Integration

### Booking Service Integration

The Review Service integrates with the Booking Service to validate bookings before creating reviews:

- **Validation**: Ensures the booking exists and belongs to the user
- **Status Check**: Verifies the booking is in "CONFIRMED" status
- **Client**: Uses HTTP client with timeout configuration

**Configuration:**
```env
BOOKING_SERVICE_URL=http://localhost:3000
```

## 🏛️ Project Structure

```
ReviewService/
├── app/
│   └── application.go          # Application setup and configuration
├── clients/
│   └── booking_client.go       # Booking Service client
├── config/
│   ├── db/
│   │   └── db.go              # Database configuration
│   └── env/
│       └── env.go             # Environment configuration
├── controllers/
│   ├── ping.go                # Health check controller
│   └── review_controller.go   # Review operations controller
├── db/
│   ├── migrations/
│   │   └── 20250727060514_create_review_table.sql
│   └── repositories/
│       ├── review.go          # Review data access
│       └── storage.go         # Database connection
├── dto/
│   └── review.go              # Data transfer objects
├── middlewares/
│   └── validator.go           # Request validation middleware
├── models/
│   └── review.go              # Review data model
├── router/
│   ├── review_router.go       # Review routes
│   └── router.go              # Main router setup
├── services/
│   └── review_service.go      # Business logic
├── utils/
│   └── json.go                # JSON response utilities
├── go.mod                     # Go module file
├── go.sum                     # Go module checksums
├── main.go                    # Application entry point
└── Makefile                   # Build and migration commands
```

## 🧪 Error Handling

The service implements comprehensive error handling:

- **HTTP Status Codes**: Appropriate status codes for different scenarios
- **JSON Error Responses**: Structured error messages
- **Validation Errors**: Input validation with detailed error messages
- **Database Errors**: Graceful handling of database operations
- **External Service Errors**: Timeout and connection error handling

## 🔧 Configuration

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `PORT` | Server port | `:3004` | No |
| `DB_HOST` | Database host | - | Yes |
| `DB_PORT` | Database port | - | Yes |
| `DB_USER` | Database user | - | Yes |
| `DB_PASSWORD` | Database password | - | Yes |
| `DB_NAME` | Database name | - | Yes |
| `BOOKING_SERVICE_URL` | Booking service URL | `http://localhost:3000` | No |

## 🚀 Deployment

### Docker (Recommended)

```dockerfile
FROM golang:1.24.4-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 3004
CMD ["./main"]
```

### Local Development

```bash
# Run with hot reload (requires air)
air

# Run tests
go test ./...

# Build binary
go build -o review-service main.go
```

## 📝 Development

### Adding New Features

1. **Create Migration**: `make migrate-create name="add_new_feature"`
2. **Update Model**: Add fields to `models/review.go`
3. **Update DTO**: Add fields to `dto/review.go`
4. **Update Repository**: Add methods to `db/repositories/review.go`
5. **Update Service**: Add business logic to `services/review_service.go`
6. **Update Controller**: Add handlers to `controllers/review_controller.go`
7. **Update Router**: Add routes to `router/review_router.go`

### Code Style

- Follow Go conventions and best practices
- Use meaningful variable and function names
- Add comments for complex logic
- Implement proper error handling
- Write unit tests for new features