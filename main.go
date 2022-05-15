package main

import (
	"fmt"

	//"github.com/joho/godotenv"

	"gitlab.com/dhf0820/ids_document_service/services"

	log "github.com/sirupsen/logrus"
	// mod "github.com/dhf0820/ids_model"
	// "gitlab.com/dhf0820/ids_core_service/protobufs/corePB"
	// "google.golang.org/grpc/reflection"
	//"strings"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials"
	//"time"
)

func main() {
	log.Infof("document starting: t220104.0\n")
	//godotenv.Load(".env.document")
	//services.Initialize("", "")
	//conf := services.GetConfig()
	//fmt.Printf("Config: %s\n", spew.Sdump(conf))
	//services.OpenDB()
	services.Start()
	fmt.Printf("Start Returned and should not have\n")

	// 	fmt.Printf("\n\n---Start restful handler\n")
	// 	restEp := services.GetMyEndpoint("restful_core")
	// 	fmt.Printf("Restful EndPoint: %s\n", spew.Sdump(restEp))
	// 	//restAddress := restEp.Address

	// 	restAddress := fmt.Sprintf("%s:%s", restEp.Address, restEp.Port)
	// 	router := NewRouter()
	// 	log.Infof("----listening for restful requests at %s", restAddress)
	// 	mainErr := http.ListenAndServe(restAddress, router)
	// 	if mainErr != nil {
	// 		logrus.Errorf("Rest Startup error: %v", mainErr)
	// 	}
	// }
}

// func Start() {
// 	run_env := os.Getenv("DOCUMENT_VERSION")
// 	fmt.Printf("Start Called: [%s]\n", run_env)

// 	fmt.Printf("32<<20 = %d   32<<24 = %d\n", 32<<20, 32<<22)
// 	//service.InitCore(run_env) //TODO: get the env value from flag
// 	//OpenDB()
// 	cfg := service.GetConfig()
// 	fmt.Printf("\n---cfg: %s\n", spew.Sdump(cfg))
// 	//ep := GetMyEndpoint("core")
// 	eps := GetMyEndpoints()

// 	for _, ep := range eps {
// 				fmt.Printf("Restful EndPoint: %s\n", spew.Sdump(ep))
// 				//restAddress := restEp.Address

// 				restAddress := fmt.Sprintf("%s:%s", ep.Address, ep.Port)
// 				router := h.NewRouter()
// 				log.Infof("listening for restful requests at %s", restAddress)
// 				mainErr := http.ListenAndServe(restAddress, router)
// 				if mainErr != nil {
// 					log.Errorf("Rest Startup error: %v", mainErr)
// 				}
// 		}
// 	}
// }
