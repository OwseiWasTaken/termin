package main

Include "termin"

func main(){
	InitTermin()
	var k string
	for ;k != "pause/break";{
		k = wgtk(Win)
		wwrite(Win, spf("%v:%s\n", lk, k))
	}
	StopTermin()
	exit(0)
}
