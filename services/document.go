package services

import (
	"context"
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"strings"
	"time"

	docMod "github.com/dhf0820/ids_model/document"

	//"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	//common "github.com/dhf0820/ids_model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"google.golang.org/protobuf/internal/errors"
)

//ProcessIngress: accepts the IngressDocument and saves the image, if any, and then saves the document
func ProcessDocument(doc *docMod.Document) (*docMod.Document, error) {

	//var checksum uint64
	//if len(ingressDoc.Image) > 0 {
	//	crc64Table := crc64.MakeTable(crc64.ECMA)
	//	checksum = crc64.Checksum(ingressDoc.Image, crc64Table)
	//	metaData := make(map[string]string)
	//	//metaData["mrn"] = udoc.MRN
	//
	//	//metaData["content_type"] = udoc.ImageType
	//	metaData["checksum"] = strconv.FormatUint(checksum, 10)
	//	//metaData["facility"] = udoc.Facility
	//	//metaData["src_id"]	= udoc.SrcID
	//
	//	fmt.Printf("Call GridFsWrite\n")
	//	id, err := WriteGridFs(metaData, ingressDoc.Image)
	//	if err != nil {
	//
	//		fmt.Printf("ProcessIngress wrote the gridfs err: %v\n", err)
	//		return nil, err
	//	}
	//	newDoc.ImageID = id
	//	newDoc.CheckSum = checksum
	//	newDoc.ImageSize = len(ingressDoc.Image)
	//
	//}
	//TODO: Check lookup up checksum of the image in documents if duplicate
	// If duplicate use the original image in this document
	// delete the new one received.
	fmt.Printf("ImageID: %s\n", doc.ImageID)
	fmt.Printf("CheckSum: %d\n", doc.CheckSum)
	fmt.Printf("ImageSize: %d\n", doc.ImageSize)

	doc.ImageID = doc.TempImageID
	doc.TempImageID = primitive.NilObjectID
	doc.ImageRepository = "GR"
	doc.StorageType = "perm"
	//doc.ImageType = "pdf"

	ndoc, err := InsertDocument(context.Background(), doc)
	if err != nil || ndoc == nil {
		err := fmt.Errorf("Error Inserting Document: %v\n", err)
		return nil, err
	}

	fmt.Printf("\n\n\n### NewDoc ready for return: %s\n", spew.Sdump(ndoc))
	return ndoc, err
}

//InsertDocument accepts a document and inserts it Required fields are the Correlation.refs
// storage type temp/perm. Optonal is ref_type of "release" and the ref_id the release id.
func InsertDocument(ctx context.Context, new_doc *docMod.Document) (*docMod.Document, error) {
	var err error
	//fmt.Printf("document: %s\n", spew.Sdump(new_doc))
	//fmt.Printf("src correlation: %s\n", spew.Sdump(new_doc.CorrelationIDs[0]))
	filter := docMod.DocumentSearchFilter{}
	filter.ReferenceType = "source"
	filter.ReferenceID = new_doc.CorrelationIDs[0].ReferenceID
	//TODO:Possibly add checkSum to filter checking document exists
	//fmt.Printf("Calling DocumentExists with %v\n", filter)
	//doc, err := DocumentExists(ctx, &filter)

	//:TODO: Remove allowing duplicate documents
	//if err != nil { // Some error other than document does not exist (It shouldn't)
	//	fmt.Printf("Document already exists error: %v\n", err)
	//	return nil, err
	//}

	//fmt.Printf("Checking if doc is nil\n")
	//if doc != nil {
	//	log.Warnf("Document for %s exists", spew.Sdump(doc))
	//	return doc, err
	//}
	//fmt.Printf("Document is new, %s\n", spew.Sdump(doc))
	collection, err := GetCollection("documents")
	if err != nil {
		//TODO Process Error
		fmt.Printf("GetCollection document failed: %v\n", err)
		return nil, err
	}
	timeNow := time.Now()
	new_doc.CreatedAt = &timeNow
	new_doc.UpdatedAt = &timeNow
	//new_doc.ID = primitive.NewObjectID()
	//fmt.Printf("Inserting document: %s\n", spew.Sdump(new_doc))
	resp, err := collection.InsertOne(ctx, new_doc)

	if err != nil {
		log.Errorf("InsertOne document failed: %v\n", err)
		return nil, err
	}
	new_doc.ID = resp.InsertedID.(primitive.ObjectID)
	return new_doc, nil
}

func GetDocument(docId primitive.ObjectID) (*docMod.Document, error) {
	var err error
	collection, err := GetCollection("documents")
	if err != nil {
		//TODO Process Error
		fmt.Printf("GetCollection document failed: %v\n", err)
		return nil, err
	}
	filter := bson.M{"_id": docId}
	singleResult := collection.FindOne(context.Background(), filter)
	if singleResult.Err() != nil {
		return nil, singleResult.Err()
	}
	doc := docMod.Document{}
	err = singleResult.Decode(&doc)
	if err != nil {
		return nil, fmt.Errorf("Decode error: %s", err.Error())
	}
	return &doc, err
}

func GetDocumentImage(docId primitive.ObjectID) (*[]byte, error) {
	doc, err := GetDocument(docId)
	if err != nil {
		fmt.Printf("GetDocumentImage error: %s\n", err.Error())
		return nil, err
	}
	//fmt.Printf("DocumentImage: %s\n", spew.Sdump(doc))
	imageId := doc.ImageID

	image, err := getGridFsImage(imageId)
	fmt.Printf("Length of image: %d\n", len(*image))
	if err != nil {
		return nil, fmt.Errorf("GetImage error: %s", err.Error())
	}
	return image, nil
}

