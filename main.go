/**
 * MIT License
 *
 * Copyright (c) 2025 Dks
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package main

import (
	"change_case/cc_lib"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	PROGRAM_NAME := "change_case"

	args := os.Args
	input, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Println("Problem with reading the file")
	}

	ext := filepath.Ext(args[1])
	languages_data, err_json := os.ReadFile(`assets/languages.json`)

	if err_json != nil {
		fmt.Println("Problem with languages assets")
	}

	var langs []cc_lib.Language
	err_lang := json.Unmarshal(languages_data, &langs)

	if err_lang != nil {
		fmt.Println("Problem with reading the file")
	}

	selectedLanguage := cc_lib.LoadLang(langs, ext)

	help_directives := "sc [snake -> camel]\n" +
		"pc [pascal -> camel]\n" +
		"sp [snake -> pascal]\n" +
		"ps [pascal -> snake]\n" +
		"cs [camel -> snake]\n" +
		"cp [camel -> pascal]"

	if len(args) <= 2 {
		fmt.Println("Specify the input output directives")
		fmt.Printf("example: %s test.js pc [pascal -> camel] \n", PROGRAM_NAME)
		os.Exit(1)
	}

	if len(args[2]) < 2 {
		fmt.Println("Select a formato like this:")
		fmt.Println(help_directives)
		os.Exit(1)
	}

	from := args[2][0]
	to := args[2][1]

	lines := cc_lib.GetLines(string(input))

	for _, line := range lines {
		tokens := cc_lib.Tokenize(line)

		for _, tok := range tokens {

			switch from {
			case 's':
				cc_lib.FromSnake(tok, to, selectedLanguage)
			case 'c':
				cc_lib.FromCamel(tok, to, selectedLanguage)
			case 'p':
				cc_lib.FromPascal(tok, to, selectedLanguage)
			default:
				fmt.Println("Specify correct input output directives")
				fmt.Println(help_directives)
				os.Exit(1)
			}

		}

	}

}
