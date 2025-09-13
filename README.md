# Tourism-FullStack-admin

<p align="center">
  <img src="https://via.placeholder.com/200x200?text=Tourism+Admin" alt="Tourism-FullStack-admin Logo" width="200" height="200">
</p>

<p align="center">
  A comprehensive full-stack management system for tourism services built with modern technology stack.
</p>

## 🌟 Overview

Tourism-FullStack-admin is a powerful, feature-rich management system designed specifically for tourism services. It provides a complete solution for managing trips, accommodations, bookings, user accounts, and payments with a focus on security, performance, and user experience.

## ✨ Features

### Core Features

- **Complete RBAC Permission Control System**
  - Role-based access control with fine-grained permissions
  - Dynamic role assignment and permission management
  - Multi-level administrative capabilities

- **Authentication System**
  - JWT-based authentication with role separation
  - Mobile authentication with OTP verification
  - Secure session management

- **Payment Processing**
  - Multi-stage payment processing (Visa, PayPal, Cash)
  - Secure card data handling (no CVV storage)
  - Payment status tracking and reporting

- **Data Security**
  - Comprehensive data validation
  - Secure file upload and storage system
  - Encrypted sensitive information

- **Developer Tools**
  - Code auto-generation functionality
  - Rich middleware support
  - Plugin architecture design
  - Swagger API documentation

### Tourism-Specific Features

- **Booking Management**
  - Trip and accommodation booking system
  - Hotel and flight reservation integration
  - Booking confirmation and e-tickets
  - Travel itinerary management

- **Content Management**
  - Interactive property/destination galleries
  - Pros and cons comparison for destinations
  - User reviews and ratings
  - Seasonal pricing and availability

## 🛠️ Technology Stack

### Frontend

- **Vue 3.5.7 + Composition API**
- **Vite 6.2.3** - Build tool
- **Pinia 2.2.2** - State management
- **Element Plus 2.10.2** - UI component library
- **UnoCSS 66.4.2** - Atomic CSS framework
- **Vue Router 4.4.3** - Routing management
- **Axios 1.8.2** - HTTP client
- **ECharts 5.5.1** - Data visualization
- **@vueuse/core** - Vue composition API toolkit

### Backend

- **Go 1.23 + Gin 1.10.0** - Web framework
- **GORM 1.25.12** - ORM framework
- **Casbin 2.103.0** - Permission management
- **Viper 1.19.0** - Configuration management
- **Zap 1.27.0** - Logging system
- **Redis 9.7.0** - Caching
- **JWT 5.2.2** - Authentication and authorization
- **PostgreSQL** - Primary database
- **MongoDB** - File storage in development
- **AWS S3** - Production file storage
- **libphonenumber** - Phone number validation

## 📂 Project Structure

### Overall Architecture

Tourism-FullStack-admin adopts a front-end and back-end separation architecture with a monolithic backend:

- **Backend (server/)**: RESTful API based on Go + Gin with a layered architecture
- **Frontend (web/)**: Single-page application based on Vue 3 + Vite
- **Mobile (mobile/)**: Mobile application for user authentication and service access
- **Deployment (deploy/)**: Docker, Kubernetes, and other deployment configurations

For detailed directory structure, please refer to the project documentation.

## 🚀 Getting Started

### Prerequisites

- Go 1.23 or higher
- Node.js 18.x or higher
- npm 9.x or higher
- PostgreSQL 14.x or higher
- Redis 7.x or higher

### Installation

1. Clone the repository

```bash
git clone https://github.com/yourusername/tourism-fullstack-admin.git
cd tourism-fullstack-admin
```

2. Install backend dependencies

```bash
cd server
go mod tidy
```

3. Install frontend dependencies

```bash
cd web
npm install
```

4. Configure the environment

```bash
cp server/config.yaml.example server/config.yaml
# Edit the config.yaml file with your database credentials and other settings
```

5. Initialize the database

```bash
cd server
go run cmd/main.go init_db
```

6. Start the development servers

```bash
# Start the backend server
cd server
go run main.go

# In another terminal, start the frontend server
cd web
npm run dev
```

### Using Docker

You can also use Docker to run the entire application:

```bash
docker-compose -f deploy/docker-compose/docker-compose.yml up
```

## 📚 API Documentation

API documentation is available via Swagger UI at `http://localhost:8080/swagger/index.html` when the server is running.

## 🧪 Testing

Run the tests with:

```bash
cd server
go test ./... -v


make docker-run
docker-compose -f deploy/docker-compose/docker-compose.yml build web
docker-compose -f deploy/docker-compose/docker-compose.yml up -d
```

## 🔒 Security

Please see our [Security Policy](./SECURITY.md) for reporting security vulnerabilities.

## 🤝 Contributing

Contributions are welcome! Please read our [Contributing Guidelines](./CONTRIBUTING.md) before submitting a pull request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

## 📞 Contact

If you have any questions or suggestions, please contact the project maintainer:

- **Name:** Khaled Mohamed
- **Email:** [mo7amed4522@gmail.com](mailto:mo7amed4522@gmail.com)
- **Phone:** +971 052 537 4187 (UAE)
- **WhatsApp:** [+971 052 537 4187](https://wa.me/971525374187)
