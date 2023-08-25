package rk

import (
	"math"
	"os"
	"reflect"
	"regexp"
	"strings"
)

var (
	// commentExpression string         = `^.*\` + comment + `.*\n|\s+`
	// (?m)[\s]+.*\%.*\n
	commentExpression string         = `(?m)[\s]*\` + comment + `.*$|\s+`
	re                *regexp.Regexp = regexp.MustCompile(commentExpression)
)

type PlagarismChecker struct {
	HashTable map[string][]int
	KGram     int
}

type RabinKarp struct {
	Base        int
	Text        string
	PatternSize int
	Start       int
	End         int
	Mod         int
	Hash        int
}

func NewRabinKarp(text string, patternSize int) *RabinKarp {
	rb := &RabinKarp{
		Base:        26,
		PatternSize: patternSize,
		Start:       0,
		End:         0,
		Mod:         5807,
		Text:        text,
	}
	rb.GetHash()
	return rb
}

func (rb *RabinKarp) GetHash() {
	hashValue := 0
	for n := 0; n < rb.PatternSize; n++ {
		runeText := []rune(rb.Text)
		value := int(runeText[n])
		mathpower := int(math.Pow(float64(rb.Base), float64(rb.PatternSize-n-1)))
		hashValue = Mod((hashValue + (value-96)*(mathpower)), rb.Mod)
	}
	rb.Start = 0
	rb.End = rb.PatternSize
	rb.Hash = hashValue
}

func (rb *RabinKarp) NextWindow() bool {
	if rb.End <= len(rb.Text)-1 {
		textBytes := []byte(rb.Text)
		mathpower := int(math.Pow(float64(rb.Base), float64(rb.PatternSize-1)))

		rb.Hash -= (int(textBytes[rb.Start]) - 96) * mathpower
		rb.Hash *= rb.Base
		rb.Hash += int(textBytes[rb.End]) - 96
		rb.Hash = Mod(rb.Hash, rb.Mod)
		rb.Start++
		rb.End++
		return true
	}
	return false
}

// CurrentWindowText return the current window text
func (rb *RabinKarp) CurrentWindowText() string {
	return rb.Text[rb.Start:rb.End]
}

func Checker(text, pattern string) string {
	textRolling := NewRabinKarp(strings.ToLower(text), len(pattern))
	patternRolling := NewRabinKarp(strings.ToLower(pattern), len(pattern))

	for i := 0; i <= len(text)-len(pattern)+1; i++ {
		if textRolling.Hash == patternRolling.Hash {
			return "Found"
		}
		textRolling.NextWindow()
	}
	return "Not Found"
}

func (pc PlagarismChecker) PrepareContent(content string) string {
	data := re.ReplaceAllString(content, "")
	data = strings.ReplaceAll(data, " ", "")
	return data
}

func NewPlagarismChecker(fileA, fileB string) *PlagarismChecker {
	checker := &PlagarismChecker{
		KGram:     5,
		HashTable: make(map[string][]int),
	}
	checker.CalculateHash(checker.GetFileContent(fileA), "a")
	checker.CalculateHash(checker.GetFileContent(fileB), "b")
	return checker
}

func (pc PlagarismChecker) CalculateHash(content, docType string) {
	text := pc.PrepareContent(content)
	textRolling := NewRabinKarp(strings.ToLower(text), pc.KGram)

	for i := 0; i <= len(text)-pc.KGram+1; i++ {
		if len(pc.HashTable[docType]) == 0 {
			pc.HashTable[docType] = []int{textRolling.Hash}
		} else {
			pc.HashTable[docType] = append(pc.HashTable[docType], textRolling.Hash)
		}
		if textRolling.NextWindow() == false {
			break
		}
	}
}

func (pc PlagarismChecker) GetRate() float64 {
	return pc.CalculatePlagarismRate()
}

func (pc PlagarismChecker) GetFileContent(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (pc PlagarismChecker) CalculatePlagarismRate() float64 {
	THA := len(pc.HashTable["a"])
	THB := len(pc.HashTable["b"])
	intersect := Intersect(pc.HashTable["a"], pc.HashTable["b"])
	SH := reflect.ValueOf(intersect).Len()

	// Formular for plagiarism rate
	// P = (2 * SH / THA + THB ) 100%
	p := float64(2*SH) / float64(THA+THB)
	return float64(p * 100)
}
