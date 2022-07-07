package streams

import (
	"bufio"
	"context"
	"errors"
	"io"
	"sync"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

// Reader is a wrapper around an io.Reader interface
type Reader struct {
	mutex      sync.Mutex
	ctx        context.Context
	forceClose func()
	reader     io.Reader
	readCloser io.ReadCloser
	bRead      uint64
	bWritten   uint64
	dependents int
	dataType   string
	dtLock     sync.Mutex
}

// NewReader creates a new Stdio.Io interface wrapper around a io.Reader interface
func NewReader(reader io.Reader) (r *Reader) {
	if reader == nil {
		panic("streams.Reader interface has nil reader interface")
	}

	r = new(Reader)
	r.reader = reader
	r.ctx, r.forceClose = context.WithCancel(context.Background())
	return
}

// IsTTY returns false because the reader interface is not a pseudo-TTY
func (r *Reader) IsTTY() bool { return false }

// Stats provides real time stream stats. Useful for progress bars etc.
func (r *Reader) Stats() (bytesWritten, bytesRead uint64) {
	//r.mutex.RLock()
	r.mutex.Lock()
	bytesWritten = r.bWritten
	bytesRead = r.bRead
	//r.mutex.RUnlock()
	r.mutex.Unlock()
	return
}

// Read is the reader interface Read() method.
func (r *Reader) Read(p []byte) (int, error) {
	select {
	case <-r.ctx.Done():
		return 0, io.EOF
	default:
	}

	r.mutex.Lock()
	i, err := r.reader.Read(p)
	r.bRead += uint64(i)
	r.mutex.Unlock()

	return i, err
}

// ReadLine returns each line in the stream as a callback function
func (r *Reader) ReadLine(callback func([]byte)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		b := scanner.Bytes()
		//r.mutex.Lock()
		//r.bRead += uint64(len(b))
		//r.mutex.Unlock()
		callback(append(b, utils.NewLineByte...))
	}

	return scanner.Err()
}

// ReadAll reads everything and dump it into one byte slice.
func (r *Reader) ReadAll() (b []byte, err error) {
	w := NewStdinWithContext(r.ctx, r.forceClose)

	_, err = w.ReadFrom(r.reader)
	if err != nil {
		return
	}

	b, err = w.ReadAll()

	r.mutex.Lock()
	r.bRead = uint64(len(b))
	r.mutex.Unlock()

	return b, err
}

// ReadArray returns a data type-specific array returned via a callback function
func (r *Reader) ReadArray(callback func([]byte)) error {
	return stdio.ReadArray(r, callback)
}

// ReadArrayWithType returns a data type-specific array returned via a callback function
func (r *Reader) ReadArrayWithType(callback func([]byte, string)) error {
	return stdio.ReadArrayWithType(r, callback)
}

// ReadMap returns a data type-specific key/values returned via a callback function
func (r *Reader) ReadMap(config *config.Config, callback func(key, value string, last bool)) error {
	return stdio.ReadMap(r, config, callback)
}

// Write is a dummy function because it's a reader interface
func (r *Reader) Write(p []byte) (int, error) {
	return 0, errors.New("Cannot write to a reader interface")
}

// Writeln is a dummy function because it's a reader interface
func (r *Reader) Writeln(b []byte) (int, error) {
	return 0, errors.New("Cannot write to a reader interface")
}

// WriteArray is a dummy function because it's a reader interface
func (r *Reader) WriteArray(dataType string) (stdio.ArrayWriter, error) {
	return nil, errors.New("Cannot write to a reader interface")
}

// Open the stream.Io interface for another dependant
func (r *Reader) Open() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.dependents++
}

// Close the stream.Io interface
func (r *Reader) Close() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.dependents--

	if r.dependents < 0 {
		panic("More closed dependents than open")
	}

	if r.dependents == 0 && r.readCloser != nil {
		r.readCloser.Close()
	}
}

// ForceClose forces the stream.Io interface to close. This should only be called by a STDIN reader
func (r *Reader) ForceClose() {
	r.forceClose()
}

// WriteTo reads from the stream.Io interface and writes to a destination
// io.Writer interface
func (r *Reader) WriteTo(w io.Writer) (int64, error) {
	return stdio.WriteTo(r, w)
}

// GetDataType returns the murex data type for the stream.Io interface
func (r *Reader) GetDataType() (dt string) {
	for {
		select {
		case <-r.ctx.Done():
			return types.Generic
		default:
		}

		r.dtLock.Lock()
		dt = r.dataType
		r.dtLock.Unlock()
		if dt != "" {
			return
		}
	}
}

// SetDataType defines the murex data type for the stream.Io interface
func (r *Reader) SetDataType(dt string) {
	r.dtLock.Lock()
	r.dataType = dt
	r.dtLock.Unlock()
}

// DefaultDataType defines the murex data type for the stream.Io interface if
// it's not already set
func (r *Reader) DefaultDataType(err bool) {
	r.dtLock.Lock()
	dt := r.dataType
	r.dtLock.Unlock()

	if dt == "" {
		if err {
			r.dtLock.Lock()
			r.dataType = types.Null
			r.dtLock.Unlock()
		} else {
			r.dtLock.Lock()
			r.dataType = types.Generic
			r.dtLock.Unlock()
		}
	}
}
