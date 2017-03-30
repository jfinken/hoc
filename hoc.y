// Copyright 2011 Bobby Powers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// based off of Appendix A from http://dinosaur.compilertools.net/yacc/

%{

package main

import (
	"fmt"
	"os"
	"runtime"
	"math"

	"github.com/chzyer/readline"
)
var regs = make(map[string]float64)
%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
	val float64 
	fn func(float64)float64
	index string
}
%token <val> DIGIT 
%token <index> VAR 
%token <fn> BLTIN
%token <val> IDENT
%token <val> ERROR

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <val> stat expr
%right '='
%left '|'
%left '&'
%left '+'  '-'
%left '*'  '/'  '%'
%left UMINUS	/*  supplies  precedence  for  unary  minus  */
%right '^'		/* exponentiation */	

%%
list	: 	/* empty */
		| 	list stat 		'\n'
		| 	list stat 		';'	
		| 	list stat 		';'	'\n'
		;

stat	:    expr 			{ fmt.Printf( "%0.4f\n", $1 ) }
		|    VAR  '='  expr { regs[$1]  =  $3 }
		;

expr	:    DIGIT 			{ $$ = $1 }
		| 	'(' expr ')'	{ $$  =  $2 }
		|    expr '+' expr	{ $$  =  $1 + $3 }
		|    expr '-' expr	{ $$  =  $1 - $3 }
		|    expr '*' expr	{ $$  =  $1 * $3 }
		|    expr '/' expr 	{ $$  =  $1 / $3 }
		|    expr '^' expr 	{ $$  =  math.Pow($1, $3) }
		|    '-'  expr     	%prec  UMINUS  { $$  = -$2  }
		|   VAR  			{ $$  = regs[$1] }
		|   BLTIN '(' expr ')' { $$  = $1($3)  }
		;
%%      

/*  start  of  programs  */

func getHomeDir() string {
    if runtime.GOOS == "windows" {
        home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
        if home == "" {
            home = os.Getenv("USERPROFILE")
        }
        return home
    }
    return os.Getenv("HOME")
}
func main() {

	// readline for the prompt and history, a go implementation of
    // the gnu libreadline library.
    rl, err := readline.NewEx(&readline.Config{
        Prompt:      "hoc> ",
        HistoryFile: getHomeDir() + "/.eq_history",
    })

    if err != nil {
        panic(err)
    }
    defer rl.Close()

	HocDebug = 1
	HocErrorVerbose = true

	for {
        var eqn string
        var ok error 

		eqn, ok = rl.Readline()
        if ok != nil {
            break
        }
		if eqn == "" {
			continue
		}
		// Readline does not include the newline required by the parser
		eqn = eqn + "\n"
		// See lex.go for the lexer
		HocParse(&HocLex{src: eqn})
	}
}
