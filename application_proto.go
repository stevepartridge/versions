package versions

import (
	"time"

	pb "github.com/stevepartridge/versions/protos"
)

func (a *Application) ToProto() *pb.Application {

	app := &pb.Application{
		Id:   a.Id,
		Name: a.Name,
	}

	if !a.UpdatedAt.IsZero() {
		app.UpdatedAt = a.UpdatedAt.Format(time.RFC3339Nano)
	}

	if !a.CreatedAt.IsZero() {
		app.CreatedAt = a.CreatedAt.Format(time.RFC3339Nano)
	}

	return app
}

func ApplicationFromProto(proto *pb.Application) Application {

	app := Application{
		Id:   proto.Id,
		Name: proto.Name,
	}

	if proto.UpdatedAt != "" {
		app.UpdatedAt, _ = time.Parse(time.RFC3339Nano, proto.UpdatedAt)
	}

	if proto.CreatedAt != "" {
		app.CreatedAt, _ = time.Parse(time.RFC3339Nano, proto.CreatedAt)
	}

	return app
}

func ApplicationsToProtos(list []Application) []*pb.Application {
	result := make([]*pb.Application, len(list))
	for i := range list {
		result[i] = list[i].ToProto()
	}
	return result
}
