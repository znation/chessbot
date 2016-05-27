package main

import (
  "fmt"
  "os"
  "github.com/nlopes/slack"
  "github.com/mxk/go-sqlite/sqlite3"
)

func main() {
  api_token := os.Getenv("CHESSBOT_SLACK_API_TOKEN")
  api := slack.New(api_token)
  api.SetDebug(true)

  rtm := api.NewRTM()
  go rtm.ManageConnection()
  for {
    select {
    case msg := <-rtm.IncomingEvents:
      fmt.Print("Event Received: ")
      switch ev := msg.Data.(type) {
      case *slack.HelloEvent:
        // Ignore hello

      case *slack.ConnectedEvent:
        fmt.Println("Infos:", ev.Info)
        fmt.Println("Connection counter:", ev.ConnectionCount)
        // Replace #general with your Channel ID
        rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "#general"))

      case *slack.MessageEvent:
        fmt.Printf("Message: %v\n", ev)

      case *slack.PresenceChangeEvent:
        fmt.Printf("Presence Change: %v\n", ev)

      case *slack.LatencyReport:
        fmt.Printf("Current latency: %v\n", ev.Value)

      case *slack.RTMError:
        fmt.Printf("Error: %s\n", ev.Error())

      case *slack.InvalidAuthEvent:
        fmt.Printf("Invalid credentials")
        break Loop

      default:

        // Ignore other events..
        // fmt.Printf("Unexpected: %v\n", msg.Data)
      }
    }
  }
}
