package facade

type BookFacade struct{}

func NewBookFacade() *BookFacade {
	return &BookFacade{}
}

func (f *BookFacade) AddBook()       {}
func (f *BookFacade) GetBook()       {}
func (f *BookFacade) RemoveBook()    {}
func (f *BookFacade) UpdateBook()    {}
func (f *BookFacade) SearchBooks()   {}
func (f *BookFacade) GetStoreStats() {}
