# ASCII-ART-WEB-STYLIZE

## Description

ASCII-ART-WEB-STYLIZE is a Go web application that converts plain text into styled ASCII art.

The project supports multiple ASCII banner styles and introduces customizable colorization, responsive frontend design, and animated visual effects.

### Supported Banner Styles
- standard
- shadow
- thinkertoy

### Features
- ASCII art generation
- Full-text or substring colorization
- Responsive UI
- Animated backgrounds
- Secure HTML rendering
- HTTP error handling
- Unit testing

---

## Requirements

- Go 1.18 or higher

---

## Installation

### Clone the Repository

```bash
git clone https://acad.learn2earn.ng/git/aiwueze/ascii-art-web-stylize
cd ascii-art-web-stylize
```

### Run the Application

Run the server from the project root:

```bash
go run ./cmd/main.go
```

### Open in Browser

```text
http://localhost:8080
```

---

## Usage

1. Enter text into the input field
2. Select a banner style:
   - Standard
   - Shadow
   - Thinkertoy
3. Optionally enter a substring to color
4. Select a color
5. Click **Generate**
6. View the generated ASCII art output

---

## Project Structure

```text
ascii-art-web-stylize/
├── banners/
│   ├── shadow.txt
│   ├── standard.txt
│   └── thinkertoy.txt
├── cmd/
│   └── main.go
├── internal/
│   ├── generator/
│   │   ├── asciiart.go
│   │   └── generator_test.go
│   └── handlers/
│       ├── handler.go
│       └── handler_test.go
├── static/
│   └── style.css
├── templates/
│   └── home.html
├── go.mod
└── README.md
```

---

## Implementation Overview

### ASCII Generation

The ASCII generation logic is implemented in:

```text
internal/generator/asciiart.go
```

The generator:
- Reads banner templates from `.txt` files
- Maps printable ASCII characters to banner rows
- Builds ASCII output row-by-row
- Supports optional substring colorization

### Static File Serving

Static assets are served using Go’s file server utilities.

### HTML Rendering

Generated output is safely rendered using Go templates and controlled HTML injection for styled ASCII output.

---

## Testing

Run all tests with:

```bash
go test ./...
```

### Included Tests

| Test | Purpose |
|---|---|
| TestSuccessfullHome200 | Verifies homepage rendering |
| TestAscii | Verifies ASCII generation correctness |
| TestSuccesfulAsciiArt200 | Verifies successful POST request handling |
| TestNotFound404 | Verifies 404 handling |
| TestBadRequest400 | Verifies invalid input handling |
| TestInternalServerError500 | Verifies template loading failure handling |

---

## Author

Dillon Ofili


---

## License

This project is intended for educational and learning purposes.