package elog

import "testing"

func TestOverwrite(t *testing.T) {
	t.Parallel()

	under := Metadata{
		"a": Metadata{
			"b": Metadata{
				"c": Metadata{
					"d":     "the original value",
					"other": "SURVIVE",
				},
			},
		},
	}
	over := Metadata{
		"a": Metadata{
			"b": Metadata{
				"c": Metadata{
					"d": "a new value",
				},
			},
		},
	}

	out := DeepMerge(under, over)

	dval := out["a"].(Metadata)["b"].(Metadata)["c"].(Metadata)["d"].(string)
	if dval != "a new value" {
		t.Fatal(dval)
	}
	surv := out["a"].(Metadata)["b"].(Metadata)["c"].(Metadata)["other"].(string)
	if surv != "SURVIVE" {
		t.Fatal(surv)
	}
}

func TestMarshalJSON(t *testing.T) {
	bs, _ := Metadata{"a": "b"}.JsonString()
	t.Log(bs)
}

func TestMetadataIsLoggable(t *testing.T) {
	func(l Loggable) {
	}(Metadata{})
}
