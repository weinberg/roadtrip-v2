package main

import (
  "bufio"
  "flag"
  "fmt"
  "github.com/brickshot/roadtrip-v2/go-client/internal/client/config"
  "github.com/brickshot/roadtrip-v2/go-client/internal/client/ui"
  "log"
  "os"
)

var id string
var reader *bufio.Reader
var screen ui.Screen
var tlsDisabled bool
var host string
var port int

func init() {
  flag.BoolVar(&tlsDisabled, "tls", false, "If false, use TLS. Defaults to false.")
  flag.StringVar(&host, "host", "0.0.0.0", "PlayerServer hostname. Defaults to localhost.")
  flag.IntVar(&port, "port", 9066, "PlayerServer port. Defaults to 9066.")
}

func setup() {
  /*
     Config file has list of characterInfo which have id
     if characterInfo list is empty, add one
     take first entry in characterInfo list and get character from server
  */
  reader = bufio.NewReader(os.Stdin)
  conf, err := config.LoadConfig()
  if err != nil {
    log.Fatalln("Config file error: ", err)
  }

  // if character list is empty, add one
  if conf.Characters == nil || len(conf.Characters) == 0 {
    // createNewCharacter()
    conf, err = config.LoadConfig()
    if err != nil {
      log.Fatalln("Config file error: ", err)
    }
  }

  // currently, only use first character
  id = conf.Characters[0].Id

  fmt.Printf("Character from config file has token: %v\n", id)

  // get character from server

}

/*
func createNewCharacter() *psgrpc.Character {
  fmt.Println("Creating a new character...")
  var name string
  for name == "" {
    fmt.Printf("What would you like your name to be?  ")
    name, _ = reader.ReadString('\n')
    name = strings.TrimRight(name, "\r\n")
    if name == "" {
      fmt.Println("That name is too short.")
    }
  }

  // create in server
  char, err := client.CreateCharacter(getCtx(), &psgrpc.CreateCharacterRequest{
    CaptchaId:     "",
    CharacterName: name,
  })
  st := status.Convert(err)
  if st != nil {
    log.Fatalf("Failed to create character: %v.\n", err)
  }

  // store in config
  _, _, err = config.NewCharacterInfo(char.Id)
  if err != nil {
    log.Fatalln("Cannot create new character: ", err)
  }

  return char
}
*/

func roadTripTitle() {
  text := `
    __ __              ______     
   '  )  )           /   /        
     /--' __ __.  __/ --/__  o _  
    /  \_(_)(_/|_(_/_(_// (_<_/_)_
                             /    
                          __/     
`
  fmt.Println(text)
}

// main
func main() {
  setup()

  screen = ui.Screen{Width: 80, Height: 25}
}
