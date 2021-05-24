package ce

import "errors"

type Version struct {
	MajorVersion int32
	MinorVersion int32
}

func (v *Version) IncrementMajorVersion() {
	v.MajorVersion = v.MinorVersion + 1
	v.MinorVersion = 0
}

func (v *Version) IncrementMinorVersion() {
	v.MinorVersion = v.MinorVersion + 1
}

func (v *Version) SetVersion(majorVersion int32, minorVersion int32) error {
	if majorVersion < v.MajorVersion {
		return errors.New("majorVersion cannot be less than " + string(v.MajorVersion))
	}

	if majorVersion == v.MajorVersion && minorVersion <= v.MinorVersion {
		return errors.New("minor version cannot be less than or equal to " + string(v.MinorVersion) +
			"when major version remains the same")
	}

	v.MajorVersion = majorVersion
	v.MinorVersion = minorVersion

	return nil
}
