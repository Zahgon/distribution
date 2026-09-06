//go:build arm || 386

package s3

import "math"

const maxChunkSize = math.MaxInt32
