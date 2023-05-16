package management_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/index"
	_ "github.com/lmorg/murex/builtins/core/io"
	_ "github.com/lmorg/murex/builtins/core/management"
	_ "github.com/lmorg/murex/builtins/core/runtime"
	_ "github.com/lmorg/murex/builtins/types/json"
	"github.com/lmorg/murex/test"
)

func TestSummary(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `out: part1
					err: part1
					summary: foobar test
					runtime --summaries -> [ foobar ]
					out: part2
					err: part2
					!summary foobar
					runtime --summaries -> [ foobar ]`,
			Stdout:  "^part1\ntestpart2\n$",
			Stderr:  "^part1\npart2\nError",
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}

func TestSource(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `tout block { out "Hello, world!" } -> source`,
			Stdout: "Hello, world!\n",
		},
		{
			Block:  `source { out "Hello, world!" }`,
			Stdout: "Hello, world!\n",
		},
		{
			Block:  `source source_test.mx`,
			Stdout: "Hello, world!\n",
		},
	}

	test.RunMurexTests(tests, t)
}
