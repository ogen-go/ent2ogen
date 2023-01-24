// Code generated by ogen, DO NOT EDIT.

package openapi

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

// Encode encodes string as json.
func (o NilString) Encode(e *jx.Encoder) {
	if o.Null {
		e.Null()
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *NilString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode NilString to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v string
		o.Value = v
		o.Null = true
		return nil
	}
	o.Null = false
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s NilString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *NilString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes bool as json.
func (o OptNilBool) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.Bool(bool(o.Value))
}

// Decode decodes bool from json.
func (o *OptNilBool) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilBool to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v bool
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	v, err := d.Bool()
	if err != nil {
		return err
	}
	o.Value = bool(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilBool) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilBool) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes SchemaAOptionalNullableEnum as json.
func (o OptNilSchemaAOptionalNullableEnum) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes SchemaAOptionalNullableEnum from json.
func (o *OptNilSchemaAOptionalNullableEnum) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilSchemaAOptionalNullableEnum to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v SchemaAOptionalNullableEnum
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilSchemaAOptionalNullableEnum) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilSchemaAOptionalNullableEnum) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes SchemaB as json.
func (o OptSchemaB) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes SchemaB from json.
func (o *OptSchemaB) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptSchemaB to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptSchemaB) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptSchemaB) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *SchemaA) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *SchemaA) encodeFields(e *jx.Encoder) {
	{

		e.FieldStart("int64")
		e.Int64(s.Int64)
	}
	{

		e.FieldStart("string_foobar_bind")
		e.Str(s.StringFoobarBind)
	}
	{

		e.FieldStart("string_optional_nullable")
		s.StringOptionalNullable.Encode(e)
	}
	{
		if s.OptionalNullableBool.Set {
			e.FieldStart("optional_nullable_bool")
			s.OptionalNullableBool.Encode(e)
		}
	}
	{

		e.FieldStart("jsontype_strings")
		e.ArrStart()
		for _, elem := range s.JsontypeStrings {
			e.Str(elem)
		}
		e.ArrEnd()
	}
	{
		if s.JsontypeStringsOptional != nil {
			e.FieldStart("jsontype_strings_optional")
			e.ArrStart()
			for _, elem := range s.JsontypeStringsOptional {
				e.Str(elem)
			}
			e.ArrEnd()
		}
	}
	{

		e.FieldStart("jsontype_ints")
		e.ArrStart()
		for _, elem := range s.JsontypeInts {
			e.Int(elem)
		}
		e.ArrEnd()
	}
	{
		if s.JsontypeIntsOptional != nil {
			e.FieldStart("jsontype_ints_optional")
			e.ArrStart()
			for _, elem := range s.JsontypeIntsOptional {
				e.Int(elem)
			}
			e.ArrEnd()
		}
	}
	{

		e.FieldStart("required_enum")
		s.RequiredEnum.Encode(e)
	}
	{
		if s.OptionalNullableEnum.Set {
			e.FieldStart("optional_nullable_enum")
			s.OptionalNullableEnum.Encode(e)
		}
	}
	{

		e.FieldStart("bytes")
		e.Base64(s.Bytes)
	}
	{

		e.FieldStart("edge_schemab_unique_required")
		s.EdgeSchemabUniqueRequired.Encode(e)
	}
	{

		e.FieldStart("edge_schemab_unique_required_bs_bind")
		s.EdgeSchemabUniqueRequiredBsBind.Encode(e)
	}
	{
		if s.EdgeSchemabUniqueOptional.Set {
			e.FieldStart("edge_schemab_unique_optional")
			s.EdgeSchemabUniqueOptional.Encode(e)
		}
	}
	{
		if s.EdgeSchemab != nil {
			e.FieldStart("edge_schemab")
			e.ArrStart()
			for _, elem := range s.EdgeSchemab {
				elem.Encode(e)
			}
			e.ArrEnd()
		}
	}
	{
		if s.EdgeSchemaaRecursive != nil {
			e.FieldStart("edge_schemaa_recursive")
			e.ArrStart()
			for _, elem := range s.EdgeSchemaaRecursive {
				elem.Encode(e)
			}
			e.ArrEnd()
		}
	}
}

var jsonFieldsNameOfSchemaA = [16]string{
	0:  "int64",
	1:  "string_foobar_bind",
	2:  "string_optional_nullable",
	3:  "optional_nullable_bool",
	4:  "jsontype_strings",
	5:  "jsontype_strings_optional",
	6:  "jsontype_ints",
	7:  "jsontype_ints_optional",
	8:  "required_enum",
	9:  "optional_nullable_enum",
	10: "bytes",
	11: "edge_schemab_unique_required",
	12: "edge_schemab_unique_required_bs_bind",
	13: "edge_schemab_unique_optional",
	14: "edge_schemab",
	15: "edge_schemaa_recursive",
}

