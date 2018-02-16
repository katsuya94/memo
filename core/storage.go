package core

type Storage interface {
	Retrieve()
	Store()
}
