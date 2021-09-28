package arraytools

import (
	"errors"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.DefineFunction("map", mkmap, types.Json)
}

func mkmap(p *lang.Process) error {
	p.Stdout.SetDataType(types.Json)

	blockKey, err := p.Parameters.Block(0)
	if err != nil {
		return err
	}

	blockValue, err := p.Parameters.Block(1)
	if err != nil {
		return err
	}

	//debug.Log("block key:", string(blockKey))
	//debug.Log("block value:", string(blockValue))

	//var wg sync.WaitGroup
	var aKeys, aValues []string

	//go func() {
	//	wg.Add(1)
	forkKeys := p.Fork(lang.F_NO_STDIN | lang.F_CREATE_STDOUT)
	_, errKeys := forkKeys.Execute(blockKey)
	//	wg.Done()
	//}()

	//go func() {
	//	wg.Add(1)
	forkValues := p.Fork(lang.F_NO_STDIN | lang.F_CREATE_STDOUT)
	_, errValues := forkValues.Execute(blockValue)
	//	wg.Done()
	//}()

	if errKeys != nil {
		return errKeys
	}
	if errValues != nil {
		return errValues
	}
	//wg.Wait()

	//go func() {
	//	wg.Add(1)
	errKeys = forkKeys.Stdout.ReadArray(func(b []byte) {
		aKeys = append(aKeys, string(b))
	})
	//	wg.Done()
	//}()

	//go func() {
	//	wg.Add(1)
	errValues = forkValues.Stdout.ReadArray(func(b []byte) {
		aValues = append(aValues, string(b))
	})
	//	wg.Done()
	//}()

	if errKeys != nil {
		return errKeys
	}
	if errValues != nil {
		return errValues
	}
	//wg.Wait()

	//debug.Json("a keys", aKeys)
	//debug.Json("a values", aValues)

	if len(aKeys) > len(aValues) {
		return errors.New("There are more keys than values (k > v)")
	}

	if len(aKeys) < len(aValues) {
		return errors.New("There are more values than keys (v > k)")
	}

	m := make(map[string]string)
	for i := range aKeys {
		m[aKeys[i]] = aValues[i]
	}
	//debug.Json("m", m)

	b, err := json.Marshal(m, p.Stdout.IsTTY())
	if err != nil {
		return err
	}
	_, err = p.Stdout.Write(b)
	return err
}
