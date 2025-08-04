package view

import ("fmt"
	"keuangan/model"
)

func Input(prompt string) string {
	var Scan string
	fmt.Print(prompt)
	fmt.Scan(&Scan)

	return Scan
}

func InputInt(prompt string) int {
	var Scan int
	fmt.Print(prompt)
	fmt.Scan(&Scan)

	return Scan
}

func Output(prompt string, out []model.Transaction) {
	fmt.Println(prompt, out)

}