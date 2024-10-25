# Project Name

Simple REST API using Go and Echo Framework

## Overview

This project is a RESTful API built with Go programming language and the Echo framework. It includes basic CRUD operations, middleware implementation, and database integration.

## Prerequisites

- Go version 1.19 or higher
- MySQL/PostgreSQL (depending on your database choice)
- Echo framework
- Air (for hot reload) [Optional]

## Tech Stack

- [Go](https://golang.org/)
- [Echo Framework](https://echo.labstack.com/)
- [GORM](https://gorm.io/)
- [JWT-Go](https://github.com/golang-jwt/jwt)
- [Godotenv](https://github.com/joho/godotenv)
- [Air](https://github.com/cosmtrek/air)

## Installation & Setup

1. Clone the repository

```bash
git clone https://github.com/zdnkarim/pteridophyte-go project-name
cd project-name
```

2. Install dependencies

```bash
go mod download
```

3. Create and configure .env file

```bash
cp .env.example .env
```

4. Run the application

```bash
# Normal run
go run main.go

# Using Air for hot reload
air
```
