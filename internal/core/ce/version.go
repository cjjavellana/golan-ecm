package ce

type Version interface {
	GetMajorVersion() uint32
	GetMinorVersion() uint32

	IncrementMajorVersion()
	IncrementMinorVersion()
	SetVersion(majorVersion uint32, minorVersion uint32) error
}
