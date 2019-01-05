package dectest

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/ericlagergren/decimal"
)

// Helper returns testing.T.Helper, if it exists.
func Helper(v interface{}) func() {
	if fn, ok := v.(interface {
		Helper()
	}); ok {
		return fn.Helper
	}
	return func() {}
}

var skipIt = map[string]struct{}{
	/* NULL reference, decimal16, decimal32, or decimal128 */
	"add900":     struct{}{},
	"add901":     struct{}{},
	"absx900":    struct{}{},
	"addx9990":   struct{}{},
	"addx9991":   struct{}{},
	"clam090":    struct{}{},
	"clam091":    struct{}{},
	"clam092":    struct{}{},
	"clam093":    struct{}{},
	"clam094":    struct{}{},
	"clam095":    struct{}{},
	"clam096":    struct{}{},
	"clam097":    struct{}{},
	"clam098":    struct{}{},
	"clam099":    struct{}{},
	"clam189":    struct{}{},
	"clam190":    struct{}{},
	"clam191":    struct{}{},
	"clam192":    struct{}{},
	"clam193":    struct{}{},
	"clam194":    struct{}{},
	"clam195":    struct{}{},
	"clam196":    struct{}{},
	"clam197":    struct{}{},
	"clam198":    struct{}{},
	"clam199":    struct{}{},
	"comx990":    struct{}{},
	"comx991":    struct{}{},
	"cotx9990":   struct{}{},
	"cotx9991":   struct{}{},
	"ctmx9990":   struct{}{},
	"ctmx9991":   struct{}{},
	"ddabs900":   struct{}{},
	"ddadd9990":  struct{}{},
	"ddadd9991":  struct{}{},
	"ddcom9990":  struct{}{},
	"ddcom9991":  struct{}{},
	"ddcot9990":  struct{}{},
	"ddcot9991":  struct{}{},
	"ddctm9990":  struct{}{},
	"ddctm9991":  struct{}{},
	"dddiv9998":  struct{}{},
	"dddiv9999":  struct{}{},
	"dddvi900":   struct{}{},
	"dddvi901":   struct{}{},
	"ddfma2990":  struct{}{},
	"ddfma2991":  struct{}{},
	"ddfma39990": struct{}{},
	"ddfma39991": struct{}{},
	"ddlogb900":  struct{}{},
	"ddmax900":   struct{}{},
	"ddmax901":   struct{}{},
	"ddmxg900":   struct{}{},
	"ddmxg901":   struct{}{},
	"ddmin900":   struct{}{},
	"ddmin901":   struct{}{},
	"ddmng900":   struct{}{},
	"ddmng901":   struct{}{},
	"ddmul9990":  struct{}{},
	"ddmul9991":  struct{}{},
	"ddnextm900": struct{}{},
	"ddnextp900": struct{}{},
	"ddnextt900": struct{}{},
	"ddnextt901": struct{}{},
	"ddqua998":   struct{}{},
	"ddqua999":   struct{}{},
	"ddred900":   struct{}{},
	"ddrem1000":  struct{}{},
	"ddrem1001":  struct{}{},
	"ddrmn1000":  struct{}{},
	"ddrmn1001":  struct{}{},
	"ddsub9990":  struct{}{},
	"ddsub9991":  struct{}{},
	"ddintx074":  struct{}{},
	"ddintx094":  struct{}{},
	"divx9998":   struct{}{},
	"divx9999":   struct{}{},
	"dvix900":    struct{}{},
	"dvix901":    struct{}{},
	"dqabs900":   struct{}{},
	"dqadd9990":  struct{}{},
	"dqadd9991":  struct{}{},
	"dqcom990":   struct{}{},
	"dqcom991":   struct{}{},
	"dqcot9990":  struct{}{},
	"dqcot9991":  struct{}{},
	"dqctm9990":  struct{}{},
	"dqctm9991":  struct{}{},
	"dqdiv9998":  struct{}{},
	"dqdiv9999":  struct{}{},
	"dqdvi900":   struct{}{},
	"dqdvi901":   struct{}{},
	"dqfma2990":  struct{}{},
	"dqfma2991":  struct{}{},
	"dqadd39990": struct{}{},
	"dqadd39991": struct{}{},
	"dqlogb900":  struct{}{},
	"dqmax900":   struct{}{},
	"dqmax901":   struct{}{},
	"dqmxg900":   struct{}{},
	"dqmxg901":   struct{}{},
	"dqmin900":   struct{}{},
	"dqmin901":   struct{}{},
	"dqmng900":   struct{}{},
	"dqmng901":   struct{}{},
	"dqmul9990":  struct{}{},
	"dqmul9991":  struct{}{},
	"dqnextm900": struct{}{},
	"dqnextp900": struct{}{},
	"dqnextt900": struct{}{},
	"dqnextt901": struct{}{},
	"dqqua998":   struct{}{},
	"dqqua999":   struct{}{},
	"dqred900":   struct{}{},
	"dqrem1000":  struct{}{},
	"dqrem1001":  struct{}{},
	"dqrmn1000":  struct{}{},
	"dqrmn1001":  struct{}{},
	"dqsub9990":  struct{}{},
	"dqsub9991":  struct{}{},
	"dqintx074":  struct{}{},
	"dqintx094":  struct{}{},
	"expx900":    struct{}{},
	"fmax2990":   struct{}{},
	"fmax2991":   struct{}{},
	"fmax39990":  struct{}{},
	"fmax39991":  struct{}{},
	"lnx900":     struct{}{},
	"logx900":    struct{}{},
	"logbx900":   struct{}{},
	"maxx900":    struct{}{},
	"maxx901":    struct{}{},
	"mxgx900":    struct{}{},
	"mxgx901":    struct{}{},
	"mnm900":     struct{}{},
	"mnm901":     struct{}{},
	"mng900":     struct{}{},
	"mng901":     struct{}{},
	"minx900":    struct{}{},
	"mulx990":    struct{}{},
	"mulx991":    struct{}{},
	"nextm900":   struct{}{},
	"nextp900":   struct{}{},
	"nextt900":   struct{}{},
	"nextt901":   struct{}{},
	"plu900":     struct{}{},
	"powx900":    struct{}{},
	"powx901":    struct{}{},
	"pwsx900":    struct{}{},
	"quax1022":   struct{}{},
	"quax1023":   struct{}{},
	"quax1024":   struct{}{},
	"quax1025":   struct{}{},
	"quax1026":   struct{}{},
	"quax1027":   struct{}{},
	"quax1028":   struct{}{},
	"quax1029":   struct{}{},
	"quax0a2":    struct{}{},
	"quax0a3":    struct{}{},
	"quax998":    struct{}{},
	"quax999":    struct{}{},
	"redx900":    struct{}{},
	"remx1000":   struct{}{},
	"remx1001":   struct{}{},
	"rmnx900":    struct{}{},
	"rmnx901":    struct{}{},
	"sqtx9900":   struct{}{},
	"subx9990":   struct{}{},
	"subx9991":   struct{}{},
	/* operand range violations, invalid context */
	"expx901":  struct{}{},
	"expx902":  struct{}{},
	"expx903":  struct{}{},
	"expx905":  struct{}{},
	"lnx901":   struct{}{},
	"lnx902":   struct{}{},
	"lnx903":   struct{}{},
	"lnx905":   struct{}{},
	"logx901":  struct{}{},
	"logx902":  struct{}{},
	"logx903":  struct{}{},
	"logx905":  struct{}{},
	"powx1183": struct{}{},
	"powx1184": struct{}{},
	"powx4001": struct{}{},
	"powx4002": struct{}{},
	"powx4003": struct{}{},
	"powx4005": struct{}{},
	"powx4008": struct{}{},
	"powx4010": struct{}{},
	"powx4012": struct{}{},
	"powx4014": struct{}{},
	"scbx164":  struct{}{},
	"scbx165":  struct{}{},
	"scbx166":  struct{}{},
	/* skipped for decNumber, too */
	"powx4302": struct{}{},
	"powx4303": struct{}{},
	"powx4342": struct{}{},
	"powx4343": struct{}{},
	"pwsx805":  struct{}{},
	/* disagreement for three arg power */
	"pwmx325": struct{}{},
	"pwmx326": struct{}{},
}

