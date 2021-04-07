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

func easyjson35b73998DecodeGithubComMailruEasyjsonTests(in *jlexer.Lexer, out *OmitEmptyDefault) {
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
		case "Field":
			out.Field = string(in.String())
		case "Str":
			out.Str = string(in.String())
		case "s":
			out.Str1 = string(in.String())
		case "Str2":
			out.Str2 = string(in.String())
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
func easyjson35b73998EncodeGithubComMailruEasyjsonTests(out *jwriter.Writer, in OmitEmptyDefault) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Field != "" {
		const prefix string = ",\"Field\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Field))
	}
	if in.Str != "" {
		const prefix string = ",\"Str\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Str))
	}
	{
		const prefix string = ",\"s\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Str1))
	}
	{
		const prefix string = ",\"Str2\":"
		out.RawString(prefix)
		out.String(string(in.Str2))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v OmitEmptyDefault) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson35b73998EncodeGithubComMailruEasyjsonTests(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OmitEmptyDefault) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson35b73998EncodeGithubComMailruEasyjsonTests(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OmitEmptyDefault) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson35b73998DecodeGithubComMailruEasyjsonTests(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OmitEmptyDefault) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson35b73998DecodeGithubComMailruEasyjsonTests(l, v)
}