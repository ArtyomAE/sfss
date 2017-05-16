package cli

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dailymuse/git-fit/config"
	"github.com/dailymuse/git-fit/transport"
	"github.com/dailymuse/git-fit/util"
)

const gitFitCacheDir = ".git/fit"

func Gc(schema *config.Config, trans transport.Transport, args []string) {
	filesDeclaredInSchema := make(map[string]bool, len(schema.Files)*2)

	for _, hash := range schema.Files {
		filesDeclaredInSchema[hash] = true
	}

	allFiles, err := ioutil.ReadDir(gitFitCacheDir)

	if err != nil {
		util.Fatal("Could not read %s: %s\n", gitFitCacheDir, err.Error())
	}

	for _, file := range allFiles {
		fileNameHash := file.Name()

		if err := sha1sumIsValidForCacheFile(sha1sumValidatorArgs{
			readDir:                 gitFitCacheDir,
			fileName:                fileNameHash,
			generateSHA1Sum:         util.FileHash,
			cacheFileHashesInSchema: filesDeclaredInSchema,
		}); err != nil {
			util.Error("%s", err.Error())
			path := fmt.Sprintf("%s/%s", gitFitCacheDir, file.Name())
			err = os.Remove(path)

			if err != nil {
				util.Error("Could not delete cached file %s: %s\n", path, err.Error())
			}
		}
	}
}

type sha1sumValidatorArgs struct {
	fileName                string
	readDir                 string
	cacheFileHashesInSchema map[string]bool
	generateSHA1Sum         sha1SumGenerator
}

// fileReader is reponsible for opening files on FS.
type fileReader func(filename string) ([]byte, error)

// sha1SumGenerator is responsible for generating a sha1sum from a byte array.
type sha1SumGenerator func(p string) (string, error)

func sha1sumIsValidForCacheFile(args sha1sumValidatorArgs) error {
	// Check to see if the file name SHA we're validating exists in the schema.
	_, cacheObjExistsInSchema := args.cacheFileHashesInSchema[args.fileName]
	if !cacheObjExistsInSchema {
		return fmt.Errorf(
			"file hash '%s' does not exist in schema",
			args.fileName,
		)
	}

	// Attempt to generate the SHA1Sum for the cache file.
	filePath := fmt.Sprintf("%s/%s", args.readDir, args.fileName)
	fileSHA1Sum, err := args.generateSHA1Sum(filePath)
	if err != nil {
		return fmt.Errorf(
			"could not generate file hash: %s",
			err.Error(),
		)
	}

	// Check to make the correlating schema's hash equals the sha1sum of the cache file.
	if fileSHA1Sum != args.fileName {
		return fmt.Errorf(
			"cache file '%s' did not equal the sha1sum of the correlating cache blob",
			args.fileName,
		)
	}

	return nil
}
