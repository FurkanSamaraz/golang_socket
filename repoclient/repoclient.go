package repoclient

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func Createclientsocket(connHost, connPort, connType, index string) {
	// İstemciyi başlatın ve sunucuya bağlanın.
	fmt.Println("Connecting to "+index+"", connType, "server", connHost+":"+connPort)
	conn, err := net.Dial(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	// Stdin'den yeni okuyucu oluşturun.
	reader := bufio.NewReader(os.Stdin)

	// çıkışa kadar döngüyü sonsuza kadar çalıştırın.
	for {
		// Yönlendirme mesajı.
		fmt.Print("Text to send-" + index + ": ")

		// Yeni satıra kadar girişi okuyun, Enter tuşu.
		input, _ := reader.ReadString('\n')

		// Soket bağlantısına gönder.
		conn.Write([]byte(input))

		//  relay dinle.
		message, _ := bufio.NewReader(conn).ReadString('\n')

		// relay yazdır.
		log.Print("Server relay-" + index + " : " + message)
	}
}