func DocumentExists(ctx context.Context, filter *docMod.DocumentSearchFilter) (*docMod.Document, error) {
	var err error
	collection, err := GetCollection("documents")
	if err != nil {
		//TODO Process Error
		fmt.Printf("GetCollection document failed: %v\n", err)
		return nil, err
	}

	u := bson.D{}
	fieldsInQuery := 0

	if filter.ReferenceType != "" {
		if filter.ReferenceID == "" {
			err := fmt.Errorf("ReferenceType requires a ReferenceID")
			return nil, err
		}
		u = bson.D{{"correlation_ids", bson.D{{"$elemMatch",
			bson.D{{Key: "reference_type", Value: strings.ToLower(filter.ReferenceType)},
				{Key: "reference_id", Value: filter.ReferenceID}}}}}}

		// e := bson.E{Key: "reference_type", Value: filter.ReferenceType}
		// u = append(u, e)
		// e = bson.E{Key: "reference_id", Value: filter.ReferenceID}
		//u = append(u, d)
		fieldsInQuery = 1
	}
	if fieldsInQuery == 0 {
		fmt.Printf("\n\n FindDocuments Search params are empty: %v\n\n", u)
		return nil, fmt.Errorf("search is empty")
	}
	fmt.Printf("DocumentExists Calling Find with %v\n", u)
	cur, err := collection.Find(context.Background(), u)

	if err != nil {
		//fmt.Printf("Find returned %v\n", err)
		return nil, nil
	}
	i := 0
	results := []*docMod.Document{}
	fmt.Printf("Process cursor length: %d\n", cur.RemainingBatchLength())
	if cur.RemainingBatchLength() > int(0) {
		for cur.Next(context.Background()) {
			i++
			d := docMod.Document{}
			fmt.Printf("Decoding\n")
			err := cur.Decode(&d)
			fmt.Printf("Decode finished\n")
			if err != nil {
				fmt.Printf("Find Document decode failed: %v", err)
				return nil, fmt.Errorf("find Document decode failed: %v", err)
			}
			//fmt.Printf("Appending document %s\n", spew.Sdump(d))
			results = append(results, &d)
		}
	} else {
		//fmt.Printf("Skipping cursor, returning nil, nil\n")
		//no existing document
		return nil, nil
	}
	if i == 0 {
		return nil, nil
	}
	//fmt.Printf("Length of results: %d\n", len(results))
	return results[0], fmt.Errorf("Document with same source, facility, and documentID exists")
}

//FindDocuments accepts a filter and returns an array of all the documents matching the filter.
// If a filter is put to source and a valid id, only the one would be returned.
// If the filter was set to release and a number, a list of all the documents for that release would be returned.
func FindDocuments(ctx context.Context, filter *docMod.DocumentSearchFilter) ([]*docMod.Document, error) {
	var err error
	collection, err := GetCollection("documents")
	if err != nil {
		//TODO Process Error
		fmt.Printf("GetCollection document failed: %v\n", err)
		return nil, err
	}

	u := bson.D{}
	fieldsInQuery := 0

	if filter.ReferenceType != "" {
		if filter.ReferenceID == "" {
			err := fmt.Errorf("ReferenceType requires a ReferenceID")
			return nil, err
		}
		u = bson.D{{"correlation_ids", bson.D{{"$elemMatch",
			bson.D{{Key: "reference_type", Value: strings.ToLower(filter.ReferenceType)},
				{Key: "reference_id", Value: filter.ReferenceID}}}}}}

		// e := bson.E{Key: "reference_type", Value: filter.ReferenceType}
		// u = append(u, e)
		// e = bson.E{Key: "reference_id", Value: filter.ReferenceID}
		//u = append(u, d)
		fieldsInQuery = 1
	}

	if fieldsInQuery == 0 {
		fmt.Printf("\n\n FindDocuments Search params are empty: %v\n\n", u)
		return nil, fmt.Errorf("search is empty")
	}
	fmt.Printf("FindRecipients Calling Find with %v\n", u)
	cur, err := collection.Find(context.Background(), u)
	if err != nil {
		return nil, nil
	}
	i := 0
	results := []*docMod.Document{}
	for cur.Next(context.Background()) {
		i++
		d := docMod.Document{}
		err := cur.Decode(&d)
		if err != nil {
			log.Errorf("Find Document decode failed: %v", err)
			return nil, fmt.Errorf("find Document decode failed: %v", err)
		}
		results = append(results, &d)
	}
	if i == 0 {
		return results, nil
	}
	return results, nil
}

//InsertDocument accepts a document and inserts it Required fields are the Correlation.refs
// storage type temp/perm. Optonal is ref_type of "release" and the ref_id the release id.
// func FindDocument(ctx context.Context, doc *m.DocumentSearchFilter) (*m.Document, error) {
// 	var err error
// 	collection, err := GetCollection("")
// 	if err != nil {
// 		//TODO Process Error
// 		fmt.Printf("GetCollection document failed: %v\n", err)
// 		return nil, err
// 	}

// 	u := bson.D{}

// 	resp, err := collection.InsertOne(ctx, doc)
// 	if err != nil {
// 		log.Errorf("InsertOne document failed: %v\n", err)
// 		return nil, err
// 	}
// 	fmt.Printf("inserted a new document: %s", spew.Sdump(resp))
// 	return doc, nil
// }
