# Introduction

This document describes how to get up and running with Steam without the need for a local running instance of YARN. These instructions will walk through the following procedures:

- Installing and starting Steam, the Compilation Service, and H2O
- Adding Roles, Workgroups, and Users to the database
- Building a simple model in Python (Optional for users who don't have an existing demo.)
- Deploying the model using Steam
- Making predictions using the Steam Prediction Service

During this demo, four terminal windows will remain open for the Steam, Scoring, H2O, and postgres services. A fifth terminal window will be used to run H2O commands in the Python or R example. 

Finally, these steps were created using H2O version 3.10.0.3, and that version resides in a Downloads folder. Wherever used, this version number and path should be adjusted to match your version and path.

