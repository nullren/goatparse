%{

package goatparse

import (
	"bytes"
	"fmt"
	"strings"
	"time"
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
	Goatparse
	Goatparse1
	Goatparse11
	Increment
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
		$$ = $2.(time.Duration)
	}

Increment:
	'+' Duration
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
	Goatparse time.Duration
	Goatparse1 time.Duration
	Goatparse11 time.Duration
	Increment time.Duration
	Start time.Time
)
