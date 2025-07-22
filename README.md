# CPU Intensive Task

A simple Go program that simulates a moderately CPU-intensive workload with periodic sleep intervals.

## Overview

This program runs a CPU-intensive task for 5 minutes, alternating between:
- **400ms of busy work**: Performing continuous mathematical calculations (square root operations)
- **600ms of sleep**: Allowing the CPU to rest

The program is designed to demonstrate controlled CPU usage patterns and graceful shutdown handling.

## Features

- **Controlled CPU Usage**: Alternates between CPU-intensive work and idle periods
- **Graceful Shutdown**: Handles SIGINT (Ctrl+C) and SIGTERM signals
- **Progress Tracking**: Displays iteration progress with elapsed time
- **Single-threaded**: Uses `GOMAXPROCS(1)` to limit execution to one CPU core

## Usage

### Running the Program

```bash
go run main.go
```

### Building and Running

```bash
go build -o cpu-task
./cpu-task
```

### Stopping the Program

Press `Ctrl+C` to gracefully stop the program. It will complete the current iteration and then exit.

## Output

The program outputs progress information for each iteration:

```
Starting moderately CPU-intensive task for 5 minutes...
Iteration 1 complete (total elapsed: 1s)
Iteration 2 complete (total elapsed: 2s)
...
^C
Interrupt received. Shutting down gracefully...
CPU task stopped.
Done.
```

## Requirements

- Go 1.18 or later (uses `atomic.Bool`)

## Use Cases

- Testing CPU performance and thermal behavior
- Simulating moderate workloads for benchmarking
- Demonstrating graceful shutdown patterns in Go
- Educational purposes for understanding CPU usage patterns
