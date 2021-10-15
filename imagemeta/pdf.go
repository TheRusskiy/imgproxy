package imagemeta

import (
	"bytes"
	"io/ioutil"

	//"bytes"
	//"encoding/binary"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	//"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

var pdfHeader = []byte("\x25\x50\x44\x46")

type PdfFormatError string

func (e PdfFormatError) Error() string { return "invalid PDF format: " + string(e) }

func DecodePdfMeta(r io.Reader) (Meta, error) {
	rsb, err := ioutil.ReadAll(r)

	if err != nil {
		return nil, err
	}

	rs := bytes.NewReader(rsb)

	ctx, err := api.ReadContext(rs, nil)

	if err != nil {
		return nil, err
	}

	dimensions, err := ctx.PageDims()

	if err != nil {
		return nil, err
	}

	d := dimensions[0]

	return &meta{
		format: "pdf",
		width:  int(d.Width),
		height: int(d.Height),
	}, nil
}

func init() {
	RegisterFormat(string(pdfHeader), DecodePdfMeta)
}
