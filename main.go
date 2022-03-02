
include "gutil"
include "vars"
include "window"

func TerminInit () () {
	InitGu()
	clear()
	InitGetCh()

}

func TerminEnd () () {
	for i:=0;i<len(wins);i++{
		end(wins[i])
	}
}
