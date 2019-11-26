package constant

type LOG_LEVEL int
type SOURCE int

const (
	LOG_LEVEL_ERROR LOG_LEVEL = 1 + iota
	LOG_LEVEL_INFO
	LOG_LEVEL_DEBUG
)

const (
	SOURCE_SFTP SOURCE = 1 + iota
	SOURCE_ESB
)
