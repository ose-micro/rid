package rid

import (
	"crypto/rand"
	"fmt"
	"strings"
	"time"
)

type Id struct {
	date      string
	time      string
	phase1    string
	phase2    string
	phase3    string
	uppercase bool
}

func (i Id) Equals(other Id) bool {
	return i.String() == other.String()
}

func (i Id) String() string {
	value := fmt.Sprintf("%s-%s-%04x-%04x-%012x",
		i.date, i.time, i.phase1, i.phase2, i.phase3)

	if i.uppercase {
		return strings.ToUpper(value)
	}
	return value
}

func New(prefix string, uppercase bool) *Id {
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

	phase1 := string(b[0:2]) // 4 hex
	phase2 := string(b[2:4]) // 4 hex
	phase3 := string(b[4:])  // 12 hex

	return &Id{
		date:      block1,
		time:      timePart,
		phase1:    phase1,
		phase2:    phase2,
		phase3:    phase3,
		uppercase: uppercase,
	}
}

func Existing(id string) *Id {
	values := strings.Split(id, "-")

	datePart := values[0]
	timePart := values[1]
	phase1 := values[2]
	phase2 := values[3]
	phase3 := values[4]

	uppercase := id == strings.ToUpper(id)

	return &Id{
		date:      datePart,
		time:      timePart,
		phase1:    phase1,
		phase2:    phase2,
		phase3:    phase3,
		uppercase: uppercase,
	}
}
