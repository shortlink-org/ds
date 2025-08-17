package anthropic

import (
	"bufio"
	"context"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/spf13/viper"
)

// New creates a new anthropic client with default options.
func New(ctx context.Context) (*Agent, error) {
	API_KEY := viper.GetString("ANTHROPIC_API_KEY") // Read the API key from the configuration

	client := anthropic.NewClient(
		option.WithAPIKey(API_KEY),
	)

	scanner := bufio.NewScanner(os.Stdin)
	getUserMessage := func() (string, bool) {
		if !scanner.Scan() {
			return "", false
		}
		return scanner.Text(), true
	}

	agent := NewAgent(&client, getUserMessage)
	err := agent.Run(ctx)
	if err != nil {
		return agent, err
	}

	return agent, nil
}
