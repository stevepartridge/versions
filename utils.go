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
	switch len(parts) {
	case 1:
		return parts[0], "", ""
	case 2:
		return parts[0], parts[1], ""
	case 3:
		return parts[0], parts[1], parts[2]
	}

	return parts[0], parts[1], strings.Join(parts[2:], "-")
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
