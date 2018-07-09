package versions

import (
	"time"

	pb "github.com/stevepartridge/versions/protos"
)

func (v *Version) ToProto() *pb.Version {
	ver := &pb.Version{
		Id:            int32(v.Id),
		Major:         v.Major,
		Minor:         v.Minor,
		Revision:      v.Revision,
		ApplicationId: v.ApplicationId,
		Name:          v.Name,
		Stable:        v.Stable,

		// Application:   v.Application,
		// Downloads: DownloadsToProtos(v.Downloads)

	}

	if !v.UpdatedAt.IsZero() {
		ver.UpdatedAt = v.UpdatedAt.Format(time.RFC3339Nano)
	}

	if !v.CreatedAt.IsZero() {
		ver.CreatedAt = v.CreatedAt.Format(time.RFC3339Nano)
	}

	return ver
}

func FromProto(proto *pb.Version) Version {

	ver := Version{
		Id:            int(proto.Id),
		Major:         proto.Major,
		Minor:         proto.Minor,
		Revision:      proto.Revision,
		ApplicationId: proto.ApplicationId,
		Name:          proto.Name,
		Stable:        proto.Stable,
	}

	if proto.Application != nil {

		app := ApplicationFromProto(proto.Application)
		if app.Id != "" {
			ver.ApplicationId = app.Id
		}
	}

	if proto.UpdatedAt != "" {
		ver.UpdatedAt, _ = time.Parse(time.RFC3339Nano, proto.UpdatedAt)
	}

	if proto.CreatedAt != "" {
		ver.CreatedAt, _ = time.Parse(time.RFC3339Nano, proto.CreatedAt)
	}

	return ver
}

func VersionsToProtos(list []Version) []*pb.Version {
	result := make([]*pb.Version, len(list))
	for i := range list {
		result[i] = list[i].ToProto()
	}
	return result
}
