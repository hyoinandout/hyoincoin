package main

import (
	"github.com/hyoinandout/hyoincoin/cli"
	"github.com/hyoinandout/hyoincoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
