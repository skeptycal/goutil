package goalgo

import (
	// . "crypto/rand"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type (
	Container interface {
		Build() error // Build builds sequence from items in X, given an iterable X
		Len() int     // Len returns return the number of stored items
	}

	DataSource interface {
		Set
		Sequence
	}
)

var (
	Now     = time.Now().UTC()
	YearNow = Now.Year()
)

// CaseFunc performs string manipulation on an input string and returns an output string
var CaseFunc = func(s string) string { return strings.ToUpper(s[:1]) + strings.ToLower(s[1:]) }
var minTime = time.Date(YearNow-30, 1, 1, 0, 0, 0, 0, time.UTC)
var maxTime = minTime.AddDate(YearNow-18, 0, 0)

// PastDay returns the time.Time representing the datetime that
// is in the past by the number of years/months/days
func PastDay(year, month, day int) time.Time { return Now.AddDate(-year, -month, -day).UTC() }

func RandomString(min, max int, caseFunc func(string) string) string {
	const offset byte = 'a'
	var letter byte

	diff := max - min
	length := rand.Intn(diff) + min
	b := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		letter = byte(rand.Intn(26)) + offset
		b = append(b, letter)
	}

	name := string(b)

	if caseFunc != nil {
		name = caseFunc(name)
	}

	return name
}

func RandomDate(min, max time.Time) time.Time {
	diff := max.Unix() - min.Unix()
	birthsecond := rand.Int63n(diff)
	return minTime.Add(time.Duration(birthsecond) * time.Second)
}
func RandomStudent() *student {
	name := RandomString(4, 12, CaseFunc)
	birthday := RandomDate(minTime, maxTime)

	return &student{name: name, birthday: birthday}
}

func GenerateRoster(n int) *roster {
	l := make([]*student, 0, n)
	for i := 0; i < n; i++ {
		l = append(l, RandomStudent())
	}

	return &roster{list: l}
}

func (r *roster) SameBirthday() bool {
	for _, person := range r.list {
		iDay := person.Birthday().YearDay()
		for _, person2 := range r.list {
			jDay := person2.Birthday().YearDay()
			if jDay == iDay {
				return true
			}
		}
	}
	return false
}

type (
	student struct {
		name     string
		birthday time.Time
	}

	// roster maintains a list of students and implements
	// sort.Interface
	roster struct {
		list []*student
	}
)

func (s *student) Name() string          { return s.name }
func (s *student) Age() time.Duration    { return time.Since(s.birthday) }
func (s *student) Birthday() time.Time   { return s.birthday }
func (s *student) Unix() int64           { return s.birthday.Unix() }
func (s *student) Month() time.Month     { return s.birthday.Month() }
func (s *student) Weekday() time.Weekday { return s.birthday.Weekday() }
func (s *student) Day() int              { return s.birthday.Day() }
func (s *student) Year() int             { return s.birthday.Year() }
func (s *student) Hour() int             { return s.birthday.Hour() }
func (s *student) Minute() int           { return s.birthday.Minute() }

func (s *roster) Less(i, j int) bool { return s.list[i].birthday.Unix() < s.list[j].birthday.Unix() }
func (s *roster) Len() int           { return len(s.list) }
func (s *roster) Swap(i, j int)      { s.list[i], s.list[j] = s.list[j], s.list[i] }
func (r *roster) String() string {
	sb := strings.Builder{}
	defer sb.Reset()

	for i, student := range r.list {
		s := fmt.Sprintf("%3d %-20s %02d/%02d/%4d\n", i, student.Name(), student.Month(), student.Day(), student.Year())
		sb.WriteString(s)
	}
	return sb.String()
}
