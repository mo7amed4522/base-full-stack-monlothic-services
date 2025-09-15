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

### to develop in this tamplate

add model for the application

```bash
cd server/model/app_name
# this use for user side not system side
mkdir app_name.go
cd app_name
mkdire request, response

```

fist add app_jwt.go inside folder utils

```bash
cd server/utils/app_jwt.go
```

```go
type AppJWT struct {
   SigninKey []byte
}

var (
   AppTokenExpired     =errors.New("AppToken is expired")
   AppTokenNotVaildYet = errors.New("AppToken not active yet")
   AppTokenMalformed   = errors.New("That's not even a AppToken")
   AppTokenInvalid     =errors.New("Couldn't handel this AppToken:")
)

func NewAppJWT() *AppJWT{
    return &AppJWT{
        []byte(global.GVA_CONFIG.JWT.SigningKey),
    }
}

func (j *AppJWT) CreateClaims(baseClaims appapiReq.BaseClaims) appapiReq.CustomClaims {
  expiresDuration , _ := ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
  claims := appapiReq.CustomClaims {
      BaseClaims: baseClaims,
      BufferTime: 60 * 60 * 24 * 7,  // 7-day buffer time
      RegisteredClaims: jwt.RegisteredClaims {
          Audience:   jwt.ClaimStrings{"AppApi"}, // Audience
          NotBeforce: jwt.NewNumberDate(time.Now().Add(-1000)), // Signature effective time
          ExpireAt:   jwt.NewNumberDate(time.Now().Add(expireDuration)), // Expiration time
          Issure:     global.GVA_CONFIG.JWT.Issuer, //Issuer of the signature
      },
  }
  return claims
}
  // Create a token
  func (j *AppJWT) CreateToken(claims appapiReq.CustomClaims)(string, error){
    token := jwt.NewWithClaims(jwt.SigninMethodHS256, claims)
    return token.SignedString(j.SigningKey)
  }

  // CreateTokenByOldToken replaces the old token with a new one. Use merge back to source to avoid concurrency issues.

  func (j *AppJWT) CreateTokenByOldToken(oldToken string, claims appapiReq.CustomClaims)(string, error){
    v, err, _ := global.GVA_Concurrency_Control.Do("AppJWT:"+oldToken, func() (interface{}, error) {
        return j.CreateToken(claims)
    })
    return v.(string), err
  }

  // Parsing the token
  func (j *AppJWT) ParseToken(tokenString string)(*appapiReq.CustomClaims, error){
    token, err := jwt.ParseWithClaims(tokenString, &appapiReq.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error){
      if err != nil {
        if ve, ok := err.(*jwt.ValidationError); ok {
          if ve.Errors&jwt.ValidationErrorMalformed != 0 {
              return nil, AppTokenMalformed
          }else if ve.Errors&jwt.ValidationErrorExpired != 0 {
              // token is Expired
              return nil, AppTokenExpired
          }else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
              return nil, AppTokenNotVaildYet
          }else {
              return nil, AppTokenInvaild
          }
        }
      }
      if token != nil {
        if claims , ok := token.Claims.(*appapiReq.CustomClaims); ok && token.Valid {
            return claims, nil
        }
        return nil , AppTokenInvalid
      }else{
        return nil, AppTokenInvalid
      }
    }
```

sec add inside folder middleware the JWT for application userSide

```bash
# this use for create the JWT for application Side
cd server/middleware/app_jwt.go
```

```go

func AppJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get a token
		token := utils.GetToken(c)
		if token == "" {
			response.NoAuth("Not logged in or illegal access", c)
			c.Abort()
			return
		}

		// Parsing tokens using AppJWT
		j := utils.NewAppJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.AppTokenExpired) {
				response.NoAuth("Authorization has expired", c)
				utils.ClearToken(c)
				c.Abort()
				return
			}
			response.NoAuth(err.Error(), c)
			utils.ClearToken(c)
			c.Abort()
			return
		}

		// Set the claims to the context using "appClaims" as the key
		c.Set("appClaims", claims)

		// Directly extract the user ID and set it to the context to avoid repeated parsing
		c.Set("appUserID", claims.BaseClaims.ID)

		// Check if the token is about to expire, and refresh it if so
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			utils.SetToken(c, newToken, int(dr.Seconds()))
		}
		c.Next()

		// Set the new token to the response header
		if newToken, exists := c.Get("new-token"); exists {
			c.Header("new-token", newToken.(string))
		}
		if newExpiresAt, exists := c.Get("new-expires-at"); exists {
			c.Header("new-expires-at", newExpiresAt.(string))
		}
	}
}

```

> then after create JWT and creat the JWT go to create a new Thing about
> **Servics** > **api** > **routes**

```bash
cd services
mkdir app_api
# this folder must create a file enter.go
 # in this must  struct the all services inside this project
```

```go
type ServicesGroup struct {
    // here add all services that you created
    // like we will create a UserServices
    UserServices
}
```

#### UserServices

```go
type UserServices struct{}

func(appUserServices *UserServices) SignInUser(u *appapi.AppUser)(userInter *appapi.AppUser, err error){
  // here we check on DB that db has init or not
  if nil == global.GVA_DB {
    return nil, fmt.Errorf("db not init ")
  }
  // here get the data from DB
  var user appapi.AppUser
  err = global.GVA_DB.Where("name = ?", u.UserName).First(&user).Error
  // here check the password is equial that exist
  if err == nil {
      if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
        return nil ,errors.New("Incorrect password")
      }
  }
  return &user, err
}
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

# this for run the docker throw Makefile
make docker-run
# this use for build the web application without server
docker-compose -f deploy/docker-compose/docker-compose.yml build web
# this use for build the server application without web application
docker-compose -f deploy/docker-compose/docker-compose.yml build server
# this for build and push the application to both web and server applications
docker-compose -f deploy/docker-compose/docker-compose.yml up -d
# this for delete the docker container
docker-compose -f deploy/docker-compose/docker-compose.yml down
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
