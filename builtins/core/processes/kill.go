package processes

import (
	"fmt"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["fid-kill"] = cmdFidKill
	lang.GoFunctions["fid-killall"] = cmdKillAll

	defaults.AppendProfile(`
private autocomplete.fid-kill {
	# Returns a 2D array FIDs 
    fid-list: --jsonl -> [ FID Command Parameters ]
}

test unit private autocomplete.fid-kill {
    "StdoutRegex": (^\["FID","Command","Parameters"\]\n.*(["[0-9]+",")),
	"StdoutType":  "jsonl",
	"ExitNum": 0,
	"StdoutBlock": ({
        -> len -> set len;
        if { = len>0 } then {
            out "Len greater than 0"
        } else {
            err "No elements returned"
        }
    })
}

autocomplete: set fid-kill {
    [{
        "DynamicDesc": ({ autocomplete.fid-kill }),
        "ListView": true,
        "AllowMultiple": true
    }]
}
`)
}

func cmdFidKill(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	for i := 0; i < p.Parameters.Len(); i++ {
		fid, err := p.Parameters.Uint32(i)
		if err != nil {
			return err
		}

		process, err := lang.GlobalFIDs.Proc(fid)
		if err != nil {
			return err
		}

		if process.Kill != nil {
			process.Kill()
		} else {
			return fmt.Errorf("fid `%d` cannot be killed. `Kill` method == `nil`", fid)
		}
	}

	return nil
}

func cmdKillAll(*lang.Process) error {
	fids := lang.GlobalFIDs.ListAll()
	for _, p := range fids {
		if p.Kill != nil /*&& !p.HasTerminated()*/ {
			procName := p.Name.String()
			procParam, _ := p.Parameters.String(0)
			if procName == "exec" {
				procName = procParam
				procParam, _ = p.Parameters.String(1)
			}
			if len(procParam) > 10 {
				procParam = procParam[:10]
			}
			lang.ShellProcess.Stderr.Write([]byte(fmt.Sprintf("!!! Sending kill signal to fid %d: %s %s !!!\n", p.Id, procName, procParam)))
			p.Kill()
		}
	}

	return nil
}
