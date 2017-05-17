package util

// SHA1SumValidatorArgs is the arguments struct for SHA1SumIsValidForCacheFile and
// SHA1SumIsValidForCacheFile.
type SHA1SumValidatorArgs struct {
	FileName                string
	ReadDir                 string
	CacheFileHashesInSchema map[string]bool
	GenerateSHA1Sum         SHA1SumGenerator
}

// SHA1SumGenerator is responsible for generating a SHA1Sum from a byte array.
type SHA1SumGenerator func(p string) (string, error)
