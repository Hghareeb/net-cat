# net-cat

TCP Chat Server
Overview
This project is a simple TCP-based chat server written in Go. It allows multiple clients to connect to the server, send messages to each other, and change their display names. The server also maintains a chat history that is sent to new clients when they join.
Features
Multi-client support: Up to 10 clients can connect to the server simultaneously.
Chat history: New clients receive the chat history upon joining.
Name change: Clients can change their display names using the /name command.
Logging: All chat messages and events are logged to a file (log.txt).
Prerequisites
Go 1.16 or later
Installation
Clone the repository:



git clone https://github.com/yourusername/tcp-chat-server.git
cd tcp-chat-server
Build the server:
sh§


go build -o TCPChat main.go
Usage
Start the server:

./TCPChat [port]
If no port is specified, the server will default to port 8080.
Example: ./TCPChat 9090 will start the server on port 9090.
Connect clients:
Use a TCP client application or script to connect to the server. For example, you can use telnet:
sh


telnet localhost 8080
Enter your name:
When prompted, enter your display name. The name should not contain spaces or special characters.
Send messages:
Type your message and press Enter to send it to all connected clients.
Change your name:
Use the /name command followed by your new name to change your display name.
Example: /name NewName
If you use the /name command without providing a new name, you will receive a usage message: Usage: /name [new_name].
Example
Start the server:
sh


./TCPChat
Connect a client:
sh


telnet localhost 8080
Enter your name:
javascript


[ENTER YOUR NAME]: Alice
Send a message:
javascript


Hello, everyone!
Change your name:
javascript


/name Bob
Receive a usage message:
javascript


// name Usage: /name [new_name]

Utility Functions
getClientName: Prompts the client to enter their name.
sendChatHistory: Sends the chat history to a new client.
recordMessage: Records a message in the chat history.
broadcast: Sends a message to all connected clients.
removeClient: Removes a client from the list of active clients.
License

This project is licensed under the MIT License. See the LICENSE file for details.


