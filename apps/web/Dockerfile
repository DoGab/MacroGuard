# =====================================
# VitalStack Web - Production Dockerfile
# =====================================
# Multi-stage build for SvelteKit with adapter-node

# ----- Build Stage -----
FROM node:25-alpine AS builder

WORKDIR /app

# Copy package files first for layer caching
COPY package*.json ./

# Install all dependencies (including devDependencies for build)
RUN npm ci

# Copy source code
COPY . .

# Build the application
RUN npm run build

# ----- Runtime Stage -----
FROM node:25-alpine

WORKDIR /app

# Create non-root user for security
RUN addgroup -g 1001 -S vitalstack && \
    adduser -u 1001 -S vitalstack -G vitalstack

# Copy built application from builder
COPY --from=builder /app/build ./build
COPY --from=builder /app/package*.json ./

# Install only production dependencies
RUN npm ci --omit=dev && npm cache clean --force

# Set ownership
RUN chown -R vitalstack:vitalstack /app

# Switch to non-root user
USER vitalstack

# Expose port
EXPOSE 3000

# Environment variables (can be overridden at runtime)
ENV NODE_ENV=production
ENV PORT=3000
ENV ORIGIN=http://localhost:3000
ENV PUBLIC_API_URL=""

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:3000 || exit 1

# Run the application
CMD ["node", "build"]
