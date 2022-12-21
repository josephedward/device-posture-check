package cli

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
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

var log = zerolog.New(os.Stdout).With().Timestamp().Logger()

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
				// log the success
				// log.Info().Msg(s)
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
				//log the error
				// log.Error().Msg(s)
			} else {
				fmt.Println(msg)
			}
		}
	}
}

func Welcome() {
	fmt.Println(Blue + "                          :          ED.                             				" + Reset)
	fmt.Println(Blue + "                         t#,         E#Wi                          .,				" + Reset)
	fmt.Println(Blue + "                  .Gt   ;##W.        E###G.       t               ,Wt				" + Reset)
	fmt.Println(Blue + "                 j#W:  :#L:WE        E#fD#W;      ED.            i#D.				" + Reset)
	fmt.Println(Blue + "               ;K#f   .KG  ,#D       E#t t##L     E#K:          f#f  				" + Reset)
	fmt.Println(Blue + "             .G#D.    EE    ;#f      E#t  .E#K,   E##W;       .D#i   				" + Reset)
	fmt.Println(Blue + "            j#K;     f#.     t#i ### E#t    j##f  E#E##t     :KW,    				" + Reset)
	fmt.Println(Blue + "          ,K#f   ,GD;:#G     GK      E#t    :E#K: E#ti##f    t#f     				" + Reset)
	fmt.Println(Blue + "           j#Wi   E#t ;#L   LW.      E#t   t##L   E#t  ;##D.  ;#G    				" + Reset)
	fmt.Println(Blue + "            .G#D: E#t  t#f f#:       E#t .D#W;    E#E  .##K:   :KE.  				" + Reset)
	fmt.Println(Blue + "              ,K#fK#t   f#D#;        E#tiW#G.     E#L;;;;;;,    .DW: 				" + Reset)
	fmt.Println(Blue + "                j###t    G#t         E#K##i       E#t             L#,				" + Reset)
	fmt.Println(Blue + "                 .G#t     t          E##D.        E#t              jt				" + Reset)
	fmt.Println(Blue + "                   ;;                E#t    **    E#t     **        )  **           " + Reset)
	fmt.Println(Blue + "                                     L:                                             " + Reset)
	fmt.Println(Blue + "                                                                                    " + Reset)
	fmt.Println(Green + "                    Golang Device Check with Tailscale and Fleet                    " + Reset)
}


func ZeroLog() zerolog.Logger {
	// fileOutput("log.txt")
	//look through all os.Args and see if one is "prod"
	for _, arg := range os.Args {
		if arg == "prod" {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
			log.Info().Msg("log level : " + zerolog.GlobalLevel().String())
			return log
		}
	}
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Info().Msg("log level : " + zerolog.GlobalLevel().String())
	return log
}