func Test(t *testing.T, file string) {
	//t.Parallel() // Call after parsing so we don't goof the scanner.
	s := open(file)
	for s.Next() {

		c := s.Case(t)
		t.Run(c.c.ID, func(t *testing.T) {

			if _, ok := skipIt[c.c.ID]; ok {
				t.SkipNow()
			}

			c.t = t

			//fmt.Println(c.c.String())
			//	t.Logf("%#v\n", c.c)
			c.execute()
		})
	}
}

var nilary = map[Op]func(z *decimal.Big) *decimal.Big{}

var unary = map[Op]func(z, x *decimal.Big) *decimal.Big{}

var binary = map[Op]func(z, x, y *decimal.Big) *decimal.Big{
	Add: (*decimal.Big).Add,
}

var ternary = map[Op]func(z, x, y, u *decimal.Big) *decimal.Big{}

func (c *scase) execute() {
	if nfn, ok := nilary[c.c.Op]; ok {
		c.Check(nfn(c.x))
	} else if ufn, ok := unary[c.c.Op]; ok {
		c.Check(ufn(c.z, c.x))
	} else if bfn, ok := binary[c.c.Op]; ok {
		if c.z == nil || c.x == nil || c.y == nil {
			//fmt.Println("🤯")
			c.t.Fatalf("input was nil: %v", []*decimal.Big{c.z, c.x, c.y})
		}
		c.Check(bfn(c.z, c.x, c.y))
	} else if tfn, ok := ternary[c.c.Op]; ok {
		c.Check(tfn(c.z, c.x, c.y, c.u))
	} else {

		panic("unknown op: " + c.c.Op.String())
		/*
					switch c.c.Op {

					case Class:
						c.Assert(c.x.Class(), c.r)
					case Cmp:
						rv := c.x.Cmp(c.y)
						r, _, snan := c.Cmp()
						c.Assert(rv, r)
						c.Assert(snan, c.x.Context.Conditions&decimal.InvalidOperation != 0)
					case Shift:
						//v, _ := c.y.Int64()
						//c.Check(misc.Shift(c.z, c.x, int(v)))
					case Quant:
						v, _ := c.y.Int64()
						c.Check(c.x.Quantize(int(v)))
					case CTR:
						r := new(big.Rat).SetFrac(c.x.Int(nil), c.y.Int(nil))
						// Given that SetRat/Rat are non-standard, I don't feel bad for
						// calling Assert(z.Cmp(r)) instead of Check(z).
						c.Assert(c.z.SetRat(r).Cmp(c.R()), 0)
					case Sign:
						c.Assert(c.x.Sign(), c.Sign())
					case CTS, CFS:
						xs := c.x.String()
						if !decimal.Regexp.MatchString(xs) {
							c.t.Fatalf("should match regexp: %q", xs)
						}
						c.Assert(xs, c.r)
					case Pow:
						math.Pow(c.z, c.x, c.y)
						r := c.R()
						if !equal(c.z, r) {
							diff := new(decimal.Big)
							eps := decimal.New(1, c.c.Prec)
							ctx := decimal.Context{Precision: -c.c.Prec}
							if ctx.Sub(diff, r, c.z).CmpAbs(eps) > 0 {
								c.t.Logf(`#%d: %s
			wanted: %q (%s:%d)
			got   : %q (%s:%d)
			`,
									c.i, c.c.ShortString(22),
									r, c.flags, -r.Scale(),
									c.z, c.z.Context.Conditions, -c.z.Scale(),
								)
							}
						}
					case Signbit:
						c.Assert(c.x.Signbit(), c.Signbit())
					default:
						panic("unknown test: " + name)
					}
		*/
	}
}

