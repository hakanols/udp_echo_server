package main
import (
    "fmt"
    "net"
)

func main() {
    localaddr := net.UDPAddr{
        Port: 1234,
        IP: net.ParseIP("0.0.0.0"),
    }
    ser, err := net.ListenUDP("udp", &localaddr)
    if err != nil {
        fmt.Printf("Some error %v\n", err)
        return
    }
    fmt.Printf("Start listening on UDP port: %d\n", localaddr.Port)
    for {
        payload := make([]byte, 2048)

        count, remoteaddr, errread := ser.ReadFromUDP(payload)
        if errread !=  nil {
            fmt.Printf("Error reading message: %v\n", errread)
            continue
        }

        fmt.Printf("Echo message \"%s\" from %v\n", payload, remoteaddr)
        _, errwrite := ser.WriteToUDP(payload[0:count], remoteaddr)
        if errwrite != nil {
            fmt.Printf("Couldn't send response: %v\n", errwrite)
        }
    }
}
