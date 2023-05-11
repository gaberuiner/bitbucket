package storage

type Storage interface {
	Save(f *File) error
	PickRandom(userName string) (*File, error)
	IsExist(f *File) (bool, error)
	Delete(f *File) error
}

type File struct {
	UserSended string
	Type       string
	FileName   string
}
