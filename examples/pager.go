package main

Include "termin"

const (
	C_inverted = "\x1b[38;2;0;0;0;48;2;255;255;255m"
	C_normal  =  "\x1b[38;2;255;255;255;48;2;0;0;0m"
)

var (
	cont []string
	fmax int
)

func Draw (y int) {
	for i:=0; i<Win.LenY-1; i++ {
		wprint(Win, i, 0, "\x1b[K")
		wwrite(Win, cont[y+i])
	}
	wprint(Win, Win.LenY, 0, "\x1b[K")
	if y == len(cont)-Win.LenY {
		wwrite(Win, C_inverted+"(END)"+C_normal)
	} else {
		wwrite(Win, spf("%s:%d/%d:", Win.name, y, len(cont)-Win.LenY))
	}
	wflush(Win)
}

// read file, update globals
func ReadAndUpdate (filename string) () {
	cont = strings.Split(filename, "/")
	Win.name = cont[len(cont)-1]

	// replace tab with 2-width space, split by \n
	cont = strings.Split( strings.Replace(
		ReadFile(filename), "\t", " "+" ",
	-1), "\n")

	fmax = len(cont)-Win.LenY
	if (fmax > 0) {
		return
	}
	if (true) { // less style
		for i:=0;i>fmax;i-- {
			cont = append(cont, "~")
		}
	} else { // shrink style
		fmax = 0
		Win.LenY = len(cont)
	}
}

func UseKey(k string, y int) (int) {
	switch (k) {
	case "down", "j":
		if y < fmax {
			y++
			Draw(y)
		}
	case "up", "k":
		if y > 0 {
			y--
			Draw(y)
		}
	case "g":
		y = 0
		Draw(y)
	case "G":
		y = fmax
		Draw(y)
	case "q":
		wclear(Win)
		exit(0)
	}
	return y
}

func main(){
	InitTermin()
	if argc == 0 {
		fprintf(stderr, "No file to read\n")
		exit(1)
	}
	var (
		file string
		y int = 0
		k string
	)

	file = argv[0]
	ReadAndUpdate(file)

	Draw(y)
	for {
		k = wgtk(Win)
		y = UseKey(k, y)
	}

	StopTermin()
	exit(0)
}

