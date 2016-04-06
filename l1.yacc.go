//line l1.y:2
package l1compiler

import __yyfmt__ "fmt"

//line l1.y:2
import (
	"fmt"
	"strconv"
)

//line l1.y:11
type yySymType struct {
	yys  int
	n    int
	s    string
	node Node
}

const LABEL = 57346
const GOLABEL = 57347
const NEG = 57348
const NAT = 57349
const NAT6 = 57350
const NAT8 = 57351
const AOP = 57352
const SOP = 57353
const CMP = 57354
const CALL = 57355
const CJUMP = 57356
const TAILCALL = 57357
const RETURN = 57358
const GOTO = 57359
const LPAREN = 57360
const RPAREN = 57361
const ALLOCATE = 57362
const PRINT = 57363
const ARRAYERROR = 57364
const ASSIGN = 57365
const MEM = 57366
const RSP = 57367
const RCX = 57368
const X = 57369
const W = 57370
const A = 57371

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LABEL",
	"GOLABEL",
	"NEG",
	"NAT",
	"NAT6",
	"NAT8",
	"AOP",
	"SOP",
	"CMP",
	"CALL",
	"CJUMP",
	"TAILCALL",
	"RETURN",
	"GOTO",
	"LPAREN",
	"RPAREN",
	"ALLOCATE",
	"PRINT",
	"ARRAYERROR",
	"ASSIGN",
	"MEM",
	"RSP",
	"RCX",
	"X",
	"W",
	"A",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line l1.y:246
func cast(y yyLexer) *Compiler { return y.(*Lexer).p }

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 87,
	19, 40,
	-2, 38,
	-1, 88,
	19, 41,
	-2, 39,
}

const yyNprod = 55
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 150

