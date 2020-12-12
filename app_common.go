package app

var (
	body       *elem
	content    UI
	renderFunc RenderFunc
	helmet     = &Helmet{}
)

type Helmet struct {
	Title       string
	Description string
	Keywords    string
	Author      string
}
