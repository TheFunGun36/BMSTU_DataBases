package main

import (
	"bandit-gen/bandit"
	"fmt"
	"log"
	"os"
)

func main() {
	var gen bandit.BanditGenerator
	inPath := "./data/"
	outFilename := "result.csv"

	gen.LoadData(inPath)
	f, e := os.Create(outFilename)
	if e != nil {
		log.Fatalf("Can't create file %s: %v", outFilename, e)
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
