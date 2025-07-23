# Airbnb Clone - Microservices Architecture

A comprehensive Airbnb Dev built using microservices architecture with modern technologies. This project demonstrates scalable, distributed system design with separate services for authentication, booking management, hotel management, and notifications.

## 🏗️ Architecture Overview

This project follows a microservices architecture pattern with the following services:

- **AuthInGo** - Authentication & User Management Service (Go)
- **BookingService** - Booking Management Service (Node.js/TypeScript)
- **HotelService** - Hotel & Room Management Service (Node.js/TypeScript)
- **NotificationService** - Email Notification Service (Node.js/TypeScript)

## 🚀 Services Overview

### 1. AuthInGo (Authentication Service)

**Language**: Go  
**Purpose**: Handles user authentication, registration, login, JWT-protected profile, request validation, rate limiting, and proxies requests to hotel/booking services.

#### Features Implemented:
- ✅ User Registration
- ✅ User Login
- ✅ JWT Token Authentication
- ✅ Protected Profile Endpoint (JWT)
- ✅ Request Validation Middleware (signup/login)
- ✅ Rate Limiting Middleware
- ✅ Database Connection Management
- ✅ RESTful API Endpoints
- ✅ Environment Configuration
- ✅ Health Check Endpoint
- ✅ Proxy to Hotel & Booking Services
- ✅ Consistent JSON Error/Success Responses
- ✅ Improved Error Handling

#### Technologies & Tools:
- **Framework**: Go with Chi Router (`github.com/go-chi/chi/v5`)
- **Database**: MySQL with MyMySQL driver (`github.com/ziutek/mymysql`)
- **Authentication**: JWT (`github.com/golang-jwt/jwt/v4`)
- **Security**: bcrypt for password hashing (`golang.org/x/crypto`)
- **Configuration**: godotenv (`github.com/joho/godotenv`)
- **Build Tool**: Makefile
- **Go Version**: 1.24.4

#### API Endpoints:
```
POST /signup     - User registration (validated)
POST /login      - User login (validated)
GET  /profile    - Get user profile (JWT protected)
GET  /ping       - Health check

# Proxy routes (internal)
/hotelservice/*      - Proxies to HotelService
/bookingservice/*    - Proxies to BookingService
```

#### Project Structure:
```
AuthInGo/
├── app/           # Application configuration
├── config/        # Database and environment config
├── controllers/   # Request handlers
├── db/            # Database connection
├── dto/           # Data transfer objects
├── errors/        # Error handling
├── middlewares/   # Auth, validation, rate limiting
├── models/        # Data models
├── router/        # Route definitions
├── services/      # Business logic
├── utils/         # Utility functions (JSON, proxy, auth)
└── main.go        # Application entry point
```

---

### 2. BookingService (Booking Management)

**Language**: Node.js/TypeScript  
**Purpose**: Manages hotel bookings with advanced features like idempotency, distributed locking, and asynchronous processing

#### Features Implemented:
- ✅ Create Booking with Distributed Locking
- ✅ Confirm Booking with Idempotency Keys
- ✅ Cancel Booking
- ✅ Booking Status Management (PENDING, CONFIRMED, CANCELLED)
- ✅ Idempotency Key Generation & Validation
- ✅ Distributed Locking with Redis (Redlock)
- ✅ Asynchronous Email Queue Processing
- ✅ Request Validation with Zod
- ✅ Correlation ID Tracking
- ✅ Comprehensive Error Handling
- ✅ Database Transactions
- ✅ Logging with Winston

#### Technologies & Tools:
- **Framework**: Express.js (`^5.1.0`)
- **Language**: TypeScript
- **Database ORM**: Prisma (`^6.8.2`)
- **Database**: MySQL
- **Queue System**: BullMQ (`^5.53.1`)
- **Cache/Lock**: Redis with IORedis (`^5.6.1`)
- **Distributed Locking**: Redlock (`^5.0.0-beta.2`)
- **Validation**: Zod (`^3.24.2`)
- **Logging**: Winston (`^3.17.0`) with MongoDB and Daily Rotate File
- **UUID Generation**: UUID (`^11.1.0`)
- **Environment**: dotenv (`^16.5.0`)
- **Additional**: Mongoose (`^8.13.2`)

#### API Endpoints:
```
POST /api/v1/bookings/                    - Create new booking
POST /api/v1/bookings/confirm/:idempotencyKey - Confirm booking
GET  /api/v1/ping/                       - Health check
GET  /api/v1/ping/health                 - Health status
```

