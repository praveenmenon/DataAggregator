# Golang Data Aggregator

This project simulates a data aggregation system for sensor streams. Multiple sensor actors (like temperature and pressure) push data into a centralized channel. 
An aggregator collects this data and outputs it, while a supervisor listens for user input to gracefully shut the system down.

## 🧠 Project Overview

- **Actors**:
  - `temperature`: Simulates temperature readings
  - `pressure`: Simulates atmospheric pressure readings
- **Aggregator**:
  - Collects and logs all sensor data in real-time
- **Supervisor**:
  - Watches for user termination (via keyboard input) and cancels the entire process

## 🚀 How It Works

1. `main.go` starts all goroutines: pressure, temperature, aggregator, and supervisor.
2. All sensors write to a shared channel.
3. Aggregator reads from the channel and prints sensor data.
4. Supervisor waits for any keyboard input (like hitting `Enter`) and stops the program.

## ✅ Features

- Concurrency using goroutines
- Graceful shutdown with `context.Context`
- Clean separation of responsibilities
- Extendable design (easy to add more sensors)

## 🛠️ Run Locally

```bash
go run ./cmd/main.go


