package main

func main() {
	client := &TCPClient{
		Host: "127.0.0.1",
		Port: 8888,
	}
	client.Start()
}
