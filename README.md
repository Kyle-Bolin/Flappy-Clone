# Flappy Bird Clone

A Flappy Bird clone built in Go using the Ebiten game engine, containerized with Docker.

## Features

- Classic Flappy Bird gameplay mechanics
- Smooth bird physics with gravity and jumping
- Randomly generated pipes with varying heights
- Score tracking
- Game over screen with restart functionality
- Cross-platform support (Windows, macOS, Linux)

## Controls

- **Space Bar** or **Left Mouse Click**: Make the bird jump
- **R Key**: Restart the game (when game over)

## Game Rules

1. Control a yellow bird by making it jump through pipes
2. Avoid hitting the pipes and the screen boundaries
3. Score points by successfully passing through pipes
4. The game ends when you collide with a pipe or hit the screen boundaries
5. Press R to restart after game over

## Prerequisites

- Go 1.21 or later
- Docker (optional, for containerized version)

## Local Development

### Running Locally

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd Flappy-Clone
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the game:
   ```bash
   go run main.go
   ```

### Building Locally

To build the executable:

```bash
go build -o flappy-bird main.go
```

Then run:
```bash
./flappy-bird
```

## Docker

### Building the Docker Image

```bash
docker build -t flappy-bird .
```

### Running with Docker

```bash
docker run -it --rm \
  -e DISPLAY=$DISPLAY \
  -v /tmp/.X11-unix:/tmp/.X11-unix \
  flappy-bird
```

**Note**: For X11 forwarding on macOS, you'll need XQuartz installed and running.

### Alternative: Using Docker Compose

Create a `docker-compose.yml` file:

```yaml
version: '3.8'
services:
  flappy-bird:
    build: .
    environment:
      - DISPLAY=${DISPLAY}
    volumes:
      - /tmp/.X11-unix:/tmp/.X11-unix
    stdin_open: true
    tty: true
```

Then run:
```bash
docker-compose up --build
```

## Project Structure

```
Flappy-Clone/
├── main.go          # Main game logic
├── go.mod           # Go module file
├── go.sum           # Go dependencies checksum
├── Dockerfile       # Docker configuration
├── .dockerignore    # Docker ignore file
└── README.md        # This file
```

## Game Architecture

The game is built using the Ebiten 2D game engine and follows a simple game loop pattern:

- **Game State**: Manages bird position, pipe positions, score, and game over state
- **Update Loop**: Handles physics, input, collision detection, and game logic
- **Draw Loop**: Renders the game state to the screen
- **Input Handling**: Processes keyboard and mouse input

## Dependencies

- `github.com/hajimehoshi/ebiten/v2`: 2D game engine
- `github.com/hajimehoshi/oto/v2`: Audio library (for future sound effects)
- `golang.org/x/image/font/basicfont`: Basic font rendering

## Contributing

Feel free to contribute to this project by:
- Adding sound effects
- Implementing high score tracking
- Adding different bird skins
- Creating power-ups
- Adding multiplayer support

## License

This project is open source and available under the MIT License. 