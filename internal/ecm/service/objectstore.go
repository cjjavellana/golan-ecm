package service

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"cjavellana.me/ecm/golan/internal/ecm/pb"
	"cjavellana.me/ecm/golan/internal/ecm/service/builder"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"strings"
	"time"
)

type ObjectStoreService struct {
	ObjectStore ce.ObjectStore
	builder.PropertyFieldHierarchy

	// Required for future compatibility see https://github.com/grpc/grpc-go/blob/master/cmd/protoc-gen-go-grpc/README.md
	pb.UnimplementedContentEngineServer
}

func (s *ObjectStoreService) CreateWorkspace(_ context.Context, in *pb.CreateWorkspaceRequest) (*pb.CreateWorkspaceResponse, error) {

	log.Infof("received create workspace request: %s", in.WorkspaceName)

	now := time.Now()

	w := s.ObjectStore.NewWorkspace(ce.ObjectDescriptor{
		Name:        in.WorkspaceName,
		Label:       in.Label,
		Description: in.Description,
	})
	w.SetCreatedBy("UserFromAuthToken")
	w.SetDateCreated(&now)

	w, _ = s.ObjectStore.SaveWorkspace(w)

	return &pb.CreateWorkspaceResponse{
		ObjectId: w.ObjectId(),
	}, nil
}

func (s *ObjectStoreService) GetWorkspace(
	_ context.Context,
	in *pb.GetWorkspaceRequest) (*pb.GetWorkspaceResponse, error) {

	log.Infof("received get workspace query: %s", in.Query)

	w, err := s.ObjectStore.GetWorkspaceByObjectId(in.GetQuery())
	if err != nil {
		return nil, err
	}

	return &pb.GetWorkspaceResponse{
		ObjectId: w.ObjectId(),
	}, nil
}

func (s *ObjectStoreService) CreateDocumentClass(
	_ context.Context,
	in *pb.CreateDocumentClassRequest) (*pb.CreateDocumentClassResponse, error) {

	docClass := s.ObjectStore.NewDocumentClass(ce.ObjectDescriptor{
		Name:        in.Name,
		Label:       in.Label,
		Description: in.Description,
	})

	err := docClass.SetWorkspaceId(in.WorkspaceId)
	if err != nil {
		// an error may be returned when the given workspace id
		// cannot converted to a valid mongodb id
		return nil, err
	}

	var propertyFields []ce.PropertyField
	for _, propertyField := range in.PropertyFields {
		fieldType := ce.FieldType(strings.ToLower(propertyField.FieldType.String()))

		p := s.ObjectStore.NewPropertyField(
			ce.ObjectDescriptor{
				Name:        propertyField.Name,
				Label:       propertyField.Label,
				Description: propertyField.Description,
			},
			fieldType,
		)
		propertyFields = append(propertyFields, p)
	}

	docClass.SetPropertyFields(propertyFields)

	docClass, err = s.ObjectStore.SaveDocumentClass(docClass)
	if err != nil {
		return nil, err
	}

	return &pb.CreateDocumentClassResponse{
		ObjectId:    docClass.ObjectId(),
		Name:        in.Name,
		Label:       in.Label,
		Description: in.Description,
	}, nil
}

func (s *ObjectStoreService) CreateDocument(_ context.Context, in *pb.CreateDocumentRequest) (*pb.CreateDocumentResponse, error) {
	doc, err := s.ObjectStore.NewDocument(
		ce.ObjectDescriptor{
			Name:        in.Name,
			Label:       in.Label,
			Description: in.Description,
		},
		in.DocumentClassId,
	)
	if err != nil {
		log.Errorf("unable to create new document %v", err)
		return nil, err
	}

	if err = doc.SetParent(in.ParentId); err != nil {
		log.Errorf("incorrect document parent id format %s", in.ParentId)
		return nil, err
	}

	doc.SetContentType(in.ContentType)
	doc.SetFilename(in.Filename)
	doc.SetUnderlyingDocument(in.UnderlyingDocument)
	if in.IsVersioningEnabled {
		doc.EnableVersioning()
	}

	var attrs = make([]*ce.Attribute, len(in.Attributes))
	for i, v := range in.Attributes {
		attrs[i] = &ce.Attribute{
			Key:   v.Key,
			Value: v.Value,
			Type:  ce.FieldType(strings.ToLower(v.FieldType.String())),
		}
	}
	doc.SetAttributes(attrs)

	doc, err = s.ObjectStore.CreateDocument(doc)
	if err != nil {
		return nil, err
	}

	return &pb.CreateDocumentResponse{
		ObjectId: doc.ObjectId(),
	}, nil
}
