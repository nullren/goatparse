//line goatparse.y:2
package goatparse

import __yyfmt__ "fmt"

//line goatparse.y:3
import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
	"time"
)

//line goatparse.y:15
type GoatparseSymType struct {
	yys  int
	item interface{}
}

const DAY = 57346
const FORTNIGHT = 57347
const HOUR = 57348
const INTEGER = 57349
const MINUTE = 57350
const SECOND = 57351
const WEEK = 57352

var GoatparseToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"DAY",
	"FORTNIGHT",
	"HOUR",
	"INTEGER",
	"MINUTE",
	"SECOND",
	"WEEK",
	"'-'",
	"'+'",
}
var GoatparseStatenames = [...]string{}

const GoatparseEofCode = 1
const GoatparseErrCode = 2
const GoatparseInitialStackSize = 16

//line goatparse.y:76

var _parserResult time.Time

type (
	Decrement   time.Duration
	Duration    time.Duration
	Duration1   time.Duration
	Increment   time.Duration
	Goatparse   time.Duration
	Goatparse1  time.Duration
	Goatparse11 time.Duration
	Start       time.Time
)

type GoatparseLex struct {
	s *scanner.Scanner
}

func (l *GoatparseLex) Lex(lval *GoatparseSymType) int {
	tok := l.s.Scan()
	if tok == scanner.EOF {
		return 0
	}
	tt := strings.ToLower(l.s.TokenText())

	if i, e := strconv.Atoi(tt); e == nil {
		lval.item = i
		return INTEGER
	}
	if tt == "d" || tt == "day" || tt == "days" {
		return DAY
	}
	if tt == "h" || tt == "hour" || tt == "hours" {
		return HOUR
	}
	if tt == "m" || tt == "min" || tt == "mins" || tt == "minute" || tt == "minutes" {
		return MINUTE
	}
	if tt == "second" || tt == "s" || tt == "seconds" {
		return SECOND
	}
	if tt == "week" || tt == "weeks" || tt == "w" {
		return WEEK
	}
	if tt == "fortnight" || tt == "fortnights" {
		return FORTNIGHT
	}
	return int(tok)
}

var err error

func (l *GoatparseLex) Error(s string) {
	err = fmt.Errorf("%v\n", s)
}

type GoatparseResult struct {
	Time   time.Time
	Offset int
}

func ParseDuration(input string) (*GoatparseResult, error) {
	var sc scanner.Scanner
	sc.Init(strings.NewReader(input))
	GoatparseParse(&GoatparseLex{s: &sc})
	return &GoatparseResult{Time: _parserResult, Offset: sc.Pos().Offset}, err
}

//line yacctab:1
var GoatparseExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const GoatparseNprod = 16
const GoatparsePrivate = 57344

var GoatparseTokenNames []string
var GoatparseStates []string

const GoatparseLast = 19

var GoatparseAct = [...]int{

	3, 10, 12, 9, 1, 8, 7, 11, 4, 17,
	16, 13, 5, 2, 14, 6, 15, 18, 19,
}
var GoatparsePact = [...]int{

	1, -1000, -1000, -1000, -3, -2, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 1, 1, -1000, -1000,
}
var GoatparsePgo = [...]int{

	0, 16, 0, 15, 14, 13, 12, 11, 4,
}
var GoatparseR1 = [...]int{

	0, 1, 2, 3, 3, 3, 3, 3, 3, 4,
	5, 6, 6, 7, 7, 8,
}
var GoatparseR2 = [...]int{

	0, 2, 2, 1, 1, 1, 1, 1, 1, 2,
	2, 0, 2, 1, 1, 1,
}
var GoatparseChk = [...]int{

	-1000, -8, -5, -2, 7, -6, -3, 9, 8, 6,
	4, 10, 5, -7, -4, -1, 12, 11, -2, -2,
}
var GoatparseDef = [...]int{

	0, -2, 15, 11, 0, 10, 2, 3, 4, 5,
	6, 7, 8, 12, 13, 14, 0, 0, 9, 1,
}
var GoatparseTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 12, 3, 11,
}
var GoatparseTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10,
}
var GoatparseTok3 = [...]int{
	0,
}

var GoatparseErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	GoatparseDebug        = 0
	GoatparseErrorVerbose = false
)

type GoatparseLexer interface {
	Lex(lval *GoatparseSymType) int
	Error(s string)
}

type GoatparseParser interface {
	Parse(GoatparseLexer) int
	Lookahead() int
}

type GoatparseParserImpl struct {
	lval  GoatparseSymType
	stack [GoatparseInitialStackSize]GoatparseSymType
	char  int
}

func (p *GoatparseParserImpl) Lookahead() int {
	return p.char
}

func GoatparseNewParser() GoatparseParser {
	return &GoatparseParserImpl{}
}

const GoatparseFlag = -1000

func GoatparseTokname(c int) string {
	if c >= 1 && c-1 < len(GoatparseToknames) {
		if GoatparseToknames[c-1] != "" {
			return GoatparseToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func GoatparseStatname(s int) string {
	if s >= 0 && s < len(GoatparseStatenames) {
		if GoatparseStatenames[s] != "" {
			return GoatparseStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func GoatparseErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !GoatparseErrorVerbose {
		return "syntax error"
	}

	for _, e := range GoatparseErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + GoatparseTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := GoatparsePact[state]
	for tok := TOKSTART; tok-1 < len(GoatparseToknames); tok++ {
		if n := base + tok; n >= 0 && n < GoatparseLast && GoatparseChk[GoatparseAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if GoatparseDef[state] == -2 {
		i := 0
		for GoatparseExca[i] != -1 || GoatparseExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; GoatparseExca[i] >= 0; i += 2 {
			tok := GoatparseExca[i]
			if tok < TOKSTART || GoatparseExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if GoatparseExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += GoatparseTokname(tok)
	}
	return res
}

func Goatparselex1(lex GoatparseLexer, lval *GoatparseSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = GoatparseTok1[0]
		goto out
	}
	if char < len(GoatparseTok1) {
		token = GoatparseTok1[char]
		goto out
	}
	if char >= GoatparsePrivate {
		if char < GoatparsePrivate+len(GoatparseTok2) {
			token = GoatparseTok2[char-GoatparsePrivate]
			goto out
		}
	}
	for i := 0; i < len(GoatparseTok3); i += 2 {
		token = GoatparseTok3[i+0]
		if token == char {
			token = GoatparseTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = GoatparseTok2[1] /* unknown char */
	}
	if GoatparseDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", GoatparseTokname(token), uint(char))
	}
	return char, token
}

func GoatparseParse(Goatparselex GoatparseLexer) int {
	return GoatparseNewParser().Parse(Goatparselex)
}

func (Goatparsercvr *GoatparseParserImpl) Parse(Goatparselex GoatparseLexer) int {
	var Goatparsen int
	var GoatparseVAL GoatparseSymType
	var GoatparseDollar []GoatparseSymType
	_ = GoatparseDollar // silence set and not used
	GoatparseS := Goatparsercvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Goatparsestate := 0
	Goatparsercvr.char = -1
	Goatparsetoken := -1 // Goatparsercvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		Goatparsestate = -1
		Goatparsercvr.char = -1
		Goatparsetoken = -1
	}()
	Goatparsep := -1
	goto Goatparsestack

ret0:
	return 0

ret1:
	return 1

Goatparsestack:
	/* put a state and value onto the stack */
	if GoatparseDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", GoatparseTokname(Goatparsetoken), GoatparseStatname(Goatparsestate))
	}

	Goatparsep++
	if Goatparsep >= len(GoatparseS) {
		nyys := make([]GoatparseSymType, len(GoatparseS)*2)
		copy(nyys, GoatparseS)
		GoatparseS = nyys
	}
	GoatparseS[Goatparsep] = GoatparseVAL
	GoatparseS[Goatparsep].yys = Goatparsestate

Goatparsenewstate:
	Goatparsen = GoatparsePact[Goatparsestate]
	if Goatparsen <= GoatparseFlag {
		goto Goatparsedefault /* simple state */
	}
	if Goatparsercvr.char < 0 {
		Goatparsercvr.char, Goatparsetoken = Goatparselex1(Goatparselex, &Goatparsercvr.lval)
	}
	Goatparsen += Goatparsetoken
	if Goatparsen < 0 || Goatparsen >= GoatparseLast {
		goto Goatparsedefault
	}
	Goatparsen = GoatparseAct[Goatparsen]
	if GoatparseChk[Goatparsen] == Goatparsetoken { /* valid shift */
		Goatparsercvr.char = -1
		Goatparsetoken = -1
		GoatparseVAL = Goatparsercvr.lval
		Goatparsestate = Goatparsen
		if Errflag > 0 {
			Errflag--
		}
		goto Goatparsestack
	}

Goatparsedefault:
	/* default state action */
	Goatparsen = GoatparseDef[Goatparsestate]
	if Goatparsen == -2 {
		if Goatparsercvr.char < 0 {
			Goatparsercvr.char, Goatparsetoken = Goatparselex1(Goatparselex, &Goatparsercvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if GoatparseExca[xi+0] == -1 && GoatparseExca[xi+1] == Goatparsestate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Goatparsen = GoatparseExca[xi+0]
			if Goatparsen < 0 || Goatparsen == Goatparsetoken {
				break
			}
		}
		Goatparsen = GoatparseExca[xi+1]
		if Goatparsen < 0 {
			goto ret0
		}
	}
	if Goatparsen == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Goatparselex.Error(GoatparseErrorMessage(Goatparsestate, Goatparsetoken))
			Nerrs++
			if GoatparseDebug >= 1 {
				__yyfmt__.Printf("%s", GoatparseStatname(Goatparsestate))
				__yyfmt__.Printf(" saw %s\n", GoatparseTokname(Goatparsetoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Goatparsep >= 0 {
				Goatparsen = GoatparsePact[GoatparseS[Goatparsep].yys] + GoatparseErrCode
				if Goatparsen >= 0 && Goatparsen < GoatparseLast {
					Goatparsestate = GoatparseAct[Goatparsen] /* simulate a shift of "error" */
					if GoatparseChk[Goatparsestate] == GoatparseErrCode {
						goto Goatparsestack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if GoatparseDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", GoatparseS[Goatparsep].yys)
				}
				Goatparsep--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if GoatparseDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", GoatparseTokname(Goatparsetoken))
			}
			if Goatparsetoken == GoatparseEofCode {
				goto ret1
			}
			Goatparsercvr.char = -1
			Goatparsetoken = -1
			goto Goatparsenewstate /* try again in the same state */
		}
	}

	/* reduction by production Goatparsen */
	if GoatparseDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Goatparsen, GoatparseStatname(Goatparsestate))
	}

	Goatparsent := Goatparsen
	Goatparsept := Goatparsep
	_ = Goatparsept // guard against "declared and not used"

	Goatparsep -= GoatparseR2[Goatparsen]
	// Goatparsep is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if Goatparsep+1 >= len(GoatparseS) {
		nyys := make([]GoatparseSymType, len(GoatparseS)*2)
		copy(nyys, GoatparseS)
		GoatparseS = nyys
	}
	GoatparseVAL = GoatparseS[Goatparsep+1]

	/* consult goto table to find next state */
	Goatparsen = GoatparseR1[Goatparsen]
	Goatparseg := GoatparsePgo[Goatparsen]
	Goatparsej := Goatparseg + GoatparseS[Goatparsep].yys + 1

	if Goatparsej >= GoatparseLast {
		Goatparsestate = GoatparseAct[Goatparseg]
	} else {
		Goatparsestate = GoatparseAct[Goatparsej]
		if GoatparseChk[Goatparsestate] != -Goatparsen {
			Goatparsestate = GoatparseAct[Goatparseg]
		}
	}
	// dummy call; replaced with literal code
	switch Goatparsent {

	case 1:
		GoatparseDollar = GoatparseS[Goatparsept-2 : Goatparsept+1]
		//line goatparse.y:50
		{
			GoatparseVAL.item = -GoatparseDollar[2].item.(time.Duration)
		}
	case 2:
		GoatparseDollar = GoatparseS[Goatparsept-2 : Goatparsept+1]
		//line goatparse.y:52
		{
			GoatparseVAL.item = time.Duration(GoatparseDollar[1].item.(int)) * GoatparseDollar[2].item.(time.Duration)
		}
	case 3:
		GoatparseDollar = GoatparseS[Goatparsept-1 : Goatparsept+1]
		//line goatparse.y:55
		{
			GoatparseVAL.item = time.Second
		}
	case 4:
		GoatparseDollar = GoatparseS[Goatparsept-1 : Goatparsept+1]
		//line goatparse.y:56
		{
			GoatparseVAL.item = time.Minute
		}
	case 5:
		GoatparseDollar = GoatparseS[Goatparsept-1 : Goatparsept+1]
		//line goatparse.y:57
		{
			GoatparseVAL.item = time.Hour
		}
	case 6:
		GoatparseDollar = GoatparseS[Goatparsept-1 : Goatparsept+1]
		//line goatparse.y:58
		{
			GoatparseVAL.item = time.Hour * 24
		}
	case 7:
		GoatparseDollar = GoatparseS[Goatparsept-1 : Goatparsept+1]
		//line goatparse.y:59
		{
			GoatparseVAL.item = time.Hour * 24 * 7
		}
	case 8:
		GoatparseDollar = GoatparseS[Goatparsept-1 : Goatparsept+1]
		//line goatparse.y:60
		{
			GoatparseVAL.item = time.Hour * 24 * 7 * 2
		}
	case 9:
		GoatparseDollar = GoatparseS[Goatparsept-2 : Goatparsept+1]
		//line goatparse.y:62
		{
			GoatparseVAL.item = GoatparseDollar[2].item.(time.Duration)
		}
	case 10:
		GoatparseDollar = GoatparseS[Goatparsept-2 : Goatparsept+1]
		//line goatparse.y:64
		{
			GoatparseVAL.item = GoatparseDollar[1].item.(time.Duration) + GoatparseDollar[2].item.(time.Duration)
		}
	case 11:
		GoatparseDollar = GoatparseS[Goatparsept-0 : Goatparsept+1]
		//line goatparse.y:67
		{
			GoatparseVAL.item = time.Duration(0)
		}
	case 12:
		GoatparseDollar = GoatparseS[Goatparsept-2 : Goatparsept+1]
		//line goatparse.y:68
		{
			GoatparseVAL.item = GoatparseDollar[1].item.(time.Duration) + GoatparseDollar[2].item.(time.Duration)
		}
	case 13:
		GoatparseDollar = GoatparseS[Goatparsept-1 : Goatparsept+1]
		//line goatparse.y:71
		{
			GoatparseVAL.item = GoatparseDollar[1].item.(time.Duration)
		}
	case 14:
		GoatparseDollar = GoatparseS[Goatparsept-1 : Goatparsept+1]
		//line goatparse.y:72
		{
			GoatparseVAL.item = GoatparseDollar[1].item.(time.Duration)
		}
	case 15:
		GoatparseDollar = GoatparseS[Goatparsept-1 : Goatparsept+1]
		//line goatparse.y:74
		{
			_parserResult = time.Now().Add(GoatparseDollar[1].item.(time.Duration))
		}
	}
	goto Goatparsestack /* stack new state and value */
}
