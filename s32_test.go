package debug

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"
)

func main() {
	routerIP := "10.133.35.158"
	routerUsername := "tcs"
	routerPassword := "tcs123"
	interfaceCheckCommand := "show interfaces"
	conn, err := connectToRouter(routerIP)
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

func connectToRouter(routerIP string) (net.Conn, error) {
	conn, err := net.Dial("tcp", routerIP+":23")
	if err != nil {
		return nil, err
	}
	return conn, nil
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

func TestSampleIntegration(t *testing.T) {
	routerIP := "10.133.35.158"
	conn, err := connectToRouter(routerIP)
	if err != nil {
		t.Fatalf("Error connecting to router: %v", err)
	}
	defer conn.Close()

	response := executeLogic(conn)

	expectedResponse := "Interface status: up"
	if !strings.Contains(response, expectedResponse) {
		t.Errorf("Unexpected router response: got %s, want %s", response, expectedResponse)
	}
}

func executeLogic(conn net.Conn) string {
	readWelcomeMessages(conn)
	sendCommand(conn, "tcs")
	sendCommand(conn, "tcs123")
	time.Sleep(2 * time.Second)
	sendCommand(conn, "show interfaces")
	return readResponse(conn)
}
