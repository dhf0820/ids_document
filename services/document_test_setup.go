package services

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/dhf0820/ids_model/common"

	"fmt"
	//"github.com/davecgh/go-spew/spew"
	//"context"
	//docPB "gitlab.com/dhf0820/ids_document_service/protobufs/docPB"
	//"net"
	//"testing"
	"os"
	//log "github.com/sirupsen/logrus"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/test/bufconn"
)

//https://stackoverflow.com/questions/42102496/testing-a-grpc-service
const bufSize = 1024 * 1024

//var bufLis *bufconn.Listener

const chunkSize = uint32(1 << 14)

var cfg *common.ServiceConfig

func InitTest() {
	fmt.Printf("\nInitTest\n\n")
	os.Setenv("CONFIG_ADDRESS", "http://localhost:19100/api/restv1/")
	os.Setenv("SERVICE_NAME", "document")
	os.Setenv("SERVICE_VERSION", "local_test")
	os.Setenv("COMPANY", "demo")
	_, err := Initialize("local_test", "demo")
	if err != nil {
		os.Exit(2)
	}
	//fmt.Printf("cfg: %s\n",spew.Sdump(cfg))
	cfg := GetConfig()
	fmt.Printf("cfg: %s\n", spew.Sdump(cfg))

	// bufLis = bufconn.Listen(bufSize)
	// s := grpc.NewServer()
	// rss := &DocumentServiceServer{}
	// docPB.RegisterDocumentServiceServer(s, rss)
	// go func() {
	// 	if err := s.Serve(bufLis); err != nil {
	// 		log.Fatalf("Server exited with error: %v", err)
	// 	}
	// }()
	//settings.SetDbName("test_release_service")
}

// func bufDialer(context.Context, string) (net.Conn, error) {
// 	return bufLis.Dial()
// }

/*func SetupDomainRelease(t *testing.T, create bool) *domain.Release {
	//settings.SetDbName("test_documents")
	data := sample.NewDomainDocument(1)
	//data.ID = primitive.NilObjectID
	// if create {
	// 	fmt.Printf("Creating new document: %s\n", spew.Sdump(data))
	// 	doc, err = domain.AddDocument(data)
	// 	if err != nil {
	// 		t.Fatalf("Error setupDomainDocument Creating: %v", err)
	// 	}
	// } else {
	// 	doc = data
	// }
	return data
}

func SetupPbCreateRelease(t *testing.T) *pb.CreateRelease {
	//settings.SetDbName("test_documents")
	data := sample.NewDocument(1)
	//imageFile := sample.ImageFileName
	//_, err := data.FromDocumentPB(pbDoc)
	// if err != nil {
	// 	err := fmt.Errorf("FromDocumentPB failed: %v", err)
	// 	t.Fatal(err)
	// }
	return data
}*/

//os.Setenv("ENV_CORE", "/Users/dhf/work/roi/services/core_service/config/config.json")
// service.InitCore("test")
// service.OpenDB()
//Start()
// fmt.Printf("\n\n---Start restful handler\n")
// restEp := service.GetMyEndpoint("restful_core")
// fmt.Printf("Restful EndPoint: %s\n", spew.Sdump(restEp))
// //restAddress := restEp.Address

// restAddress := fmt.Sprintf("%s:%s", restEp.Address, restEp.Port)
// router := h.NewRouter()
// logrus.Infof("----listening for restful requests at %s", restAddress)
// mainErr := http.ListenAndServe(restAddress, router)
// if mainErr != nil {
// 	logrus.Errorf("Rest Startup error: %v", mainErr)
// }
