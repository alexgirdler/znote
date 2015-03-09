package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/alexgirdler/golfram"
)

func main() {
	apikey := os.Getenv("wolfram")
	c := golfram.NewClient(apikey, golfram.DefaultURL)
	bio := bufio.NewReader(os.Stdin)
	var headers []string
	for {
		line, _, err := bio.ReadLine()
		if err != nil {
			break
		}
		lineHeader, err := genHeaders(line)
		if err == nil {
			headers = append(headers, lineHeader)
		}
		line = infReplace(line)
		line = sumCompile(line)
		line = intCompile(line)
		line = limCompile(line)
		line = divisionCompile(line)
		line = golframImage(line, c)
		line = linkToHeader(line, headers)
		fmt.Println(string(line))
	}

}

func genHeaders(line []byte) (string, error) {
	headerRegex := regexp.MustCompile("^\\#(.*?)$")
	header := ""
	err := errors.New("No match")
	if headerRegex.Match(line) {
		header = string(headerRegex.FindSubmatch(line)[1])
		err = nil
	}
	return header, err
}

func infReplace(line []byte) []byte {
	infRegex := regexp.MustCompile("\\\\inf")
	line = infRegex.ReplaceAll(line, []byte("\\infty"))
	return line
}

func sumCompile(line []byte) []byte {
	sumRegex := regexp.MustCompile("\\\\sum_")
	sumRegex.Longest()
	if sumRegex.Match(line) {
		line = sumRegex.ReplaceAll(line, []byte("\\sum\\limits_"))
	}
	sumNRegex := regexp.MustCompile("\\\\sum(\\d)")
	if sumNRegex.Match(line) {
		line = sumNRegex.ReplaceAll(line, []byte("\\sum\\limits_{n=$1}^{\\infty}"))
	}
	return line
}

func intCompile(line []byte) []byte {
	intRegex := regexp.MustCompile("\\\\int_")
	intRegex.Longest()
	if intRegex.Match(line) {
		line = intRegex.ReplaceAll(line, []byte("\\int\\limits_"))
	}
	intNRegex := regexp.MustCompile("\\\\int(\\d)")
	intNRegex.Longest()
	if intNRegex.Match(line) {
		line = intNRegex.ReplaceAll(line, []byte("\\int\\limits_{$1}^{\\infty}"))
	}
	return line
}

func limCompile(line []byte) []byte {
	limNRegex := regexp.MustCompile("\\\\lim(\\d)")
	if limNRegex.Match(line) {
		line = limNRegex.ReplaceAll(line, []byte("\\lim_{n\\to$1}"))
	}
	limIRegex := regexp.MustCompile("\\\\limI")
	if limIRegex.Match(line) {
		line = limIRegex.ReplaceAll(line, []byte("\\lim_{n\\to\\infty}"))
	}
	return line
}

func divisionCompile(line []byte) []byte {
	divisionRegex := regexp.MustCompile("\\((\\S*?)/(\\S*)\\)")
	var divisionReplacement []byte
	if divisionRegex.Match(line) {
		line = divisionRegex.ReplaceAll(line, []byte("\\dfrac{$1}{$2}"))
		for {
			match := divisionRegex.Match(divisionReplacement)
			if match {
				line = divisionRegex.ReplaceAll(divisionReplacement, []byte("\\dfrac{$1}{$2}"))
			} else {
				break
			}
		}
	}
	return line
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func golframImage(line []byte, c *golfram.Client) []byte {
	golframRegex := regexp.MustCompile("\\[#WA (.*?)\\]")
	if golframRegex.Match(line) {
		query := string(golframRegex.FindSubmatch(line)[1])
		filename := "imgs/" + GetMD5Hash(strings.Replace(query, " ", "_", 1))
		//filename := strings.Replace(query, " ", "_", 1)
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			//fmt.Println(filename)
			result, _ := c.NewQuery("plot " + query)
			c.GetPlot(result, 0, filename)
		}
		line = []byte("![" + query + "](" + filename + ")\\ ")
	}
	return line
}

func linkToHeader(line []byte, headers []string) []byte {
	emphRegex := regexp.MustCompile("__(.*?)__")
	matchHeader := ""
	matched := false
	if emphRegex.Match(line) {
		term := emphRegex.FindSubmatch(line)[1]
		for _, header := range headers {
			if strings.Contains(header, string(term)) {
				if matched {
					return line
				} else {
					matched = true
					matchHeader = header
				}
			}
		}
		if matched {
			line = []byte(strings.Replace(string(line), string(term), "["+string(term)+"]["+matchHeader[1:len(matchHeader)]+"]", -1))
		}
	}
	return line
}
