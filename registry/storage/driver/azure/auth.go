package azure

import (
	"context"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/service"
)

const (
	UDCGracePeriod = 30.0 * time.Minute
	UDCExpiryTime  = 48.0 * time.Hour
)

type signer interface {
	Sign(context.Context, *sas.BlobSignatureValues) (sas.QueryParameters, error)
}

type sharedKeySigner struct {
	cred *azblob.SharedKeyCredential
}

type clientTokenSigner struct {
	client    *azblob.Client
	cred      azcore.TokenCredential
	udcMutex  sync.Mutex
	udc       *service.UserDelegationCredential
	udcExpiry time.Time
}

type azureClient struct {
	container string
	client    *azblob.Client
	signer    signer
}

func newClient(params *DriverParameters) (*azureClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newTokenClient(params *DriverParameters) (*azureClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newSharedKeyCredentialsClient(params *DriverParameters) (*azureClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *azureClient) ContainerClient() *container.Client { _ = "STUB: not implemented"; return nil }

func (a *azureClient) SignBlobURL(ctx context.Context, blobURL string, expires time.Time) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *sharedKeySigner) Sign(ctx context.Context, signatureValues *sas.BlobSignatureValues) (sas.QueryParameters, error) {
	_ = "STUB: not implemented"
	return *new(sas.QueryParameters), nil
}

func (s *clientTokenSigner) refreshUDC(ctx context.Context) (*service.UserDelegationCredential, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *clientTokenSigner) Sign(ctx context.Context, signatureValues *sas.BlobSignatureValues) (sas.QueryParameters, error) {
	_ = "STUB: not implemented"
	return *new(sas.QueryParameters), nil
}
