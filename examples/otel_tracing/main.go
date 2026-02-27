// Copyright © 2025 Ping Identity Corporation

// Package main demonstrates how to configure OpenTelemetry distributed tracing
// with the PingOne Go Client SDK. It shows how to inject a TracerProvider and
// a session ID so that every outbound API call is automatically traced and
// correlated with the X-Ping-External-Transaction-ID and
// X-Ping-External-Session-ID headers.
//
// This example exports spans to stdout so you can see the raw trace data
// without requiring a running collector. Swap the exporter for an OTLP
// exporter when integrating with Jaeger, Tempo, Datadog, or another backend.
//
// Required environment variables:
//
//	PINGONE_CLIENT_ID       – OAuth2 client ID
//	PINGONE_CLIENT_SECRET   – OAuth2 client secret
//	PINGONE_ENVIRONMENT_ID  – PingOne environment UUID
package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/google/uuid"
	"github.com/pingidentity/pingone-go-client/config"
	"github.com/pingidentity/pingone-go-client/oauth2"
	"github.com/pingidentity/pingone-go-client/pingone"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	// ── Logging ──────────────────────────────────────────────────────────────
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})))

	ctx := context.Background()

	// ── OpenTelemetry setup ───────────────────────────────────────────────────
	// Export spans to stdout for easy local inspection.
	// In production, replace stdouttrace.New() with an OTLP exporter such as:
	//   otlptracegrpc.New(ctx, otlptracegrpc.WithEndpoint("localhost:4317"))
	exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		slog.Error("Failed to create trace exporter", "error", err)
		os.Exit(1)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		// Add resource attributes here (service.name, service.version, etc.):
		// sdktrace.WithResource(resource.NewWithAttributes(
		// 	semconv.SchemaURL,
		// 	semconv.ServiceName("my-pingone-app"),
		// 	semconv.ServiceVersion("1.0.0"),
		// )),
	)
	defer func() {
		// Flush and shut down the provider so all in-flight spans are exported
		// before the process exits.
		if shutdownErr := tp.Shutdown(ctx); shutdownErr != nil {
			slog.Warn("Failed to shut down tracer provider", "error", shutdownErr)
		}
	}()

	// Optionally register as the global provider so any library that uses
	// otel.GetTracerProvider() picks it up automatically.
	otel.SetTracerProvider(tp)

	// ── Environment variables ─────────────────────────────────────────────────
	clientID := os.Getenv("PINGONE_CLIENT_ID")
	clientSecret := os.Getenv("PINGONE_CLIENT_SECRET")
	authEnvironmentID := os.Getenv("PINGONE_ENVIRONMENT_ID")

	if clientID == "" || clientSecret == "" || authEnvironmentID == "" {
		slog.Error("Missing required environment variables",
			"PINGONE_CLIENT_ID", clientID != "",
			"PINGONE_CLIENT_SECRET", clientSecret != "",
			"PINGONE_ENVIRONMENT_ID", authEnvironmentID != "")
		os.Exit(1)
	}

	// ── SDK configuration ─────────────────────────────────────────────────────
	serviceCfg := config.NewConfiguration()
	serviceCfg.WithGrantType(oauth2.GrantTypeClientCredentials)
	serviceCfg.WithClientID(clientID)
	serviceCfg.WithClientSecret(clientSecret)
	serviceCfg.WithEnvironmentID(authEnvironmentID)
	serviceCfg.WithRootDomain("pingone.com")

	// Inject the TracerProvider so every API call and token acquisition is traced.
	// Spans are automatically created for:
	//   - Token acquisition  (pingone.config.AcquireTokenSource)
	//   - OAuth2 grant flows (pingone.config.OAuth2.ClientCredentials, etc.)
	//   - API operations     (pingone.Environments.GetEnvironmentById, etc.)
	//   - HTTP requests      (METHOD /path via otelhttp transport)
	serviceCfg.WithTracerProvider(tp)

	// Provide a session ID that will be sent as X-Ping-External-Session-ID on
	// every API request.  Use a value meaningful to your application, such as
	// a CLI run ID, a Terraform plan ID, or a request correlation ID.
	serviceCfg.WithSessionID("example-session-001")

	// ── Create the API client ─────────────────────────────────────────────────
	// Start a root span so that the token acquisition and all subsequent API
	// calls appear as children under a single trace.
	tracer := tp.Tracer("github.com/pingidentity/pingone-go-client/examples/otel_tracing")
	ctx, rootSpan := tracer.Start(ctx, "example.run")
	defer rootSpan.End()

	httpClient, err := serviceCfg.Client(ctx, nil)
	if err != nil {
		slog.Error("Failed to create HTTP client", "error", err)
		os.Exit(1)
	}

	configuration := pingone.NewConfiguration(serviceCfg)
	apiClient, err := pingone.NewAPIClient(ctx, configuration)
	if err != nil {
		slog.Error("Failed to create API client", "error", err)
		os.Exit(1)
	}
	apiClient.GetConfig().HTTPClient = httpClient

	// ── Make an API call ──────────────────────────────────────────────────────
	environmentID, err := uuid.Parse(authEnvironmentID)
	if err != nil {
		slog.Error("Invalid PINGONE_ENVIRONMENT_ID", "error", err)
		os.Exit(1)
	}

	// The SDK will:
	//  1. Start a "pingone.Environments.GetEnvironmentById" child span.
	//  2. Automatically populate X-Ping-External-Transaction-ID with the trace ID.
	//  3. Automatically populate X-Ping-External-Session-ID with "example-session-001".
	//  4. Create a child HTTP span via otelhttp for the outbound request.
	env, httpResp, err := apiClient.EnvironmentsApi.GetEnvironmentById(ctx, environmentID).Execute()
	if err != nil {
		slog.Error("API call failed", "error", err, "status", httpResp.Status)
		os.Exit(1)
	}

	slog.Info("Environment retrieved",
		"name", env.Name,
		"type", env.Type,
		"region", env.Region,
	)
}
