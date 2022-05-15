package services

import (
	"context"
	"fmt"
	"hash/crc64"
	"io/ioutil"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	docMod "github.com/dhf0820/ids_model/document"

	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//"os"
	//log "github.com/sirupsen/logrus"
	"testing"

	dm "github.com/dhf0820/ids_document/pkg"
	"github.com/joho/godotenv"
	//m "github.com/dhf0820/ids_model"
)

func TestProcessDoc(t *testing.T) {
	godotenv.Load(".env.document_test")
	fmt.Printf("\n\nTestProcessDocument\n")
	InitTest()
	OpenDB()
	Convey("Subject: Save the image", t, func() {
		doc := MakeTestDocument()
		tId, _ := primitive.ObjectIDFromHex("61aa40dd995d52fdb964290c")
		doc.ImageID = tId

		//doc.TempImageID = tId

		image, err := ioutil.ReadFile("./ClinDocMammo.pdf")
		So(err, ShouldBeNil)
		imageLen := len(image)
		So(imageLen, ShouldNotEqual, 0)
		ingress := dm.IngressDocument{}
		ingress.Image = image
		ingress.Document = doc
		document, err := ProcessDocument(doc)
		So(err, ShouldBeNil)
		So(document, ShouldNotBeNil)
		//fmt.Printf("\n\n$$$\nFinal Document: %s\n", spew.Sdump(document))
	})

}

func TestWriteGridfs(t *testing.T) {
	godotenv.Load(".env.document_test")
	fmt.Printf("\n\nTestWriteGridFs\n")
	InitTest()
	OpenDB()
	//conf := GetConfig()

	Convey("Subject: Save the image", t, func() {
		Convey("Image does not exist", func() {
			//var checksum uint64

			//doc := MakeTestDocument()
			image, err := ioutil.ReadFile("./ClinDocMammo.pdf")
			So(err, ShouldBeNil)
			imageLen := len(image)
			So(imageLen, ShouldNotEqual, 0)
			crc64Table := crc64.MakeTable(crc64.ECMA)
			checksum := crc64.Checksum(image, crc64Table)
			fmt.Printf("CheckSum: %d\n", checksum)
			metaData := make(map[string]string)
			metaData["checksum"] = strconv.FormatUint(checksum, 10)
			id, err := WriteGridFs(metaData, image)
			So(err, ShouldBeNil)
			So(id, ShouldNotEqual, primitive.NilObjectID)
			fmt.Printf("Image ID : %s\n", id)
		})
	})
}

func TestGetImage(t *testing.T) {
	godotenv.Load(".env.document_test")
	fmt.Printf("\n\nTestGetImage\n")
	InitTest()
	OpenDB()
	Convey("Subject: Save the image", t, func() {
		doc := MakeTestDocument()
		tId, _ := primitive.ObjectIDFromHex("61b6d05d3bc9576296633351")
		doc.ImageID = tId
		//doc.TempImageID = tId
		//
		//imageId := doc.TempImageID

		//image, err := ioutil.ReadFile("./ClinDocMammo.pdf")
		//So(err, ShouldBeNil)
		//imageLen := len(image)
		//So(imageLen, ShouldNotEqual, 0)
		//ingress := dm.IngressDocument{}
		//ingress.Image = image
		//ingress.Document = doc
		//document, err := ProcessDocument(doc)
		//So(err, ShouldBeNil)
		//So(document, ShouldNotBeNil)
		//fmt.Printf("\n\n$$$\nFinal Document: %s\n", spew.Sdump(document))
		newImage, err := getGridFsImage(doc.ImageID)
		So(err, ShouldBeNil)
		So(newImage, ShouldNotBeNil)
		ioutil.WriteFile("debbie.pdf", *newImage, 0664)

		//So(len(newImage), ShouldEqual,imageLen)
	})
}

func TestGetDocument(t *testing.T) {
	godotenv.Load(".env.document_test")
	fmt.Printf("\n\nTestGetDocument\n")
	InitTest()
	OpenDB()
	var docId primitive.ObjectID
	//var imageId primitive.ObjectID
	//var doc *pkg.Document
	Convey("Subject: GetDocument", t, func() {
		Convey("Get document details only", func() {
			fmt.Println("GetDocument Data Only")
			docId, _ = primitive.ObjectIDFromHex("61b6c9f6ed6bd5a1bbc9e9c0")
			doc, err := GetDocument(docId)
			So(err, ShouldBeNil)
			So(doc, ShouldNotBeNil)
			fmt.Printf("Document Found: %s\n", spew.Sdump(doc))
			//imageId = doc.ImageID
		})

		Convey("Get document Image Only", func() {
			//docId, _ = primitive.ObjectIDFromHex("61b6c89bd8cc9541dde09f1b")
			image, err := GetDocumentImage(docId)
			So(err, ShouldBeNil)
			So(image, ShouldNotBeNil)
			So(len(*image), ShouldEqual, 771234)
		})
	})
}