// Decode decodes SchemaA from json.
func (s *SchemaA) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SchemaA to nil")
	}
	var requiredBitSet [2]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "int64":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int64()
				s.Int64 = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"int64\"")
			}
		case "string_foobar_bind":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.StringFoobarBind = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"string_foobar_bind\"")
			}
		case "string_optional_nullable":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				if err := s.StringOptionalNullable.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"string_optional_nullable\"")
			}
		case "optional_nullable_bool":
			if err := func() error {
				s.OptionalNullableBool.Reset()
				if err := s.OptionalNullableBool.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"optional_nullable_bool\"")
			}
		case "jsontype_strings":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				s.JsontypeStrings = make([]string, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem string
					v, err := d.Str()
					elem = string(v)
					if err != nil {
						return err
					}
					s.JsontypeStrings = append(s.JsontypeStrings, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"jsontype_strings\"")
			}
		case "jsontype_strings_optional":
			if err := func() error {
				s.JsontypeStringsOptional = make([]string, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem string
					v, err := d.Str()
					elem = string(v)
					if err != nil {
						return err
					}
					s.JsontypeStringsOptional = append(s.JsontypeStringsOptional, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"jsontype_strings_optional\"")
			}
		case "jsontype_ints":
			requiredBitSet[0] |= 1 << 6
			if err := func() error {
				s.JsontypeInts = make([]int, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem int
					v, err := d.Int()
					elem = int(v)
					if err != nil {
						return err
					}
					s.JsontypeInts = append(s.JsontypeInts, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"jsontype_ints\"")
			}
		case "jsontype_ints_optional":
			if err := func() error {
				s.JsontypeIntsOptional = make([]int, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem int
					v, err := d.Int()
					elem = int(v)
					if err != nil {
						return err
					}
					s.JsontypeIntsOptional = append(s.JsontypeIntsOptional, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"jsontype_ints_optional\"")
			}
		case "required_enum":
			requiredBitSet[1] |= 1 << 0
			if err := func() error {
				if err := s.RequiredEnum.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"required_enum\"")
			}
		case "optional_nullable_enum":
			if err := func() error {
				s.OptionalNullableEnum.Reset()
				if err := s.OptionalNullableEnum.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"optional_nullable_enum\"")
			}
		case "bytes":
			requiredBitSet[1] |= 1 << 2
			if err := func() error {
				v, err := d.Base64()
				s.Bytes = []byte(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"bytes\"")
			}
		case "edge_schemab_unique_required":
			requiredBitSet[1] |= 1 << 3
			if err := func() error {
				if err := s.EdgeSchemabUniqueRequired.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"edge_schemab_unique_required\"")
			}
		case "edge_schemab_unique_required_bs_bind":
			requiredBitSet[1] |= 1 << 4
			if err := func() error {
				if err := s.EdgeSchemabUniqueRequiredBsBind.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"edge_schemab_unique_required_bs_bind\"")
			}
		case "edge_schemab_unique_optional":
			if err := func() error {
				s.EdgeSchemabUniqueOptional.Reset()
				if err := s.EdgeSchemabUniqueOptional.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"edge_schemab_unique_optional\"")
			}
		case "edge_schemab":
			if err := func() error {
				s.EdgeSchemab = make([]SchemaB, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem SchemaB
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.EdgeSchemab = append(s.EdgeSchemab, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"edge_schemab\"")
			}
		case "edge_schemaa_recursive":
			if err := func() error {
				s.EdgeSchemaaRecursive = make([]SchemaA, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem SchemaA
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.EdgeSchemaaRecursive = append(s.EdgeSchemaaRecursive, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"edge_schemaa_recursive\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SchemaA")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [2]uint8{
		0b01010111,
		0b00011101,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfSchemaA) {
					name = jsonFieldsNameOfSchemaA[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SchemaA) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SchemaA) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes SchemaAOptionalNullableEnum as json.
func (s SchemaAOptionalNullableEnum) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes SchemaAOptionalNullableEnum from json.
func (s *SchemaAOptionalNullableEnum) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SchemaAOptionalNullableEnum to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch SchemaAOptionalNullableEnum(v) {
	case SchemaAOptionalNullableEnumC:
		*s = SchemaAOptionalNullableEnumC
	case SchemaAOptionalNullableEnumD:
		*s = SchemaAOptionalNullableEnumD
	default:
		*s = SchemaAOptionalNullableEnum(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s SchemaAOptionalNullableEnum) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SchemaAOptionalNullableEnum) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes SchemaARequiredEnum as json.
func (s SchemaARequiredEnum) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes SchemaARequiredEnum from json.
func (s *SchemaARequiredEnum) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SchemaARequiredEnum to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch SchemaARequiredEnum(v) {
	case SchemaARequiredEnumA:
		*s = SchemaARequiredEnumA
	case SchemaARequiredEnumB:
		*s = SchemaARequiredEnumB
	default:
		*s = SchemaARequiredEnum(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s SchemaARequiredEnum) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SchemaARequiredEnum) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *SchemaB) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *SchemaB) encodeFields(e *jx.Encoder) {
	{

		e.FieldStart("id")
		e.Int64(s.ID)
	}
}

var jsonFieldsNameOfSchemaB = [1]string{
	0: "id",
}

// Decode decodes SchemaB from json.
func (s *SchemaB) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SchemaB to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int64()
				s.ID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SchemaB")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfSchemaB) {
					name = jsonFieldsNameOfSchemaB[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SchemaB) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SchemaB) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
