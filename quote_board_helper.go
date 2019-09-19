package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/urfave/cli"
	"golang.org/x/xerrors"
)

func main() {
	app := cli.NewApp()
	app.Name = "Quote Board Helper"
	app.Usage = "Helps get the letters you need to create a quote"
	app.Action = mainAction
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "filename, f",
			Usage: "name of the file you want to parse",
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func mainAction(c *cli.Context) error {

	var phrase string
	var err error
	filename := c.GlobalString("filename")
	if filename != "" {
		fmt.Println(filename)
		// assume argument present
		phrase, err = getPhraseFromFile(filename)
	} else {
		phrase, err = getPhraseFromInteractive()
	}
	if err != nil {
		return err
	}

	letterCounts := getLetterCounts(phrase)
	printLetterCounts(letterCounts, os.Stdout)
	return nil
}

func getPhraseFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var s string
	for scanner.Scan() {
		s += scanner.Text()
	}

	return s, nil
}

func getPhraseFromInteractive() (string, error) {
	fmt.Print("Enter quote here: ")
	reader := bufio.NewReader(os.Stdin)
	phrase, err := reader.ReadString('\n')
	if err != nil {
		return "", xerrors.Errorf("error reading user's string in interactive mode: %w", err)
	}
	return phrase, nil
}

type letterCount struct {
	letter string
	count  int
}

// getLetterCount counts the number of letters in a given string
// returns in sorted order
// omits blank spaces
func getLetterCounts(phrase string) []letterCount {
	counts := make(map[string]int)
	phrase = cleanseString(phrase)
	for _, letter := range phrase {
		counts[string(letter)]++
	}

	results := make([]letterCount, 0, len(counts))
	for k, v := range counts {
		results = append(results, letterCount{
			letter: k,
			count:  v,
		})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].letter < results[j].letter
	})

	return results
}

// remove spaces, new lines
func cleanseString(input string) string {
	input = strings.Replace(input, " ", "", -1)
	input = strings.TrimSuffix(input, "\n")
	return input
}

func printLetterCounts(letterCounts []letterCount, writer io.Writer) {
	for _, letterCount := range letterCounts {
		s := fmt.Sprintf("%v : %v\n", letterCount.letter, letterCount.count)
		writer.Write([]byte(s))
	}
}
