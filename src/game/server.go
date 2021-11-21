package game

import (
	"io"
	"log"
	"net"
)

func Server() {
	l, err := net.Listen("tcp", ":8000") // 소캣열어줌
	if nil != err {
		log.Println(err)
	}
	defer l.Close() // 메인프로세스 종료시 소켓 종료

	for {
		conn, err := l.Accept()
		if nil != err {
			log.Println(err)
			continue
		}
		defer conn.Close()
		go ConnHandler(conn) // 헨들러 고루틴으로 돌림
	}
}

func ConnHandler(conn net.Conn) {
	recvBuf := make([]byte, 4096) // 4096 바이트로 데이터 전송하는데 가장 효율적이라고함
	for {
		n, err := conn.Read(recvBuf) // client가 값을 줄때까지 값읽어들임
		if nil != err {
			if io.EOF == err {
				log.Println(err)
				return
			}
			log.Println(err)
			return
		}
		if 0 < n { // 읽어들인 값이 있다면
			data := recvBuf[:n]           // 받은 데이터 byte배열에 저장
			log.Println(string(data))     // 출력
			_, err = conn.Write(data[:n]) // 연결되었던 client에게 write시킴
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
