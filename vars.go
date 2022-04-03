
var (
	termy, termx = GetTerminalSize()
	wins = []*Window{}
	Win = MakeWin("Main", stdout, 0, termy, 0, termx)
	nw = &Window{} // NoWin (used in end())
	lk []byte
)

func Compress( x []byte ) ( string ) {
	buff := ""
	for i:=0;i!=6;i++{
		if (x[i] == 0) { break }
		buff+=spf("%.3d,", x[i])
		lk = append(lk, x[i])
	}
	return buff
}

