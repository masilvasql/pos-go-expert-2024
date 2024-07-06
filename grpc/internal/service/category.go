package service

import (
	"context"
	"io"

	"github.com/masilvasql/gprc-pos/internal/database"
	"github.com/masilvasql/gprc-pos/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(db database.Category) *CategoryService {
	return &CategoryService{CategoryDB: db}
}

func (s *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := s.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categryResponse,
	}, nil
}

func (s *CategoryService) ListCategory(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	category, err := s.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var categories []*pb.Category
	for _, c := range category {
		categories = append(categories, &pb.Category{
			Id:          c.ID,
			Name:        c.Name,
			Description: c.Description,
		})
	}

	return &pb.CategoryList{
		Categories: categories,
	}, nil
}

func (s *CategoryService) GetCategoryById(ctx context.Context, in *pb.CategoryRequestById) (*pb.CategoryResponseById, error) {
	category, err := s.CategoryDB.FindByID(in.Id)
	if err != nil {
		return nil, err
	}

	categryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponseById{
		Category: categryResponse,
	}, nil
}

func (s *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := []*pb.Category{}
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.CategoryList{
				Categories: categories,
			})
		}
		if err != nil {
			return err
		}

		c, err := s.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		categories = append(categories, &pb.Category{
			Id:          c.ID,
			Name:        c.Name,
			Description: c.Description,
		})

	}
}

func (s *CategoryService) CreateCategoryStreamBidirectional(stream pb.CategoryService_CreateCategoryStreamBidirectionalServer) error {

	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		c, err := s.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		err = stream.Send(&pb.Category{
			Id:          c.ID,
			Name:        c.Name,
			Description: c.Description,
		})

		if err != nil {
			return err
		}

	}
}
