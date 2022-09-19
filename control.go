
//mapping

var Control = map[string]string{
	"":"NULL",

	"001,": "^A",
	"002,": "^B",
	"003,": "^C", // disable ^C exit
	"004,": "^D",
	"005,": "^E",
	"006,": "^F",
	"007,": "^G",
	"008,": "^H",
	//"009,": "^I", // tab
	//"010,": "^J", // enter if keyboard.go is off
	"011,": "^K",
	"012,": "^L",
	"013,": "^M", // enter if keyboard.go is on
	"014,": "^N",
	"015,": "^O",
	"016,": "^P",
	"017,": "^Q",
	"018,": "^R",
	"019,": "^S",
	"020,": "^T",
	"021,": "^U",
	"022,": "^V",
	"023,": "^W",
	"024,": "^X",
	"025,": "^Y",
	"026,": "^Z",

	// move
	"027,091,065,":"up",
	"027,091,066,":"down",
	"027,091,068,":"left",
	"027,091,067,":"right",

	// special
	"127,":"backspace",
	"027,091,051,126,":"delete",

	//alt+(
	"027,040,":"()",
	"027,091,":"[]",
	"027,123,":"{}",

	//func
	"027,079,080,":"f1",
	"027,079,081,":"f2",
	"027,079,082,":"f3",
	"027,079,083,":"f4",
	"027,091,049,053,126,":"f5",
	"027,091,049,055,126,":"f6",
	"027,091,049,056,126,":"f7",
	"027,091,049,057,126,":"f8",
	"027,091,050,048,126,":"f9",
	"027,091,050,049,126,":"f10",


	"010,":"enter",
	"009,":"tab",
	"027,":"esc",
	"032,":"space", "033,":"!", "034,":"\"", "035,":"#",
	"036,":"$", "037,":"%", "038,":"&", "039,":"'",
	"040,":"(", "041,":")", "042,":"*", "043,":"+",
	"044,":",", "045,":"-", "046,":".", "047,":"/",
	"048,":"0", "049,":"1", "050,":"2", "051,":"3", "052,":"4",
	"053,":"5", "054,":"6", "055,":"7", "056,":"8", "057,":"9",
	"058,":":", "059,":";",
	"060,":"<", "061,":"=", "062,":">",
	"063,":"?", "064,":"@",

	"065,":"A", "066,":"B", "067,":"C", "068,":"D",
	"069,":"E", "070,":"F", "071,":"G", "072,":"H",
	"073,":"I", "074,":"J", "075,":"K", "076,":"L",
	"077,":"M", "078,":"N", "079,":"O", "080,":"P",
	"081,":"Q", "082,":"R", "083,":"S", "084,":"T",
	"085,":"U", "086,":"V", "087,":"W", "088,":"X",
	"089,":"Y", "090,":"Z",

	"091,":"[", "092,":"\\", "093,":"]",
	"094,":"^", "095,":"_", "096,":"`",

	"097,":"a", "098,":"b", "099,":"c", "100,":"d",
	"101,":"e", "102,":"f", "103,":"g", "104,":"h",
	"105,":"i", "106,":"j", "107,":"k", "108,":"l",
	"109,":"m", "110,":"n", "111,":"o", "112,":"p",
	"113,":"q", "114,":"r", "115,":"s", "116,":"t",
	"117,":"u", "118,":"v", "119,":"w", "120,":"x",
	"121,":"y", "122,":"z",
}
