package services

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/dhf0820/ids_model/common"

	//coreModel "gitlab.com/dhf0820/ids_core_service/pkg"
	//"strings"
	//log "github.com/sirupsen/logrus"
	//"gitlab.com/dhf0820/ids_release_service/messaging"
	"net/http"
	"os"

	//mod "github.com/dhf0820/ids_model"

	//"strconv"
	//"time"
	//"os"
	"io/ioutil"
	//mod "github.com/dhf0820/ids_model"
	//"gitlab.com/dhf0820/ca_link_service/messaging"
)

/* DeliveryConfig is the basic configuration for this all services available to the current service
what it can talk to and who can talk to it*/
// type ServiceConfig struct {
// 	Name string `json:"name"`
// 	//Messaging  		Messaging					// move to endpoints
// 	DataConnector    *mod.DataConnector  `json:"dataconnector"`
// 	Services         []*mod.ServiceScope `json:"services" bson:"services"`
// 	MyEndPoints      []*mod.EndPoint     `json:"myendpoints"`
// 	ServiceEndPoints []mod.EndPoint     `json:"serviceendpoints" bson:"endpoints"`
// 	ConnectInfo      []*mod.ConnectInfo  `json:"connect_info" bson:"connect_info"`
// }

var (
	Conf *common.ServiceConfig //*config
	//MsgClients []*Messaging
	//*messaging.MessagingClient
	//GWConfig *m.DeliveryConfig
)

func GetServiceEndPoint(value string) *common.EndPoint {
	endPoints := GetConfig().ServiceEndPoints
	for _, ep := range endPoints {
		//fmt.Printf("Looking at %s for %s\n", ep.Name, value)
		if ep.Name == value {
			//fmt.Printf("Found EndPoint: %s\n", ep.Name)
			return ep
		}
	}
	return nil
}

func GetMyEndPoint(value string) *common.EndPoint {
	endPoints := GetConfig().MyEndPoints
	for _, ep := range endPoints {
		//fmt.Printf("Looking at %s for %s\n", ep.Name, value)
		if ep.Name == value {
			//fmt.Printf("Found EndPoint: %s\n", ep.Name)
			return ep
		}
	}
	return nil
}

//func GetMsgClient(name string) *messaging.MessagingClient {
//	for _, msg := range MsgClients {
//		if msg.Name == name {
//			return msg.Client
//		}
//	}
//	return nil
//}

//func InitMessaging(amqp *EndPoint) *messaging.MessagingClient {
//	//fmt.Printf("\n---Initializing %s\n\n", spew.Sdump(amqp))
//	msg := Messaging{}
//	msg.Name = amqp.Name
//	msg.VsAMQP = amqp.Address
//	msg.Client = &messaging.MessagingClient{}
//	msg.Client.ConnectToBroker(msg.VsAMQP)
//	MsgClients = append(MsgClients, &msg)
//	return msg.Client
//}
//TODO: Replace GetConfig in core with calls to core to get the centeralized information

func Initialize(serviceVersion, company string) (*common.ServiceConfig, error) {
	var err error
	//var delay time.Duration
	//var maxAttempts int
	fmt.Printf("\n\n\n -------Initiallizing DocumentService -----\n\n")
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceVersion == "" {
		serviceVersion = os.Getenv("SERVICE_VERSION")
	}
	if serviceName == "" {
		serviceName = "document"
		os.Setenv("SERVICE_NAME", "document")
	}

	if company == "" {
		company = os.Getenv("COMPANY")
	}
	if company == "" {
		company = "demo"
		os.Setenv("COMPANY", "demo")
	}

	if serviceVersion == "" {
		serviceVersion = "local_test"
		os.Setenv("VERSION", serviceVersion)
	}
	configAddress := os.Getenv("CONFIG_ADDRESS")
	fmt.Printf("Core address: %s\n", configAddress)
	fmt.Printf("waiting to retrieve the confguration from core\n")
	Conf, err = GetServiceConfig(serviceName, serviceVersion, company, "")
	if err != nil {
		return nil, err
	}
	//fmt.Printf("\n\n---GetServiceConfig returned:\n%s\n",  spew.Sdump(Conf))
	fmt.Printf("\n----config: %s]\n", spew.Sdump(Conf))
	OpenDB()
	return Conf, err
}

