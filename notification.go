package contracts

import (
	"encoding/json"
	"fmt"
	"time"
)

type NotificationType string

const (
	NotificationTypeUnknown       NotificationType = ""
	NotificationTypeVulnerability NotificationType = "vulnerability"
	NotificationTypeWarning       NotificationType = "warning"
)

type NotificationLevel string

const (
	NotificationLevelUnknown  NotificationLevel = ""
	NotificationLevelCritical NotificationLevel = "critical"
	NotificationLevelHigh     NotificationLevel = "high"
	NotificationLevelMedium   NotificationLevel = "medium"
	NotificationLevelLow      NotificationLevel = "low"
)

type NotificationLinkType string

const (
	NotificationLinkTypeUnknown   NotificationLinkType = ""
	NotificationLinkTypePipeline  NotificationLinkType = "pipeline"
	NotificationLinkTypeContainer NotificationLinkType = "container"
)

type Notification struct {
	Type    NotificationType  `json:"type,omitempty"`
	Level   NotificationLevel `json:"level,omitempty"`
	Message string            `json:"message,omitempty"`
}

type NotificationRecord struct {
	ID       string               `json:"id,omitempty"`
	LinkType NotificationLinkType `json:"linkType,omitempty"`
	LinkID   string               `json:"linkID,omitempty"`

	// fields mapped to link_detail column
	PipelineDetail  *PipelineLinkDetail  `json:"pipelineDetail,omitempty"`
	ContainerDetail *ContainerLinkDetail `json:"containerDetail,omitempty"`

	Source        string          `json:"source,omitempty"`
	Notifications []Notification  `json:"notifications,omitempty"`
	InsertedAt    *time.Time      `json:"insertedAt,omitempty"`
	Groups        []*Group        `json:"groups,omitempty"`
	Organizations []*Organization `json:"organizations,omitempty"`
}

type PipelineLinkDetail struct {
	Branch   string `json:"branch,omitempty"`
	Revision string `json:"revision"`
	Version  string `json:"version,omitempty"`
	Status   Status `json:"status,omitempty"`
}

type ContainerLinkDetail struct {
	Tag         string `json:"tag,omitempty"`
	PublicImage bool   `json:"publicImage,omitempty"`
}

func (nr *NotificationRecord) GetLinkDetail() ([]byte, error) {
	switch nr.LinkType {
	case NotificationLinkTypePipeline:
		return json.Marshal(nr.PipelineDetail)
	case NotificationLinkTypeContainer:
		return json.Marshal(nr.ContainerDetail)
	}

	return []byte{}, nil
}

func (nr *NotificationRecord) SetLinkDetail(linkDetail []byte) error {
	if len(linkDetail) == 0 {
		return nil
	}

	switch nr.LinkType {
	case NotificationLinkTypePipeline:
		if err := json.Unmarshal(linkDetail, &nr.PipelineDetail); err != nil {
			return fmt.Errorf("LinkDetail for NotificationRecord %v of type %v is not of type PipelineLinkDetail: %w", nr.LinkID, nr.LinkType, err)
		}
	case NotificationLinkTypeContainer:
		if err := json.Unmarshal(linkDetail, &nr.ContainerDetail); err != nil {
			return fmt.Errorf("LinkDetail for NotificationRecord %v of type %v is not of type ContainerLinkDetail: %w", nr.LinkID, nr.LinkType, err)
		}
	}

	return nil
}

func (nr *NotificationRecord) GetNotifications() ([]byte, error) {
	return json.Marshal(nr.Notifications)
}

func (nr *NotificationRecord) SetNotifications(notifications []byte) error {
	if len(notifications) == 0 {
		return nil
	}

	if err := json.Unmarshal(notifications, &nr.Notifications); err != nil {
		return fmt.Errorf("Notifications for NotificationRecord %v of type %v is not of type []Notification: %w", nr.LinkID, nr.LinkType, err)
	}

	return nil
}

func (nr *NotificationRecord) GetGroups() ([]byte, error) {
	return json.Marshal(nr.Groups)
}

func (nr *NotificationRecord) SetGroups(groups []byte) error {
	if len(groups) == 0 {
		return nil
	}

	if err := json.Unmarshal(groups, &nr.Groups); err != nil {
		return fmt.Errorf("Groups for NotificationRecord %v of type %v is not of type []*Group: %w", nr.LinkID, nr.LinkType, err)
	}

	return nil
}

func (nr *NotificationRecord) GetOrganizations() ([]byte, error) {
	return json.Marshal(nr.Organizations)
}

func (nr *NotificationRecord) SetOrganizations(organizations []byte) error {
	if len(organizations) == 0 {
		return nil
	}

	if err := json.Unmarshal(organizations, &nr.Organizations); err != nil {
		return fmt.Errorf("Organizations for NotificationRecord %v of type %v is not of type []*Group: %w", nr.LinkID, nr.LinkType, err)
	}

	return nil
}
