package reposerver

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func Createserversocket(connHost, connPort, connType, index string) {
	// Sunucuyu başlatın ve gelen bağlantıları dinleyin.
	fmt.Println("Starting " + index + "" + connType + " server on " + connHost + ":" + connPort)
	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Uygulama kapandığında dinleyiciyi kapatın.
	defer l.Close()

	// çıkışa kadar döngüyü sonsuza kadar çalıştırın.
	for {
		// Gelen bir bağlantıyı dinleyin.
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		fmt.Println("Client connected.")

		// İstemci bağlantı adresini yazdırın.
		fmt.Println("Client " + index + "" + c.RemoteAddr().String() + " connected.")

		// Bağlantıları aynı anda yeni bir programda yönetin.
		go handleConnection(c, index)
	}
}

// handleConnection, tek bir bağlantı isteği için mantığı işler.
func handleConnection(conn net.Conn, index string) {
	// İstemci girişini yeni bir satıra kadar tamponlayın.
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	// Sol istemcileri kapatın.
	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}

	// Yeni satır karakterini sıyırarak yanıt mesajını yazdırın.
	log.Println("Client message- "+index+":", string(buffer[:len(buffer)-1]))

	// İstemciye yanıt mesajı gönderin.
	conn.Write(buffer)

	//İşlemi yeniden başlatın.
	handleConnection(conn, index)
}
