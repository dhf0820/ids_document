package services

//"fmt"
//pkg "github.com/dhf0820/ids_document/pkg"

//	"bytes"
//	"fmt"
//	"os"
//	"io/ioutil"
//	"net/http"
//	"time"
//	"github.com/google/uuid"
//
//	log "github.com/sirupsen/logrus"
//	dm "github.com/dhf0820/ids_document/pkg"
//	//"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	//"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/gridfs"
//	"go.mongodb.org/mongo-driver/mongo/options"
//	"google.golang.org/grpc/codes"
//	"google.golang.org/grpc/status"

//
////ImageStore and interface to handle storing a documents image various ways
//type ImageSaver interface {
//	//Save the image
//	Save(image *bytes.Buffer) error
//}
//
////InRecordStore manages storeing a document image inside the document record.
//// type InRecordStore struct {
//// 	Document  *dm.Document
//// 	ImageType string
//// }
//
////GridFsStore saves the document image in GridFS and puts the id in the document record
//type GridFsStore struct {
//	Document  *dm.Document
//	ImageType string
//}
//
////FileStore saves the document image in a file in specified folder and extension
//type FileStore struct {
//	FileType string
//	Path     string
//	ImageID  string
//}
//
////NewFileStore returns a new FileStore with the provided params in the struct
//func NewFileStore(fileType, path, imageID string) *FileStore {
//	return &FileStore{
//		FileType: fileType,
//		Path:     path,
//		ImageID:  imageID,
//	}
//}
//
////NewInRecordStore returns a new InRecordStore
//// func NewInRecordStore(doc *dm.Document, imageType string) *InRecordStore {
//// 	return &InRecordStore{
//// 		Document:  doc,
//// 		ImageType: imageType,
//// 	}
//// }
//
////NewGridFsStore returns a new InRecordStore
//func NewGridFsStore(doc *dm.Document, imageType string) *GridFsStore {
//	return &GridFsStore{
//		Document:  doc,
//		ImageType: imageType,
//	}
//}
//
//
//// DO NOT PLAN ON SUPPORTING THIS METHOD OF STORING.
////Save saves the document image using InRecordStore
//// func (store *InRecordStore) Save(imageData *bytes.Buffer, imageType string) error {
//// 	//log.Printf("StoringDocument: %s\n", spew.Sdump(store.Document))
//// 	//bytes := imageData.Bytes()
//// 	//store.Document.Image = bytes
//// 	//store.Document.Facility = "TEST"
//// 	store.ImageType = imageType
//// 	err := store.Document.Update()// FInd the document and update it in total with this
//// 	if err != nil {
//// 		return logError(status.Errorf(codes.Internal, "cannot write image to document: %v", err))
//// 	}
//// 	//log.Printf("Stored image in document: %s\n", store.Document.ID.Hex())
//// 	return nil
//// }
//
////Save saves the document image using GridFS and placing the id in the record
////func (store *GridFsStore) Save(imageData *bytes.Buffer) error {
////	log.Printf("Save image in GridFs\n")
////	bytes := imageData.Bytes()
////	fmt.Printf("Document: %v\n", store.Document)
////	imageID, err := store.Document.WriteGridFs(bytes)
////	if err != nil {
////		return logError(status.Errorf(codes.Internal, "GridFs failed: %v", err))
////	}
////	//fmt.Printf("store.ImageType: %s\n", store.ImageType)
////	store.Document.ImageID = imageID
////	store.Document.ImageType = store.ImageType
////	store.Document.ImageRepository = "GR"
////	store.Document.URL = "https://VertiSoft.com"
////	err = store.Document.Update()
////	if err != nil {
////		return logError(status.Errorf(codes.Internal, "Update Document.ImageID failed: %v", err))
////	}
////	fmt.Printf("Image stored in Gridfs: %s\n", imageID)
////	return nil
////}
//
////Save saves the image in to a regular file
//func (store *FileStore) Save(imageData *bytes.Buffer) error {
//	imageID, err := uuid.NewRandom()
//	if err != nil {
//		return fmt.Errorf("cannot generate image id: %w", err)
//	}
//	outFilePath := fmt.Sprintf("%s/%s%s", store.Path, imageID, store.FileType)
//	file, err := os.Create(outFilePath)
//	if err != nil {
//		return fmt.Errorf("cannot create image file: %w", err)
//	}
//	_, err = imageData.WriteTo(file)
//	if err != nil {
//		file.Close()
//		return fmt.Errorf("cannot write image to file: %w", err)
//	}
//	file.Close()
//	return nil
//
//}
//
//func (doc *dm.Document) DeleteGridFs() error {
//	if doc.ImageID == "" {
//		return nil
//	}
//	bucket, err := gridfs.NewBucket(
//		db.DB.Database,
//	)
//	imgId, _ := primitive.ObjectIDFromHex(doc.ImageID)
//	//fmt.Printf("Deleting gridfs id: %s\n", imgId)
//	if err != nil {
//		err = fmt.Errorf("Unable to get GridFS Bucket: %s", err)
//		log.Errorf("%s", err)
//		return err
//	}
//
//	err = bucket.Delete(imgId)
//	if err != nil {
//		fmt.Printf("Delete image %s failed: %s\n", imgId, err)
//	}
//	return nil
//}
//
//func (doc *dm.Document) WriteGridFs(imageData []byte) (string, error) {
//	startTime := time.Now()
//	//mdb := db.DB.Database
//	//fmt.Printf("Database: %s\n", spew.Sdump(mdb))
//
//	bucket, err := gridfs.NewBucket(
//		db.DB.Database,
//	)
//	if err != nil {
//		err = fmt.Errorf("Unable to get GridFS Bucket: %s", err)
//		log.Errorf("%s", err)
//		return "", err
//	}
//	metaData := make(map[string]string)
//	metaData["content_type"] = "pdf"
//	fileName := fmt.Sprintf("%s_%s_%s_%s", doc.Client, doc.Facility, doc.MRN, doc.ID.Hex())
//	saveImage, err := bucket.OpenUploadStream(
//		fileName,
//		options.GridFSUpload().SetMetadata(metaData),
//	)
//	if err != nil {
//		err = fmt.Errorf("OpenUploadStream failed: %s", err)
//		log.Errorln(err)
//		return "", err
//	}
//	defer saveImage.Close()
//
//	fileSize, err := saveImage.Write(imageData)
//	if err != nil {
//		err = fmt.Errorf("Save GridFS failed: %v", err)
//		log.Errorln(err)
//		return "", err
//	}
//	log.Debugf("Gridfs Saved %d bytes in %f seconds", fileSize, time.Since(startTime).Seconds())
//	imageIDHex := saveImage.FileID.(primitive.ObjectID).Hex()
//	return imageIDHex, nil
//}
//
//func (doc *pkg.Document) GetImage() (*[]byte, error) {
//	switch doc.ImageRepository {
//	case "GR":
//		return getGridFsImage(doc.ImageID)
//	case "CA":
//		return GetCAImage(doc.URL, doc.Facility, doc.ImageID)
//	}
//	return nil, fmt.Errorf("Invalid ImageRepository: [%s]\n", doc.ImageRepository)
//	//return GetImage(doc.ImageRepository, doc.ImageID)
//}
//
//func GetImage(repo string, imageID string) (*[]byte, error) {
//	switch repo {
//	case "GR":
//		return getGridFsImage(imageID)
//	}
//	return nil, fmt.Errorf("Invialid repository: %s", repo)
//}
//
//func getGridFsImage(imageID string) (*[]byte, error) {
//	db := db.DB.Database
//
//	// fsFiles := db.Collection("fs.files")
//	bucket, err := gridfs.NewBucket(
//		db,
//	)
//	if err != nil {
//		log.Errorf("New Bucket failed: %s", err)
//	}
//	var buf bytes.Buffer
//	//fmt.Printf("Looking for imageid: %s\n", imageID)
//	//oid, _ := primitive.ObjectIDFromHex("5ed406756cd109a290b17231")
//	oid, err := primitive.ObjectIDFromHex(imageID)
//	if err != nil {
//		err := fmt.Errorf("Document ImageID is not valid: %s  err: %s", imageID, err)
//		log.Errorln(err)
//		return nil, err
//	}
//	_, err = bucket.DownloadToStream(oid, &buf)
//	if err != nil {
//		log.Errorf("DownloadToStream failed: %s", err)
//	}
//	image := buf.Bytes()
//	//dStream, err := bucket.DownloadToStreamByName("Create_test_123456_5ed4003e030ea422d7210766", &buf)
//
//	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//	// var results bson.M
//	// fmt.Printf("Looking for imageid: %s\n", doc.ImageID)
//	// //iod, _ := primitive.ObjectIDFromHex(doc.ImageID)
//	// filter := bson.D{}
//	// err := fsFiles.FindOne(ctx, filter).Decode(&results)
//	// if err != nil {
//	// 	err := fmt.Errorf("GetGridFsImage did not find image: %s  Err: %s", filter, err)
//	// 	log.Errorln(err)
//	// 	return nil, err
//	// }
//	// fmt.Printf("Found Image: %s\n", spew.Sdump(results))
//	return &image, nil
//}
//
//func GetCAImage(url string, facility string, imageId string) (*[]byte, error) {
//	//TODO: Need to get the base image from the client table
//	baseURL := "http://localhost:4567/api/v1"
//	fullURL := baseURL + url
//	req, err := http.NewRequest("GET", fullURL, nil)
//	if err != nil {
//		return nil, err
//	}
//	req.Header.Set("AUTHORIZATION", "37")
//	req.Header.Set("facility", "demo")
//	client := &http.Client{Timeout: time.Second * 10}
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Fatal("Error reading response. ", err)
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatal("Error reading body. ", err)
//	}
//	imageFileName := fmt.Sprintf("./%s.pdf", imageId)
//	err = ioutil.WriteFile(imageFileName, body, 0644)
//	if err != nil {
//		fmt.Printf("Write to file error: %v\n", err)
//		return nil, err
//	}
//	return &body, nil
//}
//
//func logError(err error) error {
//	if err != nil {
//		log.Print(err)
//	}
//	return err
//}