func open(fpath string) (c *scanner) {
	file, err := os.Open(fpath)
	if err != nil {
		panic(err)
	}
	/*gzr, err := gzip.NewReader(file)
	if err != nil {
		panic(err)
	}*/
	return &scanner{
		//s:     bufio.NewScanner(gzr),
		s: bufio.NewScanner(file),
		close: func() {
			//gzr.Close();
			file.Close()
		},
	}
}

type scanner struct {
	i     int
	s     *bufio.Scanner
	close func()
}

func (c *scanner) Next() bool {
	if !c.s.Scan() {
		c.close()
		return false
	}
	c.i++
	return true
}

func (c *scanner) Case(t *testing.T) *scase {
	cs, err := ParseCase(c.s.Bytes())
	if err != nil {
		panic(err)
	}
	if strings.HasPrefix(cs.ID, "addx") {
		cs.Prec = 9
		cs.MaxScale = 384
		cs.MinScale = -384
		cs.Mode = big.ToPositiveInf

		switch {
		case cs.ID >= "addx046" && cs.ID <= "add051":
			cs.Prec = 15
		case cs.ID >= "add060" && cs.ID <= "add077":
			cs.Prec = 6

		case cs.ID >= "add161" && cs.ID <= "add183":
			cs.Prec = 15
		}
		//cs.Trap = ^(Inexact | Rounded | Subnormal)
	} else if strings.HasPrefix(cs.ID, "add") {
		cs.Prec = 9
		cs.MaxScale = 384
		cs.MinScale = -384
		cs.Mode = big.ToNearestEven
		//cs.Trap = ^(Inexact | Rounded | Subnormal)
	} else if strings.HasPrefix(cs.ID, "dddiv") {
		cs.Prec = 16
		cs.MaxScale = 384
		cs.MinScale = -384
		cs.Mode = big.ToNearestEven
		//	cs.Trap = ^(Inexact | Rounded | Subnormal)
	}
	return parse(t, cs, c.i)
}

