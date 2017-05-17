package cli

import (
	"os"
	"time"

	"github.com/dailymuse/git-fit/config"
	"github.com/dailymuse/git-fit/util"
)

// GCArgs TODO write comment
type GCArgs struct {
	Args             []string
	Schema           *config.Config
	ReadDir          dirFileReader
	RemoveFile       fileRemover
	LogError         errorLogger
	SHA1SumIsValid   util.SHA1sumValidatorArgs
	SHA1SumGenerator util.SHA1SumGenerator
}

type dirFileReader func(dirname string) ([]os.FileInfo, error)

type fileRemover func(name string) error

type errorLogger func(format string, args ...interface{})

// MockFileInfo TODO fillout
type MockFileInfo struct {
	name string
}

// NewMockFileInfo TODO fill out
func NewMockFileInfo(mockFileName string) os.FileInfo {
	return MockFileInfo{
		name: mockFileName,
	}
}

func (m MockFileInfo) Name() string {
	return m.name
}

func (m MockFileInfo) Size() int64 {
	return 0
}

func (m MockFileInfo) Mode() os.FileMode {
	return os.ModeDir
}

func (m MockFileInfo) ModTime() time.Time {
	return time.Time{}
}

func (m MockFileInfo) IsDir() bool {
	return false
}

func (m MockFileInfo) Sys() interface{} {
	return nil
}
