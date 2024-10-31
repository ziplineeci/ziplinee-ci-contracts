package contracts

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {

	t.Run("JSONMarshalPayloadSingleBuild", func(t *testing.T) {

		build := Build{
			ID:           "3",
			RepoSource:   "github.com",
			RepoOwner:    "estafette",
			RepoName:     "estafette-ci-api",
			RepoBranch:   "master",
			RepoRevision: "as23456",
			BuildVersion: "1.0.0",
			BuildStatus:  StatusSucceeded,
			Labels: []Label{
				Label{
					Key:   "app",
					Value: "estafette-ci-api",
				},
				Label{
					Key:   "team",
					Value: "estafette-team",
				},
				Label{
					Key:   "language",
					Value: "golang",
				},
			},
			Manifest: "",
			Commits: []GitCommit{
				GitCommit{
					Message: "First commit",
					Author: GitAuthor{
						Email:    "name@server.com",
						Name:     "Name",
						Username: "MyName",
					},
				},
			},
			InsertedAt: time.Date(2018, 4, 17, 8, 3, 0, 0, time.UTC),
			UpdatedAt:  time.Date(2018, 4, 17, 8, 15, 0, 0, time.UTC),
		}

		// act
		bytes, err := json.Marshal(&build)

		assert.Nil(t, err)
		assert.Equal(t, "{\"id\":\"3\",\"repoSource\":\"github.com\",\"repoOwner\":\"estafette\",\"repoName\":\"estafette-ci-api\",\"repoBranch\":\"master\",\"repoRevision\":\"as23456\",\"buildVersion\":\"1.0.0\",\"buildStatus\":\"succeeded\",\"labels\":[{\"key\":\"app\",\"value\":\"estafette-ci-api\"},{\"key\":\"team\",\"value\":\"estafette-team\"},{\"key\":\"language\",\"value\":\"golang\"}],\"commits\":[{\"message\":\"First commit\",\"author\":{\"email\":\"name@server.com\",\"name\":\"Name\",\"username\":\"MyName\"}}],\"insertedAt\":\"2018-04-17T08:03:00Z\",\"updatedAt\":\"2018-04-17T08:15:00Z\",\"duration\":0}", string(bytes))
	})

	t.Run("JSONMarshalPayloadArrayOfBuilds", func(t *testing.T) {

		builds := make([]*Build, 0)

		builds = append(builds, &Build{
			ID:           "3",
			RepoSource:   "github.com",
			RepoOwner:    "estafette",
			RepoName:     "estafette-ci-api",
			RepoBranch:   "master",
			RepoRevision: "as23456",
			BuildVersion: "1.0.0",
			BuildStatus:  StatusSucceeded,
			Labels: []Label{
				Label{
					Key:   "app",
					Value: "estafette-ci-api",
				},
				Label{
					Key:   "team",
					Value: "estafette-team",
				},
				Label{
					Key:   "language",
					Value: "golang",
				},
			},
			Manifest: "",
			Commits: []GitCommit{
				GitCommit{
					Message: "First commit",
					Author: GitAuthor{
						Email:    "name@server.com",
						Name:     "Name",
						Username: "MyName",
					},
				},
			},
			InsertedAt: time.Date(2018, 4, 17, 8, 3, 0, 0, time.UTC),
			UpdatedAt:  time.Date(2018, 4, 17, 8, 15, 0, 0, time.UTC),
		})
		builds = append(builds, &Build{
			ID:           "8",
			RepoSource:   "github.com",
			RepoOwner:    "estafette",
			RepoName:     "estafette-ci-api",
			RepoBranch:   "master",
			RepoRevision: "as23456",
			BuildVersion: "1.0.0",
			BuildStatus:  StatusSucceeded,
			Labels: []Label{
				Label{
					Key:   "app",
					Value: "estafette-ci-api",
				},
				Label{
					Key:   "team",
					Value: "estafette-team",
				},
				Label{
					Key:   "language",
					Value: "golang",
				},
			},
			Manifest: "",
			Commits: []GitCommit{
				GitCommit{
					Message: "Second commit",
					Author: GitAuthor{
						Email:    "othername@server.com",
						Name:     "Other Name",
						Username: "OtherName",
					},
				},
			},
			InsertedAt: time.Date(2018, 4, 17, 8, 3, 0, 0, time.UTC),
			UpdatedAt:  time.Date(2018, 4, 17, 8, 15, 0, 0, time.UTC),
		})

		// act
		bytes, err := json.Marshal(&builds)

		assert.Nil(t, err)
		assert.Equal(t, "[{\"id\":\"3\",\"repoSource\":\"github.com\",\"repoOwner\":\"estafette\",\"repoName\":\"estafette-ci-api\",\"repoBranch\":\"master\",\"repoRevision\":\"as23456\",\"buildVersion\":\"1.0.0\",\"buildStatus\":\"succeeded\",\"labels\":[{\"key\":\"app\",\"value\":\"estafette-ci-api\"},{\"key\":\"team\",\"value\":\"estafette-team\"},{\"key\":\"language\",\"value\":\"golang\"}],\"commits\":[{\"message\":\"First commit\",\"author\":{\"email\":\"name@server.com\",\"name\":\"Name\",\"username\":\"MyName\"}}],\"insertedAt\":\"2018-04-17T08:03:00Z\",\"updatedAt\":\"2018-04-17T08:15:00Z\",\"duration\":0},{\"id\":\"8\",\"repoSource\":\"github.com\",\"repoOwner\":\"estafette\",\"repoName\":\"estafette-ci-api\",\"repoBranch\":\"master\",\"repoRevision\":\"as23456\",\"buildVersion\":\"1.0.0\",\"buildStatus\":\"succeeded\",\"labels\":[{\"key\":\"app\",\"value\":\"estafette-ci-api\"},{\"key\":\"team\",\"value\":\"estafette-team\"},{\"key\":\"language\",\"value\":\"golang\"}],\"commits\":[{\"message\":\"Second commit\",\"author\":{\"email\":\"othername@server.com\",\"name\":\"Other Name\",\"username\":\"OtherName\"}}],\"insertedAt\":\"2018-04-17T08:03:00Z\",\"updatedAt\":\"2018-04-17T08:15:00Z\",\"duration\":0}]", string(bytes))
	})
}
