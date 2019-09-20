package repository

import (
	"context"
	pb "grpctest/server/proto"
)

type StudentRepo interface {
	Create(ctx context.Context, in *pb.Student) (*pb.StudentID, error)
	Read(ctx context.Context, in *pb.StudentID) (*pb.Student, error)
	Update(ctx context.Context, in *pb.Student) (*pb.StudentID, error)
	Delete(ctx context.Context, in *pb.StudentID) (*pb.StudentID, error)
}
