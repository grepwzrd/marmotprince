package player

import (
	"math/rand"
	"strings"
	"time"
)

func makeName() string {
	firstPrefixes := []string{"bazz", "dev", "mik", "meep", "skwee", "chip", "cheep", "mip", "beez", "buzz", "chizz", "chazz", "cheez"}
	firstSuffixes := []string{"ali", "chee", "chik", "mik", "chak", "clik", "bee", "bay", "bok", "bik", "ak", "aak", "yip", "yap", "yop"}
	secondPrefixes := []string{"wik", "wak", "wok", "qwik", "qwak", "gwa", "guh", "guj", "guk", "gut", "gum", "wuk", "rak", "rok", "raq", "ratt", "jii", "jik", "jak", "jok"}
	secondSuffixes := []string{"ayy", "ee", "ei", "ii", "iii", "lee", "kee", "kii", "qee", "qey", "jee", "gee", "wee", "ree", "wii", "rii", "qi"}
	fp := firstPrefixes[rand.Intn(len(firstPrefixes))]
	fs := firstSuffixes[rand.Intn(len(firstSuffixes))]
	sp := secondPrefixes[rand.Intn(len(secondPrefixes))]
	ss := secondSuffixes[rand.Intn(len(secondSuffixes))]
	firstNames := []string{fp, fs}
	lastNames := []string{sp, ss}
	firstName := strings.Join(firstNames, "")
	lastName := strings.Join(lastNames, "")
	fullName := []string{firstName, lastName}
	return strings.Join(fullName, " ")
}

func GenerateName() string {
	rand.Seed(time.Now().Unix())
	return makeName()
}
