package bosh

import (
	"fmt"
	"strings"
	
	"golang.org/x/net/context"
	"github.com/cloudfoundry-community/gogobosh"
	"github.com/cloudfoundry-community/gogobosh/api"
	"github.com/cloudfoundry-community/gogobosh/models"
	"github.com/cloudfoundry-community/gogobosh/net"
)

type BoshClient struct {
	UUID string
	Name string
	Version string
	CPI string
	
	director models.Director
}

func NewBoshClient(ctx context.Context, target string, user string, password string) (*BoshClient, error) {
	
	c := &BoshClient {
		director: gogobosh.NewDirector(target, user, password),
	}
	
	err := c.Connect()
	if err != nil && !strings.Contains(err.Error(), "connection refused") {
		return nil, err
	}
	
	return c, nil 
}

func (b *BoshClient) Connect() error {
	
	repo := api.NewBoshDirectorRepository(&b.director, net.NewDirectorGateway())

	info, resp := repo.GetInfo()
	if resp.IsNotSuccessful() {
		return fmt.Errorf("directory information query was unsuccessful: %s - %s", resp.ErrorCode, resp.Message)
	}
	
	b.UUID = info.UUID
	b.Name = info.Name
	b.Version = info.Version
	b.CPI = info.CPI
	
	return nil
}

func (b *BoshClient) IsConnected() bool {
	return b.UUID != ""
}
