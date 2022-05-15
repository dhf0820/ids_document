package services

import (
	"fmt"

	//"github.com/davecgh/go-spew/spew"

	//docMod "github.com/dhf0820/ids_model/document"

	//"github.com/davecgh/go-spew/spew"
	. "github.com/smartystreets/goconvey/convey"
	//deliv "gitlab.com/dhf0820/delivery_service/connect_deliver"
	//pb "gitlab.com/dhf0820/delivery_service/protobufs/delPB"
	"testing"
)

func TestMergeDocument(t *testing.T) {
	t.Parallel()
	fmt.Printf("----TestMergeDocument\n")
	InitTest()
	Convey("GetDocument via Restful", t, func() {
		infiles := []string{}
		doc1 := "./debbie.pdf"
		doc2 := "./harman.pdf"
		infiles = append(infiles, doc1, doc2)

		err := MergeToCombined(infiles, "./merged.pdf", false)
		So(err, ShouldBeNil)
	})
}
