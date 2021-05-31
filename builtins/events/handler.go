package events

import (
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

type eventType interface {
	Add(name, interrupt string, block []rune, fileRef *ref.File) (err error)
	Remove(interrupt string) (err error)
	Dump() (dump interface{})
}

var events = make(map[string]eventType)

// AddEventType registers your event type handlers
func AddEventType(eventTypeName string, handlerInterface eventType) {
	events[eventTypeName] = handlerInterface
}

type j struct {
	Name      string
	Interrupt interface{}
}

// Callback is a generic function your event handlers types should hook into so
// murex functions can remain consistent.
func Callback(name string, interrupt interface{}, block []rune, fileRef *ref.File, stdout stdio.Io) {
	if fileRef == nil {
		panic("fileRef should not be nil value")
	}

	json, err := json.Marshal(&j{
		Name:      name,
		Interrupt: interrupt,
	}, false)
	if err != nil {
		lang.ShellProcess.Stderr.Writeln([]byte("error building event input: " + err.Error()))
		return
	}

	fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_NEW_MODULE | lang.F_BACKGROUND | lang.F_CREATE_STDIN)
	fork.Stdin.SetDataType(types.Json)
	fork.Name.Set("(event)")
	fork.FileRef = fileRef
	_, err = fork.Stdin.Write(json)
	if err != nil {
		lang.ShellProcess.Stderr.Writeln([]byte("error writing event input: " + err.Error()))
		return
	}

	debug.Log("Event callback:", string(json), string(block))

	fork.Stdout = stdout

	_, err = fork.Execute(block)
	if err != nil {
		lang.ShellProcess.Stderr.Writeln([]byte("error compiling event callback: " + err.Error()))
	}
}

// DumpEvents is used for `runtime` to output all the saved events
func DumpEvents() (dump map[string]interface{}) {
	dump = make(map[string]interface{})

	for et := range events {
		dump[et] = events[et].Dump()
	}

	return
}
