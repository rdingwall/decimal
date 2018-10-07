
//line parser.rl:1
package suite

import (
    "fmt"
)

func ParseCase(data []byte) (c Case, err error) {
    cs, p, pe, eof := 0, 0, len(data), len(data)

    var (
        ok   bool // for mode and op
        mark int
    )

    
//line parser.go:19
const parser_start int = 1
const parser_first_final int = 141
const parser_error int = 0

const parser_en_main int = 1


//line parser.go:27
	{
	cs = parser_start
	}

//line parser.go:32
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 141:
		goto st_case_141
	case 21:
		goto st_case_21
	case 142:
		goto st_case_142
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 143:
		goto st_case_143
	case 93:
		goto st_case_93
	case 144:
		goto st_case_144
	case 94:
		goto st_case_94
	case 145:
		goto st_case_145
	case 146:
		goto st_case_146
	case 95:
		goto st_case_95
	case 96:
		goto st_case_96
	case 147:
		goto st_case_147
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 99:
		goto st_case_99
	case 100:
		goto st_case_100
	case 101:
		goto st_case_101
	case 102:
		goto st_case_102
	case 148:
		goto st_case_148
	case 103:
		goto st_case_103
	case 104:
		goto st_case_104
	case 105:
		goto st_case_105
	case 106:
		goto st_case_106
	case 107:
		goto st_case_107
	case 108:
		goto st_case_108
	case 109:
		goto st_case_109
	case 110:
		goto st_case_110
	case 111:
		goto st_case_111
	case 112:
		goto st_case_112
	case 113:
		goto st_case_113
	case 114:
		goto st_case_114
	case 115:
		goto st_case_115
	case 116:
		goto st_case_116
	case 117:
		goto st_case_117
	case 118:
		goto st_case_118
	case 119:
		goto st_case_119
	case 120:
		goto st_case_120
	case 121:
		goto st_case_121
	case 122:
		goto st_case_122
	case 123:
		goto st_case_123
	case 124:
		goto st_case_124
	case 125:
		goto st_case_125
	case 126:
		goto st_case_126
	case 127:
		goto st_case_127
	case 128:
		goto st_case_128
	case 129:
		goto st_case_129
	case 130:
		goto st_case_130
	case 131:
		goto st_case_131
	case 132:
		goto st_case_132
	case 133:
		goto st_case_133
	case 134:
		goto st_case_134
	case 135:
		goto st_case_135
	case 136:
		goto st_case_136
	case 137:
		goto st_case_137
	case 138:
		goto st_case_138
	case 139:
		goto st_case_139
	case 140:
		goto st_case_140
	}
	goto st_out
	st_case_1:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr0
			}
		case data[p] >= 48:
			goto tr0
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
tr0:
//line parser.rl:18
 mark = p 
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line parser.go:361
		if data[p] == 32 {
			goto tr2
		}
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 122 {
				goto st2
			}
		case data[p] >= 48:
			goto st2
		}
		goto st0
tr2:
//line parser.rl:20

            c.ID = string(data[mark:p])
        
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line parser.go:385
		switch data[p] {
		case 32:
			goto st3
		case 97:
			goto tr5
		case 100:
			goto tr6
		}
		goto st0
