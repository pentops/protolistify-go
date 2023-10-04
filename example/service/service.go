package service

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"time"

	sspb "github.com/pentops/protoc-gen-listify/example/api/services/sample/v1"
	"github.com/pentops/protoc-gen-listify/listify/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// A sample service implementation that uses listfy with bare sql
type Service struct {
	*sspb.UnimplementedSampleServiceServer

	db              *sql.DB
	defaultPageSize int
	maxPageSize     int
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db:              db,
		defaultPageSize: 5,
		maxPageSize:     100,
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

	widgets, err := s.getWidgets(ctx, tx, q, []interface{}{req.Id})
	if err != nil {
		return nil, err
	}

	return &sspb.GetWidgetResponse{Widget: widgets[0]}, nil
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

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Commit() //nolint:errcheck

	total, err := s.totalWidgets(ctx, tx)
	if err != nil {
		return nil, err
	}

	filters := req.GetFilterClauses()

	if req.Page != "" {
		decoded, err := base64.StdEncoding.DecodeString(req.Page)
		if err != nil {
			log.Printf("failed to decode page token: %s", err)
			return nil, status.Error(codes.InvalidArgument, "invalid page token")
		}

		c := &listify.FilterClause{
			Predicate: "created >= ?",
			Arguments: []*listify.FilterArgument{
				{
					Kind: nil,
				},
			},
		}

		filters.Clauses = append(filters.Clauses, c)
	}

	q := `
	SELECT
		id,
		customer_id,
		details,
		status,
		created
	FROM widgets`

	if len(args) > 0 {
		q += "\n WHERE " + strings.Join(stmts, " AND ")
	}

	q += "\n ORDER BY created"

	limit := int64(s.defaultPageSize)
	if req.Limit > 0 && req.Limit <= int64(s.maxPageSize) {
		limit = req.Limit
	}

	q += fmt.Sprintf("\n LIMIT %d", limit+1)

	widgets, err := s.getWidgets(ctx, tx, q, args)
	if err != nil {
		return nil, err
	}

	page := &listify.Page{
		NextPage:         "",
		FinalPage:        true,
		TotalPageRecords: int64(len(widgets)),
		TotalRecords:     total,
	}

	if int64(len(widgets)) > limit {
		last := widgets[len(widgets)-1]
		widgets = widgets[:len(widgets)-1]

		page.NextPage = base64.StdEncoding.EncodeToString([]byte(last.Created.AsTime().Format(time.RFC3339Nano)))
		page.FinalPage = false
	}

	resp := &sspb.ListWidgetsResponse{
		Widgets: widgets,
		Page:    page,
	}

	return resp, nil
}

func (s *Service) totalWidgets(ctx context.Context, tx *sql.Tx) (int64, error) {
	q := "SELECT COUNT(*) FROM widgets"

	var total int64
	err := tx.QueryRowContext(ctx, q).Scan(&total)
	if err != nil {
		log.Printf("failed to scan total: %s", err)
		return 0, status.Error(codes.Internal, "failed to list widgets")
	}

	return total, nil
}

func (s *Service) getWidgets(ctx context.Context, tx *sql.Tx, query string, args []interface{}) ([]*sspb.Widget, error) {
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		log.Printf("failed to query widgets: %s", err)
		return nil, status.Error(codes.Internal, "failed to list widgets")
	}
	defer rows.Close()

	var widgets []*sspb.Widget
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

		widgets = append(widgets, &w)
	}

	if rows.Err() != nil {
		log.Printf("failed to scan widgets: %s", rows.Err())
		return nil, status.Error(codes.Internal, "failed to list widgets")
	}

	return widgets, nil
}
