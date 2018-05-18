// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"strings"
)

// tagOptions is the string following a comma in a struct field's "json"
// tag, or the empty string. It does not include the leading comma.
type tagOptions string

// parseTag splits a struct field's json tag into its name and
// comma-separated options.
func parseTag(tag string) (string, tagOptions) {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], tagOptions(tag[idx+1:])
	}
	return tag, tagOptions("")
}

// Contains reports whether a comma-separated list of options
// contains a particular substr flag. substr must be surrounded by a
// string boundary or commas.
func (o tagOptions) Contains(optionName string) bool {
	if len(o) == 0 {
		return false
	}
	s := string(o)
	for s != "" {
		var next string
		i := strings.Index(s, ",")
		if i >= 0 {
			s, next = s[:i], s[i+1:]
		}
		if s == optionName {
			return true
		}
		s = next
	}
	return false
}

const filterOptsKey = "filter:"
type filterOpts map[string]string

func parseFilterOpts(optsStr string) filterOpts {
	opts := filterOpts{}

	if len(optsStr) == 0 {
		return opts
	}

	s, optsStr := optsStr, ""
	for s != "" {
		var next string
		i := strings.Index(s, ",")
		if i >= 0 {
			s, next = s[:i], s[i+1:]
		}

		j := strings.Index(s, filterOptsKey)
		if j >= 0 {
			optsStr = s[j+len(filterOptsKey):]
			break
		} else {
			s = next
		}
	}

	s = optsStr
	for s != "" {
		var next string
		i := strings.Index(s, ";")
		if i >= 0 {
			s, next = s[:i], s[i+1:]
		}

		j := strings.Index(s, ".")
		if j >= 0 {
			key, value := s[:j], s[j+1:]
			opts[key] = value
		} else if len(s) > 0 {
			opts[s] = ""
		}

		s = next
	}

	return opts
}
