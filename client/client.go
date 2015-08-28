package main

import (
  "fmt"
  "log"
  "net"
  // "net/http"
  "io/ioutil"
)

func check(err error) {
    if err != nil {
      panic(err)
    }
}

func handleConnection(connection net.Conn) {

  log.Printf("Server %v connected.", connection.RemoteAddr())

  // stuff to do... like read data from client, process it, write back to client
  // see what you can do with (c net.Conn) at
  // http://golang.org/pkg/net/#Conn
 
  buffer := make([]byte, 4096)
  
  dat, err := ioutil.ReadFile("../media/Wzlogo.mp3")
  check(err)
  max := 0
  //buffers every 4096 bytes
  for i := 0; i < len(dat); i += 4096 {
    if i + 4096 < len(dat) {
      max = i + 4096
    } else {
      max = len(dat) - 1
    }
    //sending to server
    length, err := connection.Write([]byte(dat[i:max]))
    if err != nil {
      connection.Close()
    }
    //server reads the file
    length, err = connection.Read(buffer)
    if err != nil || length == 0 {
      connection.Close()
    }
    fmt.Println(buffer);
  }
  
  log.Printf("Connection from %v closed.", connection.RemoteAddr())
}

func getFiles(dir string) {
  files, err := ioutil.ReadDir(dir)
  check(err)
  for _, file := range files { 
    fmt.Println(file.Name())
  }
}

func main() {
  // resp, err := http.Get("http://media-stream-1049.appspot.com/test")

  hostName := "" // change this
  portNum := "3000"
  conn, err := net.Dial("tcp", hostName+":"+portNum)
  

  if err != nil {
    fmt.Println(err)
    return
  }

  // defer resp.Body.Close()
  // body, err := ioutil.ReadAll(resp.Body)
  // s := string(body[:])
  // fmt.Println(s)

  handleConnection(conn)
  // getFiles("../")

  fmt.Printf("Connection established between %s and localhost.\n", hostName)
  fmt.Printf("Remote Address : %s \n", conn.RemoteAddr().String())
  fmt.Printf("Local Address : %s \n", conn.LocalAddr().String())

}