type Ordenate struct {
	a int // y
	b int // x
}

type Window struct {
	name string
	min Ordenate
	max Ordenate
	stream *bufio.Writer
}

func MakeWin( name string, stream *bufio.Writer, MinY, MaxY, MinX, MaxX int ) ( Window ) {
	w := Window{name, Ordenate{MinY, MinX}, Ordenate{MaxY, MaxX}, stream}
	wins = append(wins, &w)
	return w
}

func read(w Window) ([]byte) {
	w.stream.Flush()
	return GetChByte()
}

func end(w *Window) () {
	w.stream.Flush()
	w = nw
	runtime.GC()
}

func wprint (w Window, y int, x int, stuff string) () {
	panic( _wprint(w, y, x, stuff))
}

func _wprint (w Window, y int, x int, stuff string) (error) {
	var ny = y+w.min.b
	var nx = x+w.min.a
	if (ny <= w.max.a && nx <= w.max.b && nx >= w.min.b-1 && ny >= w.min.a-1 && nx+len(stuff) <= w.max.b) {
		w.stream.Write([]byte(spos(ny, nx)+stuff))
		return nil
	} else {
		return errors.New(
			spf("tried to write out of window `%s`s bounds\nx(max:%d, input:%d(%d+len(msg)))\ny(max:%d, input:%d)",
			w.name, w.max.b, nx+len(stuff)-1, nx, w.max.a, ny,
		))
	}
}

func wputc(w Window, y int, x int, stuff byte) ( error ) {
	var ny = y+w.min.b
	var nx = x+w.min.a
	if (ny < w.max.b && nx < w.max.a && nx > w.min.a-1 && ny > w.min.b-1) {
		w.stream.Write(append([]byte(spos(ny, nx)), stuff))
		return nil
	} else {
		return errors.New(spf("tried to write out of window `%s`s bounds", w.name))
	}
}

func wuputc(w Window, y int, x int, stuff byte) () {
	w.stream.Write(append([]byte(spos(y, x)), stuff))
}

func wmove(w Window, y int, x int) () {
	var ny = y+w.min.b
	var nx = x+w.min.a
	if (ny < w.max.b && nx < w.max.a && nx > w.min.a-1 && ny > w.min.b-1) {
		w.stream.Write([]byte(spos(y, x)))
	} else {
		panic(errors.New(spf("tried to wmove out of window `%s`s bounds", w.name)))
	}
}

func wumove(w Window, y int, x int) () {
	w.stream.Write([]byte(spos(y, x)))
}
