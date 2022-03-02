
var (
	wins = []*Window{}
	termy, termx = GetTerminalSize()
	Win = MakeWin("Main", stdout, 0, termy, 0, termx)
	nw = &Window{}
)
