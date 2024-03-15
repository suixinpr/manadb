package main

import (
	"flag"
	"fmt"
	"github/suixinpr/manadb/internal/server"
	"github/suixinpr/manadb/internal/storage/smngr"
	"log"
	"os/user"
	"path/filepath"
)

var (
	host      string
	port      int
	directory string
)

func init() {
	usr, err := user.Current()
	if err != nil {
		log.Panic(err)
	}
	flag.StringVar(&host, "h", "127.0.0.1", "host, default is 127.0.0.1")
	flag.IntVar(&port, "p", 9315, "port number, default is 9315")
	flag.StringVar(&directory, "d", filepath.Join(usr.HomeDir, "manadata"), "database directory, the default path is user directory")
}

func main() {
	flag.Parse()
	smngr.Init(directory)
	server.ListenAndServe(fmt.Sprintf("%s:%d", host, port))
}
