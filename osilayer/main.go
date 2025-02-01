package main

import (
	"encoding/json"
	"log"
	"net"
	"time"
)

// Message represents a simple payload being sent across layers
type Message struct {
	Data string `json:"data"`
}

func main() {
	go runServer()
	time.Sleep(1 * time.Second) // Wait for server to start
	runClient()
}

// Server simulates receiving and processing data
func runServer() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listener.Close()

	log.Println("Server is running on port 8080...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Println("[Session Layer] Connection established")

	// Read data from client (Transport Layer)
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error reading data: %v", err)
		return
	}

	// Decode data (Presentation Layer)
	var msg Message
	err = json.Unmarshal(buffer[:n], &msg)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		return
	}

	// Application Layer Processing
	log.Printf("[Application Layer] Received: %s", msg.Data)

	// Respond to client
	response := Message{Data: "Message received: " + msg.Data}
	respBytes, _ := json.Marshal(response)
	conn.Write(respBytes)
	log.Println("[Session Layer] Response sent")
}

// Client simulates sending data
func runClient() {
	log.Println("Client is starting...")

	// Establish connection (Session Layer)
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	defer conn.Close()

	log.Println("[Session Layer] Connected to server")

	// Prepare data (Application Layer)
	message := Message{Data: "Hello from Client"}
	data, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	// Send data (Transport Layer)
	_, err = conn.Write(data)
	if err != nil {
		log.Fatalf("Error sending data: %v", err)
	}
	log.Println("[Transport Layer] Data sent")

	// Receive response (Transport Layer)
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	// Decode response (Presentation Layer)
	var response Message
	err = json.Unmarshal(buffer[:n], &response)
	if err != nil {
		log.Fatalf("Error decoding response JSON: %v", err)
	}

	// Log response (Application Layer)
	log.Printf("[Application Layer] Server response: %s", response.Data)
}
