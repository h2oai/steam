# User Management Overview

Before using Steam, it is important to understand User Management within your YARN environment. In Steam, User Management is supported in a PostgreSQL database. The User Management functions in Steam determine the level of access that users have for Steam features. The Steam database supports setup via CLI commands. Refer to the ***CLI Command Reference Appendix*** in the Steam User Guide for a list of all available CLI commands. 

For more information on Steam User Management, refer to the following sections. 

- [Terms](#terms)
- [Privileges/Access Control](#privileges)
- [Authorization](#authorization)
- [User Management Workflow](#usermgmtworkflow)
- [User Management Example](#usermgmtexample)
- [Next Steps](#nextsteps)

## <a name="terms"></a>Terms

The following lists common terms used when describing Steam User Management.  

- **Entities** represent *objects* in Steam. Examples of entities include Roles, Workgroups, Identities, Clusters, Projects, Models, and Services (engines). 

- **Identities** represent *users* in Steam. Users sign in using an Identity, and then perform operations in Steam.

- **Permissions** determine what operations you can perform. Examples of permissions include *Manage Clusters*, *View Clusters*, *Manage Models*, *View Models*, and so on.

- **Privileges** determine the entities that you can perform operations on (i.e., data / access control).



## <a name="privileges"></a>Privileges/Access Control

Privileges are uniquely identified by the entity in question and the kind of privilege you have on the entity.

The following privileges are available on an entity:

- **Own** privileges allow you to share, view, edit, and delete entities.

- **Edit** privileges allow you to view and edit entities, but not share or delete them.

- **View** privileges allow you to view entities, but not share, edit, or delete them.

When you create an entity, you immediately *Own* it. You can then share this entity with others and award them either *Edit* or *View* privileges. Entities are allowed to have more than one owner, so you can also add additional owners to entities. 

The following table lists the kind of privileges you need in order to perform specific operations on entities:


        Entity               Own  Edit View
        -----------------------------------
        Role
          Read               x    x    x
          Update             x    x
          Assign Permission  x    x
          Delete             x
          Share              x
          
        Workgroup
          Read               x    x    x
          Update             x    x
          Delete             x
          Share              x
        
        Identity
          Read               x    x    x
          Assign Role        x    x
          Assign Workgroup   x    x
          Update             x    x
          Delete             x
          Share              x
        
        Cluster
          Read               x    x    x
          Start/Stop         x
        
        Project
          Read               x    x    x
          Assign Model       x    x
          Update             x    x
          Delete             x
          Share              x
        
        Engine, Model
          Read               x    x    x
          Update             x    x
          Delete             x
          Share              x

## <a name="authorization"></a>Authorization

Permissions and privileges are set up using Roles and Workgroups, respectively.

- Identities cannot be linked directly to permissions. For that, you'll need Roles.

- Identities cannot be linked directly to privileges on entities. For that, you'll need Workgroups, i.e. when you share entities with others, you would be sharing those entities with workgroups, not individuals.

### Roles
A **Role** is a named set of permissions. Roles allow you define a cohesive set of permissions into operational roles and then have multiple identities *play* those roles, regardless of access control.
For example:

- a *Data Scientist* role can be composed of the permissions *View Clusters*, *Manage Models*, *View Models*.
- an *Operations* role can be composed of the permissions *View Models*, *View Services*, *Manage Services*,
- a *Manager* role can be composed of the permissions *Manage Roles*, *View Roles*, *Manage Workgroups*, *View Workgroups*

### Workgroups
A **Workgroup** is a named set of identities. Workgroups allow you to form collections of identities for access control purposes. For example, a *Demand Forecasting* workgroup can be composed of all the users working on demand forecasting, regardless of their role. This workgroup can be then used to control access to all the clusters, projects, models and services that are used for demand forecasting. 


## <a name="usermgmtworkflow"></a>User Management Workflow

The steps below provide a common workflow to follow when creating users. This workflow is followed in the example that follows.

1. Define roles based on operational needs.
2. Define workgroups based on data / access control needs.
3. Then add a new user:

 -	Create the user's identity.
 - Associate the user with one or more roles.
 - Optionally, associate the user with one or more workgroups. 

## <a name="usermgmtexample"></a>User Management Example

The following example creates sample roles, workgroups, and users using the CLI. Refer to the **CLI Command Reference Appendix** in the *Steam User Guide* for information about all of the commands available in the CLI. 

1. Log in as the Steam superuser on the machine that is running Steam.

 ```./steam login <yarn_edge_node>:<port> --username=superuser --password=superuser ```

1. Create an engineer role, and link that role to permissions. Note that you can run ``./steam get permissions`` to view a list of available permissions.

 ```./steam create role engineer --desc="a default engineer role"```
 
 ```./steam link role engineer ViewModel ViewProject ViewWorkgroup```
		
1. Create a data scientist role, and link that role to permissions.

 ```./steam create role datascience --desc="a default data scientist role"```
 
 ```./steam link role datascience ManageProject ManageModel ViewCluster```
		
1. Create preparation and production workgroups.

 ```./steam create workgroup preparation --desc="data prep group"```
 
 ```./steam create workgroup production --desc="production group"```
		
1. Create two users - Bob and Jim.

 ```./steam create identity bob bobSpassword```
 
 ```./steam create identity jim j1mSpassword```
		
1. Link Bob to engineer role; link Jim to datascience role.

 ```./steam link identity bob role engineer```
 
 ```./steam link identity jim role datascience```
		
1. Link Bob to preparation workgroup; link Jim to production workgroup.

 ```./steam link identity bob workgroup preparation```
 
 ```./steam link identity jim workgroup production```

## <a name="nextsteps"></a>Next Steps

Now that you understand User Management, you can create your own roles, workgroups, and users. Once created, be sure to provide your users with their Steam login credentials.  

