package cloudformation

import "encoding/json"

// Join returns a new instance of JoinFunc that joins items with separator.
func Join(separator string, items ...StringExpr) JoinFunc {
	return JoinFunc{Separator: separator, Items: items}
}

// JoinFunc represents an invocation of the Fn::Join intrinsic.
//
// The intrinsic function Fn::Join appends a set of values into a single
// value, separated by the specified delimiter. If a delimiter is the empty
// string, the set of values are concatenated with no delimiter.
//
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-join.html
type JoinFunc struct {
	Separator string
	Items     []StringExpr
}

// MarshalJSON returns a JSON representation of the object
func (f JoinFunc) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FnJoin []interface{} `json:"Fn::Join"`
	}{FnJoin: []interface{}{f.Separator, f.Items}})
}

// UnmarshalJSON sets the object from the provided JSON representation
func (f *JoinFunc) UnmarshalJSON(buf []byte) error {
	v := struct {
		FnJoin [2]json.RawMessage `json:"Fn::Join"`
	}{}
	if err := json.Unmarshal(buf, &v); err != nil {
		return err
	}
	if err := json.Unmarshal(v.FnJoin[0], &f.Separator); err != nil {
		return err
	}
	if err := json.Unmarshal(v.FnJoin[1], &f.Items); err != nil {
		return err
	}

	return nil
}

func (f JoinFunc) String() *StringExpr {
	return &StringExpr{Func: f}
}

var _ StringFunc = JoinFunc{} // JoinFunc must implement StringFunc