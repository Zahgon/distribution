package metrics

import "github.com/docker/go-metrics"

const (
	NamespacePrefix = "registry"
)

var (
	StorageNamespace = metrics.NewNamespace(NamespacePrefix, "storage", nil)

	NotificationsNamespace = metrics.NewNamespace(NamespacePrefix, "notifications", nil)

	ProxyNamespace = metrics.NewNamespace(NamespacePrefix, "proxy", nil)
)
