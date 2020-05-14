package main
import (
    "fmt"
    "net"
    "log"
)
func main(){
    conn, err := net.Dial("tcp", "localhost:3000")
     if err!= nil{
        log.Fatalln(err)
    }
 //#1
    _, err = conn.Write([]byte("Hello server!"))
     if err!= nil{
        log.Fatalln(err)
    }
    fmt.Println("Message sent: Hello Server!")

//#2
_, err = conn.Write([]byte("How are you!"))
     if err!= nil{
        log.Fatalln(err)
    }
    fmt.Println("Message sent: How are you!")

    for{
         buffer := make([]byte, 1400)
        dataSize, err := conn.Read(buffer)
         if err!= nil{
        fmt.Println("Connection closed")
        return
    }
    data:= buffer[:dataSize]
    fmt.Println("recieved message: ",string(data))
    }

}
