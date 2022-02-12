// +build windows

package defaults

/*
   WARNING:
   --------

   This Go source file has been automatically generated from
   profile_windows.mx using docgen.

   Please do not manually edit this file because it will be automatically
   overwritten by the build pipeline. Instead please edit the aforementioned
   profile_windows.mx file located in the same directory.
*/

func init() {
	murexProfile = append(murexProfile, "autocomplete set go { [{\n    \"Flags\": [ \"build\", \"clean\", \"doc\", \"env\", \"bug\", \"fix\", \"fmt\", \"generate\", \"get\", \"install\", \"list\", \"run\", \"test\", \"tool\", \"version\", \"vet\", \"help\" ]\n}] }\n\nautocomplete set cast { [{\n    \"Dynamic\": ({ runtime: --unmarshallers })\n}] }\n\nautocomplete set tout { [{\n    \"Dynamic\": ({ runtime: --marshallers })\n}] }")
}
