package negotiate

import (
	"io"

	"github.com/unrolled/render"
)

// Respond switches on Content Type to determine
// the response
func (n *Negotiator) Respond(w io.Writer, status int, v interface{}) {
	switch n.ContentType {
	case render.ContentJSON:
		err := n.Render.JSON(w, status, v)
		if err != nil {
			panic(err)
		}
	case render.ContentXML:
		err := n.Render.XML(w, status, v)
		if err != nil {
			panic(err)
		}
	default:
		err := n.Render.JSON(w, status, v)
		if err != nil {
			panic(err)
		}
	}
}
