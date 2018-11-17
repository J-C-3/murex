package shell

import (
	"fmt"
	"os"
	"sync"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/shell/history"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/home"
	"github.com/lmorg/murex/utils/readline"
)

var (
	// Interactive describes whether murex is running as an interactive shell or not
	Interactive bool

	// Prompt is the readline instance
	Prompt = readline.NewInstance()

	// PromptGoProc is an custom defined ID for each prompt Goprocess so we don't accidentally end up with multiple prompts running
	PromptGoProc = new(mutexCounter)
)

// Start the interactive shell
func Start() {
	/*defer func() {
		if debug.Enable {
			return
		}
		if r := recover(); r != nil {
			os.Stderr.WriteString(fmt.Sprintln("Panic caught:", r))
			Start()
		}
	}()*/

	var err error

	Interactive = true
	Prompt.TempDirectory = consts.TempDir
	Prompt.TabCompleter = tabCompletion
	Prompt.SyntaxCompleter = syntaxCompletion
	Prompt.HistoryAutoWrite = false

	h, err := history.New(home.MyDir + consts.PathSlash + ".murex_history")
	if err != nil {
		//ansi.Stderrln(proc.ShellProcess, ansi.FgRed, "Error opening history file: "+err.Error())
		proc.ShellProcess.Stderr.Writeln([]byte("Error opening history file: " + err.Error()))
	} else {
		Prompt.History = h
	}

	SignalHandler(true)

	go autocomplete.UpdateGlobalExeList()

	v, err := proc.ShellProcess.Config.Get("shell", "max-suggestions", types.Integer)
	if err != nil {
		v = 4
	}
	Prompt.MaxTabCompleterRows = v.(int)

	prompt()

	noQuit := make(chan bool)
	<-noQuit
}

func prompt() {
	thisProc := PromptGoProc.Add()

	nLines := 1
	var merged string
	var block []rune
	Prompt.GetMultiLine = func(r []rune) []rune {
		var multiLine []rune
		if len(block) == 0 {
			multiLine = r
		} else {
			multiLine = append(append(block, []rune(utils.NewLineString)...), r...)
		}

		expanded, err := history.ExpandVariables(multiLine, Prompt)
		if err != nil {
			expanded = multiLine
		}
		return expanded
	}

	for {
		getSyntaxHighlighting()
		getShowHintText()
		cachedHintText = []rune{}

		if nLines > 1 {
			getMultilinePrompt(nLines)
		} else {
			block = []rune{}
			getPrompt()
		}

		line, err := Prompt.Readline()
		if err != nil {
			switch err.Error() {
			case readline.ErrCtrlC:
				merged = ""
				nLines = 1
				fmt.Println(PromptSIGINT)
				continue
			case readline.ErrEOF:
				fmt.Println(utils.NewLineString)
				//return
				os.Exit(0)
			default:
				panic(err)
			}
		}

		if nLines > 1 {
			block = append(block, []rune(utils.NewLineString+line)...)
		} else {
			block = []rune(line)
		}

		expanded, err := history.ExpandVariables(block, Prompt)
		if err != nil {
			//ansi.Stderrln(proc.ShellProcess, ansi.FgRed, err.Error())
			proc.ShellProcess.Stderr.Writeln([]byte(err.Error()))
			merged = ""
			nLines = 1
			continue
		}

		if string(expanded) != string(block) {
			os.Stdout.WriteString(ansi.ExpandConsts("{GREEN}") + string(expanded) + ansi.ExpandConsts("{RESET}") + utils.NewLineString)
		}

		pt, _ := parse(block)
		switch {
		case pt.NestedBlock > 0:
			nLines++
			merged += line + `^\n`

		case pt.Escaped:
			nLines++
			merged += line[:len(line)-1] + `^\n`

		case pt.QuoteSingle, pt.QuoteBrace > 0:
			nLines++
			merged += line + `^\n`

		case pt.QuoteDouble:
			nLines++
			merged += line + `\n`

		case len(block) == 0:
			continue

		default:
			merged += line
			mergedExp, err := history.ExpandVariablesInLine([]rune(merged), Prompt)
			if err == nil {
				merged = string(mergedExp)
			}

			Prompt.History.Write(merged)

			nLines = 1
			merged = ""

			lang.ShellExitNum, _ = lang.RunBlockShellConfigSpaceWithPrompt(expanded, nil, new(streams.TermOut), streams.NewTermErr(ansi.IsAllowed()), thisProc)
			streams.CrLf.Write()

			if PromptGoProc.NotEqual(thisProc) {
				return
			}
		}
	}
}

func getSyntaxHighlighting() {
	highlight, err := proc.ShellProcess.Config.Get("shell", "syntax-highlighting", types.Boolean)
	if err != nil {
		highlight = false
	}
	if highlight.(bool) == true {
		Prompt.SyntaxHighlighter = syntaxHighlight
	} else {
		Prompt.SyntaxHighlighter = nil
	}
}

func getShowHintText() {
	showHintText, err := proc.ShellProcess.Config.Get("shell", "show-hint-text", types.Boolean)
	if err != nil {
		showHintText = false
	}
	if showHintText.(bool) == true {
		Prompt.HintText = hintText
	} else {
		Prompt.HintText = nil
	}
}

type mutexCounter struct {
	i int
	m sync.Mutex
}

func (mc *mutexCounter) Add() int {
	mc.m.Lock()
	defer mc.m.Unlock()

	mc.i++
	return mc.i
}

func (mc *mutexCounter) Set(i int) {
	mc.m.Lock()
	mc.i = i
	mc.m.Unlock()
}

func (mc *mutexCounter) NotEqual(i int) bool {
	mc.m.Lock()
	defer mc.m.Unlock()

	//debug.Log(mc.i, i)

	if mc.i != i {
		return true
	}

	return false
}
