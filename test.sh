#!/bin/bash

# Build the TCP chat server
go build -o TCPChat main.go

check_server_response() {
  sleep 2
  local response=$(nc -zv localhost $1 2>&1)
  if [[ $response == *"succeeded"* ]]; then
    echo "Server is listening on port $1"
  else
    echo "Server is not listening on port $1"
  fi
}

check_usage_response() {
  local response=$(./TCPChat 2525 localhost 2>&1)
  if [[ $response == *"[USAGE]: ./TCPChat \$port"* ]]; then
    echo "Server responded with usage message"
  else
    echo "Server did not respond with usage message"
  fi
}

simulate_client() {
  local client_name=$1
  local commands=$2
  echo "Simulating client $client_name"
  {
    echo "$client_name"
    sleep 1
    cat $commands
  } | nc localhost $3 > "client_${client_name}.log" &
}

# Test default port listening
echo "Testing default port listening..."
./TCPChat &
SERVER_PID=$!
check_server_response 8080
kill $SERVER_PID

# Test usage response
echo "Testing usage response..."
check_usage_response

# Test specific port listening
echo "Testing specific port listening on 2525..."
./TCPChat 2525 &
SERVER_PID=$!
check_server_response 2525
kill $SERVER_PID

# Simulate multiple clients
echo "Simulating multiple client connections..."
./TCPChat 2525 &
SERVER_PID=$!
sleep 2

# Create command files for clients
echo "Hello, this is Client1" > client1_commands.txt
echo "/name NewClient1" >> client1_commands.txt
echo "This is a message from NewClient1" >> client1_commands.txt

echo "Hello, this is Client2" > client2_commands.txt
echo "This is a message from Client2" >> client2_commands.txt

simulate_client "Client1" "client1_commands.txt" 2525
simulate_client "Client2" "client2_commands.txt" 2525

sleep 5

# Check client logs
echo "Client1 log:"
cat client_Client1.log

echo "Client2 log:"
cat client_Client2.log

# Clean up
echo "Cleaning up..."
kill $SERVER_PID
rm TCPChat
rm client1_commands.txt client2_commands.txt client_Client1.log client_Client2.log

echo "Test completed."