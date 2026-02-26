# rocketGo

A space rocket game built with Go and Ebiten game engine.

## Overview

rocketGo is an interactive game where you control a rocket's altitude using the spacebar. The longer you hold the space key, the more thrust is applied, increasing your acceleration. Release the key and gravity pulls you back down.

**⚠️ Note:** This project is currently in active development. More functionality and features will be added soon!

## Requirements

- Go 1.16 or higher
- Ebiten v2 game engine

## Installation

1. Clone or download this project to your local machine
2. Navigate to the project directory:
   ```bash
   cd c:\DATA\rocketGo
   ```
3. Install dependencies:
   ```bash
   go mod download
   ```

## Running the Game

### Option 1: Using Make (Recommended)
```bash
make run
```

### Option 2: Direct Go Command
```bash
go run main.go
```

## Building

To build an executable:
```bash
make build
```

This creates an executable file in the project directory.

## Controls

- **Spacebar:** Apply thrust to increase altitude
- Hold longer for greater acceleration

## Project Structure

```
rocketGo/
├── main.go          # Main game logic
├── assets/          # Game assets (images)
│   ├── background.png
│   └── player.png
├── README.md        # This file
├── Makefile         # Build automation
└── go.mod           # Go module dependencies
```

## Upcoming Features

- [ ] Better physics simulation
- [ ] Score/altitude tracking
- [ ] Obstacles and enemies
- [ ] Sound effects
- [ ] Multiple levels
- [ ] High score system

## Development

This project is actively being developed. Check back soon for updates!# spaceXGo
