package cloudformation

import "encoding/json"

// Select returns a new instance of SelectFunc chooses among items via selector.
func Select(selector string, items ...StringExpr) SelectFunc {
	return SelectFunc{Selector: selector, Items: StringListExpr{Literal: items}}
}

// SelectFunc represents an invocation of the Fn::Select intrinsic.
//
// The intrinsic function Fn::Select returns a single object from a
// list of objects by index.
//
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-select.html
type SelectFunc struct {
	Selector string // XXX int?
	Items    StringListExpr
}

// MarshalJSON returns a JSON representation of the object
func (f SelectFunc) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FnSelect []interface{} `json:"Fn::Select"`
	}{FnSelect: []interface{}{f.Selector, f.Items}})
}

// UnmarshalJSON sets the object from the provided JSON representation
func (f *SelectFunc) UnmarshalJSON(buf []byte) error {
	v := struct {
		FnSelect [2]json.RawMessage `json:"Fn::Select"`
	}{}
	if err := json.Unmarshal(buf, &v); err != nil {
		return err
	}
	if err := json.Unmarshal(v.FnSelect[0], &f.Selector); err != nil {
		return err
	}
	if err := json.Unmarshal(v.FnSelect[1], &f.Items); err != nil {
		return err
	}

	return nil
}

func (f SelectFunc) String() *StringExpr {
	return &StringExpr{Func: f}
}

var _ StringFunc = SelectFunc{} // SelectFunc must implement StringFunc