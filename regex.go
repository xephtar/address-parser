package addressParser

import (
	"regexp"
)

func getDistrictParser() *regexp.Regexp {
	return regexp.MustCompile(`([a-z0-9].*(\bmahallesi|\bmahalle|\bmah|\bmh))`)
}

func getStreetParser() *regexp.Regexp {
	return regexp.MustCompile(`([a-z0-9].*(\bcaddesi|\bcadde|\bcad|\bcd|\bbulvar|\bblv|\bbulvari))`)
}

func getSubStreetParser() *regexp.Regexp {
	return regexp.MustCompile(`([a-z0-9].*(\bsokagi|\bsokak|\bsok|\bsk|\bsite|\bsitesi|\bkonutlari\bevleri|\blojman|\blojmani|\bplaza|\bis merkezi|\bis hani)|\brezidans)`)
}