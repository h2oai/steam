/**
 * @author justin on 9/30/16.
 */
var AWS = require('aws-sdk');
var express = require('express');
var bodyParser = require('body-parser');
var cors = require('cors');
var _ = require('lodash');
var app = express();

AWS.config.region = 'us-east-1';

app.use(cors());
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

app.post('/steam', function(req, res) {
  var ecs = new AWS.ECS();
  var params = {
    cluster: 'steam',
    taskDefinition: 'steam:12'
  };

  ecs.runTask(params, function(err, data) {
    if (err) {
      console.log(err, err.stack);
    } else {
      console.log('runtask', data.tasks[0].containers[0].taskArn);
      var id = new Buffer(data.tasks[0].containers[0].taskArn).toString('base64');
    }
    res.send(JSON.stringify({id: id}));
  });
});

app.get('/steam', function(req, res) {
  var ecs = new AWS.ECS();
  var params = {
    cluster: 'steam'
  };

  ecs.listTasks(params, function(err, data) {
    if (err) {
      console.log(err, err.stack);
    }
    console.log(data);
    var tasks = data.taskArns;
    tasks = _.map(tasks, function(a) {
      return new Buffer(a).toString('base64')
    });
    res.send(JSON.stringify(tasks));
  });
});

app.delete('/steam/:id', function(req, res) {
  var ecs = new AWS.ECS();
  var params = {
    task: id,
    cluster: 'steam'
  };
  ecs.stopTask({

  })
});

app.listen(3000, function() {
  console.log("Started on PORT 3000");
});

