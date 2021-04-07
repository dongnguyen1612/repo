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

func easyjson4b43ff1fDecodeGithubComMailruEasyjsonTests(in *jlexer.Lexer, out *CustomMapKeyType) {
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
		case "Map":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Map = make(map[customKeyType]int)
				for !in.IsDelim('}') {
					var key customKeyType
					if data := in.Raw(); in.Ok() {
						in.AddError((key).UnmarshalJSON(data))
					}
					in.WantColon()
					var v1 int
					v1 = int(in.Int())
					(out.Map)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
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
func easyjson4b43ff1fEncodeGithubComMailruEasyjsonTests(out *jwriter.Writer, in CustomMapKeyType) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Map\":"
		out.RawString(prefix[1:])
		if in.Map == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v2First := true
			for v2Name, v2Value := range in.Map {
				if v2First {
					v2First = false
				} else {
					out.RawByte(',')
				}
				out.Raw((v2Name).MarshalJSON())
				out.RawByte(':')
				out.Int(int(v2Value))
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CustomMapKeyType) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4b43ff1fEncodeGithubComMailruEasyjsonTests(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CustomMapKeyType) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4b43ff1fEncodeGithubComMailruEasyjsonTests(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CustomMapKeyType) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4b43ff1fDecodeGithubComMailruEasyjsonTests(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CustomMapKeyType) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4b43ff1fDecodeGithubComMailruEasyjsonTests(l, v)
}