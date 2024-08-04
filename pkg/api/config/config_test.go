package config_test

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"zotregistry.dev/zot/pkg/api/config"
)

func TestConfig(t *testing.T) {
	Convey("Test config utils", t, func() {
		firstStorageConfig := config.StorageConfig{
			GC: true, Dedupe: true,
			GCDelay: 1 * time.Minute, GCInterval: 1 * time.Hour,
		}
		secondStorageConfig := config.StorageConfig{
			GC: true, Dedupe: true,
			GCDelay: 1 * time.Minute, GCInterval: 1 * time.Hour,
		}

		So(firstStorageConfig.ParamsEqual(secondStorageConfig), ShouldBeTrue)

		firstStorageConfig.GC = false

		So(firstStorageConfig.ParamsEqual(secondStorageConfig), ShouldBeFalse)

		firstStorageConfig.GC = true
		firstStorageConfig.Dedupe = false

		So(firstStorageConfig.ParamsEqual(secondStorageConfig), ShouldBeFalse)

		firstStorageConfig.Dedupe = true
		firstStorageConfig.GCDelay = 2 * time.Minute

		So(firstStorageConfig.ParamsEqual(secondStorageConfig), ShouldBeFalse)

		firstStorageConfig.GCDelay = 1 * time.Minute
		firstStorageConfig.GCInterval = 2 * time.Hour

		So(firstStorageConfig.ParamsEqual(secondStorageConfig), ShouldBeFalse)

		firstStorageConfig.GCInterval = 1 * time.Hour

		So(firstStorageConfig.ParamsEqual(secondStorageConfig), ShouldBeTrue)

		isSame, err := config.SameFile("test-config", "test")
		So(err, ShouldNotBeNil)
		So(isSame, ShouldBeFalse)

		dir1 := t.TempDir()

		isSame, err = config.SameFile(dir1, "test")
		So(err, ShouldNotBeNil)
		So(isSame, ShouldBeFalse)

		dir2 := t.TempDir()

		isSame, err = config.SameFile(dir1, dir2)
		So(err, ShouldBeNil)
		So(isSame, ShouldBeFalse)

		isSame, err = config.SameFile(dir1, dir1)
		So(err, ShouldBeNil)
		So(isSame, ShouldBeTrue)
	})

	Convey("Test DeepCopy() & Sanitize()", t, func() {
		conf := config.New()
		So(conf, ShouldNotBeNil)
		authConfig := &config.AuthConfig{LDAP: (&config.LDAPConfig{}).SetBindPassword("oina")}
		conf.HTTP.Auth = authConfig
		So(func() { conf.Sanitize() }, ShouldNotPanic)
		conf = conf.Sanitize()
		So(conf.HTTP.Auth.LDAP.BindPassword(), ShouldEqual, "******")

		// negative
		obj := make(chan int)
		err := config.DeepCopy(conf, obj)
		So(err, ShouldNotBeNil)
		err = config.DeepCopy(obj, conf)
		So(err, ShouldNotBeNil)
	})

	Convey("Test IsRetentionEnabled()", t, func() {
		conf := config.New()
		So(conf.IsRetentionEnabled(), ShouldBeFalse)

		conf.Storage.Retention.Policies = []config.RetentionPolicy{
			{
				Repositories: []string{"repo"},
			},
		}

		So(conf.IsRetentionEnabled(), ShouldBeFalse)

		policies := []config.RetentionPolicy{
			{
				Repositories: []string{"repo"},
				KeepTags: []config.KeepTagsPolicy{
					{
						Patterns:                []string{"tag"},
						MostRecentlyPulledCount: 2,
					},
				},
			},
		}

		conf.Storage.Retention = config.ImageRetention{
			Policies: policies,
		}

		So(conf.IsRetentionEnabled(), ShouldBeTrue)

		subPaths := make(map[string]config.StorageConfig)

		subPaths["/a"] = config.StorageConfig{
			GC: true,
			Retention: config.ImageRetention{
				Policies: policies,
			},
		}

		conf.Storage.SubPaths = subPaths

		So(conf.IsRetentionEnabled(), ShouldBeTrue)
	})

	Convey("Test IsClusteringEnabled()", t, func() {
		conf := config.New()
		So(conf.Cluster, ShouldBeNil)
		So(conf.IsClusteringEnabled(), ShouldBeFalse)

		conf.Cluster = &config.ClusterConfig{
			Members: []string{
				"127.0.0.1:9090",
				"127.0.0.1:9001",
			},
			HashKey: "loremipsumdolors",
		}

		So(conf.IsClusteringEnabled(), ShouldBeTrue)
	})

	Convey("Test IsSharedStorageEnabled()", t, func() {
		conf := config.New()
		So(conf.Storage.RemoteCache, ShouldBeFalse)
		So(conf.Storage.CacheDriver, ShouldBeNil)
		So(conf.Storage.StorageDriver, ShouldBeNil)
		So(conf.IsSharedStorageEnabled(), ShouldBeFalse)

		conf.Storage.RemoteCache = true
		So(conf.Storage.RemoteCache, ShouldBeTrue)
		So(conf.IsSharedStorageEnabled(), ShouldBeFalse)

		storageDriver := map[string]interface{}{"name": "s3"}
		conf.Storage.StorageDriver = storageDriver
		So(conf.Storage.StorageDriver, ShouldNotBeNil)
		So(conf.IsSharedStorageEnabled(), ShouldBeFalse)

		cacheDriver := map[string]interface{}{"name": "dynamodb"}
		conf.Storage.CacheDriver = cacheDriver
		So(conf.Storage.CacheDriver, ShouldNotBeNil)
		So(conf.IsSharedStorageEnabled(), ShouldBeTrue)
	})
}
