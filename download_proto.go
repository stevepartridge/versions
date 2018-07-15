package versions

import (
	"time"

	pb "github.com/stevepartridge/versions/protos"
)

func (d *Download) ToProto() *pb.Download {

	dl := &pb.Download{
		Id:             int32(d.Id),
		StorageType:    d.StorageType,
		VersionId:      int32(d.VersionId),
		Os:             d.OS,
		Arch:           d.Arch,
		Filename:       d.Filename,
		Extension:      d.Ext,
		TotalDownloads: int32(d.Downloads),
	}

	if !d.UpdatedAt.IsZero() {
		dl.UpdatedAt = d.UpdatedAt.Format(time.RFC3339Nano)
	}

	if !d.CreatedAt.IsZero() {
		dl.CreatedAt = d.CreatedAt.Format(time.RFC3339Nano)
	}

	return dl
}

func DownloadFromProto(proto *pb.Download) Download {

	dl := Download{
		Id:          int(proto.Id),
		VersionId:   int(proto.VersionId),
		StorageType: proto.StorageType,
		OS:          proto.Os,
		Arch:        proto.Arch,
		Filename:    proto.Filename,
		Ext:         proto.Extension,
	}

	if proto.Data != nil {
		dl.Data = proto.Data
	}

	if proto.UpdatedAt != "" {
		dl.UpdatedAt, _ = time.Parse(time.RFC3339Nano, proto.UpdatedAt)
	}

	if proto.CreatedAt != "" {
		dl.CreatedAt, _ = time.Parse(time.RFC3339Nano, proto.CreatedAt)
	}

	return dl
}

func DownloadsToProtos(list []Download) []*pb.Download {
	result := make([]*pb.Download, len(list))
	for i := range list {
		result[i] = list[i].ToProto()
	}
	return result
}
