package dectest

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"math/bits"
	"strings"
)

var roundingModes = map[string]big.RoundingMode{
	"ceiling": big.ToPositiveInf,

	// (Round toward 0; truncate.) The discarded digits are ignored; the
	// result is unchanged.
	"down": big.ToZero,

	// (Round toward -∞.) If all of the discarded digits are zero or if the sign
	// is 0 the result is unchanged. Otherwise, the sign is 1 and the result
	// coefficient should be incremented by 1.
	"floor": big.ToNegativeInf,

	// If the discarded digits represent greater than half (0.5) the value of a
	// one in the next left position then the result coefficient should be
	// incremented by 1 (rounded up). If they represent less than half, then
	// the result coefficient is not adjusted (that is, the discarded digits are
	// ignored).
	"half_even": big.ToNearestEven,

	// If the discarded digits represent greater than or equal to half (0.5) of
	// the value of a one in the next left position then the result coefficient
	// should be incremented by 1 (rounded up). Otherwise the discarded
	// digits are ignored.
	"half_up": big.ToNearestAway,
}

// ParseCases returns a slice of test cases in .decTest form read from r.
func ParseCases(r io.Reader) (cases []Case, err error) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	var precision int
	var mode big.RoundingMode
	var maxscale, minscale int
	var clamp bool
	var modeUnsupported bool

	for s.Scan() {
		p := s.Bytes()
		// Skip empty lines and comments.
		if len(p) == 0 || p[0] == '-' {
			continue
		}

		fmt.Println(strings.ToLower(string(p)))

		// TODO: ragel-ify: move this state to a scanner.Next() struct
		var prec int
		if n, _ := fmt.Sscanf(strings.ToLower(string(p)), "precision: %d", &prec); n == 1 {
			fmt.Printf("😃 Change prec '%d'\n", prec)
			precision = prec
		}

		var mxs int
		if n, _ := fmt.Sscanf(strings.ToLower(string(p)), "maxexponent: %d", &mxs); n == 1 {
			fmt.Printf("😃 Change max scale '%d'\n", mxs)
			maxscale = mxs
		}

		var mns int
		if n, _ := fmt.Sscanf(strings.ToLower(string(p)), "minexponent: %d", &mns); n == 1 {
			fmt.Printf("😃 Change min scale '%d'\n", mns)
			minscale = mns
		}

		var cl int
		if n, _ := fmt.Sscanf(strings.ToLower(string(p)), "clamp: %d", &cl); n == 1 {
			fmt.Printf("😃 Change clamp mode '%d' --> %v\n", cl, clamp)
			if cl == 1 {
				clamp = true
			} else {
				clamp = false
			}
		}

		var rm string
		if n, _ := fmt.Sscanf(strings.ToLower(string(p)), "rounding: %s", &rm); n == 1 {
			r, ok := roundingModes[rm]
			fmt.Printf("😃 Change rounding mode '%s' --> %s\n", rm, r)
			if !ok {
				// unsupported rounding mode
				// @TODO: model these unsupported rounding modes and explicitly t.Skip() them
				modeUnsupported = true
			} else {
				mode = r
				modeUnsupported = false
			}
		}

		c, err := ParseCase(p)
		if err != nil {
			return nil, err
		}

		if c.ID == "" || modeUnsupported {
			continue
		}

		c.Prec = precision
		c.Mode = mode
		c.MaxScale = maxscale
		c.MinScale = minscale
		if clamp {
			c.Trap |= Clamped
		}
		cases = append(cases, c)
	}
	return cases, s.Err()
}

type Case struct {
	Prefix   string
	Prec     int
	Op       Op
	Mode     big.RoundingMode
	Trap     Condition
	Inputs   []Data
	Output   Data
	Excep    Condition
	ID       string
	MaxScale int
	MinScale int
}

// TODO(eric): String should print the same format as the input

func (c Case) String() string {
	return fmt.Sprintf("%s%d [%s, %s]: %s(%s) = %s %s",
		c.Prefix, c.Prec, c.Trap, c.Mode, c.Op,
		join(c.Inputs, ", ", -1), c.Output, c.Excep)
}

// ShortString returns the same as String, except long data values are capped at
// length digits.
func (c Case) ShortString(length int) string {
	return fmt.Sprintf("%s%d [%s, %s]: %s(%s) = %s %s",
		c.Prefix, c.Prec, c.Trap, c.Mode, c.Op,
		join(c.Inputs, ", ", length), trunc(c.Output, length), c.Excep)
}

func trunc(d Data, l int) Data {
	if l <= 0 || l >= len(d) {
		return d
	}
	return d[:l/2] + "..." + d[len(d)-(l/2):]
}

