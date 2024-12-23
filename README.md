# managing_user_records_app

Managing User Records App

# README: User Record App

## Overview

This project is a **User Record Management System** built using **Golang**, **Echo framework**, and **Next.js**. The backend is designed to handle RESTful APIs, while the frontend provides a user-friendly interface for managing user records.

## Features

- **Backend**:

  - Developed in **Golang** with the **Echo** framework.
  - RESTful API for user CRUD operations (Create, Read, Update, Delete).
  - Input validation and error handling.
  - Database integration with support for relational databases (e.g., PostgreSQL) and non-relational databases (e.g., MongoDB).
  - Follows clean architecture principles.

- **Frontend**:

  - Built with **Next.js**.
  - Modern UI for managing user records.
  - Server-side rendering (SSR).
  - API integration with the backend for seamless communication.

- **Additional Features**:
  - Environment-based configuration.
  - User authentication.
  - Asncynchronous communication systel for logging mechanism.

---

## Prerequisites

Ensure the following are installed on your machine:

- **Golang** (v1.21+)
- **Node.js** (v18.x+)
- **PostgreSQL**
- **MongoDB**
- **Git**
- **npm or bun or pnpm**

---

## Getting Started

### Backend Setup

1. **Clone the repository:**

   ```bash
   git clone https://github.com/kgminhtet-dev/managing_user_records_app.git
   cd managing_user_records_app/backend
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
   ```

3. **Configure the environment:**
   Create a `.env` file in the `backend` directory and define the following:

   ```env
   ENV=development
   HOST=localhost
   PORT=8080
   JWT_SECRET_TOKEN=secret
   USER_CONFIG_PATH=~/config/config.yaml (must include file name)
   RECORD_CONFIG_PATH=~/config (just path)
   ```

   create a config.yaml file for user

   ```yaml
   development:
   database:
       name: postgres
       host: localhost
       port: 5432
       user: postgres
       password: 12345678
       dbname: mur_user
       sslmode: disable
       timezone: Asia/Yangon

   testing:
   database:
       name: sqlite
   ```

   create a config.yaml file for record

   ```yaml
   development:
   database:
       url: mongodb://localhost:27017
       name: mur
       collection: records

   testing:
   database:
       url: mongodb://localhost:27017
       name: test
   ```

4. **Start the backend server:**

   ```bash
   make all or
   make build && make run
   ```

   The backend will be available at `http://localhost:8080`.

---

### Frontend Setup

1. **Navigate to the frontend directory:**

   ```bash
   cd ../frontend
   ```

2. **Install dependencies:**

   ```bash
   npm install or
   bun install or
   pnpm install
   ```

3. **Configure the environment:**
   Create a `.env` file in the `frontend` directory and define the following:

   ```env
    URI=http://localhost:8080
   ```

4. **Start the development server:**

   ```bash
   npm run dev or
   bun dev or
   pnpm dev
   ```

   The frontend will be available at `http://localhost:3000`.

---

## API Endpoints

### Base URL: `/api/v1/users`

| Method | Endpoint | Description               |
| ------ | -------- | ------------------------- |
| GET    | `/`      | Get all users             |
| GET    | `/:id`   | Get user by ID            |
| POST   | `/`      | Create a new user         |
| PUT    | `/:id`   | Update user details by ID |
| DELETE | `/:id`   | Delete user by ID         |

---

## Tech Stack

### Backend:

- **Language**: Golang
- **Framework**: Echo
- **Database**: PostgreSQL/MySQL
- **Others**: GORM, Echo middleware

### Frontend:

- **Framework**: Next.js
- **Styling**: ShacdCN/UI, Tailwind CSS

---

Feel free to suggest improvements or report issues in the [Issues](https://github.com/your-username/user-record-app/issues) section. Happy coding! 🎉
