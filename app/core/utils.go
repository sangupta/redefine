/*

Redefine - UI component documentation

MIT License.
Copyright (c) 2022, Sandeep Gupta.

Use of this source code is governed by a MIT style license
that can be found in LICENSE file in the code repository.

*/

package core

import (
	"errors"
	"os"
)

// Check if a file should be included in the list
// of files to be parsed, against the inclusion rules
// defined in user supplied configuation
func isFileIncluded(path string, includes []string) bool {
	for _, pattern := range includes {
		if isWildCardMatch(path, pattern) {
			return true
		}
	}

	return false
}

// Check if the given string is a wildcard match against
// the given wildcard pattern. Returns a `bool`.
func isWildCardMatch(str string, pattern string) bool {
	i := 0
	j := 0
	starIndex := -1
	iIndex := -1

	strLen := len(str)
	patternLen := len(pattern)

	for i < strLen {
		if j < patternLen && (string(pattern[j]) == "?" || string(pattern[j]) == string(str[i])) {
			i++
			j++
		} else if j < patternLen && string(pattern[j]) == "*" {
			starIndex = j
			iIndex = i
			j++
		} else if starIndex != -1 {
			j = starIndex + 1
			i = iIndex + 1
			iIndex++
		} else {
			return false
		}
	}

	for j < patternLen && string(pattern[j]) == "*" {
		j++
	}

	return j == patternLen
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	if err == nil {
		return true
	}

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return false
}
