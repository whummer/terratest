package aws

//Reference for mocking purposes https://github.com/aws/aws-sdk-go/blob/master/aws/config.go

import (
	"net/http"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
)

type RequestRetryer interface{}

type AwsConfigOverrides struct {
	CredentialsChainVerboseErrors *bool

	Credentials *credentials.Credentials

	Endpoint *string

	EndpointResolver endpoints.Resolver

	EnforceShouldRetryCheck *bool

	Region *string

	DisableSSL *bool

	HTTPClient *http.Client

	MaxRetries *int

	Retryer RequestRetryer

	DisableParamValidation *bool

	DisableComputeChecksums *bool

	S3ForcePathStyle *bool

	S3Disable100Continue *bool

	S3UseAccelerate *bool

	S3DisableContentMD5Validation *bool

	S3UseARNRegion *bool

	EC2MetadataDisableTimeoutOverride *bool

	UseDualStack *bool

	SleepDelay func(time.Duration)

	DisableRestProtocolURICleaning *bool

	EnableEndpointDiscovery *bool

	DisableEndpointHostPrefix *bool

	STSRegionalEndpoint endpoints.STSRegionalEndpoint

	S3UsEast1RegionalEndpoint endpoints.S3UsEast1RegionalEndpoint
}

var once sync.Once

var defaultConfig *AwsConfigOverrides

func SetAwsConfigOverrides(config *AwsConfigOverrides) *AwsConfigOverrides {
	once.Do(func() {
		defaultConfig = config
	})

	return defaultConfig
}

func GetAwsConfigOverrides() *AwsConfigOverrides {
	return defaultConfig
}
