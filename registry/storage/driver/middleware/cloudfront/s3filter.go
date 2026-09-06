package middleware

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"
)

const (
	defaultIPRangesURL = "https://ip-ranges.amazonaws.com/ip-ranges.json"

	defaultUpdateFrequency = time.Hour * 12
)

func newAWSIPs(ctx context.Context, host string, updateFrequency time.Duration, awsRegion []string) (*awsIPs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type awsIPs struct {
	host            string
	updateFrequency time.Duration
	ipv4            []net.IPNet
	ipv6            []net.IPNet
	mutex           sync.RWMutex
	awsRegion       []string
	updaterStopChan chan bool
	initialized     bool
}

type awsIPResponse struct {
	Prefixes   []prefixEntry `json:"prefixes"`
	V6Prefixes []prefixEntry `json:"ipv6_prefixes"`
}

type prefixEntry struct {
	IPV4Prefix string `json:"ip_prefix"`
	IPV6Prefix string `json:"ipv6_prefix"`
	Region     string `json:"region"`
	Service    string `json:"service"`
}

func fetchAWSIPs(ctx context.Context, url string) (awsIPResponse, error) {
	_ = "STUB: not implemented"
	return *new(awsIPResponse), nil
}

func (s *awsIPs) tryUpdate(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *awsIPs) updater() { _ = "STUB: not implemented"; return }

func (s *awsIPs) getCandidateNetworks(ip net.IP) []net.IPNet { _ = "STUB: not implemented"; return nil }

func (s *awsIPs) contains(ip net.IP) bool { _ = "STUB: not implemented"; return false }

func parseIPFromRequest(request *http.Request) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

func eligibleForS3(request *http.Request, awsIPs *awsIPs) bool {
	_ = "STUB: not implemented"
	return false
}
