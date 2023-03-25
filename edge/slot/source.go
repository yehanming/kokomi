package slot

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"snple.com/kokomi/pb"
	"snple.com/kokomi/pb/edges"
	"snple.com/kokomi/pb/slots"
)

type SourceService struct {
	ss *SlotService

	slots.UnimplementedSourceServiceServer
}

func newSourceService(ss *SlotService) *SourceService {
	return &SourceService{
		ss: ss,
	}
}

func (s *SourceService) Create(ctx context.Context, in *pb.Source) (*pb.Source, error) {
	var output pb.Source
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetSource().Create(ctx, in)
}

func (s *SourceService) Update(ctx context.Context, in *pb.Source) (*pb.Source, error) {
	var output pb.Source
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetSource().Update(ctx, in)
}

func (s *SourceService) View(ctx context.Context, in *pb.Id) (*pb.Source, error) {
	var output pb.Source
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetSource().View(ctx, in)
}

func (s *SourceService) ViewByName(ctx context.Context, in *pb.Name) (*pb.Source, error) {
	var output pb.Source
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetSource().ViewByName(ctx, in)
}

func (s *SourceService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetSource().Delete(ctx, in)
}

func (s *SourceService) List(ctx context.Context, in *slots.ListSourceRequest) (*slots.ListSourceResponse, error) {
	var err error
	var output slots.ListSourceResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &edges.ListSourceRequest{
		Page:   in.GetPage(),
		Tags:   in.GetTags(),
		Type:   in.GetType(),
		Source: in.GetSource(),
	}

	reply, err := s.ss.es.GetSource().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.Source = reply.GetSource()

	return &output, nil
}

func (s *SourceService) Link(ctx context.Context, in *slots.LinkSourceRequest) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	slotID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request2 := &edges.LinkSourceRequest{Id: in.GetId(), Status: in.GetStatus()}

	reply, err := s.ss.es.GetSource().Link(ctx, request2)
	if err != nil {
		return &output, err
	}

	s.ss.es.GetControl().LinkSource(slotID, in.GetId(), in.GetStatus())

	return reply, nil
}

func (s *SourceService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Source, error) {
	var output pb.Source
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ss.es.GetSource().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	return reply, nil
}

func (s *SourceService) Pull(ctx context.Context, in *slots.PullSourceRequest) (*slots.PullSourceResponse, error) {
	var err error
	var output slots.PullSourceResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &edges.PullSourceRequest{
		After: in.GetAfter(),
		Limit: in.GetLimit(),
	}

	reply, err := s.ss.es.GetSource().Pull(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Source = reply.GetSource()

	return &output, nil
}
