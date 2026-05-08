package generator

import (
	"reflect"
	"testing"
)

// Unit test for ascii-generator function.
func TestAscii(t *testing.T) {
	input := "hello"

	// Using standard strings with explicit \n to prevent invisible spacing bugs
	expected := "" +
		" _              _   _          \n" +
		"| |            | | | |         \n" +
		"| |__     ___  | | | |   ___   \n" +
		"|  _ \\   / _ \\ | | | |  / _ \\  \n" +
		"| | | | |  __/ | | | | | (_) | \n" +
		"|_| |_|  \\___| |_| |_|  \\___/  \n" +
		"                               \n" +
		"                               \n"

	output := AsciiGen(input, "standard", "", "")

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Failed.\n Expected:\n[%v]\n Got:\n[%v]", expected, output)
	}
}
