package envsubst

import "os"

// Eval replaces ${var} in the string based on the mapping function.
func Eval(s string, mapping func(string, map[string]string) string, envs map[string]string) (string, error) {
	t, err := Parse(s)
	if err != nil {
		return s, err
	}
	return t.Execute(mapping, envs)
}

// EvalEnv replaces ${var} in the string according to the values of the
// current environment variables. References to undefined variables are
// replaced by the empty string.
func EvalEnv(s string) (string, error) {
	return Eval(s, func(s string, m map[string]string) string {
		return os.Getenv(s)
	}, make(map[string]string))
}
