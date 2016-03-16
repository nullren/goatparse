%{

package goatparse

import (
	"time"
	"fmt"
	"strings"
	"text/scanner"
  "strconv"
)

%}

%union {
	item interface{}
}

%token	DAY
%token	FORTNIGHT
%token	HOUR
%token	INTEGER
%token	MINUTE
%token	SECOND
%token	WEEK

%type	<item>
	DAY
	FORTNIGHT
	HOUR
	INTEGER
	MINUTE
	SECOND
	WEEK

%type	<item>
	Decrement
	Duration
	Duration1
	Increment
	Goatparse
	Goatparse1
	Goatparse11
	Start

%start Start

%%

Decrement:
	'-' Duration
	{
		$$ = -$2.(time.Duration)
	}

Duration:
	INTEGER Duration1
	{
		$$ = time.Duration($1.(int)) * $2.(time.Duration)
	}

Duration1:
	SECOND
	{
		$$ = time.Second
	}
|	MINUTE
	{
		$$ = time.Minute
	}
|	HOUR
	{
		$$ = time.Hour
	}
|	DAY
	{
		$$ = time.Hour * 24
	}
|	WEEK
	{
		$$ = time.Hour * 24 * 7
	}
|	FORTNIGHT
	{
		$$ = time.Hour * 24 * 7 * 2
	}

Increment:
	'+' Duration
	{
		$$ = $2.(time.Duration)
	}

Goatparse:
	Duration Goatparse1
	{
		$$ = $1.(time.Duration) + $2.(time.Duration)
	}

Goatparse1:
	/* EMPTY */
	{
		$$ = time.Duration(0)
	}
|	Goatparse1 Goatparse11
	{
		$$ = $1.(time.Duration) + $2.(time.Duration)
	}

Goatparse11:
	Increment
	{
		$$ = $1.(time.Duration)
	}
|	Decrement
	{
		$$ = $1.(time.Duration)
	}

Start:
	Goatparse
	{
		_parserResult = time.Now().Add($1.(time.Duration))
	}

%%

var _parserResult time.Time

type (
	Decrement time.Duration
	Duration time.Duration
	Duration1 time.Duration
	Increment time.Duration
	Goatparse time.Duration
	Goatparse1 time.Duration
	Goatparse11 time.Duration
	Start time.Time
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

  if tt == "d" || tt == "day" || tt == "days" {
       return DAY
  }

if tt == "h" || tt == "hour" || tt == "hours" {
       return HOUR
  }

if i, e := strconv.Atoi(tt); e == nil {
lval.item = i
       return INTEGER
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
  Time time.Time
  Offset int
}

func ParseDurationFromNow(input string) (*GoatparseResult, error) {
  var sc scanner.Scanner
  sc.Init(strings.NewReader(input))
  GoatparseParse(&GoatparseLex{s: &sc})
  return &GoatparseResult{Time:_parserResult, Offset:sc.Pos().Offset}, err
}
