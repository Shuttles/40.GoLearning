package books

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) CategoryByLenth() string {
	if b.Pages >= 300 {
		return "NOVEL"
	}

	return "SHORT STORY"
}

func (b *Book) GetAuthor() string {
	return b.Author
}
