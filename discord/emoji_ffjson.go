// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: discord/emoji.go

package discord

import (
	"bytes"
	"errors"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *Emoji) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *Emoji) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "id":`)

	{

		obj, err = j.ID.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"name":`)
	fflib.WriteJsonString(buf, string(j.Name))
	buf.WriteByte(',')
	if len(j.RoleIDs) != 0 {
		buf.WriteString(`"roles":`)
		if j.RoleIDs != nil {
			buf.WriteString(`[`)
			for i, v := range j.RoleIDs {
				if i != 0 {
					buf.WriteString(`,`)
				}

				{

					obj, err = v.MarshalJSON()
					if err != nil {
						return err
					}
					buf.Write(obj)

				}
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if true {
		buf.WriteString(`"user":`)

		{

			err = j.User.MarshalJSONBuf(buf)
			if err != nil {
				return err
			}

		}
		buf.WriteByte(',')
	}
	if j.RequireColons != false {
		if j.RequireColons {
			buf.WriteString(`"require_colons":true`)
		} else {
			buf.WriteString(`"require_colons":false`)
		}
		buf.WriteByte(',')
	}
	if j.Managed != false {
		if j.Managed {
			buf.WriteString(`"managed":true`)
		} else {
			buf.WriteString(`"managed":false`)
		}
		buf.WriteByte(',')
	}
	if j.Animated != false {
		if j.Animated {
			buf.WriteString(`"animated":true`)
		} else {
			buf.WriteString(`"animated":false`)
		}
		buf.WriteByte(',')
	}
	if j.Available != false {
		if j.Available {
			buf.WriteString(`"available":true`)
		} else {
			buf.WriteString(`"available":false`)
		}
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtEmojibase = iota
	ffjtEmojinosuchkey

	ffjtEmojiID

	ffjtEmojiName

	ffjtEmojiRoleIDs

	ffjtEmojiUser

	ffjtEmojiRequireColons

	ffjtEmojiManaged

	ffjtEmojiAnimated

	ffjtEmojiAvailable
)

var ffjKeyEmojiID = []byte("id")

var ffjKeyEmojiName = []byte("name")

var ffjKeyEmojiRoleIDs = []byte("roles")

var ffjKeyEmojiUser = []byte("user")

var ffjKeyEmojiRequireColons = []byte("require_colons")

var ffjKeyEmojiManaged = []byte("managed")

var ffjKeyEmojiAnimated = []byte("animated")

var ffjKeyEmojiAvailable = []byte("available")

// UnmarshalJSON umarshall json - template of ffjson
func (j *Emoji) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *Emoji) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtEmojibase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtEmojinosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyEmojiAnimated, kn) {
						currentKey = ffjtEmojiAnimated
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyEmojiAvailable, kn) {
						currentKey = ffjtEmojiAvailable
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyEmojiID, kn) {
						currentKey = ffjtEmojiID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffjKeyEmojiManaged, kn) {
						currentKey = ffjtEmojiManaged
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffjKeyEmojiName, kn) {
						currentKey = ffjtEmojiName
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffjKeyEmojiRoleIDs, kn) {
						currentKey = ffjtEmojiRoleIDs
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyEmojiRequireColons, kn) {
						currentKey = ffjtEmojiRequireColons
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'u':

					if bytes.Equal(ffjKeyEmojiUser, kn) {
						currentKey = ffjtEmojiUser
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyEmojiAvailable, kn) {
					currentKey = ffjtEmojiAvailable
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyEmojiAnimated, kn) {
					currentKey = ffjtEmojiAnimated
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyEmojiManaged, kn) {
					currentKey = ffjtEmojiManaged
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyEmojiRequireColons, kn) {
					currentKey = ffjtEmojiRequireColons
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyEmojiUser, kn) {
					currentKey = ffjtEmojiUser
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyEmojiRoleIDs, kn) {
					currentKey = ffjtEmojiRoleIDs
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyEmojiName, kn) {
					currentKey = ffjtEmojiName
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyEmojiID, kn) {
					currentKey = ffjtEmojiID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtEmojinosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtEmojiID:
					goto handle_ID

				case ffjtEmojiName:
					goto handle_Name

				case ffjtEmojiRoleIDs:
					goto handle_RoleIDs

				case ffjtEmojiUser:
					goto handle_User

				case ffjtEmojiRequireColons:
					goto handle_RequireColons

				case ffjtEmojiManaged:
					goto handle_Managed

				case ffjtEmojiAnimated:
					goto handle_Animated

				case ffjtEmojiAvailable:
					goto handle_Available

				case ffjtEmojinosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_ID:

	/* handler: j.ID type=discord.EmojiID kind=uint64 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.ID.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Name:

	/* handler: j.Name type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Name = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_RoleIDs:

	/* handler: j.RoleIDs type=[]discord.RoleID kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.RoleIDs = nil
		} else {

			j.RoleIDs = []RoleID{}

			wantVal := true

			for {

				var tmpJRoleIDs RoleID

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJRoleIDs type=discord.RoleID kind=uint64 quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						tbuf, err := fs.CaptureField(tok)
						if err != nil {
							return fs.WrapErr(err)
						}

						err = tmpJRoleIDs.UnmarshalJSON(tbuf)
						if err != nil {
							return fs.WrapErr(err)
						}
					}
					state = fflib.FFParse_after_value
				}

				j.RoleIDs = append(j.RoleIDs, tmpJRoleIDs)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_User:

	/* handler: j.User type=discord.User kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			err = j.User.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_RequireColons:

	/* handler: j.RequireColons type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				j.RequireColons = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				j.RequireColons = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Managed:

	/* handler: j.Managed type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				j.Managed = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				j.Managed = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Animated:

	/* handler: j.Animated type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				j.Animated = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				j.Animated = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Available:

	/* handler: j.Available type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				j.Available = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				j.Available = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}