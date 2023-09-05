/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sqlparser

import (
	"testing"
)

// Ensure there is no corruption from using a pooled yyParserImpl in Parse.
func TestExtendParallelValid(t *testing.T) {
	validExtendSQL := []struct {
		input                string
		output               string
		partialDDL           bool
		ignoreNormalizerTest bool
	}{
		//{
		//	input: "select 1 from t where a ilike b",
		//},
		{
			input: "SELECT distinct_id FROM Transfer WHERE timestamp >= toDateTime(now() - 86400) AND timestamp < toDateTime(now()) EXCEPT SELECT distinct_id FROM Transfer WHERE timestamp < toDateTime(now() - 86400)",
		},
	}

	for i := 0; i < len(validExtendSQL); i++ {
		tcase := validExtendSQL[i]
		if tcase.output == "" {
			tcase.output = tcase.input
		}
		tree, err := Parse(tcase.input)
		if err != nil {
			t.Errorf("Parse(%q) err: %v, want nil", tcase.input, err)
			continue
		}
		out := String(tree)
		if out != tcase.output {
			t.Errorf("Parse(%q) = %q, want: %q", tcase.input, out, tcase.output)
		}
	}
}
