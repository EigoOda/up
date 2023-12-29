package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli/v2"
)

func main(){
  app := &cli.App{
    Name: "boom",
    Usage: "make an explosive entrance",
    Action: func(*cli.Context) error {
      fmt.Println("BOOM!")
      return nil
    },
  }

  if err := app.Run(os.Args); err != nil {
    log.Fatal(err)
  }
}
