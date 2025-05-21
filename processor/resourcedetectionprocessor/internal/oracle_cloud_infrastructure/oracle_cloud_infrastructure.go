// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package oracle_cloud_infrastructure // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/oracle_cloud_infrastructure"

import (
	"context"
	"os"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/oracle_cloud_infrastructure/internal/metadata"

	"fmt"
	"io"
	"net/http"

	"github.com/tidwall/gjson"
)

const (
	// TypeStr is type of detector.
	TypeStr = "oracle_cloud_infrastructure"
)

// NewDetector returns a detector which can detect resource attributes on Oracle Cloud Infrastructure
func NewDetector(set processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	cfg := dcfg.(Config)
	return &detector{
		logger: set.Logger,
		rb:     metadata.NewResourceBuilder(cfg.ResourceAttributes),
	}, nil
}

type detector struct {
	logger *zap.Logger
	rb     *metadata.ResourceBuilder
}

// Detect detects oracle_cloud_infrastructure metadata and returns a resource with the available ones
func (d *detector) Detect(_ context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	dynoIDMissing := false
	if dynoID, ok := os.LookupEnv("HEROKU_DYNO_ID"); ok {
		d.rb.SetServiceInstanceID(dynoID)
	} else {
		dynoIDMissing = true
	}

	oracle_cloud_infrastructureAppIDMissing := false
	if v, ok := os.LookupEnv("HEROKU_APP_ID"); ok {
		d.rb.SetHerokuAppID(v)
	} else {
		oracle_cloud_infrastructureAppIDMissing = true
	}
	if dynoIDMissing {
		if oracle_cloud_infrastructureAppIDMissing {
			d.logger.Debug("Heroku metadata is missing. Please check metadata is enabled.")
		} else {
			// some oracle_cloud_infrastructure deployments will enable some of the metadata.
			d.logger.Debug("Partial Heroku metadata is missing. Please check metadata is supported.")
		}
	}
	if !oracle_cloud_infrastructureAppIDMissing {
		d.rb.SetCloudProvider("oracle_cloud_infrastructure")
	}
	if v, ok := os.LookupEnv("HEROKU_APP_NAME"); ok {
		d.rb.SetServiceName(v)
	}
	if v, ok := os.LookupEnv("HEROKU_RELEASE_CREATED_AT"); ok {
		d.rb.SetHerokuReleaseCreationTimestamp(v)
	}
	if v, ok := os.LookupEnv("HEROKU_RELEASE_VERSION"); ok {
		d.rb.SetServiceVersion(v)
	}
	if v, ok := os.LookupEnv("HEROKU_SLUG_COMMIT"); ok {
		d.rb.SetHerokuReleaseCommit(v)
	}

	return d.rb.Emit(), conventions.SchemaURL, nil
}

const metadataEndpoint = "http://169.254.169.254/opc/v2/instance/"

var allowList = []string{
	"instanceId",
	"displayName",
	"shape",
	"regionInfo.regionIdentifier",
	"availabilityDomain",
}

func fetchMetadata(ctx context.Context, url string) (string, error) {
	// Create HTTP client
	client := &http.Client{}

	// Create request to query the OCI metadata endpoint
	req, err := http.NewRequest("GET", metadataEndpoint, nil)
	if err != nil {
		return "", fmt.Errorf("Error creating request: %w\n", err)
	}

	// Add required Auth header for IMDSv2
	req.Header.Add("Authorization", "Bearer Oracle")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error making request: %w\n", err)
	}

	// Ensure that we don't leave dangling HTTP connections
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Received non-OK response: %d\n", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response data: %w\n", err)
	}

	return string(body), nil
}

func printFilteredMetadata(jsonData string, allowList []string) {
	fmt.Println("Requested Metadata:")
	for _, key := range allowList {
		value := gjson.Get(jsonData, key)
		if value.Exists() {
			fmt.Printf("%s: %v\n", key, value)
		}
	}
}

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	jsonData, err := fetchMetadata(ctx, metadataEndpoint)
// 	if err != nil {
// 		log.Fatalf("Failed to fetch metadata: %v\n", err)
// 	}

// 	printFilteredMetadata(jsonData, allowList)
// }
