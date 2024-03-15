package main

import (
	"flag"
	"github/suixinpr/manadb/internal/storage/smngr"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

var (
	directory string
	remove    bool
)

func init() {
	usr, err := user.Current()
	if err != nil {
		log.Panic(err)
	}
	flag.StringVar(&directory, "d", filepath.Join(usr.HomeDir, "manadata"), "database directory, the default path is user directory")
	flag.BoolVar(&remove, "r", false, "remove the directory when the database directory exists")
}

func main() {
	flag.Parse()

	if remove {
		exist, err := pathExists(directory)
		if err != nil {
			log.Panic(err)
		}
		if exist {
			os.RemoveAll(directory)
		}
	}

	err := os.Mkdir(directory, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
	smngr.Init(directory)
	initMetadata()
}

func pathExists(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
