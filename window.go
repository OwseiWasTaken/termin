type Ordenate struct {
	y int
	x int
}

type Window struct {
	name string
	min Ordenate
	max Ordenate
	stream *bufio.Writer
}

func MakeWin( name string, stream *bufio.Writer, MinY, MaxY, MinX, MaxX int ) ( *Window ) {
	w, e :=_MakeWin(name, stream, MinY, MaxY, MinX, MaxX)
	panic(e)
	wins = append(wins, w)
	return w
}

func _MakeWin( name string, stream *bufio.Writer, MinY, MaxY, MinX, MaxX int ) ( *Window, error ) {
	if (MaxX > termx || MaxY > termy) {
		return nil, errors.New(spf(
"can't make window `%s`, MaxX ($5:%d) or MaxY ($3:%s) is bigger than terminal size (termx: %d, termy: %d)",
name, MaxX, MaxY, termx, termy,
		))
	}
	w := &Window{name, Ordenate{MinY, MinX}, Ordenate{MaxY, MaxX}, stream}
	return w, nil
}

func read(w *Window) ([]byte) {
	w.stream.Flush()
	return GetChByte()
}

func end(w *Window) () {
	w.stream.Flush()
	w = nw
	runtime.GC()
}

func wprint (w *Window, y int, x int, stuff string) () {
	panic( _wprint(w, y, x, stuff))
}

func _wprint (w *Window, y int, x int, stuff string) (error) {
	var ny = y+w.min.y
	var nx = x+w.min.x
	if (ny <= w.max.y && nx <= w.max.x && nx >= w.min.x-1 && ny >= w.min.y-1 && nx+len(stuff) <= w.max.x) {
		w.stream.WriteString(spos(ny, nx)+stuff)
		return nil
	} else {
		return errors.New(
			spf("tried to write out of window `%s`s bounds\nx(max:%d, input:%d(%d+len(msg)))\ny(max:%d, input:%d)",
			w.name, w.max.x, nx+len(stuff)-1, nx, w.max.y, ny,
		))
	}
}

// window's unsafe print
func wuprint ( w *Window, y int, x int, stuff string ) () {
	var ny = y+w.min.y
	var nx = x+w.min.x
	if (ny <= w.max.y && nx <= w.max.x && nx >= w.min.x-1 && ny >= w.min.y-1) {
		w.stream.WriteString(spos(ny, nx)+stuff)
	} else {
		panic( errors.New(
			spf("tried to write out of window `%s`s bounds\nx(max:%d, input:%d(%d+len(msg)))\ny(max:%d, input:%d)",
			w.name, w.max.x, nx+len(stuff)-1, nx, w.max.y, ny,
		)))
	}
}

func wputc(w *Window, y int, x int, stuff string) ( error ) {
	var ny = y+w.min.y
	var nx = x+w.min.x
	if (ny < w.max.y && nx < w.max.x && nx > w.min.x-1 && ny > w.min.y-1) {
		w.stream.WriteString(spos(ny, nx)+stuff)
		return nil
	} else {
		return errors.New(spf("tried to write out of window `%s`s bounds", w.name))
	}
}

func wuputc(w *Window, y int, x int, stuff string) () {
	if (len(stuff)==1) {
		w.stream.WriteString(spos(y, x)+stuff)
	} else {
		panic(errors.New(spf("func wuputc char to write's len != 1")))
	}
}

func wmove(w *Window, y int, x int) () {
	var ny = y+w.min.y
	var nx = x+w.min.x
	if (ny < w.max.y && nx < w.max.x && nx > w.min.x-1 && ny > w.min.y-1) {
		w.stream.WriteString(spos(y, x))
	} else {
		clear()
		stdout.Flush()
		stderr.Flush()
		panic(errors.New(
			spf(
				"tried to wmove out of window `%s`s bounds\nx(max:%d, input:%d)\ny(max:%d, input:%d)",
				w.name, w.max.x, nx, w.max.y, ny,
			),
		))
	}
}

func wumove(w *Window, y int, x int) () {
	w.stream.Write([]byte(spos(y, x)))
}

func wDrawLine(w *Window, y int, char string) () {
	w.stream.WriteString(spos(y, 0) + strings.Repeat(char, w.max.x+1))
}

func wDrawCollum( w *Window, x int, char string ) () {
	var c string
	if len(char) == 1 {
		c=string(char[0])
	} else {
		clear()
		w.stream.Flush()
		panic(errors.New(spf("func wDrawCollum char must be len 1 (len(char) -> %d)\n", len(char))))
	}
	for i:=w.min.y;i<w.max.y;i++ {
		w.stream.WriteString(spos(i, x)+c)
	}
}

func Compress( x []byte ) ( string ) {
	buff := ""
	for i:=0;i!=6;i++{
		if (x[i] == 0) { break }
		buff+=spf("%.3d,", x[i])
		lk = append(lk, x[i])
	}
	return buff
}

func wgtk ( w *Window ) ( string ) {
	x:=read(w)
	lk = []byte{}
	e, ok := Control[Compress(x)]
	if (!ok) {
		// if this happens, update https://github.com/owseiwastaken/termin
		//	   get the key u pressed and get Compress' result ( e.g. [103,0,0 ...] -> "103," )
		//	   add the key to the HashMap, and it's name ( e.g. "103,":"g", "32":"space",)
		// if you don't want to do this just send me a msg ow#2183 (discord)

		// btw, lk saves the byte array of the key (on func Compress)
		return "NULL"
	}
	return e
}

