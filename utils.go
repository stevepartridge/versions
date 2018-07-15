package versions

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func Parse(ver string) (string, string, string) {
	if ver == "" {
		return "", "", ""
	}

	parts := strings.Split(ver, ".")
	if len(parts) == 1 {
		return parts[0], "", ""
	} else if len(parts) == 2 {
		return parts[0], parts[1], ""
	} else if len(parts) == 3 {
		return parts[0], parts[1], parts[2]
	} else if len(parts) > 3 {
		return parts[0], parts[1], strings.Join(parts[2:], "-")
	} else {
		return "", "", ""
	}

}

func Sha1ToString(in []byte) string {

	hasher := sha1.New()
	hasher.Write(in)
	out := hasher.Sum(nil)

	return hex.EncodeToString(out)
}

func Sha256ToString(in []byte) string {

	hasher := sha256.New()
	hasher.Write(in)
	out := hasher.Sum(nil)

	return hex.EncodeToString(out)
}
