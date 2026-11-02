# ASCII Studio

A web-based ASCII art generator that converts text into stylized ASCII art using multiple banner styles.

## Overview

ASCII Studio is a Go web application that transforms user input text into ASCII art. It supports multiple banner styles (standard, shadow, thinkertoy) and provides a clean web interface for input and output. The application runs in a browser and includes features for copying and downloading generated art.

## Features

- **ASCII Art Generation**: Converts text input into ASCII art using customizable banner styles
- **Multiple Banner Engines**: Supports standard, shadow, and thinkertoy banner styles
- **Multi-line Input Support**: Handles text with multiple lines and newlines
- **Web UI**: Clean HTML interface with CSS styling
- **Copy to Clipboard**: Easy copying of generated ASCII art
- **Export/Download**: Download ASCII art as a .txt file
- **HTTP Error Handling**: Proper 400, 404, 405, and 500 error responses
- **Static File Serving**: Serves CSS and other static assets
- **Docker Support**: Fully containerized application with multi-stage build

## Project Structure

```
ascii-studio/
├── main.go           # Application entry point and server setup
├── handlers.go       # HTTP request handlers
├── banner.go         # ASCII art rendering and banner loading logic
├── templates/        # HTML templates
│   ├── index.html
│   ├── 400.html
│   ├── 404.html
│   ├── 405.html
│   └── 500.html
├── static/           # Static assets
│   └── style.css
├── banners/          # ASCII banner files
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── Dockerfile        # Multi-stage Docker build
├── go.mod            # Go module dependencies
└── README.md
```

## How It Works

### Request Flow

- **GET /**: Serves the home page with input form
- **POST /ascii-art**: Processes text input and banner selection, generates ASCII art, renders result
- **POST /download**: Exports the generated ASCII art as a downloadable .txt file

### Rendering Flow

1. User submits text and selects banner style via web form
2. Server loads the selected banner file from `banners/` directory
3. Text is processed line-by-line, converting each character to its ASCII representation
4. Result is rendered in the HTML template with copy/download options

### Banner Loading

Banner files contain character mappings for ASCII art generation. Each file defines 8-line representations for printable characters (32-126 ASCII range).

### Export Mechanism

The download feature uses HTTP headers to trigger file download:

- `Content-Disposition: attachment; filename=ascii-art.txt`
- `Content-Type: text/plain`

## Installation & Run (Local)

### Prerequisites

- Go 1.19+
- Access to port 8080

### Run Locally

```bash
go run .
```

Open http://localhost:8080 in your browser.

## Deploy to Vercel

Vercel supports Go backends directly. This project reads Vercel's `PORT` setting and embeds its templates, CSS, and banner files into the Go binary, so no separate static-file setup is needed.

1. Push this project to a GitHub, GitLab, or Bitbucket repository.
2. In Vercel, choose **Add New Project** and import the repository.
3. Keep the project root set to this repository root. No framework preset, build command, output directory, or `vercel.json` is required.
4. Deploy. Vercel will build the Go backend; the resulting URL serves the app.

You can also deploy from the project directory with the Vercel CLI using `vercel` and follow its prompts.

## Docker Setup

### Why Multi-Stage Build?

Multi-stage builds separate the build environment (with Go compiler) from the runtime environment (minimal Alpine), resulting in smaller, more secure images.

### Why Alpine?

Alpine Linux provides a minimal base image (~5MB) with necessary runtime dependencies, keeping the final image size small and secure.

### Why Port 8080?

The application listens on port 8080 internally. Docker maps this to host port 8080 for external access.

### Build Image

```bash
docker build -t ascii-studio .
```

### Run Container

```bash
docker run -p 8080:8080 ascii-studio
```

Access the application at http://localhost:8080.

## Export Feature

The export functionality allows users to download generated ASCII art as a text file.

### How Download Works

1. User clicks download button on result page
2. Server receives POST /download with text and banner data
3. ASCII art is generated server-side
4. Response includes appropriate headers for file download

### HTTP Headers

- `Content-Disposition: attachment; filename=ascii-art.txt` - Triggers browser download dialog
- `Content-Type: text/plain` - Specifies plain text content

## API Endpoints

| Route      | Method | Description        |
| ---------- | ------ | ------------------ |
| /          | GET    | Home page          |
| /ascii-art | POST   | Generate ASCII art |
| /download  | POST   | Download ASCII art |

## Docker Internals

### Image vs Container

- **Image**: Read-only template containing application code and dependencies
- **Container**: Runnable instance of an image with its own filesystem and network

### Container Lifecycle

Containers can be started, stopped, and restarted. Stopped containers retain their state until removed.

### Common Commands

```bash
# List running containers
docker ps

# List all images
docker images

# View container logs
docker logs <container_id>

# Stop a running container
docker stop <container_id>
```

## Troubleshooting

### Port Already Allocated

**Error**: "Bind for 0.0.0.0:8080 failed: port is already allocated"
**Fix**: Stop conflicting container or use different port

```bash
docker ps -a
docker stop <container_id>
# Or run on different port
docker run -p 9090:8080 ascii-studio
```

### CSS Not Loading

**Issue**: Styles not applying in browser
**Fix**: Check static file serving and paths

```bash
docker exec -it <container_id> ls -la /app/static
# Ensure style.css exists and is readable
```

### Container Exits Immediately

**Issue**: Container stops right after starting
**Fix**: Check for runtime errors

```bash
docker logs <container_id>
# Look for panic messages or bind errors
```

### Download Not Working

**Issue**: Download button doesn't trigger file save
**Fix**: Verify Content-Disposition header is set correctly in handlers.go

## Tech Stack

- **Go**: Backend server using net/http
- **HTML**: Frontend templates
- **CSS**: Styling for web interface
- **Docker**: Containerization

## Author

#ASSEM
#AMMAR
#RAHMA
**ASCII Studio Team**
