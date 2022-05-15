package pkg

import (
	//m "github.com/dhf0820/ids_model"
	"time"

	cm "github.com/dhf0820/ids_model/common"
	docMod "github.com/dhf0820/ids_model/document"
	patientModel "github.com/dhf0820/ids_model/patient"
)

type DocumentResponse struct {
	Document *docMod.Document `json:"document"`
	Status   int              `json:"status"`
	Message  string           `json:"message"`
}

type CorrelationId struct {
	ReferenceType string `json:"reference_type" bson:"reference_type"` // source, Release, delivery Module in the correnlation (Release, Delivery)
	ReferenceID   string `json:"reference_id" bson:"reference_id"`     // ID of the source or id of ref type
	System        string `json:"system" bson:"system"`                 // Originating System (CA, Cerner,...)
	Version       string `json:"version" bson:"version"`               // Provided by source or blank
	//SystemID			string 			`json:"system_id" bson:"system_id"`  // the document/imageid in the sending system
	SystemImageURL    string     `json:"system_image_url" bson:"system_image_url"`       //System url to retrieve documen_image
	SystemDocumentURL string     `json:"system_document_url" bson:"system_document_url"` //System url to retrieve the document_metadata
	SystemFacility    string     `json:"system_facility" bson:"system_facility"`
	OriginatingIP     string     `json:"originating_ip" bson:"originating_ip"`
	Removable         string     `json:"removable" bson:"removable"`
	CreateTime        *time.Time `json:"create_time" bson:"create_time"`
}

type UploadDocument struct {
	SrcReleaseId   string               `json:"src_release_id"`
	SrcReleaseURL  string               `json:"src_release_url"`
	SrcID          string               `json:"src_id"`
	SrcImageURL    string               `json:"src_image_url"`
	SrcDocURL      string               `json:"src_doc_url"`
	DocClass       string               `json:"doc_class"`
	DocDescription string               `json:"doc_description"`
	DateOfService  string               `json:"date_of_service"`
	ImageType      string               `json:"image_type"`
	ImageSize      int                  `json:"image_size"`
	StorageType    string               `json:"storage_type"`
	Patient        patientModel.Patient `json:"patient"`
	Facility       string               `json:"facility"`
	Options        []*cm.Option         `json:"option:"`
	CheckSum       uint64               `json:"check_sum"`
	Image          []byte               `json:"image"`
}

//type IngressDocument struct {
//	CorrelationIDs    	[]*CorrelationId	`json:"correlation_ids" bson:"correlation_ids"`
//	//ImageID  			primitive.ObjectID	`json:"image_id omitempty" bson:"image_id omitempty"` //GridFS Image ID
//	//TempImageID 		primitive.ObjectID	`json:"temp_image_id omitempty" bson:"temp_image_id omitempty"`
//	CheckSum			string				`json:"check_sum" bson:"check_sum"`
//	//ImageRepository   	string				`json:"image_repository" bson:"image_repository"`
//	StorageType       	string				`json:"storage_type" bson:"storage_type"` // temp/perm
//	ImageType			string 				`json:"image_type" bson:"image_type"`
//	ImageSize 			int64 				`json:"image_size" bson:"image_size"`
//}
type IngressDocument struct {
	Document *docMod.Document `json:"document"`
	Image    []byte           `json:"image"`
}

//type Document struct {
//	ID                	primitive.ObjectID	`json:"id,omitempty" bson:"_id,omitempty"` // Database Id for this document
//	CorrelationIDs    	[]*CorrelationId	`json:"correlation_ids" bson:"correlation_ids"`
//	ImageID  			primitive.ObjectID	`json:"image_id omitempty" bson:"image_id omitempty"` //GridFS Image ID
//	TempImageID 		primitive.ObjectID	`json:"temp_image_id omitempty" bson:"temp_image_id omitempty"`
//	CheckSum			uint64				`json:"check_sum" bson:"check_sum"`
//	ImageRepository   	string				`json:"image_repository" bson:"image_repository"`
//	StorageType       	string				`json:"storage_type" bson:"storage_type"` // temp/perm
//	ImageType			string 				`json:"image_type" bson:"image_type"`
//	ImageSize 			int 				`json:"image_size" bson:"image_size"`
//	CreatedAt         	*time.Time			`json:"created_at" bson:"created_at"`
//	UpdatedAt         	*time.Time			`json:"updated_at" bson:"updated_at"`
//	DeletedAt         	*time.Time			`json:"deleted_at" bson:"deleted_at"`
//}

