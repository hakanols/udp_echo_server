package main
import (
    "fmt"
    "net"
    "bufio"
	"os"
)

func main() {
    ip_address := "127.0.0.1:1234"
	message := "Hi mr UDP Server, How are you doing?"

	if 2 < len(os.Args) {
	   fmt.Println("Unvalid numer of arguments")
	   return
	} else if len(os.Args) == 2 {
		ip := os.Args[1]
		if net.ParseIP(ip) == nil {
			fmt.Printf("Invalid IP Address: %s\n", ip)
			return
		}
		ip_address=ip+":1234"
    }

    payload :=  make([]byte, 2048)
    conn, err := net.Dial("udp", ip_address)
    if err != nil {
        fmt.Printf("Some error %v", err)
        return
    }

    fmt.Fprintf(conn, message)
	fmt.Printf("Sent to address: \"%s\" message: \"%s\"\n", ip_address, message)
    _, err = bufio.NewReader(conn).Read(payload)
    if err == nil {
        fmt.Printf("Received message: \"%s\"\n", payload)
    } else {
        fmt.Printf("Some error %v\n", err)
    }
    conn.Close()
}