#### Key Features:
- **Idempotency**: Prevents duplicate bookings using UUID-based idempotency keys
- **Distributed Locking**: Uses Redlock algorithm to prevent race conditions
- **Queue Processing**: Asynchronous email notifications via BullMQ
- **Database Transactions**: Ensures data consistency
- **Request Correlation**: Tracks requests across services

#### Project Structure:
```
BookingService/
├── src/
│   ├── config/         # Configuration files
│   ├── controllers/    # Request handlers
│   ├── dto/           # Data transfer objects
│   ├── middlewares/   # Express middlewares
│   ├── prisma/        # Database schema & migrations
│   ├── producers/     # Queue producers
│   ├── queues/        # Queue definitions
│   ├── repositories/  # Data access layer
│   ├── routers/       # Route definitions
│   ├── services/      # Business logic
│   ├── utils/         # Utility functions
│   └── validators/    # Request validation
└── package.json
```

---

### 3. HotelService (Hotel Management)

**Language**: Node.js/TypeScript  
**Purpose**: Manages hotels, rooms, and room categories with full CRUD operations

#### Features Implemented:
- ✅ Hotel CRUD Operations (Create, Read, Update, Delete)
- ✅ Soft Delete Implementation
- ✅ Hotel Search & Filtering
- ✅ Room Management
- ✅ Room Category Management
- ✅ Hotel Rating System
- ✅ Database Migrations with Sequelize
- ✅ Repository Pattern Implementation
- ✅ Request Validation
- ✅ Comprehensive Error Handling
- ✅ Logging System
- ✅ Correlation ID Tracking

#### Technologies & Tools:
- **Framework**: Express.js (`^5.1.0`)
- **Language**: TypeScript
- **Database ORM**: Sequelize (`^6.37.7`)
- **Database**: MySQL (`mysql2 ^3.14.0`)
- **Migration Tool**: Sequelize CLI (`^6.6.2`)
- **Validation**: Zod (`^3.24.2`)
- **Logging**: Winston (`^3.17.0`) with MongoDB and Daily Rotate File
- **UUID Generation**: UUID (`^11.1.0`)
- **Environment**: dotenv (`^16.5.0`)
- **Additional**: Mongoose (`^8.13.2`)

#### API Endpoints:
```
POST   /api/v1/hotels/     - Create new hotel
GET    /api/v1/hotels/:id  - Get hotel by ID
GET    /api/v1/hotels/     - Get all hotels
PUT    /api/v1/hotels/:id  - Update hotel
DELETE /api/v1/hotels/:id  - Soft delete hotel
GET    /api/v1/ping/       - Health check
```

#### Database Models:
- **Hotel**: id, name, address, location, rating, ratingCount, timestamps, soft delete
- **Room**: id, hotelId, roomCategoryId, dateOfAvailability, price, bookingId
- **RoomCategory**: id, hotelId, category details

#### Key Features:
- **Repository Pattern**: Clean separation of data access logic
- **Soft Delete**: Hotels are marked as deleted rather than physically removed
- **Base Repository**: Generic repository class for common operations
- **Database Migrations**: Version-controlled database schema changes

#### Project Structure:
```
HotelService/
├── src/
│   ├── config/         # Configuration files
│   ├── controllers/    # Request handlers
│   ├── db/            # Database models & migrations
│   ├── dto/           # Data transfer objects
│   ├── middlewares/   # Express middlewares
│   ├── repositories/  # Data access layer
│   ├── routers/       # Route definitions
│   ├── services/      # Business logic
│   ├── utils/         # Utility functions
│   └── validators/    # Request validation
└── package.json
```

---

### 4. NotificationService (Email Notifications)

**Language**: Node.js/TypeScript  
**Purpose**: Handles email notifications with template-based messaging and queue processing

#### Features Implemented:
- ✅ Email Sending with Nodemailer
- ✅ Template-based Email System (Handlebars)
- ✅ Asynchronous Queue Processing with BullMQ
- ✅ Redis Queue Management
- ✅ Email Template Rendering
- ✅ Gmail SMTP Integration
- ✅ Queue Worker Setup
- ✅ Email Job Processing
- ✅ Error Handling & Logging
- ✅ Correlation ID Tracking
- ✅ Health Check Endpoints

