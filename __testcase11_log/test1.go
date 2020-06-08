package main

import "log"

func main() {
	log.Println(123)
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Printf("afdf")

	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)
	log.Printf("dddd")

	log.SetPrefix("【print log:】")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	log.Printf("999")

	log.Panic("8888")
}
