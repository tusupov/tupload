package main

import (
	"net/http"
	"handle"
	"fmt"
	"flag"
)

var (
	argv struct {
		port string
		rootPath string
	}
)

func initArgv() {

	flag.StringVar(&argv.port, "port", "8080", "port")
	flag.StringVar(&argv.rootPath, "path", "/var/www/", "path to project")

}

func main()  {

	initArgv()

	//image upload handle
	http.HandleFunc("/image", func (w http.ResponseWriter, r *http.Request) {
		handle.UploadImage(w, r, argv.rootPath)
	})

	//images public folder
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	//Start web server
	err := http.ListenAndServe(":" + argv.port, nil)
	if err != nil {
		fmt.Println(err)
	}

}
