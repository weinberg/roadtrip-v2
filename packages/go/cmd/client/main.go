package main

import (
  "bufio"
  "context"
  "flag"
  "fmt"
  types "github.com/brickshot/roadtrip-v2/internal/client"
  "github.com/brickshot/roadtrip-v2/internal/client/config"
  "github.com/brickshot/roadtrip-v2/internal/client/ui"
  "github.com/hasura/go-graphql-client"
  "log"
  "net/http"
  "os"
  "strconv"
  "time"
)

var token string
var reader *bufio.Reader
var screen ui.Screen
var tlsDisabled bool
var host string
var port int
var client *graphql.Client

var currentCharacter types.CurrentCharacter

func init() {
  flag.BoolVar(&tlsDisabled, "tls", false, "If false, use TLS. Defaults to false.")
  flag.StringVar(&host, "host", "0.0.0.0", "PlayerServer hostname. Defaults to localhost.")
  flag.IntVar(&port, "port", 8080, "PlayerServer port. Defaults to 8080.")
}

/**
 * Code to add Authorization header with token to http client transport
 */
type AddHeaderTransport struct {
  T http.RoundTripper
}

func (adt *AddHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
  req.Header.Add("Authorization", token)
  return adt.T.RoundTrip(req)
}

func NewAddHeaderTransport(T http.RoundTripper) *AddHeaderTransport {
  if T == nil {
    T = http.DefaultTransport
  }
  return &AddHeaderTransport{T}
}

/**
 * Load initial graphql payload
 */
func getCurrentCharacter() {

  var query struct {
    CurrentCharacter types.CurrentCharacter
  }

  err := client.Query(context.Background(), &query, nil)
  if err != nil {
    fmt.Printf("%v\n", err)
    // Handle error.
    return
  }

  currentCharacter = query.CurrentCharacter
}

/**
 * Load location
 */
func getLocation() {

  var query struct {
    CurrentCharacter struct {
      Car struct {
        Location types.Location
      }
    }
  }

  err := client.Query(context.Background(), &query, nil)
  if err != nil {
    fmt.Printf("%v\n", err)
    // Handle error.
    return
  }

  currentCharacter.Car.Location = query.CurrentCharacter.Car.Location
}

/**
 * setup
 */
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
  token = conf.Characters[0].Token

  fmt.Printf("Character from config file has token: %v\n", token)

  // get character from server
  httpClient := http.Client{Transport: NewAddHeaderTransport(nil)}
  client = graphql.NewClient("http://"+host+":"+strconv.Itoa(port), &httpClient)

  getCurrentCharacter()
}

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

  for {
    time.Sleep(time.Second)
    getLocation()
    ui.Render(ui.RenderData{CurrentCharacter: currentCharacter})
  }
}
