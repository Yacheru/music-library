package constants

import "errors"

var (
	EmptyConfigVarError    = errors.New("empty config variable: ")
	NoLyricsFoundError     = errors.New("no lyrics found")
	TimeOutError           = errors.New("request to get lyrics took too long. Maybe lyrics doesn't exists...")
	SongNotFoundError      = errors.New("song not found")
	SongAlreadyExistsError = errors.New(`song already exists`)
)
