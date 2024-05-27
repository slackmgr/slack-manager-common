package common

type FindOption func(*FindOptions)

type FindOptions struct {
	equals    map[string]interface{}
	notEquals map[string]interface{}
}

func NewFindOptions() *FindOptions {
	return &FindOptions{
		equals:    make(map[string]interface{}),
		notEquals: make(map[string]interface{}),
	}
}

func WithKeyEquals(key string, value interface{}) FindOption {
	return func(o *FindOptions) {
		o.equals[key] = value
	}
}

func WithKeyNotEquals(key string, value interface{}) FindOption {
	return func(o *FindOptions) {
		o.notEquals[key] = value
	}
}

func (o *FindOptions) Equals() map[string]interface{} {
	return o.equals
}

func (o *FindOptions) NotEquals() map[string]interface{} {
	return o.notEquals
}
