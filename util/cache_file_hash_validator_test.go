package util

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestValidateCacheFile(t *testing.T) {
	Convey("Given a cache file hash", t, func() {

		Convey("if the file hash does not exist in the schema, it should return an error", func() {
			err := SHA1sumIsValidForCacheFile(SHA1sumValidatorArgs{
				ReadDir:                 "./git/fit",
				FileName:                "69e542360fa6f81b704199685432ddea1dc6094",
				CacheFileHashesInSchema: map[string]bool{},
			})
			So(err.Error(), ShouldEqual, "file hash '69e542360fa6f81b704199685432ddea1dc6094' does not exist in schema")
			So(err, ShouldNotBeNil)
		})

		Convey("if the SHA1Sum generator fails to generate a SHA from cache file, it should return an error", func() {
			err := SHA1sumIsValidForCacheFile(SHA1sumValidatorArgs{
				ReadDir:  "./git/fit",
				FileName: "69e542360fa6f81b704199685432ddea1dc60944",
				GenerateSHA1Sum: func(p string) (string, error) {
					So(p, ShouldEqual, fmt.Sprintf("./git/fit/69e542360fa6f81b704199685432ddea1dc60944"))
					return "", errors.New("error generating file SHA1Sum")
				},
				CacheFileHashesInSchema: map[string]bool{"69e542360fa6f81b704199685432ddea1dc60944": true},
			})
			So(err, ShouldNotBeNil)
		})

		Convey("if the SHA1Sum of the cache file does not match the correlating schema SHA1Sum, it should return an error", func() {
			err := SHA1sumIsValidForCacheFile(SHA1sumValidatorArgs{
				ReadDir:  "./git/fit",
				FileName: "69e542360fa6f81b704199685432ddea1dc60944",
				GenerateSHA1Sum: func(p string) (string, error) {
					So(p, ShouldEqual, fmt.Sprintf("./git/fit/69e542360fa6f81b704199685432ddea1dc60944"))
					return "69e542360fa6", nil
				},
				CacheFileHashesInSchema: map[string]bool{"69e542360fa6f81b704199685432ddea1dc60944": true},
			})
			So(err, ShouldNotBeNil)
		})

		Convey("if the SHA1Sum of the cache file matches the correlating schema SHA1Sum, it should return nil", func() {
			err := SHA1sumIsValidForCacheFile(SHA1sumValidatorArgs{
				ReadDir:  "./git/fit",
				FileName: "69e542360fa6f81b704199685432ddea1dc60944",
				GenerateSHA1Sum: func(p string) (string, error) {
					So(p, ShouldEqual, fmt.Sprintf("./git/fit/69e542360fa6f81b704199685432ddea1dc60944"))
					return "69e542360fa6f81b704199685432ddea1dc60944", nil
				},
				CacheFileHashesInSchema: map[string]bool{"69e542360fa6f81b704199685432ddea1dc60944": true},
			})
			So(err, ShouldBeNil)
		})
	})
}
