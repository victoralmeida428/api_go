package model

type Livro struct {
	Nome string
	Ano  int
}

func (m *Model) getLivro() *Livro {
	return &Livro{}
}
