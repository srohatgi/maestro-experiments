package docker

import (
	"bytes"
	"github.com/fsouza/go-dockerclient"
)

const (
	dockerSock = "unix:///var/run/docker.sock"
)

type DockerClient struct {
	client *docker.Client
}

func NewDockerClient() *DockerClient {
	client, err := docker.NewClient(dockerSock)
	if err != nil {
		panic(err)
	}
	return &DockerClient{
		client: client,
	}
}

func (c *DockerClient) BuildImage(imageName string, contextDir string, dockerfilePath string) (*string, error) {
	var buf bytes.Buffer
	err := c.client.BuildImage(docker.BuildImageOptions{
		Name:           imageName,
		ContextDir:     contextDir,
		Dockerfile:     dockerfilePath,
		RmTmpContainer: true,
		OutputStream:   &buf,
		SuppressOutput: false,
	})
	if err != nil {
		return nil, err
	}
	output := buf.String()
	return &output, nil
}

func (c *DockerClient) getNumImages() (*int, error) {
	images, err := c.client.ListImages(docker.ListImagesOptions{All: false})
	if err != nil {
		return nil, err
	}
	numImages := len(images)
	return &numImages, nil
}

func (c *DockerClient) removeImage(imageName string) {
	c.client.RemoveImage(imageName)
}
