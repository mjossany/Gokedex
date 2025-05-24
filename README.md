# Gokedex

Gokedex is a CLI application that allows you to interact with the PokeAPI. You can explore different locations, find Pokémon, and catch them.

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

*   `help`: Displays a help message with all available commands.
*   `exit`: Exits the Pokedex application.
*   `map`: Lists the next 20 location areas in the Pokémon world.
*   `mapb`: Lists the previous 20 location areas in the Pokémon world.
*   `explore <location_name>`: Lists all Pokémon encounters in the specified location area.
*   `catch <pokemon_name>`: Attempts to catch the specified Pokémon.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.

## License

This project is licensed under the MIT License.
