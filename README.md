# ASCII-ART-WEB

## Description

ASCII-ART-WEB is a web-based application that transforms plain text into visually styled ASCII art. Users can type any text into the browser, choose from three banner styles — **Standard**, **Shadow**, and **Thinkertoy** — and instantly generate the corresponding ASCII art rendered right in the page. The project is built with Go on the backend and served via a lightweight HTTP server, with a clean HTML frontend for user interaction.

---

## Authors

| Role | Username | 
|------|----------|
| Team Captain | **fnya** |
| Team Member | **micsamuel** |
| Team Member | **dofili** |

### Contributions

- **fnya** *(Team Captain)* — Developed the core ASCII-art generation logic (`internal/asciiart.go`) and handled the integration layer that links the backend processing to the frontend rendering pipeline.
- **micsamuel** — Built and styled the frontend (`templates/home.html`), including the input form, banner selector, and the result display section.
- **dofili** — Engineered the backend HTTP server (`cmd/main.go` and `cmd/handler.go`), defining routes, handling requests, and wiring up template execution with proper error handling.

---

## Usage

### Prerequisites

- [Go](https://golang.org/dl/) 1.18 or higher installed on your machine.

### Running the Application

1. **Clone the repository:**

   ```bash
   git clone <repository-url>
   cd ascii-art-web
   ```

2. **Navigate to the server entry point:**

   ```bash
   cd cmd
   ```

3. **Start the server:**

   ```bash
   go run .
   ```

4. **Open your browser and visit:**

   ```
   http://localhost:8080
   ```

### Using the App

1. Type your desired text into the input textarea.
2. Select a banner style from the dropdown: **Standard**, **Shadow**, or **Thinkertoy**.
3. Click the **GENERATE** button.
4. Your ASCII art will be displayed below the form.

> **Note:** Use actual line breaks (press `Enter`) in the textarea to render multi-line ASCII art. Only standard printable ASCII characters (space through `~`) are accepted — submitting non-printable characters or emoji will return a `400 Bad Request` error.

---

## Implementation Details

### Algorithm

The ASCII art generation is handled by the `AsciiGen` function in `internal/asciiart.go`. Here is a step-by-step breakdown of how it works:

#### 1. Banner File Loading
The function takes two arguments — the input `sentence` and a `bannerFile` name (e.g., `"standard"`, `"shadow"`, `"thinkertoy"`). It constructs the file path and reads the corresponding `.txt` banner file from the `banners/` directory.

```
file := "../banners/" + bannerFile + ".txt"
```

Each banner file contains ASCII art representations for all printable characters (ASCII codes 32–126), where each character spans **9 lines** (1 blank separator + 8 art lines).

#### 2. Input Preprocessing
Windows-style line endings (`\r\n`) are normalised to Unix-style (`\n`) using `strings.ReplaceAll` before any further processing. The normalised string is then split on `\n` to produce a slice of individual lines to render.

#### 3. Character Validation
Before any art line is looked up, each character is checked against the printable ASCII range:

```go
if ch < ' ' || ch > '~' {
    return ""
}
```

Any character outside the range `32–126` (space through `~`) — including emoji, control characters, and extended Unicode — causes the function to immediately return an empty string. The HTTP handler treats an empty result as a **400 Bad Request** and returns an appropriate error response to the client.

#### 4. Character Rendering
For each line of input:
- If the line is empty, a blank newline is written directly to the output builder.
- Otherwise, the function iterates over **8 art rows** (rows 1–8 of each character block).
  - For each row, it loops over every character in the line.
  - It calculates the character's starting index in the banner file using the formula:

    ```
    start := ((int(ch) - 32) * 9) + i
    ```

    Where `int(ch) - 32` maps the character to its zero-based position among printable ASCII characters, multiplied by 9 (lines per character), offset by the current row `i`.
  - The corresponding banner line is appended to a `strings.Builder` for efficient string construction.
  - After all characters in a row are written, a newline is appended.

#### 5. Output
The completed `strings.Builder` is returned as a single string, which is passed to the HTML template and rendered inside a `<pre>` block to preserve spacing and alignment.

### HTTP Error Handling

The server responds with the following HTTP status codes:

| Status Code | Condition |
|-------------|-----------|
| `200 OK` | Request succeeded and ASCII art was generated. |
| `400 Bad Request` | Input text is empty or contains non-printable / non-ASCII characters. |
| `404 Not Found` | A route that does not exist was requested. |
| `500 Internal Server Error` | The HTML template could not be parsed or executed (e.g., missing file). |

---

## Testing

The project includes a comprehensive test suite in `cmd/web_test.go` that covers the key behaviours of both the HTTP handlers and the core generation logic.

### Running the Tests

```bash
cd cmd
go test -v
```

### Test Cases

| Test Function | What it verifies |
|---------------|-----------------|
| `TestSuccessfullHome` | `GET /` returns **200 OK** when the template is found and rendered successfully. |
| `TestAscii` | The `AsciiGen` function produces the exact expected ASCII art output for the input `"hello"` using the `standard` banner. |
| `TestAsciiArtHandler` | A valid `POST /ascii-art` request with `inputText=hello&bannerType=standard` returns **200 OK**. |
| `TestNotFound404` | A `GET` request to an unregistered route (e.g. `/this-route-is-fake`) returns **404 Not Found**. |
| `TestBadRequest` | Submitting non-printable characters (emoji) as input returns **400 Bad Request**. |
| `TestInternalServerError500` | When the working directory is changed so the template file cannot be found, `GET /` returns **500 Internal Server Error**. |

---

### Project Structure

```
ascii-art-web/
├── banners/
│   ├── shadow.txt
│   ├── standard.txt
│   └── thinkertoy.txt
├── cmd/
│   ├── handler.go       # HTTP handlers (Home, AsciiArt)
│   ├── main.go          # Server setup and routing
│   └── web_test.go      # Test suite for handlers and ASCII generation
├── internal/
│   └── asciiart.go      # Core ASCII art generation logic
├── templates/
│   └── home.html        # Frontend HTML template
├── go.mod
└── README.md
```