package game

import (
	"fmt"
	"log"
	"net"
)

func Client() {
	conn, err := net.Dial("tcp", ":8000") // 서버와 연결시도
	if nil != err {
		log.Println(err)
	}

	go func() { // 고루틴 생성하고 바로 실행
		data := make([]byte, 4096) // byte 배열 만듬
		//data2 := FSend

		for {
			n, err := conn.Read(data) // 커낵션에서 데이터 Read
			if err != nil {
				log.Println(err)
				return
			}

			log.Println("Server send : " + string(data[:n])) // read 한 데이터 출력
		}
	}()

	for {
		var s string
		fmt.Scanln(&s)        // 프롬포트 입력 스캔
		conn.Write([]byte(s)) // 커낵션에 Write
	}
}

func ClientInit() {
	fmt.Println("Client.go GOGO")
}
