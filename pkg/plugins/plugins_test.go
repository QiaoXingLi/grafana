package plugins

import (
	"path/filepath"
	"testing"

	"github.com/grafana/grafana/pkg/setting"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/ini.v1"
)

func TestPluginScans(t *testing.T) {

	Convey("When scanning for plugins", t, func() {
		setting.StaticRootPath, _ = filepath.Abs("../../public/")
		setting.Raw = ini.Empty()

		pm := &PluginManager{
			Cfg: &setting.Cfg{
				FeatureToggles: map[string]bool{},
			},
		}
		err := pm.Init()

		So(err, ShouldBeNil)
		So(len(DataSources), ShouldBeGreaterThan, 1)
		So(len(Panels), ShouldBeGreaterThan, 1)

		Convey("Should set module automatically", func() {
			So(DataSources["graphite"].Module, ShouldEqual, "app/plugins/datasource/graphite/module")
		})
	})

	Convey("When reading app plugin definition", t, func() {
		pm := &PluginManager{
			Cfg: &setting.Cfg{
				FeatureToggles: map[string]bool{},
				PluginSettings: setting.PluginSettings{
					"nginx-app": map[string]string{
						"path": "testdata/test-app",
					},
				},
			},
		}
		err := pm.Init()
		So(err, ShouldBeNil)

		So(len(Apps), ShouldBeGreaterThan, 0)
		So(Apps["test-app"].Info.Logos.Large, ShouldEqual, "public/plugins/test-app/img/logo_large.png")
		So(Apps["test-app"].Info.Screenshots[1].Path, ShouldEqual, "public/plugins/test-app/img/screenshot2.png")
	})

	Convey("When checking if renderer is backend only plugin", t, func() {
		pluginScanner := &PluginScanner{}
		result := pluginScanner.IsBackendOnlyPlugin("renderer")

		So(result, ShouldEqual, true)
	})

	Convey("When checking if app is backend only plugin", t, func() {
		pluginScanner := &PluginScanner{}
		result := pluginScanner.IsBackendOnlyPlugin("app")

		So(result, ShouldEqual, false)
	})

}
