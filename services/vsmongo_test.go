package services

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	"os"

	. "github.com/smartystreets/goconvey/convey"

	//log "github.com/sirupsen/logrus"
	"testing"
	//"github.com/joho/godotenv"
)

func TestOpenDB(t *testing.T) {
	//t.Parallel()
	InitTest()

	fmt.Printf("\n\nTestOpenDB\n")
	Convey("Subject: Open the mongo DB", t, func() {
		//os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
		fmt.Printf("\n\n--- Convey")
		os.Setenv("CONFIG_ADDRESS", "http://test_core:19900/api/v1/")
		os.Setenv("SERVICE_NAME", "document")
		os.Setenv("SERVICE_VERSION", "local")
		os.Setenv("COMPANY", "demo")
		InitTest()
		conf := GetConfig()

		So(conf, ShouldNotBeNil)
		mongo := OpenDB()
		So(mongo, ShouldNotBeNil)
		c, err := GetCollection("documents")
		So(err, ShouldBeNil)
		So(c, ShouldNotBeNil)
	})
}

func TestConnectToDB(t *testing.T) {
	//t.Parallel()
	InitTest()
	//godotenv.Load(".env.core")
	fmt.Printf("\n\nTestConnectToDB \n")
	Convey("Subject: Open the mongo DB", t, func() {
		//os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
		fmt.Printf("\n\n--- Convey TestConnectToDB\n")

		mongo, err := ConnectToDB()
		So(mongo, ShouldNotBeNil)
		So(err, ShouldBeNil)
		c, err := GetCollection("configs")
		So(err, ShouldBeNil)
		So(c, ShouldNotBeNil)
	})
}
