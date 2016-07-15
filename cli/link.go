package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var linkHelp = `
link [permission-type]
Add authentication permissions.
Examples:

	$ steam link permission
`

func link(c *context) *cobra.Command {
	cmd := newCmd(c, linkHelp, nil)
	cmd.AddCommand(linkRole(c))
	cmd.AddCommand(linkIdentity(c))
	return cmd
}

/* getPermissionIds
 *
 * Support function to return ids from codes
 */
func getPermissionIds(c *context, codes ...string) ([]int64, error) {
	permissions, err := c.remote.GetSupportedPermissions()
	if err != nil {
		return nil, err //FIXME format error
	}

	permissionsMap := make(map[string]int64)
	for _, p := range permissions {
		permissionsMap[p.Code] = p.Id
	}

	ids := make([]int64, len(codes))
	for i, c := range codes {
		var ok bool
		ids[i], ok = permissionsMap[c]
		if !ok {
			return nil, fmt.Errorf("Failed getting permissions: cannot locate permission with corresponding code: %s", c)
		}
	}

	return ids, nil
}
