package storage

type Storage interface {
	Save(f *Film) error
	PickRandom(userName string) (*Film, error)
	IsExist(f *Film) (bool, error)
	Delete(f *Film) error
}

type Film struct {
	UserSended string
	FilmName   string
}
