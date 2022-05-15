package services

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//"context"
	"encoding/json"
	"fmt"

	docMod "github.com/dhf0820/ids_model/document"

	//"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
	//m "github.com/dhf0820/ROIPrint/pkg/model"
	"io/ioutil"
	"net/http"
	//"github.com/oleiade/reflections"
)

//
//type DocumentResponse struct{
//	Status  int    `json:"status"`
//	Message string `json:"message"`
//	Document  *docMod.Document `json:"document"`
//
//}
//
//type ImageResponse struct{
//	Status  int    `json:"status"`
//	Message string `json:"message"`
//	Image  *[]byte `json:"image"`
//}

//type IngressDocument struct {
//	Document 	dm.UploadDocument 	`json:"document"`
//	Image 		[]byte 				`json:"image"`
//}

//############################## Response Writers ######################

func WriteDocumentResponse(w http.ResponseWriter, status int, resp *docMod.DocumentResponse) error {
	w.Header().Set("Content-Type", "application/json")

	switch status {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err

	}
	return nil
}

func WriteImageResponse(w http.ResponseWriter, status int, resp *docMod.DocumentResponse) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		err = fmt.Errorf("500|Error marshaling JSON: %v", err.Error())
		HandleError(w, "WriteDocumentResponse", err)
	}
	return nil
}

//############################  Route Handlers  ###########################
func addDocument(w http.ResponseWriter, r *http.Request) {
	log.Infof("addDocument Handler is called")
	//ingressDoc := dm.IngressDocument{}
	doc := docMod.Document{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("400|addDocument read body: %s", err.Error())
		log.Errorf("%s", err.Error())
		HandleError(w, "addDocument", err)
		return
	}
	//fmt.Printf("addDocument Body: %s\n", string(b))
	if err := json.Unmarshal(b, &doc); err != nil {
		err = fmt.Errorf("400|Error marshaling JSON: %v", err)
		log.Errorf("addDocument error: %s", err)
		HandleError(w, "addDocument", err)
		return
	}

	ndoc, err := ProcessDocument(&doc)
	if err != nil {
		err = fmt.Errorf("400|addDocument read body: %s", err.Error())
		log.Errorf("%s", err.Error())
		HandleError(w, "addDocument", err)
		return
	}
	resp := docMod.DocumentResponse{}
	resp.Status = 201
	resp.Document = ndoc
	resp.Message = "Added Document"

	//fmt.Printf("resp: %s\n", spew.Sdump(resp))
	WriteDocumentResponse(w, 201, &resp)

}

func getDocumentWithImage(w http.ResponseWriter, r *http.Request) {
	log.Infof("getDocument Handler is called")
	fmt.Printf("Image: %s\n", r.FormValue("image"))
	fmt.Printf("format: %s\n", r.FormValue("type"))

	params := mux.Vars(r)
	fmt.Printf("params: %v\n", params)
	id := params["doc_id"]
	//fmt.Printf("Received doc_id: %s\n", id)
	documentId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error on fromHex: %vn", err)
		HandleError(w, "getDocument", err)
		return
	}
	fmt.Printf("Received docId: %s\n", documentId)
	switch r.FormValue("image") {
	case "include": //Also return the image
		document, err := GetDocument(documentId)
		if err != nil {
			err = fmt.Errorf("400|getDocument read body: %s", err.Error())
			log.Errorf("%s", err.Error())
			HandleError(w, "getRelease", err)
			return
		}
		resp := docMod.DocumentResponse{
			Status:   200,
			Document: document,
			Message:  "Ok",
		}
		image, err := GetDocumentImage(documentId)
		if err != nil {
			log.Errorf("Image not retrieved")
			err := fmt.Errorf("Image %s not found: %s", documentId, err)
			resp.Message = err.Error()
			resp.Status = 400
		}
		resp.Image = image
		WriteDocumentResponse(w, resp.Status, &resp)
	case "only": // Return the image only
		resp := docMod.DocumentResponse{
			Status:  200,
			Message: "Ok",
		}
		image, err := GetDocumentImage(documentId)
		if err != nil {
			log.Errorf("Image not retrieved")
			err := fmt.Errorf("Image %s not found: %s", documentId, err)
			resp.Message = err.Error()
			resp.Status = 400
		}
		resp.Image = image
		WriteDocumentResponse(w, resp.Status, &resp)
	case "none": // No image
		fmt.Printf("none: No Image\n")
		document, err := GetDocument(documentId)
		if err != nil {
			err = fmt.Errorf("400|getDocument read body: %s", err.Error())
			fmt.Printf("GetDocument error: %s", err.Error())
			HandleError(w, "getDocument", err)
			return
		}
		//fmt.Printf("Received Document: %s\n", spew.Sdump(document))
		resp := docMod.DocumentResponse{
			Status:   200,
			Document: document,
			Message:  "Ok",
		}
		fmt.Printf("WriteDocumentResponse: %s\n", spew.Sdump(resp))
		WriteDocumentResponse(w, resp.Status, &resp)
	default:
		fmt.Printf("DEFAULT: No Image\n")
		document, err := GetDocument(documentId)
		if err != nil {
			err = fmt.Errorf("400|getDocument read body: %s", err.Error())
			fmt.Printf("GetDocument error: %s", err.Error())
			HandleError(w, "getDocument", err)
			return
		}
		resp := docMod.DocumentResponse{
			Status:   200,
			Document: document,
			Message:  "Ok",
		}
		WriteDocumentResponse(w, resp.Status, &resp)
	}

	//document, err := GetDocument(documentId)
	//if err != nil {
	//	err = fmt.Errorf("400|getDocument read body: %s", err.Error())
	//	log.Errorf("%s", err.Error())
	//	HandleError(w, "getRelease", err)
	//	return
	//}
	//resp := docMod.DocumentResponse{
	//	Status: 200,
	//	Document: document,
	//	Message: "Ok",
	//}
	//WriteDocumentResponse(w, resp.Status, &resp)
	//if query["image"] == "include" {
	//	image, err := GetDocumentImage(documentId)
	//	if err != nil {
	//		log.Errorf("Image not retrieved")
	//		err := fmt.Errorf("Image %s not found: %s", documentId, err )
	//		resp.Message = err.Error()
	//	}
	//	resp.Image = image
	//}

	//WriteDocumentResponse(w, 200, &resp)

}

