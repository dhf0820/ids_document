package services

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	//log "github.com/sirupsen/logrus"
	"testing"
	//"github.com/joho/godotenv"
)

func TestGetServiceConfig(t *testing.T) {
	//t.Parallel()
	InitTest()

	fmt.Printf("\n\nGetServiceConfig\n")
	Convey("Subject: GetServiceConfig", t, func() {
		//os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
		fmt.Printf("\n\n--- Convey GetServiceConfig\n")
		os.Setenv("CONFIG_ADDRESS", "http://localhost:19100/api/v1/")
		os.Setenv("SERVICE_NAME", "document")
		os.Setenv("SERVICE_VERSION", "local_test")
		os.Setenv("COMPANY", "demo")
		cfg, err := GetServiceConfig("document", "local_test", "demo", "")
		So(err, ShouldBeNil)
		So(cfg, ShouldNotBeNil)
		So(cfg.Version, ShouldEqual, "local_test")
		So(cfg.Name, ShouldEqual, "document")
		//fmt.Printf("\nReceived cfg: = %s\n", spew.Sdump(cfg))
	})
}

func TestInitialize(t *testing.T) {
	//t.Parallel()
	// InitTest()

	fmt.Printf("\n\nInitialize\n")
	Convey("Subject: Initialize", t, func() {
		fmt.Printf("\n\n--- Convey Initalize\n")
		os.Setenv("CONFIG_ADDRESS", "http://localhost:19100/api/rest/v1/")
		os.Setenv("SERVICE_NAME", "document")
		os.Setenv("SERVICE_VERSION", "local_test")
		cfg, err := Initialize("local_test", "demo")
		So(err, ShouldBeNil)
		So(cfg, ShouldNotBeNil)
		So(cfg.Version, ShouldEqual, "local_test")
		So(cfg.Name, ShouldEqual, "document")
		//fmt.Printf("\nReceived cfg: = %s\n", spew.Sdump(cfg))
	})
}
