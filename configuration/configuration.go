package configuration

import (
	"io"
	"net/http"
	"time"
)

const (
	defaultMaxEntries = 1000

	defaultMaxTags = 1000
)

type Configuration struct {
	Version Version `yaml:"version"`

	Log Log `yaml:"log"`

	Loglevel Loglevel `yaml:"loglevel,omitempty"`

	Storage Storage `yaml:"storage"`

	Auth Auth `yaml:"auth,omitempty"`

	Middleware map[string][]Middleware `yaml:"middleware,omitempty"`

	HTTP HTTP `yaml:"http,omitempty"`

	Notifications Notifications `yaml:"notifications,omitempty"`

	Redis Redis `yaml:"redis,omitempty"`

	Health Health `yaml:"health,omitempty"`

	Catalog Catalog `yaml:"catalog,omitempty"`

	Tags Tags `yaml:"tags,omitempty"`

	Proxy Proxy `yaml:"proxy,omitempty"`

	Validation Validation `yaml:"validation,omitempty"`

	Policy Policy `yaml:"policy,omitempty"`
}

type Policy struct {
	Repository Repository `yaml:"repository,omitempty"`
}

type Repository struct {
	Classes []string `yaml:"classes"`
}

type Catalog struct {
	MaxEntries int `yaml:"maxentries,omitempty"`
}

type Log struct {
	AccessLog AccessLog `yaml:"accesslog,omitempty"`

	Level Loglevel `yaml:"level,omitempty"`

	Formatter string `yaml:"formatter,omitempty"`

	Fields map[string]any `yaml:"fields,omitempty"`

	Hooks []LogHook `yaml:"hooks,omitempty"`

	ReportCaller bool `yaml:"reportcaller,omitempty"`
}

type AccessLog struct {
	Disabled bool `yaml:"disabled,omitempty"`
}

type HTTP struct {
	Addr string `yaml:"addr,omitempty"`

	Net string `yaml:"net,omitempty"`

	Host string `yaml:"host,omitempty"`

	Prefix string `yaml:"prefix,omitempty"`

	Secret string `yaml:"secret,omitempty"`

	RelativeURLs bool `yaml:"relativeurls,omitempty"`

	DrainTimeout time.Duration `yaml:"draintimeout,omitempty"`

	TLS TLS `yaml:"tls,omitempty"`

	Headers http.Header `yaml:"headers,omitempty"`

	Debug Debug `yaml:"debug,omitempty"`

	HTTP2 HTTP2 `yaml:"http2,omitempty"`

	H2C H2C `yaml:"h2c,omitempty"`
}

type Debug struct {
	Addr string `yaml:"addr,omitempty"`

	Prometheus Prometheus `yaml:"prometheus,omitempty"`
}

type Prometheus struct {
	Enabled bool `yaml:"enabled,omitempty"`

	Path string `yaml:"path,omitempty"`
}

type HTTP2 struct {
	Disabled bool `yaml:"disabled,omitempty"`
}

type H2C struct {
	Enabled bool `yaml:"enabled,omitempty"`
}

type TLS struct {
	Certificate string `yaml:"certificate,omitempty"`

	Key string `yaml:"key,omitempty"`

	ClientCAs []string `yaml:"clientcas,omitempty"`

	ClientAuth ClientAuth `yaml:"clientauth,omitempty"`

	MinimumTLS string `yaml:"minimumtls,omitempty"`

	CipherSuites []string `yaml:"ciphersuites,omitempty"`

	LetsEncrypt LetsEncrypt `yaml:"letsencrypt,omitempty"`
}

type LetsEncrypt struct {
	CacheFile string `yaml:"cachefile,omitempty"`

	Email string `yaml:"email,omitempty"`

	Hosts []string `yaml:"hosts,omitempty"`

	DirectoryURL string `yaml:"directoryurl,omitempty"`
}

type Tags struct {
	MaxTags int `yaml:"maxtags,omitempty"`
}

type LogHook struct {
	Disabled bool `yaml:"disabled,omitempty"`

	Type string `yaml:"type,omitempty"`

	Levels []string `yaml:"levels,omitempty"`

	MailOptions MailOptions `yaml:"options,omitempty"`
}

type MailOptions struct {
	SMTP SMTP `yaml:"smtp,omitempty"`

	From string `yaml:"from,omitempty"`

	To []string `yaml:"to,omitempty"`
}

