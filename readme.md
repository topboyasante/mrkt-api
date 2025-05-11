# MRKT

MRKT is a marketplace for people to share products they want to sell. Potential buyers can reach out to you via `call` or `whatsapp` and discussions can be taken from there. Open Source, Open Code.

### Stack
  - Frontend: React, NextJS, TypeScript
  - Backend: Golang, Gin, GORM
  - Database: Supabase
  - Hosting: Fly.io

### Features

- [x] Post Listings
- [x] Contact listing owners by Phone
- [x] Contact listing owners by WhatsApp
- [x] View All Listings
- [x] View Featured Listings
- [x] Edit Listings
- [x] Delete Listings
- [ ] Make Listings Featured
- [ ] Add Categories to Listings

### How to set up
1. Clone the repository
2. Fill in your environment variables
3. Spin up a PostgreSQL DB by removing the comments from the `docker-compose.yml` file
4. Run app with `docker compose up --build`
