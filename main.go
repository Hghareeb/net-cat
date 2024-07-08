package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type Client struct {
	NetConn net.Conn
	Name    string
	Reader  *bufio.Reader
}

var clients []Client
var mutex sync.Mutex
var history []string
var pendingClients int

func main() {
	logFile, err := os.Create("log.txt")
	if err != nil {
		log.Fatal("Error opening logfile:", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile) // Save logs to file
	log.Printf("Welcome to TCP-Chat!\n")

	var port string
	if len(os.Args) == 1 {
		port = ":8989"
		fmt.Printf("Listening on the port %s\n", port)
	} else if len(os.Args) == 2 {
		port = ":" + os.Args[1]
		fmt.Printf("Listening on the port %s\n", port)
	} else {
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(0)
	}

	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		mutex.Lock()
		if len(clients) >= 3 || pendingClients >= 3 {
			mutex.Unlock()
			log.Println("Maximum number of clients reached or too many pending clients. Connection rejected.")
			conn.Write([]byte("Maximum number of clients reached or too many pending clients. Connection rejected.\n"))
			conn.Close()
			continue
		}
		pendingClients++
		mutex.Unlock()

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Get client's name
	name := getClientName(conn)
	if name == "" {
		mutex.Lock()
		pendingClients--
		mutex.Unlock()
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
	pendingClients--
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

func getClientName(conn net.Conn) string {
	// Create a reader to read from the connection
	reader := bufio.NewReader(conn)
	penguin, err := os.ReadFile("ping.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	var name string
	for {
		// Prompt the client to enter their name
		mutex.Lock()
		conn.Write([]byte((penguin)))
		conn.Write([]byte("[ENTER YOUR NAME]: "))
		mutex.Unlock()

		// Read a line of input from the client
		nameBytes, err := reader.ReadBytes('\n')
		if err != nil {
			log.Printf("Error reading client name: %v", err)
			return ""
		}
		name = strings.TrimSpace(string(nameBytes))

		// Check if the name contains spaces or is empty
		if name == "" || strings.Contains(name, " ") || strings.Contains(name, "\033") {
			conn.Write([]byte("Please enter a valid name.\n"))
			continue // Ask the user to re-enter the name
		}

		// If the name is valid, break out of the loop
		break
	}

	return name
}

func sendChatHistory(client Client) {
	// sends chat history to clients network
	for _, historyMsg := range history {
		_, err := client.NetConn.Write([]byte(historyMsg + "\n"))
		if err != nil {
			log.Printf("Error sending chat history to %s: %v", client.Name, err)
		}
	}
}

func recordMessage(msg string) {
	// save new msgs in history
	log.Println(msg)
	history = append(history, msg)
}

func broadcast(msg string) {
	for _, client := range clients {
		// iterates over clients
		_, err := client.NetConn.Write([]byte(msg + "\n"))
		// write the msgs in the clients network
		if err != nil {
			log.Printf("Error broadcasting message to %s: %v", client.Name, err)
		}
	}
}

func removeClient(client Client) {
	for i, c := range clients {
		if c == client {
			mutex.Lock()
			clients = append(clients[:i], clients[i+1:]...)
			// append without the index of the client that will get removed
			mutex.Unlock()
			break
		}
	}
}