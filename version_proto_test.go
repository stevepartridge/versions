package versions

import (
	// "fmt"
	"testing"
	"time"

	pb "github.com/stevepartridge/versions/protos"
)

var (
	testVersionProto1 = &pb.Version{
		Id:            int32(testVersion1.Id),
		Major:         testVersion1.Major,
		Minor:         testVersion1.Minor,
		Revision:      testVersion1.Revision,
		ApplicationId: testVersion1.ApplicationId,
		Name:          testVersion1.Name,
		Stable:        testVersion1.Stable,
		UpdatedAt:     now.Format(time.RFC3339Nano),
		CreatedAt:     now.Format(time.RFC3339Nano),
	}
)

func Test_Unit_ToProtoMatch(t *testing.T) {
	proto := testVersion1.ToProto()
	if proto.Id != testVersionProto1.Id {
		t.Errorf("Proto id should have matched: %d", testVersionProto1.Id)
	}
	if proto.Major != testVersionProto1.Major {
		t.Errorf("Proto major should have matched: %s", testVersionProto1.Major)
	}
	if proto.Minor != testVersionProto1.Minor {
		t.Errorf("Proto minor should have matched: %s", testVersionProto1.Minor)
	}
	if proto.Revision != testVersionProto1.Revision {
		t.Errorf("Proto revision should have matched: %s", testVersionProto1.Revision)
	}
	if proto.ApplicationId != testVersionProto1.ApplicationId {
		t.Errorf("Proto application id should have matched: %s", testVersionProto1.ApplicationId)
	}
	if proto.Name != testVersionProto1.Name {
		t.Errorf("Proto name should have matched: %s", testVersionProto1.Name)
	}
	if proto.Stable != testVersionProto1.Stable {
		t.Errorf("Proto stable should have matched: %t", testVersionProto1.Stable)
	}
	if proto.UpdatedAt != testVersionProto1.UpdatedAt {
		t.Errorf("Proto updated at %s should have matched: %s",
			proto.UpdatedAt,
			testVersionProto1.UpdatedAt,
		)
	}
	if proto.CreatedAt != testVersionProto1.CreatedAt {
		t.Errorf("Proto created at %s should have matched: %s",
			proto.CreatedAt,
			testVersionProto1.CreatedAt,
		)
	}
}

func Test_Unit_FromProtoMatch(t *testing.T) {

	testVersionProto1.Application = &pb.Application{
		Id:   "1234",
		Name: "Application Name",
	}

	ver := FromProto(testVersionProto1)
	if int32(ver.Id) != testVersionProto1.Id {
		t.Errorf("Version id should have matched: %d", testVersionProto1.Id)
	}
	if ver.Major != testVersionProto1.Major {
		t.Errorf("Version major should have matched: %s", testVersionProto1.Major)
	}
	if ver.Minor != testVersionProto1.Minor {
		t.Errorf("Version minor should have matched: %s", testVersionProto1.Minor)
	}
	if ver.Revision != testVersionProto1.Revision {
		t.Errorf("Version revision should have matched: %s", testVersionProto1.Revision)
	}
	if ver.ApplicationId != testVersionProto1.ApplicationId {
		t.Errorf("Version application id should have matched: %s", testVersionProto1.ApplicationId)
	}
	if ver.Name != testVersionProto1.Name {
		t.Errorf("Version name should have matched: %s", testVersionProto1.Name)
	}
	if ver.Stable != testVersionProto1.Stable {
		t.Errorf("Version stable should have matched: %t", testVersionProto1.Stable)
	}
	if ver.UpdatedAt.Format(time.RFC3339Nano) != testVersionProto1.UpdatedAt {
		t.Errorf("Version updated at %s should have matched: %s",
			ver.UpdatedAt,
			testVersionProto1.UpdatedAt,
		)
	}
	if ver.CreatedAt.Format(time.RFC3339Nano) != testVersionProto1.CreatedAt {
		t.Errorf("Version created at %s should have matched: %s",
			ver.CreatedAt,
			testVersionProto1.CreatedAt,
		)
	}
}

func Test_Unit_VersionsToProtoMatch(t *testing.T) {
	vers := []Version{testVersion1}
	protos := VersionsToProtos(vers)

	if len(protos) == 0 {
		t.Error("Protos is length of 0, should be 1")
	}

	for i := range vers {
		if vers[i].Id != int(protos[0].Id) {
			t.Errorf("Proto index 0 id should equal %d", vers[i].Id)
		}
	}
}