func getImage(w http.ResponseWriter, r *http.Request) {
	log.Infof("getImage Handler is called")
	params := mux.Vars(r)
	fmt.Printf("params: %v\n", params)
	id := params["doc_id"]
	//fmt.Printf("Received doc_id: %s\n", id)
	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error on fromHex: %vn", err)
		HandleError(w, "getImage", err)
		return
	}
	fmt.Printf("Received docId: %s\n", docId)
	image, err := GetDocumentImage(docId)
	//if err != nil {
	//	err = fmt.Errorf("400|getDocument read body: %s", err.Error())
	//	log.Errorf("%s", err.Error())
	//	HandleError(w, "getRelease", err)
	//	return
	//}
	resp := docMod.DocumentResponse{}
	if err != nil {
		resp.Status = 400
		resp.Message = fmt.Sprintf("GetImage failed: %s", err.Error())
	} else {
		resp.Image = image
		resp.Status = 200
		resp.Message = "Ok"
	}

	//resp := ImageResponse{
	//	Image: image,
	//	Status:  200,
	//	Message: "Ok",
	//}
	WriteImageResponse(w, resp.Status, &resp)

}

//func processDocument(ingressDoc *dm.IngressDocument) (*dm.Document, error) {
//	newDoc := ingressDoc.Document
//	var checksum uint64
//	if len(ingressDoc.Image) > 0 {
//		crc64Table := crc64.MakeTable(crc64.ECMA)
//		checksum = crc64.Checksum(ingressDoc.Image, crc64Table)
//		metaData := make(map[string]string)
//		//metaData["mrn"] = udoc.MRN
//
//		//metaData["content_type"] = udoc.ImageType
//		metaData["checksum"] = strconv.FormatUint(checksum, 10)
//		//metaData["facility"] = udoc.Facility
//		//metaData["src_id"]	= udoc.SrcID
//		id, err := WriteGridFs(metaData, ingressDoc.Image)
//		if err != nil {
//			return nil, err
//		}
//		newDoc.ImageID = id
//		newDoc.CheckSum = checksum
//		newDoc.ImageSize = len(ingressDoc.Image)
//	}
//
//	ndoc, err := InsertDocument(context.Background(), newDoc)
//	if err != nil || ndoc == nil {
//		err := fmt.Errorf("Error Inserting Document: %v\n", err)
//		return nil, err
//	}
//
//	//fmt.Printf("\n\n\n$$ NewDoc ready for return: %s\n", spew.Sdump(ndoc))
//	return ndoc, err
//}
//fmt.Printf("image: %v\n",spew.Sdump(newUpload) )
// token := r.Header.Get("AUTHORIZATION")
// log.Debugf("Handler token: %s", token)
//fmt.Printf("\n\nCreating new recipient: %s\n", spew.Sdump(newRecip))

//_, err = http.Post("http://localhost:29900/api/v1/release/27/document", "application/json",
//	bytes.NewBuffer(body))
// err = usr.Create(token)
// if err != nil {
// 	log.Errorf("Create User failed: %s", err.Error())
// 	HandleError(w, "Create", err)
// 	return
// }

//Check if the recipient already exists. There should only be one.
// if there were more than one found return error and array of what was found. Otherwise return the exissting one.
// filter := m.SearchRecipientsValues{}
// filter.Source = newRecip.Client.Source
// filter.Customer = newRecip.Customer.Code
// filter.ForeignId = newRecip.Client.SourceId
// filter.Facility = newRecip.Customer.Facility
// ctx := context.Background() //TODO: need to set timeout on create
// recips, err := service.FindRecipients(ctx, &filter)
// if err == nil {
// 	if len(recips) > 0 {
// 		resp := SummaryResponse{}
// 		resp.Recipient = recips[0]
// 		resp.Status = 409
// 		resp.Message = "Already exists"
// 		WriteRecipientSummaryResponse(w, 200, &resp)
// 		return
// 	}
// }

// recip, err := service.CreateRecipient(ctx, &newRecip)
// resp := SummaryResponse{}
// if err != nil {
// 	if err.Error() != "Recipient Exists" {
// 		fmt.Printf("Create error: %v\n", err)
// 		respondWithError(w, 500, "err.Error()")
// 		return
// 	} else {
// 		resp.Message = "Recipient Exist "
// 	}
// } else {
// 	resp.Message = "Created"
// }
// resp.Recipient = recip
// WriteRecipientSummaryResponse(w, 200, &resp)
//}
