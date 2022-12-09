type Window struct {
	name string
	MinX int
	MaxX int
	MinY int
	MaxY int
	LenX int
	LenY int
	stream *bufio.Writer
	instream *bufio.Reader
}

func MakeWin( name string, stream *bufio.Writer, instream *bufio.Reader, MinY, MaxY, MinX, MaxX int ) ( *Window ) {
	w, e :=_MakeWin(name, stream, instream, MinY, MaxY, MinX, MaxX)
	panic(e)
	wins = append(wins, w)
	return w
}

func _MakeWin( name string, stream *bufio.Writer, instream *bufio.Reader, MinY, MaxY, MinX, MaxX int ) ( *Window, error ) {
	if (MaxX > termx || MaxY > termy) {
		return nil, errors.New(spf(
"can't make window `%s`, MaxX ($5:%d) or MaxY ($3:%d) is bigger than terminal size (termx: %d, termy: %d)",
name, MaxX, MaxY, termx, termy,
		))
	}
	w := &Window{name, MinX, MaxX, MinY, MaxY, MaxX-MinX, MaxY-MinY, stream, instream}
	return w, nil
}

func wread(w *Window) ([]byte) {
	w.stream.Flush()
	return GetChByte(w.instream)
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
	var ny = y+w.MinY
	var nx = x+w.MinX
	if (ny <= w.MaxY && nx <= w.MaxX && nx >= w.MinX-1 && ny >= w.MinY-1 && nx+len(stuff) <= w.MaxX) {
		w.stream.WriteString(spos(ny, nx)+stuff)
		return nil
	} else {
		return errors.New(
			spf("tried to write out of window %s`s bounds\nx(max:%d, input:%d(%d+len(msg)))\ny(max:%d, input:%d)",
			w.name, w.MaxX, nx+len(stuff)-1, nx, w.MaxY, ny,
		))
	}
}

// window's unsafe print
func wuprint ( w *Window, y int, x int, stuff string ) () {
	var ny = y+w.MinY
	var nx = x+w.MinX
	if (ny <= w.MaxY && nx <= w.MaxX && nx >= w.MinX-1 && ny >= w.MinY-1) {
		w.stream.WriteString(spos(ny, nx)+stuff)
	} else {
		panic( errors.New(
			spf("tried to write out of window `%s`s bounds\nx(max:%d, input:%d(%d+len(msg)))\ny(max:%d, input:%d)",
			w.name, w.MaxX, nx+len(stuff)-1, nx, w.MaxY, ny,
		)))
	}
}

func wputc(w *Window, y int, x int, stuff rune) ( error ) {
	var ny = y+w.MinY
	var nx = x+w.MinX
	if (ny < w.MaxY && nx < w.MaxX && nx > w.MinX-1 && ny > w.MinY-1) {
		w.stream.WriteString(spos(ny, nx)+string(stuff))
		return nil
	} else {
		return errors.New(spf("tried to write out of window `%s`s bounds", w.name))
	}
}

func wuputc(w *Window, y int, x int, stuff rune) () {
	w.stream.WriteString(spos(y, x)+string(stuff))
}

func wmove(w *Window, y int, x int) () {
	var ny = y+w.MinY
	var nx = x+w.MinX
	if (ny < w.MaxY && nx < w.MaxX && nx > w.MinX-1 && ny > w.MinY-1) {
		w.stream.WriteString(spos(y, x))
	} else {
		clear()
		stdout.Flush()
		stderr.Flush()
		panic(errors.New(
			spf(
				"tried to wmove out of window `%s`s bounds\nx(max:%d, input:%d)\ny(max:%d, input:%d)",
				w.name, w.MaxX, nx, w.MaxY, ny,
			),
		))
	}
}

func wumove(w *Window, y int, x int) () {
	w.stream.WriteString(spos(w.MinY+y, w.MinX+x))
}

func wDrawLine(w *Window, y int, char rune) () {
	if y == -1 {
		y = w.MaxY-w.MinY
	}
	if w.MaxY < (y+w.MinY) || y < 0 {
		panic(errors.New(
			spf(
				"wDrawLine's (%d) y can't be bigger than MaxY (%d) nor smaller than 0",
				y, w.MaxY,
			),
		))
	}
	w.stream.WriteString(spos(w.MinY + y, w.MinX) + strings.Repeat(string(char), w.LenX))
}

func wDrawCollum( w *Window, x int, char rune) () {
	if x == -1 {
		x = w.MaxX-w.MinX
	}
	if w.MaxX < (x+w.MinX) || x < 0 {
		panic(errors.New(
			spf(
				"wDrawCollum's (%d) x can't be bigger than MaxX (%d) nor smaller than 0",
				x, w.MaxX,
			),
		))
	}
	for i:=w.MinY;i<w.MaxY;i++ {
		w.stream.WriteString(spos(i, w.MinX+x)+string(char))
	}
}

//TODO(3) unsafe sized: make 'Sized' funcs safe
func wDrawSizedLine( w *Window, y int, s int, char rune ) () {
	w.stream.WriteString(spos(w.MinY + y, w.MinX) + strings.Repeat(string(char), s))
}

func wDrawSizedCollum( w *Window, x int, s int, char rune ) () {
	for i:=0;i<s;i++ {
		w.stream.WriteString(spos(w.MinY+i, w.MinX+x)+string(char))
	}
}

func Compress( x []byte ) ( string ) {
	buff := ""
	for i:=0;i!=6;i++{
		if (x[i] == 0) { break }
		buff+=spf("%.3d,", x[i])
		//lk = append(lk, x[i])
	}
	return buff
}

func wgtk ( w *Window ) ( string ) {
	lk = wread(w)
	e, ok := Control[Compress(lk)]
	if (!ok) {
		// if this happens, update https://github.com/owseiwastaken/termin
		//	   get the key u pressed and get Compress' result ( e.g. [103,0,0 ...] -> "103," )
		//	   add the key to the HashMap, and it's name ( e.g. "103,":"g", "32":"space",)
		// if you don't want to do this just send me a msg ow#2183 (discord)

		// btw, lk saves the byte array of the key
		// tho, if u only want the byte array, consider using wread
		return "NULL"
	}
	return e
}

func wDrawBorder ( w *Window, char rune) () {
	wDrawCollum(w, 0 , char)
	wDrawCollum(w, -1, char)
	wDrawLine(	w, 0 , char)
	wDrawLine(	w, -1, char)
}

func wDrawBorderName ( w *Window, char rune ) () {
	wDrawCollum(w, 0 , char)
	wDrawCollum(w, -1, char)
	wDrawLine(	w, 0 , char)
	wDrawLine(	w, -1, char)
	wprint(		w, 0 , 1, " " + w.name + " ")
}

func wClear ( w *Window ) () {
	for i:=w.MinY;i<w.MaxY;i++ {
		wDrawLine(w, i, ' ')
	}
	w.stream.Flush()
	wmove(w, 0, 0)
}

func wFlush( w *Window ) () {
	w.stream.Flush()
}

func wWrite ( w *Window, s string ) () {
	w.stream.WriteString(s)
}

func wPut ( w *Window, s rune) () {
	w.stream.WriteByte(byte(s))
}

func wPutb ( w *Window, s byte) () {
	w.stream.WriteByte(s)
}

func wColor ( s string ) () {
	stdout.WriteString(s)
}
