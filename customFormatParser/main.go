package main

import (
	"fmt"
	"os/user"
	"strings"
	"time"
)

// "[%{module}] [%{file} - %{line}] [%{level}] %{message}"
var (
	// Default format of log message
	defFmt = "#%[1]d %[2]s %[4]s:%[5]d ▶ %.3[6]s %[7]s"

	// Default format of time
	defTimeFmt = "2006-01-02 15:04:05"

	// Map from format's placeholders to printf verbs
	phfs map[string]string
)

// init pkg
func init() {
	initFormatPlaceholders()
}

// Analyze and represent format string as printf format string and time format
func parseVersioningFormat(format string) (msgfmt, timefmt string) {
	if len(format) < 10 /* (len of "%{message} */ {
		fmt.Println("len(format) ", len(format))
		return defFmt, defTimeFmt
	}
	fmt.Println("off we go")
	timefmt = defTimeFmt
	idx := strings.IndexRune(format, '%') //find the index of the (1st??) % symbol
	for idx != -1 {
		msgfmt += format[:idx] //msgfmt is the format up to the first %
		format = format[idx:]  //format becomes everything after the first %
		if len(format) > 2 {   //check that format is now longer than 2 characters
			if format[1] == '{' { //check if the second character is a {}
				// end of curr verb pos
				if jdx := strings.IndexRune(format, '}'); jdx != -1 { //if its the end of the current 'verb'
					// next verb pos
					idx = strings.Index(format[1:], "%{") //update the index to the next verb start position
					// incorrect verb found ("...%{wefwef ...") but after
					// this, new verb (maybe) exists ("...%{inv %{verb}...")
					if idx != -1 && idx < jdx {
						msgfmt += "%%"
						format = format[1:]
						continue
					}
					// get verb and arg
					verb, arg := ph2verb(format[:jdx+1])
					msgfmt += verb
					// check if verb is time
					// here you can handle args for other verbs
					if verb == `%[2]s` && arg != "" /* %{time} */ {
						timefmt = arg
					}
					format = format[jdx+1:]
				} else {
					format = format[1:]
				}
			} else {
				msgfmt += "%%"
				format = format[1:]
			}
		}
		idx = strings.IndexRune(format, '%')
	}
	msgfmt += format
	return
}

// translate format placeholder to printf verb and some argument of placeholder
// (now used only as time format)
func ph2verb(ph string) (verb string, arg string) {
	n := len(ph)
	if n < 4 {
		return ``, ``
	}
	if ph[0] != '%' || ph[1] != '{' || ph[n-1] != '}' {
		return ``, ``
	}
	idx := strings.IndexRune(ph, ':')
	if idx == -1 {
		return phfs[ph], ``
	}
	verb = phfs[ph[:idx]+"}"]
	arg = ph[idx+1 : n-1]
	fmt.Printf("verb: %s, arg: %s", verb, arg)
	return
}

// Initializes the map of placeholders
//"[%{bigVersion}] [%{littleVersion} - %{microVersion}] [%{time}] %{message}"
func initFormatPlaceholders() {
	phfs = map[string]string{
		"%{bigVersion}":    "%[1]d",
		"%{littleVersion}": "%[2]d",
		"%{microVersion}":  "%[3]d",
		"%{time}":          "%[4]s",
		"%{timeUnix}":      "%[4]d",
		"%{client}":        "%[5]s",
		"%{job}":           "%[6]s",
		"%{creator}":       "%[7]s",
		"%{owner}":         "%[8]s",
		"%{hash}":          "%[9]s",
		"%{message}":       "%[10]s",
	}
}

func main() {

	msgfmt, timefmt := parseVersioningFormat("%{bigVersion}.%{littleVersion}.%{microVersion}_%{timeUnix}_%{client}_%{job}_%{creator}_%{owner}_%{hash}_%{message}")
	fmt.Printf("msgfmt: %s, timefmt: %s\r\n", msgfmt, timefmt)
	bigVersion := 1
	littleVersion := 2
	microVersion := 3
	currentTime := time.Now()
	client := "ACME Corp"
	job := "Sylvestor's hammer"
	user, _ := user.Current()
	userId := user.Name + user.Uid
	owner := "Sylverstor"
	hash := "0xDEADBEEF"
	message := "cultural adjustment"
	msg := fmt.Sprintf(msgfmt,
		bigVersion,         // %[1] // %{bigVersion}
		littleVersion,      // %[2] // %{littleVersion}
		microVersion,       // %[3] // %{microVersion}
		currentTime.Unix(), // %[4] // %{timeUnix} // -> currentTime, // %[4] // %{time}
		client,             // %[5] // %{client}
		job,                // %[6] // %{job}
		userId,             // %[7] // %{creator}
		owner,              // %[8] // %{owner}
		hash,               // %[9] // %{hash}
		message,            // %[19] // %{message}
	)
	fmt.Println("msg: ", msg)

	for i, _ := range phfs {
		fmt.Println(i)
	}
}