#### Technologies & Tools:
- **Framework**: Express.js (`^5.1.0`)
- **Language**: TypeScript
- **Email Service**: Nodemailer (`^7.0.3`)
- **Template Engine**: Handlebars (`^4.7.8`)
- **Queue System**: BullMQ (`^5.53.1`)
- **Cache**: Redis with IORedis (`^5.6.1`)
- **Validation**: Zod (`^3.24.2`)
- **Logging**: Winston (`^3.17.0`) with MongoDB and Daily Rotate File
- **UUID Generation**: UUID (`^11.1.0`)
- **Environment**: dotenv (`^16.5.0`)
- **Additional**: Mongoose (`^8.13.2`)

#### API Endpoints:
```
GET /api/v1/ping/       - Health check
GET /api/v1/ping/health - Health status
```

#### Email Templates:
- **Welcome Template**: User welcome emails with personalization
- **Booking Confirmation**: Booking confirmation emails
- **Extensible**: Easy to add new templates

#### Key Features:
- **Queue-based Processing**: Asynchronous email sending via Redis queues
- **Template System**: Dynamic email content with Handlebars
- **Gmail Integration**: SMTP configuration for Gmail
- **Worker Pattern**: Dedicated workers for email processing
- **Error Handling**: Comprehensive error handling for email failures

#### Project Structure:
```
NotificationService/
├── src/
│   ├── config/         # Configuration files
│   ├── controllers/    # Request handlers
│   ├── dto/           # Data transfer objects
│   ├── middlewares/   # Express middlewares
│   ├── processors/    # Queue processors
│   ├── producers/     # Queue producers
│   ├── queues/        # Queue definitions
│   ├── routers/       # Route definitions
│   ├── services/      # Business logic
│   ├── templates/     # Email templates
│   ├── utils/         # Utility functions
│   └── validators/    # Request validation
└── package.json
```

---

## 🛠️ Common Technologies Across Services

### Shared Technologies:
- **Logging**: Winston with MongoDB and Daily Rotate File
- **Validation**: Zod for request validation
- **Environment Configuration**: dotenv
- **UUID Generation**: UUID library
- **Error Handling**: Custom error classes and middleware
- **Correlation Tracking**: Request correlation IDs
- **Health Checks**: Ping endpoints for service monitoring

### Infrastructure:
- **Database**: MySQL (primary database for all services)
- **Cache/Queue**: Redis (for caching, queuing, and distributed locking)
- **Email**: Gmail SMTP (for email notifications)

---

## 🚀 Getting Started

### Prerequisites
- Node.js (v18+)
- Go (v1.24.4+)
- MySQL
- Redis
- Gmail account (for email notifications)

### Environment Setup

Each service requires environment variables. Create `.env` files in each service directory:

#### AuthInGo (.env)
```env
DB_HOST=localhost
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=auth_db
JWT_SECRET=your_jwt_secret
```

#### BookingService (.env)
```env
PORT=3001
DATABASE_URL="mysql://user:password@localhost:3306/booking_db"
REDIS_SERVER_URL=redis://localhost:6379
LOCK_TTL=1000
```

#### HotelService (.env)
```env
PORT=3002
DB_HOST=localhost
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=hotel_db
MONGODB_URI=mongodb://localhost:27017/hotel_logs
```

#### NotificationService (.env)
```env
PORT=3003
MONGODB_URI=mongodb://localhost:27017/notification_logs
REDIS_HOST=localhost
REDIS_PORT=6379
MAIL_USER=your_gmail@gmail.com
MAIL_PASSWORD=your_app_password
```

### Installation & Running

#### AuthInGo
```bash
cd AuthInGo
go mod download
make run
```

#### BookingService
```bash
cd BookingService
npm install
npm run dev
```

#### HotelService
```bash
cd HotelService
npm install
npm run migrate  # Run database migrations
npm run dev
```

#### NotificationService
```bash
cd NotificationService
npm install
npm run dev
```

---

## 📊 Service Communication

The services communicate through:
- **HTTP APIs**: RESTful endpoints for synchronous communication
- **Message Queues**: BullMQ with Redis for asynchronous communication
- **Shared Database**: MySQL for data persistence
- **Correlation IDs**: For request tracking across services

---

## 🔧 Development Features

### Code Quality:
- **TypeScript**: Strong typing for Node.js services
- **Go**: Compiled language with strong typing
- **Validation**: Zod schemas for request validation
- **Error Handling**: Comprehensive error handling across all services
- **Logging**: Structured logging with Winston

### Database:
- **Migrations**: Database version control with Sequelize and Prisma
- **Transactions**: ACID compliance for critical operations
- **Soft Deletes**: Data preservation with soft delete patterns
- **Indexing**: Optimized database queries

