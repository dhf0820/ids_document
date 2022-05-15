package services

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	common "github.com/dhf0820/ids_model/common"
	log "github.com/sirupsen/logrus"

	//common "gitlab.com/dhf0820/ids_core_service/pkg"
	//"github.com/joho/godotenv"
	//"github.com/dhf0820/ids_document"
	"net/http"
	"os"
	"sync"
	// mod "github.com/dhf0820/ids_model"
	//"net"
	// "gitlab.com/dhf0820/ids_core_service/protobufs/corePB"
	// "google.golang.org/grpc/reflection"
	// "strings"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials"
)

func Start() {
	run_env := os.Getenv("SERVICE_VERSION")
	company := os.Getenv("COMPANY")
	fmt.Printf("Start Called: [%s]\n", run_env)
	Initialize(run_env, company)
	//service.InitCore(run_env) //TODO: get the env value from flag
	//OpenDB()
	//cfg := GetConfig()
	//fmt.Printf("\n---cfg: %s\n", spew.Sdump(cfg))
	//ep := mod.GetMyEndpoint("core")
	GetConfig()
	eps := GetMyEndPoints()

	var wg sync.WaitGroup
	var restAddress string
	for _, ep := range eps {
		if ep.Protocol == "http" {
			fmt.Printf("Start Restful listener\n")
			defer wg.Done()
			wg.Add(1)

			fmt.Printf("Starting Restful EndPoint: %s\n", spew.Sdump(ep))
			restAddress = fmt.Sprintf("%s:%s", "0.0.0.0", ep.Port)
			go restful_worker(&wg, restAddress)
		} else {
			fmt.Printf("connection %s found. Not implemented\n", ep.Name)
		}
	}
	fmt.Printf("Wait for restful listeners: %s\n", restAddress)
	wg.Wait()
	fmt.Printf("Restful Servers stopping\n")
}

func restful_worker(wg *sync.WaitGroup, restAddress string) {
	//defer wg.Done()
	fmt.Printf("Restful Worker:58 -- Restful Address: %s\n", restAddress)
	//restAddress := restEp.Address
	//restAddress := fmt.Sprintf("%s:%s", "0.0.0.0", ep.Port)
	//restAddress := fmt.Sprintf("%s:%s", "0.0.0.0", "29912")
	router := NewRouter()
	log.Infof("Document starting listening for restful requests at %s", restAddress)
	http.ListenAndServe(restAddress, router)
	// if mainErr != nil {
	// 	log.Errorf("Rest Startup error: %v", mainErr)
	// }
}

func GetMyEndPoints() []*common.EndPoint {
	fmt.Printf("Config: %s\n\n", spew.Sdump(GetConfig()))
	endPoints := GetConfig().MyEndPoints
	//fmt.Printf("Release Endpoints: %s", spew.Sdump(endPoints))
	return endPoints
}
