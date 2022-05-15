package services

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	Route{
		"HealthCheck",
		"GET",
		"/api/rest/v1/healthcheck",
		HealthCheck,
	},
	Route{
		"GetDocument",
		"GET",
		"/api/rest/v1/document/{doc_id}",
		getDocumentWithImage,
	},
	//Route{
	//	"GetDocumentImage",
	//	"GET",
	//	"/api/v1/document/{doc_id}?with_image=true",
	//	getDocument_with_image,
	//},
	Route{
		"GetImage",
		"GET",
		"/api/rest/v1/image/{doc_id}",
		getImage,
	},
	// Route{
	// 	"TestForm",
	// 	"POST",
	// 	"/api/v1/test",
	// 	testForm,
	// },
	// Route{
	// 	"CreateRecipient",
	// 	"POST",
	// 	"/api/v1/recipient",
	// 	createRecipient,
	// },
	// Route{
	// 	"GetConfig",
	// 	"GET",
	// 	"/api/v1/config",
	// 	getConfig,
	// },
	// Route{
	// 	"GetRecipient",
	// 	"GET",
	// 	"/api/v1/recipient/{recipient_id}",
	// 	getRecipient,
	// },
	// Route{
	// 	"FindRecipient",
	// 	"GET",
	// 	"/api/v1/recipient",
	// 	findRecipients,
	// },
	// Route{
	// 	"FindRecipientByForeign",
	// 	"GET",
	// 	"/api/v1/recipient_foreign_id",
	// 	getRecipientByForeignId,
	// },
	// Route{
	// 	"Admin",
	// 	"PUT",
	// 	"/api/v1/admin",
	// 	UpdateEnv,
	// },
	// Route{
	// 	"CreateRelease",
	// 	"POST",
	// 	"/api/v1/release",
	// 	createRelease,
	// },
	// Route{
	// 	"UploadDocument",
	// 	"POST",
	// 	"/api/v1/document/{document_id}",
	// 	receiveDocument,
	// },
	Route{
		"AddDocument",
		"POST",
		"/api/rest/v1/document",
		addDocument,
	},
	// Route{
	// 	"SubmitDelivery",
	// 	"POST",
	// 	"/api/v1/release/deliver",
	// 	submit,
	// },
	// Route{
	// 	"CreateDevice",
	// 	"POST",
	// 	"/api/v1/device/{recipient_id}",
	// 	createDevice,
	// },
	// Route{
	// 	"GetDevice",
	// 	"GET",
	// 	"/api/v1/device",
	// 	getDevice,
	// },
	// Route{
	// 	"GetConnectorConfig",
	// 	"GET",
	// 	"/api/v1/connector_config",
	// 	getConnectorConf,
	// },

	//Route{
	//	"SetLogLevel",
	//	"GET",
	//	"/api/v1/log_level/{level}",
	//	GetLogLevel,
	//},
	//Route{
	//	"SetLogLevel",
	//	"PUT",
	//	"/api/v1/log_level/{level}",
	//	SetLogLevel,
	//},
}
