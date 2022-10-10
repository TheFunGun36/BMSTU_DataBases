package bandit

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type BanditGenerator struct {
	ready            bool
	nicknameCounter  int
	firstNamesMale   []string
	firstNamesFemale []string
	lastNames        []string
	midNamesMale     []string
	midNamesFemale   []string
	influence        []string
	nicknamesMale    []string
	nicknamesFemale  []string
}

func (b *BanditGenerator) loadSingle(dst *[]string, filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Open file error: %v", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		*dst = append(*dst, strings.Trim(scanner.Text(), " \r\n\t"))
	}

	return 1
}

func (bg *BanditGenerator) LoadData(path string) {
	var c = 0
	c += bg.loadSingle(&bg.firstNamesMale, path+"first_names_male.txt")
	c += bg.loadSingle(&bg.firstNamesFemale, path+"first_names_female.txt")
	c += bg.loadSingle(&bg.lastNames, path+"last_names.txt")
	c += bg.loadSingle(&bg.midNamesMale, path+"mid_names_male.txt")
	c += bg.loadSingle(&bg.midNamesFemale, path+"mid_names_female.txt")
	c += bg.loadSingle(&bg.influence, path+"influence_spheres.txt")
	c += bg.loadSingle(&bg.nicknamesMale, path+"nicknames_male.txt")
	c += bg.loadSingle(&bg.nicknamesFemale, path+"nicknames_female.txt")

	bg.ready = c == 8
	bg.nicknameCounter = 0

	rand.Seed(time.Now().Unix())
}

func randDate() time.Time {
	min := time.Date(1950, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2004, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func (bg *BanditGenerator) BanditNext() *Bandit {
	if bg.nicknameCounter >= len(bg.nicknamesMale)+len(bg.nicknamesFemale) {
		return nil
	}

	res := new(Bandit)
	res.Gender = bg.nicknameCounter >= len(bg.nicknamesMale)

	// gender independent
	res.BirthDate = randDate()
	res.Influence.Int16 = int16(rand.Intn(len(bg.influence)))
	res.LastName = bg.lastNames[rand.Intn(len(bg.lastNames))]
	res.MidName.Valid = true

	if res.Gender {
		// female
		res.LastName += "а"
		res.Nickname = bg.nicknamesFemale[bg.nicknameCounter-len(bg.nicknamesMale)]
		res.FirstName = bg.firstNamesFemale[rand.Intn(len(bg.firstNamesFemale))]
		res.MidName.String = bg.midNamesFemale[rand.Intn(len(bg.midNamesFemale))]
	} else {
		// male
		res.Nickname = bg.nicknamesMale[bg.nicknameCounter]
		res.FirstName = bg.firstNamesMale[rand.Intn(len(bg.firstNamesMale))]
		res.MidName.String = bg.midNamesMale[rand.Intn(len(bg.midNamesMale))]
	}

	bg.nicknameCounter += 1
	return res
}
