package value

import "github.com/go-restruct/restruct/internal/expr/typing"

var (
	_ = Value(Boolean{})
	_ = Comparer(Boolean{})
	_ = Logical(Boolean{})
)

// Boolean represents a boolean constant.
type Boolean struct {
	value bool
}

// NewBoolean creates a new boolean constant.
func NewBoolean(value bool) Boolean {
	return Boolean{value}
}

func (c Boolean) String() string {
	if c.value {
		return "true"
	}
	return "false"
}

// Value implements Value
func (c Boolean) Value() interface{} { return c.value }

// Type implements Value
func (c Boolean) Type() (typing.Type, error) {
	return typing.PrimitiveType(typing.Boolean, typing.TypeInfo{}), nil
}

// Equal implements Comparer
func (c Boolean) Equal(right Comparer) (Value, error) {
	r, ok := right.(Boolean)
	if !ok {
		return nil, ErrTypeMismatch
	}
	return NewBoolean(c.value == r.value), nil
}

// NotEqual implements Comparer
func (c Boolean) NotEqual(right Comparer) (Value, error) {
	r, ok := right.(Boolean)
	if !ok {
		return nil, ErrTypeMismatch
	}
	return NewBoolean(c.value != r.value), nil
}

// LogicalNot implements Logical
func (c Boolean) LogicalNot() (Value, error) {
	return NewBoolean(!c.value), nil
}

// LogicalAnd implements Logical
func (c Boolean) LogicalAnd(right Logical) (Value, error) {
	r, ok := right.(Boolean)
	if !ok {
		return nil, ErrTypeMismatch
	}
	return NewBoolean(c.value && r.value), nil
}

// LogicalOr implements Logical
func (c Boolean) LogicalOr(right Logical) (Value, error) {
	r, ok := right.(Boolean)
	if !ok {
		return nil, ErrTypeMismatch
	}
	return NewBoolean(c.value || r.value), nil
}