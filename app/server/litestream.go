package server

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"

	litestream "github.com/supersupersimple/litestream-lib"
	lss3 "github.com/supersupersimple/litestream-lib/s3"
)

const (
	EnvLitestreamUrl             = "LITESTREAM_URL"
	EnvLitestreamAccessKeyID     = "LITESTREAM_ACCESS_KEY_ID"
	EnvLitestreamSecretAccessKey = "LITESTREAM_SECRET_ACCESS_KEY"
)

func ParseReplicaURL(s string) (scheme, host, urlpath string, err error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", "", "", err
	}

	switch u.Scheme {
	case "file":
		scheme, u.Scheme = u.Scheme, ""
		return scheme, "", path.Clean(u.String()), nil

	case "":
		return u.Scheme, u.Host, u.Path, fmt.Errorf("replica url scheme required: %s", s)

	default:
		return u.Scheme, u.Host, strings.TrimPrefix(path.Clean(u.Path), "/"), nil
	}
}

func replicate(ctx context.Context, dsn string) (*litestream.DB, error) {
	// Create Litestream DB reference for managing replication.
	lsdb := litestream.NewDB(dsn)
	lsdb.SetDriverName("sqlite")

	// Build S3 replica and attach to database.
	client := lss3.NewReplicaClient()
	url := os.Getenv(EnvLitestreamUrl)
	_, host, path, err := ParseReplicaURL(url)
	if err != nil {
		return nil, err
	}
	client.Path = path
	client.Bucket, client.Region, client.Endpoint, client.ForcePathStyle = lss3.ParseHost(host)
	client.AccessKeyID = os.Getenv(EnvLitestreamAccessKeyID)
	client.SecretAccessKey = os.Getenv(EnvLitestreamSecretAccessKey)

	replica := litestream.NewReplica(lsdb, "s3")
	replica.Client = client

	lsdb.Replicas = append(lsdb.Replicas, replica)

	if err := restore(ctx, replica); err != nil {
		return nil, err
	}

	// Initialize database.
	if err := lsdb.Open(); err != nil {
		return nil, err
	}

	return lsdb, nil
}

func restore(ctx context.Context, replica *litestream.Replica) (err error) {
	// Skip restore if local database already exists.
	if _, err := os.Stat(replica.DB().Path()); err == nil {
		fmt.Println("local database already exists, skipping restore")
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}

	// Configure restore to write out to DSN path.
	opt := litestream.NewRestoreOptions()
	opt.DriverName = "sqlite"
	opt.OutputPath = replica.DB().Path()

	// Determine the latest generation to restore from.
	if opt.Generation, _, err = replica.CalcRestoreTarget(ctx, opt); err != nil {
		return err
	}

	// Only restore if there is a generation available on the replica.
	// Otherwise we'll let the application create a new database.
	if opt.Generation == "" {
		fmt.Println("no generation found, creating new database")
		return nil
	}

	fmt.Printf("restoring replica for generation %s\n", opt.Generation)
	if err := replica.Restore(ctx, opt); err != nil {
		return err
	}
	fmt.Println("restore complete")
	return nil
}
