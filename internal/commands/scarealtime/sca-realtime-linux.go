//go:build linux

package scarealtime

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/checkmarx/ast-cli/internal/logger"
)

const temporaryProjectPathToScan = "../../small-project/"

var linuxSCARealTime = ScaRealTime{
	ExecutableFilePath:         filepath.Join(scaResolverWorkingDir, "ScaResolver"),
	HashFilePath:               filepath.Join(scaResolverWorkingDir, "ScaResolver.tar.gz.sha256sum"),
	SCAResolverDownloadURL:     "https://sca-downloads.s3.amazonaws.com/cli/latest/ScaResolver-linux64.tar.gz",
	SCAResolverHashDownloadURL: "https://sca-downloads.s3.amazonaws.com/cli/latest/ScaResolver-linux64.tar.gz.sha256sum",
	SCAResolverFileName:        "ScaResolver.tar.gz",
	SCAResolverHashFileName:    "ScaResolver.tar.gz.sha256sum",
}

// getScaResolver Gets SCA Resolver file path to run SCA Realtime
func getScaResolver() (string, error) {
	err := downloadSCAResolverAndHashFileIfNeeded(&linuxSCARealTime)
	if err != nil {
		return "", err
	}

	return linuxSCARealTime.ExecutableFilePath, nil
}

// unzipOrExtractFiles Extracts SCA Resolver files
func unzipOrExtractFiles() error {
	logger.PrintIfVerbose("Extracting files in: " + scaResolverWorkingDir)
	gzipStream, err := os.Open(filepath.Join(scaResolverWorkingDir, linuxSCARealTime.SCAResolverFileName))
	if err != nil {
		fmt.Println("error")
	}
	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		log.Fatal("ExtractTarGz: NewReader failed")
	}

	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(header.Name, 0755); err != nil { //nolint:gomnd
				log.Fatalf("ExtractTarGz: Mkdir() failed: %s", err.Error())
			}
		case tar.TypeReg:
			extractedFilePath := filepath.Join(scaResolverWorkingDir, header.Name)
			outFile, err := os.Create(extractedFilePath)
			if err != nil {
				log.Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
			}
			if _, err = io.Copy(outFile, tarReader); err != nil {
				log.Fatalf("ExtractTarGz: Copy() failed: %s", err.Error())
			}
			err = outFile.Close()
			if err != nil {
				return err
			}
			err = os.Chmod(extractedFilePath, fs.ModePerm)
			if err != nil {
				return err
			}
		default:
			log.Fatalf(
				"ExtractTarGz: uknown type: %v in %s",
				header.Typeflag,
				header.Name)
		}
	}

	return nil
}
