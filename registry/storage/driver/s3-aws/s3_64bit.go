//go:build !arm && !386

package s3

const maxChunkSize = 5 * 1024 * 1024 * 1024
