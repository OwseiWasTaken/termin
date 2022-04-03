
var (
	termy, termx = GetTerminalSize()
	wins = []*Window{}
	Win = MakeWin("Main", stdout, 0, termy, 0, termx)
	nw = &Window{} // NoWin (used in end())
	lk []byte
)

