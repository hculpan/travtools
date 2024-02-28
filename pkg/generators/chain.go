package generators

import (
	"bufio"
	"math/rand"
	"strings"
	"unicode"

	"github.com/hculpan/travtools/pkg/embed"
)

const (
	start = "_start_"
)

type Chain struct {
	chain     map[string][]string
	list      []string
	minLength int
	maxLength int
}

func NewChainFromFile(filename string, minLength, maxLength int) (*Chain, error) {
	list, err := readFileLines(filename)
	if err != nil {
		return nil, err
	}

	return NewChain(list, minLength, maxLength), nil
}

func NewChain(list []string, minLength, maxLength int) *Chain {
	return &Chain{
		chain:     make(map[string][]string),
		list:      list,
		minLength: minLength,
		maxLength: maxLength,
	}
}

// Build takes a slice of strings and builds the Markov Chain with it.
func (c *Chain) Build() {
	for _, name := range c.list {
		processedName := strings.ToLower(name) + " " // Append a space to signify the end of a name
		c.add(start, string(processedName[0]))
		for i := 0; i < len(processedName)-1; i++ {
			c.add(string(processedName[i]), string(processedName[i+1]))
		}
	}
}

func (c *Chain) add(key, next string) {
	c.chain[key] = append(c.chain[key], next)
}

func (c *Chain) GenerateName() string {
	name := ""
	if rand.Intn(100) > 50 {
		name = c.list[rand.Intn(len(c.list))]
	} else {
		current := start
		for {
			for {
				if len(name) >= c.maxLength {
					break
				}

				nextChars, exists := c.chain[current]
				if !exists || len(nextChars) == 0 {
					break
				}

				next := nextChars[rand.Intn(len(nextChars))]
				if next == " " {
					break
				}
				name += next
				current = next
			}
			if len(name) >= c.minLength {
				break
			}
		}
	}
	return string(unicode.ToUpper(rune(name[0]))) + name[1:]
}

func readFileLines(filename string) ([]string, error) {
	fileData, err := embed.ReadDataFile(filename)
	if err != nil {
		return nil, err
	}

	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(string(fileData)))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
