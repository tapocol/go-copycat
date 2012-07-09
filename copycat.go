package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "hub"
)

type CopyCatMessage struct {
  sender string
  message string
}

func (m CopyCatMessage) Format() string {
  return m.sender + ": " + m.message
}

var EXIT_MESSAGE string = "stfu\n"
var sender, message string
var cc_message CopyCatMessage
var err error

func main() {
  sender = "you"
  h := hub.Hub{
    Handler: func(m hub.HubMessage) { fmt.Println("Carbon Copy says to " + m.Format()) },
    Broadcast: make(chan hub.HubMessage),
  }
  go h.Run()
  in := bufio.NewReader(os.Stdin)
  message := "Welcome to CopyCat. Go ahead and talk to me. Just don't tell me to stfu."
  var err error = nil
  for err == nil && !strings.Contains(EXIT_MESSAGE, message) {
    cc_message = CopyCatMessage{
      sender: sender,
      message: message,
    }
    h.Broadcast <- cc_message
    message, err = in.ReadString('\n')
  }
  cc_message = CopyCatMessage{
    sender: sender,
    message: "Understood",
  }
  h.Handler(cc_message)
}


