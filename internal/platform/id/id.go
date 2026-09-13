package id

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func New(prefix string) string {
	entropy := ulid.Monotonic(rand.Reader, 0)

	value := ulid.MustNew(
		ulid.Timestamp(time.Now()),
		entropy,
	)

	return prefix + "_" + value.String()
}
