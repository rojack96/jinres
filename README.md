# Jinres

**Jinres** is a small Go utility that provides a clean builder for structured JSON HTTP responses in **Gin**, plus a CLI tool that generates Swagger-friendly response structs for all HTTP status codes.

---

## 📦 Installation

```bash
go install github.com/rojack96/jinres@latest
```

This installs both:

* the Go package (`import "github.com/rojack96/jinres"`)
* the CLI command: `jinres`

---

## 🚀 Usage (Gin response builder)

```go
j := jinres.NewJinres()

j.OK().
    Message("success").
    Response(gin.H{"id": 1}).
    Custom("debug", true).
    Done(ctx)
```

Produces:

```json
{
  "status": 200,
  "message": "success",
  "response": { "id": 1 },
  "debug": true
}
```

### Builder features

* `Message("text")`
* `Response(any)`
* `Custom(key, value)`
* `Done(ctx)`

Includes helpers for all HTTP codes (`OK()`, `BadRequest()`, `NotFound()`, etc.).

---

## 🔧 CLI — Generate Swagger structs

Jinres includes a CLI command used to generate a Go file containing Swagger-compatible response structs such as:

```go
// Continue represents HTTP 100 responses.
type Continue struct {
    Status   uint16 `json:"status" example:"100"`
    Message  string `json:"message" example:"Continue"`
    Response any    `json:"response,omitempty"`
} // @name ContinueResponse
```

Run:

```bash
jinres init
```

This creates a file in your project root:

```
status_structs.go
```

---

## ⚙ CLI Options

### Output path

```bash
jinres init -o internal/structs.go
```

Generates a folder (if doesn't exists) whitin the file.

### Change package name

```bash
jinres init -p api
```

Default package is:

```
package jinres
```

or the target folder name if `-o` is used.

---

## 📝 Envelope structure

The runtime response format:

```go
type Jinres struct {
    Status   int    `json:"status"`
    Message  string `json:"message"`
    Response any    `json:"response,omitempty"`
}
```

---

## 🧩 Why Jinres?

* Simple, unified JSON response format
* Fluent builder for Gin handlers
* Auto-generated schemas for Swagger/Swaggo
* Zero-dependency, tiny footprint

---

## 📄 License

[MIT](LICENSE)