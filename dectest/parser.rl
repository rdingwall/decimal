package dectest

import (
    "fmt"
)

func ParseCase(data []byte) (c Case, err error) {
    cs, p, pe, eof := 0, 0, len(data), len(data)

    var (
        ok   bool // for mode and op
        mark int
    )

    %%{
        machine parser;

        action mark { mark = fpc }

        action set_id {
            c.ID = string(data[mark:fpc])
        }

        action set_op {
            if c.Op, ok = valToOp[string(data[mark:fpc])]; !ok {
                return c, fmt.Errorf("invalid op: %q", data[mark:fpc])
            }
        }
        action set_mode {
	    	if c.Mode, ok = valToMode[string(data[mark:fpc])]; !ok {
				return c, fmt.Errorf("invalid mode: %q", data[mark:fpc])
	    	}
        }
        action set_trap {
			c.Trap = ConditionFromString(string(data[mark:fpc]))
        }
        action add_input  { c.Inputs = append(c.Inputs, Data(data[mark:fpc])) }
        action set_output { c.Output = Data(data[mark:fpc]) }
        action add_excep  {
          c.Excep |= ConditionFromString(string(data[mark:fpc]))
        }

       op = [a-z0-9]+ >mark %set_op;

        condition = (
              'Inexact' # Inexact
            | 'Underflow' # Underflow
            | 'Rounded' # Underflow
            | 'Overflow' # Overflow
            | 'Division_by_zero' # DivisionByZero
            | 'Invalid_operation' # InvalidOperation

			# Custom
			| 'Clamped' # Clamped
			| 'Rounded' # Rounded
			| 'y' # ConversionSyntax
			| 'm' # DivisionImpossible
			| 'Division_undefined' # DivisionUndefined
			| 't' # InsufficientStorage
			| '?' # InvalidContext
			| 'Subnormal' # Subnormal
        );
        excep = condition >mark %add_excep;

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
        quote = '\'' | '"';
        input =  ( numeric_string | '#') >mark %add_input;
        output = ( numeric_string | class | '#') >mark %set_output;

        id = ([a-z0-9]+) >mark %set_id;

        main := (
            (id ' '+) # A short name which identifies the test
            (op ' '+) # A case-independent keyword which describes the operation
            (quote? input quote? ' '+)+ # The first (or only) operand required for the operation.
            '->' ' '+
            quote? output quote?
             (' '+ excep)*
        );

        write data;
        write init;
        write exec;
    }%%
    return c, nil
}
