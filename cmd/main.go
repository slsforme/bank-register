package main

import "bankapp/config"

func init() {
	config.InitLogger()
	config.LoadEnv()
}

func main() {

	// cfg := config.Load()
}
