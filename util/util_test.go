//
package util

import (
	"fmt"
	"testing"
)

//
func TestMD5sMatch(t *testing.T) {
}

// tests if StringToIP works as intended
func TestStringToIP(t *testing.T) {
	name := "test-app"
	target := "172.16.7.24" // pre-calculated IP for "test-app"
	actual := StringToIP(name)

	if actual != target {
		t.Error(fmt.Sprintf("Expected IP '%s' got '%s'", target, actual))
	}
}

// tests if StringToPort works as intended
func TestStringToPort(t *testing.T) {
	name := "test-app"
	target := "11816" // pre-calculated port for "test-app"
	actual := StringToPort(name)

	if actual != target {
		t.Error(fmt.Sprintf("Expected port '%s' got '%s'", target, actual))
	}
}

// tests to ensure that IPs generated by StringToIP are unique for similar names
func TestStringToIPUnique(t *testing.T) {
	var name string

	name = "test-app"
	first := StringToIP(name)

	name = "app-test"
	second := StringToIP(name)

	if first == second {
		t.Error("IPs arent unique")
	}
}

// tests to ensure that ports generated by StringToPort are unique for similar
// names
func TestStringToPortUnique(t *testing.T) {
	var name string

	name = "test-app"
	first := StringToPort(name)

	name = "app-test"
	second := StringToPort(name)

	if first == second {
		t.Error("Ports arent unique")
	}
}