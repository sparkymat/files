// Code generated by qtc from "list.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line view/list.qtpl:1
package view

//line view/list.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line view/list.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line view/list.qtpl:1
func StreamList(qw422016 *qt422016.Writer) {
//line view/list.qtpl:1
	qw422016.N().S(`
  <h1>Hello</h1>
`)
//line view/list.qtpl:3
}

//line view/list.qtpl:3
func WriteList(qq422016 qtio422016.Writer) {
//line view/list.qtpl:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line view/list.qtpl:3
	StreamList(qw422016)
//line view/list.qtpl:3
	qt422016.ReleaseWriter(qw422016)
//line view/list.qtpl:3
}

//line view/list.qtpl:3
func List() string {
//line view/list.qtpl:3
	qb422016 := qt422016.AcquireByteBuffer()
//line view/list.qtpl:3
	WriteList(qb422016)
//line view/list.qtpl:3
	qs422016 := string(qb422016.B)
//line view/list.qtpl:3
	qt422016.ReleaseByteBuffer(qb422016)
//line view/list.qtpl:3
	return qs422016
//line view/list.qtpl:3
}
