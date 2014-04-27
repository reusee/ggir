package glib

import (
	"strings"
	"testing"
)

func TestReturnBytes(t *testing.T) {
	n, res := Base64Decode("Zm9vYmFy")
	if n != 6 || string(res) != "foobar" {
		t.Fatal("Base64Decode")
	}

	contents, n, r, err := FileGetContents("foo")
	if string(contents) != "foobar\n" || n != 7 || r != true || err != nil {
		t.Fatal("FileGetContents")
	}

	s := "foobarbaz"
	nRead, nWritten, name, err := FilenameFromUtf8(s, int64(len(s)))
	if nRead != 9 || nWritten != 9 || string(name) != s || err != nil {
		t.Fatal("FilenameFromUtf8")
	}
}

func TestReturnStringSlice(t *testing.T) {
	envs := GetEnviron()
	if len(envs) == 0 {
		t.Fatal("GetEnviron")
	}

	names := GetLanguageNames()
	if len(names) == 0 {
		t.Fatal("GetLanguageNames")
	}

	variants := GetLocaleVariants("zh_CN")
	if strings.Join(variants, " ") != "zh_CN zh" {
		t.Fatal("GetLocaleVariants")
	}

	envNames := Listenv()
	if len(envNames) == 0 {
		t.Fatal("Listenv")
	}

	parts := RegexSplitSimple("a", "ababab", 0, 0)
	if strings.Join(parts, "|") != "|b|b|b" {
		t.Fatal("RegexSplitSimple")
	}

	n, parts, ok, err := ShellParseArgv("foo bar baz")
	if n != 3 || strings.Join(parts, "|") != "foo|bar|baz" || !ok || err != nil {
		t.Fatal("ShellParseArgv")
	}

	parts = Strsplit("foo bar baz", " ", -1)
	if strings.Join(parts, "|") != "foo|bar|baz" {
		t.Fatal("Strsplit")
	}
}
