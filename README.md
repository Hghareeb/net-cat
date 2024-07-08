Sure! Here's a `README.md` file for your TCP chat application:

```markdown
# TCP Chat Application

## Description

This is a TCP-based chat server written in Go that allows multiple clients to connect and chat with each other in real-time. The server maintains a chat history and ensures that usernames are unique. It also allows clients to change their usernames during the chat session.

## Features

- Real-time messaging between multiple clients.
- Chat history is sent to new clients upon joining.
- Clients can change their usernames using the `/name [new_name]` command.
- Maximum of 3 clients can be connected concurrently or pending to join.
- Server logs all activity to a `log.txt` file.

## Usage

### Running the Server

To run the server, use the following command:

```sh
go run main.go [port]
```

- If no port is specified, the server will listen on the default port `8989`.
- If a port is specified, the server will listen on that port.

### Connecting a Client

To connect a client, use a TCP client (like `telnet` or `nc`) to connect to the server:

```sh
telnet localhost [port]
```

or

```sh
nc localhost [port]
```

### Commands

- `[ENTER YOUR NAME]:` - Clients are prompted to enter their name when they first connect.
- `/name [new_name]` - Clients can change their username using this command.

## Dependencies

- Go (https://golang.org)

## Installation

1. Install Go: Follow the instructions at https://golang.org/doc/install to install Go on your system.
2. Clone the repository:
   ```sh
   git clone https://github.com/your-repo/tcp-chat.git
   ```
3. Navigate to the project directory:
   ```sh
   cd tcp-chat
   ```

## Running the Application

1. Start the server:
   ```sh
   go run main.go [port]
   ```
2. Connect clients using `telnet` or `nc`:
   ```sh
   telnet localhost [port]
   ```
   or
   ```sh
   nc localhost [port]
   ```

## Code Structure

- `main.go` - The main file containing the server logic and client handling.
- `log.txt` - Log file where server logs are saved.
- `ping.txt` - for the pinguin pic.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Acknowledgements

- Go Programming Language (https://golang.org)
- Inspiration from various TCP chat applications and tutorials.

## Contact

If you have any questions or suggestions, please open an issue or contact me at [hassannghareeb@gmail.com].
