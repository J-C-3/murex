package typemgmt

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/Knetic/govaluate"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	//lang.GoFunctions["="] = cmdEqu
	//lang.GoFunctions["let"] = cmdLet

	lang.DefineMethod("=", cmdEqu, types.Math, types.Math)
	lang.DefineMethod("let", cmdLet, types.Math, types.Null)
}

var (
	rxLet   = regexp.MustCompile(`^([_a-zA-Z0-9]+)\s*=(.*)$`)
	rxMinus = regexp.MustCompile(`^([_a-zA-Z0-9]+)--$`)
	rxPlus  = regexp.MustCompile(`^([_a-zA-Z0-9]+)\+\+$`)
)

func cmdEqu(p *lang.Process) (err error) {
	if p.Parameters.Len() == 0 {
		return errors.New("Missing expression")
	}

	var leftSide string

	if p.IsMethod {
		if !debug.Enabled {
			defer func() {
				if r := recover(); r != nil {
					p.ExitNum = 2
					err = errors.New(fmt.Sprint("Panic caught: ", r))
				}
			}()
		}

		dt := p.Stdin.GetDataType()
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}

		v, err := types.ConvertGoType(b, dt)
		if err != nil {
			return err
		}

		switch dt {
		case types.Integer:
			leftSide = strconv.Itoa(v.(int))

		case types.Float, types.Number:
			leftSide = types.FloatToString(v.(float64))

		case types.Boolean:
			if v.(bool) {
				leftSide = "true"
			} else {
				leftSide = "false"
			}

		default:
			leftSide = `"` + v.(string) + `"`
		}
	}

	value, dt, err := evaluate(p, leftSide+p.Parameters.StringAll())
	if err != nil {
		return err
	}

	s, err := types.ConvertGoType(value, types.String)
	if err != nil {
		return fmt.Errorf("Unable to convert result to text: %s", err.Error())
	}

	p.Stdout.SetDataType(dt)
	_, err = p.Stdout.Write([]byte(s.(string)))
	return err
}

func cmdLet(p *lang.Process) (err error) {
	//p.Stdout.SetDataType(types.Null)

	if !debug.Enabled {
		defer func() {
			if r := recover(); r != nil {
				p.ExitNum = 2
				err = errors.New(fmt.Sprint("Panic caught: ", r))
			}
		}()
	}

	params := p.Parameters.StringAll()
	var variable, expression string

	switch {
	case rxLet.MatchString(params):
		match := rxLet.FindAllStringSubmatch(params, -1)
		variable = match[0][1]
		expression = match[0][2]

	case rxPlus.MatchString(params):
		match := rxPlus.FindAllStringSubmatch(params, -1)
		variable = match[0][1]
		expression = variable + "+1"

	case rxMinus.MatchString(params):
		match := rxMinus.FindAllStringSubmatch(params, -1)
		variable = match[0][1]
		expression = variable + "-1"

	default:
		return errors.New("Invalid syntax for `let`. Should be `let variable-name = expression`")
	}

	value, dt, err := evaluate(p, expression)
	if err != nil {
		return err
	}

	err = p.Variables.Set(p, variable, value, dt)
	return err
}

func evaluate(p *lang.Process, expression string) (value interface{}, dataType string, err error) {
	if !debug.Enabled {
		defer func() {
			if r := recover(); r != nil {
				p.ExitNum = 2
				err = errors.New(fmt.Sprint("Panic caught: ", r))
			}
		}()
	}

	eval, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return
	}

	value, err = eval.Evaluate(lang.DumpVariables(p))
	if err != nil {
		return
	}

	dataType = types.DataTypeFromInterface(value)

	return
}
