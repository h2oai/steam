# Deployment

The Deployment page lists all available deployed servie. 


## <a name="makepredictions"></a>Making Predictions

1. To reach the scoring service, click the IP address link listed under the Deployed Services. This opens Steam Prediction Service tool. The fields that display on the Prediction Service tool are automatically populated with field information from the deployed model.

 ![Prediction Services tool](images/prediction_service.png)

2. Make predictions by specifying input values based on column data from the original dataset. This automatically populates the fields in the query string. (Note that you can optionally include input parameters directly in the query string instead of specifying parameters.)

3. Click **Predict** when you are done. 

>**Note**: Use the **Clear** button to clear all entries and begin a new prediction. Use the **More Stats** button to view additional statistics about the scoring service results.