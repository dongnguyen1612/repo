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

func easyjson82f18266DecodeGithubComMailruEasyjsonTests(in *jlexer.Lexer, out *MembersUnescaped) {
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
		key := in.UnsafeFieldName(true)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "漢語":
			out.A = string(in.String())
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
func easyjson82f18266EncodeGithubComMailruEasyjsonTests(out *jwriter.Writer, in MembersUnescaped) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"漢語\":"
		out.RawString(prefix[1:])
		out.String(string(in.A))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MembersUnescaped) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson82f18266EncodeGithubComMailruEasyjsonTests(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MembersUnescaped) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson82f18266EncodeGithubComMailruEasyjsonTests(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MembersUnescaped) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson82f18266DecodeGithubComMailruEasyjsonTests(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MembersUnescaped) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson82f18266DecodeGithubComMailruEasyjsonTests(l, v)
}
