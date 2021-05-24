package ce

type Document struct {
	Object

	VersionEnabled bool
	Version
}

func (d *Document) EnableVersioning() {
	d.VersionEnabled = true
	d.MajorVersion = 1
	d.MinorVersion = 0
}
