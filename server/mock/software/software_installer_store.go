// Automatically generated by mockimpl. DO NOT EDIT!

package mock

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/fleetdm/fleet/v4/server/fleet"
)

var _ fleet.SoftwareInstallerStore = (*SoftwareInstallerStore)(nil)

type GetFunc func(ctx context.Context, installerID string) (io.ReadCloser, int64, error)

type PutFunc func(ctx context.Context, installerID string, content io.ReadSeeker) error

type ExistsFunc func(ctx context.Context, installerID string) (bool, error)

type CleanupFunc func(ctx context.Context, usedInstallerIDs []string, removeCreatedBefore time.Time) (int, error)

type SignFunc func(ctx context.Context, fileID string) (string, error)

type SoftwareInstallerStore struct {
	GetFunc        GetFunc
	GetFuncInvoked bool

	PutFunc        PutFunc
	PutFuncInvoked bool

	ExistsFunc        ExistsFunc
	ExistsFuncInvoked bool

	CleanupFunc        CleanupFunc
	CleanupFuncInvoked bool

	SignFunc        SignFunc
	SignFuncInvoked bool

	mu sync.Mutex
}

func (s *SoftwareInstallerStore) Get(ctx context.Context, installerID string) (io.ReadCloser, int64, error) {
	s.mu.Lock()
	s.GetFuncInvoked = true
	s.mu.Unlock()
	return s.GetFunc(ctx, installerID)
}

func (s *SoftwareInstallerStore) Put(ctx context.Context, installerID string, content io.ReadSeeker) error {
	s.mu.Lock()
	s.PutFuncInvoked = true
	s.mu.Unlock()
	return s.PutFunc(ctx, installerID, content)
}

func (s *SoftwareInstallerStore) Exists(ctx context.Context, installerID string) (bool, error) {
	s.mu.Lock()
	s.ExistsFuncInvoked = true
	s.mu.Unlock()
	return s.ExistsFunc(ctx, installerID)
}

func (s *SoftwareInstallerStore) Cleanup(ctx context.Context, usedInstallerIDs []string, removeCreatedBefore time.Time) (int, error) {
	s.mu.Lock()
	s.CleanupFuncInvoked = true
	s.mu.Unlock()
	return s.CleanupFunc(ctx, usedInstallerIDs, removeCreatedBefore)
}

func (s *SoftwareInstallerStore) Sign(ctx context.Context, fileID string) (string, error) {
	s.mu.Lock()
	s.SignFuncInvoked = true
	s.mu.Unlock()
	return s.SignFunc(ctx, fileID)
}
