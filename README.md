# Go URL Shortener Library

A lightweight, high-performance, and production-ready URL shortener library written in Go. This library uses **GORM** for persistent storage abstraction and **Redis** for blazing-fast caching mechanism ($O(1)$ lookups).

## Features

- **Database Agnostic**: Supports PostgreSQL, MySQL, SQLite, and more via GORM.
- **Built-in Cache**: Automatic caching using Redis to reduce database load.
- **Auto-Migration**: Automatically manages database schema updates for URL mappings.
- **Secure Code Generation**: Utilizes cryptographically secure random string generation (`crypto/rand`).

## Installation

```bash
go get [github.com/Firyal-dev/url-shorterner](https://github.com/Firyal-dev/url-shorterner)