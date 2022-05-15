package services

import (
	"bytes"
	//"context"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/smartystreets/goconvey/convey"

	//log "github.com/sirupsen/logrus"
	"testing"
)

type Upload struct {
	ID             string `json:"ID"`
	Size           int    `json:"Size"`
	SizeIsDeferred bool   `json:"SizeIsDeferred"`
	Offset         int
	MetaData       map[string]interface{}
	IsFinal        bool
	IsPartial      bool
	PartialUploads bool
}

// type HookData struct {
// 	Upload 	Upload 	`json:"Upload"`
// }
//func TestProcessJson(t *testing.T) {
//	//t.Parallel()
//	fmt.Printf("\n\nTestPocessJson\n")
//	Convey("unmarshal a json string", t, func() {
//		u := m.Upload{}
//		u.ID = "32f46e9f3052486a022c00548e3b27d1"
//		u.Size = 27181
//		u.SizeIsDeferred = true
//		hd := HookData{}
//		hd.Upload = u
//		hb, err := json.Marshal(hd)
//		ud := HookData{}
//		err = json.Unmarshal(hb, &ud)
//		fmt.Printf("ud := %s\n", hb)
//		fmt.Printf("um: %s\n", spew.Sdump(ud))
//
//		b := []byte("{\"Upload\":{\"ID\":\"32f46e9f3052486a022c00548e3b27d1\",\"Size\":27181,\"SizeIsDeferred\":true}}") //,"Offset":27181,"MetaData":{"Client":"test","Description":"Xray Left Knee"}}}"
//
//		fmt.Printf("sample: %s\n", b)
//		results := []byte("{\"Upload\":{\"ID\":\"32f46e9f3052486a022c00548e3b27d1\",\"Size\":27181,\"SizeIsDeferred\":false,\"Offset\":27181,\"MetaData\":{\"Client\":\"test\",\"Description\":\"Xray Left Knee\"}}}")
//
//		//os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
//		fmt.Printf("\n\n--- Convey\n")
//		hook := HookData{}
//		err = json.Unmarshal(results, &hook)
//		So(err, ShouldBeNil)
//		fmt.Printf("hook: %s\n", spew.Sdump(hook))
//	})
//}

//var r  {"Upload":{"ID":"32f46e9f3052486a022c00548e3b27d1","Size":27181,"SizeIsDeferred":false,"Offset":27181,"MetaData":{"Client":"test","Description":"Xray Left Knee"}}}

func TestUploadDoc(t *testing.T) {
	//t.Parallel()
	fmt.Printf("\n\nTestUploadDoc\n")
	Convey("Upload a document", t, func() {
		InitTest()
		OpenDB()

		//doc := dm.UploadDocument{}
		//ingress := dm.IngressDocument{}
		doc := MakeTestDocument()
		//maxId := 99999
		//minId := 10000
		//srcID := rand.Intn(maxId - minId) + minId
		//doc.SrcID = strconv.Itoa(srcID)
		//doc.DocDescription = "Mammography"
		//doc.DocClass = "Radiology"

		//////doc.TempImageID,_ = primitive.ObjectIDFromHex("61aa40dd995d52fdb964290c")
		//image, err := ioutil.ReadFile("./ClinDocMammo.pdf")
		//ingress.Document = doc
		//ingress.Image = image
		//doc.Image =  image
		body, err := json.Marshal(doc)
		if err != nil {
			So(err, ShouldBeNil)
		}
		resp, err := http.Post("http://localhost:29912/api/v1/document", "application/json",
			bytes.NewBuffer(body))
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeNil)
		//fmt.Printf("resp : %s\n", spew.Sdump(resp))
		//req, _ := http.NewRequest("POST", "http://localhost:29900/api/v1/release/27/document", bytes.NewBuffer(body))
	})
}


//func executeRequest(req *http.Request) *httptest.ResponseRecorder {
//	rr := httptest.NewRecorder()
//	http.Post
//	a.Router.ServeHTTP(rr, req)
//
//	return rr
//}

/*func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}*/