type DocumentSearchFilter struct {
	Limit         int32  `json:"limit" schema:"limit"`
	OffSet        int32  `json:"offset" schema:"offset"`
	Order         string `json:"order" schema:"order"`
	SortBy        string `json:"sortby" schema:"sortby"` // string of names normal =increasing, leading - for descending
	ReferenceType string `json:"reference_type" schema:"reference_type"`
	ReferenceID   string `json:"reference_id" schema:"reference_id"` // When used must be paired with Ref Type
	Version       string `json:"version" schema:"version"`           // When used must be paired with RefType and ID
	CheckSum      uint64 `json:"check_sum" schema:"check_sum"`
	ImageID       string `json:"image_id" schema:"image_id"` // IDS actual Image ID
	ID            string `json:"id" schema:"id"`             // primitive ObjectId as hex string Unique index
}

//ImageInfo  Page, ImageSize, ImageType, CheckSum
//StorageInfo ImageRepo, storagetype

// type DocumentMetadata struct {
// 	SrcDocURL			string 						 		 `json:"src_doc_url" bson:"src_doc_url"`
// 	SrcImageURL    string						 		 `json:"src_image_url" bson:"src_image_url"`
// 	ImageURL
// 	ImageID
// 	Status
// 	DocClass          string             `json:"doc_class" bson:"doc_class"`
// 	Description       string             `json:"description" bson:"description"`
// 	ImageType         string             `json:"image_type" bson:"image_type"`
// 	DateOfService     string             `json:"date_of_service" bson:"date_of_service"`

// }

// type DbDocument struct {
// 	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` // Database Id for this document
// 	Client          Client             `json:"client"`
// 	Customer        Customer           `josn:"customer"`
// 	ReleaseID       string             `json:"release_id,omitempty"`       // Release this documents belongs to.
// 	ImageRepository string             `json:"image_repository,omitempty"` //What is the repository for the image (gridfs,)
// 	ImageID         string             `json:"image_id,omitempty"`         // Use ImageRepository to locate the image
// 	ImageType       string             `json:"image_type,omitempty"`       // Type of image uploaded. pdf, fhir
// 	SourceDocID     string             `son:"source_doc_id,omitempty"`     // source document id
// 	MRN             string             `json:"mrn,omitempty"`              // MRN of the patient
// 	Facility        string             `json:"facility,omitempty"`         // facility generating the document
// 	Source          string             `json:"source"`
// 	DocClass        string             `json:"doc_class,omitempty"` // Document class of image. "ROI-Auth" for the authorization
// 	// "ROI-Cover" for any cover sheet. ROI-Index
// 	Description   string     `json:"description,omitempty"` // Description provided by facility
// 	URL           string     `json:"url,omitempty"`         // URL of where to retrieve the document
// 	CreatedAt     *time.Time `json:"created_at" bson:"created_at"`
// 	UpdatedAt     *time.Time `json:"updated_at" bson:"updated_at"`
// 	DeletedAt     *time.Time `json:"deletedAt" bson:"deleted_at"`
// 	DateOfService string     `json:"date_of_service"`
// }

// type FullDocument struct {
// 	Id            string   `json:"id" ` // Original Document Id
// 	Client        Client   `json:"client"`
// 	Customer      Customer `json:"customer"`
// 	DocClass      string   `json:"doc_class"`
// 	Description   string   `json:"description"`
// 	DateOfService string
// 	ImageType     string
// 	ReleaseId     string
// 	URL           string
// 	DocumentId    string // DocumentService id of the master document
// }

type DeliveryDocumentFile struct {
	ImageRepository string `json:"image_repository" bson:"image_repository"`
	ImageType       string `json:"image_type" bson:"image_type"`
	ImageURL        string `json:"image_url" bson:"image_url"`
	DocumentURL     string `json:"document_url" bson:"document_url"`
	DocClass        string `json:"doc_class" bson:"doc_class"`
	Description     string `json:"description" bson:"description"`
	Version         string `json:"version" bson:"version"`
	DateOfService   string `json:"date_of_service" bson:"date_of_service"`
	FileName        string `json:"file_name" bson:"file_name"`
}

type DeliveryDocumentImage struct {
	ImageRepository string `json:"image_repository" bson:"image_repository"`
	ImageType       string `json:"image_type" bson:"image_type"`
	ImageURL        string `json:"image_url" bson:"image_url"`
	DocumentURL     string `json:"document_url" bson:"document_url"`
	DocClass        string `json:"doc_class" bson:"doc_class"`
	Description     string `json:"description" bson:"description"`
	Version         string `json:"version" bson:"version"`
	DateOfService   string `json:"date_of_service" bson:"date_of_service"`
	Image           []byte `json:"file_name" bson:"file_name"`
}
