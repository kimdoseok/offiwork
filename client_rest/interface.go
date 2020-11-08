package main

type (
	Reader interface {
		Get(id uint) Client
		List(page uint) []Client
		Search(fltstr string) []Client
	}

	Writer interface {
		Save(Client) bool
		Delete(id uint) bool
	}

	Repository interface {
		Reader
		Writer
	}
)
