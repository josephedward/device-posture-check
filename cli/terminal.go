package cli

import (
	"fmt"
	"github.com/rs/zerolog"
	// "github.com/rs/zerolog/log"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func PrintIfErr(err error) {
	if err != nil {
		Error(err.Error())
	}
}

func Success(message ...interface{}) {

	//if log level is debug, print success messages
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		for _, msg := range message {
			s, ok := msg.(string) // the "ok" boolean will flag success.
			if ok {
				fmt.Println(Green + string(s) + Reset)
			} else {
				fmt.Println(msg)
			}
		}
	}
}

func Error(message ...interface{}) {
	//if log level is debug, print err messages
	if zerolog.GlobalLevel() == zerolog.DebugLevel {

		for _, msg := range message {
			s, ok := msg.(string) // the "ok" boolean will flag success.
			if ok {
				fmt.Println(Red + string(s) + Reset)
			} else {
				fmt.Println(msg)
			}
		}
	}
}


func Welcome() {
	fmt.Println("	                                                    		")
	fmt.Println("                :      ED.                             		")
	fmt.Println("               t#,     E#Wi                          .,		")
	fmt.Println("        .Gt   ;##W.    E###G.       t               ,Wt		")
	fmt.Println("       j#W:  :#L:WE    E#fD#W;      ED.            i#D.		")
	fmt.Println("     ;K#f   .KG  ,#D   E#t t##L     E#K:          f#f  		")
	fmt.Println("   .G#D.    EE    ;#f  E#t  .E#K,   E##W;       .D#i   		")
	fmt.Println("  j#K;     f#.     t#i E#t    j##f  E#E##t     :KW,    		")
	fmt.Println(",K#f   ,GD;:#G     GK  E#t    :E#K: E#ti##f    t#f     		")
	fmt.Println(" j#Wi   E#t ;#L   LW.  E#t   t##L   E#t  ;##D.  ;#G    		")
	fmt.Println("  .G#D: E#t  t#f f#:   E#t .D#W;    E#E  .##K:   :KE.  		")
	fmt.Println("    ,K#fK#t   f#D#;    E#tiW#G.     E#L;;;;;;,    .DW: 		")
	fmt.Println("      j###t    G#t     E#K##i       E#t             L#,		")
	fmt.Println("       .G#t     t      E##D.        E#t              jt		")
	fmt.Println("         ;;            E#t          E#t                ) 		")
	fmt.Println("                       L:                              		")
	
	}
	