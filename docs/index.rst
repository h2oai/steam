.. sampledoc documentation master file, created by
   sphinx-quickstart on Mon Apr 25 15:24:53 2016.
   You can adapt this file completely to your liking, but it should at least
   contain the root `toctree` directive.

========
Overview
========

Steam is an "instant on" platform that streamlines the entire process of building and
deploying applications. It is the industry's first data science hub that lets data scientists
and developers collaboratively build, deploy, and refine predictive applications across
large scale datasets. Data scientists can publish Python and R code as REST APIs and
easily integrate with production applications.

This document describes how to start and use Steam and the Steam Scoring Service.
Note that this document assumes that an admin has successfully installed and started
Steam on a YARN edge node using the instructions provided in the Steam Installation
and Setup steps.

**Note**: Before you begin using Steam, be sure that your minimum version of H2O
is 3.10.0.3. Earlier versions are not supported. If necessary, follow the instructions
on the `H2O Download page <http://h2o.ai/download>`__ for your platform to upgrade H2O.

.. toctree::
   :maxdepth: 2

   Installation

.. toctree::
   :maxdepth: 2

   UserManagement

.. toctree::
   :maxdepth: 2

   Login
   Projects
   Clusters
   Users
   UseSteamWithFlow
   StopSteam

.. toctree::
   :maxdepth: 2
   
   UsingCLI
   CLIAppendix