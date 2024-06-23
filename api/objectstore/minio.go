package objectstore

import (
  "github.com/minio/minio-go/v7"
  "github.com/minio/minio-go/v7/pkg/credentials"
  "github.com/charmbracelet/log"
  "io"
  "os"
  "strconv"
  "context"
)

const (
  SALZ_DEFAULT_BUCKET = "salz"
)

type ObjectStoreHandler interface {
  UploadBotFile (botPath string, r io.Reader) error
}

type MinIOHandler struct {
  Client *minio.Client
  ObjectStoreEndpoint string
  BucketName string
}


func NewMinIOHandler() *MinIOHandler {
  minioEndpoint := os.Getenv("MINIO_ENDPOINT")
  minioAccessKey := os.Getenv("MINIO_ACCESS_KEY")
  minioSecretKey := os.Getenv("MINIO_SECRET_KEY")
  minioUseSSL := os.Getenv("MINIO_USE_SSL")

  if minioEndpoint == "" || minioAccessKey == "" || minioSecretKey == "" {
    log.Fatal("MinIO environment variables not set")
  }

  if minioUseSSL == "" {
    log.Debug("MINIO_USE_SSL not set, defaulting to true")
    minioUseSSL = "true"
  }

  minioSSL, err := strconv.ParseBool(minioUseSSL)
  if err != nil {
    log.Fatal("Invalid value for MINIO_USE_SSL", "MINIO_USE_SSL", minioUseSSL, "error", err)
  }

  minioClient, err := minio.New(minioEndpoint, &minio.Options{
    Creds:  credentials.NewStaticV4(minioAccessKey, minioSecretKey, ""),
    Secure: minioSSL,
  })
  if err != nil {
    log.Fatal("Unable to create MinIO client", err)
  }

  salzBucket := os.Getenv("SALZ_BUCKET")
  if salzBucket == "" {
    log.Debug("SALZ_BUCKET not set, using default", "default_bucket", SALZ_DEFAULT_BUCKET)
    salzBucket = SALZ_DEFAULT_BUCKET
  }

  ostore := &MinIOHandler{
    Client: minioClient,
    ObjectStoreEndpoint: minioEndpoint,
    BucketName: salzBucket,
  }

  // Check if the bucket exists -- throw if not
  exists, err := MinIOBucketExists(context.Background(), minioClient, salzBucket)
  if err != nil {
    log.Fatal("Error checking if bucket exists", "bucket", salzBucket, "error", err)
  }
  if !exists {
    log.Fatal("Bucket does not exist", "bucket", salzBucket)
  }
  
  return ostore
}

func MinIOBucketExists(ctx context.Context, client *minio.Client, bucketName string) (bool, error) {
  exists, err := client.BucketExists(ctx, bucketName)
  if err != nil {
    log.Error("Error checking if bucket exists", "bucket", bucketName, "error", err)
    return false, err
  }
  return exists, nil
}


func (o *MinIOHandler) UploadBotFile(botPath string, r io.Reader) error {
  // TODO: content type, content encoding, etc?
  info, err := o.Client.PutObject(context.Background(), o.BucketName, botPath, r, -1, minio.PutObjectOptions{})
  if err != nil {
    log.Error("Error uploading bot file", "bot_path", botPath, "error", err)
    return err
  }
  log.Info("Uploaded bot file", "bot_path", botPath, "size", info.Size)
  return err
}
