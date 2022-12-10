package main

Include "termin"

func main(){
	InitTermin()
	var k string
	for ;k != "pause/break";{
		k = wgtk(Win) // wait for keyboard key
		// wgtk sets lk (last key []byte) and returns the key's name
		wwrite(Win, spf("%v:%s\n", lk, k))
	}
	StopTermin()
	exit(0)
}
