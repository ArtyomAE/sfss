package util

import "fmt"

// SHA1sumIsValidForCacheFile verifies the integreity of a given cache file by generating
// and checking the sha1sum of the cache file compared to the correlating schema's hash.
func SHA1sumIsValidForCacheFile(args SHA1sumValidatorArgs) error {
	// Check to see if the file name SHA we're validating exists in the schema.
	_, cacheObjExistsInSchema := args.CacheFileHashesInSchema[args.FileName]
	if !cacheObjExistsInSchema {
		return fmt.Errorf(
			"file hash '%s' does not exist in schema",
			args.FileName,
		)
	}

	// Attempt to generate the SHA1Sum for the cache file.
	filePath := fmt.Sprintf("%s/%s", args.ReadDir, args.FileName)
	fileSHA1Sum, err := args.GenerateSHA1Sum(filePath)
	if err != nil {
		return fmt.Errorf(
			"could not generate file hash: %s",
			err.Error(),
		)
	}

	// Check to make the correlating schema's hash equals the sha1sum of the cache file.
	if fileSHA1Sum != args.FileName {
		return fmt.Errorf(
			"cache file '%s' did not equal the sha1sum of the correlating cache blob",
			args.FileName,
		)
	}

	return nil
}
