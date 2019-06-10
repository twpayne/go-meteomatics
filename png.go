package meteomatics

import "context"

// PNGRequest requests a PNG.
func (c *Client) PNGRequest(ctx context.Context, ts TimeStringer, ps ParameterStringer, ls LocationStringer, options *RequestOptions) ([]byte, error) {
	return c.RawRequest(ctx, ts, ps, ls, FormatPNG, options)
}
