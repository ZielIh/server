# HammerDown AntiCheat Server

## Description

HammerDown AntiCheat Server is the server component of the HammerDown cheat detection system.  The server receives data from clients, analyzes it using multiple detection methods, and responds with information about possible detected cheats.

## Features

**Hash-based Detection**: Uses CRC32 checksums for rapid identification of files known as cheats. 

**Signature-based Detection**: Analyzes specific byte patterns in received data to detect cheats based on signatures. 

**Data Compression**: Handles zlib-compressed data from clients. 

**MySQL Database**: Stores cheat information, hashes, and signatures. 

**Intelligent Caching**: Saves hashes from signature detections to accelerate future detections. 

## Requirements

- Go 1.19 or higher
- MySQL 5.7 or higher
- Configured environment variables

## Installation

1. Clone the repository
2. Install dependencies with `go mod tidy`
3. Configure environment variables
4. Run with `go run main.go`

## Configuration

The server requires these environment variables in a `.env` file: 

```
DB_USER=your_mysql_username
DB_PASS=your_mysql_password
DB_HOST=localhost:3306
DB_NAME=hammerdown_db
```

## API

**POST /detect**: Main endpoint for cheat detection. 

Accepts binary data compressed with zlib and returns JSON with detection results.