### Performance:
- **Distributed Locking**: Prevents race conditions in booking
- **Queue Processing**: Asynchronous processing for better performance
- **Connection Pooling**: Efficient database connections
- **Caching**: Redis for caching and session management

---

## 🏆 Key Highlights

1. **Microservices Architecture**: Properly separated concerns with independent services
2. **Scalability**: Each service can be scaled independently
3. **Reliability**: Comprehensive error handling and logging
4. **Performance**: Asynchronous processing and distributed locking
5. **Maintainability**: Clean code architecture with repository patterns
6. **Monitoring**: Health checks and correlation tracking
7. **Security**: JWT authentication and input validation
8. **Data Consistency**: Database transactions and idempotency

---

## 📝 API Documentation

### 1. AuthInGo Service API

**Base URL:** `http://localhost:3000`

#### Authentication Endpoints

##### POST /users/register
Register a new user account.

**Request Body:**
```json
{
  "username": "string",
  "email": "string",
  "password": "string"
}
```

**Response (201 Created):**
```json
{
  "message": "User registered successfully",
  "userId": "integer"
}
```

**Response (400 Bad Request):**
```json
{
  "error": "Validation error message"
}
```

##### POST /users/login
Authenticate user and get JWT token.

**Request Body:**
```json
{
  "email": "string",
  "password": "string"
}
```

**Response (200 OK):**
```json
{
  "message": "Login successful",
  "token": "jwt_token_string",
  "user": {
    "id": "integer",
    "email": "string",
    "username": "string"
  }
}
```

**Response (401 Unauthorized):**
```json
{
  "error": "Invalid credentials"
}
```

##### GET /ping
Health check endpoint.

**Response (200 OK):**
```json
{
  "message": "pong",
  "timestamp": "2024-01-01T00:00:00Z"
}
```

---

### 2. BookingService API

**Base URL:** `http://localhost:3001`

#### Booking Endpoints

##### POST /api/v1/bookings/
Create a new booking with distributed locking.

**Request Body:**
```json
{
  "userId": 1,
  "hotelId": 1,
  "totalGuests": 2,
  "bookingAmount": 15000
}
```

**Response (201 Created):**
```json
{
  "bookingId": 1,
  "idempotencyKey": "550e8400-e29b-41d4-a716-446655440000"
}
```

**Response (400 Bad Request):**
```json
{
  "message": "Invalid request body",
  "success": false,
  "error": {
    "issues": [
      {
        "path": ["userId"],
        "message": "User ID must be present"
      }
    ]
  }
}
```

**Response (500 Internal Server Error):**
```json
{
  "message": "Failed to acquire lock for booking resource",
  "success": false
}
```

##### POST /api/v1/bookings/confirm/:idempotencyKey
Confirm a booking using idempotency key.

**Path Parameters:**
- `idempotencyKey`: UUID string from booking creation

**Response (200 OK):**
```json
{
  "bookingId": 1,
  "status": "CONFIRMED"
}
```

**Response (404 Not Found):**
```json
{
  "message": "Idempotency key not found",
  "success": false
}
```

**Response (400 Bad Request):**
```json
{
  "message": "Idempotency key already finalized",
  "success": false
}
```

#### Health Check Endpoints

##### GET /api/v1/ping/
Health check with request validation.

**Request Body:**
```json
{
  "message": "test"
}
```

**Response (200 OK):**
```json
{
  "message": "pong",
  "correlationId": "550e8400-e29b-41d4-a716-446655440000"
}
```

##### GET /api/v1/ping/health
Simple health status check.

**Response (200 OK):**
```json
"OK"
```

---

### 3. HotelService API

**Base URL:** `http://localhost:3002`

#### Hotel Management Endpoints

##### POST /api/v1/hotels/
Create a new hotel.

**Request Body:**
```json
{
  "name": "Grand Hotel",
  "address": "123 Main Street, City",
  "location": "Downtown",
  "rating": 4.5,
  "ratingCount": 150
}
```

**Response (201 Created):**
```json
{
  "message": "Hotel created successfully",
  "data": {
    "id": 1,
    "name": "Grand Hotel",
    "address": "123 Main Street, City",
    "location": "Downtown",
    "rating": 4.5,
    "ratingCount": 150,
    "createdAt": "2024-01-01T00:00:00Z",
    "updatedAt": "2024-01-01T00:00:00Z",
    "deletedAt": null
  },
  "success": true
}
```

