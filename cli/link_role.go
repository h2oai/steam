package cli

// FIXME use ByName remote calls

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var linkRoleHelp = `
role [roleId] [permissionIds]
Add permissions to a role. 
Examples:

	$ steam link role 2 2,4,6
`

func linkRole(c *context) *cobra.Command {
	cmd := newCmd(c, linkRoleHelp, func(c *context, args []string) {
		if len(args) != 2 {
			log.Fatalln("Invalid usage. See 'steam help link role'.")
		}

		roleId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatalln("Invalid usage for roleId %s: expecting int: %v", args[0], err)
		}

		rawIds := strings.Split(args[1], ",")
		permissionIds := make([]int64, len(rawIds))
		for i, id := range rawIds {
			permissionIds[i], err = strconv.ParseInt(id, 10, 64)
			if err != nil {
				log.Fatalln("Invalid usage for permissionIds %s: expecting int: %v", id, err)
			}
		}

		c.remote.LinkRoleAndPermissions(roleId, permissionIds)
		fmt.Println("Role", roleId, "linked to permissions:", permissionIds)
	})

	return cmd
}
