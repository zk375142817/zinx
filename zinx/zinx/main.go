package main

import "zinxMmo/zinx/znet"

func main() {
	s := znet.NewServer("s")
	s.Serve()
}
