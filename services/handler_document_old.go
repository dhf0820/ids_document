package services

//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//
//	"github.com/gorilla/mux"
//	//"gitlab.com/dhf0820/ids_core_service/service"
//
//	"github.com/davecgh/go-spew/spew"
//	log "github.com/sirupsen/logrus"
//	dm "github.com/dhf0820/ids_document/pkg"
//	m "github.com/dhf0820/ids_model"
//	//relConn "gitlab.com/dhf0820/ids_release_service/connect"
//	//rpb "gitlab.com/dhf0820/ids_release_service/protobufs/relPB"
//	//"gitlab.com/dhf0820/idsCore/pkg/release"
//)
//
//type ReleaseResponse struct {
//	Status  int       `json:"status"`
//	Message string    `json:"message"`
//	Release m.Release `json:"release"`
//}
//
//func WriteDocumentResponse(w http.ResponseWriter, status int, resp *DocumentResponse) error {
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(status)
//	err := json.NewEncoder(w).Encode(resp)
//	if err != nil {
//		err = fmt.Errorf("500|Error marshaling JSON: %v", err.Error())
//		HandleError(w, "WriteDocumentResponse", err)
//	}
//	return nil
//}
//
//type DocumentResponse struct {
//	Status  int       		`json:"status"`
//	Message string    		`json:"message"`
//	Document dm.Document		`json:"document"`
//}
//
//type DocumentsResponse struct {
//	Status  int       `json:"status"`
//	Message string    `json:"message"`
//	Documents []dm.Document `json:"documents"`
//}
//
//func createDocument(w http.ResponseWriter, r *http.Request) {
//	log.Infof("Add Document Handler called")
//	params := mux.Vars(r)
//	fmt.Printf("params: %v\n", params)
//	// fmt.Printf("ReleaseId: %v\n", params["release_id"])
//	// relId := params["release_id"]
//	// fmt.Printf("Set releaseId: %s\n", relId)
//	doc := dm.Document{}
//	b, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		err = fmt.Errorf("400|addDocument read body: %s", err.Error())
//		log.Errorf("%s", err.Error())
//		HandleError(w, "addDocument", err)
//		return
//	}
//	fmt.Printf("Document Body: %s\n", string(b))
//	if err := json.Unmarshal(b, &doc); err != nil {
//		err = fmt.Errorf("400|Error marshaling JSON: %s", err.Error())
//		log.Errorf("addDocument error: %s", err)
//		HandleError(w, "addDocument", err)
//		return
//	}
//	fmt.Printf("Adding new document: %s\n", spew.Sdump(doc))
//
//	// err = usr.Create(token)
//	// if err != nil {
//	// 	log.Errorf("Create User failed: %s", err.Error())
//	// 	HandleError(w, "Create", err)
//	// 	return
//	// }
//	ctx := context.Background() //TODO: need to set timeout on create
//	newDoc, err := InsertDocument(ctx,  &doc)
//	fmt.Printf("Service.AddDocument returned\n")
//	if err != nil {
//		fmt.Printf("\n--Create error: %v\n", err)
//		respondWithError(w, 500, err.Error())
//		return
//	}
//	resp := DocumentResponse{}
//	resp.Document = *newDoc
//	resp.Message = "Added"
//	WriteDocumentResponse(w, 200, &resp)
//}
