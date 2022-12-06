// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/uri"
)

// GetKeyboardParams is parameters of getKeyboard operation.
type GetKeyboardParams struct {
	ID int64
}

func unpackGetKeyboardParams(packed middleware.Parameters) (params GetKeyboardParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeGetKeyboardParams(args [1]string, r *http.Request) (params GetKeyboardParams, _ error) {
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: id: parse")
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	return params, nil
}