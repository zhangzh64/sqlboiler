package randomize

import (
	"crypto/md5"
	"fmt"
	"math/rand"
)

const alphabetAll = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphabetLowerAlpha = "abcdefghijklmnopqrstuvwxyz"

func randStr(s *Seed, ln int) string {
	str := make([]byte, ln)
	for i := 0; i < ln; i++ {
		str[i] = byte(alphabetAll[s.NextInt()%len(alphabetAll)])
	}

	return string(str)
}

func randByteSlice(s *Seed, ln int) []byte {
	str := make([]byte, ln)
	for i := 0; i < ln; i++ {
		str[i] = byte(s.NextInt() % 256)
	}

	return str
}

func randPoint(s *Seed) string {
	a := s.NextInt() % 100
	b := a + 1
	return fmt.Sprintf("(%d,%d)", a, b)
}

func randBox(s *Seed) string {
	a := s.NextInt() % 100
	b := a + 1
	c := a + 2
	d := a + 3
	return fmt.Sprintf("(%d,%d),(%d,%d)", a, b, c, d)
}

func randCircle(s *Seed) string {
	a, b, c := s.NextInt()%100, s.NextInt()%100, s.NextInt()%100
	return fmt.Sprintf("((%d,%d),%d)", a, b, c)
}

func randNetAddr(s *Seed) string {
	return fmt.Sprintf(
		"%d.%d.%d.%d",
		s.NextInt()%254+1,
		s.NextInt()%254+1,
		s.NextInt()%254+1,
		s.NextInt()%254+1,
	)
}

func randMacAddr(s *Seed) string {
	buf := make([]byte, 6)
	for i := range buf {
		buf[i] = byte(s.NextInt())
	}

	// Set the local bit
	buf[0] |= 2
	return fmt.Sprintf(
		"%02x:%02x:%02x:%02x:%02x:%02x",
		buf[0], buf[1], buf[2], buf[3], buf[4], buf[5],
	)
}

func randLsn(s *Seed) string {
	a := s.NextInt() % 9000000
	b := s.NextInt() % 9000000
	return fmt.Sprintf("%d/%d", a, b)
}

func randTxID(s *Seed) string {
	// Order of integers is relevant
	a := s.NextInt()%200 + 100
	b := a + 100
	c := a
	d := a + 50
	return fmt.Sprintf("%d:%d:%d,%d", a, b, c, d)
}

func randMoney(s *Seed) string {
	return fmt.Sprintf("%d.00", s.NextInt())
}

// StableDBName takes a database name in, and generates
// a random string using the database name as the rand Seed.
// getDBNameHash is used to generate unique test database names.
func StableDBName(input string) string {
	return randStrFromSource(stableSource(input), 40)
}

// stableSource takes an input value, and produces a random
// seed from it that will produce very few collisions in
// a 40 character random string made from a different alphabet.
func stableSource(input string) *rand.Rand {
	sum := md5.Sum([]byte(input))
	var seed int64
	for i, byt := range sum {
		seed ^= int64(byt) << uint((i*4)%64)
	}
	return rand.New(rand.NewSource(seed))
}

func randStrFromSource(r *rand.Rand, length int) string {
	ln := len(alphabetLowerAlpha)

	output := make([]rune, length)
	for i := 0; i < length; i++ {
		output[i] = rune(alphabetLowerAlpha[r.Intn(ln)])
	}

	return string(output)
}