func TestAddDocument(t *testing.T) {
	//t.Parallel()
	godotenv.Load(".env.document_test")
	fmt.Printf("\n\nTestAddDocument\n")
	InitTest()
	OpenDB()
	//conf := GetConfig()

	Convey("Subject: AddDocument", t, func() {
		Convey("New document no correlations", func() {
			//os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
			fmt.Printf("\n\n--- Convey New document no corelation")
			//conf, err :=Initialize("_test")
			//So(conf, ShouldNotBeNil)
			//fmt.Printf("Config received:\n%s\n", spew.Sdump(conf))
			mongo := OpenDB()
			So(mongo, ShouldNotBeNil)
			c, err := GetCollection("documents")
			So(err, ShouldBeNil)
			So(c, ShouldNotBeNil)
			doc := MakeTestDocument()
			fmt.Printf("39 - New document id: %s\n", doc.ID.Hex())
			d, err := InsertDocument(context.Background(), doc)
			So(err, ShouldBeNil)
			So(d, ShouldNotBeNil)
			//fmt.Printf("insdoc: %s\n", spew.Sdump(d))
		})
		Convey("Valid New Document", func() {
			//os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
			fmt.Printf("\n\n--- Convey  Valid New Document\n")
			//conf, err :=Initialize("local_test")
			//So(conf, ShouldNotBeNil)
			//fmt.Printf("Config received:\n%s\n", spew.Sdump(conf))
			// mongo := OpenDB()
			// So(mongo, ShouldNotBeNil)
			c, err := GetCollection("documents")
			So(err, ShouldBeNil)
			So(c, ShouldNotBeNil)
			doc := MakeTestDocument()
			fmt.Printf("New document id: %s\n", doc.ID.Hex())
			d, err := InsertDocument(context.Background(), doc)
			So(err, ShouldBeNil)
			So(d, ShouldNotBeNil)
			//fmt.Printf("insdoc: %s\n", spew.Sdump(d))
		})
		Convey("Duplicate New Document", func() {
			//os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
			fmt.Printf("\n\n--- Convey Duplicate New Document")
			//conf, err :=Initialize("local_test")
			// So(conf, ShouldNotBeNil)
			//fmt.Printf("Config received:\n%s\n", spew.Sdump(conf))
			mongo := OpenDB()
			So(mongo, ShouldNotBeNil)
			c, err := GetCollection("documents")
			So(err, ShouldBeNil)
			So(c, ShouldNotBeNil)
			doc := MakeTestDocument()
			//fmt.Printf("New Document: %s\n", spew.Sdump(doc))
			//fmt.Printf("New document id: %s\n", doc.ID.Hex())
			d, err := InsertDocument(context.Background(), doc)
			So(err, ShouldBeNil)
			So(d, ShouldNotBeNil)
			//fmt.Printf("insdoc: %s\n", spew.Sdump(d))
		})
	})
}

func TestFindDocuments(t *testing.T) {
	//t.Parallel()
	godotenv.Load(".env.core")
	fmt.Printf("\n\nTestFindDocuments\n")
	InitTest()
	OpenDB()
	//conf := GetConfig()

	Convey("Subject: Determine if document already is in repository ", t, func() {
		Convey("Document Exist", func() {
			fmt.Printf("\n\n--- Convey FindDocuments\n")
			c, err := GetCollection("documents")
			So(err, ShouldBeNil)
			So(c, ShouldNotBeNil)
			filter := docMod.DocumentSearchFilter{}
			filter.ReferenceType = "source"
			filter.ReferenceID = "1111"
			doc, err := DocumentExists(context.Background(), &filter)
			So(err, ShouldNotBeNil)
			So(doc, ShouldNotBeNil)
			//fmt.Printf("found: %s\n", spew.Sdump(doc))
		})
		Convey("List of matchng documents", func() {
			fmt.Printf("\n\n--- Convey List of matching documents\n")
			docs := []*docMod.Document{}
			c, err := GetCollection("documents")
			So(err, ShouldBeNil)
			So(c, ShouldNotBeNil)
			filter := docMod.DocumentSearchFilter{}
			filter.ReferenceType = "release"
			filter.ReferenceID = "R-001"
			docs, err = FindDocuments(context.Background(), &filter)
			So(err, ShouldBeNil)
			ShouldEqual(2, len(docs))
			//fmt.Printf("found: %s\n", spew.Sdump(docs))
		})
		Convey("Document does not Exist", func() {
			fmt.Printf("\n\n--- Convey Document does not exist\n")
			c, err := GetCollection("documents")
			So(err, ShouldBeNil)
			So(c, ShouldNotBeNil)
			filter := docMod.DocumentSearchFilter{}
			filter.ReferenceType = "source"
			filter.ReferenceID = "1110"
			doc, err := DocumentExists(context.Background(), &filter)
			So(err, ShouldBeNil)
			So(doc, ShouldBeNil)
		})
	})
}

func MakeTestDocument() *docMod.Document {
	doc := docMod.Document{}
	src := docMod.CorrelationId{}
	rel := docMod.CorrelationId{}
	src.ReferenceType = "source"
	src.ReferenceID = "9999"
	src.OriginatingIP = "192.168.1.26"
	src.SystemFacility = "demo"
	doc.CorrelationIDs = append(doc.CorrelationIDs, &src)
	rel.ReferenceType = "source"
	rel.ReferenceID = "1271958"
	rel.SystemFacility = "demo"
	doc.ID = primitive.NewObjectID()
	doc.CorrelationIDs = append(doc.CorrelationIDs, &rel)
	doc.ImageType = "pdf"
	return &doc
}
