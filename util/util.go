package util

import (
  // external dependencies
  "github.com/Sirupsen/logrus"
)

func CheckCondition(condition bool, msg string) {
  if !condition {
    logrus.Fatal(msg)
  }
}

func CheckErr(err error) {
  if err != nil {
    logrus.Fatal(err.Error())
  }
}
