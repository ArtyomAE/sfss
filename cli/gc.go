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
	fileHashesDeclaredInSchema := make(map[string]bool, len(schema.Files)*2)

	for _, hash := range schema.Files {
		fileHashesDeclaredInSchema[hash] = true
	}

	cacheFiles, err := ioutil.ReadDir(gitFitCacheDir)

	if err != nil {
		util.Fatal("Could not read %s: %s\n", gitFitCacheDir, err.Error())
	}

	for _, cacheFile := range cacheFiles {
		cacheFileName := cacheFile.Name()

		if err := util.SHA1sumIsValidForCacheFile(util.SHA1sumValidatorArgs{
			ReadDir:                 gitFitCacheDir,
			FileName:                cacheFileName,
			GenerateSHA1Sum:         util.FileHash,
			CacheFileHashesInSchema: fileHashesDeclaredInSchema,
		}); err != nil {
			util.Error("%s", err.Error())
			path := fmt.Sprintf("%s/%s", gitFitCacheDir, cacheFile.Name())
			err = os.Remove(path)

			if err != nil {
				util.Error("Could not delete cach file %s: %s\n", path, err.Error())
			}
		}
	}
}
