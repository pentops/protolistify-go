package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	sspb "github.com/pentops/protoc-gen-listify/example/api/services/sample/v1"
	"github.com/pentops/protoc-gen-listify/example/service"

	"github.com/gomicro/blockit/cbblocker"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"gopkg.daemonl.com/envconf"
)

var config = struct {
	Port        int    `env:"PORT" default:"8080"`
	JsonPort    int    `env:"JSON_PORT" default:"8081"`
	PostgresURL string `env:"POSTGRES_URL" default:"postgres://test:test@localhost:5432/test?sslmode=disable"`
}{}

func main() {
	if err := envconf.Parse(&config); err != nil {
		log.Fatalf("failed to parse config: %s", err)
	}

	err := startService()
	if err != nil {
		log.Fatalf("failed to start service: %s", err)
	}
}

func startService() error {
	ctx := context.Background()

	db, err := sql.Open("postgres", config.PostgresURL)
	if err != nil {
		return fmt.Errorf("failed to open postgres connection: %w", err)
	}

	<-cbblocker.New(func() error {
		log.Print("waiting for postgres to be ready")
		return db.PingContext(ctx)
	}, 1*time.Second).BlockitWithContext(ctx)

	<-cbblocker.New(func() error {
		err := initData(ctx, db)
		if err != nil {
			log.Printf("failed to init data: %s", err)
		}

		return err
	}, 1*time.Second).BlockitWithContext(ctx)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(grpcServer(ctx, db))
	g.Go(jsonServer(ctx))

	return g.Wait()
}

func initData(ctx context.Context, db *sql.DB) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Commit()

	_, err = tx.ExecContext(ctx, `
		DROP TABLE IF EXISTS widgets;
		DROP TYPE IF EXISTS widget_status;

		CREATE TYPE widget_status AS ENUM (
			'WIDGET_STATUS_UNSPECIFIED',
			'WIDGET_STATUS_REQUESTED',
			'WIDGET_STATUS_CREATED'
		);

		CREATE TABLE IF NOT EXISTS widgets (
			id text PRIMARY KEY,
			customer_id text NOT NULL,
			details jsonb NOT NULL,
			status widget_status NOT NULL,
			created timestamptz NOT NULL DEFAULT now()
		);

		INSERT INTO widgets (id, customer_id, details, status, created)
		VALUES
			('c6583e80-83c8-4366-8597-b9cbb75feb46', 'b4911f49-bc4d-43ed-bbdd-edfeef7f2530', '{"name": "widget 1", "weight": "100", "square": {"width": "5", "height": "5"}}', 'WIDGET_STATUS_REQUESTED', '2023-07-10T08:00:00Z'),
			('51e48853-0921-4c3c-ab24-ad9360e2efcc', 'b4911f49-bc4d-43ed-bbdd-edfeef7f2530', '{"name": "widget 2", "weight": "88", "square": {"width": "5", "height": "5"}}', 'WIDGET_STATUS_REQUESTED', '2023-07-11T08:00:00Z'),
			('37ddef59-139e-4c87-98a5-92923f7abaed', 'b4911f49-bc4d-43ed-bbdd-edfeef7f2530', '{"name": "widget 3", "weight": "42", "round":{"diameter": "2"}}', 'WIDGET_STATUS_CREATED', '2023-07-12T08:00:00Z'),
			('565c3595-c67a-4008-b28a-d75fdfd6c428', 'b4911f49-bc4d-43ed-bbdd-edfeef7f2530', '{"name": "widget 4", "weight": "314", "round":{"diameter": "15"}}', 'WIDGET_STATUS_CREATED', '2023-07-13T08:00:00Z'),
			('93b85b94-73ea-4da2-83c0-e7780ec77e5c', 'b4911f49-bc4d-43ed-bbdd-edfeef7f2530', '{"name": "widget 5", "weight": "100", "square": {"width": "5", "height": "5"}}', 'WIDGET_STATUS_CREATED', '2023-07-14T08:00:00Z'),
			('7740b4d6-d266-4ecc-ab31-0c2a4681ae05', '7889974f-a676-4248-97b6-f03338035d18', '{"name": "widget 6", "weight": "88", "square": {"width": "5", "height": "5"}}', 'WIDGET_STATUS_CREATED', '2023-07-15T09:00:00Z'),
			('4caa87c3-2763-4dfb-94ed-0383894d4356', '7889974f-a676-4248-97b6-f03338035d18', '{"name": "widget 7", "weight": "42", "round":{"diameter": "2"}}', 'WIDGET_STATUS_CREATED', '2023-07-16T09:00:00Z'),
			('9c80f827-80d0-40b6-88df-932016ca5c00', '7889974f-a676-4248-97b6-f03338035d18', '{"name": "widget 8", "weight": "314", "round":{"diameter": "15"}}', 'WIDGET_STATUS_CREATED', '2023-07-17T09:00:00Z')
		;
	`)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}

func grpcServer(ctx context.Context, db *sql.DB) func() error {
	return func() error {
		serv := service.NewService(db)

		grpcServer := grpc.NewServer(
			grpc.UnaryInterceptor(
				grpc_recovery.UnaryServerInterceptor(),
			),
		)

		reflection.Register(grpcServer)

		sspb.RegisterSampleServiceServer(grpcServer, serv)

		log.Printf("starting service on port %d", config.Port)
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
		if err != nil {
			return fmt.Errorf("failed to listen: %w", err)
		}

		go func() {
			<-ctx.Done()
			log.Println("shutting down service")
			grpcServer.GracefulStop()
		}()

		return grpcServer.Serve(lis)
	}
}

func jsonServer(ctx context.Context) func() error {
	return func() error {
		c, err := grpc.Dial(fmt.Sprintf("localhost:%d", config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
		if err != nil {
			return fmt.Errorf("failed to dial: %w", err)
		}

		mux := runtime.NewServeMux()

		err = sspb.RegisterSampleServiceHandler(ctx, mux, c)
		if err != nil {
			return fmt.Errorf("failed to register sample service handler: %w", err)
		}

		var handler http.Handler = mux

		srv := &http.Server{
			Addr:    fmt.Sprintf(":%d", config.JsonPort),
			Handler: handler,
		}

		go func() {
			<-ctx.Done()
			log.Println("shutting down json service")
			srv.Shutdown(ctx)
		}()

		log.Printf("starting json service on port %d", config.JsonPort)
		return srv.ListenAndServe()
	}
}
