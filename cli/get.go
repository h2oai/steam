package cli

import "github.com/spf13/cobra"

var getHelp = `
get [resource-type]
List or view resources of the specified type.
Examples:

    $ steam get clusters
`

func get(c *context) *cobra.Command {
	cmd := newCmd(c, getHelp, nil)
	cmd.AddCommand(getClusters(c))
	cmd.AddCommand(getCluster(c))
	cmd.AddCommand(getEngines(c))
	cmd.AddCommand(getEngine(c))
	cmd.AddCommand(getModels(c))
	cmd.AddCommand(getModel(c))
	cmd.AddCommand(getServices(c))
	cmd.AddCommand(getService(c))
	cmd.AddCommand(getIdentities(c))
	cmd.AddCommand(getIdentity(c))
	cmd.AddCommand(getRoles(c))
	cmd.AddCommand(getRole(c))
	cmd.AddCommand(getWorkgroups(c))
	cmd.AddCommand(getWorkgroup(c))
	cmd.AddCommand(getPermissions(c))
	cmd.AddCommand(getEntities(c))
	cmd.AddCommand(getHistory(c))
	return cmd
}

/* getClusterTypes
 *
 * Support function to map TypeIds to cluster types (i.e. yarn, external, etc.)
 */
func getClusterTypes(c *context) (map[int64]string, error) {
	raw, err := c.remote.GetSupportedClusterTypes()
	if err != nil {
		return nil, err //FIXME format error
	}

	ret := make(map[int64]string)
	for _, ct := range raw {
		ret[ct.Id] = ct.Name
	}

	return ret, nil
}

/* identityStatus
 *
 * Support function to return active or inactive from bool
 */
func identityStatus(isActive bool) string {
	if isActive {
		return "Active"
	}

	return "Inactive"
}
