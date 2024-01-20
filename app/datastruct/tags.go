package datastruct

type UpdateTagCommand bool
type UpdateTagMap map[string]UpdateTagCommand

const (
	ADD_TAG    UpdateTagCommand = true
	REMOVE_TAG UpdateTagCommand = false
)

type Tag struct {
	Id   string
	Name string
}