type SMTP struct {
	Addr string `yaml:"addr,omitempty"`

	Username string `yaml:"username,omitempty"`

	Password string `yaml:"password,omitempty"`

	Insecure bool `yaml:"insecure,omitempty"`
}

type FileChecker struct {
	Interval time.Duration `yaml:"interval,omitempty"`

	File string `yaml:"file,omitempty"`

	Threshold int `yaml:"threshold,omitempty"`
}

type HTTPChecker struct {
	Timeout time.Duration `yaml:"timeout,omitempty"`

	StatusCode int

	Interval time.Duration `yaml:"interval,omitempty"`

	URI string `yaml:"uri,omitempty"`

	Headers http.Header `yaml:"headers"`

	Threshold int `yaml:"threshold,omitempty"`
}

type TCPChecker struct {
	Timeout time.Duration `yaml:"timeout,omitempty"`

	Interval time.Duration `yaml:"interval,omitempty"`

	Addr string `yaml:"addr,omitempty"`

	Threshold int `yaml:"threshold,omitempty"`
}

type Health struct {
	FileCheckers []FileChecker `yaml:"file,omitempty"`

	HTTPCheckers []HTTPChecker `yaml:"http,omitempty"`

	TCPCheckers []TCPChecker `yaml:"tcp,omitempty"`

	StorageDriver StorageDriver `yaml:"storagedriver,omitempty"`
}

type StorageDriver struct {
	Enabled bool `yaml:"enabled,omitempty"`

	Interval time.Duration `yaml:"interval,omitempty"`

	Threshold int `yaml:"threshold,omitempty"`
}

type Platform struct {
	Architecture string `yaml:"architecture,omitempty"`

	OS string `yaml:"os,omitempty"`
}

type v0_1Configuration Configuration

func (version *Version) UnmarshalYAML(unmarshal func(any) error) error {
	_ = "STUB: not implemented"
	return nil
}

var CurrentVersion = MajorMinorVersion(0, 1)

type Loglevel string

func (loglevel *Loglevel) UnmarshalYAML(unmarshal func(any) error) error {
	_ = "STUB: not implemented"
	return nil
}

type Parameters map[string]any

type Storage map[string]Parameters

func (storage Storage) Type() string { _ = "STUB: not implemented"; return "" }

func (storage Storage) TagParameters() Parameters {
	_ = "STUB: not implemented"
	return *new(Parameters)
}

func (storage Storage) setTagParameter(key string, value any) { _ = "STUB: not implemented"; return }

func (storage Storage) Parameters() Parameters { _ = "STUB: not implemented"; return *new(Parameters) }

func (storage Storage) setParameter(key string, value any) { _ = "STUB: not implemented"; return }

