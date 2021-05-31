package parse

import (
	"fmt"

	"github.com/libp2p/go-routing-language/syntax"
)

func ParseBool(ctx *ParseCtx, src syntax.Node) (bool, error) {
	x, ok := src.(syntax.Bool)
	if !ok {
		return false, fmt.Errorf("not a bool")
	}
	return x.Value, nil
}

func ParseBytes(ctx *ParseCtx, src syntax.Node) ([]byte, error) {
	x, ok := src.(syntax.Bytes)
	if !ok {
		return nil, fmt.Errorf("not a bytes")
	}
	return x.Bytes, nil
}

func ParseString(ctx *ParseCtx, src syntax.Node) (string, error) {
	xstr, ok := src.(syntax.String)
	if !ok {
		return "", fmt.Errorf("not a string")
	}
	return xstr.Value, nil
}

func ParseInt64(ctx *ParseCtx, src syntax.Node) (int64, error) {
	xint, ok := src.(syntax.Int)
	if !ok {
		return 0, fmt.Errorf("not an int")
	}
	return xint.Int64(), nil
}

func ParseFloat64(ctx *ParseCtx, src syntax.Node) (float64, error) {
	xfloat, ok := src.(syntax.Float)
	if !ok {
		return 0, fmt.Errorf("not a float")
	}
	f, _ := xfloat.Float64()
	return f, nil
}
