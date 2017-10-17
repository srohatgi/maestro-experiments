package docker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	imageName      = "testimage"
	contextDir     = "."
	dockerfilePath = "Dockerfile-test"
)

func TestBuildImage(t *testing.T) { // this test will actually create an image on your machine and then remove it
	dockerClient := NewDockerClient()
	numImages, err := dockerClient.getNumImages()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		dockerClient.removeImage(imageName)
	}()
	output, err := dockerClient.BuildImage(imageName, contextDir, dockerfilePath)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(*output)
	newNumImages, err := dockerClient.getNumImages()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, *numImages+1, *newNumImages)
}
