package versions

import (
	// "fmt"
	"testing"
	"time"

	pb "github.com/stevepartridge/versions/protos"
)

var (
	testApplicationProto1 = &pb.Application{
		Id:        testApplication1.Id,
		Name:      testApplication1.Name,
		UpdatedAt: now.Format(time.RFC3339Nano),
		CreatedAt: now.Format(time.RFC3339Nano),
	}
)

func Test_Unit_ApplicationToProtoMatch(t *testing.T) {

	proto := testApplication1.ToProto()

	if proto.Id != testApplicationProto1.Id {
		t.Errorf("Proto id should have matched: %d", testApplicationProto1.Id)
	}

	if proto.Name != testApplicationProto1.Name {
		t.Errorf("Proto name should have matched: %s", testApplicationProto1.Name)
	}

	if proto.UpdatedAt != testApplicationProto1.UpdatedAt {
		t.Errorf("Proto updated at %s should have matched: %s",
			proto.UpdatedAt,
			testApplicationProto1.UpdatedAt,
		)
	}

	if proto.CreatedAt != testApplicationProto1.CreatedAt {
		t.Errorf("Proto created at %s should have matched: %s",
			proto.CreatedAt,
			testApplicationProto1.CreatedAt,
		)
	}
}

func Test_Unit_ApplicationFromProtoMatch(t *testing.T) {

	app := ApplicationFromProto(testApplicationProto1)
	if app.Id != testApplicationProto1.Id {
		t.Errorf("Application id should have matched: %d", testApplicationProto1.Id)
	}

	if app.Name != testApplicationProto1.Name {
		t.Errorf("Application name should have matched: %s", testApplicationProto1.Name)
	}

	if app.UpdatedAt.Format(time.RFC3339Nano) != testApplicationProto1.UpdatedAt {
		t.Errorf("Application updated at %s should have matched: %s",
			app.UpdatedAt,
			testApplicationProto1.UpdatedAt,
		)
	}

	if app.CreatedAt.Format(time.RFC3339Nano) != testApplicationProto1.CreatedAt {
		t.Errorf("Application created at %s should have matched: %s",
			app.CreatedAt,
			testApplicationProto1.CreatedAt,
		)
	}
}

func Test_Unit_ApplicationsToProtoMatch(t *testing.T) {
	apps := []Application{testApplication1}
	protos := ApplicationsToProtos(apps)

	if len(protos) == 0 {
		t.Error("Protos is length of 0, should be 1")
	}

	for i := range apps {
		if apps[i].Id != protos[0].Id {
			t.Errorf("Proto index 0 id should equal %d", apps[i].Id)
		}
	}
}