##### GET /api/v1/hotels/:id
Get hotel details by ID.

**Path Parameters:**
- `id`: Hotel ID (integer)

**Response (200 OK):**
```json
{
  "message": "Hotel retrieved successfully",
  "data": {
    "id": 1,
    "name": "Grand Hotel",
    "address": "123 Main Street, City",
    "location": "Downtown",
    "rating": 4.5,
    "ratingCount": 150,
    "createdAt": "2024-01-01T00:00:00Z",
    "updatedAt": "2024-01-01T00:00:00Z",
    "deletedAt": null
  },
  "success": true
}
```

**Response (404 Not Found):**
```json
{
  "message": "Record with id 1 not found",
  "success": false
}
```

##### GET /api/v1/hotels/
Get all hotels (excluding soft-deleted).

**Response (200 OK):**
```json
{
  "message": "Hotels retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Grand Hotel",
      "address": "123 Main Street, City",
      "location": "Downtown",
      "rating": 4.5,
      "ratingCount": 150,
      "createdAt": "2024-01-01T00:00:00Z",
      "updatedAt": "2024-01-01T00:00:00Z",
      "deletedAt": null
    }
  ],
  "success": true
}
```

##### PUT /api/v1/hotels/:id
Update hotel information.

**Path Parameters:**
- `id`: Hotel ID (integer)

**Request Body:**
```json
{
  "name": "Updated Grand Hotel",
  "address": "456 New Street, City",
  "location": "Uptown",
  "rating": 4.8,
  "ratingCount": 200
}
```

**Response (200 OK):**
```json
{
  "message": "Hotel updated successfully",
  "data": {
    "id": 1,
    "name": "Updated Grand Hotel",
    "address": "456 New Street, City",
    "location": "Uptown",
    "rating": 4.8,
    "ratingCount": 200,
    "createdAt": "2024-01-01T00:00:00Z",
    "updatedAt": "2024-01-01T12:00:00Z",
    "deletedAt": null
  },
  "success": true
}
```

##### DELETE /api/v1/hotels/:id
Soft delete a hotel.

**Path Parameters:**
- `id`: Hotel ID (integer)

**Response (200 OK):**
```json
{
  "message": "Hotel deleted successfully",
  "success": true
}
```

**Response (404 Not Found):**
```json
{
  "message": "Hotel with id 1 not found",
  "success": false
}
```

#### Health Check Endpoints

##### GET /api/v1/ping/
Health check endpoint.

**Response (200 OK):**
```json
{
  "message": "pong",
  "correlationId": "550e8400-e29b-41d4-a716-446655440000"
}
```

---

### 4. NotificationService API

**Base URL:** `http://localhost:3003`

#### Health Check Endpoints

##### GET /api/v1/ping/
Health check with request validation.

**Request Body:**
```json
{
  "message": "test"
}
```

**Response (200 OK):**
```json
{
  "message": "pong",
  "correlationId": "550e8400-e29b-41d4-a716-446655440000"
}
```

##### GET /api/v1/ping/health
Simple health status check.

**Response (200 OK):**
```json
"OK"
```

#### Email Queue Processing

The NotificationService processes email jobs asynchronously through Redis queues. Email jobs are added to the queue by other services and processed by background workers.

**Email Job Payload:**
```json
{
  "to": "user@example.com",
  "subject": "Welcome to Airbnb Clone",
  "templateId": "Welcome",
  "params": {
    "name": "John Doe",
    "appName": "Airbnb Clone"
  }
}
```

**Available Email Templates:**
- `Welcome`: User welcome emails
- `BookingConfirmation`: Booking confirmation emails
- Custom templates can be added in `/src/templates/mailer/`

---

## 🔧 Common Response Patterns

### Success Response Format
```json
{
  "message": "Operation successful",
  "data": { /* response data */ },
  "success": true
}
```

### Error Response Format
```json
{
  "message": "Error description",
  "success": false,
  "error": { /* error details */ }
}
```

### HTTP Status Codes Used
- `200 OK`: Successful GET, PUT operations
- `201 Created`: Successful POST operations
- `400 Bad Request`: Invalid request data
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Access denied
- `404 Not Found`: Resource not found
- `409 Conflict`: Resource conflict
- `500 Internal Server Error`: Server error

### Request Headers
- `Content-Type: application/json`
- `Authorization: Bearer <jwt_token>` (for protected endpoints)
- `x-correlation-id`: Auto-generated correlation ID for request tracking

---

