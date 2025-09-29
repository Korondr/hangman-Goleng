package main
import "os"

func main() {
	args := os.Args[1:]

	if len(args) == 2 {
		unknown := args[0]
		guessed := args[1]

		NoInterect(unknown, guessed)

	} else {
		Interect()
	}
}
