//+build illumos

package freeport

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

/*
Sun Solaris Operating Environment
Solaris uses the ndd utility program to configure the IP stack. Solaris uses two separate ephemeral port ranges, one for TCP and UDP. Both ports default to values from 32768 through 65535.
The following examples illustrate how you can query and change the settings:

Example 1: Query the Current Settings

# ipadm show-prop  -co CURRENT  -p smallest_anon_port,largest_anon_port tcp
32768
65535
*/

const ephemeralPortRangenddKey = "smallest_anon_port,largest_anon_port"

var ephemeralPortRangePatt = regexp.MustCompile(`^\s*(\d+)\n(\d+)`)

func getEphemeralPortRange() (int, int, error) {
	cmd := exec.Command("/usr/sbin/ipadm", "show-prop", "-co", "CURRENT", "-p",  ephemeralPortRangenddKey, "tcp")
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}

	val := string(out)

	m := ephemeralPortRangePatt.FindStringSubmatch(val)
	if m != nil {
		min, err1 := strconv.Atoi(m[1])
		max, err2 := strconv.Atoi(m[2])

		if err1 == nil && err2 == nil {
			return min, max, nil
		}
	}

	return 0, 0, fmt.Errorf("unexpected  value %q for key %q", val, ephemeralPortRangenddKey)
}
