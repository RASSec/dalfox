package printing

import (
	"fmt"
	"os"
	"sync"

	"github.com/hahwul/dalfox/pkg/model"
	"github.com/logrusorgru/aurora"
)

var (
	mutex = &sync.Mutex{}
)

// DalLog is log fomatting for DalFox
func DalLog(level, text string, options model.Options) {
	var ftext string
	if level == "INFO" {
		if options.Debug {
			ftext = "[I] " + text
		}
		text = aurora.Blue("[I] ").String() + text

	}
	if level == "WEAK" {
		if options.Debug {
			ftext = "[W] " + text
		}
		text = aurora.Yellow("[W] ").String() + text

	}
	if level == "VULN" {
		if options.Debug {
			ftext = "[V] " + text
		}
		text = aurora.Red("[V] ").String() + text

	}
	if level == "SYSTEM" {
		if options.Debug {
			ftext = "[*] " + text
		}
		text = aurora.White("[*] ").String() + text

	}
	if level == "GREP" {
		if options.Debug {
			ftext = "[G] " + text
		}
		text = aurora.Green("[G] ").String() + text

	}

	if level == "CODE" {
		if options.Debug {
			ftext = "    " + text
		}
		text = aurora.Gray(16-1, "    "+text).String()
	}

	if level == "ERROR" {
		if options.Debug {
			ftext = "[E] " + text
		}
		text = aurora.Yellow("[E] ").String() + text
	}

	if level == "YELLOW" {
		text = aurora.BrightYellow(text).String()
	}

	//mutex.Lock()
	if options.Silence == true {
		if level == "PRINT" {
			if options.Format == "json" {
				ftext = text
				//fmt.Println(aurora.BrightGreen(text))
				fmt.Println(text)

			} else {
				ftext = "[POC] " + text
				if options.NoColor {
					fmt.Println("[POC]" + text)
				} else {
					fmt.Println(aurora.BrightGreen("[POC]" + text))
				}
			}
		}
	} else {
		if level == "PRINT" {
			if options.Format == "json" {
				ftext = text
				//fmt.Println(aurora.BrightGreen(text))
				fmt.Println(text)

			} else {
				ftext = "[POC] " + text
				if options.NoColor {
					fmt.Println("[POC]" + text)
				} else {
					fmt.Println(aurora.BrightGreen("[POC]" + text))
				}
			}
		} else {
			if options.NoColor {
				text = "\r" + ftext
			} else {
				text = "\r" + text
			}
			fmt.Fprintln(os.Stderr, text)
		}
	}

	if options.OutputFile != "" {
		var fdtext string
		if ftext != "" {
			fdtext = ftext
			f, err := os.OpenFile(options.OutputFile,
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Fprintln(os.Stderr, "output file error (file)")
			}
			defer f.Close()
			if _, err := f.WriteString(fdtext + "\n"); err != nil {
				fmt.Fprintln(os.Stderr, "output file error (write)")
			}
		}
	}
}
