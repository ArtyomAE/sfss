package util

import (
	"os"
	"time"
)

// MockFileInfo mocks the os.FileInfo interface.
type MockFileInfo struct {
	name string
}

// NewMockFileInfo returns a new instance of MockFileInfo.
func NewMockFileInfo(mockFileName string) os.FileInfo {
	return MockFileInfo{
		name: mockFileName,
	}
}

// Name returns the mock file name.
func (m MockFileInfo) Name() string {
	return m.name
}

// Size returns a zero value for size.
func (m MockFileInfo) Size() int64 {
	return 0
}

// Mode returns an empty os.ModeDir for the file's mode bits.
func (m MockFileInfo) Mode() os.FileMode {
	return os.ModeDir
}

// ModTime returns an empty time.Time for the last modification time.
func (m MockFileInfo) ModTime() time.Time {
	return time.Time{}
}

// IsDir returns a zero value for if fileInfo is a directory.
func (m MockFileInfo) IsDir() bool {
	return false
}

// Sys returns nil for the underlying file data source.
func (m MockFileInfo) Sys() interface{} {
	return nil
}
