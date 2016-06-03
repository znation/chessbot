package bot

import (
  // language built-ins
  "strings"

  // external dependencies
  "github.com/Sirupsen/logrus"
)

func misunderstood() string {
  return "Sorry, I didn't understand your request.\n\n" + help()
}

func help() string {
  return `I support the following commands:
* help
`
}

func HandleMessage(userId string, msg string) string {
  prefix := "<@" + userId + ">: "
  if strings.HasPrefix(msg, prefix) {
    msgContents := strings.TrimPrefix(msg, prefix)
    switch msgContents {
      case "help":
        return help()
      default:
        return misunderstood()
    }
  } else {
    logrus.Info("Ignoring message \"", msg, "\". Did not match prefix \"", prefix, "\".")
    return ""
  }
}
