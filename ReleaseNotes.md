# Steam Release Notes

Thank you for you interest in Steam, the industry’s first data science hub that lets data scientists and developers collaboratively build, deploy, and refine predictive applications across large scale datasets. 

## Contacting Support

If you're an Open Source community member, you can contact H2O using one of the following methods:

- Click the Support link in the Steam UI to send an e-mail message
- Send an e-mail message directly to <a href="mailto:support@h2o.ai">support@h2o.ai</a>
- Ask your question on the [H2O Community](https://community.h2o.ai/spaces/540/index.html) site (create an account if necessary)

## v1.1.5 Changes

- [STEAM-613] The ``./steam add engine`` command has been replaced with ``./steam upload engine``. 
- [STEAM-604] When adding clusters to Steam, node sizes are now only specified in GB.
- [STEAM-619] Adds a confirmation dialog on deletion of a cluster
- [STEAM-601] Adds functionality to import a yaml imports file for Anaconda

## v1.1.4 Changes

- [STEAM-496] Early release support for LDAP basic authentication using the ``--authentication-provider`` and ``--authentication-config`` flags when starting Steam.
- [STEAM-562] Steam superusers can add and maintain users and roles in the UI.
- [PUBSTEAM-4] When a dataframe that is used to train a model is deleted from the cluster, the model metadata will remain in the Steam database.
- [PUBSTEAM-3] Servers that are down are no longer shown as green in the main view.

## v1.1.3 Changes

- [STEAM-567] Updated Prediction Service UI to reflect DeepWater support
- [STEAM-578] Added support for importing deploying DeepWater models
- [STEAM-579] Added a default limit for CLI items to 10000
- [STEAM-504] Modify existing permissions from UMS screen 
- [STEAM-559] Prediction Service now allows binary input (images, audio, etc...)

## v1.1.2 Changes

- [STEAM-574] Fixes missing protocol in Cluster listing addresses

## v1.1.1 Changes

- [STEAM-517] The cluster URL to H2O Flow now works correctly.
- [N/A] - The **Clusters** page now shows the number of nodes on the cluster, the version of H2O that is running on the cluster, and the cluster health state.  

## v1.1.0 Changes

- [STEAM-450] Projects can be deleted through the UI. Note that only projects without models can be deleted.
- [STEAM-507] Models can now be deleted through the UI.
- [STEAM-518] Models can be downloaded as MOJOs.


