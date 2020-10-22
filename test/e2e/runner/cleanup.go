package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	e2e "github.com/tendermint/tendermint/test/e2e/pkg"
)

// Cleanup removes the Docker Compose containers and testnet directory.
func Cleanup(testnet *e2e.Testnet) error {
	err := cleanupDocker()
	if err != nil {
		return err
	}
	err = cleanupDir(testnet.Dir)
	if err != nil {
		return err
	}
	return nil
}

// cleanupDocker removes all E2E resources (with label e2e=True), regardless
// of testnet.
func cleanupDocker() error {
	logger.Info("Removing Docker containers and networks")

	// '' is required since xargs runs even when there's no input, and
	// macOS' xargs does not support the -r switch.
	err := exec("sh", "-c", "docker container ls -q --filter label=e2e | xargs docker container rm -f ''")
	if err != nil {
		return err
	}

	err = exec("sh", "-c", "docker network ls -q --filter label=e2e | xargs docker network rm ''")
	if err != nil {
		return err
	}

	return nil
}

// cleanupDir cleans up a testnet directory
func cleanupDir(dir string) error {
	if dir == "" {
		return errors.New("no directory set")
	}

	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("Removing testnet directory %q", dir))

	// On Linux, some local files in the volume will be owned by root since Tendermint
	// runs as root inside the container, so we need to clean them up from within a
	// container running as root too.
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return err
	}
	err = execDocker("run", "--rm", "--entrypoint", "", "-v", fmt.Sprintf("%v:/network", absDir),
		"tendermint/e2e-node", "sh", "-c", "rm -rf /network/*/")
	if err != nil {
		return err
	}

	err = os.RemoveAll(dir)
	if err != nil {
		return err
	}

	return nil
}
