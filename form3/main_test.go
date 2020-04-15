package form3

import (
	"flag"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/logging"
)

func TestMain(m *testing.M) {
	flag.Parse()

	if testing.Verbose() {
		logging.SetOutput()
	}

	os.Exit(m.Run())
}
