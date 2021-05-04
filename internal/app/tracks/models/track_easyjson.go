// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	models2 "2021_1_Noskool_team/internal/app/album/models"
	models1 "2021_1_Noskool_team/internal/app/musicians/models"
	models "2021_1_Noskool_team/internal/models"
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

func easyjsonAe118d8fDecode20211NoskoolTeamInternalAppTracksModels(in *jlexer.Lexer, out *Tracks) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Tracks, 0, 8)
			} else {
				*out = Tracks{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 *Track
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(Track)
				}
				(*v1).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonAe118d8fEncode20211NoskoolTeamInternalAppTracksModels(out *jwriter.Writer, in Tracks) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			if v3 == nil {
				out.RawString("null")
			} else {
				(*v3).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Tracks) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAe118d8fEncode20211NoskoolTeamInternalAppTracksModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Tracks) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAe118d8fEncode20211NoskoolTeamInternalAppTracksModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Tracks) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAe118d8fDecode20211NoskoolTeamInternalAppTracksModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Tracks) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAe118d8fDecode20211NoskoolTeamInternalAppTracksModels(l, v)
}
func easyjsonAe118d8fDecode20211NoskoolTeamInternalAppTracksModels1(in *jlexer.Lexer, out *Track) {
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
		case "track_id":
			out.TrackID = int(in.Int())
		case "tittle":
			out.Tittle = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "audio":
			out.Audio = string(in.String())
		case "picture":
			out.Picture = string(in.String())
		case "release_date":
			out.ReleaseDate = string(in.String())
		case "duration":
			out.Duration = string(in.String())
		case "in_mediateka":
			out.InMediateka = bool(in.Bool())
		case "in_favorite":
			out.InFavorite = bool(in.Bool())
		case "genres":
			if in.IsNull() {
				in.Skip()
				out.Genres = nil
			} else {
				in.Delim('[')
				if out.Genres == nil {
					if !in.IsDelim(']') {
						out.Genres = make([]*models.Genre, 0, 8)
					} else {
						out.Genres = []*models.Genre{}
					}
				} else {
					out.Genres = (out.Genres)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *models.Genre
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(models.Genre)
						}
						easyjsonAe118d8fDecode20211NoskoolTeamInternalModels(in, v4)
					}
					out.Genres = append(out.Genres, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "musicians":
			if in.IsNull() {
				in.Skip()
				out.Musicians = nil
			} else {
				in.Delim('[')
				if out.Musicians == nil {
					if !in.IsDelim(']') {
						out.Musicians = make([]*models1.Musician, 0, 8)
					} else {
						out.Musicians = []*models1.Musician{}
					}
				} else {
					out.Musicians = (out.Musicians)[:0]
				}
				for !in.IsDelim(']') {
					var v5 *models1.Musician
					if in.IsNull() {
						in.Skip()
						v5 = nil
					} else {
						if v5 == nil {
							v5 = new(models1.Musician)
						}
						(*v5).UnmarshalEasyJSON(in)
					}
					out.Musicians = append(out.Musicians, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "album":
			if in.IsNull() {
				in.Skip()
				out.Albums = nil
			} else {
				in.Delim('[')
				if out.Albums == nil {
					if !in.IsDelim(']') {
						out.Albums = make([]*models2.Album, 0, 8)
					} else {
						out.Albums = []*models2.Album{}
					}
				} else {
					out.Albums = (out.Albums)[:0]
				}
				for !in.IsDelim(']') {
					var v6 *models2.Album
					if in.IsNull() {
						in.Skip()
						v6 = nil
					} else {
						if v6 == nil {
							v6 = new(models2.Album)
						}
						(*v6).UnmarshalEasyJSON(in)
					}
					out.Albums = append(out.Albums, v6)
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
func easyjsonAe118d8fEncode20211NoskoolTeamInternalAppTracksModels1(out *jwriter.Writer, in Track) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"track_id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.TrackID))
	}
	{
		const prefix string = ",\"tittle\":"
		out.RawString(prefix)
		out.String(string(in.Tittle))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"audio\":"
		out.RawString(prefix)
		out.String(string(in.Audio))
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
		const prefix string = ",\"duration\":"
		out.RawString(prefix)
		out.String(string(in.Duration))
	}
	{
		const prefix string = ",\"in_mediateka\":"
		out.RawString(prefix)
		out.Bool(bool(in.InMediateka))
	}
	{
		const prefix string = ",\"in_favorite\":"
		out.RawString(prefix)
		out.Bool(bool(in.InFavorite))
	}
	{
		const prefix string = ",\"genres\":"
		out.RawString(prefix)
		if in.Genres == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v7, v8 := range in.Genres {
				if v7 > 0 {
					out.RawByte(',')
				}
				if v8 == nil {
					out.RawString("null")
				} else {
					easyjsonAe118d8fEncode20211NoskoolTeamInternalModels(out, *v8)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"musicians\":"
		out.RawString(prefix)
		if in.Musicians == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.Musicians {
				if v9 > 0 {
					out.RawByte(',')
				}
				if v10 == nil {
					out.RawString("null")
				} else {
					(*v10).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"album\":"
		out.RawString(prefix)
		if in.Albums == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Albums {
				if v11 > 0 {
					out.RawByte(',')
				}
				if v12 == nil {
					out.RawString("null")
				} else {
					(*v12).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Track) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAe118d8fEncode20211NoskoolTeamInternalAppTracksModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Track) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAe118d8fEncode20211NoskoolTeamInternalAppTracksModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Track) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAe118d8fDecode20211NoskoolTeamInternalAppTracksModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Track) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAe118d8fDecode20211NoskoolTeamInternalAppTracksModels1(l, v)
}
func easyjsonAe118d8fDecode20211NoskoolTeamInternalModels(in *jlexer.Lexer, out *models.Genre) {
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
		case "genre_id":
			out.GenreID = int(in.Int())
		case "title":
			out.Title = string(in.String())
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
func easyjsonAe118d8fEncode20211NoskoolTeamInternalModels(out *jwriter.Writer, in models.Genre) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"genre_id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.GenreID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	out.RawByte('}')
}
