package testsuites

import (
	"context"
	crand "crypto/rand"
	"sync"
	"testing"

	storagedriver "github.com/distribution/distribution/v3/registry/storage/driver"
	"github.com/stretchr/testify/suite"
)

var randomBytes = make([]byte, 128<<20)

func init() {
	_, _ = crand.Read(randomBytes)
}

type DriverConstructor func() (storagedriver.StorageDriver, error)

type DriverTeardown func() error

type DriverSuite struct {
	suite.Suite
	Constructor DriverConstructor
	Teardown    DriverTeardown
	storagedriver.StorageDriver
	ctx        context.Context
	skipVerify bool
}

func Driver(t *testing.T, driverConstructor DriverConstructor, skipVerify bool) {
	_ = "STUB: not implemented"
	return
}

func (suite *DriverSuite) SetupSuite() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TearDownSuite() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TearDownTest() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestRootExists() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestValidPaths() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) deletePath(path string) { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestInvalidPaths() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestWriteRead1() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestWriteRead2() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestWriteRead3() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestWriteRead4() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestWriteReadNonUTF8() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestTruncate() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestReadNonexistent() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestWriteReadStreams1() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestWriteReadStreams2() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestWriteReadStreams3() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestWriteReadStreams4() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestWriteReadStreamsNonUTF8() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestWriteReadLargeStreams() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestReaderWithOffset() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestContinueStreamAppendLarge() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestContinueStreamAppendSmall() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) testContinueStreamAppend(chunkSize int64) {
	_ = "STUB: not implemented"
	return
}

func (suite *DriverSuite) TestReadNonexistentStream() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestWriteZeroByteStreamThenAppend() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestWriteZeroByteContentThenAppend() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestList() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestMove() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestMoveOverwrite() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestMoveNonexistent() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestMoveInvalid() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestDelete() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestRedirectURL() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestDeleteNonexistent() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestDeleteFolder() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestDeleteOnlyDeletesSubpaths() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestStatCall() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestPutContentMultipleTimes() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestConcurrentStreamReads() { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) TestConcurrentFileStreams() { _ = "STUB: not implemented"; return }

type DriverBenchmarkSuite struct {
	DriverSuite
}

func BenchDriver(b *testing.B, driverConstructor DriverConstructor) {
	_ = "STUB: not implemented"
	return
}

func (s *DriverBenchmarkSuite) BenchmarkPutGetEmptyFiles(b *testing.B) {
	_ = "STUB: not implemented"
	return
}

func (s *DriverBenchmarkSuite) BenchmarkPutGet1KBFiles(b *testing.B) {
	_ = "STUB: not implemented"
	return
}

func (s *DriverBenchmarkSuite) BenchmarkPutGet1MBFiles(b *testing.B) {
	_ = "STUB: not implemented"
	return
}

func (s *DriverBenchmarkSuite) BenchmarkPutGet1GBFiles(b *testing.B) {
	_ = "STUB: not implemented"
	return
}

func (s *DriverBenchmarkSuite) benchmarkPutGetFiles(b *testing.B, size int64) {
	_ = "STUB: not implemented"
	return
}

func (s *DriverBenchmarkSuite) BenchmarkStreamEmptyFiles(b *testing.B) {
	_ = "STUB: not implemented"
	return
}

func (s *DriverBenchmarkSuite) BenchmarkStream1KBFiles(b *testing.B) {
	_ = "STUB: not implemented"
	return
}

func (s *DriverBenchmarkSuite) BenchmarkStream1MBFiles(b *testing.B) {
	_ = "STUB: not implemented"
	return
}

func (s *DriverBenchmarkSuite) BenchmarkStream1GBFiles(b *testing.B) {
	_ = "STUB: not implemented"
	return
}

func (s *DriverBenchmarkSuite) benchmarkStreamFiles(b *testing.B, size int64) {
	_ = "STUB: not implemented"
	return
}

func (s *DriverBenchmarkSuite) BenchmarkList5Files(b *testing.B) { _ = "STUB: not implemented"; return }

func (s *DriverBenchmarkSuite) BenchmarkList50Files(b *testing.B) {
	_ = "STUB: not implemented"
	return
}

func (s *DriverBenchmarkSuite) benchmarkListFiles(b *testing.B, numFiles int64) {
	_ = "STUB: not implemented"
	return
}

func (s *DriverBenchmarkSuite) BenchmarkDelete5Files(b *testing.B) {
	_ = "STUB: not implemented"
	return
}

func (s *DriverBenchmarkSuite) BenchmarkDelete50Files(b *testing.B) {
	_ = "STUB: not implemented"
	return
}

func (s *DriverBenchmarkSuite) benchmarkDeleteFiles(b *testing.B, numFiles int64) {
	_ = "STUB: not implemented"
	return
}

func (suite *DriverSuite) testFileStreams(size int64) { _ = "STUB: not implemented"; return }

func (suite *DriverSuite) writeReadCompare(filename string, contents []byte) {
	_ = "STUB: not implemented"
	return
}

func (suite *DriverSuite) writeReadCompareStreams(filename string, contents []byte) {
	_ = "STUB: not implemented"
	return
}

var (
	filenameChars  = []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	separatorChars = []byte("-")
)

func randomPath(length int64) string { _ = "STUB: not implemented"; return "" }

func randomFilename(length int64) string { _ = "STUB: not implemented"; return "" }

func randomContents(length int64) []byte { _ = "STUB: not implemented"; return nil }

type randReader struct {
	r int64
	m sync.Mutex
}

func (rr *randReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func newRandReader(n int64) *randReader { _ = "STUB: not implemented"; return nil }

func firstPart(filePath string) string { _ = "STUB: not implemented"; return "" }
