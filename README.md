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
sh


git clone https://github.com/yourusername/tcp-chat-server.git
cd tcp-chat-server
Build the server:
sh


go build -o TCPChat main.go
Usage
Start the server:
sh


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


/name
Usage: /name [new_name]
Code Explanation
Main Function
The main function initializes the server, sets up logging, and listens for incoming connections:
go


func main() {
    // ... (initialization and setup code)
    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Printf("Error accepting connection: %v", err)
            continue
        }

        // Limit the number of clients to 10
        if len(clients) >= 10 {
            log.Println("Maximum number of clients reached. Connection rejected.")
            conn.Write([]byte("Maximum number of clients reached. Connection rejected."))
            conn.Close()
            continue
        }

        go handleClient(conn)
    }
}
Handling Clients
The handleClient function manages client interactions, including receiving messages and handling name changes:
go


func handleClient(conn net.Conn) {
    defer conn.Close()

    // Get client's name
    name := getClientName(conn)
    if name == "" {
        return
    }

    // Create a new client object
    client := Client{
        NetConn: conn,
        Name:    name,
        Reader:  bufio.NewReader(conn),
    }

    // Add the client to the list of active clients
    mutex.Lock()
    clients = append(clients, client)
    mutex.Unlock()

    // Send chat history to the client
    sendChatHistory(client)

    // Notify all clients that a new client has joined
    broadcast(fmt.Sprintf("[%s] %s joined the chat", time.Now().Format("2006-01-02 15:04:05"), name))
    log.Printf("%s has joined the chat.", name)

    // Read and broadcast messages from the client
    for {
        msg, err := client.Reader.ReadString('\n')
        if err != nil {
            log.Printf("%s has left the chat.", name)
            break
        }
        msg = strings.TrimSpace(msg)
        if msg == "" || strings.Contains(msg, "\033") {
            continue
        }

        if strings.HasPrefix(msg, "/name ") {
            // Extract the new name from the message
            newName := strings.TrimSpace(strings.TrimPrefix(msg, "/name "))
            if newName == "" {
                client.NetConn.Write([]byte("Usage: /name [new_name]\n"))
                continue
            }
            // Validate the new name
            if strings.Contains(newName, " ") || strings.Contains(newName, "\033") {
                client.NetConn.Write([]byte("Please enter a valid name.\n"))
                continue
            }
            // Store the old name and update the client's name
            oldName := client.Name
            client.Name = newName
            // Broadcast the name change to all clients
            broadcast(fmt.Sprintf("[%s] %s changed their name to %s", time.Now().Format("2006-01-02 15:04:05"), oldName, newName))
            log.Printf("%s changed their name to %s.", oldName, newName)
            continue
        }

        formattedMessage := fmt.Sprintf("[%s][%s]: %s", time.Now().Format("2006-01-02 15:04:05"), client.Name, msg)
        broadcast(formattedMessage)
        recordMessage(formattedMessage)
    }

    // Remove the client from the list of active clients
    removeClient(client)

    // Notify all clients that a client has left
    broadcast(fmt.Sprintf("[%s] %s has left the chat", time.Now().Format("2006-01-02 15:04:05"), name))
}
Utility Functions
getClientName: Prompts the client to enter their name.
sendChatHistory: Sends the chat history to a new client.
recordMessage: Records a message in the chat history.
broadcast: Sends a message to all connected clients.
removeClient: Removes a client from the list of active clients.
License

This project is licensed under the MIT License. See the LICENSE file for details.


