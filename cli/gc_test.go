package cli

import (
	"errors"
	"os"
	"testing"

	"github.com/dailymuse/git-fit/config"
	. "github.com/smartystreets/goconvey/convey"
)

func TestValidateCacheFile(t *testing.T) {
	Convey("Given a file configuration", t, func() {
		Convey("if reading the cache directory fails, it should return an error", func() {
			config := config.Config{
				Version: 1,
				Files: map[string]string{
					"etc/pg.dump.tar.gz": "69e542360fa6f81b704199685432ddea1dc60944",
				},
			}

			err := Gc(GCArgs{
				Args:   []string{"1"},
				Schema: &config,
				ReadDir: func(dirname string) ([]os.FileInfo, error) {
					So(dirname, ShouldEqual, ".git/fit")
					return nil, errors.New("could not read cache dir")
				},
				LogError: func(format string, args ...interface{}) {},
			})

			So(err, ShouldNotBeNil)
		})

		Convey("if the cache file sha does not exist in the schema, remove the cache and return nil", func() {
			config := config.Config{
				Version: 1,
				Files: map[string]string{
					"etc/pg.dump.tar.gz": "69e542360fa6f81b704199685432ddea1dc60944",
				},
			}

			file := NewMockFileInfo("not-a-real-sha-this-should-not-exist-in-schema")
			files := []os.FileInfo{file}

			err := Gc(GCArgs{
				Args:   []string{"1"},
				Schema: &config,
				ReadDir: func(dirname string) ([]os.FileInfo, error) {
					So(dirname, ShouldEqual, ".git/fit")
					return files, nil
				},
				LogError: func(format string, args ...interface{}) {
					return
				},
				RemoveFile: func(filePath string) error {
					So(filePath, ShouldEqual, ".git/fit/not-a-real-sha-this-should-not-exist-in-schema")
					return nil
				},
			})

			So(err, ShouldBeNil)
		})

		Convey("if the sha1sum of a cached file does not match the sha1sum correlating to the schema, it should delete the cache", func() {
			Convey("if it fails to delete the cache file, it should return an error", func() {
				config := config.Config{
					Version: 1,
					Files: map[string]string{
						"etc/pg.dump.tar.gz": "69e542360fa6f81b704199685432ddea1dc60944",
					},
				}

				file := NewMockFileInfo("69e542360fa6f81b704199685432ddea1dc60944")
				files := []os.FileInfo{file}

				err := Gc(GCArgs{
					Args:   []string{"1"},
					Schema: &config,
					ReadDir: func(dirname string) ([]os.FileInfo, error) {
						So(dirname, ShouldEqual, ".git/fit")
						return files, nil
					},
					LogError: func(format string, args ...interface{}) {
						return
					},
					SHA1SumGenerator: func(p string) (string, error) {
						So(p, ShouldEqual, ".git/fit/69e542360fa6f81b704199685432ddea1dc60944")
						return "69e542360fa6f81b704", nil
					},
					RemoveFile: func(filePath string) error {
						So(filePath, ShouldEqual, ".git/fit/69e542360fa6f81b704199685432ddea1dc60944")
						return errors.New("failed to delete cache file")
					},
				})

				So(err, ShouldNotBeNil)
			})

			Convey("if the cache was successfully deleted it should return nil", func() {
				config := config.Config{
					Version: 1,
					Files: map[string]string{
						"etc/pg.dump.tar.gz": "69e542360fa6f81b704199685432ddea1dc60944",
					},
				}

				file := NewMockFileInfo("69e542360fa6f81b704199685432ddea1dc60944")
				files := []os.FileInfo{file}

				err := Gc(GCArgs{
					Args:   []string{"1"},
					Schema: &config,
					ReadDir: func(dirname string) ([]os.FileInfo, error) {
						So(dirname, ShouldEqual, ".git/fit")
						return files, nil
					},
					LogError: func(format string, args ...interface{}) {
						return
					},
					SHA1SumGenerator: func(p string) (string, error) {
						So(p, ShouldEqual, ".git/fit/69e542360fa6f81b704199685432ddea1dc60944")
						return "69e542360fa6f81b704", nil
					},
					RemoveFile: func(filePath string) error {
						So(filePath, ShouldEqual, ".git/fit/69e542360fa6f81b704199685432ddea1dc60944")
						return nil
					},
				})

				So(err, ShouldBeNil)
			})
		})
	})
}
