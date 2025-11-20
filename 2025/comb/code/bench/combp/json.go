package combp

import (
	"strconv"

	"github.com/flowdev/comb"
	"github.com/flowdev/comb/cmb"
)

// break initialization cycle:
func init() {
	valuep = comb.LazyBranchParser(ParseValue)
}

type (
	// JSONValue represents any value that can be encountered in
	// JSON, including complex types like objects and arrays.
	JSONValue any

	// JSONString represents a JSON string value.
	JSONString string

	// JSONNumber represents a JSON number value, which internally is treated as float64.
	JSONNumber float64

	// JSONObject represents a JSON object, which is a collection of key-value pairs.
	JSONObject map[string]JSONValue

	// JSONArray represents a JSON array, which is a list of JSON values.
	JSONArray []JSONValue

	// JSONBool represents a JSON boolean value.
	JSONBool bool

	// JSONNull represents the JSON null value.
	JSONNull struct{}
)

var NullValue = JSONNull(struct{}{})

// kv is a struct representing a key-value pair in a JSON object.
//
// 'key' holds the string key, and 'value' holds the corresponding JSON value.
type kv struct {
	key   string
	value JSONValue
}

// 1 OMIT
// ParseValue is a parser that attempts to parse different types of
// JSON values (object, array, string, etc.).
func ParseValue() comb.Parser[JSONValue] {
	return cmb.FirstSuccessful(
		parseObject(),
		parseArray(),
		parseString(),
		cmb.Map(cmb.Float64(true, 10, false), func(f float64) (JSONValue, error) { return f, nil }),
		cmb.Assign(JSONValue(JSONBool(true)), cmb.String("true")),
		cmb.Assign(JSONValue(JSONBool(false)), cmb.String("false")),
		cmb.Assign(JSONValue(NullValue), cmb.String("null")),
	)
}
// 2 OMIT

var valuep comb.Parser[JSONValue]

// parseObject parses a JSON object, which starts and ends with
// curly braces and contains key-value pairs.
func parseObject() comb.Parser[JSONValue] {
	return cmb.Map3(
		cmb.Char('{'),
		cmb.Separated0(memberParser(), cmb.Char(','), false),
		cmb.Char('}'),
		func(_ rune, kvs []kv, _ rune) (JSONValue, error) {
			obj := make(JSONObject)
			for _, kv := range kvs {
				obj[kv.key] = kv.value
			}
			return obj, nil
		},
	)
}

// 3 OMIT
// parseArray parses a JSON array, which starts and ends with
// square brackets and contains a list of values.
func parseArray() comb.Parser[JSONValue] {
	return cmb.Map(
		cmb.Delimited(
			cmb.Char('['),
			cmb.Separated0(elementParser(), cmb.String(","), false),
			cmb.Char(']'),
		),
		func(elements []JSONValue) (JSONValue, error) {
			return JSONArray(elements), nil
		},
	)
}
// 4 OMIT

// parseString parses a JSON string.
func parseString() comb.Parser[JSONValue] {
	return cmb.Map3(
		cmb.Char('"'),
		cmb.Many0(characterParser()),
		cmb.Char('"'),
		func(_ rune, chars []rune, _ rune) (JSONValue, error) {
			return JSONString(string(chars)), nil
		},
	)
}

// member creates a parser for a single key-value pair in a JSON object.
//
// It expects a string followed by a colon and then a JSON value.
// The result is a kv struct with the parsed key and value.
func memberParser() comb.Parser[kv] {
	return cmb.Map3(
		cmb.Delimited(wsp(), parseString(), wsp()),
		cmb.Char(':'),
		elementParser(),
		func(k JSONValue, _ rune, v JSONValue) (kv, error) {
			strKey, _ := k.(JSONString)
			return kv{key: string(strKey), value: v}, nil
		},
	)
}

// element creates a parser for a single element in a JSON array.
//
// It wraps the element with optional whitespace on either side.
func elementParser() comb.Parser[JSONValue] {
	return cmb.Delimited(wsp(), valuep, wsp())
}

// character creates a parser for a single JSON string character.
//
// It distinguishes between regular characters and escape sequences.
func characterParser() comb.Parser[rune] {
	return cmb.FirstSuccessful(
		cmb.Satisfy("normal character", func(c rune) bool {
			return c != '"' && c != '\\' && c >= 0x20 && c <= 0x10FFFF
		}),

		// escape
		escapeParser(),
	)
}

// escape creates a parser for escaped characters in a JSON string.
//
// It handles common escape sequences like '\n', '\t', etc., and unicode escapes.
func escapeParser() comb.Parser[rune] {
	mapFunc := func(char1, char2 rune) (rune, error) {
		// char1 will always be '\\'
		switch char2 {
		case 'b':
			return '\b', nil
		case 'f':
			return '\f', nil
		case 'n':
			return '\n', nil
		case 'r':
			return '\r', nil
		case 't':
			return '\t', nil
		default: // for unicode escapes
			return char2, nil
		}
	}

	return cmb.Map2(
		cmb.Char('\\'),
		cmb.FirstSuccessful(
			cmb.OneOfRunes('"', '\\', '/', 'b', 'f', 'n', 'r', 't'),
			unicodeEscapeParser(),
		),
		mapFunc,
	)
}

// unicodeEscape creates a parser for a Unicode escape sequence in a JSON string.
//
// It expects a sequence starting with 'u' followed by four hexadecimal digits and
// converts them to the corresponding rune.
func unicodeEscapeParser() comb.Parser[rune] {
	mapFunc := func(_ rune, hex string) (rune, error) {
		// char will always be 'u'
		codePoint, err := strconv.ParseInt(hex, 16, 32)
		if err != nil {
			return 0, err
		}
		return rune(codePoint), nil
	}

	return cmb.Map2(
		cmb.Char('u'),
		cmb.SatisfyMN("hex digit", 4, 4, cmb.IsHexDigit),
		mapFunc,
	)
}

// wsp creates a parser for whitespace in JSON.
//
// It can handle spaces, tabs, newlines, and carriage returns.
// The parser accumulates all whitespace characters and returns them as a single string.
func wsp() comb.Parser[string] {
	mapFunc := func(runes []rune) (string, error) {
		return string(runes), nil
	}

	return cmb.Map(cmb.Many0(
		cmb.OneOfRunes(' ', '\t', '\n', '\r'),
	), mapFunc)
}
