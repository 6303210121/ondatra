package debug

import (
	"bufio"
	"fmt"
	"net"
	"time"
	"strings"
	"testing"
)

func main() {
	routerIP := "10.133.35.158"
	routerUsername := "tcs"
	routerPassword := "tcs123"
	interfaceCheckCommand := "show interfaces"
	conn, err := net.Dial("tcp", routerIP+":23")
	if err != nil {
		fmt.Println("Error connecting to the router:", err)
		return
	}
	defer conn.Close()
	readWelcomeMessages(conn)
	sendCommand(conn, routerUsername)
	sendCommand(conn, routerPassword)
	time.Sleep(2 * time.Second)
	sendCommand(conn, interfaceCheckCommand)
	response := readResponse(conn)
	fmt.Println("Router response:", response)
}

func readWelcomeMessages(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil || containsLoginPrompt(message) {
			break
		}
	}
}

func sendCommand(conn net.Conn, command string) {
	fmt.Fprintln(conn, command)
	time.Sleep(2 * time.Second)
}

=====================================

func TestSample(t *testing.T) {
	// Setup mock TCP connection
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to start mock TCP server: %v", err)
	}
	defer listener.Close()

	go func() {
		// Accept the incoming connection
		conn, err := listener.Accept()
		if err != nil {
			t.Fatalf("failed to accept connection: %v", err)
		}
		defer conn.Close()

		// Simulate router responses
		fmt.Fprintln(conn, "Welcome to the router")
		fmt.Fprintln(conn, "Username:")
		fmt.Fprintln(conn, "Password:")
		fmt.Fprintln(conn, "show interfaces")
		fmt.Fprintln(conn, "Interface status: up")
	}()

	// Connect to the mock TCP server
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatalf("failed to connect to mock TCP server: %v", err)
	}
	defer conn.Close()

	// Execute the main function
	routerIP := listener.Addr().String()
	routerUsername := "tcs"
	routerPassword := "tcs123"
	response := executeMainFunction(conn, routerIP, routerUsername, routerPassword)

	// Validate the response
	expectedResponse := "Interface status: up"
	if response != expectedResponse {
		t.Errorf("unexpected router response: got %s, want %s", response, expectedResponse)
	}
}

func executeMainFunction(conn net.Conn, routerIP, routerUsername, routerPassword string) string {
	readWelcomeMessages(conn)
	sendCommand(conn, routerUsername)
	sendCommand(conn, routerPassword)
	time.Sleep(2 * time.Second)
	sendCommand(conn, "show interfaces")
	return readResponse(conn)
}

func readResponse(conn net.Conn) string {
	reader := bufio.NewReader(conn)
	response, _ := reader.ReadString('\n')
	return response
}

func containsLoginPrompt(message string) bool {
	return strings.Contains(message, "login") || strings.Contains(message, "password")
}


func isInterfaceUp(response string) bool {
	return strings.Contains(response, "up")
}

func isInterfaceDown(response string) bool {
	return strings.Contains(response, "down")
}

func Testsample (t *testing.T) {
        //return "nil"
        
}
