# Bunker

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.19-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

**Bunker** is a lightweight, modular Go utility toolkit designed to streamline common backend tasks, including working-day/SLA calculations, structured HTTP application errors, a fluent chainable HTTP client, and core utilities.

---

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Packages & Usage](#packages--usage)
  - [1. Chronos (Business Day & SLA Calculator)](#1-chronos)
  - [2. Errors (Structured HTTP Errors)](#2-errors)
  - [3. GoRest (Fluent HTTP Client)](#3-gorest)
  - [4. Utils (Core Helpers & Logging)](#4-utils)
- [Testing](#testing)
- [License](#license)

---

## Features

- **🕒 Chronos**: Calculate target dates by adding business days (weekdays) while automatically skipping weekends and custom public holidays.
- **🚨 Errors**: Predefined, standardized HTTP application errors with HTTP status codes, customizable error codes, and messages.
- **🌐 GoRest**: Fluent, chainable HTTP REST client with support for payloads (struct, map, string), query parameters, custom headers, bearer tokens, basic auth, cookie jar, and formatted debug logging.
- **🛠️ Utils**: Handy helper functions for string validations and color-coded console logs (`DEBUG-INFO`, `DEBUG-WARN`, `DEBUG-ERROR`, `DEBUG-FATAL`).

---

## Installation

```bash
go get github.com/olympus-forge/bunker
```

---

## Packages & Usage

### 1. Chronos

Calculate future business dates by skipping weekends (Saturday and Sunday) and any user-specified holidays:

```go
package main

import (
	"fmt"
	"time"

	"github.com/olympus-forge/bunker/chronos"
)

func main() {
	today := time.Now()

	// List of public or company holidays to skip
	holidays := []time.Time{
		time.Date(2026, time.December, 25, 0, 0, 0, 0, time.UTC),
	}

	// Calculate date after 5 working days
	targetDate := chronos.AddWeekDay(5, &today, holidays)
	fmt.Println("SLA Due Date:", targetDate.Format("2006-01-02"))

	// Or use fluent chaining
	dueDate := chronos.New().SetToday(&today).SetHolidays(holidays).AddWeekDay(5)
	fmt.Println("Fluent Due Date:", dueDate.Format("2006-01-02"))
}
```

---

### 2. Errors

Standardize API error handling with mapped HTTP status codes and custom application error codes:

```go
package main

import (
	"fmt"

	"github.com/olympus-forge/bunker/errors"
)

func main() {
	// Create from predefined HTTP error constants
	err := errors.New(errors.StatusBadRequest)

	// Customize error code and message
	err.SetErrorCode("AUTH-4001").SetMessage("Invalid user credentials provided")

	fmt.Printf("HTTP Status: %d\n", err.HttpStatusCode) // 400
	fmt.Printf("App Code: %s\n", err.ErrorCode)         // AUTH-4001
	fmt.Printf("Message: %s\n", err.Error())            // Invalid user credentials provided
}
```

Pre-defined error constants include `BadRequest`, `UnAuthorized`, `Forbidden`, `NotFound`, `Conflict`, `UnprocessableEntity`, `TooManyRequests`, `InternalServerError`, `BadGateway`, `ServiceUnavailable`, and more.

---

### 3. GoRest

Perform HTTP requests cleanly using method chaining:

```go
package main

import (
	"fmt"
	"time"

	"github.com/olympus-forge/bunker/gorest"
)

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	client := gorest.New("https://api.example.com")

	// GET Request with Query Parameters & Custom Headers
	resp := client.Get().
		AddPath("/v1/users").
		Query(map[string]string{
			"page":  "1",
			"limit": "20",
		}).
		SetHeader("Accept", "application/json").
		SetTimeout(10 * time.Second).
		SetDebug(true). // Enables colored debug output
		Do()

	if resp.HaveError() {
		fmt.Printf("Request failed: %v\n", resp.Errors)
		return
	}

	fmt.Printf("Status Code: %d\n", resp.Response.StatusCode)

	// POST Request with Payload & Bearer Token
	payload := CreateUserRequest{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	postResp := gorest.New("https://api.example.com").
		Post().
		AddPath("/v1/users").
		SetToken("your-jwt-token").
		SetPayload(&payload).
		Do()

	if postResp.HaveError() {
		fmt.Printf("Post failed: %v\n", postResp.Errors)
		return
	}
}
```

#### GoRest Key Capabilities:
- **HTTP Methods**: `Get()`, `Post()`, `Put()`, `Patch()`, `Delete()`, `Head()`, `Options()`
- **Authentication**: `SetToken("jwt")`, `SetBasicAuth("user", "pass")`
- **Payloads**: Automatic JSON serialization for structs, maps, or raw strings
- **Debug Mode**: `SetDebug(true)` prints a formatted console log showing the request details and response latency
- **Cookie Jar**: Built-in cookie jar support using `publicsuffix`

---

### 4. Utils

Quick string validation and ANSI-colored terminal loggers:

```go
package main

import (
	"github.com/olympus-forge/bunker"
)

func main() {
	// String check
	isEmpty := bunker.IsEmptyString("   ") // true

	// Colored logs
	bunker.LogInfo("Service started on port :8080")
	bunker.LogWarn("Rate limit threshold reached 80%")
	bunker.PrintErr("Failed to connect to cache server")
	// bunker.FatalPrintErr("Unrecoverable error") // Prints red log & calls os.Exit(1)
}
```

---

## Testing

Run tests across all packages:

```bash
go test -v ./...
```

---

## License

This project is licensed under the [MIT License](LICENSE).
