package service

import (
	"context"
	"database/sql"
	"log"
	"time"

	sspb "github.com/pentops/protoc-gen-listify/example/api/services/sample/v1"
	"github.com/pentops/protoc-gen-listify/listify/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	*sspb.UnimplementedSampleServiceServer

	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) GetWidget(ctx context.Context, req *sspb.GetWidgetRequest) (*sspb.GetWidgetResponse, error) {
	log.Println("getting widget")

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	q := `
	SELECT
		id,
		customer_id,
		details,
		status,
		created
	FROM widgets
	WHERE id = $1`

	row := tx.QueryRowContext(ctx, q, req.Id)

	var w sspb.Widget
	var details []byte
	var stat string
	var created time.Time

	err = row.Scan(&w.Id, &w.CustomerId, &details, &stat, &created)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "widget not found")
		}

		log.Printf("failed to scan widget: %s", err)
		return nil, status.Error(codes.Internal, "failed to get widget")
	}

	w.Status = sspb.WidgetStatus(sspb.WidgetStatus_value[stat])
	w.Created = timestamppb.New(created)

	var d sspb.WidgetDetails
	err = protojson.Unmarshal(details, &d)
	if err != nil {
		log.Printf("failed to unmarshal widget details: %s", err)
		return nil, status.Error(codes.Internal, "failed to get widget")
	}

	w.Details = &d

	return &sspb.GetWidgetResponse{Widget: &w}, nil
}

func (s *Service) ListWidgets(ctx context.Context, req *sspb.ListWidgetsRequest) (*sspb.ListWidgetsResponse, error) {
	log.Println("listing widgets")

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	err = req.ValidateFilters()
	if err != nil {
		return nil, err
	}

	//q, args, err := req.FiltersToSql()

	log.Printf("filters: %v", req.Filters)

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	q := "SELECT COUNT(*) FROM widgets"

	var total int64
	err = tx.QueryRowContext(ctx, q).Scan(&total)
	if err != nil {
		log.Printf("failed to scan total: %s", err)
		return nil, status.Error(codes.Internal, "failed to list widgets")
	}

	q = `
	SELECT
		id,
		customer_id,
		details,
		status,
		created
	FROM widgets`

	rows, err := tx.QueryContext(ctx, q)
	if err != nil {
		log.Printf("failed to query widgets: %s", err)
		return nil, status.Error(codes.Internal, "failed to list widgets")
	}
	defer rows.Close()

	resp := &sspb.ListWidgetsResponse{}

	for rows.Next() {
		var w sspb.Widget
		var details []byte
		var stat string
		var created time.Time

		err = rows.Scan(&w.Id, &w.CustomerId, &details, &stat, &created)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, status.Error(codes.NotFound, "widget not found")
			}

			log.Printf("failed to scan widget: %s", err)
			return nil, status.Error(codes.Internal, "failed to get widget")
		}

		w.Status = sspb.WidgetStatus(sspb.WidgetStatus_value[stat])
		w.Created = timestamppb.New(created)

		var d sspb.WidgetDetails
		err = protojson.Unmarshal(details, &d)
		if err != nil {
			log.Printf("failed to unmarshal widget details: %s", err)
			return nil, status.Error(codes.Internal, "failed to get widget")
		}

		w.Details = &d

		resp.Widgets = append(resp.Widgets, &w)
	}

	if rows.Err() != nil {
		log.Printf("failed to scan widgets: %s", rows.Err())
		return nil, status.Error(codes.Internal, "failed to list widgets")
	}

	resp.Page = &listify.Page{
		TotalRecords: total,
	}

	return resp, nil
}