func GetConfig() *common.ServiceConfig {
	return Conf
}

type ConfigResp struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Config  common.ServiceConfig `json:"config"`
}

// Get the service configfrom core. Default is Restful, Optional is GRPC
func GetServiceConfig(name, version, company, mode string) (*common.ServiceConfig, error) {
	fmt.Printf("GetServiceConfig:140 version: %s\n", version)
	var cfg ConfigResp
	var err error
	//var unmarshalErr *json.UnmarshalTypeError
	var bdy []byte
	cfg = ConfigResp{}
	api := os.Getenv("CONFIG_ADDRESS")
	//coreName := strings.ReplaceAll(os.Getenv("CORE_NAME_PORT"), " ","")
	//api :=os.Getenv("API")
	if mode == "" || mode == "RESTFUL" {
		url := fmt.Sprintf("%s/config?name=%s&version=%s&company=%s", api, name, version, company)
		//url := fmt.Sprintf("http://localhost:19900/api/v1/config?name=%s&version=%s&company=%s", name, version, company)
		fmt.Printf("Config url: %s\n", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf(": %s  %v\n", url, err)
			return nil, err
		}
		defer resp.Body.Close()
		//cfg = coreModel.ServiceConfig{}
		bdy, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("raw string: %s\n",string(bdy))
		err = json.Unmarshal(bdy, &cfg)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("Config JSON: %s\n", spew.Sdump(cfg))

	} else {
		return nil, fmt.Errorf("%s is not supported yet", mode)
	}
	//fmt.Printf("CFG: = %s\n", spew.Sdump(cfg.Config))
	return &cfg.Config, err
}

//func getConf() *ServiceConfig {
//	V = viper.New()
//	os.Getenv("ROI_PREFIX")
//	V.SetEnvPrefix("roi")
//	V.AutomaticEnv()
//
//	V.SetConfigName("config.json")
//	V.AddConfigPath(os.Getenv("ENV_ROI")) // specifies where the location of the config file
//	//V.AddConfigPath("../config")
//	//V.AddConfigPath("./config")
//	//V.AddConfigPath("/etc/config")
//
//	V.SetConfigType("json")
//
//	err := V.ReadInConfig()
//
//	if err != nil {
//		fmt.Printf("%v", err)
//	}
//
//	Conf = &ServiceConfig{}
//	err = V.Unmarshal(Conf)
//	if err != nil {
//		fmt.Printf("unable to decode into config struct, %v", err)
//	}
//	fmt.Printf("\n\n--Core Conf: %s\n", spew.Sdump(Conf))
//
//	return Conf
//}
//
//func GetConfig() *ServiceConfig{
//	fmt.Printf("\n\n--GetConf: %s\n", spew.Sdump(Conf))
//	return Conf
//}

//func DeliveryConfigure() *DeliveryConfig {
//	//Config = fillConfig()
//	fmt.Printf("CaLinkConfigure called\n")
//	if Conf == nil {
//		fmt.Printf("\nConfiguring CaLink Service\n")
//		Conf = getConf()
//		Conf.Messaging.Client = &messaging.MessagingClient{}
//		//GWConfig.MessagingClient = &messaging.MessagingClient{}
//		if Conf.Messaging.Client == nil {
//			fmt.Printf("Messaging.Client is not configured\n")
//		}
//		fmt.Printf("AMQP: [%s]\n", Conf.Messaging.VsAMQP)
//		Conf.Messaging.Client.ConnectToBroker(Conf.Messaging.VsAMQP)
//	} else {
//		fmt.Printf("ChartArchive Link Service Is Currently configured: %s\n", spew.Sdump(GetDelivery()))
//	}
//	return Conf
//}
//
///*GetDelivery returns the active Delivery configation. THis includes CA Database, Messaging and list of services
//available.
//*/
//func GetDelivery() *DeliveryConfig {
//	return Conf
//}
//
//// //Config is the easy methog to access the configurations
//// func GetConfig() *DeliveryConfig {
//// 	return Conf
//// }

//var (
//	V *viper.Viper
//)

// Read the config file from the current directory and marshal
// into the conf config struct.
