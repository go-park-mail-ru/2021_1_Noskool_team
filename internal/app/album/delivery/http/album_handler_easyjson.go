// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package http

import (
	models "2021_1_Noskool_team/internal/app/musicians/models"
	models1 "2021_1_Noskool_team/internal/app/tracks/models"
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

func easyjson345a258eDecode20211NoskoolTeamInternalAppAlbumDeliveryHttp(in *jlexer.Lexer, out *AlbumsHandler) {
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
func easyjson345a258eEncode20211NoskoolTeamInternalAppAlbumDeliveryHttp(out *jwriter.Writer, in AlbumsHandler) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (handler AlbumsHandler) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson345a258eEncode20211NoskoolTeamInternalAppAlbumDeliveryHttp(&w, handler)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (handler AlbumsHandler) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson345a258eEncode20211NoskoolTeamInternalAppAlbumDeliveryHttp(w, handler)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (handler *AlbumsHandler) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson345a258eDecode20211NoskoolTeamInternalAppAlbumDeliveryHttp(&r, handler)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (handler *AlbumsHandler) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson345a258eDecode20211NoskoolTeamInternalAppAlbumDeliveryHttp(l, handler)
}
func easyjson345a258eDecode20211NoskoolTeamInternalAppAlbumDeliveryHttp1(in *jlexer.Lexer, out *AlbumWithExtraInform) {
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
		case "album_id":
			out.AlbumID = int(in.Int())
		case "tittle":
			out.Tittle = string(in.String())
		case "picture":
			out.Picture = string(in.String())
		case "release_date":
			out.ReleaseDate = string(in.String())
		case "musician":
			if in.IsNull() {
				in.Skip()
				out.Musician = nil
			} else {
				if out.Musician == nil {
					out.Musician = new([]models.Musician)
				}
				if in.IsNull() {
					in.Skip()
					*out.Musician = nil
				} else {
					in.Delim('[')
					if *out.Musician == nil {
						if !in.IsDelim(']') {
							*out.Musician = make([]models.Musician, 0, 1)
						} else {
							*out.Musician = []models.Musician{}
						}
					} else {
						*out.Musician = (*out.Musician)[:0]
					}
					for !in.IsDelim(']') {
						var v1 models.Musician
						(v1).UnmarshalEasyJSON(in)
						*out.Musician = append(*out.Musician, v1)
						in.WantComma()
					}
					in.Delim(']')
				}
			}
		case "tracks":
			if in.IsNull() {
				in.Skip()
				out.Tracks = nil
			} else {
				in.Delim('[')
				if out.Tracks == nil {
					if !in.IsDelim(']') {
						out.Tracks = make([]*models1.Track, 0, 8)
					} else {
						out.Tracks = []*models1.Track{}
					}
				} else {
					out.Tracks = (out.Tracks)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *models1.Track
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(models1.Track)
						}
						(*v2).UnmarshalEasyJSON(in)
					}
					out.Tracks = append(out.Tracks, v2)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson345a258eEncode20211NoskoolTeamInternalAppAlbumDeliveryHttp1(out *jwriter.Writer, in AlbumWithExtraInform) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"album_id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.AlbumID))
	}
	{
		const prefix string = ",\"tittle\":"
		out.RawString(prefix)
		out.String(string(in.Tittle))
	}
	{
		const prefix string = ",\"picture\":"
		out.RawString(prefix)
		out.String(string(in.Picture))
	}
	{
		const prefix string = ",\"release_date\":"
		out.RawString(prefix)
		out.String(string(in.ReleaseDate))
	}
	{
		const prefix string = ",\"musician\":"
		out.RawString(prefix)
		if in.Musician == nil {
			out.RawString("null")
		} else {
			if *in.Musician == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
				out.RawString("null")
			} else {
				out.RawByte('[')
				for v3, v4 := range *in.Musician {
					if v3 > 0 {
						out.RawByte(',')
					}
					(v4).MarshalEasyJSON(out)
				}
				out.RawByte(']')
			}
		}
	}
	{
		const prefix string = ",\"tracks\":"
		out.RawString(prefix)
		if in.Tracks == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Tracks {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AlbumWithExtraInform) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson345a258eEncode20211NoskoolTeamInternalAppAlbumDeliveryHttp1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AlbumWithExtraInform) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson345a258eEncode20211NoskoolTeamInternalAppAlbumDeliveryHttp1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AlbumWithExtraInform) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson345a258eDecode20211NoskoolTeamInternalAppAlbumDeliveryHttp1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AlbumWithExtraInform) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson345a258eDecode20211NoskoolTeamInternalAppAlbumDeliveryHttp1(l, v)
}
