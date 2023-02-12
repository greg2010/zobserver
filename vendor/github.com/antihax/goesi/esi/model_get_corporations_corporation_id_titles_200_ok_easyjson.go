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

func easyjsonDf6cf6fdDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCorporationsCorporationIdTitles200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationsCorporationIdTitles200OkList, 0, 0)
			} else {
				*out = GetCorporationsCorporationIdTitles200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationsCorporationIdTitles200Ok
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
func easyjsonDf6cf6fdEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCorporationsCorporationIdTitles200OkList) {
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
func (v GetCorporationsCorporationIdTitles200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDf6cf6fdEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdTitles200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDf6cf6fdEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdTitles200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDf6cf6fdDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdTitles200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDf6cf6fdDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonDf6cf6fdDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCorporationsCorporationIdTitles200Ok) {
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
		case "grantable_roles":
			if in.IsNull() {
				in.Skip()
				out.GrantableRoles = nil
			} else {
				in.Delim('[')
				if out.GrantableRoles == nil {
					if !in.IsDelim(']') {
						out.GrantableRoles = make([]string, 0, 4)
					} else {
						out.GrantableRoles = []string{}
					}
				} else {
					out.GrantableRoles = (out.GrantableRoles)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.GrantableRoles = append(out.GrantableRoles, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "grantable_roles_at_base":
			if in.IsNull() {
				in.Skip()
				out.GrantableRolesAtBase = nil
			} else {
				in.Delim('[')
				if out.GrantableRolesAtBase == nil {
					if !in.IsDelim(']') {
						out.GrantableRolesAtBase = make([]string, 0, 4)
					} else {
						out.GrantableRolesAtBase = []string{}
					}
				} else {
					out.GrantableRolesAtBase = (out.GrantableRolesAtBase)[:0]
				}
				for !in.IsDelim(']') {
					var v5 string
					v5 = string(in.String())
					out.GrantableRolesAtBase = append(out.GrantableRolesAtBase, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "grantable_roles_at_hq":
			if in.IsNull() {
				in.Skip()
				out.GrantableRolesAtHq = nil
			} else {
				in.Delim('[')
				if out.GrantableRolesAtHq == nil {
					if !in.IsDelim(']') {
						out.GrantableRolesAtHq = make([]string, 0, 4)
					} else {
						out.GrantableRolesAtHq = []string{}
					}
				} else {
					out.GrantableRolesAtHq = (out.GrantableRolesAtHq)[:0]
				}
				for !in.IsDelim(']') {
					var v6 string
					v6 = string(in.String())
					out.GrantableRolesAtHq = append(out.GrantableRolesAtHq, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "grantable_roles_at_other":
			if in.IsNull() {
				in.Skip()
				out.GrantableRolesAtOther = nil
			} else {
				in.Delim('[')
				if out.GrantableRolesAtOther == nil {
					if !in.IsDelim(']') {
						out.GrantableRolesAtOther = make([]string, 0, 4)
					} else {
						out.GrantableRolesAtOther = []string{}
					}
				} else {
					out.GrantableRolesAtOther = (out.GrantableRolesAtOther)[:0]
				}
				for !in.IsDelim(']') {
					var v7 string
					v7 = string(in.String())
					out.GrantableRolesAtOther = append(out.GrantableRolesAtOther, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "name":
			out.Name = string(in.String())
		case "roles":
			if in.IsNull() {
				in.Skip()
				out.Roles = nil
			} else {
				in.Delim('[')
				if out.Roles == nil {
					if !in.IsDelim(']') {
						out.Roles = make([]string, 0, 4)
					} else {
						out.Roles = []string{}
					}
				} else {
					out.Roles = (out.Roles)[:0]
				}
				for !in.IsDelim(']') {
					var v8 string
					v8 = string(in.String())
					out.Roles = append(out.Roles, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "roles_at_base":
			if in.IsNull() {
				in.Skip()
				out.RolesAtBase = nil
			} else {
				in.Delim('[')
				if out.RolesAtBase == nil {
					if !in.IsDelim(']') {
						out.RolesAtBase = make([]string, 0, 4)
					} else {
						out.RolesAtBase = []string{}
					}
				} else {
					out.RolesAtBase = (out.RolesAtBase)[:0]
				}
				for !in.IsDelim(']') {
					var v9 string
					v9 = string(in.String())
					out.RolesAtBase = append(out.RolesAtBase, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "roles_at_hq":
			if in.IsNull() {
				in.Skip()
				out.RolesAtHq = nil
			} else {
				in.Delim('[')
				if out.RolesAtHq == nil {
					if !in.IsDelim(']') {
						out.RolesAtHq = make([]string, 0, 4)
					} else {
						out.RolesAtHq = []string{}
					}
				} else {
					out.RolesAtHq = (out.RolesAtHq)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.RolesAtHq = append(out.RolesAtHq, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "roles_at_other":
			if in.IsNull() {
				in.Skip()
				out.RolesAtOther = nil
			} else {
				in.Delim('[')
				if out.RolesAtOther == nil {
					if !in.IsDelim(']') {
						out.RolesAtOther = make([]string, 0, 4)
					} else {
						out.RolesAtOther = []string{}
					}
				} else {
					out.RolesAtOther = (out.RolesAtOther)[:0]
				}
				for !in.IsDelim(']') {
					var v11 string
					v11 = string(in.String())
					out.RolesAtOther = append(out.RolesAtOther, v11)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "title_id":
			out.TitleId = int32(in.Int32())
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
func easyjsonDf6cf6fdEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCorporationsCorporationIdTitles200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.GrantableRoles) != 0 {
		const prefix string = ",\"grantable_roles\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v12, v13 := range in.GrantableRoles {
				if v12 > 0 {
					out.RawByte(',')
				}
				out.String(string(v13))
			}
			out.RawByte(']')
		}
	}
	if len(in.GrantableRolesAtBase) != 0 {
		const prefix string = ",\"grantable_roles_at_base\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v14, v15 := range in.GrantableRolesAtBase {
				if v14 > 0 {
					out.RawByte(',')
				}
				out.String(string(v15))
			}
			out.RawByte(']')
		}
	}
	if len(in.GrantableRolesAtHq) != 0 {
		const prefix string = ",\"grantable_roles_at_hq\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v16, v17 := range in.GrantableRolesAtHq {
				if v16 > 0 {
					out.RawByte(',')
				}
				out.String(string(v17))
			}
			out.RawByte(']')
		}
	}
	if len(in.GrantableRolesAtOther) != 0 {
		const prefix string = ",\"grantable_roles_at_other\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v18, v19 := range in.GrantableRolesAtOther {
				if v18 > 0 {
					out.RawByte(',')
				}
				out.String(string(v19))
			}
			out.RawByte(']')
		}
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if len(in.Roles) != 0 {
		const prefix string = ",\"roles\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v20, v21 := range in.Roles {
				if v20 > 0 {
					out.RawByte(',')
				}
				out.String(string(v21))
			}
			out.RawByte(']')
		}
	}
	if len(in.RolesAtBase) != 0 {
		const prefix string = ",\"roles_at_base\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v22, v23 := range in.RolesAtBase {
				if v22 > 0 {
					out.RawByte(',')
				}
				out.String(string(v23))
			}
			out.RawByte(']')
		}
	}
	if len(in.RolesAtHq) != 0 {
		const prefix string = ",\"roles_at_hq\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v24, v25 := range in.RolesAtHq {
				if v24 > 0 {
					out.RawByte(',')
				}
				out.String(string(v25))
			}
			out.RawByte(']')
		}
	}
	if len(in.RolesAtOther) != 0 {
		const prefix string = ",\"roles_at_other\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v26, v27 := range in.RolesAtOther {
				if v26 > 0 {
					out.RawByte(',')
				}
				out.String(string(v27))
			}
			out.RawByte(']')
		}
	}
	if in.TitleId != 0 {
		const prefix string = ",\"title_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.TitleId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationsCorporationIdTitles200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDf6cf6fdEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdTitles200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDf6cf6fdEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdTitles200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDf6cf6fdDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdTitles200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDf6cf6fdDecodeGithubComAntihaxGoesiEsi1(l, v)
}
