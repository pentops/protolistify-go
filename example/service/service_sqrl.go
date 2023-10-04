package service

import (
	"context"
	"database/sql"
	"encoding/base64"
	"log"
	"time"

	sq "github.com/elgris/sqrl"
	sspb "github.com/pentops/protoc-gen-listify/example/api/services/sample/v1"
	"github.com/pentops/protoc-gen-listify/listify/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gopkg.daemonl.com/sqrlx"
)

// An alternate example of a service implementation that uses listify with sqrl
type ServiceSqrl struct {
	*sspb.UnimplementedSampleServiceServer

	db              *sqrlx.Wrapper
	defaultPageSize int
	maxPageSize     int
}

func NewServiceSqrl(conn sqrlx.Connection) (*ServiceSqrl, error) {
	db, err := sqrlx.New(conn, sq.Dollar)
	if err != nil {
		return nil, err
	}

	return &ServiceSqrl{
		db:              db,
		defaultPageSize: 5,
		maxPageSize:     100,
	}, nil
}

func (s *ServiceSqrl) GetWidget(ctx context.Context, req *sspb.GetWidgetRequest) (*sspb.GetWidgetResponse, error) {
	log.Println("getting widget")

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	q := sq.
		Select("id", "customer_id", "details", "status", "created").
		From("widgets").
		Where("id = ?", req.Id)

	var widgets []*sspb.Widget
	err = s.db.Transact(ctx, nil, func(ctx context.Context, tx sqrlx.Transaction) error {
		widgets, err = s.getWidgets(ctx, tx, q)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &sspb.GetWidgetResponse{Widget: widgets[0]}, nil
}

func (s *ServiceSqrl) ListWidgets(ctx context.Context, req *sspb.ListWidgetsRequest) (*sspb.ListWidgetsResponse, error) {
	log.Println("listing widgets")

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	err = req.ValidateFilters()
	if err != nil {
		return nil, err
	}

	var total int64
	err = s.db.Transact(ctx, nil, func(ctx context.Context, tx sqrlx.Transaction) error {
		total, err = s.totalWidgets(ctx, tx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	filters := req.GetFilterClauses()
	limit := int64(s.defaultPageSize)

	if req.Page != nil {
		page := req.Page

		if page.Offset != "" {
			decoded, err := base64.StdEncoding.DecodeString(page.Offset)
			if err != nil {
				log.Printf("failed to decode page token: %s", err)
				return nil, status.Error(codes.InvalidArgument, "invalid page token")
			}

			c := &listify.FilterClause{
				Predicate: "created >= ?",
				Arguments: []*listify.FilterArgument{
					{
						Kind: &listify.FilterArgument_String_{
							String_: string(decoded),
						},
					},
				},
			}

			filters.Clauses = append(filters.Clauses, c)
		}

		if page.Limit > 0 && page.Limit <= int64(s.maxPageSize) {
			limit = page.Limit
		}
	}

	q := sq.
		Select("id", "customer_id", "details", "status", "created").
		From("widgets").
		OrderBy("created").
		Limit(uint64(limit))

	stmts, args := filters.ToSqrl()

	for i := range stmts {
		q.Where(stmts[i], args[i])
	}

	var widgets []*sspb.Widget
	err = s.db.Transact(ctx, nil, func(ctx context.Context, tx sqrlx.Transaction) error {
		widgets, err = s.getWidgets(ctx, tx, q)
		if err != nil {
			return err
		}

		return nil
	})

	page := &listify.PageResponse{
		NextOffset:       "",
		FinalPage:        true,
		TotalPageRecords: int64(len(widgets)),
		TotalRecords:     total,
	}

	if int64(len(widgets)) > limit {
		last := widgets[len(widgets)-1]
		widgets = widgets[:len(widgets)-1]

		page.NextOffset = base64.StdEncoding.EncodeToString([]byte(last.Created.AsTime().Format(time.RFC3339Nano)))
		page.FinalPage = false
	}

	resp := &sspb.ListWidgetsResponse{
		Widgets: widgets,
		Page:    page,
	}

	return resp, nil
}

func (s *ServiceSqrl) totalWidgets(ctx context.Context, tx sqrlx.Transaction) (int64, error) {
	q := sq.
		Select("COUNT(*)").
		From("widgets")

	var total int64
	err := tx.QueryRow(ctx, q).Scan(&total)
	if err != nil {
		log.Printf("failed to scan total: %s", err)
		return 0, status.Error(codes.Internal, "failed to list widgets")
	}

	return total, nil
}

func (s *ServiceSqrl) getWidgets(ctx context.Context, tx sqrlx.Transaction, query sqrlx.Sqlizer) ([]*sspb.Widget, error) {
	rows, err := tx.Query(ctx, query)
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
