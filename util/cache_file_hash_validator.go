package util

import "fmt"

// SHA1SumIsValidForCacheFile verifies the integrity of a given cache file by generating
// and checking the SHA1Sum of the cache file compared to the correlating schema's hash.
func SHA1SumIsValidForCacheFile(args SHA1SumValidatorArgs) error {
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

	// Verify the correlating schema's hash equals the SHA1Sum of the cache file.
	if fileSHA1Sum != args.FileName {
		return fmt.Errorf(
			"cache file '%s' did not equal the SHA1Sum of the correlating cache blob",
			args.FileName,
		)
	}

	return nil
}
