package docs

// Digest stores a 1 line summary of each builtins
var Digest map[string]string = map[string]string{
	`alter`:           `Change a value within a structured data-type and pass that change along the pipeline without altering the original source input`,
	`swivel-datatype`: `Converts tabulated data into a map of values for serialised data-types such as JSON and YAML`,
	`murex-docs`:      `Displays the man pages for _murex_ builtins`,
	`prepend`:         `Add data to the start of an array`,
	`f`:               `Lists objects (eg files) in the current working directory`,
	`set`:             `Define a variable and set it's value`,
	`ttyfd`:           `Returns the TTY device of the parent.`,
	`err`:             `'echo' a string to the STDERR`,
	`out`:             `'echo' a string to the STDOUT`,
	`try`:             `Handles errors inside a block of code`,
	`print`:           `Write a string to the OS STDOUT (bypassing _murex_ pipelines)`,
	`pt`:              `Pipe telemetry. Writes data-types and bytes written`,
	`catch`:           `Handles the exception code raised by 'try' or 'trypipe'`,
	`post`:            `HTTP POST request with a JSON-parsable return`,
	`rx`:              `Regexp pattern matching for file system objects (eg '.*\.txt')`,
	`trypipe`:         `Checks state of each function in a pipeline and exits block on error`,
	`unset`:           `Deallocates an environmental variable (aliased to '!export')`,
	`event`:           `Event driven programming for shell scripts`,
	`swivel-table`:    `Rotates a table by 90 degrees`,
	`getfile`:         `Makes a standard HTTP request and return the contents as _murex_-aware data type for passing along _murex_ pipelines.`,
	`>`:               `Writes STDIN to disk - overwriting contents if file already exists`,
	`>>`:              `Writes STDIN to disk - appending contents if file already exists`,
	`if`:              `Conditional statement to execute different blocks of code depending on the result of the condition`,
	`append`:          `Add data to the end of an array`,
	`get`:             `Makes a standard HTTP request and returns the result as a JSON object`,
	`tout`:            `'echo' a string to the STDOUT and set it's data-type`,
	`g`:               `Glob pattern matching for file system objects (eg *.txt)`,
}
