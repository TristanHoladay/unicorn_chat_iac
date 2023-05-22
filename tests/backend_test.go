package backend_test

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestS3BackendCreated() {
	// get output that backend mwas created
}

func TestS3BackendHoldsTFState() {
	// validate tf state saved to s3 backend .... is this test in the right spot / module?
}