func join(a []Data, sep Data, l int) Data {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return trunc(a[0], l)
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(trunc(a[i], l))
	}

	b := make([]byte, n)
	bp := copy(b, trunc(a[0], l))
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], trunc(s, l))
	}
	return Data(b)
}

// Data is input or output from a test case.
type Data string

// NoData is output when the operation throws some sort of Condition and does
// not "return" any data.
const NoData Data = "?"

func (i Data) TrimQuotes() Data {
	return Data(strings.Trim(string(i), "'"))
}

// IsNaN returns two booleans indicating whether the data is a NaN value and
// whether it's signaling or not.
func (i Data) IsNaN() (nan, signal bool) {
	i = i.TrimQuotes()
	if len(i) == 1 {
		return (i == "S" || i == "Q"), i == "S"
	}
	if i[0] == '-' {
		i = i[1:]
	}
	return strings.EqualFold(string(i), "nan") ||
		strings.EqualFold(string(i), "qnan") ||
		strings.EqualFold(string(i), "snan"), i[0] == 's' || i[0] == 'S'
}

// IsInf returns a boolean indicating whether the data is an Infinity and an
// int indicating the signedness of the Infinity.
func (i Data) IsInf() (int, bool) {
	i = i.TrimQuotes()
	if len(i) != 4 {
		return 0, false
	}
	if strings.EqualFold(string(i), "-Inf") {
		return -1, true
	}
	if strings.EqualFold(string(i), "+Inf") {
		return +1, true
	}
	return 0, false
}

// Condition is a bitmask value raised after or during specific operations.
type Condition uint32

const (
	Clamped Condition = 1 << iota
	ConversionSyntax
	DivisionByZero
	DivisionImpossible
	DivisionUndefined
	Inexact
	InsufficientStorage
	InvalidContext
	InvalidOperation
	Overflow
	Rounded
	Subnormal
	Underflow
)

func (c Condition) String() string {
	if c == 0 {
		return "NoConditions"
	}

	// Each condition is one bit, so this saves some allocations.
	a := make([]string, 0, bits.OnesCount32(uint32(c)))
	for i := Condition(1); c != 0; i <<= 1 {
		if c&i == 0 {
			continue
		}
		switch c ^= i; i {
		case Clamped:
			a = append(a, "clamped")
		case ConversionSyntax:
			a = append(a, "conversion syntax")
		case DivisionByZero:
			a = append(a, "division by zero")
		case DivisionImpossible:
			a = append(a, "division impossible")
		case Inexact:
			a = append(a, "inexact")
		case InsufficientStorage:
			a = append(a, "insufficient storage")
		case InvalidContext:
			a = append(a, "invalid context")
		case InvalidOperation:
			a = append(a, "invalid operation")
		case Overflow:
			a = append(a, "overflow")
		case Rounded:
			a = append(a, "rounded")
		case Subnormal:
			a = append(a, "subnormal")
		case Underflow:
			a = append(a, "underflow")
		default:
			a = append(a, fmt.Sprintf("unknown(%d)", i))
		}
	}
	return strings.Join(a, ", ")
}

func ConditionFromString(s string) (r Condition) {
	var ok bool
	r, ok = valToCondition[s]
	if !ok {
		panic(fmt.Errorf("unknown condition '%s'", s))
	}
	return
}

var valToCondition = map[string]Condition{
	"Inexact":           Inexact,
	"Underflow":         Underflow, // tininess and 'extraordinary' error
	"Overflow":          Overflow,
	"Division_by_zero":  DivisionByZero,
	"Invalid_operation": InvalidOperation,

	// custom
	"Clamped":              Clamped,
	"Rounded":              Rounded,
	"Conversion_syntax":    ConversionSyntax,
	"Division_impossible":  DivisionImpossible,
	"Division_undefined":   DivisionUndefined,
	"Insufficient_storage": InsufficientStorage,
	"Invalid_context":      InvalidContext,
	"Subnormal":            Subnormal,
}

var valToMode = map[string]big.RoundingMode{
	">":  big.ToPositiveInf,
	"<":  big.ToNegativeInf,
	"0":  big.ToZero,
	"=0": big.ToNearestEven,
	"=^": big.ToNearestAway,
	"^":  big.AwayFromZero,
}

// Op is a specific operation the test case must perform.
type Op uint8

const (
	Add Op = iota // add
	Div
	Sub // subtract
	Apply
)

var valToOp = map[string]Op{
	"add":      Add,
	"divide":   Div,
	"subtract": Sub,
	"apply":    Apply,
}

//go:generate stringer -type=Op
