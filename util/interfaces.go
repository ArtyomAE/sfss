package util

// SHA1sumValidatorArgs is the arguments struct for SHA1sumIsValidForCacheFile and
// SHA1sumIsValidForCacheFile.
type SHA1sumValidatorArgs struct {
	FileName                string
	ReadDir                 string
	CacheFileHashesInSchema map[string]bool
	GenerateSHA1Sum         sha1SumGenerator
}

// sha1SumGenerator is responsible for generating a sha1sum from a byte array.
type sha1SumGenerator func(p string) (string, error)
