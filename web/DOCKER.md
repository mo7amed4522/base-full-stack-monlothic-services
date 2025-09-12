# Docker Setup for Landing Page Management

This document provides instructions for running the Landing Page Management application using Docker and you have 3 of producation and dev and nginx.

## Prerequisites

- Docker Engine 20.10+
- Docker Compose 2.0+
- At least 2GB of available RAM
- At least 1GB of available disk space

## Quick Start

### Development Environment

```bash
# Start development environment
npm run docker:dev

# Stop development environment
npm run docker:dev:down
```

### Production Environment

```bash
# Start production environment
npm run docker:prod

# Stop production environment
npm run docker:prod:down
```

## Docker Commands

### Building Images

```bash
# Build development image
npm run docker:build:dev

# Build production image
npm run docker:build:prod

# Build both images
npm run docker:build
```

### Running Containers

```bash
# Run development container
npm run docker:run:dev

# Run production container
npm run docker:run:prod
```

### Maintenance

```bash
# Clean up Docker system
npm run docker:clean
```

## Docker Compose Files

### docker-compose.yml
Main compose file with profiles for both development and production environments.

### docker-compose.dev.yml
Dedicated development environment with:
- Hot reload enabled
- Volume mounting for live code changes
- Development server on port 5180

### docker-compose.prod.yml
Production environment with:
- Nginx serving static files
- Optimized build
- Health checks
- Production server on port 80

## Environment Variables

### Development
- `NODE_ENV=development`
- `VITE_API_BASE_URL=http://172.232.96.159:13340/api/`

### Production
- `NODE_ENV=production`

## Ports

- **Development**: 5180
- **Production**: 80

## Volumes

### Development
- Source code is mounted as a volume for hot reload
- Node modules are preserved in a separate volume

### Production
- No volumes mounted (static files served from container)

## Multi-stage Build

The Dockerfile uses multi-stage builds:

1. **Builder stage**: Installs dependencies and builds the application
2. **Development stage**: Sets up development environment with hot reload
3. **Production stage**: Serves built files with Nginx

## Nginx Configuration

The production environment uses a custom Nginx configuration that includes:
- Gzip compression
- Client-side routing support
- Static asset caching
- Security headers
- API proxy configuration (if backend is available)

## Troubleshooting

### Common Issues

1. **Port already in use**
   ```bash
   # Check what's using the port
   lsof -i :5180  # for development
   lsof -i :80    # for production
   ```

2. **Permission denied**
   ```bash
   # Make sure Docker daemon is running
   sudo systemctl start docker  # Linux
   # or start Docker Desktop on macOS/Windows
   ```

3. **Out of disk space**
   ```bash
   # Clean up Docker system
   npm run docker:clean
   
   # Remove unused images
   docker image prune -a
   ```

4. **Build fails**
   ```bash
   # Clear Docker build cache
   docker builder prune
   
   # Rebuild without cache
   docker-compose build --no-cache
   ```

### Logs

```bash
# View development logs
docker-compose -f docker-compose.dev.yml logs -f

# View production logs
docker-compose -f docker-compose.prod.yml logs -f

# View specific service logs
docker logs landingpage-management-dev
docker logs landingpage-management-prod
```

### Accessing the Container

```bash
# Access development container
docker exec -it landingpage-management-dev sh

# Access production container
docker exec -it landingpage-management-prod sh
```

## Performance Optimization

### Development
- Uses volume mounting for instant code changes
- Preserves node_modules in separate volume
- Enables hot module replacement

### Production
- Multi-stage build reduces image size
- Nginx serves static files efficiently
- Gzip compression enabled
- Static asset caching configured

## Security Considerations

- Production image runs as non-root user
- Security headers configured in Nginx
- No sensitive data in Docker images
- Environment variables for configuration

## Deployment

### Local Development
```bash
npm run docker:dev
```
Access at: http://localhost:5180

### Production Deployment
```bash
npm run docker:prod
```
Access at: http://localhost

### Custom Deployment
For custom deployments, modify the docker-compose files or create new ones based on your infrastructure requirements.

## Support

For issues related to Docker setup, please check:
1. Docker and Docker Compose versions
2. Available system resources
3. Network connectivity
4. Log files for specific error messages