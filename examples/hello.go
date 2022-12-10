package main

Include "termin"

func main(){
	InitTermin()

	wprint(Win, 0, 0, "hello") // to the main window, at 0,0, write "hello"
	wgtk(Win) // get key, but don't store output

	StopTermin()
	exit(0)
}
