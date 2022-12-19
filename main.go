package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"regexp"
	"syscall"
	"time"
	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

func main() {
	var core zapcore.Core
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core = ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	var goModFile string
	flag.StringVar(&goModFile, "f", "go.mod", "Specify how the file with the go.mod content is named")

	flag.Usage = func() {
		fmt.Printf("Usage of our Program: \n")
		fmt.Printf("./main -f Export-81311be8-cfa2-4981-ad0f-b169eaf9933f.zip\n")
		flag.PrintDefaults() // prints default usage
	}
	flag.Parse()

	byteToken, _ := terminal.ReadPassword(int(syscall.Stdin))
	println()
	token := string(byteToken)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	repositories, resp, err := client.Repositories.Get(ctx, "united-manufacturing-hub", "united-manufacturing-hub")
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	fmt.Printf("%+v", repositories)

	goModContent, err := os.ReadFile(goModFile)
	if err != nil {
		zap.S().Errorf("cannot read go.mod file, check -f flag")
	}
	regexmatch1 := regexp.MustCompile(`([a-zA-Z0-9\/\.\-\_]*)\s(v[0-9\.\-a-zA-Z]*)`) // this regex matches the dependencies, the first capturing group is the name of the dependency, the second one is the version
	regexmatch2 := regexp.MustCompile(`(go)\s([0-9\.]*)`)                            // this regex matches the go version with the first group just being "go" and the second group being the version
	dependencyList := regexmatch1.FindAllStringSubmatch(string(goModContent), -1)

	Jsonoutputlist := os.WriteFile("output.json", []byte(""), 0644)

	type Entry struct {
		ID              uint64    `json:"id"`
		Name            string    `json:"name"`
		Version         string    `json:"version"`
		GoVersion       string    `json:"goversion"`
		Vulnerabilities []string  `json:"vulnerabilities"`
		Licence         string    `json:"licence"`
		LastUpdate      time.Time `json:"lastupdate"`
	}
	var testData Entry
	testData.ID = 1
	testData.Name = "united-manufacturing-hub"
	testData.Version = "0.9.9"
	testData.GoVersion = regexmatch2.FindString("$2")
	testData.Vulnerabilities = nil
	testData.Licence =



		json :=	json.Unmarshal()
	os.WriteFile("output.json", ,0644)
}
