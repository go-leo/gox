package mapstructure

import (
	"github.com/mitchellh/mapstructure"

	"github.com/go-leo/gox/encodingx"
)

// options is the configuration that is used to create a new decoder
// and allows customization of various aspects of decoding.
type options struct {
	DecoderConfig        mapstructure.DecoderConfig
	DecodeHooks          []mapstructure.DecodeHookFunc
	DecodeHookFuncTypes  []mapstructure.DecodeHookFuncType
	DecodeHookFuncKinds  []mapstructure.DecodeHookFuncKind
	DecodeHookFuncValues []mapstructure.DecodeHookFuncValue
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {
	for _, f := range o.DecodeHookFuncTypes {
		o.DecodeHooks = append(o.DecodeHooks, f)
	}
	for _, f := range o.DecodeHookFuncKinds {
		o.DecodeHooks = append(o.DecodeHooks, f)
	}
	for _, f := range o.DecodeHookFuncValues {
		o.DecodeHooks = append(o.DecodeHooks, f)
	}
	if len(o.DecodeHooks) > 0 {
		o.DecoderConfig.DecodeHook = mapstructure.ComposeDecodeHookFunc(o.DecodeHooks...)
	}
}

type Option func(*options)

// ErrorUnused If set to true, then it is an error for there to exist
// keys in the original map that were unused in the decoding process
// (extra keys).
func ErrorUnused() Option {
	return func(o *options) {
		o.DecoderConfig.ErrorUnused = true
	}
}

// ErrorUnset If set to true, then it is an error for there to exist
// fields in the result that were not set in the decoding process
// (extra fields). This only applies to decoding to a struct. This
// will affect all nested structs as well.
func ErrorUnset() Option {
	return func(o *options) {
		o.DecoderConfig.ErrorUnused = true
	}
}

// ZeroFields if set to true, will zero fields before writing them.
// For example, a map will be emptied before decoded values are put in
// it. If this is false, a map will be merged.
func ZeroFields() Option {
	return func(o *options) {
		o.DecoderConfig.ZeroFields = true
	}
}

// WeaklyTypedInput if set to true, the decoder will make the following
// "weak" conversions:
//
//   - bools to string (true = "1", false = "0")
//   - numbers to string (base 10)
//   - bools to int/uint (true = 1, false = 0)
//   - strings to int/uint (base implied by prefix)
//   - int to bool (true if value != 0)
//   - string to bool (accepts: 1, t, T, TRUE, true, True, 0, f, F,
//     FALSE, false, False. Anything else is an error)
//   - empty array = empty map and vice versa
//   - negative numbers to overflowed uint values (base 10)
//   - slice of maps to a merged map
//   - single values are converted to slices if required. Each
//     element is weakly decoded. For example: "4" can become []int{4}
//     if the target type is an int slice.
func WeaklyTypedInput() Option {
	return func(o *options) {
		o.DecoderConfig.WeaklyTypedInput = true
	}
}

// Squash will squash embedded structs.  A squash tag may also be
// added to an individual struct field using a tag.  For example:
//
//	type Parent struct {
//	    Child `mapstructure:",squash"`
//	}
func Squash() Option {
	return func(o *options) {
		o.DecoderConfig.Squash = true
	}
}

// Metadata is the struct that will contain extra metadata about
// the decoding. If this is nil, then no metadata will be tracked.
func Metadata(md *mapstructure.Metadata) Option {
	return func(o *options) {
		o.DecoderConfig.Metadata = md
	}
}

// TagName the tag name that mapstructure reads for field names. This
// defaults to "mapstructure"
func TagName(name string) Option {
	return func(o *options) {
		o.DecoderConfig.TagName = name
	}
}

// IgnoreUntaggedFields ignores all struct fields without explicit
// TagName, comparable to `mapstructure:"-"` as default behaviour.
func IgnoreUntaggedFields() Option {
	return func(o *options) {
		o.DecoderConfig.IgnoreUntaggedFields = true
	}
}

// MatchName is the function used to match the map key to the struct
// field name or tag. Defaults to `strings.EqualFold`. This can be used
// to implement case-sensitive tag values, support snake casing, etc.
func MatchName(f func(mapKey, fieldName string) bool) Option {
	return func(o *options) {
		o.DecoderConfig.MatchName = f
	}
}

// DecodeHookFuncTypes is a DecodeHookFunc which has complete information about the source and target types.
func DecodeHookFuncTypes(fs ...mapstructure.DecodeHookFuncType) Option {
	return func(o *options) {
		o.DecodeHookFuncTypes = append(o.DecodeHookFuncTypes, fs...)
	}
}

// DecodeHookFuncKinds is a DecodeHookFunc which knows only the Kinds of the source and target types.
func DecodeHookFuncKinds(fs ...mapstructure.DecodeHookFuncKind) Option {
	return func(o *options) {
		o.DecodeHookFuncKinds = append(o.DecodeHookFuncKinds, fs...)
	}
}

// DecodeHookFuncValues is a DecodeHookFunc which has complete access to both the source and target values.
func DecodeHookFuncValues(fs ...mapstructure.DecodeHookFuncValue) Option {
	return func(o *options) {
		o.DecodeHookFuncValues = append(o.DecodeHookFuncValues, fs...)
	}
}

// Unmarshal takes an input structure and uses reflection to translate it to
// the output structure. output must be a pointer to a map or struct.
func Unmarshal(input, output any, opts ...Option) error {
	o := &options{
		DecoderConfig: mapstructure.DecoderConfig{
			Result: output,
		},
	}
	o.apply(opts...)
	o.init()
	decoder, err := mapstructure.NewDecoder(&o.DecoderConfig)
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}

// NewDecoder returns a new decoder.
func NewDecoder(output any, opts ...Option) encodingx.Decoder {
	o := &options{
		DecoderConfig: mapstructure.DecoderConfig{
			Result: output,
		},
	}
	o.apply(opts...)
	o.init()
	return &decoder{options: o, output: output}
}

type decoder struct {
	output  any
	options *options
}

func (d *decoder) Decode(input interface{}) error {
	decoder, err := mapstructure.NewDecoder(&d.options.DecoderConfig)
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}
