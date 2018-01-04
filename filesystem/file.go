package filesystem

type File interface {
	Create() error
	Path() string
	Exists() bool
}
