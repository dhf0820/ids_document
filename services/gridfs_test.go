package services


//func TestGetImage(t *testing.T) {
//	godotenv.Load(".env.document_test")
//	fmt.Printf("\n\nTestGetImage\n")
//	InitTest()
//	OpenDB()
//	Convey("Subject: Save the image", t, func() {
//		doc := MakeTestDocument()
//		doc.TempImageID,_ = primitive.ObjectIDFromHex("61aa40dd995d52fdb964290c")
//		image, err := ioutil.ReadFile("./ClinDocMammo.pdf")
//		So(err, ShouldBeNil)
//		imageLen := len(image)
//		So(imageLen, ShouldNotEqual, 0)
//		ingress := dm.IngressDocument{}
//		ingress.Image = image
//		ingress.Document = doc
//		document, err := ProcessDocument(doc)
//		So(err, ShouldBeNil)
//		So(document, ShouldNotBeNil)
//		fmt.Printf("\n\n$$$\nFinal Document: %s\n", spew.Sdump(document))
//	})
//
//}

//func TestWriteGridfs(t *testing.T) {
//godotenv.Load(".env.document_test")
//fmt.Printf("\n\nTestWriteGridFs\n")
//InitTest()
//OpenDB()
////conf := GetConfig()
//
//Convey("Subject: Save the image", t, func() {
//Convey("Image does not exist", func() {
////var checksum uint64
//
////doc := MakeTestDocument()
//image, err :=ioutil.ReadFile("./ClinDocMammo.pdf")
//So(err,ShouldBeNil)
//imageLen := len(image)
//So(imageLen, ShouldNotEqual, 0)
//crc64Table := crc64.MakeTable(crc64.ECMA)
//checksum := crc64.Checksum(image, crc64Table)
//fmt.Printf("CheckSum: %d\n", checksum)
//metaData := make(map[string]string)
//metaData["checksum"] = strconv.FormatUint(checksum, 10)
//id, err :=WriteGridFs(metaData, image)
//So(err, ShouldBeNil)
//So(id, ShouldNotEqual, primitive.NilObjectID)
//fmt.Printf("Image ID : %s\n", id)
//})
//})
//}
//
//func TestAddDocument(t *testing.T) {
////t.Parallel()
//godotenv.Load(".env.document_test")
//fmt.Printf("\n\nTestAddDocument\n")
//InitTest()
//OpenDB()
////conf := GetConfig()
//
//Convey("Subject: AddDocument", t, func() {
//Convey("New document no correlations", func() {
////os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
//fmt.Printf("\n\n--- Convey New document no corelation")
////conf, err :=Initialize("local")
////So(conf, ShouldNotBeNil)
////fmt.Printf("Config received:\n%s\n", spew.Sdump(conf))
//mongo := OpenDB()
//So(mongo, ShouldNotBeNil)
//c, err := GetCollection("documents")
//So(err, ShouldBeNil)
//So(c, ShouldNotBeNil)
//doc := MakeTestDocument()
//fmt.Printf("39 - New document id: %s\n", doc.ID.Hex())
//d, err := InsertDocument(context.Background(), doc)
//So(err, ShouldBeNil)
//So(d, ShouldNotBeNil)
//fmt.Printf("insdoc: %s\n", spew.Sdump(d))
//})
//Convey("Valid New Document", func() {
////os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
//fmt.Printf("\n\n--- Convey  Valid New Document\n")
////conf, err :=Initialize("local")
////So(conf, ShouldNotBeNil)
////fmt.Printf("Config received:\n%s\n", spew.Sdump(conf))
//// mongo := OpenDB()
//// So(mongo, ShouldNotBeNil)
//c, err := GetCollection("documents")
//So(err, ShouldBeNil)
//So(c, ShouldNotBeNil)
//doc := MakeTestDocument()
//fmt.Printf("New document id: %s\n", doc.ID.Hex())
//d, err := InsertDocument(context.Background(), doc)
//So(err, ShouldBeNil)
//So(d, ShouldNotBeNil)
//fmt.Printf("insdoc: %s\n", spew.Sdump(d))
//})
//Convey("Duplicate New Document", func() {
////os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
//fmt.Printf("\n\n--- Convey Duplicate New Document")
////conf, err :=Initialize("local")
//// So(conf, ShouldNotBeNil)
////fmt.Printf("Config received:\n%s\n", spew.Sdump(conf))
//mongo := OpenDB()
//So(mongo, ShouldNotBeNil)
//c, err := GetCollection("documents")
//So(err, ShouldBeNil)
//So(c, ShouldNotBeNil)
//doc := MakeTestDocument()
//fmt.Printf("New Document: %s\n", spew.Sdump(doc))
//fmt.Printf("New document id: %s\n", doc.ID.Hex())
//d, err := InsertDocument(context.Background(), doc)
//So(err, ShouldBeNil)
//So(d, ShouldNotBeNil)
//fmt.Printf("insdoc: %s\n", spew.Sdump(d))
//})
//})
//}
//
//func TestFindDocuments(t *testing.T) {
////t.Parallel()
//godotenv.Load(".env.core")
//fmt.Printf("\n\nTestFindDocuments\n")
//InitTest()
//OpenDB()
////conf := GetConfig()
//
//Convey("Subject: Determine if document already is in repository ", t, func() {
//Convey("Document Exist", func() {
//fmt.Printf("\n\n--- Convey FindDocuments\n")
//c, err := GetCollection("documents")
//So(err, ShouldBeNil)
//So(c, ShouldNotBeNil)
//filter := dm.DocumentSearchFilter{}
//filter.ReferenceType = "source"
//filter.ReferenceID = "1111"
//doc, err := DocumentExists(context.Background(), &filter)
//So(err, ShouldNotBeNil)
//So(doc, ShouldNotBeNil)
//fmt.Printf("found: %s\n", spew.Sdump(doc))
//})
//Convey("List of matchng documents", func() {
//fmt.Printf("\n\n--- Convey List of matching documents\n")
//docs := []*dm.Document{}
//c, err := GetCollection("documents")
//So(err, ShouldBeNil)
//So(c, ShouldNotBeNil)
//filter := dm.DocumentSearchFilter{}
//filter.ReferenceType = "release"
//filter.ReferenceID = "R-001"
//docs, err = FindDocuments(context.Background(), &filter)
//So(err, ShouldBeNil)
//ShouldEqual(2, len(docs))
//fmt.Printf("found: %s\n", spew.Sdump(docs))
//})
//Convey("Document does not Exist", func() {
//fmt.Printf("\n\n--- Convey Document does not exist\n")
//c, err := GetCollection("documents")
//So(err, ShouldBeNil)
//So(c, ShouldNotBeNil)
//filter := dm.DocumentSearchFilter{}
//filter.ReferenceType = "source"
//filter.ReferenceID = "1110"
//doc, err := DocumentExists(context.Background(), &filter)
//So(err, ShouldBeNil)
//So(doc, ShouldBeNil)
//})
//})
//}
//
//func MakeTestDocument() *dm.Document {
//	doc := dm.Document{}
//	src := dm.CorrelationId{}
//	rel := dm.CorrelationId{}
//	src.ReferenceType = "source"
//	src.ReferenceID = "9999"
//	src.OriginatingIP = "192.168.1.26"
//	src.SystemFacility = "demo"
//	doc.CorrelationIDs = append(doc.CorrelationIDs, &src)
//	rel.ReferenceType = "source"
//	rel.ReferenceID = "1271958"
//	rel.SystemFacility = "demo"
//	doc.ID = primitive.NewObjectID()
//	doc.CorrelationIDs = append(doc.CorrelationIDs, &rel)
//	doc.ImageType = "pdf"
//	return &doc
//}
//
//
