package main

import (
  // built-ins
  "database/sql"
  "os"

  // external dependencies
  _ "github.com/mxk/go-sqlite/sqlite3"
  "github.com/Sirupsen/logrus"
  "github.com/go-gorp/gorp"
  "github.com/nlopes/slack"

  // internal dependencies
  "github.com/znation/chessbot/bot"
  "github.com/znation/chessbot/chess"
  "github.com/znation/chessbot/util"
)

func initDb() *gorp.DbMap {
  // connect to db using standard Go database/sql API
  // use whatever database/sql driver you wish
  db, err := sql.Open("sqlite3", "chessbot.db")
  util.CheckErr(err)

  // construct a gorp DbMap
  dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

  // add a table, setting the table name to 'posts' and
  // specifying that the Id property is an auto incrementing PK
  dbmap.AddTableWithName(chess.Game{}, "games").SetKeys(true, "Id")
  dbmap.AddTableWithName(chess.Board{}, "boards").SetKeys(true, "Id")

  // create the table. in a production system you'd generally
  // use a migration tool, or create the tables via scripts
  err = dbmap.CreateTablesIfNotExists()
  util.CheckErr(err)

  return dbmap
}

func main() {
  api_token := os.Getenv("CHESSBOT_SLACK_API_TOKEN")
  util.CheckCondition(api_token != "", "Environment variable CHESSBOT_SLACK_API_TOKEN must be defined.")

  api := slack.New(api_token)
  rtm := api.NewRTM()


  // initialize the DbMap
  dbmap := initDb()
  defer dbmap.Db.Close()

  go rtm.ManageConnection()

  for {
    select {
      case msg := <-rtm.IncomingEvents:
        switch ev := msg.Data.(type) {
          case *slack.ConnectedEvent:
            logrus.Info("Connection counter:", ev.ConnectionCount)

          case *slack.MessageEvent:
            logrus.Info("Message: ", ev)
            response := bot.HandleMessage(rtm.GetInfo().User.ID, ev.Text)
            if response != "" {
              rtm.SendMessage(rtm.NewOutgoingMessage(response, ev.Channel))
            }

          case *slack.RTMError:
            logrus.Fatal("Error: ", ev.Error())

          case *slack.InvalidAuthEvent:
            logrus.Fatal("Invalid credentials")

          default:
            // Ignore all other events
        }
    }
  }
}
