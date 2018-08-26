package DeepSort

import "errors"

// ErrAlreadyRunning is thrown by ClassificationService.Load
// if another service with the same ID is already loaded.
var ErrAlreadyRunning = errors.New("already running")

// ErrStartFailed is thrown if the ClassificationService
// could not start and had to abort.
var ErrStartFailed = errors.New("failed to start")
