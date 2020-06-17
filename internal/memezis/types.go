package memezis

// sources
const (
	SourceMemezisBot = "memezis_bot"
	SourcePostman    = "postman"
)

type clientError interface {
	Error() string
	Code() int
}
