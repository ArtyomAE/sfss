package cli

import (
	"fmt"

	"github.com/dailymuse/git-fit/util"
)

const gitFitCacheDir = ".git/fit"

// Gc is responsible for removing invalid cache objects not declared in the git-fit schema.json.
func Gc(args GcArgs) error {
	fileHashesDeclaredInSchema := make(map[string]bool, len(args.Schema.Files)*2)

	for _, hash := range args.Schema.Files {
		fileHashesDeclaredInSchema[hash] = true
	}

	cacheFiles, err := args.ReadDir(gitFitCacheDir)

	if err != nil {
		args.LogError("could not read %s: %s", gitFitCacheDir, err.Error())
		return err
	}

	for _, cacheFile := range cacheFiles {
		cacheFileName := cacheFile.Name()
		if err := util.SHA1sumIsValidForCacheFile(util.SHA1sumValidatorArgs{
			ReadDir:                 gitFitCacheDir,
			FileName:                cacheFileName,
			GenerateSHA1Sum:         args.SHA1SumGenerator,
			CacheFileHashesInSchema: fileHashesDeclaredInSchema,
		}); err != nil {
			args.LogError("%s", err.Error())
			path := fmt.Sprintf("%s/%s", gitFitCacheDir, cacheFile.Name())
			err = args.RemoveFile(path)

			if err != nil {
				args.LogError("could not delete cach file %s: %s", path, err.Error())
				return err
			}
		}
	}

	return nil
}
