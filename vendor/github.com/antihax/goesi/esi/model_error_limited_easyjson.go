// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

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

func easyjson4c2ed35dDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *ErrorLimitedList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(ErrorLimitedList, 0, 4)
			} else {
				*out = ErrorLimitedList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 ErrorLimited
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson4c2ed35dEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in ErrorLimitedList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v ErrorLimitedList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4c2ed35dEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ErrorLimitedList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4c2ed35dEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ErrorLimitedList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4c2ed35dDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ErrorLimitedList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4c2ed35dDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson4c2ed35dDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *ErrorLimited) {
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
		case "error":
			out.Error_ = string(in.String())
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
func easyjson4c2ed35dEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in ErrorLimited) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Error_ != "" {
		const prefix string = ",\"error\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Error_))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ErrorLimited) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4c2ed35dEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ErrorLimited) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4c2ed35dEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ErrorLimited) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4c2ed35dDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ErrorLimited) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4c2ed35dDecodeGithubComAntihaxGoesiEsi1(l, v)
}
