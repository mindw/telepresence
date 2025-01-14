package socket_test

import (
	"context"
	"io/fs"
	"net"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"github.com/datawire/dlib/dgroup"
	"github.com/datawire/dlib/dhttp"
	"github.com/datawire/dlib/dlog"
	"github.com/telepresenceio/telepresence/v2/pkg/client/socket"
)

func TestDialSocket(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		sockname := filepath.Join(t.TempDir(), "ok.sock")
		listener, err := net.Listen("unix", sockname)
		if !assert.NoError(t, err) {
			return
		}
		defer listener.Close()

		ctx := dlog.NewTestContext(t, false)
		grp := dgroup.NewGroup(ctx, dgroup.GroupConfig{
			EnableWithSoftness: true,
			ShutdownOnNonError: true,
			DisableLogging:     true,
		})

		grp.Go("server", func(ctx context.Context) error {
			sc := &dhttp.ServerConfig{
				Handler: grpc.NewServer(),
			}
			return sc.Serve(ctx, listener)
		})

		grp.Go("client", func(ctx context.Context) error {
			conn, err := socket.Dial(ctx, sockname, true)
			assert.NoError(t, err)
			if assert.NotNil(t, conn) {
				assert.NoError(t, conn.Close())
			}
			return nil
		})

		assert.NoError(t, grp.Wait())
	})
	t.Run("Hang", func(t *testing.T) {
		sockname := filepath.Join(t.TempDir(), "hang.sock")
		listener, err := net.Listen("unix", sockname)
		if !assert.NoError(t, err) {
			return
		}
		defer listener.Close()

		ctx := dlog.NewTestContext(t, false)
		conn, err := socket.Dial(ctx, sockname, true)
		assert.Nil(t, conn)
		assert.Error(t, err)
		t.Log(err)
		assert.ErrorIs(t, err, fs.ErrNotExist)
		assert.Contains(t, err.Error(), "dial unix "+sockname)
		assert.Contains(t, err.Error(), "this usually means that the process has locked up")
	})
	t.Run("NotExist", func(t *testing.T) {
		ctx := dlog.NewTestContext(t, false)
		sockname := filepath.Join(t.TempDir(), "not-exist.sock")
		conn, err := socket.Dial(ctx, sockname, true)
		assert.Nil(t, conn)
		assert.Error(t, err)
		t.Log(err)
		assert.ErrorIs(t, err, os.ErrNotExist)
		assert.Contains(t, err.Error(), "dial unix "+sockname)
		assert.Contains(t, err.Error(), "this usually means that the process is not running")
	})
}
