
var (
	termy, termx = GetTerminalSize()
	wins = []*Window{}
	Win = MakeWin("Main", stdout, stdin, 0, termy, 0, termx)
	nw = &Window{} // NoWin (used in end())
	lk []byte
)

