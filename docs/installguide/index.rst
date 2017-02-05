.. sampledoc documentation master file, created by
   sphinx-quickstart on Mon Apr 25 15:24:53 2016.
   You can adapt this file completely to your liking, but it should at least
   contain the root `toctree` directive.

============================
Steam Installation and Setup
============================

Steam is an "instant on" platform that streamlines the entire process of building and
deploying applications. It is the industry's first data science hub that lets data scientists
and developers collaboratively build, deploy, and refine predictive applications across
large scale datasets. Data scientists can publish Python and R code as REST APIs and
easily integrate with production applications.

This document applies to Administrators and describes how to install and start Steam in a Hadoop environment and make it accessible to a set of users. The process includes uploading an H2O driver and adding users via an LDAP config file. 

**Note**: Before you begin using Steam, be sure that your minimum version of H2O
is 3.10.3.2. Earlier versions are not supported. If necessary, follow the instructions
on the `H2O Download page <http://h2o.ai/download>`__ for your platform to upgrade H2O.

.. toctree::
   :maxdepth: 2

   Installation
   Start_Steam
   h2o_driver
   Configurations
