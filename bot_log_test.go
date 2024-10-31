package contracts

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBotLog(t *testing.T) {
	t.Run("JSONMarshalSingleBotLog", func(t *testing.T) {
		releaseLog := BotLog{
			ID:         "5",
			RepoSource: "github.com",
			RepoOwner:  "estafette",
			RepoName:   "estafette-ci-api",
			BotID:      "123445",
			Steps: []*BuildLogStep{
				&BuildLogStep{
					Step: "deploy",
					Image: &BuildLogStepDockerImage{
						Name:         "golang",
						Tag:          "1.10.2-alpine3.7",
						IsPulled:     false,
						ImageSize:    135000,
						PullDuration: 2 * time.Second,
						Error:        "",
					},
					Duration: 91 * time.Second,
					LogLines: []BuildLogLine{
						BuildLogLine{
							Timestamp:  time.Date(2018, 4, 17, 8, 3, 0, 0, time.UTC),
							StreamType: "stdout",
							Text: "ok  	github.com/estafette/estafette-ci-contracts	0.017s",
						},
					},
					ExitCode: 0,
					Status:   "SUCCEEDED",
				},
			},
			InsertedAt: time.Date(2018, 4, 17, 8, 3, 0, 0, time.UTC),
		}
		// act
		bytes, err := json.Marshal(&releaseLog)
		assert.Nil(t, err)
		assert.Equal(t, "{\"id\":\"5\",\"repoSource\":\"github.com\",\"repoOwner\":\"estafette\",\"repoName\":\"estafette-ci-api\",\"botID\":\"123445\",\"steps\":[{\"step\":\"deploy\",\"image\":{\"name\":\"golang\",\"tag\":\"1.10.2-alpine3.7\",\"isPulled\":false,\"imageSize\":135000,\"pullDuration\":2000000000},\"duration\":91000000000,\"logLines\":[{\"timestamp\":\"2018-04-17T08:03:00Z\",\"streamType\":\"stdout\",\"text\":\"ok  \\tgithub.com/estafette/estafette-ci-contracts\\t0.017s\"}],\"exitCode\":0,\"status\":\"SUCCEEDED\"}],\"insertedAt\":\"2018-04-17T08:03:00Z\"}", string(bytes))
	})
}
