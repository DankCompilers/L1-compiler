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
const N8 = 57351
const NEGN8 = 57352
const AOP = 57353
const SOP = 57354
const CMP = 57355
const CALL = 57356
const CJUMP = 57357
const TAILCALL = 57358
const RETURN = 57359
const GOTO = 57360
const LPAREN = 57361
const RPAREN = 57362
const ALLOCATE = 57363
const PRINT = 57364
const ARRAYERROR = 57365
const ASSIGN = 57366
const MEM = 57367
const RSP = 57368
const RCX = 57369
const X = 57370
const W = 57371
const A = 57372

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LABEL",
	"GOLABEL",
	"NEG",
	"NAT",
	"NAT6",
	"N8",
	"NEGN8",
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

//line l1.y:214
func cast(y yyLexer) *Compiler { return y.(*Lexer).p }

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 66,
	20, 31,
	-2, 29,
	-1, 67,
	20, 32,
	-2, 30,
}

const yyNprod = 51
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 135

var yyAct = [...]int{

	51, 44, 10, 45, 18, 43, 63, 62, 46, 55,
	35, 41, 19, 14, 50, 49, 54, 52, 53, 86,
	24, 36, 47, 31, 34, 37, 20, 33, 76, 7,
	6, 2, 81, 60, 60, 36, 47, 31, 34, 61,
	80, 69, 66, 68, 65, 73, 68, 67, 72, 71,
	74, 70, 50, 49, 54, 52, 53, 19, 77, 64,
	15, 19, 54, 52, 53, 19, 82, 50, 49, 54,
	52, 53, 17, 36, 47, 31, 34, 21, 84, 11,
	12, 13, 79, 85, 36, 78, 31, 34, 36, 47,
	31, 34, 28, 26, 29, 30, 25, 33, 19, 50,
	49, 54, 52, 53, 48, 36, 83, 31, 34, 75,
	42, 9, 39, 40, 3, 57, 56, 58, 32, 27,
	36, 36, 23, 31, 34, 38, 4, 22, 16, 5,
	1, 0, 8, 59, 59,
}
var yyPact = [...]int{

	12, -1000, 110, 11, 9, 11, 107, -1000, -1000, 72,
	72, -1000, -1000, -1000, 53, 6, 53, 78, -1000, -1000,
	-1000, -1000, 5, 101, -13, 106, 46, -1000, 94, 57,
	-1000, -1000, -1000, -18, -1000, -1000, -1000, -1000, 8, 46,
	93, 61, -1000, 105, 15, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 72, 77, 74, 32, -1000,
	-1000, 24, -6, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 102, 46, -1000, -1000, -1000,
	-1000, -1000, 54, -1000, -1000, -1, -1000,
}
var yyPgo = [...]int{

	0, 130, 126, 129, 60, 128, 127, 20, 5, 119,
	6, 104, 9, 1, 3, 118, 10, 8, 4, 2,
	0,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 3, 4, 4, 5, 5, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 9, 9, 9, 8, 7, 12, 12, 13,
	13, 10, 10, 10, 14, 14, 11, 11, 15, 15,
	16, 17, 17, 17, 19, 19, 19, 20, 20, 20,
	18,
}
var yyR2 = [...]int{

	0, 4, 1, 2, 6, 1, 2, 3, 1, 3,
	3, 3, 3, 3, 3, 3, 2, 4, 1, 3,
	3, 1, 3, 3, 3, 3, 5, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1,
}
var yyChk = [...]int{

	-1000, -1, 19, 4, -2, -3, 19, 20, -2, 4,
	-19, 7, 8, 9, -19, -4, -5, 19, -18, 4,
	20, -4, -6, -11, -7, 18, 15, -9, 14, 16,
	17, 29, -15, 19, 30, -16, 27, 20, 24, 11,
	12, 24, 4, -8, -13, -14, -17, 28, -11, 7,
	6, -20, 9, 10, 8, -12, 22, 21, 23, -11,
	-18, -12, 25, -10, -7, -8, -14, -17, -18, -13,
	-16, -17, -10, -14, -17, 4, 13, -19, 8, 8,
	8, 8, -14, 4, -13, -20, 20,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 2, 0, 1, 3, 0,
	0, 44, 45, 46, 0, 0, 5, 0, 8, 50,
	4, 6, 0, 0, 0, 0, 0, 18, 0, 0,
	21, 36, 37, 0, 38, 39, 40, 7, 0, 0,
	0, 0, 16, 0, 0, 29, 30, 34, 35, 41,
	42, 43, 47, 48, 49, 0, 0, 0, 0, 27,
	28, 0, 0, 9, 10, 15, -2, -2, 33, 12,
	13, 14, 11, 31, 32, 0, 0, 19, 22, 23,
	24, 20, 0, 17, 25, 0, 26,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30,
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
		//line l1.y:32
		{
			fmt.Println("Detected program: %+v", yyDollar[2].s)
			yyVAL.node = newProgramNode(yyDollar[2].s, yyDollar[3].node)
			cast(yylex).SetAstRoot(yyVAL.node)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:39
		{
			fmt.Println("Detected end subprogram")
			yyVAL.node = newSubProgramNode(yyDollar[1].node, nil)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line l1.y:44
		{
			fmt.Printf("Detected subprogram\n")
			yyVAL.node = newSubProgramNode(yyDollar[1].node, yyDollar[2].node)
		}
	case 4:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line l1.y:51
		{
			fmt.Println("Detected func: %+v", yyDollar[2].s)
			yyVAL.node = newFunctionNode(yyDollar[2].s, yyDollar[3].n, yyDollar[4].n, yyDollar[5].node)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:58
		{
			fmt.Printf("Detected end instruction\n")
			yyVAL.node = newInstructionNode(yyDollar[1].node, nil)
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line l1.y:63
		{
			fmt.Printf("Detected instruction\n")
			yyVAL.node = newInstructionNode(yyDollar[1].node, yyDollar[2].node)
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:70
		{
			yyVAL.node = yyDollar[2].node
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:74
		{
			yyVAL.node = yyDollar[1].node
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
			//fmt.Printf("Detected assign \n")
			yyVAL.node = newAssignNode(yyDollar[1].node, yyDollar[3].node)
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:95
		{
			//fmt.Printf("Detected aop \n")
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
			//fmt.Printf("Detected sop \n")
			yyVAL.node = newOpNode(yyDollar[2].s, yyDollar[1].node, yyDollar[3].node)
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:110
		{
			yyVAL.node = newAssignNode(yyDollar[1].node, yyDollar[3].node)
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line l1.y:114
		{
			//fmt.Printf("Detected goto: %q\n", $2)
			yyVAL.node = newGotoNode(yyDollar[2].s)
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line l1.y:119
		{
			yyVAL.node = newCjumpNode(yyDollar[2].node, yyDollar[3].s, yyDollar[4].s)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:123
		{
			yyVAL.node = yyDollar[1].node
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:127
		{
			yyVAL.node = newCallNode(yyDollar[2].node, yyDollar[3].n)
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:131
		{
			yyVAL.node = newTailcallNode(yyDollar[2].node, yyDollar[3].n)
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:135
		{
			yyVAL.node = newReturnNode()
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:143
		{
			yyVAL.node = newSysCallNode(yyDollar[2].s, 1)
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:147
		{
			yyVAL.node = newSysCallNode(yyDollar[2].s, 2)
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:151
		{
			yyVAL.node = newSysCallNode(yyDollar[2].s, 3)
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line l1.y:157
		{
			yyVAL.node = newCmpopNode(yyDollar[2].s, yyDollar[1].node, yyDollar[3].node)
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line l1.y:163
		{
			yyVAL.node = newMemNode(yyDollar[3].node, yyDollar[4].n)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:170
		{
			yyVAL.node = yyDollar[1].node
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:171
		{
			yyVAL.node = yyDollar[1].node
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:174
		{
			yyVAL.node = yyDollar[1].node
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:175
		{
			yyVAL.node = yyDollar[1].node
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:178
		{
			yyVAL.node = yyDollar[1].node
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:179
		{
			yyVAL.node = yyDollar[1].node
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:180
		{
			yyVAL.node = yyDollar[1].node
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:182
		{
			yyVAL.node = newTokenNode(yyDollar[1].s)
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:183
		{
			yyVAL.node = yyDollar[1].node
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:186
		{
			yyVAL.node = newTokenNode(yyDollar[1].s)
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:187
		{
			yyVAL.node = yyDollar[1].node
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:190
		{
			yyVAL.node = newTokenNode(yyDollar[1].s)
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:191
		{
			yyVAL.node = yyDollar[1].node
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:194
		{
			yyVAL.node = newTokenNode(yyDollar[1].s)
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:196
		{
			yyVAL.node = newTokenNode(strconv.Itoa(yyDollar[1].n))
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:197
		{
			yyVAL.node = newTokenNode(strconv.Itoa(yyDollar[1].n))
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:198
		{
			yyVAL.node = newTokenNode(strconv.Itoa(yyDollar[1].n))
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:201
		{
			yyVAL.n = yyDollar[1].n
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:202
		{
			yyVAL.n = yyDollar[1].n
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:203
		{
			yyVAL.n = yyDollar[1].n
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:205
		{
			yyVAL.n = yyDollar[1].n
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:206
		{
			yyVAL.n = yyDollar[1].n
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:207
		{
			yyVAL.n = yyDollar[1].n
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line l1.y:211
		{
			yyVAL.node = newLabelNode(yyDollar[1].s)
		}
	}
	goto yystack /* stack new state and value */
}
