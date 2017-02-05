Adding an H2O Driver
====================

Now that Steam is installed and running, the next step is to add an H2O Driver. 

**Note**: Currently, only the H2O bleeding edge release is supported on Steam. 

1. On your local machine, download the h2odriver from the `H2O Download page <http://h2o.ai/download>`__. Be sure to select your version of Hadoop. For example:

  ::

    wget http://h2o-release.s3.amazonaws.com/h2o/master/3756/h2o-3.11.0.3756-hdp2.2.zip

2. In the Steam UI, navigate to the **Clusters** page and select **Launch New Cluster**.

3. In the H2O Version section of the **Launch New Cluster page**, click the **Choose File** button and browse to the H2O driver that you just downloaded.

4. Click **Upload Engine** to add the egine to Steam.

   .. figure:: images/upload_engine.png
      :alt: Upload Engine
    
A message will display indicating that the engine was successfully uploaded.
