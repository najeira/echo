package echoutil

import "testing"

type escapeTest struct {
	in  string
	out string
	err string
}

var unescapeTests = []escapeTest{
	{
		"",
		"",
		"",
	},
	{
		"abc",
		"abc",
		"",
	},
	{
		"1%41",
		"1A",
		"",
	},
	{
		"1%41%42%43",
		"1ABC",
		"",
	},
	{
		"%4a",
		"J",
		"",
	},
	{
		"%6F",
		"o",
		"",
	},
	{
		"%", // not enough characters after %
		"",
		"%",
	},
	{
		"%a", // not enough characters after %
		"",
		"%a",
	},
	{
		"%1", // not enough characters after %
		"",
		"%1",
	},
	{
		"123%45%6", // not enough characters after %
		"",
		"%6",
	},
	{
		"%zzzzz", // invalid hex digits
		"",
		"%zz",
	},
	{
		"a%20b",
		"a b",
		"",
	},
}

func TestUnescape(t *testing.T) {
	for _, tt := range unescapeTests {
		in := tt.in
		out := tt.out
		actual, err := pathUnescape(in)
		if actual != out || (err != nil) != (tt.err != "") {
			t.Errorf("pathUnescape(%q) = %q, %s; want %q, %s", in, actual, err, out, tt.err)
		}
	}
}
