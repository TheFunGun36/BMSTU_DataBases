package main

import (
	"bandit-gen/bandit"
	"fmt"
	"log"
	"os"
)

func generateBandits(filename string, data string) {
	var gen bandit.BanditGenerator
	gen.LoadData(data)
	f, e := os.Create(filename)
	if e != nil {
		log.Fatalf("Can't create file %s: %v", filename, e)
	}
	defer f.Close()

	line := fmt.Sprintln(
		"nickname,lastname,firstname,midname,birth,gender,influence")
	f.WriteString(line)

	b := gen.BanditNext()
	for b != nil {
		line = fmt.Sprintf("%s,%s,%s,%s,%s,%t,%d\n",
			b.Nickname,
			b.LastName,
			b.FirstName,
			b.MidName.String,
			b.BirthDate.Format("2006-01-02"),
			b.Gender,
			int(b.Influence.Int16))

		f.WriteString(line)
		b = gen.BanditNext()
	}
}

func generateFactions(filename string, data string) {

}

func main() {
	data := "./data/"

	generateBandits("./out/bandits.csv", data)
	os.Mkdir("./out", 0755)

}
