package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`!if`:             `if`,
	`!catch`:          `catch`,
	`!set`:            `set`,
	`!event`:          `event`,
	`(`:               `brace-quote`,
	`echo`:            `out`,
	`!and`:            `and`,
	`!or`:             `or`,
	`tread`:           `tread`,
	`try`:             `try`,
	`event`:           `event`,
	`swivel-datatype`: `swivel-datatype`,
	`read`:            `read`,
	`rx`:              `rx`,
	`or`:              `or`,
	`set`:             `set`,
	`out`:             `out`,
	`>`:               `>`,
	`pt`:              `pt`,
	`alter`:           `alter`,
	`append`:          `append`,
	`and`:             `and`,
	`trypipe`:         `trypipe`,
	`swivel-table`:    `swivel-table`,
	`>>`:              `>>`,
	`if`:              `if`,
	`murex-docs`:      `murex-docs`,
	`ttyfd`:           `ttyfd`,
	`unset`:           `unset`,
	`prepend`:         `prepend`,
	`f`:               `f`,
	`tout`:            `tout`,
	`g`:               `g`,
	`get`:             `get`,
	`post`:            `post`,
	`err`:             `err`,
	`catch`:           `catch`,
	`getfile`:         `getfile`,
	`brace-quote`:     `brace-quote`,
}
