package services

import (
	"bytes"
	"fmt"
	"time"

	docMod "github.com/dhf0820/ids_model/document"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//func (doc *Document) WriteGridFs(metaData map[string]string, imageData []byte) (string, error) {
func WriteGridFs(metaData map[string]string, imageData []byte) (primitive.ObjectID, error) {
	fmt.Printf("WriteGridFs entered\n")
	//fmt.Printf("Database: %s\n", spew.Sdump(DB))
	startTime := time.Now()
	//mdb := db.DB.Database
	//fmt.Printf("Database: %s\n", spew.Sdump(mdb))
	//tst := dm.Document{}
	//fmt.Printf("dm: %s\n", spew.Sdump(tst))
	bucket, err := gridfs.NewBucket(
		DB.Database,
		options.GridFSBucket().SetName("fs"),
	)
	if err != nil {
		err = fmt.Errorf("Unable to get GridFS Bucket: %s", err)
		fmt.Printf("%s", err)
		return primitive.NilObjectID, err
	}
	//client := "harman"
	facility := metaData["facility"]
	mrn := metaData["mrn"]
	srcID := metaData["src_id"]
	//docId := primitive.NewObjectID()
	//metaData := make(map[string]string)
	//metaData["content_type"] = "pdf"
	//metaData["mrn"] = mrn
	//metadata["checksum] = "]
	//fileName := fmt.Sprintf("%s_%s_%s_%s", doc.Client, doc.Facility, doc.MRN, doc.ID.Hex())
	fileName := fmt.Sprintf("%s_%s_%s", facility, srcID, mrn)
	saveImage, err := bucket.OpenUploadStream(
		fileName,
		options.GridFSUpload().SetMetadata(metaData),
	)
	if err != nil {
		err = fmt.Errorf("OpenUploadStream failed: %s", err)
		fmt.Println(err)
		return primitive.NilObjectID, err
	}
	defer saveImage.Close()

	fileSize, err := saveImage.Write(imageData)
	if err != nil {
		err = fmt.Errorf("Save GridFS failed: %v", err)
		fmt.Println(err)
		return primitive.NilObjectID, err
	}
	fmt.Printf("Gridfs Saved %d bytes in %f seconds", fileSize, time.Since(startTime).Seconds())
	//imageIDHex := saveImage.FileID.(primitive.ObjectID).Hex()
	id := saveImage.FileID.(primitive.ObjectID)
	return id, nil
}

func getGridFsImage(imageID primitive.ObjectID) (*[]byte, error) {
	db := DB.Database
	startTime := time.Now()
	// fsFiles := db.Collection("fs.files")
	bucket, err := gridfs.NewBucket(
		db,
		options.GridFSBucket().SetName("fs"),
	)
	if err != nil {
		fmt.Errorf("New Bucket failed: %s", err)
	}
	var buf bytes.Buffer
	//fmt.Printf("Looking for imageid: %s\n", imageID)
	//oid, _ := primitive.ObjectIDFromHex("5ed406756cd109a290b17231")
	fmt.Printf("Retrieving ID:%s\n", imageID)
	//oid, err := primitive.ObjectIDFromHex(imageID)
	//if err != nil {
	//	err := fmt.Errorf("Document ImageID is not valid: %s  err: %s", imageID, err)
	//	fmt.Println(err)
	//	return nil, err
	//}
	_, err = bucket.DownloadToStream(imageID, &buf)
	if err != nil {
		fmt.Printf("DownloadToStream failed: %v", err)
	}
	image := buf.Bytes()
	fmt.Printf("Gridfs retrieved in %f seconds", time.Since(startTime).Seconds())
	fmt.Printf("Length: %d\n", len(image))
	//dStream, err := bucket.DownloadToStreamByName("Create_test_123456_5ed4003e030ea422d7210766", &buf)

	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// var results bson.M
	// fmt.Printf("Looking for imageid: %s\n", doc.ImageID)
	// //iod, _ := primitive.ObjectIDFromHex(doc.ImageID)
	// filter := bson.D{}
	// err := fsFiles.FindOne(ctx, filter).Decode(&results)
	// if err != nil {
	// 	err := fmt.Errorf("GetGridFsImage did not find image: %s  Err: %s", filter, err)
	// 	log.Errorln(err)
	// 	return nil, err
	// }
	// fmt.Printf("Found Image: %s\n", spew.Sdump(results))
	return &image, nil
}

func DeleteGridFs(doc *docMod.Document) error {
	if doc.ImageID == primitive.NilObjectID {
		return nil
	}

	bucket, err := gridfs.NewBucket(

		DB.Database,
	)
	imgId := doc.ImageID //primitive.ObjectIDFromHex(doc.ImageID)
	//fmt.Printf("Deleting gridfs id: %s\n", imgId)
	if err != nil {
		err = fmt.Errorf("Unable to get GridFS Bucket: %s", err)
		log.Errorf("%s", err)
		return err
	}

	err = bucket.Delete(imgId)
	if err != nil {
		fmt.Printf("Delete image %s failed: %s\n", imgId, err)
	}
	return nil
}
