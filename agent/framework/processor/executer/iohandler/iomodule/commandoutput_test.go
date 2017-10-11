package iomodule

import (
	"testing"

	"io"

	"github.com/aws/amazon-ssm-agent/agent/log"
	"github.com/stretchr/testify/assert"
)

// TestInputCases is a list of strings which we test multi-writer on.
var TestInputCases = [...]string{
	"Test input text.",
	"A sample \ninput text.",
	"\b5Ὂg̀9! ℃ᾭG",
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. In fermentum cursus mi, sed placerat tellus condimentum non. " +
		"Pellentesque vel volutpat velit. Sed eget varius nibh. Sed quis nisl enim. Nulla faucibus nisl a massa fermentum porttitor. " +
		"Integer at massa blandit, congue ligula ut, vulputate lacus. Morbi tempor tellus a tempus sodales. Nam at placerat odio, " +
		"ut placerat purus. Donec imperdiet venenatis orci eu mollis. Phasellus rhoncus bibendum lacus sit amet cursus. Aliquam erat" +
		" volutpat. Phasellus auctor ipsum vel efficitur interdum. Duis sed elit tempor, convallis lacus sed, accumsan mi. Integer" +
		" porttitor a nunc in porttitor. Vestibulum felis enim, pretium vel nulla vel, commodo mollis ex. Sed placerat mollis leo, " +
		"at varius eros elementum vitae. Nunc aliquet velit quis dui facilisis elementum. Etiam interdum lobortis nisi, vitae " +
		"convallis libero tincidunt at. Nam eu velit et velit dignissim aliquet facilisis id ipsum. Vestibulum hendrerit, arcu " +
		"id gravida facilisis, felis leo malesuada eros, non dignissim quam turpis a massa. ",
}

var logger = log.NewMockLog()

func testCommandOuput(testCase string, limit int) string {
	r, w := io.Pipe()
	readerClosed := make(chan bool)
	var stdout string

	// Initialize console output module
	stdoutConsole := CommandOutput{
		OutputLimit:  limit,
		OutputString: &stdout,
	}

	go stdoutConsole.Read(logger, r, readerClosed)
	w.Write([]byte(testCase))
	w.Close()
	<-readerClosed
	return stdout
}

// TestCommandOuput tests the CommandOutput module
func TestCommandOuput(t *testing.T) {
	for _, testInput := range TestInputCases {
		stdout := testCommandOuput(testInput, 20)
		if len(testInput) > 20 {
			assert.Equal(t, stdout, testInput[:20])
		} else {
			assert.Equal(t, stdout, testInput)
		}
	}
}
