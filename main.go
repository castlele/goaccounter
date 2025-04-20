package main

import (
	"github.com/castlele/goaccounter/src/storage"
	"github.com/castlele/goaccounter/src/storage/database"
)


var db storage.Storage

func main() {
    db = &database.LocalDB{}
}
