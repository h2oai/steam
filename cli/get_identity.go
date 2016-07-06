package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var getIdentityHelp = `
identity [username]
View detailed user information.
Examples:

	$ steam get identity minsky
`

func getIdentity(c *context) *cobra.Command {
	cmd := newCmd(c, getIdentityHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Invalid usage. See 'steam help get identity'.")
		}

		// -- Args --

		username := args[0]

		// -- Execution --

		identity, err := c.remote.GetIdentityByName(username)
		if err != nil {
			log.Fatalln(err) //FIXME format error
		}

		workgroups, err := c.remote.GetWorkgroupsForIdentity(identity.Id)
		if err != nil {
			log.Fatalln(err) //FIXME format error
		}

		roles, err := c.remote.GetRolesForIdentity(identity.Id)
		if err != nil {
			log.Fatalln(err) //FIXME format error
		}

		permissions, err := c.remote.GetPermissionsForIdentity(identity.Id)
		if err != nil {
			log.Fatalln(err) //FIXME format error
		}

		// -- Formatting --

		base := []string{
			fmt.Sprintf("STATUS:\t%s", identityStatus(identity.IsActive)),
			fmt.Sprintf("LAST LOGIN:\t%s", fmtAgo(identity.LastLogin)),
			fmt.Sprintf("ID:\t%d", identity.Id),
			fmt.Sprintf("AGE:\t%s", fmtAgo(identity.Created)),
		}

		c.printt("\t"+identity.Name, base)

		ws := make([]string, len(workgroups))
		for i, w := range workgroups {
			ws[i] = fmt.Sprintf("%s\t%s", w.Name, w.Description)
		}
		c.printt("WORKGROUP\tDESCRIPTION", ws)

		rs := make([]string, len(roles))
		for i, r := range roles {
			rs[i] = fmt.Sprintf("%s\t%s", r.Name, r.Description)
		}
		c.printt("ROLE\tDESCRIPTION", rs)

		fmt.Println("PERMISSIONS")
		for _, p := range permissions {
			fmt.Println(p.Description)
		}

	})

	return cmd
}
