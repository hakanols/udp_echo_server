package main
import (
    "fmt"
    "net"
    "bufio"
	"os"
)

func main() {
    ip_address := "127.0.0.1"

	if 2 < len(os.Args) {
	   fmt.Println("Unvalid numer of arguments")
	   return
	} else if len(os.Args) == 2 {
		ip_address = os.Args[1]
		if net.ParseIP(ip_address) == nil {
			fmt.Printf("Invalid IP Address: %s\n", ip_address)
			return
		}
    }

	fmt.Printf("IP Address: %s\n", ip_address)

    p :=  make([]byte, 2048)
    conn, err := net.Dial("udp", ip_address+":1234")
    if err != nil {
        fmt.Printf("Some error %v", err)
        return
    }
    fmt.Fprintf(conn, "Hi UDP Server, How are you doing?")
    _, err = bufio.NewReader(conn).Read(p)
    if err == nil {
        fmt.Printf("%s\n", p)
    } else {
        fmt.Printf("Some error %v\n", err)
    }
    conn.Close()
}