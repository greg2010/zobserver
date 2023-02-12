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

func easyjson89630228DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *DeleteCharactersCharacterIdMailLabelsLabelIdUnprocessableEntityList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(DeleteCharactersCharacterIdMailLabelsLabelIdUnprocessableEntityList, 0, 4)
			} else {
				*out = DeleteCharactersCharacterIdMailLabelsLabelIdUnprocessableEntityList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 DeleteCharactersCharacterIdMailLabelsLabelIdUnprocessableEntity
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
func easyjson89630228EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in DeleteCharactersCharacterIdMailLabelsLabelIdUnprocessableEntityList) {
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
func (v DeleteCharactersCharacterIdMailLabelsLabelIdUnprocessableEntityList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89630228EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeleteCharactersCharacterIdMailLabelsLabelIdUnprocessableEntityList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89630228EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeleteCharactersCharacterIdMailLabelsLabelIdUnprocessableEntityList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89630228DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeleteCharactersCharacterIdMailLabelsLabelIdUnprocessableEntityList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89630228DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson89630228DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *DeleteCharactersCharacterIdMailLabelsLabelIdUnprocessableEntity) {
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
func easyjson89630228EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in DeleteCharactersCharacterIdMailLabelsLabelIdUnprocessableEntity) {
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
func (v DeleteCharactersCharacterIdMailLabelsLabelIdUnprocessableEntity) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89630228EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeleteCharactersCharacterIdMailLabelsLabelIdUnprocessableEntity) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89630228EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeleteCharactersCharacterIdMailLabelsLabelIdUnprocessableEntity) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89630228DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeleteCharactersCharacterIdMailLabelsLabelIdUnprocessableEntity) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89630228DecodeGithubComAntihaxGoesiEsi1(l, v)
}
