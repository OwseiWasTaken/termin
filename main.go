include "gutil"
include "termin.vars"
include "termin.control"
include "termin.window"

func InitTermin() () {
	InitGu()
	clear()
	InitGetCh()
}

func StopTermin() () {
	for i:=0;i<len(wins);i++{
		end(wins[i])
	}
}

func EndTermin() () {
	for i:=0;i<len(wins);i++{
		end(wins[i])
	}
}
