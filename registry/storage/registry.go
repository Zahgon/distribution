package storage

import (
	"context"
	"regexp"
	"runtime"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/registry/storage/cache"
	storagedriver "github.com/distribution/distribution/v3/registry/storage/driver"
	"github.com/distribution/reference"
)

var (
	DefaultConcurrencyLimit = runtime.GOMAXPROCS(0)
)

type registry struct {
	blobStore                    *blobStore
	blobServer                   *blobServer
	statter                      *blobStatter
	blobDescriptorCacheProvider  cache.BlobDescriptorCacheProvider
	deleteEnabled                bool
	tagLookupConcurrencyLimit    int
	resumableDigestEnabled       bool
	blobDescriptorServiceFactory distribution.BlobDescriptorServiceFactory
	driver                       storagedriver.StorageDriver

	manifestURLs         manifestURLs
	validateImageIndexes validateImageIndexes
}

type manifestURLs struct {
	allow *regexp.Regexp
	deny  *regexp.Regexp
}

type validateImageIndexes struct {
	imagesExist bool

	imagePlatforms []platform
}

type platform struct {
	architecture string
	os           string
}

type RegistryOption func(*registry) error

func EnableRedirect(registry *registry) error { _ = "STUB: not implemented"; return nil }

func TagLookupConcurrencyLimit(concurrencyLimit int) RegistryOption {
	_ = "STUB: not implemented"
	return *new(RegistryOption)
}

func EnableDelete(registry *registry) error { _ = "STUB: not implemented"; return nil }

func DisableDigestResumption(registry *registry) error { _ = "STUB: not implemented"; return nil }

func ManifestURLsAllowRegexp(r *regexp.Regexp) RegistryOption {
	_ = "STUB: not implemented"
	return *new(RegistryOption)
}

func ManifestURLsDenyRegexp(r *regexp.Regexp) RegistryOption {
	_ = "STUB: not implemented"
	return *new(RegistryOption)
}

func EnableValidateImageIndexImagesExist(registry *registry) error {
	_ = "STUB: not implemented"
	return nil
}

func AddValidateImageIndexImagesExistPlatform(architecture string, os string) RegistryOption {
	_ = "STUB: not implemented"
	return *new(RegistryOption)
}

func BlobDescriptorServiceFactory(factory distribution.BlobDescriptorServiceFactory) RegistryOption {
	_ = "STUB: not implemented"
	return *new(RegistryOption)
}

func BlobDescriptorCacheProvider(blobDescriptorCacheProvider cache.BlobDescriptorCacheProvider) RegistryOption {
	_ = "STUB: not implemented"
	return *new(RegistryOption)
}

func NewRegistry(ctx context.Context, driver storagedriver.StorageDriver, options ...RegistryOption) (distribution.Namespace, error) {
	_ = "STUB: not implemented"
	return *new(distribution.Namespace), nil
}

func (reg *registry) Scope() distribution.Scope {
	_ = "STUB: not implemented"
	return *new(distribution.Scope)
}

func (reg *registry) Repository(ctx context.Context, canonicalName reference.Named) (distribution.Repository, error) {
	_ = "STUB: not implemented"
	return *new(distribution.Repository), nil
}

func (reg *registry) Blobs() distribution.BlobEnumerator {
	_ = "STUB: not implemented"
	return *new(distribution.BlobEnumerator)
}

func (reg *registry) BlobStatter() distribution.BlobStatter {
	_ = "STUB: not implemented"
	return *new(distribution.BlobStatter)
}

type repository struct {
	*registry
	ctx             context.Context
	name            reference.Named
	descriptorCache distribution.BlobDescriptorService
}

func (repo *repository) Named() reference.Named {
	_ = "STUB: not implemented"
	return *new(reference.Named)
}

func (repo *repository) Tags(ctx context.Context) distribution.TagService {
	_ = "STUB: not implemented"
	return *new(distribution.TagService)
}

func (repo *repository) Manifests(ctx context.Context, options ...distribution.ManifestServiceOption) (distribution.ManifestService, error) {
	_ = "STUB: not implemented"
	return *new(distribution.ManifestService), nil
}

func (repo *repository) Blobs(ctx context.Context) distribution.BlobStore {
	_ = "STUB: not implemented"
	return *new(distribution.BlobStore)
}
