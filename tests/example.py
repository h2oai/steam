import sys
import logging

# Turn on verbose logging so that we observe client-server communication.
# Can also redirect this to a file. Refer to 'logging' documentation.
logging.basicConfig(level=logging.DEBUG)

sys.path.insert(0, '../python')

import steam

# Connect as superuser
as_superuser = steam.RPCClient(steam.HTTPConnection('172.16.2.77', 9000, 'superuser', 'superuser'))

# Fetch a list of all available permissions
permission_list = as_superuser.get_supported_permissions()

permissions = {}
for permission in permission_list:
    permissions[permission.description] = permission.id

print permissions

# Next, we'll use superuser credentials to create a new user, a new role and assign the user to the role.
viewer_role_id = as_superuser.create_role("viewer", "Can only view clusters and models")

as_superuser.link_role_and_permissions(viewer_role_id, [
    permissions['View models'],
    permissions['View clusters']
])

roles = as_superuser.get_roles(0, 1000)

# Fetch roles
for role in roles:
    print role

zaphod_id = as_superuser.create_identity("zaphod", "beeblebrox")

as_superuser.link_identity_and_role(zaphod_id, viewer_role_id)

# Now connect as Zaphod
as_zaphod = steam.RPCClient(steam.HTTPConnection('172.16.2.76', 9000, 'zaphod', 'beeblebrox'))

# Fetch models
models = as_zaphod.get_models(0, 1000)

for model in models:
    print model

