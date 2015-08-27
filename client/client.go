package main

import (
  "fmt"
  "log"
  "net"
)

 func handleConnection(c net.Conn) {

  log.Printf("Server %v connected.", c.RemoteAddr())

  // stuff to do... like read data from client, process it, write back to client
  // see what you can do with (c net.Conn) at
  // http://golang.org/pkg/net/#Conn

  buffer := make([]byte, 4096)
  
  n, err := c.Write([]byte("hello world"))
  if err != nil {
    c.Close()
  }
  n, err = c.Read(buffer)
  if err != nil || n == 0 {
    c.Close()
  }
  
  fmt.Println(buffer);
  log.Printf("Connection from %v closed.", c.RemoteAddr())
}

func main() {
  hostName := "localhost" // change this
  portNum := "3000"

  conn, err := net.Dial("tcp", hostName+":"+portNum)

  if err != nil {
    fmt.Println(err)
    return
  }

  handleConnection(conn)

  fmt.Printf("Connection established between %s and localhost.\n", hostName)
  fmt.Printf("Remote Address : %s \n", conn.RemoteAddr().String())
  fmt.Printf("Local Address : %s \n", conn.LocalAddr().String())

}