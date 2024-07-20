package backend

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func listContainers(cli *client.Client) ([]string, error) {
	var results []string
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, ctr := range containers {
		var portsStr []string
		for _, p := range ctr.Ports {
			portsStr = append(portsStr, fmt.Sprintf("%d:%d", p.PublicPort, p.PrivatePort))
		}
		portsJoined := strings.Join(portsStr, ", ")

		if strings.HasPrefix(ctr.Image, "sha256:") {
			ctr.Image = ctr.Image[7:20]
		}

		containerInfo := fmt.Sprintf("%s %s %s %s %s %s %d\n", ctr.ID[:12], ctr.Image, ctr.Names[0], ctr.Command, portsJoined, ctr.Status, ctr.Created)
		results = append(results, containerInfo)
	}
	fmt.Println(results)
	return results, nil
}

func (b *Backend) ListContainers() ([]string, error) {
	return listContainers(b.cli)
}

func CreateContainer(ctx context.Context, cli *client.Client, imageName string) {
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
	}, nil, nil, nil, "")
	if err != nil {

		if strings.Contains(err.Error(), "No such image") {
			fmt.Printf("Image %s not found locally, pulling...\n", imageName)
			out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
			if err != nil {
				panic(err)
			}
			io.Copy(io.Discard, out)
			out.Close()

			resp, err = cli.ContainerCreate(ctx, &container.Config{
				Image: imageName,
			}, nil, nil, nil, "")
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		fmt.Printf("Error starting container: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Container created and started successfully")
}

func getContainerLogs(ctx context.Context, cli *client.Client, containerID string) (string, error) {
	out, err := cli.ContainerLogs(ctx, containerID, container.LogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)
	out.Close()

	logs, err := io.ReadAll(out)
	if err != nil {
		return "", err
	}
	return string(logs), nil
}

func (b *Backend) GetContainerLogs(containerID string) (string, error) {
	return getContainerLogs(b.ctx, b.cli, containerID)
}

func (b *Backend) RemoveImage(imageID string) error {
	_, err := b.cli.ImageRemove(b.ctx, imageID, image.RemoveOptions{})
	return err
}

func listImages(ctx context.Context, cli *client.Client) ([]string, error) {
	var results []string
	images, err := cli.ImageList(ctx, image.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	for _, img := range images {
		var name string
		if len(img.RepoDigests) > 0 {
			name = strings.Split(img.RepoDigests[0], "@")[0]
		} else {
			name = "<none>"
		}

		tags := "<none>"
		if len(img.RepoTags) > 0 {
			tags = strings.Join(img.RepoTags, ", ")
		}

		imageInfo := fmt.Sprintf("%s %s %s %d %d", name, img.ID[7:], tags, img.Created, img.Size)
		results = append(results, imageInfo)
	}
	return results, nil
}

func (b *Backend) ListImages() ([]string, error) {
	return listImages(b.ctx, b.cli)
}
