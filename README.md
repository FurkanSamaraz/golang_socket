# golang_socket

# SERVER


"""
package main

import (
	"main/repoclient"
)

var wait chan string

func main() {
	go repoclient.Createclientsocket("localhost", "8080", "tcp", "1")
	go repoclient.Createclientsocket("localhost", "8000", "tcp", "2")
	<-wait
}
"""
