package aws

import (
	"errors"
	"strconv"
)

type Version struct {
	majorVersion uint32
	minorVersion uint32
}

func NewVersion(majorVersion uint32, minorVersion uint32) *Version {
	return &Version{majorVersion: majorVersion, minorVersion: minorVersion}
}

func (v *Version) GetMajorVersion() uint32 {
	return v.majorVersion
}

func (v *Version) GetMinorVersion() uint32 {
	return v.minorVersion
}

func (v *Version) IncrementMajorVersion() {
	v.majorVersion = v.minorVersion + 1
	v.minorVersion = 0
}

func (v *Version) IncrementMinorVersion() {
	v.minorVersion = v.minorVersion + 1
}

func (v *Version) SetVersion(majorVersion uint32, minorVersion uint32) error {
	if majorVersion < v.majorVersion {
		return errors.New("majorVersion cannot be less than " + strconv.Itoa(int(v.majorVersion)))
	}

	if majorVersion == v.majorVersion && minorVersion <= v.minorVersion {
		return errors.New("minor version cannot be less than or equal to " + strconv.Itoa(int(v.minorVersion)) +
			"when major version remains the same")
	}

	v.majorVersion = majorVersion
	v.minorVersion = minorVersion

	return nil
}
