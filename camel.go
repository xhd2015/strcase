/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2015 Ian Coleman
 * Copyright (c) 2018 Ma_124, <github.com/Ma124>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, Subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or Substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package strcase

import (
	"strings"
)

// Converts a string to CamelCase
func toCamelInitCase(s string, initCase bool) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	if a, ok := uppercaseAcronym[s]; ok {
		s = a
	}

	n := strings.Builder{}
	n.Grow(len(s))
	capNext := initCase
	for i, v := range []byte(s) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if capNext {
			if vIsLow {
				v += 'A'
				v -= 'a'
			}
		} else if i == 0 {
			if vIsCap {
				v += 'a'
				v -= 'A'
			}
		}
		if vIsCap || vIsLow {
			n.WriteByte(v)
			capNext = false
		} else if vIsNum := v >= '0' && v <= '9'; vIsNum {
			n.WriteByte(v)
			capNext = true
		} else {
			capNext = v == '_' || v == ' ' || v == '-' || v == '.'
		}
	}
	return n.String()
}

const diff_toLower byte = 'a' - 'A' // 32
const diff_toUpper = 'A' - 'a'      // -32

// only split at boundary: '_. -'
func toCamelInitCaseV2(s string, initCase bool, toCamelMap map[string]string) string {
	if len(toCamelMap) == 0 {
		return toCamelInitCase(s, initCase)
	}
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	n := strings.Builder{}
	n.Grow(len(s))

	capNext := initCase
	sz := len(s)
	if sz > 12 {
		sz = 12
	}
	word := make([]byte, 0, sz)

	for i, v := range []byte(s) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if capNext {
			if vIsLow {
				if len(word) > 0 {
					newWord, ok := toCamelMap[string(word)]
					if ok {
						n.WriteString(newWord)
					} else {
						n.Write(word)
					}
					word = word[0:0]
				}
				v -= diff_toLower // to upper
			}
		} else if i == 0 {
			if vIsCap {
				v += diff_toLower
			}
		}
		if vIsCap || vIsLow {
			word = append(word, v)
			capNext = false
		} else if vIsNum := v >= '0' && v <= '9'; vIsNum {
			word = append(word, v)
			capNext = true
		} else {
			capNext = v == '_' || v == ' ' || v == '-' || v == '.'
		}
	}
	// the finaly word
	if len(word) > 0 {
		newWord, ok := toCamelMap[string(word)]
		if ok {
			n.WriteString(newWord)
		} else {
			n.Write(word)
		}
	}
	return n.String()
}

// ToCamel converts a string to CamelCase
func ToCamel(s string) string {
	return toCamelInitCase(s, true)
}

// ToCamel converts a string to CamelCase
func ToCamelWithMap(s string, wordMap map[string]string) string {
	return toCamelInitCaseV2(s, true, wordMap)
}

// ToLowerCamel converts a string to lowerCamelCase
func ToLowerCamel(s string) string {
	return toCamelInitCase(s, false)
}

// ToLowerCamel converts a string to lowerCamelCase
func ToLowerCamelWithMap(s string, wordMap map[string]string) string {
	return toCamelInitCaseV2(s, false, wordMap)
}