var yyAct = [...]int{

	67, 84, 68, 81, 71, 76, 50, 69, 83, 48,
	64, 66, 37, 48, 109, 73, 72, 74, 75, 51,
	70, 45, 49, 108, 97, 60, 59, 47, 44, 36,
	48, 51, 58, 45, 49, 51, 70, 45, 49, 41,
	39, 42, 43, 38, 47, 24, 80, 80, 82, 62,
	63, 57, 51, 56, 45, 49, 48, 55, 73, 72,
	74, 75, 61, 90, 87, 89, 93, 94, 89, 88,
	91, 92, 95, 86, 85, 48, 54, 53, 51, 70,
	45, 49, 73, 72, 74, 75, 103, 73, 72, 74,
	75, 78, 77, 79, 22, 98, 52, 51, 105, 45,
	49, 33, 51, 70, 45, 49, 7, 51, 6, 25,
	26, 27, 28, 29, 30, 31, 32, 2, 34, 19,
	21, 20, 16, 17, 18, 13, 14, 15, 10, 11,
	12, 107, 106, 102, 101, 100, 99, 4, 104, 96,
	65, 9, 3, 8, 46, 40, 35, 23, 5, 1,
}
var yyPact = [...]int{

	99, -1000, 138, 90, 87, 90, 137, -1000, -1000, 121,
	118, 115, 112, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 82, 27, 26, 77, 58, 57, 38, 34,
	32, 13, 7, -1000, -1000, 6, 39, -13, 136, 76,
	-1000, 71, 5, -1000, -1000, -1000, -1000, -16, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 9, 76, 81, 52, -1000, 135, 12, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 88, 128, 127, 126,
	-1000, -1000, 125, -7, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 134, 76, -1000, -1000,
	-1000, -1000, -1000, 123, -1000, -1000, 4, -5, -1000, -1000,
}
var yyPgo = [...]int{

	0, 149, 137, 148, 94, 147, 146, 12, 11, 145,
	1, 4, 5, 0, 2, 144, 6, 7, 3,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 4, 5, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	9, 9, 9, 8, 7, 7, 12, 12, 13, 13,
	10, 10, 10, 14, 14, 11, 11, 15, 15, 16,
	17, 17, 17, 17, 18,
}
var yyR2 = [...]int{

	0, 4, 1, 2, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 1, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 4, 1, 3, 3, 1, 1,
	3, 3, 3, 3, 5, 5, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, 18, 4, -2, -3, 18, 19, -2, 4,
	7, 8, 9, 7, 8, 9, 7, 8, 9, 7,
	9, 8, -4, -5, 18, -4, -4, -4, -4, -4,
	-4, -4, -4, 19, -4, -6, -11, -7, 17, 14,
	-9, 13, 15, 16, -18, 28, -15, 18, 4, 29,
	-16, 26, 19, 19, 19, 19, 19, 19, 19, 19,
	19, 23, 10, 11, 23, 4, -8, -13, -14, -17,
	27, -11, 7, 6, 8, 9, -12, 21, 20, 22,
	-11, -18, -12, 24, -10, -7, -8, -14, -17, -18,
	-13, -16, -17, -10, -14, -17, 4, 12, 7, 8,
	8, 8, 8, -14, 4, -13, 9, 8, 19, 19,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 2, 0, 1, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 13, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 4, 14, 0, 0, 0, 0, 0,
	25, 0, 0, 28, 29, 45, 46, 0, 54, 47,
	48, 49, 6, 9, 5, 7, 11, 8, 10, 12,
	15, 0, 0, 0, 0, 23, 0, 0, 38, 39,
	43, 44, 50, 51, 52, 53, 0, 0, 0, 0,
	36, 37, 0, 0, 16, 17, 22, -2, -2, 42,
	19, 20, 21, 18, 40, 41, 0, 0, 26, 30,
	31, 32, 27, 0, 24, 33, 0, 0, 34, 35,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line l1.y:31
		{
			fmt.Println("Detected program: %+v", yyDollar[2].s)
			yyVAL.node = newProgramNode(yyDollar[2].s, yyDollar[3].node)
			cast(yylex).SetAstRoot(yyVAL.node)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:38
		{
			fmt.Println("Detected end subprogram")
			yyVAL.node = newSubProgramNode(yyDollar[1].node, nil)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line l1.y:43
		{
			fmt.Printf("Detected subprogram\n")
			yyVAL.node = newSubProgramNode(yyDollar[1].node, yyDollar[2].node)
		}
	case 4:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line l1.y:50
		{
			fmt.Println("Detected func: %+v", yyDollar[2].s)
			yyVAL.node = newFunctionNode(yyDollar[2].s, yyDollar[3].n, yyDollar[4].n, yyDollar[5].node)
		}
	case 5:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line l1.y:55
		{
			//fmt.Printf("Detected func: %q\n", $2)
			yyVAL.node = newFunctionNode(yyDollar[2].s, yyDollar[3].n, yyDollar[4].n, yyDollar[5].node)
		}
	case 6:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line l1.y:60
		{
			//fmt.Printf("Detected func: %q\n", $2)
			yyVAL.node = newFunctionNode(yyDollar[2].s, yyDollar[3].n, yyDollar[4].n, yyDollar[5].node)
		}
	case 7:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line l1.y:65
		{
			//fmt.Printf("Detected func: %q\n", $2)
			yyVAL.node = newFunctionNode(yyDollar[2].s, yyDollar[3].n, yyDollar[4].n, yyDollar[5].node)
		}
	case 8:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line l1.y:70
		{
			//fmt.Printf("Detected func: %q\n", $2)
			yyVAL.node = newFunctionNode(yyDollar[2].s, yyDollar[3].n, yyDollar[4].n, yyDollar[5].node)
		}
	case 9:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line l1.y:75
		{
			//fmt.Printf("Detected func: %q\n", $2)
			yyVAL.node = newFunctionNode(yyDollar[2].s, yyDollar[3].n, yyDollar[4].n, yyDollar[5].node)
		}
	case 10:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line l1.y:80
		{
			//fmt.Printf("Detected func: %q\n", $2)
			yyVAL.node = newFunctionNode(yyDollar[2].s, yyDollar[3].n, yyDollar[4].n, yyDollar[5].node)
		}
	case 11:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line l1.y:85
		{
			//fmt.Printf("Detected func: %q\n", $2)
			yyVAL.node = newFunctionNode(yyDollar[2].s, yyDollar[3].n, yyDollar[4].n, yyDollar[5].node)
		}
	case 12:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line l1.y:90
		{
			//fmt.Printf("Detected func: %q\n", $2)
			yyVAL.node = newFunctionNode(yyDollar[2].s, yyDollar[3].n, yyDollar[4].n, yyDollar[5].node)
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:96
		{
			fmt.Printf("Detected end instruction\n")
			yyVAL.node = newInstructionNode(yyDollar[1].node, nil)
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line l1.y:101
		{
			fmt.Printf("Detected instruction\n")
			yyVAL.node = newInstructionNode(yyDollar[1].node, yyDollar[2].node)
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:108
		{
			yyVAL.node = yyDollar[2].node
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:114
		{
			//fmt.Printf("Detected assign \n")
			yyVAL.node = newAssignNode(yyDollar[1].node, yyDollar[3].node)
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:119
		{
			//fmt.Printf("Detected assign \n")
			yyVAL.node = newAssignNode(yyDollar[1].node, yyDollar[3].node)
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:124
		{
			//fmt.Printf("Detected assign \n")
			yyVAL.node = newAssignNode(yyDollar[1].node, yyDollar[3].node)
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:129
		{
			//fmt.Printf("Detected aop \n")
			yyVAL.node = newOpNode(yyDollar[2].s, yyDollar[1].node, yyDollar[3].node)
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:134
		{
			//fmt.Printf("Detected sop \n")
			yyVAL.node = newOpNode(yyDollar[2].s, yyDollar[1].node, yyDollar[3].node)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:139
		{
			//fmt.Printf("Detected sop \n")
			yyVAL.node = newOpNode(yyDollar[2].s, yyDollar[1].node, yyDollar[3].node)
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:144
		{
			yyVAL.node = newAssignNode(yyDollar[1].node, yyDollar[3].node)
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line l1.y:148
		{
			//fmt.Printf("Detected goto: %q\n", $2)
			yyVAL.node = newGotoNode(yyDollar[2].s)
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line l1.y:153
		{
			yyVAL.node = newCjumpNode(yyDollar[2].node, yyDollar[3].s, yyDollar[4].s)
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:157
		{
			yyVAL.node = yyDollar[1].node
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:161
		{
			yyVAL.node = newCallNode(yyDollar[2].node, yyDollar[3].n)
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:165
		{
			yyVAL.node = newTailcallNode(yyDollar[2].node, yyDollar[3].n)
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:169
		{
			yyVAL.node = newReturnNode()
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:173
		{
			yyVAL.node = yyDollar[1].node
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:180
		{
			yyVAL.node = newSysCallNode(yyDollar[2].s, 1)
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:184
		{
			yyVAL.node = newSysCallNode(yyDollar[2].s, 2)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:188
		{
			yyVAL.node = newSysCallNode(yyDollar[2].s, 3)
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:194
		{
			yyVAL.node = newCmpopNode(yyDollar[2].s, yyDollar[1].node, yyDollar[3].node)
		}
	case 34:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line l1.y:200
		{
			yyVAL.node = newMemNode(yyDollar[3].node, yyDollar[4].n)
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line l1.y:204
		{
			yyVAL.node = newMemNode(yyDollar[3].node, yyDollar[4].n)
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:209
		{
			yyVAL.node = yyDollar[1].node
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:210
		{
			yyVAL.node = yyDollar[1].node
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:213
		{
			yyVAL.node = yyDollar[1].node
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:214
		{
			yyVAL.node = yyDollar[1].node
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:217
		{
			yyVAL.node = yyDollar[1].node
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:218
		{
			yyVAL.node = yyDollar[1].node
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:219
		{
			yyVAL.node = yyDollar[1].node
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:221
		{
			yyVAL.node = newTokenNode(yyDollar[1].s)
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:222
		{
			yyVAL.node = yyDollar[1].node
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:225
		{
			yyVAL.node = newTokenNode(yyDollar[1].s)
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:226
		{
			yyVAL.node = yyDollar[1].node
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:229
		{
			yyVAL.node = newTokenNode(yyDollar[1].s)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:230
		{
			yyVAL.node = yyDollar[1].node
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:233
		{
			yyVAL.node = newTokenNode(yyDollar[1].s)
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:236
		{
			yyVAL.node = newTokenNode(strconv.Itoa(yyDollar[1].n))
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:237
		{
			yyVAL.node = newTokenNode(strconv.Itoa(yyDollar[1].n))
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:238
		{
			fmt.Println("Yacc got NAT6")
			yyVAL.node = newTokenNode(strconv.Itoa(yyDollar[1].n))
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:239
		{
			yyVAL.node = newTokenNode(strconv.Itoa(yyDollar[1].n))
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:243
		{
			yyVAL.node = newLabelNode(yyDollar[1].s)
		}
	}
	goto yystack /* stack new state and value */
}
