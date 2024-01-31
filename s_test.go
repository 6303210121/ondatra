package debug

import (
        "fmt"
        "golang.org/x/crypto/ssh"
        "log"
	"testing"
)

func main2() {

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


        output, err := session.CombinedOutput("show interface status")
        if err != nil {
                log.Fatalf("Failed to execute command: %v", err)
        }

        fmt.Println("Logged into the router.")
        fmt.Printf("Interface Status:\n%s\n", output)
}
func Testsample2 (t *testing.T) {
	//return "nil"
        //ds
}
