Using the Steam CLI
-------------------

The CLI is an optional utility that can be used to maintain a Steam
environment and to create new roles, workgroups, and users. The CLI will
primarily be used by admins and/or Steam superusers. The steps below
describe how to start the Steam CLI.

Perform the following steps to start the Steam CLI.

1. Open a terminal window and ssh to the machine running Steam. Be sure
   to provide the correct password for the node when prompted.

   ::

		ssh <user>@<yarn_edge_node>

2. Change directories to the Steam folder. From within this folder, log
   in to the machine running Steam. Use the password that you provided
   when you created superuser. The exmaple below logs in a user named
   **Bob**.

   ::

		cd steam--darwin-amd64/
		./steam login 192.168.2.182:8080 --username=bob --password=bobSpassword

3. Run the following to verify that the CLI is working correctly.

   ::

       ./steam --help

Refer to the `CLI Command Reference Appendix <CLIAppendix.md>`__ for
information on the commands available in the CLI.
