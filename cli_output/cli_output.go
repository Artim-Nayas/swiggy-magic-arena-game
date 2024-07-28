package cli_output

import "fmt"

func Render(output string) {
	fmt.Printf("%s\n", output)
}

func RenderInvalidOperation() {
	fmt.Println("Invalid operation")
}

func RenderCliInput() {
	fmt.Printf("\n>")
}
