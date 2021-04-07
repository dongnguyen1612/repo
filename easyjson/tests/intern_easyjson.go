// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package tests

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC56581f4DecodeGithubComMailruEasyjsonTests(in *jlexer.Lexer, out *NoIntern) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "field":
			out.Field = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC56581f4EncodeGithubComMailruEasyjsonTests(out *jwriter.Writer, in NoIntern) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"field\":"
		out.RawString(prefix[1:])
		out.String(string(in.Field))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NoIntern) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC56581f4EncodeGithubComMailruEasyjsonTests(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NoIntern) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC56581f4EncodeGithubComMailruEasyjsonTests(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NoIntern) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC56581f4DecodeGithubComMailruEasyjsonTests(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NoIntern) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC56581f4DecodeGithubComMailruEasyjsonTests(l, v)
}
func easyjsonC56581f4DecodeGithubComMailruEasyjsonTests1(in *jlexer.Lexer, out *Intern) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "field":
			out.Field = string(in.StringIntern())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC56581f4EncodeGithubComMailruEasyjsonTests1(out *jwriter.Writer, in Intern) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"field\":"
		out.RawString(prefix[1:])
		out.String(string(in.Field))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Intern) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC56581f4EncodeGithubComMailruEasyjsonTests1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Intern) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC56581f4EncodeGithubComMailruEasyjsonTests1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Intern) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC56581f4DecodeGithubComMailruEasyjsonTests1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Intern) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC56581f4DecodeGithubComMailruEasyjsonTests1(l, v)
}