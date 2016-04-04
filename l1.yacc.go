//line l1.y:2
package L1

import __yyfmt__ "fmt"

//line l1.y:2
import (
	//"fmt"
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
	"'1'",
	"'2'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line l1.y:203
func cast(y yyLexer) *Compiler { return y.(*Lexer).p }

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 61,
	19, 31,
	-2, 29,
	-1, 62,
	19, 32,
	-2, 30,
}

const yyNprod = 46
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 122

var yyAct = [...]int{

	41, 58, 42, 55, 75, 74, 30, 43, 47, 46,
	48, 49, 32, 47, 46, 48, 49, 73, 26, 57,
	29, 44, 33, 50, 27, 31, 38, 44, 33, 40,
	27, 31, 44, 33, 19, 27, 31, 64, 61, 63,
	67, 68, 63, 62, 30, 66, 69, 81, 56, 34,
	65, 30, 15, 23, 21, 24, 25, 20, 29, 30,
	77, 47, 46, 48, 49, 60, 33, 45, 27, 31,
	59, 30, 79, 33, 14, 27, 31, 7, 6, 2,
	44, 33, 18, 27, 31, 71, 80, 52, 51, 53,
	12, 54, 54, 33, 76, 27, 31, 47, 46, 48,
	49, 72, 36, 37, 16, 11, 10, 4, 78, 70,
	39, 9, 3, 8, 28, 35, 22, 33, 17, 13,
	5, 1,
}
var yyPact = [...]int{

	61, -1000, 108, 60, 58, 60, 107, -1000, -1000, 99,
	98, 56, 33, 56, 40, -1000, -1000, 30, 92, 3,
	106, 7, -1000, 67, 47, -1000, -1000, -1000, -1000, -5,
	-1000, -1000, -1000, -1000, -1000, 2, 7, 91, 55, -1000,
	105, 73, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	94, -13, -26, -27, -1000, -1000, 86, -4, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	104, 7, -1000, -1000, -1000, -1000, -1000, 77, -1000, -1000,
	28, -1000,
}
var yyPgo = [...]int{

	0, 121, 107, 120, 90, 119, 118, 34, 29, 116,
	1, 67, 23, 0, 2, 114, 12, 7, 3,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 3, 4, 4, 5, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 9, 9, 9, 8, 7, 12, 12, 13,
	13, 10, 10, 10, 14, 14, 11, 11, 15, 15,
	16, 17, 17, 17, 17, 18,
}
var yyR2 = [...]int{

	0, 4, 1, 2, 6, 1, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 4, 1, 3, 3,
	1, 1, 3, 3, 3, 3, 5, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, 18, 4, -2, -3, 18, 19, -2, 4,
	7, 7, -4, -5, 18, 19, -4, -6, -11, -7,
	17, 14, -9, 13, 15, 16, -18, 28, -15, 18,
	4, 29, -16, 26, 19, 23, 10, 11, 23, 4,
	-8, -13, -14, -17, 25, -11, 7, 6, 8, 9,
	-12, 21, 20, 22, -11, -18, -12, 24, -10, -7,
	-8, -14, -17, -18, -13, -16, -17, -10, -14, -17,
	4, 12, 7, 30, 31, 31, 8, -14, 4, -13,
	9, 19,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 2, 0, 1, 3, 0,
	0, 0, 0, 5, 0, 4, 6, 0, 0, 0,
	0, 0, 17, 0, 0, 20, 21, 36, 37, 0,
	45, 38, 39, 40, 7, 0, 0, 0, 0, 15,
	0, 0, 29, 30, 34, 35, 41, 42, 43, 44,
	0, 0, 0, 0, 27, 28, 0, 0, 8, 9,
	14, -2, -2, 33, 11, 12, 13, 10, 31, 32,
	0, 0, 18, 22, 23, 24, 19, 0, 16, 25,
	0, 26,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 30,
	31,
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
			//fmt.Printf("Detected program: %q\n", $2)
			yyVAL.node = newProgramNode(yyDollar[2].s, yyDollar[3].node)
			cast(yylex).SetAstRoot(yyVAL.node)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:38
		{
			//fmt.Printf("Detected end subprogram\n")
			yyVAL.node = newSubProgramNode(yyDollar[1].node, nil)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line l1.y:43
		{
			//fmt.Printf("Detected subprogram\n")
			yyVAL.node = newSubProgramNode(yyDollar[1].node, yyDollar[2].node)
		}
	case 4:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line l1.y:50
		{
			//fmt.Printf("Detected func: %q\n", $2)
			yyVAL.node = newFunctionNode(yyDollar[2].s, yyDollar[3].n, yyDollar[4].n, yyDollar[5].node)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:57
		{
			//fmt.Printf("Detected end instruction\n")
			yyVAL.node = newInstructionNode(yyDollar[1].node, nil)
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line l1.y:62
		{
			//fmt.Printf("Detected instruction\n")
			yyVAL.node = newInstructionNode(yyDollar[1].node, yyDollar[2].node)
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:69
		{
			yyVAL.node = yyDollar[2].node
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:75
		{
			//fmt.Printf("Detected assign \n")
			yyVAL.node = newAssignNode(yyDollar[1].node, yyDollar[3].node)
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:80
		{
			//fmt.Printf("Detected assign \n")
			yyVAL.node = newAssignNode(yyDollar[1].node, yyDollar[3].node)
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:85
		{
			//fmt.Printf("Detected assign \n")
			yyVAL.node = newAssignNode(yyDollar[1].node, yyDollar[3].node)
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:90
		{
			//fmt.Printf("Detected aop \n")
			yyVAL.node = newOpNode(yyDollar[2].s, yyDollar[1].node, yyDollar[3].node)
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:95
		{
			//fmt.Printf("Detected sop \n")
			yyVAL.node = newOpNode(yyDollar[2].s, yyDollar[1].node, yyDollar[3].node)
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:100
		{
			//fmt.Printf("Detected sop \n")
			yyVAL.node = newOpNode(yyDollar[2].s, yyDollar[1].node, yyDollar[3].node)
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:105
		{
			yyVAL.node = newAssignNode(yyDollar[1].node, yyDollar[3].node)
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line l1.y:109
		{
			//fmt.Printf("Detected goto: %q\n", $2)
			yyVAL.node = newGotoNode(yyDollar[2].s)
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line l1.y:114
		{
			yyVAL.node = newCjumpNode(yyDollar[2].node, yyDollar[3].s, yyDollar[4].s)
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:118
		{
			yyVAL.node = yyDollar[1].node
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:122
		{
			yyVAL.node = newCallNode(yyDollar[2].node, yyDollar[3].n)
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:126
		{
			yyVAL.node = newTailcallNode(yyDollar[2].node, yyDollar[3].n)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:130
		{
			yyVAL.node = newReturnNode()
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:134
		{
			yyVAL.node = yyDollar[1].node
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:141
		{
			yyVAL.node = newSysCallNode(yyDollar[2].s, 1)
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:145
		{
			yyVAL.node = newSysCallNode(yyDollar[2].s, 2)
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:149
		{
			yyVAL.node = newSysCallNode(yyDollar[2].s, 3)
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:155
		{
			yyVAL.node = newCmpopNode(yyDollar[2].s, yyDollar[1].node, yyDollar[3].node)
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line l1.y:161
		{
			yyVAL.node = newMemNode(yyDollar[3].node, yyDollar[4].n)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:166
		{
			yyVAL.node = yyDollar[1].node
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:167
		{
			yyVAL.node = yyDollar[1].node
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:170
		{
			yyVAL.node = yyDollar[1].node
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:171
		{
			yyVAL.node = yyDollar[1].node
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:174
		{
			yyVAL.node = yyDollar[1].node
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:175
		{
			yyVAL.node = yyDollar[1].node
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:176
		{
			yyVAL.node = yyDollar[1].node
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:178
		{
			yyVAL.node = newTokenNode(yyDollar[1].s)
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:179
		{
			yyVAL.node = yyDollar[1].node
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:182
		{
			yyVAL.node = newTokenNode(yyDollar[1].s)
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:183
		{
			yyVAL.node = yyDollar[1].node
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:186
		{
			yyVAL.node = newTokenNode(yyDollar[1].s)
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:187
		{
			yyVAL.node = yyDollar[1].node
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:190
		{
			yyVAL.node = newTokenNode(yyDollar[1].s)
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:193
		{
			yyVAL.node = newTokenNode(strconv.Itoa(yyDollar[1].n))
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:194
		{
			yyVAL.node = newTokenNode(strconv.Itoa(yyDollar[1].n))
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:195
		{
			yyVAL.node = newTokenNode(strconv.Itoa(yyDollar[1].n))
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:196
		{
			yyVAL.node = newTokenNode(strconv.Itoa(yyDollar[1].n))
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:200
		{
			yyVAL.node = newLabelNode(yyDollar[1].s)
		}
	}
	goto yystack /* stack new state and value */
}
