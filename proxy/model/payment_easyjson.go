// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

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

func easyjson377dcee4DecodeGithubComDinel13ThesisAcTestModel(in *jlexer.Lexer, out *VerifyPaymentRequest) {
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
		case "id_mahasiswa":
			out.IdMahasiswa = int(in.Int())
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
func easyjson377dcee4EncodeGithubComDinel13ThesisAcTestModel(out *jwriter.Writer, in VerifyPaymentRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id_mahasiswa\":"
		out.RawString(prefix[1:])
		out.Int(int(in.IdMahasiswa))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v VerifyPaymentRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson377dcee4EncodeGithubComDinel13ThesisAcTestModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v VerifyPaymentRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson377dcee4EncodeGithubComDinel13ThesisAcTestModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *VerifyPaymentRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson377dcee4DecodeGithubComDinel13ThesisAcTestModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *VerifyPaymentRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson377dcee4DecodeGithubComDinel13ThesisAcTestModel(l, v)
}
func easyjson377dcee4DecodeGithubComDinel13ThesisAcTestModel1(in *jlexer.Lexer, out *PaymentWraperResponse) {
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
		case "payment":
			(out.PaymentResponse).UnmarshalEasyJSON(in)
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
func easyjson377dcee4EncodeGithubComDinel13ThesisAcTestModel1(out *jwriter.Writer, in PaymentWraperResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"payment\":"
		out.RawString(prefix[1:])
		(in.PaymentResponse).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PaymentWraperResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson377dcee4EncodeGithubComDinel13ThesisAcTestModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PaymentWraperResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson377dcee4EncodeGithubComDinel13ThesisAcTestModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PaymentWraperResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson377dcee4DecodeGithubComDinel13ThesisAcTestModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PaymentWraperResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson377dcee4DecodeGithubComDinel13ThesisAcTestModel1(l, v)
}
func easyjson377dcee4DecodeGithubComDinel13ThesisAcTestModel2(in *jlexer.Lexer, out *PaymentResponse) {
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
		case "isPay":
			out.IsPay = bool(in.Bool())
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
func easyjson377dcee4EncodeGithubComDinel13ThesisAcTestModel2(out *jwriter.Writer, in PaymentResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"isPay\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.IsPay))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PaymentResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson377dcee4EncodeGithubComDinel13ThesisAcTestModel2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PaymentResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson377dcee4EncodeGithubComDinel13ThesisAcTestModel2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PaymentResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson377dcee4DecodeGithubComDinel13ThesisAcTestModel2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PaymentResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson377dcee4DecodeGithubComDinel13ThesisAcTestModel2(l, v)
}
func easyjson377dcee4DecodeGithubComDinel13ThesisAcTestModel3(in *jlexer.Lexer, out *PaymentRequest) {
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
		case "id_mahasiswa":
			out.IdMahasiswa = int(in.Int())
		case "jumlah":
			out.Jumlah = float64(in.Float64())
		case "metode":
			out.Metode = string(in.String())
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
func easyjson377dcee4EncodeGithubComDinel13ThesisAcTestModel3(out *jwriter.Writer, in PaymentRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id_mahasiswa\":"
		out.RawString(prefix[1:])
		out.Int(int(in.IdMahasiswa))
	}
	{
		const prefix string = ",\"jumlah\":"
		out.RawString(prefix)
		out.Float64(float64(in.Jumlah))
	}
	{
		const prefix string = ",\"metode\":"
		out.RawString(prefix)
		out.String(string(in.Metode))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PaymentRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson377dcee4EncodeGithubComDinel13ThesisAcTestModel3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PaymentRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson377dcee4EncodeGithubComDinel13ThesisAcTestModel3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PaymentRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson377dcee4DecodeGithubComDinel13ThesisAcTestModel3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PaymentRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson377dcee4DecodeGithubComDinel13ThesisAcTestModel3(l, v)
}