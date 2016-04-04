package l1compiler

import (
	"fmt"
	"strconv"
)
import (
	"bufio"
	"io"
	"strings"
)

type frame struct {
	i            int
	s            string
	line, column int
}
type Lexer struct {
	// The lexer runs in its own goroutine, and communicates via channel 'ch'.
	ch chan frame
	// We record the level of nesting because the action could return, and a
	// subsequent call expects to pick up where it left off. In other words,
	// we're simulating a coroutine.
	// TODO: Support a channel-based variant that compatible with Go's yacc.
	stack []frame
	stale bool

	// The 'l' and 'c' fields were added for
	// https://github.com/wagerlabs/docker/blob/65694e801a7b80930961d70c69cba9f2465459be/buildfile.nex
	// Since then, I introduced the built-in Line() and Column() functions.
	l, c int

	parseResult interface{}

	// The following line makes it easy for scripts to insert fields in the
	// generated code.
p *Compiler
	// [NEX_END_OF_LEXER_STRUCT]
}

// NewLexerWithInit creates a new Lexer object, runs the given callback on it,
// then returns it.
func NewLexerWithInit(in io.Reader, initFun func(*Lexer)) *Lexer {
	type dfa struct {
		acc          []bool           // Accepting states.
		f            []func(rune) int // Transitions.
		startf, endf []int            // Transitions at start and end of input.
		nest         []dfa
	}
	yylex := new(Lexer)
	if initFun != nil {
		initFun(yylex)
	}
	yylex.ch = make(chan frame)
	var scan func(in *bufio.Reader, ch chan frame, family []dfa, line, column int)
	scan = func(in *bufio.Reader, ch chan frame, family []dfa, line, column int) {
		// Index of DFA and length of highest-precedence match so far.
		matchi, matchn := 0, -1
		var buf []rune
		n := 0
		checkAccept := func(i int, st int) bool {
			// Higher precedence match? DFAs are run in parallel, so matchn is at most len(buf), hence we may omit the length equality check.
			if family[i].acc[st] && (matchn < n || matchi > i) {
				matchi, matchn = i, n
				return true
			}
			return false
		}
		var state [][2]int
		for i := 0; i < len(family); i++ {
			mark := make([]bool, len(family[i].startf))
			// Every DFA starts at state 0.
			st := 0
			for {
				state = append(state, [2]int{i, st})
				mark[st] = true
				// As we're at the start of input, follow all ^ transitions and append to our list of start states.
				st = family[i].startf[st]
				if -1 == st || mark[st] {
					break
				}
				// We only check for a match after at least one transition.
				checkAccept(i, st)
			}
		}
		atEOF := false
		for {
			if n == len(buf) && !atEOF {
				r, _, err := in.ReadRune()
				switch err {
				case io.EOF:
					atEOF = true
				case nil:
					buf = append(buf, r)
				default:
					panic(err)
				}
			}
			if !atEOF {
				r := buf[n]
				n++
				var nextState [][2]int
				for _, x := range state {
					x[1] = family[x[0]].f[x[1]](r)
					if -1 == x[1] {
						continue
					}
					nextState = append(nextState, x)
					checkAccept(x[0], x[1])
				}
				state = nextState
			} else {
			dollar: // Handle $.
				for _, x := range state {
					mark := make([]bool, len(family[x[0]].endf))
					for {
						mark[x[1]] = true
						x[1] = family[x[0]].endf[x[1]]
						if -1 == x[1] || mark[x[1]] {
							break
						}
						if checkAccept(x[0], x[1]) {
							// Unlike before, we can break off the search. Now that we're at the end, there's no need to maintain the state of each DFA.
							break dollar
						}
					}
				}
				state = nil
			}

			if state == nil {
				lcUpdate := func(r rune) {
					if r == '\n' {
						line++
						column = 0
					} else {
						column++
					}
				}
				// All DFAs stuck. Return last match if it exists, otherwise advance by one rune and restart all DFAs.
				if matchn == -1 {
					if len(buf) == 0 { // This can only happen at the end of input.
						break
					}
					lcUpdate(buf[0])
					buf = buf[1:]
				} else {
					text := string(buf[:matchn])
					buf = buf[matchn:]
					matchn = -1
					ch <- frame{matchi, text, line, column}
					if len(family[matchi].nest) > 0 {
						scan(bufio.NewReader(strings.NewReader(text)), ch, family[matchi].nest, line, column)
					}
					if atEOF {
						break
					}
					for _, r := range text {
						lcUpdate(r)
					}
				}
				n = 0
				for i := 0; i < len(family); i++ {
					state = append(state, [2]int{i, 0})
				}
			}
		}
		ch <- frame{-1, "", line, column}
	}
	go scan(bufio.NewReader(in), yylex.ch, []dfa{
		// \(
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 40:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 40:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \)
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 41:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 41:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// ;[^\r\n]*[\r\n]
		{[]bool{false, false, true, false}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 10:
					return -1
				case 13:
					return -1
				case 59:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 10:
					return 2
				case 13:
					return 2
				case 59:
					return 3
				}
				return 3
			},
			func(r rune) int {
				switch r {
				case 10:
					return -1
				case 13:
					return -1
				case 59:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 10:
					return 2
				case 13:
					return 2
				case 59:
					return 3
				}
				return 3
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// [ \t]+
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 9:
					return 1
				case 32:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 9:
					return 1
				case 32:
					return 1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// return
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 114:
					return 1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 2
				case 110:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 116:
					return 3
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				case 117:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 114:
					return 5
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return 6
				case 114:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// print
		{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 110:
					return -1
				case 112:
					return 1
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 110:
					return -1
				case 112:
					return -1
				case 114:
					return 2
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return 3
				case 110:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 110:
					return 4
				case 112:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 110:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 116:
					return 5
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 110:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

		// allocate
		{[]bool{false, false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return 1
				case 99:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 108:
					return 2
				case 111:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 108:
					return 3
				case 111:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 111:
					return 4
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return 5
				case 101:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 6
				case 99:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 116:
					return 7
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return 8
				case 108:
					return -1
				case 111:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// array-error
		{[]bool{false, false, false, false, false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 97:
					return 1
				case 101:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 97:
					return -1
				case 101:
					return -1
				case 111:
					return -1
				case 114:
					return 2
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 97:
					return -1
				case 101:
					return -1
				case 111:
					return -1
				case 114:
					return 3
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 97:
					return 4
				case 101:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 97:
					return -1
				case 101:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 121:
					return 5
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return 6
				case 97:
					return -1
				case 101:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 97:
					return -1
				case 101:
					return 7
				case 111:
					return -1
				case 114:
					return -1
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 97:
					return -1
				case 101:
					return -1
				case 111:
					return -1
				case 114:
					return 8
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 97:
					return -1
				case 101:
					return -1
				case 111:
					return -1
				case 114:
					return 9
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 97:
					return -1
				case 101:
					return -1
				case 111:
					return 10
				case 114:
					return -1
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 97:
					return -1
				case 101:
					return -1
				case 111:
					return -1
				case 114:
					return 11
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 97:
					return -1
				case 101:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 121:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// cjump
		{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 99:
					return 1
				case 106:
					return -1
				case 109:
					return -1
				case 112:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 106:
					return 2
				case 109:
					return -1
				case 112:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 106:
					return -1
				case 109:
					return -1
				case 112:
					return -1
				case 117:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 106:
					return -1
				case 109:
					return 4
				case 112:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 106:
					return -1
				case 109:
					return -1
				case 112:
					return 5
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 106:
					return -1
				case 109:
					return -1
				case 112:
					return -1
				case 117:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

		// goto
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 103:
					return 1
				case 111:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 111:
					return 2
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 111:
					return -1
				case 116:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 111:
					return 4
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 111:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// mem
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 109:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 2
				case 109:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 109:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 109:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// <-
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 60:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return 2
				case 60:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 60:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// call
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return 1
				case 108:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 2
				case 99:
					return -1
				case 108:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 108:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 108:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 108:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// tail-call
		{[]bool{false, false, false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 116:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 97:
					return 2
				case 99:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 105:
					return 3
				case 108:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 105:
					return -1
				case 108:
					return 4
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return 5
				case 97:
					return -1
				case 99:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 97:
					return -1
				case 99:
					return 6
				case 105:
					return -1
				case 108:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 97:
					return 7
				case 99:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 105:
					return -1
				case 108:
					return 8
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 105:
					return -1
				case 108:
					return 9
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// \n
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 10:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 10:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// 0|-?([1-9][0-9]*)
		{[]bool{false, false, true, true, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 45:
					return 1
				case 48:
					return 2
				}
				switch {
				case 48 <= r && r <= 48:
					return -1
				case 49 <= r && r <= 57:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 48:
					return -1
				}
				switch {
				case 48 <= r && r <= 48:
					return -1
				case 49 <= r && r <= 57:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 48:
					return -1
				}
				switch {
				case 48 <= r && r <= 48:
					return -1
				case 49 <= r && r <= 57:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 48:
					return 4
				}
				switch {
				case 48 <= r && r <= 48:
					return 4
				case 49 <= r && r <= 57:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 48:
					return 4
				}
				switch {
				case 48 <= r && r <= 48:
					return 4
				case 49 <= r && r <= 57:
					return 4
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// rcx
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 114:
					return 1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return 2
				case 114:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 114:
					return -1
				case 120:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 114:
					return -1
				case 120:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// (rdi)|(rsi)|(rdx)|(r8)|(r9)
		{[]bool{false, false, true, true, false, false, true, true, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 56:
					return -1
				case 57:
					return -1
				case 100:
					return -1
				case 105:
					return -1
				case 114:
					return 1
				case 115:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 56:
					return 2
				case 57:
					return 3
				case 100:
					return 4
				case 105:
					return -1
				case 114:
					return -1
				case 115:
					return 5
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 56:
					return -1
				case 57:
					return -1
				case 100:
					return -1
				case 105:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 56:
					return -1
				case 57:
					return -1
				case 100:
					return -1
				case 105:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 56:
					return -1
				case 57:
					return -1
				case 100:
					return -1
				case 105:
					return 7
				case 114:
					return -1
				case 115:
					return -1
				case 120:
					return 8
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 56:
					return -1
				case 57:
					return -1
				case 100:
					return -1
				case 105:
					return 6
				case 114:
					return -1
				case 115:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 56:
					return -1
				case 57:
					return -1
				case 100:
					return -1
				case 105:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 56:
					return -1
				case 57:
					return -1
				case 100:
					return -1
				case 105:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 56:
					return -1
				case 57:
					return -1
				case 100:
					return -1
				case 105:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 120:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// (rax)|(rbx)|(rbp)|(r10)|(r11)|(r12)|(r13)|(r14)|(r15)
		{[]bool{false, false, false, false, false, true, true, true, true, true, true, true, true, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 48:
					return -1
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 53:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 112:
					return -1
				case 114:
					return 1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 48:
					return -1
				case 49:
					return 2
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 53:
					return -1
				case 97:
					return 3
				case 98:
					return 4
				case 112:
					return -1
				case 114:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 48:
					return 8
				case 49:
					return 9
				case 50:
					return 10
				case 51:
					return 11
				case 52:
					return 12
				case 53:
					return 13
				case 97:
					return -1
				case 98:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 48:
					return -1
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 53:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 120:
					return 7
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 48:
					return -1
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 53:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 112:
					return 5
				case 114:
					return -1
				case 120:
					return 6
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 48:
					return -1
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 53:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 48:
					return -1
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 53:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 48:
					return -1
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 53:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 48:
					return -1
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 53:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 48:
					return -1
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 53:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 48:
					return -1
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 53:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 48:
					return -1
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 53:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 48:
					return -1
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 53:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 48:
					return -1
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 53:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 120:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// rsp
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 112:
					return -1
				case 114:
					return 1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 112:
					return 3
				case 114:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// (<=?)|=
		{[]bool{false, true, true, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 60:
					return 1
				case 61:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 60:
					return -1
				case 61:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 60:
					return -1
				case 61:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 60:
					return -1
				case 61:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// (<<=)|(>>=)
		{[]bool{false, false, false, false, true, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 60:
					return 1
				case 61:
					return -1
				case 62:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 60:
					return 5
				case 61:
					return -1
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 60:
					return -1
				case 61:
					return 4
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 60:
					return -1
				case 61:
					return 6
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// (\+=)|(-=)|(\*=)|(&=)
		{[]bool{false, false, false, false, false, true, true, true, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 38:
					return 1
				case 42:
					return 2
				case 43:
					return 3
				case 45:
					return 4
				case 61:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 38:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 61:
					return 8
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 38:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 61:
					return 7
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 38:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 61:
					return 6
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 38:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 61:
					return 5
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 38:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 61:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 38:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 61:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 38:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 61:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 38:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 61:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// :[a-zA-Z_][a-zA-Z0-9_]*
		{[]bool{false, false, true, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 58:
					return 1
				case 95:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 90:
					return -1
				case 97 <= r && r <= 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 58:
					return -1
				case 95:
					return 2
				}
				switch {
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 90:
					return 2
				case 97 <= r && r <= 122:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 58:
					return -1
				case 95:
					return 3
				}
				switch {
				case 48 <= r && r <= 57:
					return 3
				case 65 <= r && r <= 90:
					return 3
				case 97 <= r && r <= 122:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 58:
					return -1
				case 95:
					return 3
				}
				switch {
				case 48 <= r && r <= 57:
					return 3
				case 65 <= r && r <= 90:
					return 3
				case 97 <= r && r <= 122:
					return 3
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// .
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				return 1
			},
			func(r rune) int {
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},
	}, 0, 0)
	return yylex
}

func NewLexer(in io.Reader) *Lexer {
	return NewLexerWithInit(in, nil)
}

// Text returns the matched text.
func (yylex *Lexer) Text() string {
	return yylex.stack[len(yylex.stack)-1].s
}

// Line returns the current line number.
// The first line is 0.
func (yylex *Lexer) Line() int {
	if len(yylex.stack) == 0 {
		return 0
	}
	return yylex.stack[len(yylex.stack)-1].line
}

// Column returns the current column number.
// The first column is 0.
func (yylex *Lexer) Column() int {
	if len(yylex.stack) == 0 {
		return 0
	}
	return yylex.stack[len(yylex.stack)-1].column
}

func (yylex *Lexer) next(lvl int) int {
	if lvl == len(yylex.stack) {
		l, c := 0, 0
		if lvl > 0 {
			l, c = yylex.stack[lvl-1].line, yylex.stack[lvl-1].column
		}
		yylex.stack = append(yylex.stack, frame{0, "", l, c})
	}
	if lvl == len(yylex.stack)-1 {
		p := &yylex.stack[lvl]
		*p = <-yylex.ch
		yylex.stale = false
	} else {
		yylex.stale = true
	}
	return yylex.stack[lvl].i
}
func (yylex *Lexer) pop() {
	yylex.stack = yylex.stack[:len(yylex.stack)-1]
}
func (yylex Lexer) Error(e string) {
	panic(e)
}

// Lex runs the lexer. Always returns 0.
// When the -s option is given, this function is not generated;
// instead, the NN_FUN macro runs the lexer.
func (yylex *Lexer) Lex(lval *yySymType) int {
OUTER0:
	for {
		switch yylex.next(0) {
		case 0:
			{
				fmt.Println("LPAREN", yylex.Text())
				return LPAREN
			}
		case 1:
			{
				fmt.Println("RPAREN", yylex.Text())
				return RPAREN
			}
		case 2:
			{
				fmt.Println("COMMENT", yylex.Text()) /* ignore comments */
			}
		case 3:
			{ /*ignore whitespace */
			}
		case 4:
			{
				fmt.Println("RETURN", yylex.Text())
				return RETURN
			}
		case 5:
			{
				fmt.Println("PRINT", yylex.Text())
				return PRINT
			}
		case 6:
			{
				fmt.Println("ALLOCATE", yylex.Text())
				return ALLOCATE
			}
		case 7:
			{
				fmt.Println("ARRAYERROR", yylex.Text())
				return ARRAYERROR
			}
		case 8:
			{
				fmt.Println("CJUMP", yylex.Text())
				return CJUMP
			}
		case 9:
			{
				fmt.Println("GOTO", yylex.Text())
				return GOTO
			}
		case 10:
			{
				fmt.Println("MEM", yylex.Text())
				return MEM
			}
		case 11:
			{
				fmt.Println("ASSIGN", yylex.Text())
				return ASSIGN
			}
		case 12:
			{
				fmt.Println("CALL", yylex.Text())
				return CALL
			}
		case 13:
			{
				fmt.Println("TAILCALL", yylex.Text())
				return TAILCALL
			}
		case 14:
			{ /* ignore */
			}
		case 15:
			{
				fmt.Println("NUM", yylex.Text())
				lval.s = yylex.Text()
				//lval.n,_ = strconv.Atoi(yylex.Text())
				lvaln, _ := strconv.Atoi(yylex.Text())

				if lvaln >= 0 {
					if lvaln <= 6 {
						fmt.Println("NAT6")
						return NAT6
					} else if lvaln <= 255 {
						return NAT8
					} else {
						return NAT
					}
				} else {
					return NEG
				}
			}
		case 16:
			{
				fmt.Println("RCX", yylex.Text())
				lval.s = yylex.Text()
				return RCX
			}
		case 17:
			{
				fmt.Println("A", yylex.Text())
				lval.s = yylex.Text()
				return A
			}
		case 18:
			{
				fmt.Println("W", yylex.Text())
				lval.s = yylex.Text()
				return W
			}
		case 19:
			{
				fmt.Println("X", yylex.Text())
				lval.s = yylex.Text()
				return X
			}
		case 20:
			{
				fmt.Println("CMP", yylex.Text())
				lval.s = yylex.Text()
				return CMP
			}
		case 21:
			{
				fmt.Println("SOP", yylex.Text())
				lval.s = yylex.Text()
				return SOP
			}
		case 22:
			{
				fmt.Println("AOP", yylex.Text())
				lval.s = yylex.Text()
				return AOP
			}
		case 23:
			{
				fmt.Println("LABEL", yylex.Text())
				lval.s = yylex.Text()
				return LABEL
			}
		case 24:
			{
				fmt.Printf("Found invalid input on line: %d\n", yylex.Line())
			}
		default:
			break OUTER0
		}
		continue
	}
	yylex.pop()

	return 0
}
