package ce

type Document interface {
	EnableVersioning()
	VersionEnabled() bool

	Object
	Version
}
