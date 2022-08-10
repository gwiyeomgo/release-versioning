package main

import (
	"flag"
	"gihub.com/gwiyeomgo/release-versioning/db"
	"gihub.com/gwiyeomgo/release-versioning/rest"
)

func main() {
	defer db.Close()
	port := flag.Int("port", 3000, "Set port of the server")
	flag.Parse()
	db.InitDB()
	rest.Start(*port)

}
