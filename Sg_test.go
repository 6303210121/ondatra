
package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
)

func main() {
	routerIP := "10.133.35.158"
	sshUsername := "admin"
	sshPassword := "tcs123"

	config := &ssh.ClientConfig{
		User: sshUsername,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", routerIP), config)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session.Close()

	// Execute the first command
	output1, err := session.CombinedOutput("show interface status")
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}

	fmt.Println("Logged into the router.")
	fmt.Printf("Interface Status:\n%s\n", output1)

	// Execute another command
	output2, err := session.CombinedOutput("show system information")
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}

	fmt.Printf("System Information:\n%s\n", output2)
}
