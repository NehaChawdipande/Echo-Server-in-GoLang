package main
import (
    "fmt"
    "log"
    "net"
)
func main(){
    fmt.Println("Listening on port 3000")
    listener, err := net.Listen("tcp", "localhost:3000")
    if err != nil {
        log.Fatalln(err)
    }
    defer listener.Close()
    for{
        conn, err := listener.Accept()
         if err!= nil{
        log.Fatalln(err)
    }
    fmt.Println("new Connection")

    go listenConnection(conn)
    }
}
func listenConnection(conn net.Conn){
    for{ //listen for messages
        buffer := make([]byte, 1400) 
        dataSize, err := conn.Read(buffer) //reads message and passes to buffer
         if err!= nil{
        fmt.Println("Connection closed")
        return
    }
    data := buffer[:dataSize]
    fmt.Println("recieved message: ", string(data))

    _, err = conn.Write(data) //echo back data
     if err!= nil{
        log.Fatalln(err)
    }
    
    fmt.Println("message sent" , string(data))
    }


}