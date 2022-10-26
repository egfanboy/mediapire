package media

import (
	"context"

	"github.com/egfanboy/mediapire-manager/internal/node"
	"github.com/rs/zerolog/log"

	"github.com/egfanboy/mediapire-media-host/pkg/api"
	mhApi "github.com/egfanboy/mediapire-media-host/pkg/api"
	"github.com/egfanboy/mediapire-media-host/pkg/types"
)

type mediaApi interface {
	GetMedia(ctx context.Context) (map[string][]types.MediaItem, error)
	StreamMedia(ctx context.Context, nodeId string, filePath string) ([]byte, error)
}

type mediaService struct {
	nodeRepo node.NodeRepo
}

func (s *mediaService) GetMedia(ctx context.Context) (result map[string][]types.MediaItem, err error) {
	log.Info().Msg("Getting all media from all nodes")
	result = map[string][]types.MediaItem{}

	nodes, err := s.nodeRepo.GetAllNodes(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get all nodes")
		return
	}

	for _, node := range nodes {

		items, _, err2 := api.NewClient(ctx).GetMedia(node)
		if err2 != nil {
			err = err2
			log.Error().Err(err).Msgf("Failed to get media from node %s", node.NodeHost)
			return
		}

		result[node.Host()] = items
	}

	return
}

func (s *mediaService) StreamMedia(ctx context.Context, nodeId string, filePath string) ([]byte, error) {
	log.Info().Msgf("Streaming media %s from node %s", filePath, nodeId)
	node, err := s.nodeRepo.GetNode(ctx, nodeId)

	if err != nil {
		log.Error().Err(err).Msgf("Failed to get node with id %s", nodeId)
		return nil, err
	}

	client := mhApi.NewClient(ctx)

	b, _, err := client.StreamMedia(node, filePath)

	if err != nil {
		log.Error().Err(err).Msgf("Failed stream media on node %s", nodeId)
	}

	return b, err
}

func newMediaService() (mediaApi, error) {
	repo, err := node.NewNodeRepo()

	if err != nil {
		return nil, err
	}

	return &mediaService{
		nodeRepo: repo,
	}, nil
}
