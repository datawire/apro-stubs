package testutil

import (
	"net"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/datawire/dlib/dlog"
)

// Start a Redis instance for tests.  Return a shutdown function for
// Redis.  Error handling is done by calling t.Fatalf().
func StartRedis(t testing.TB, sockname string) func() {
	t.Helper()

	if _, err := os.Stat(sockname); err == nil {
		t.Fatalf("socket already exists: %q", sockname)
	} else if !os.IsNotExist(err) {
		t.Fatalf("checking if socket exists: %q: %v", sockname, err)
	}

	ctx := dlog.NewTestContext(t, true)

	cmdline := []string{"redis-server",
		"--loglevel", "warning",
		"--port", "0", // don't listen on TCP
		"--unixsocket", sockname,
	}
	cmd := exec.CommandContext(ctx, cmdline[0], cmdline[1:]...)
	cmd.Stdout = dlog.StdLogger(ctx, dlog.LogLevelDebug).Writer()
	cmd.Stderr = cmd.Stdout

	if err := cmd.Start(); err != nil {
		t.Fatalf("%q: %v", cmdline, err)
	}

	var waitErr error
	waitDone := make(chan struct{})
	go func() {
		waitErr = cmd.Wait()
		close(waitDone)
	}()

	// Wait for it to become ready
	for {
		conn, err := new(net.Dialer).DialContext(ctx, "unix", sockname)
		if err == nil {
			conn.Close()
			break
		}
		select {
		case <-waitDone:
			t.Fatalf("%q: %v", cmdline, waitErr)
		case <-ctx.Done():
			t.Fatalf("%q: %v", cmdline, ctx.Err())
		default:
		}
		time.Sleep(time.Second / 100)
	}

	// Return a shutdown function
	return func() {
		// Tell Redis to shut down
		cmd.Process.Signal(os.Interrupt)
		// Wait for Redis to shut down
		<-waitDone
		if waitErr != nil {
			t.Fatalf("%q: %v", cmdline, waitErr)
		}
	}
}
