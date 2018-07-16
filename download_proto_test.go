package versions

import (
	// "fmt"
	"testing"
	"time"

	pb "github.com/stevepartridge/versions/protos"
)

var (
	testDownloadProto1 = &pb.Download{
		Id:        int32(testDownload1.Id),
		Filename:  testDownload1.Filename,
		UpdatedAt: testDownload1.UpdatedAt.Format(time.RFC3339Nano),
		CreatedAt: testDownload1.CreatedAt.Format(time.RFC3339Nano),
	}
)

func Test_Unit_DownloadToProtoMatch(t *testing.T) {

	proto := testDownload1.ToProto()

	if proto.Id != testDownloadProto1.Id {
		t.Errorf("Proto id should have matched: %d", testDownloadProto1.Id)
	}

	if proto.Filename != testDownloadProto1.Filename {
		t.Errorf("Proto name should have matched: %s", testDownloadProto1.Filename)
	}

	if proto.UpdatedAt != testDownloadProto1.UpdatedAt {
		t.Errorf("Proto updated at %s should have matched: %s",
			proto.UpdatedAt,
			testDownloadProto1.UpdatedAt,
		)
	}

	if proto.CreatedAt != testDownloadProto1.CreatedAt {
		t.Errorf("Proto created at %s should have matched: %s",
			proto.CreatedAt,
			testDownloadProto1.CreatedAt,
		)
	}
}

func Test_Unit_DownloadFromProtoMatch(t *testing.T) {

	dl := DownloadFromProto(testDownloadProto1)
	if dl.Id != int(testDownloadProto1.Id) {
		t.Errorf("Download id should have matched: %d", testDownloadProto1.Id)
	}

	if dl.Filename != testDownloadProto1.Filename {
		t.Errorf("Download name should have matched: %s", testDownloadProto1.Filename)
	}

	if dl.UpdatedAt.Format(time.RFC3339Nano) != testDownloadProto1.UpdatedAt {
		t.Errorf("Download updated at %s should have matched: %s",
			dl.UpdatedAt,
			testDownloadProto1.UpdatedAt,
		)
	}

	if dl.CreatedAt.Format(time.RFC3339Nano) != testDownloadProto1.CreatedAt {
		t.Errorf("Download created at %s should have matched: %s",
			dl.CreatedAt,
			testDownloadProto1.CreatedAt,
		)
	}
}

func Test_Unit_DownloadsToProtoMatch(t *testing.T) {
	dls := []Download{testDownload1}
	protos := DownloadsToProtos(dls)

	if len(protos) == 0 {
		t.Error("Protos is length of 0, should be 1")
	}

	for i := range dls {
		if dls[i].Id != int(protos[0].Id) {
			t.Errorf("Proto index 0 id should equal %d", dls[i].Id)
		}
	}
}
