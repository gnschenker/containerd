// +build !linux

package server

import "context"

func apply(_ context.Context, _ *Config) error {
	return nil
}
