package dectest

import (
    "bufio"
    "fmt"
    "io"
    "strconv"
)

type scanner struct {
    precision   int
	maxExponent int
	minExponent int
	clamp       int
	rounding    RoundingMode
	extended    int
	c           *testCase
	s           *bufio.Scanner
    err error
}

type testCase struct {
    ID         string
	Op         Op
	Inputs     []Data
	Output     Data
	Conditions Condition
}

func NewScanner(r io.Reader) *scanner {
    return &scanner{s: bufio.NewScanner(r)}
}

func (s *scanner) Scan() bool {
    s.c = nil
    for s.c == nil {
        if !s.s.Scan() {
            return false
        }
        s.err = s.parse(s.s.Bytes())
        if s.err != nil {
            return false
        }
        if s.c != nil {
            return true
        }
    }
    return false
}

func (s *scanner) Case() *testCase {
    return s.c
}

func (s *scanner) Err() error {
    return s.err
}

func (s *scanner) parse(data []byte) (err error) {
    cs, p, pe, eof := 0, 0, len(data), len(data)
    var (
        mark int
        ok bool
    )

    %%{
        machine parser;

        action mark { mark = fpc }

        action set_id {
            s.c = &testCase{ID: string(data[mark:fpc])}
        }

        action set_op {
            fmt.Printf("operations(%s) = %s\n", data[mark:fpc], operations[string(data[mark:fpc])])
            if s.c.Op, ok = operations[string(data[mark:fpc])]; !ok {
                return fmt.Errorf("dectest: invalid op: %q", data[mark:fpc])
            }
        }

        action set_precision {
            s.precision, err = strconv.Atoi(string(data[mark:fpc]))
            if err != nil {
                return err
            }
        }

        action set_max_exponent {
            s.maxExponent, err = strconv.Atoi(string(data[mark:fpc]))
            if err != nil {
                return err
            }
        }

        action set_min_exponent {
            s.minExponent, err = strconv.Atoi(string(data[mark:fpc]))
            if err != nil {
                return err
            }
        }

        action set_rounding {
            if s.rounding, ok = roundingModes[string(data[mark:fpc])]; !ok {
                return fmt.Errorf("unknown rounding mode: %s", data[mark:fpc])
            }
        }

        action add_input  { s.c.Inputs = append(s.c.Inputs, Data(data[mark:fpc])) }
        action set_output { s.c.Output = Data(data[mark:fpc]) }
        action add_condition { s.c.Conditions |= conditionFromString(string(data[mark:fpc])) }

        precision = ( digit+ ) >mark %set_precision;
        max_exponent = ( digit+ ) >mark %set_max_exponent;
        min_exponent = ( digit+ ) >mark %set_min_exponent;

        rounding = (
            'ceiling'
            | 'down'
            | 'floor'
            | 'half_down'
            | 'half_even'
            | 'half_up'
            | 'up'
            | '05up'
        ) >mark %set_rounding;

        id = ( alpha{3,} digit{3,} ) >mark %set_id;

        op = (
            'abs'
            | 'add'
            | 'and'
            | 'apply'
            | 'canonical'
            | 'class'
            | 'compare'
            | 'comparesig'
            | 'comparetotal'
            | 'comparetotalmag'
            | 'copy'
            | 'copyabs'
            | 'copynegate'
            | 'copysign'
            | 'divide'
            | 'divideint'
            | 'exp'
            | 'fma'
            | 'invert'
            | 'ln'
            | 'log10'
            | 'logb'
            | 'max'
            | 'min'
            | 'maxmag'
            | 'minmag'
            | 'minus'
            | 'multiply'
            | 'nextminus'
            | 'nextplus'
            | 'nexttoward'
            | 'or'
            | 'plus'
            | 'power'
            | 'quantize'
            | 'reduce'
            | 'remainder'
            | 'remaindernear'
            | 'rescale'
            | 'rotate'
            | 'samequantum'
            | 'scaleb'
            | 'shift'
            | 'squareroot'
            | 'subtract'
            | 'toEng'
            | 'tointegral'
            | 'tointegralx'
            | 'toSci'
            | 'trim'
            | 'xor'
        ) >mark %set_op;

        quote = '\'' | '"';
        sign       = '+' | '-';
		indicator  = 'e' | 'E';
		exponent   = indicator? sign? digit+;
        number     = (digit+ '.' digit* | '.' digit+ | digit+ | digit) exponent?;
		nan_prefix = [sSqQ];
		nan        = (nan_prefix? 'nan'i digit* | '?');
		class      = (nan_prefix? 'nan'i | sign? (
				  'Subnormal'
				| 'Normal'
				| 'Zero'
				| 'Infinity')
		);
        numeric_string = sign? (
			  nan                  # S, Q, NaN, sNaN, ...
            | ('inf'i 'inity'i?)   # +inf, -inf, ...
            | number               # 10, 10.1, +0e-392, ...
        );
        input =  ( numeric_string | '#') >mark %add_input;
        output = ( numeric_string | class | '#') >mark %set_output;

        condition = (
            'Clamped'i
            | 'Conversion_syntax'
            | 'Division_by_zero'
            | 'Division_impossible'
            | 'Division_undefined'
            | 'Inexact'
            | 'Insufficient_storage'
            | 'Invalid_context'
            | 'Invalid_operation'
            | 'Lost_digits'
            | 'Overflow'
            | 'Rounded'
            | 'Subnormal'
            | 'Underflow'
        ) >mark %add_condition;

        main := (
            ('precision:' space+ precision)
            | ('maxexponent:' space+ max_exponent)
            | ('minexponent:' space+ min_exponent)
            | ('rounding:' space+ rounding)
            | (id space+ op space+ (quote? input quote? space+)+ '->' space+ quote? output quote? (space+ condition)*)
        );


        write data;
        write init;
        write exec;
    }%%

    return nil
}