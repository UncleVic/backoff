# Go Backoff Utility

A lightweight, configurable Go package for implementing **Exponential Backoff with Jitter**. This utility is designed to help distributed systems handle retries gracefully, preventing "thundering herd" issues through randomized delay scaling.

---

## Features

* **Exponential Scaling**: Increases delay between retries based on a configurable multiplier (`Factor`).
* **Normalized Jitter**: Uses `math/rand/v2` with a normal distribution to ensure retry attempts across multiple clients are spread out.
* **Safety Bounds**: Provides strict enforcement of `MaxDelay` and `MinDelay`.
* **Smart Defaults**: Automatically applies sensible production-ready defaults if configuration fields are omitted.

---

## Installation

```bash
go get github.com/unclevic/backoff
```

## Usage

### Basic Example

```go
package main

import (
    "fmt"
    "time"
    "github.com/unclevic/backoff"
)

func main() {
    cfg := backoff.Config{
        MaxAttempts:    5,
        MinDelay:       1 * time.Second,
        MaxDelay:       30 * time.Second,
        JitterFraction: 0.1,
        Factor:         2.0,
    }

    bo := backoff.NewBackoff(cfg)

    for {
        delay, err := bo.NextDelay()
        if err != nil {
            fmt.Println("Retries exhausted:", err)
            break
        }

        fmt.Printf("Attempting retry in %v...\n", delay)
        time.Sleep(delay)

        // Your logic (API call, DB connection, etc.) goes here
    }
}
```
