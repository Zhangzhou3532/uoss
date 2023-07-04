package aliyunoss

import (
	"context"
	"fmt"
	"gitlab.mihoyo.com/infosys/uoss/internal/uoss"
	"io/ioutil"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Client struct {
	c      *oss.Bucket
	bucket string
}

var (
	Endpoint string

	Bucket          string
	AccessKeyID     string
	AccessKeySecret string
)

// NewClient creates a new object storage client.
func NewClient() (*Client, error) {
	if !strings.Contains(Endpoint, "aliyuncs.com") {
		return nil, uoss.ErrTryNext
	}

	c, err := oss.New(
		Endpoint,
		AccessKeyID,
		AccessKeySecret,
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", uoss.ErrNewClient, err)
	}

	b, err := c.Bucket(Bucket)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", uoss.ErrNewClient, err)
	}

	return &Client{
		c:      b,
		bucket: Bucket,
	}, nil
}

// Put will create a new object and it will overwrite the original one if it exists already.
// Note that currently we can't use ctx to control timeout or cancel of Aliyun OSS requests.
func (c *Client) Put(name, data string) (err error) {
	return c.c.PutObject(name, strings.NewReader(data))
}

// Get downloads the object from oss.
// Note that currently we can't use ctx to control timeout or cancel of Aliyun OSS requests.
func (c *Client) Get(ctx context.Context, name string) ([]byte, error) {
	body, err := c.c.GetObject(name)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", uoss.ErrGetObject, err)
	}
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", uoss.ErrGetObject, err)
	}

	return data, nil
}

// GetObjectAsFile downloads the data to a local file.
// Note that currently we can't use ctx to control timeout or cancel of Aliyun OSS requests.
func (c *Client) GetObjectAsFile(ctx context.Context, name string, filePath string) error {
	err := c.c.GetObjectToFile(name, filePath)

	if err != nil {
		return fmt.Errorf("%w: %s", uoss.ErrGetObjectAsFile, err)
	}

	return nil
}

// ListObjectsOfCurrentDir list all files and sub directories of current directory.
// Note that currently we can't use ctx to control timeout or cancel of Aliyun OSS requests.
func (c *Client) ListObjectsOfCurrentDir(ctx context.Context, name string) ([]string, error) {
	marker := oss.Marker("")
	prefix := oss.Prefix(name)
	var objects = make([]string, 0, 100)
	for {
		l, err := c.c.ListObjects(marker, prefix, oss.Delimiter("/"))
		if err != nil {
			return objects, fmt.Errorf("%w: %s", uoss.ErrListObjectsOfCurrentDir, err)
		}
		for _, dirName := range l.CommonPrefixes {
			objects = append(objects, dirName)
		}
		for _, fileName := range l.Objects {
			objects = append(objects, fileName.Key)
		}
		if l.IsTruncated {
			marker = oss.Marker(l.NextMarker)
		} else {
			break
		}
	}
	return objects, nil
}
