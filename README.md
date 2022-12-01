# golang_socket

# SERVER


```
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
```

# CLIENT
```
package main

import (
	"main/reposerver"
)

var wait chan string

func main() {
	go reposerver.Createserversocket("localhost", "8080", "tcp", "1")
	go reposerver.Createserversocket("localhost", "8000", "tcp", "2")
	<-wait
}
```

