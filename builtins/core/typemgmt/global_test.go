package typemgmt

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/io"
	"github.com/lmorg/murex/lang"
)

func TestGlobalFunctionPositive(t *testing.T) {
	lang.InitEnv()

	set := []Test{
		{
			Block:    "global: f=b",
			Name:     "f",
			Value:    "b",
			DataType: "str",
		},
		{
			Block:    "global: foo=b",
			Name:     "foo",
			Value:    "b",
			DataType: "str",
		},
		{
			Block:    "global: f=bar",
			Name:     "f",
			Value:    "bar",
			DataType: "str",
		},
		{
			Block:    "global: foo=bar",
			Name:     "foo",
			Value:    "bar",
			DataType: "str",
		},
		{
			Block:    "global: _b=foobar",
			Name:     "_b",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "global: f_=foobar",
			Name:     "f_",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "global: f_b=foobar",
			Name:     "f_b",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "global: foo_b=foobar",
			Name:     "foo_b",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "global: f_bar=foobar",
			Name:     "f_bar",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "global: foo_bar=foobar",
			Name:     "foo_bar",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "global: foobar=foobar",
			Name:     "foobar",
			Value:    "foobar",
			DataType: "str",
		},
	}

	unset := []string{
		"f",
		"foo",
		"_b",
		"f_",
		"f_b",
		"foo_b",
		"f_bar",
		"foo_bar",
		"foobar",
	}

	VariableTests(set, t)
	UnSetTests("!global", unset, t)
}

func TestGlobalMethodPositive(t *testing.T) {
	lang.InitEnv()

	set := []Test{
		{
			Block:    "out: b -> global: f",
			Name:     "f",
			Value:    "b",
			DataType: "str",
		},
		{
			Block:    "out: b -> global: foo",
			Name:     "foo",
			Value:    "b",
			DataType: "str",
		},
		{
			Block:    "out: bar -> global: f",
			Name:     "f",
			Value:    "bar",
			DataType: "str",
		},
		{
			Block:    "out: bar -> global: foo",
			Name:     "foo",
			Value:    "bar",
			DataType: "str",
		},
		{
			Block:    "out: foobar -> global: _b",
			Name:     "_b",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "out: foobar -> global: f_",
			Name:     "f_",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "out: foobar -> global: f_b",
			Name:     "f_b",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "out: foobar -> global: foo_b",
			Name:     "foo_b",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "out: foobar -> global: f_bar",
			Name:     "f_bar",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "out: foobar -> global: foo_bar",
			Name:     "foo_bar",
			Value:    "foobar",
			DataType: "str",
		},
		{
			Block:    "out: foobar -> global: foobar",
			Name:     "foobar",
			Value:    "foobar",
			DataType: "str",
		},
	}

	unset := []string{
		"f",
		"foo",
		"_b",
		"f_",
		"f_b",
		"foo_b",
		"f_bar",
		"foo_bar",
		"foobar",
	}

	VariableTests(set, t)
	UnSetTests("!global", unset, t)
}

func TestGlobalFunctionNegative(t *testing.T) {
	lang.InitEnv()

	tests := []Test{
		/*{ TODO: this should fail
			Block: "global: =foobar",
			Fail:  true,
		},*/
		{
			Block: "global: -=foobar",
			Fail:  true,
		},
		{
			Block: "global: _=foobar",
			Fail:  true,
		},
		{
			Block: "global: foo-bar=foobar",
			Fail:  true,
		},
		{
			Block: "global: foo\\-bar=foobar",
			Fail:  true,
		},
	}

	VariableTests(tests, t)
}

func TestGlobalMethodNegative(t *testing.T) {
	lang.InitEnv()

	tests := []Test{
		{
			Block: "out: foobar -> set",
			Fail:  true,
		},
		/*{ // TODO: this should fail
			Block: "out: foobar -> global: =",
			Fail:  true,
		},*/
		{
			Block: "out: foobar -> global: -",
			Fail:  true,
		},
		{
			Block: "out: foobar -> global: _",
			Fail:  true,
		},
		{
			Block: "out: foobar -> global: foo-bar",
			Fail:  true,
		},
		{
			Block: "out: foobar -> global: foo\\-bar",
			Fail:  true,
		},
		{
			Block: "out: foobar -> global: foo=bar",
			Fail:  true,
		},
	}

	VariableTests(tests, t)
}

func TestGlobalFunctionDataTypes(t *testing.T) {
	lang.InitEnv()

	set := []Test{
		{
			Block:    "global: foobar=123",
			Name:     "foobar",
			Value:    "123",
			DataType: "str",
		},
		{
			Block:    "global: foobar=123.456",
			Name:     "foobar",
			Value:    "123.456",
			DataType: "str",
		},
		{
			Block:    "global: foobar=true",
			Name:     "foobar",
			Value:    "true",
			DataType: "str",
		},
		{
			Block:    "global: foobar=false",
			Name:     "foobar",
			Value:    "false",
			DataType: "str",
		},
		{
			Block:    "global: int foobar=123",
			Name:     "foobar",
			Value:    "123",
			DataType: "int",
		},
		{
			Block:    "global: num foobar=123.456",
			Name:     "foobar",
			Value:    "123.456",
			DataType: "num",
		},
		{
			Block:    "global: bool foobar=true",
			Name:     "foobar",
			Value:    "true",
			DataType: "bool",
		},
		{
			Block:    "global: bool foobar=false",
			Name:     "foobar",
			Value:    "false",
			DataType: "bool",
		},
	}

	unset := []string{
		"foobar",
	}

	VariableTests(set, t)
	UnSetTests("!global", unset, t)
}

func TestGlobalMethodDataTypes(t *testing.T) {
	lang.InitEnv()

	set := []Test{
		{
			Block:    "tout: int 123 -> global: foobar",
			Name:     "foobar",
			Value:    "123",
			DataType: "int",
		},
		{
			Block:    "tout: num 123.456 -> global: foobar",
			Name:     "foobar",
			Value:    "123.456",
			DataType: "num",
		},
		{
			Block:    "tout: bool true -> global: foobar",
			Name:     "foobar",
			Value:    "true",
			DataType: "bool",
		},
		{
			Block:    "tout: bool false -> global: foobar",
			Name:     "foobar",
			Value:    "false",
			DataType: "bool",
		},
		{
			Block:    "out: 123 -> global: int foobar",
			Name:     "foobar",
			Value:    "123",
			DataType: "int",
		},
		{
			Block:    "out: 123.456 -> global: num foobar",
			Name:     "foobar",
			Value:    "123.456",
			DataType: "num",
		},
		{
			Block:    "tout: int true -> global: bool foobar",
			Name:     "foobar",
			Value:    "true",
			DataType: "bool",
		},
		{
			Block:    "out: false -> global: bool foobar",
			Name:     "foobar",
			Value:    "false",
			DataType: "bool",
		},
	}

	unset := []string{
		"foobar",
	}

	VariableTests(set, t)
	UnSetTests("!global", unset, t)
}
