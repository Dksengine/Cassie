package cc_lib

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
	"unicode"
)

type Language struct {
	Extension string   `json:extension`
	Keys      []string `json:keys`
}

// This is an implementation of a function that transforms the input to snake and camel.
// Due to the similitary between camel and pascal cases, it's not necessary duplicate the logic
func OKTransformFromSToCOrP(token string, camel bool) string {
	res := ""

	if token == "OKTransformFromSToCOrP" {
		fmt.Print("OK")
	}

	re := regexp.MustCompile(`^[a-z]+(?:_[a-zA-Z0-9]+)+$`)
	parts := re.FindAllString(token, -1)

	if len(parts) == 0 {
		return token
	}

	for _, part := range parts {
		ReSplit := regexp.MustCompile(`_`)
		PartsSplit := ReSplit.Split(part, -1)

		for j, p := range PartsSplit {
			runes := []rune(p)

			if camel && j == 0 {
				continue
			}
			runes[0] = unicode.ToUpper(runes[0])
			PartsSplit[j] = string(runes)
		}

		res = strings.Join(PartsSplit, "")

	}

	return res
}

func FromSnakeToCamel(token string) string {
	return OKTransformFromSToCOrP(token, true)
}

func FromSnakeToPascal(token string) string {
	return OKTransformFromSToCOrP(token, false)
}

func FromCamelOrPascalToSnake(token string) string {

	res := ""

	re := regexp.MustCompile(`\b\w+[A-Z]\w+\b`)
	parts := re.FindAllString(token, -1)

	if len(parts) == 0 {
		return token
	}

	for _, part := range parts {
		ReSplit := regexp.MustCompile(`([a-z]+|[A-Z][a-z]+)`)
		PartsSplit := ReSplit.FindAllString(part, -1)

		for j, p := range PartsSplit {
			runes := []rune(p)

			runes[0] = unicode.ToLower(runes[0])
			PartsSplit[j] = string(runes)
		}

		res = strings.Join(PartsSplit, "_")
	}
	return res
}

func TransformFromCToPViceversa(token string, camel bool) string {
	res := []rune(token)

	if camel {
		res[0] = unicode.ToUpper(res[0])
	} else {
		res[0] = unicode.ToLower(res[0])
	}

	return string(res)
}

func FromCamelToPascal(token string) string {
	return TransformFromCToPViceversa(token, true)
}

func FromPascalToCamel(token string) string {
	return TransformFromCToPViceversa(token, false)
}

func GetLines(input string) []string {
	reader := bufio.NewReader(strings.NewReader(input))
	lines := []string{}

	for {
		line, err := reader.ReadString('\n')
		lines = append(lines, line)

		if err == io.EOF {
			if len(line) > 0 && !strings.HasSuffix(line, "\n") {
				lines[len(lines)-1] = line
			}
			break
		}
	}

	return lines
}

func Tokenize(tokens string) []string {
	current := tokens

	re := regexp.MustCompile(`\w+|\W+`)
	return re.FindAllString(current, -1)

}

func FromSnake(token string, output byte, lang *Language) {
	if output == 'c' {
		if IsKeyWord(lang, token) {
			fmt.Printf("%s", (token))
		} else {
			fmt.Printf("%s", FromSnakeToCamel(token))
		}
	} else if output == 'p' {
		if IsKeyWord(lang, token) {
			fmt.Printf("%s", (token))
		} else {
			fmt.Printf("%s", FromSnakeToPascal(token))
		}
	} else {
		fmt.Printf("%s", token)
	}
}

func FromCamel(token string, output byte, lang *Language) {
	if output == 's' {
		if IsKeyWord(lang, token) {
			fmt.Printf("%s", (token))
		} else {
			fmt.Printf("%s", FromCamelOrPascalToSnake(token))
		}
	} else if output == 'p' {
		if IsKeyWord(lang, token) {
			fmt.Printf("%s", (token))
		} else {
			fmt.Printf("%s", FromCamelToPascal(token))
		}
	} else {
		fmt.Printf("%s", token)
	}
}

func FromPascal(token string, output byte, lang *Language) {
	if output == 'c' {
		if IsKeyWord(lang, token) {
			fmt.Printf("%s", (token))
		} else {
			fmt.Printf("%s", FromPascalToCamel(token))
		}
	} else if output == 's' {
		if IsKeyWord(lang, token) {
			fmt.Printf("%s", (token))
		} else {
			fmt.Printf("%s", FromCamelOrPascalToSnake(token))
		}
	} else {
		fmt.Printf("%s", token)
	}
}

func IsKeyWord(lang *Language, token string) bool {
	if lang == nil {
		return false
	}
	for _, key := range lang.Keys {
		if token == key {
			return true
		}
	}
	return false
}

func LoadLang(langs []Language, ext string) *Language {
	for _, language := range langs {
		if language.Extension == ext {
			return &language
		}
	}
	return nil
}