tr5:
//line parser.rl:18
 mark = p 
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line parser.go:404
		if data[p] == 100 {
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if data[p] == 100 {
			goto st6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 32 {
			goto tr9
		}
		goto st0
tr9:
//line parser.rl:24

            if c.Op, ok = valToOp[string(data[mark:p])]; !ok {
                return c, fmt.Errorf("invalid op: %q", data[mark:p])
            }
        
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line parser.go:440
		switch data[p] {
		case 32:
			goto st7
		case 39:
			goto tr11
		case 43:
			goto tr12
		case 45:
			goto tr12
		case 46:
			goto tr13
		case 63:
			goto tr15
		case 73:
			goto tr16
		case 78:
			goto tr17
		case 81:
			goto tr18
		case 83:
			goto tr18
		case 105:
			goto tr16
		case 110:
			goto tr17
		case 113:
			goto tr18
		case 115:
			goto tr18
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr14
		}
		goto st0
tr11:
//line parser.rl:18
 mark = p 
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line parser.go:484
		switch data[p] {
		case 43:
			goto st9
		case 45:
			goto st9
		case 46:
			goto st10
		case 63:
			goto st125
		case 73:
			goto st126
		case 78:
			goto st133
		case 81:
			goto st135
		case 83:
			goto st135
		case 105:
			goto st126
		case 110:
			goto st133
		case 113:
			goto st135
		case 115:
			goto st135
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st124
		}
		goto st0
tr12:
//line parser.rl:18
 mark = p 
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line parser.go:524
		switch data[p] {
		case 46:
			goto st10
		case 63:
			goto st125
		case 73:
			goto st126
		case 78:
			goto st133
		case 81:
			goto st135
		case 83:
			goto st135
		case 105:
			goto st126
		case 110:
			goto st133
		case 113:
			goto st135
		case 115:
			goto st135
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st124
		}
		goto st0
tr13:
//line parser.rl:18
 mark = p 
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line parser.go:560
		if 48 <= data[p] && data[p] <= 57 {
			goto st11
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 32:
			goto tr27
		case 39:
			goto st120
		case 43:
			goto st121
		case 45:
			goto st121
		case 69:
			goto st123
		case 101:
			goto st123
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st11
		}
		goto st0
tr27:
//line parser.rl:37
 c.Inputs = append(c.Inputs, Data(data[mark:p])) 
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line parser.go:597
		switch data[p] {
		case 32:
			goto st12
		case 39:
			goto tr32
		case 43:
			goto tr33
		case 45:
			goto tr34
		case 46:
			goto tr35
		case 63:
			goto tr37
		case 73:
			goto tr38
		case 78:
			goto tr39
		case 81:
			goto tr40
		case 83:
			goto tr40
		case 105:
			goto tr38
		case 110:
			goto tr39
		case 113:
			goto tr40
		case 115:
			goto tr40
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr36
		}
		goto st0
tr32:
//line parser.rl:18
 mark = p 
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line parser.go:641
		switch data[p] {
		case 43:
			goto st14
		case 45:
			goto st14
		case 46:
			goto st15
		case 63:
			goto st108
		case 73:
			goto st109
		case 78:
			goto st116
		case 81:
			goto st118
		case 83:
			goto st118
		case 105:
			goto st109
		case 110:
			goto st116
		case 113:
			goto st118
		case 115:
			goto st118
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st107
		}
		goto st0
tr33:
//line parser.rl:18
 mark = p 
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line parser.go:681
		switch data[p] {
		case 46:
			goto st15
		case 63:
			goto st108
		case 73:
			goto st109
		case 78:
			goto st116
		case 81:
			goto st118
		case 83:
			goto st118
		case 105:
			goto st109
		case 110:
			goto st116
		case 113:
			goto st118
		case 115:
			goto st118
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st107
		}
		goto st0
tr35:
//line parser.rl:18
 mark = p 
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line parser.go:717
		if 48 <= data[p] && data[p] <= 57 {
			goto st16
		}
		goto st0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch data[p] {
		case 32:
			goto tr49
		case 39:
			goto st103
		case 43:
			goto st104
		case 45:
			goto st104
		case 69:
			goto st106
		case 101:
			goto st106
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st16
		}
		goto st0
tr49:
//line parser.rl:37
 c.Inputs = append(c.Inputs, Data(data[mark:p])) 
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line parser.go:754
		switch data[p] {
		case 32:
			goto st17
		case 45:
			goto st18
		}
		goto st0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		if data[p] == 62 {
			goto st19
		}
		goto st0
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if data[p] == 32 {
			goto st20
		}
		goto st0
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch data[p] {
		case 32:
			goto st20
		case 35:
			goto tr57
		case 39:
			goto tr58
		case 43:
			goto tr59
		case 45:
			goto tr59
		case 46:
			goto tr60
		case 63:
			goto tr62
		case 73:
			goto tr63
		case 78:
			goto tr64
		case 81:
			goto tr65
		case 83:
			goto tr65
		case 105:
			goto tr63
		case 110:
			goto tr64
		case 113:
			goto tr65
		case 115:
			goto tr65
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr61
		}
		goto st0
tr57:
//line parser.rl:18
 mark = p 
	goto st141
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
//line parser.go:830
		if data[p] == 32 {
			goto tr173
		}
		goto st0
tr173:
//line parser.rl:38
 c.Output = Data(data[mark:p]) 
	goto st21
tr174:
//line parser.rl:39

          c.Excep |= ConditionFromString(string(data[mark:p]))
        
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line parser.go:850
		switch data[p] {
		case 32:
			goto st21
		case 63:
			goto tr67
		case 67:
			goto tr68
		case 68:
			goto tr69
		case 73:
			goto tr70
		case 79:
			goto tr71
		case 82:
			goto tr72
		case 83:
			goto tr73
		case 85:
			goto tr74
		case 109:
			goto tr67
		case 116:
			goto tr67
		case 121:
			goto tr67
		}
		goto st0
tr67:
//line parser.rl:18
 mark = p 
	goto st142
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
//line parser.go:887
		if data[p] == 32 {
			goto tr174
		}
		goto st0
tr68:
//line parser.rl:18
 mark = p 
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line parser.go:901
		if data[p] == 108 {
			goto st23
		}
		goto st0
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		if data[p] == 97 {
			goto st24
		}
		goto st0
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if data[p] == 109 {
			goto st25
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		if data[p] == 112 {
			goto st26
		}
		goto st0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		if data[p] == 101 {
			goto st27
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		if data[p] == 100 {
			goto st142
		}
		goto st0
tr69:
//line parser.rl:18
 mark = p 
	goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line parser.go:960
		if data[p] == 105 {
			goto st29
		}
		goto st0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		if data[p] == 118 {
			goto st30
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		if data[p] == 105 {
			goto st31
		}
		goto st0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		if data[p] == 115 {
			goto st32
		}
		goto st0
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		if data[p] == 105 {
			goto st33
		}
		goto st0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		if data[p] == 111 {
			goto st34
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		if data[p] == 110 {
			goto st35
		}
		goto st0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		if data[p] == 95 {
			goto st36
		}
		goto st0
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		switch data[p] {
		case 98:
			goto st37
		case 117:
			goto st43
		}
		goto st0
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		if data[p] == 121 {
			goto st38
		}
		goto st0
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		if data[p] == 95 {
			goto st39
		}
		goto st0
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		if data[p] == 122 {
			goto st40
		}
		goto st0
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		if data[p] == 101 {
			goto st41
		}
		goto st0
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		if data[p] == 114 {
			goto st42
		}
		goto st0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		if data[p] == 111 {
			goto st142
		}
		goto st0
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		if data[p] == 110 {
			goto st44
		}
		goto st0
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		if data[p] == 100 {
			goto st45
		}
		goto st0
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		if data[p] == 101 {
			goto st46
		}
		goto st0
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		if data[p] == 102 {
			goto st47
		}
		goto st0
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		if data[p] == 105 {
			goto st48
		}
		goto st0
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		if data[p] == 110 {
			goto st26
		}
		goto st0
tr70:
//line parser.rl:18
 mark = p 
	goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
//line parser.go:1157
		if data[p] == 110 {
			goto st50
		}
		goto st0
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		switch data[p] {
		case 101:
			goto st51
		case 118:
			goto st55
		}
		goto st0
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		if data[p] == 120 {
			goto st52
		}
		goto st0
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		if data[p] == 97 {
			goto st53
		}
		goto st0
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		if data[p] == 99 {
			goto st54
		}
		goto st0
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		if data[p] == 116 {
			goto st142
		}
		goto st0
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		if data[p] == 97 {
			goto st56
		}
		goto st0
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		if data[p] == 108 {
			goto st57
		}
		goto st0
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		if data[p] == 105 {
			goto st58
		}
		goto st0
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		if data[p] == 100 {
			goto st59
		}
		goto st0
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		if data[p] == 95 {
			goto st60
		}
		goto st0
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		if data[p] == 111 {
			goto st61
		}
		goto st0
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		if data[p] == 112 {
			goto st62
		}
		goto st0
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		if data[p] == 101 {
			goto st63
		}
		goto st0
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		if data[p] == 114 {
			goto st64
		}
		goto st0
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		if data[p] == 97 {
			goto st65
		}
		goto st0
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		if data[p] == 116 {
			goto st66
		}
		goto st0
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		if data[p] == 105 {
			goto st67
		}
		goto st0
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		if data[p] == 111 {
			goto st68
		}
		goto st0
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		if data[p] == 110 {
			goto st142
		}
		goto st0
tr71:
//line parser.rl:18
 mark = p 
	goto st69
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
//line parser.go:1345
		if data[p] == 118 {
			goto st70
		}
		goto st0
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		if data[p] == 101 {
			goto st71
		}
		goto st0
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		if data[p] == 114 {
			goto st72
		}
		goto st0
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		if data[p] == 102 {
			goto st73
		}
		goto st0
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		if data[p] == 108 {
			goto st74
		}
		goto st0
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		if data[p] == 111 {
			goto st75
		}
		goto st0
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		if data[p] == 119 {
			goto st142
		}
		goto st0
tr72:
//line parser.rl:18
 mark = p 
	goto st76
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
//line parser.go:1413
		if data[p] == 111 {
			goto st77
		}
		goto st0
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		if data[p] == 117 {
			goto st78
		}
		goto st0
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		if data[p] == 110 {
			goto st79
		}
		goto st0
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		if data[p] == 100 {
			goto st26
		}
		goto st0
tr73:
//line parser.rl:18
 mark = p 
	goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
//line parser.go:1454
		if data[p] == 117 {
			goto st81
		}
		goto st0
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		if data[p] == 98 {
			goto st82
		}
		goto st0
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		if data[p] == 110 {
			goto st83
		}
		goto st0
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		if data[p] == 111 {
			goto st84
		}
		goto st0
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		if data[p] == 114 {
			goto st85
		}
		goto st0
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		if data[p] == 109 {
			goto st86
		}
		goto st0
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		if data[p] == 97 {
			goto st87
		}
		goto st0
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		if data[p] == 108 {
			goto st142
		}
		goto st0
tr74:
//line parser.rl:18
 mark = p 
	goto st88
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
//line parser.go:1531
		if data[p] == 110 {
			goto st89
		}
		goto st0
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		if data[p] == 100 {
			goto st70
		}
		goto st0
tr58:
//line parser.rl:18
 mark = p 
	goto st90
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
//line parser.go:1554
		switch data[p] {
		case 43:
			goto st91
		case 45:
			goto st91
		case 46:
			goto st92
		case 63:
			goto st146
		case 73:
			goto st95
		case 78:
			goto st101
		case 81:
			goto st148
		case 83:
			goto st148
		case 105:
			goto st95
		case 110:
			goto st101
		case 113:
			goto st148
		case 115:
			goto st148
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st145
		}
		goto st0
tr59:
//line parser.rl:18
 mark = p 
	goto st91
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
//line parser.go:1594
		switch data[p] {
		case 46:
			goto st92
		case 63:
			goto st146
		case 73:
			goto st95
		case 78:
			goto st101
		case 81:
			goto st148
		case 83:
			goto st148
		case 105:
			goto st95
		case 110:
			goto st101
		case 113:
			goto st148
		case 115:
			goto st148
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st145
		}
		goto st0
tr60:
//line parser.rl:18
 mark = p 
	goto st92
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
//line parser.go:1630
		if 48 <= data[p] && data[p] <= 57 {
			goto st143
		}
		goto st0
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
		switch data[p] {
		case 32:
			goto tr173
		case 39:
			goto st141
		case 43:
			goto st93
		case 45:
			goto st93
		case 69:
			goto st94
		case 101:
			goto st94
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st143
		}
		goto st0
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		if 48 <= data[p] && data[p] <= 57 {
			goto st144
		}
		goto st0
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
		switch data[p] {
		case 32:
			goto tr173
		case 39:
			goto st141
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st144
		}
		goto st0
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		switch data[p] {
		case 43:
			goto st93
		case 45:
			goto st93
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st144
		}
		goto st0
tr61:
//line parser.rl:18
 mark = p 
	goto st145
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
//line parser.go:1706
		switch data[p] {
		case 32:
			goto tr173
		case 39:
			goto st141
		case 43:
			goto st93
		case 45:
			goto st93
		case 46:
			goto st143
		case 69:
			goto st94
		case 101:
			goto st94
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st145
		}
		goto st0
tr62:
//line parser.rl:18
 mark = p 
	goto st146
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
//line parser.go:1736
		switch data[p] {
		case 32:
			goto tr173
		case 39:
			goto st141
		}
		goto st0
tr63:
//line parser.rl:18
 mark = p 
	goto st95
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
//line parser.go:1753
		switch data[p] {
		case 78:
			goto st96
		case 110:
			goto st96
		}
		goto st0
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		switch data[p] {
		case 70:
			goto st147
		case 102:
			goto st147
		}
		goto st0
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
		switch data[p] {
		case 32:
			goto tr173
		case 39:
			goto st141
		case 73:
			goto st97
		case 105:
			goto st97
		}
		goto st0
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		switch data[p] {
		case 78:
			goto st98
		case 110:
			goto st98
		}
		goto st0
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		switch data[p] {
		case 73:
			goto st99
		case 105:
			goto st99
		}
		goto st0
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		switch data[p] {
		case 84:
			goto st100
		case 116:
			goto st100
		}
		goto st0
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		switch data[p] {
		case 89:
			goto st146
		case 121:
			goto st146
		}
		goto st0
tr64:
//line parser.rl:18
 mark = p 
	goto st101
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
//line parser.go:1846
		switch data[p] {
		case 65:
			goto st102
		case 97:
			goto st102
		}
		goto st0
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		switch data[p] {
		case 78:
			goto st146
		case 110:
			goto st146
		}
		goto st0
tr65:
//line parser.rl:18
 mark = p 
	goto st148
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
//line parser.go:1875
		switch data[p] {
		case 32:
			goto tr173
		case 39:
			goto st141
		case 78:
			goto st101
		case 110:
			goto st101
		}
		goto st0
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		if data[p] == 32 {
			goto tr49
		}
		goto st0
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
		if 48 <= data[p] && data[p] <= 57 {
			goto st105
		}
		goto st0
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		switch data[p] {
		case 32:
			goto tr49
		case 39:
			goto st103
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st105
		}
		goto st0
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		switch data[p] {
		case 43:
			goto st104
		case 45:
			goto st104
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st105
		}
		goto st0
tr36:
//line parser.rl:18
 mark = p 
	goto st107
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
//line parser.go:1944
		switch data[p] {
		case 32:
			goto tr49
		case 39:
			goto st103
		case 43:
			goto st104
		case 45:
			goto st104
		case 46:
			goto st16
		case 69:
			goto st106
		case 101:
			goto st106
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st107
		}
		goto st0
tr37:
//line parser.rl:18
 mark = p 
	goto st108
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
//line parser.go:1974
		switch data[p] {
		case 32:
			goto tr49
		case 39:
			goto st103
		}
		goto st0
tr38:
//line parser.rl:18
 mark = p 
	goto st109
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
//line parser.go:1991
		switch data[p] {
		case 78:
			goto st110
		case 110:
			goto st110
		}
		goto st0
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		switch data[p] {
		case 70:
			goto st111
		case 102:
			goto st111
		}
		goto st0
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
		switch data[p] {
		case 32:
			goto tr49
		case 39:
			goto st103
		case 73:
			goto st112
		case 105:
			goto st112
		}
		goto st0
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		switch data[p] {
		case 78:
			goto st113
		case 110:
			goto st113
		}
		goto st0
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		switch data[p] {
		case 73:
			goto st114
		case 105:
			goto st114
		}
		goto st0
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		switch data[p] {
		case 84:
			goto st115
		case 116:
			goto st115
		}
		goto st0
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		switch data[p] {
		case 89:
			goto st108
		case 121:
			goto st108
		}
		goto st0
tr39:
//line parser.rl:18
 mark = p 
	goto st116
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
//line parser.go:2084
		switch data[p] {
		case 65:
			goto st117
		case 97:
			goto st117
		}
		goto st0
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		switch data[p] {
		case 78:
			goto st108
		case 110:
			goto st108
		}
		goto st0
tr40:
//line parser.rl:18
 mark = p 
	goto st118
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
//line parser.go:2113
		switch data[p] {
		case 32:
			goto tr49
		case 39:
			goto st103
		case 78:
			goto st116
		case 110:
			goto st116
		}
		goto st0
tr34:
//line parser.rl:18
 mark = p 
	goto st119
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
//line parser.go:2134
		switch data[p] {
		case 46:
			goto st15
		case 62:
			goto st19
		case 63:
			goto st108
		case 73:
			goto st109
		case 78:
			goto st116
		case 81:
			goto st118
		case 83:
			goto st118
		case 105:
			goto st109
		case 110:
			goto st116
		case 113:
			goto st118
		case 115:
			goto st118
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st107
		}
		goto st0
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		if data[p] == 32 {
			goto tr27
		}
		goto st0
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		if 48 <= data[p] && data[p] <= 57 {
			goto st122
		}
		goto st0
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		switch data[p] {
		case 32:
			goto tr27
		case 39:
			goto st120
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st122
		}
		goto st0
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		switch data[p] {
		case 43:
			goto st121
		case 45:
			goto st121
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st122
		}
		goto st0
tr14:
//line parser.rl:18
 mark = p 
	goto st124
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
//line parser.go:2220
		switch data[p] {
		case 32:
			goto tr27
		case 39:
			goto st120
		case 43:
			goto st121
		case 45:
			goto st121
		case 46:
			goto st11
		case 69:
			goto st123
		case 101:
			goto st123
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st124
		}
		goto st0
tr15:
//line parser.rl:18
 mark = p 
	goto st125
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
//line parser.go:2250
		switch data[p] {
		case 32:
			goto tr27
		case 39:
			goto st120
		}
		goto st0
tr16:
//line parser.rl:18
 mark = p 
	goto st126
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
//line parser.go:2267
		switch data[p] {
		case 78:
			goto st127
		case 110:
			goto st127
		}
		goto st0
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
		switch data[p] {
		case 70:
			goto st128
		case 102:
			goto st128
		}
		goto st0
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		switch data[p] {
		case 32:
			goto tr27
		case 39:
			goto st120
		case 73:
			goto st129
		case 105:
			goto st129
		}
		goto st0
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		switch data[p] {
		case 78:
			goto st130
		case 110:
			goto st130
		}
		goto st0
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		switch data[p] {
		case 73:
			goto st131
		case 105:
			goto st131
		}
		goto st0
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
		switch data[p] {
		case 84:
			goto st132
		case 116:
			goto st132
		}
		goto st0
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		switch data[p] {
		case 89:
			goto st125
		case 121:
			goto st125
		}
		goto st0
tr17:
//line parser.rl:18
 mark = p 
	goto st133
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
//line parser.go:2360
		switch data[p] {
		case 65:
			goto st134
		case 97:
			goto st134
		}
		goto st0
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
		switch data[p] {
		case 78:
			goto st125
		case 110:
			goto st125
		}
		goto st0
tr18:
//line parser.rl:18
 mark = p 
	goto st135
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
//line parser.go:2389
		switch data[p] {
		case 32:
			goto tr27
		case 39:
			goto st120
		case 78:
			goto st133
		case 110:
			goto st133
		}
		goto st0
tr6:
//line parser.rl:18
 mark = p 
	goto st136
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
//line parser.go:2410
		if data[p] == 105 {
			goto st137
		}
		goto st0
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
		if data[p] == 118 {
			goto st138
		}
		goto st0
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
		if data[p] == 105 {
			goto st139
		}
		goto st0
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
		if data[p] == 100 {
			goto st140
		}
		goto st0
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
		if data[p] == 101 {
			goto st6
		}
		goto st0
	st_out:
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof141: cs = 141; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof142: cs = 142; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
	_test_eof59: cs = 59; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
	_test_eof62: cs = 62; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof65: cs = 65; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof
	_test_eof84: cs = 84; goto _test_eof
	_test_eof85: cs = 85; goto _test_eof
	_test_eof86: cs = 86; goto _test_eof
	_test_eof87: cs = 87; goto _test_eof
	_test_eof88: cs = 88; goto _test_eof
	_test_eof89: cs = 89; goto _test_eof
	_test_eof90: cs = 90; goto _test_eof
	_test_eof91: cs = 91; goto _test_eof
	_test_eof92: cs = 92; goto _test_eof
	_test_eof143: cs = 143; goto _test_eof
	_test_eof93: cs = 93; goto _test_eof
	_test_eof144: cs = 144; goto _test_eof
	_test_eof94: cs = 94; goto _test_eof
	_test_eof145: cs = 145; goto _test_eof
	_test_eof146: cs = 146; goto _test_eof
	_test_eof95: cs = 95; goto _test_eof
	_test_eof96: cs = 96; goto _test_eof
	_test_eof147: cs = 147; goto _test_eof
	_test_eof97: cs = 97; goto _test_eof
	_test_eof98: cs = 98; goto _test_eof
	_test_eof99: cs = 99; goto _test_eof
	_test_eof100: cs = 100; goto _test_eof
	_test_eof101: cs = 101; goto _test_eof
	_test_eof102: cs = 102; goto _test_eof
	_test_eof148: cs = 148; goto _test_eof
	_test_eof103: cs = 103; goto _test_eof
	_test_eof104: cs = 104; goto _test_eof
	_test_eof105: cs = 105; goto _test_eof
	_test_eof106: cs = 106; goto _test_eof
	_test_eof107: cs = 107; goto _test_eof
	_test_eof108: cs = 108; goto _test_eof
	_test_eof109: cs = 109; goto _test_eof
	_test_eof110: cs = 110; goto _test_eof
	_test_eof111: cs = 111; goto _test_eof
	_test_eof112: cs = 112; goto _test_eof
	_test_eof113: cs = 113; goto _test_eof
	_test_eof114: cs = 114; goto _test_eof
	_test_eof115: cs = 115; goto _test_eof
	_test_eof116: cs = 116; goto _test_eof
	_test_eof117: cs = 117; goto _test_eof
	_test_eof118: cs = 118; goto _test_eof
	_test_eof119: cs = 119; goto _test_eof
	_test_eof120: cs = 120; goto _test_eof
	_test_eof121: cs = 121; goto _test_eof
	_test_eof122: cs = 122; goto _test_eof
	_test_eof123: cs = 123; goto _test_eof
	_test_eof124: cs = 124; goto _test_eof
	_test_eof125: cs = 125; goto _test_eof
	_test_eof126: cs = 126; goto _test_eof
	_test_eof127: cs = 127; goto _test_eof
	_test_eof128: cs = 128; goto _test_eof
	_test_eof129: cs = 129; goto _test_eof
	_test_eof130: cs = 130; goto _test_eof
	_test_eof131: cs = 131; goto _test_eof
	_test_eof132: cs = 132; goto _test_eof
	_test_eof133: cs = 133; goto _test_eof
	_test_eof134: cs = 134; goto _test_eof
	_test_eof135: cs = 135; goto _test_eof
	_test_eof136: cs = 136; goto _test_eof
	_test_eof137: cs = 137; goto _test_eof
	_test_eof138: cs = 138; goto _test_eof
	_test_eof139: cs = 139; goto _test_eof
	_test_eof140: cs = 140; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 141, 143, 144, 145, 146, 147, 148:
//line parser.rl:38
 c.Output = Data(data[mark:p]) 
		case 142:
//line parser.rl:39

          c.Excep |= ConditionFromString(string(data[mark:p]))
        
//line parser.go:2611
		}
	}

	_out: {}
	}

//line parser.rl:107

    return c, nil
}
