package cli

import (
	"os"

	"github.com/dailymuse/git-fit/config"
	"github.com/dailymuse/git-fit/util"
)

// GcArgs is the arguments struct for Gc.
type GcArgs struct {
	Args             []string
	Schema           *config.Config
	ReadDir          dirFileReader
	RemoveFile       fileRemover
	LogError         errorLogger
	SHA1SumIsValid   util.SHA1sumValidatorArgs
	SHA1SumGenerator util.SHA1SumGenerator
}

// dirFileReader is responsible for reading a directory and return an array
// of FileInfo.
type dirFileReader func(dirname string) ([]os.FileInfo, error)

// fileRemover is responsible for removing a file from the file system.
type fileRemover func(name string) error

// errorLogger is responsible for logging errors.
type errorLogger func(format string, args ...interface{})
