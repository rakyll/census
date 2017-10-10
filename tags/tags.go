package tags

import "context"

type Key interface {
	Name() string
	StringValue(b []byte) string
}

type StringKey struct {
	name string
}

func NewStringKey(name string) (StringKey, error) {
	panic("not implemented")
}

func (s StringKey) Name() string {
	return s.name
}

func (s StringKey) StringValue(b []byte) string {
	panic("not implemented")
}

func (s StringKey) Update(v string) Modifier {
	panic("not implemented")
}

func (s StringKey) Insert(v string) Modifier {
	panic("not implemented")
}

func (s StringKey) Upsert(v string) Modifier {
	panic("not implemented")
}

func Decode(bytes []byte) (*TagSet, error) {
	panic("not implemented")
}

func Encode(ts *TagSet) []byte {
	panic("not implemented")
}

type Tag struct {
	Key   Key
	Value []byte
}

type TagSet struct{}

type Modifier func(t *TagSet) *TagSet

func NewTagSet(orig *TagSet, m ...Modifier) *TagSet {
	panic("not implemented")
}

// FromContext returns the TagSet stored in the context.
func FromContext(ctx context.Context) *TagSet {
	panic("not implemented")
}

// NewContext creates a new context from the old one replacing any existing
// TagSet with the new parameter TagSet ts.
func NewContext(ctx context.Context, ts *TagSet) context.Context {
	panic("not implemented")
}
