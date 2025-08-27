package rid

import (
	"crypto/rand"
	"fmt"
	"strings"
	"time"
)

type RId struct {
	block1    string
	block2    string
	phase1    []byte
	phase2    []byte
	phase3    []byte
	uppercase bool
}

func (i RId) Equals(other RId) bool {
	return i.String() == other.String()
}

func (i RId) String() string {
	value := fmt.Sprintf("%s-%s-%04x-%04x-%012x",
		i.block1, i.block2, i.phase1, i.phase2, i.phase3)

	if i.uppercase {
		return strings.ToUpper(value)
	}
	return value
}

func New(prefix string, uppercase bool) *RId {
	now := time.Now()

	// Block1 = prefix + YYMMDD (trim/pad to 8 chars)
	datePart := now.Format("060102")
	block1 := fmt.Sprintf("%s%s", prefix, datePart)
	if len(block1) > 8 {
		block1 = block1[:8]
	} else if len(block1) < 8 {
		block1 = fmt.Sprintf("%-8s", block1) // pad if too short
	}

	// Block2 = time HHMM
	timePart := now.Format("1504")

	// Random filler: 2 + 2 + 6 = 10 bytes
	b := make([]byte, 10)
	_, _ = rand.Read(b)

	phase1 := b[0:2] // 4 hex
	phase2 := b[2:4] // 4 hex
	phase3 := b[4:]  // 12 hex

	return &RId{
		block1:    block1,
		block2:    timePart,
		phase1:    phase1,
		phase2:    phase2,
		phase3:    phase3,
		uppercase: uppercase,
	}
}
