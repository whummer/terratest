package aws

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/stretchr/testify/assert"
)

func TestNewAuthenticatedSessionFromDefaultCredentials(t *testing.T) {
	localEndpoints := map[string]string{
		"apigateway":     "http://localhost:4566",
		"apigatewayv2":   "http://localhost:4566",
		"cloudformation": "http://localhost:4566",
		"cloudwatch":     "http://localhost:4566",
		"dynamodb":       "http://localhost:4566",
		"es":             "http://localhost:4566",
		"firehose":       "http://localhost:4566",
		"iam":            "http://localhost:4566",
		"kinesis":        "http://localhost:4566",
		"lambda":         "http://localhost:4566",
		"route53":        "http://localhost:4566",
		"redshift":       "http://localhost:4566",
		"s3":             "http://localhost:4566",
		"secretsmanager": "http://localhost:4566",
		"ses":            "http://localhost:4566",
		"sns":            "http://localhost:4566",
		"sqs":            "http://localhost:4566",
		"ssm":            "http://localhost:4566",
		"stepfunctions":  "http://localhost:4566",
		"sts":            "http://localhost:4566",
		"ec2":            "http://localhost:4566",
	}

	signingRegionExpected := "custom-signing-region"
	region := GetRandomStableRegion(t, nil, nil)

	SetAwsEndpointsOverrides(localEndpoints)
	session, errors := NewAuthenticatedSessionFromDefaultCredentials(region)
	resolver := session.Config.EndpointResolver
	resolverExpected, errors := resolver.EndpointFor(endpoints.S3ServiceID, signingRegionExpected)

	assert.NotNil(t, session)
	assert.NoError(t, errors)
	assert.Equal(t, signingRegionExpected, resolverExpected.SigningRegion)
	assert.Equal(t, localEndpoints[endpoints.S3ServiceID], resolverExpected.URL)
}
