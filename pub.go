package pop3

const (
	CRLF = "\r\n"

	USER     = "USER"
	PASSWORD = "PASS"
	NOOP     = "NOOP"
	RESET    = "RSET"
	DELETE   = "DELE"
	QUIT     = "QUIT"
	STATUS   = "STAT"
	LIST     = "LIST"
	RETRIEVE = "RETR"
)

type Pop3Error struct {
	prefix string
	msg    string
}

func (pe Pop3Error) Error() string {
	return "Pop3Error: " + pe.prefix + ": " + pe.msg
}
