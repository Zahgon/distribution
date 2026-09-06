package s3

import (
	"net/http"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
)

type signer struct {
	Request      *http.Request
	Time         time.Time
	Credentials  *credentials.Credentials
	Query        url.Values
	stringToSign string
	signature    string
}

var s3ParamsToSign = map[string]bool{
	"acl":                          true,
	"location":                     true,
	"logging":                      true,
	"notification":                 true,
	"partNumber":                   true,
	"policy":                       true,
	"requestPayment":               true,
	"torrent":                      true,
	"uploadId":                     true,
	"uploads":                      true,
	"versionId":                    true,
	"versioning":                   true,
	"versions":                     true,
	"response-content-type":        true,
	"response-content-language":    true,
	"response-expires":             true,
	"response-cache-control":       true,
	"response-content-disposition": true,
	"response-content-encoding":    true,
	"website":                      true,
	"delete":                       true,
}

func setv2Handlers(svc *s3.S3) { _ = "STUB: not implemented"; return }

func Sign(req *request.Request) { _ = "STUB: not implemented"; return }

func (v2 *signer) Sign() error { _ = "STUB: not implemented"; return nil }
