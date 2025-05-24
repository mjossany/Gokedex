# Gokedex

Gokedex is a command-line interface (CLI) application written in Go that allows you to interact with the [PokeAPI](https://pokeapi.co/). You can explore different locations in the Pokémon world, discover Pokémon in those locations, catch them, and manage your personal Pokédex collection.

## Features

- **Location Exploration**: Browse through different areas in the Pokémon world
- **Pokémon Discovery**: Find Pokémon encounters in specific locations  
- **Pokémon Catching**: Attempt to catch Pokémon you encounter
- **Personal Pokédex**: Keep track of all the Pokémon you've caught
- **Pokémon Inspection**: View detailed information about your caught Pokémon
- **Response Caching**: Built-in caching system for improved performance (5-minute cache interval)

## Prerequisites

- Go 1.24.2 or later

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/mjossany/Gokedex.git
   ```
2. Navigate to the project directory:
   ```bash
   cd Gokedex
   ```
3. Build the executable:
   ```bash
   go build
   ```

## Usage

Run the application from the command line:

```bash
./Gokedex
```

Once the application is running, you will see the `Pokedex >` prompt. You can then enter commands to interact with the application.

## Commands

The following commands are available:

- **`help`**: Displays a help message with all available commands and their descriptions.
- **`exit`**: Exits the Pokédex application.
- **`map`**: Lists the next 20 location areas in the Pokémon world.
- **`mapb`**: Lists the previous 20 location areas in the Pokémon world.
- **`explore <location_name>`**: Lists all Pokémon encounters in the specified location area.
- **`catch <pokemon_name>`**: Attempts to catch the specified Pokémon. Success is not guaranteed!
- **`pokedex`**: Displays all the Pokémon you have successfully caught.
- **`inspect <pokemon_name>`**: Shows detailed information about a caught Pokémon, including its stats, types, height, and weight.

### Example Usage

```bash
Pokedex > map
Pokedex > explore canalave-city-area
Pokedex > catch pikachu
Pokedex > pokedex
Pokedex > inspect pikachu
```

## Project Structure

```
Gokedex/
├── internal/
│   ├── pokeapi/     # PokeAPI client implementation
│   ├── pokecache/   # Caching system
│   └── pokedex/     # Pokédex data management
├── command_*.go     # Individual command implementations
├── main.go          # Application entry point
├── repl.go          # Read-Eval-Print Loop implementation
└── go.mod           # Go module file
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.

## License

This project is licensed under the MIT License.
