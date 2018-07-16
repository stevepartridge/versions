package versions

import (
	"testing"
)

var (
	verTest1 = "1.2.3"
	verTest2 = "1.2"
	verTest3 = "1.2.3.4"
	verTest4 = "1."
	verTest5 = "1"

	sha1TestBytes = []byte("1234")
	sha1TestHash  = "7110eda4d09e062aa5e4a390b0a572ac0d2c0220"

	sha256TestBytes = []byte("1234")
	sha256TestHash  = "03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4"
)

func Test_Unit_Parse_Valid(t *testing.T) {

	major, minor, revision := Parse(verTest1)
	if major != "1" || minor != "2" || revision != "3" {
		t.Errorf("Expected %s but saw %d.%d.%d",
			verTest1,
			major,
			minor,
			revision,
		)
	}
}

func Test_Unit_Parse_Empty(t *testing.T) {

	major, minor, revision := Parse("")
	if major != "" && minor != "" && revision != "" {
		t.Errorf("Expected %s but saw %d.%d.%d",
			verTest1,
			major,
			minor,
			revision,
		)
	}
}

func Test_Unit_Parse_TwoComponents(t *testing.T) {

	major, minor, revision := Parse(verTest2)
	if major != "1" && minor != "2" && revision != "" {
		t.Errorf("Expected %s but saw %d.%d.%d",
			verTest2,
			major,
			minor,
			revision,
		)
	}
}

func Test_Unit_Parse_TooManyComponents(t *testing.T) {

	major, minor, revision := Parse(verTest3)
	if major != "1" && minor != "2" && revision != "3.4" {
		t.Errorf("Expected %s but saw %d.%d.%d",
			verTest3,
			major,
			minor,
			revision,
		)
	}
}

func Test_Unit_Parse_EmptyMinor(t *testing.T) {

	major, minor, revision := Parse(verTest5)
	if major != "1" && minor != "" && revision != "" {
		t.Errorf("Expected %s but saw %d.%d.%d",
			verTest5,
			major,
			minor,
			revision,
		)
	}
}

func Test_Unit_Sha1ToString(t *testing.T) {

	hash := Sha1ToString(sha1TestBytes)

	if hash != sha1TestHash {
		t.Errorf("Expected %s but saw %s", sha1TestHash, hash)
	}
}

func Test_Unit_Sha256ToString(t *testing.T) {

	hash := Sha256ToString(sha256TestBytes)

	if hash != sha256TestHash {
		t.Errorf("Expected %s but saw %s", sha256TestHash, hash)
	}
}
