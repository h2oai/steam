/**
 * @author justin on 9/30/16.
 */

var http = require('http');
var url = require('url');
var httpProxy = require('http-proxy');
var fs = require('fs');
var express = require('express');
var _ = require('lodash');


var DOCKER_GEN_FILE = './docker-gen/output';
var containers = [];

fs.watchFile(DOCKER_GEN_FILE, function (curr, prev) {
  fs.readFile(DOCKER_GEN_FILE, 'utf8', function(err, data) {
    console.log(data);
    containers = JSON.parse(data);
  });
});

var proxy = httpProxy.createProxyServer({});

function parseCookies (request) {
  var list = {},
    rc = request.headers.cookie;

  rc && rc.split(';').forEach(function( cookie ) {
    var parts = cookie.split('=');
    list[parts.shift().trim()] = decodeURI(parts.join('='));
  });

  return list;
}

var server = http.createServer(function(req, res) {
  var requestType = req.headers['Content-Type'];
  if ((req.method == 'POST' || req.method == 'PUT') && requestType === 'multipart/form-data') {
    req.body = '';

    req.addListener('data', function(chunk) {
      req.body += chunk;
    });

    req.addListener('end', function() {
      res.end();
    });
  } else {
    var queryObject = url.parse(req.url,true).query;
    if (queryObject.id) {
      var id = Buffer.from(queryObject.id, 'base64').toString();
    }
    var cookies = parseCookies(req);
    var target = null;
    var y = [];
    if (id != undefined) {
      res.setHeader('Set-Cookie', 'id=' + id);
      y = containers.filter(function(a) {
        if (a.Labels) {
          return a.Labels['com.amazonaws.ecs.task-arn'] === id;
        }
      });
      if (y[0]) {
        target = y[0].IP;
      }
    } else if (cookies.id) {
      y = containers.filter(function(a) {
        if (a.Labels) {
          return a.Labels['com.amazonaws.ecs.task-arn'] === cookies.id;
        }
      });
      if (y[0]) {
        target = y[0].IP;
      }
    }
    if (target) {
      proxy.web(req, res, { target: 'http://' + target + ':9002'});
    } else {
      res.end();
    }
  }
});

console.log("listening on port 9000");
server.listen(9000);