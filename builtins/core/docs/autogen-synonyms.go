package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`!event`:          `event`,
	`echo`:            `out`,
	`!or`:             `or`,
	`!export`:         `export`,
	`unset`:           `export`,
	`!global`:         `global`,
	`!set`:            `set`,
	`(`:               `brace-quote`,
	`!and`:            `and`,
	`!if`:             `if`,
	`!catch`:          `catch`,
	`get`:             `get`,
	`f`:               `f`,
	`g`:               `g`,
	`try`:             `try`,
	`err`:             `err`,
	`rx`:              `rx`,
	`read`:            `read`,
	`or`:              `or`,
	`tread`:           `tread`,
	`if`:              `if`,
	`append`:          `append`,
	`swivel-table`:    `swivel-table`,
	`>>`:              `>>`,
	`ttyfd`:           `ttyfd`,
	`tout`:            `tout`,
	`pt`:              `pt`,
	`alter`:           `alter`,
	`prepend`:         `prepend`,
	`post`:            `post`,
	`out`:             `out`,
	`brace-quote`:     `brace-quote`,
	`trypipe`:         `trypipe`,
	`event`:           `event`,
	`getfile`:         `getfile`,
	`>`:               `>`,
	`global`:          `global`,
	`set`:             `set`,
	`swivel-datatype`: `swivel-datatype`,
	`and`:             `and`,
	`murex-docs`:      `murex-docs`,
	`catch`:           `catch`,
	`export`:          `export`,
}
