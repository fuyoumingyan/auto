package initpkg

import "fmt"

func initBanner() {
	var banner = `
   ▄▄▄       █    ██ ▄▄▄█████▓ ▒█████  
  ▒████▄     ██  ▓██▒▓  ██▒ ▓▒▒██▒  ██▒
  ▒██  ▀█▄  ▓██  ▒██░▒ ▓██░ ▒░▒██░  ██▒
  ░██▄▄▄▄██ ▓▓█  ░██░░ ▓██▓ ░ ▒██   ██░
   ▓█   ▓██▒▒▒█████▓   ▒██▒ ░ ░ ████▓▒░
   ▒▒   ▓▒█░░▒▓▒ ▒ ▒   ▒ ░░   ░ ▒░▒░▒░ 
    ▒   ▒▒ ░░░▒░ ░ ░     ░      ░ ▒ ▒░ 
    ░   ▒    ░░░ ░ ░   ░      ░ ░ ░ ▒  
        ░  ░   ░                  ░ ░

`
	fmt.Println(banner)
}
