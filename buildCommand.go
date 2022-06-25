package spacetrack

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"spacetrack/utils"
	"strings"
)

type BuildCommand struct {
	Cookies    []*http.Cookie
	Paths      []string
	Method     string
	StreamData bool
}

func (b *BuildCommand) Predicates(column string) *BuildCommand {
	column = utils.PathCleanSep(column)
	b.Paths = append(b.Paths, column)
	return b
}

func (b *BuildCommand) GT(value string) *BuildCommand {
	value = utils.PathCleanSep(value)
	b.Paths = append(b.Paths, url.PathEscape(RESTOperatorsGT)+value)
	return b
}

func (b *BuildCommand) LT(value string) *BuildCommand {
	value = utils.PathCleanSep(value)
	b.Paths = append(b.Paths, url.PathEscape(RESTOperatorsLT)+value)
	return b
}

func (b *BuildCommand) NE(value string) *BuildCommand {
	value = utils.PathCleanSep(value)
	b.Paths = append(b.Paths, url.PathEscape(RESTOperatorsNE)+value)
	return b
}

func (b *BuildCommand) OR(value ...string) *BuildCommand {
	b.Paths = append(b.Paths, strings.Join(value, RESTOperatorsOR))
	return b
}

func (b *BuildCommand) Range(start, end string) *BuildCommand {
	start = utils.PathCleanSep(start)
	end = utils.PathCleanSep(end)
	b.Paths = append(b.Paths, start+RESTOperatorsRange+end)
	return b
}

func (b *BuildCommand) Like(value string) *BuildCommand {
	value = utils.PathCleanSep(value)
	b.Paths = append(b.Paths, RESTOperatorsLike+value)
	return b
}

func (b *BuildCommand) Wildcard(value string) *BuildCommand {
	value = utils.PathCleanSep(value)
	b.Paths = append(b.Paths, RESTOperatorsWildcard+value)
	return b
}

func (b *BuildCommand) NullValue() *BuildCommand {
	b.Paths = append(b.Paths, RESTOperatorsNullValue)
	return b
}

func (b *BuildCommand) MetadataShow() *BuildCommand {
	b.Paths = append(b.Paths, RESTPredicatesMetaData, "true")
	return b
}

func (b *BuildCommand) Distinct() *BuildCommand {
	b.Paths = append(b.Paths, RESTPredicatesDistinct, "true")
	return b
}

func (b *BuildCommand) OrderBy(value string, sort ...string) *BuildCommand {
	value = utils.PathCleanSep(value)

	if value == "" {
		return b
	}

	sb := new(strings.Builder)
	sb.WriteString(value)

	if len(sort) > 0 {
		sb.WriteString(" " + sort[0])
	}

	b.Paths = append(b.Paths, RESTPredicatesOrderBy, url.PathEscape(sb.String()))
	return b
}

func (b *BuildCommand) Limit(size int, offset ...int) *BuildCommand {

	sb := new(strings.Builder)
	sb.WriteString(utils.ToStr(size))

	if len(offset) > 0 {
		sb.WriteString(",")
		sb.WriteString(utils.ToStr(offset[0]))
	}

	b.Paths = append(b.Paths, RESTPredicatesLimit, sb.String())
	return b
}

func (b *BuildCommand) EmptyResultShow() *BuildCommand {
	b.Paths = append(b.Paths, RESTPredicatesEmptyResultShow, "show")
	return b
}

func (b *BuildCommand) FormatJSON() *BuildCommand {
	b.Paths = append(b.Paths, RESTPredicatesFormat, "json")
	return b
}

func (b *BuildCommand) FormatXML() *BuildCommand {
	b.Paths = append(b.Paths, RESTPredicatesFormat, "xml")
	return b
}

func (b *BuildCommand) FormatCSV() *BuildCommand {
	b.Paths = append(b.Paths, RESTPredicatesFormat, "csv")
	return b
}

func (b *BuildCommand) FormatTLE() *BuildCommand {
	b.Paths = append(b.Paths, RESTPredicatesFormat, "tle")
	return b
}

func (b *BuildCommand) FormatHTML() *BuildCommand {
	b.Paths = append(b.Paths, RESTPredicatesFormat, "html")
	return b
}

func (b *BuildCommand) Format3LE() *BuildCommand {
	b.Paths = append(b.Paths, RESTPredicatesFormat, "3le")
	return b
}

func (b *BuildCommand) FormatKVN() *BuildCommand {
	b.Paths = append(b.Paths, RESTPredicatesFormat, "kvn")
	return b
}

func (b *BuildCommand) FormatStream() *BuildCommand {
	b.Paths = append(b.Paths, RESTPredicatesFormat, "stream")
	b.StreamData = true
	return b
}

type BuildCodec interface {
	Unmarshal(data []byte, v any) error
	Marshal(v any) ([]byte, error)
	Write(p []byte) (n int, err error)
}

func (b *BuildCommand) Do(entity any, codec ...BuildCodec) error {

	reqUrl := strings.Join(b.Paths, utils.URLSeparator)

	resp, err := utils.NewHttpClient(
		b.Method,
		reqUrl,
		nil,
		&utils.HttpClientOpts{
			DisableCompression: false,
			DisableKeepAlives:  false,
			Cookies:            b.Cookies,
		},
	)

	if err != nil {
		return err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	switch resp.StatusCode {
	case http.StatusOK:
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if len(codec) > 0 {
			return codec[0].Unmarshal(body, &entity)
		} else {
			return json.Unmarshal(body, &entity)
		}
	case http.StatusNoContent:
		if len(codec) > 0 {
			_, err := io.Copy(codec[0], resp.Body)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf(
				utils.HttpResultFormat,
				600,
				"Unavailable BuildCodec!!!",
			)
		}
		return nil
	case http.StatusForbidden:
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf(
			utils.HttpResultFormat,
			resp.StatusCode,
			string(body),
		)
	default:
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf(
			utils.HttpResultFormat,
			resp.StatusCode,
			string(body),
		)
	}
}
