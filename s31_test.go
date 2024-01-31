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