func (storage *Storage) UnmarshalYAML(unmarshal func(any) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (storage Storage) MarshalYAML() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

type Auth map[string]Parameters

func (auth Auth) Type() string { _ = "STUB: not implemented"; return "" }

func (auth Auth) Parameters() Parameters { _ = "STUB: not implemented"; return *new(Parameters) }

func (auth Auth) setParameter(key string, value any) { _ = "STUB: not implemented"; return }

func (auth *Auth) UnmarshalYAML(unmarshal func(any) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (auth Auth) MarshalYAML() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

type Notifications struct {
	EventConfig Events `yaml:"events,omitempty"`

	Endpoints []Endpoint `yaml:"endpoints,omitempty"`
}

type Endpoint struct {
	Name              string        `yaml:"name"`
	Disabled          bool          `yaml:"disabled"`
	URL               string        `yaml:"url"`
	Headers           http.Header   `yaml:"headers"`
	Timeout           time.Duration `yaml:"timeout"`
	Threshold         int           `yaml:"threshold"`
	Backoff           time.Duration `yaml:"backoff"`
	IgnoredMediaTypes []string      `yaml:"ignoredmediatypes"`
	Ignore            Ignore        `yaml:"ignore"`
}

type Events struct {
	IncludeReferences bool `yaml:"includereferences"`
}

type Ignore struct {
	MediaTypes []string `yaml:"mediatypes"`
	Actions    []string `yaml:"actions"`
}

type Middleware struct {
	Name string `yaml:"name"`

	Disabled bool `yaml:"disabled,omitempty"`

	Options Parameters `yaml:"options"`
}

type Proxy struct {
	RemoteURL string `yaml:"remoteurl"`

	Username string `yaml:"username"`

	Password string `yaml:"password"`

	Exec *ExecConfig `yaml:"exec,omitempty"`

	TTL *time.Duration `yaml:"ttl,omitempty"`

	CacheWriteTimeout *time.Duration `yaml:"cachewritetimeout,omitempty"`
}

type ExecConfig struct {
	Command string `yaml:"command"`

	Lifetime *time.Duration `yaml:"lifetime,omitempty"`
}

type Validation struct {
	Enabled bool `yaml:"enabled,omitempty"`

	Disabled bool `yaml:"disabled,omitempty"`

	Manifests ValidationManifests `yaml:"manifests,omitempty"`
}

type ValidationManifests struct {
	URLs URLs `yaml:"urls,omitempty"`

	Indexes ValidationIndexes `yaml:"indexes,omitempty"`
}

type URLs struct {
	Allow []string `yaml:"allow,omitempty"`

	Deny []string `yaml:"deny,omitempty"`
}

type ValidationIndexes struct {
	Platforms Platforms `yaml:"platforms"`

	PlatformList []Platform `yaml:"platformlist,omitempty"`
}

type Platforms string

func (platforms *Platforms) UnmarshalYAML(unmarshal func(any) error) error {
	_ = "STUB: not implemented"
	return nil
}

func Parse(rd io.Reader) (*Configuration, error) { _ = "STUB: not implemented"; return nil, nil }

type RedisOptions struct {
	Addrs []string `yaml:"addrs,omitempty"`

	ClientName string `yaml:"clientname,omitempty"`

	DB int `yaml:"db,omitempty"`

	Protocol int `yaml:"protocol,omitempty"`

	Username string `yaml:"username,omitempty"`

	Password string `yaml:"password,omitempty"`

	SentinelUsername string `yaml:"sentinelusername,omitempty"`

	SentinelPassword string `yaml:"sentinelpassword,omitempty"`

	MaxRetries int `yaml:"maxretries,omitempty"`

	MinRetryBackoff time.Duration `yaml:"minretrybackoff,omitempty"`

	MaxRetryBackoff time.Duration `yaml:"maxretrybackoff,omitempty"`

	DialTimeout time.Duration `yaml:"dialtimeout,omitempty"`

	ReadTimeout time.Duration `yaml:"readtimeout,omitempty"`

	WriteTimeout time.Duration `yaml:"writetimeout,omitempty"`

	ContextTimeoutEnabled bool `yaml:"contexttimeoutenabled,omitempty"`

	PoolFIFO bool `yaml:"poolfifo,omitempty"`

	PoolSize int `yaml:"poolsize,omitempty"`

	PoolTimeout time.Duration `yaml:"pooltimeout,omitempty"`

	MinIdleConns int `yaml:"minidleconns,omitempty"`

	MaxIdleConns int `yaml:"maxidleconns,omitempty"`

	MaxActiveConns int `yaml:"maxactiveconns,omitempty"`

	ConnMaxIdleTime time.Duration `yaml:"connmaxidletime,omitempty"`

	ConnMaxLifetime time.Duration `yaml:"connmaxlifetime,omitempty"`

	MaxRedirects int `yaml:"maxredirects,omitempty"`

	ReadOnly bool `yaml:"readonly,omitempty"`

	RouteByLatency bool `yaml:"routebylatency,omitempty"`

	RouteRandomly bool `yaml:"routerandomly,omitempty"`

	MasterName string `yaml:"mastername,omitempty"`

	DisableIdentity bool `yaml:"disableidentity,omitempty"`

	IdentitySuffix string `yaml:"identitysuffix,omitempty"`

	UnstableResp3 bool `yaml:"unstableresp3,omitempty"`
}

type RedisTLSOptions struct {
	Certificate string `yaml:"certificate,omitempty"`

	Key string `yaml:"key,omitempty"`

	RootCAs []string `yaml:"rootcas,omitempty"`
}

type Redis struct {
	Options RedisOptions `yaml:",inline"`

	TLS RedisTLSOptions `yaml:"tls,omitempty"`
}

const (
	ClientAuthRequestClientCert          = "request-client-cert"
	ClientAuthRequireAnyClientCert       = "require-any-client-cert"
	ClientAuthVerifyClientCertIfGiven    = "verify-client-cert-if-given"
	ClientAuthRequireAndVerifyClientCert = "require-and-verify-client-cert"
)

type ClientAuth string

func (clientAuth *ClientAuth) UnmarshalYAML(unmarshal func(any) error) error {
	_ = "STUB: not implemented"
	return nil
}
