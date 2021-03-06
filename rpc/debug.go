// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

/*
	Some HTML presented at http://machine:port/debug/rpc
	Lists services, their methods, and some statistics, still rudimentary.
*/

import (
	"html/template"
)

const debugText = `<html>
	<body>
	<title>Services</title>
	{{range .}}
	<hr>
	Service {{.Name}}
	<hr>
		<table>
		<th align=center>Method</th><th align=center>Calls</th>
		{{range .Method}}
			<tr>
			<td align=left font=fixed>{{.Name}}({{.Type.ArgType}}, {{.Type.ReplyType}}) error</td>
			<td align=center>{{.Type.NumCalls}}</td>
			</tr>
		{{end}}
		</table>
	{{end}}
	</body>
	</html>`

var debug = template.Must(template.New("RPC debug").Parse(debugText))

// If set, print log statements for internal and I/O errors.
var debugLog = false

type debugMethod struct {
	Type *methodType
	Name string
}

type methodArray []debugMethod

type debugService struct {
	Service *service
	Name    string
	Method  methodArray
}

type serviceArray []debugService

func (s serviceArray) Len() int           { return len(s) }
func (s serviceArray) Less(i, j int) bool { return s[i].Name < s[j].Name }
func (s serviceArray) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (m methodArray) Len() int           { return len(m) }
func (m methodArray) Less(i, j int) bool { return m[i].Name < m[j].Name }
func (m methodArray) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

type debugHTTP struct {
	*Server
}