func ctx(c Case) decimal.Context {
	return decimal.Context{
		Precision:     c.Prec,
		OperatingMode: decimal.GDA,
		RoundingMode:  decimal.RoundingMode(c.Mode),
		Traps:         decimal.Condition(c.Trap),
		MinScale:      c.MinScale,
		MaxScale:      c.MaxScale,
	}
}

func parse(t *testing.T, c Case, i int) *scase {
	ctx := ctx(c)
	s := scase{
		t:     t,
		ctx:   ctx,
		i:     i,
		c:     c,
		z:     decimal.WithContext(ctx),
		r:     string(c.Output.TrimQuotes()),
		flags: decimal.Condition(c.Excep),
	}
	if c.ID == "addx257" {
		t.Logf("😨 %#v", c)
	}
	switch len(c.Inputs) {
	case 3:
		s.u, _ = decimal.WithContext(ctx).SetString(string(c.Inputs[2].TrimQuotes()))
		fallthrough
	case 2:
		s.y, _ = decimal.WithContext(ctx).SetString(string(c.Inputs[1].TrimQuotes()))
		fallthrough
	case 1:
		s.x, _ = decimal.WithContext(ctx).SetString(string(c.Inputs[0].TrimQuotes()))
	default:
		t.Logf("%s\n%d inputs: %#v", s.c, len(c.Inputs), c)
	}
	return &s
}

func (c *scase) Assert(a, b interface{}) {
	Helper(c.t)()
	if !reflect.DeepEqual(a, b) {
		c.t.Fatalf(`#%d: %s
wanted: %v
got   : %v
`, c.i, c.c.ShortString(22), b, a)
	}
}

func (c *scase) Check(z *decimal.Big) {
	Helper(c.t)()
	r := c.R()
	if !equal(z, r) {
		str := fmt.Sprintf(`#%d: %s
wanted: %q (%s:%d)
got   : %q (%s:%d)
`,
			c.i, c.c.ShortString(10000),
			r, c.flags, -r.Scale(),
			z, z.Context.Conditions, -z.Scale(),
		)
		c.t.Error(str)
	}
}

type scase struct {
	z, x, y, u *decimal.Big
	c          Case
	i          int
	r          string
	t          *testing.T
	flags      decimal.Condition
	ctx        decimal.Context
}

func (s *scase) R() *decimal.Big {
	if Data(s.r) == NoData {
		return decimal.New(0, 0).SetInf(false)
	}
	r, ok := decimal.WithContext(s.ctx).SetString(s.r)
	if !ok {
		s.t.Fatalf("SetString(%s) returned failure", s.r)
	}
	r.Context.Conditions = s.flags
	return r
}

func (s *scase) Str() string { return s.r }

func (s *scase) Sign() int {
	r, err := strconv.Atoi(s.r)
	if err != nil {
		Helper(s.t)()
		s.t.Fatal(err)
	}
	return r
}

func (s *scase) Cmp() (int, bool, bool) {
	qnan, snan := Data(s.r).IsNaN()
	if qnan || snan {
		return 0, qnan, snan
	}
	r, err := strconv.Atoi(s.r)
	if err != nil {
		Helper(s.t)()
		s.t.Fatal(err)
	}
	return r, false, false
}

func (s *scase) Signbit() bool {
	r, err := strconv.ParseBool(s.r)
	if err != nil {
		Helper(s.t)()
		s.t.Fatal(err)
	}
	return r
}

func equal(x, y *decimal.Big) bool {
	if x.Signbit() != y.Signbit() {
		return false
	}
	if x.IsFinite() != y.IsFinite() {
		return false
	}
	if !x.IsFinite() {
		return (x.IsInf(0) && y.IsInf(0)) || (x.IsNaN(0) && y.IsNaN(0))
	}
	// Python doesn't have DivisionUndefined.
	if (x.Context.Conditions & ^decimal.DivisionUndefined) != y.Context.Conditions {
		return false
	}
	cmp := x.Cmp(y) == 0
	scl := x.Scale() == y.Scale()
	prec := x.Precision() == y.Precision()
	return cmp && scl && prec
}
