// predict.js
// H2O Scoring Service

// TODO
// check for timeouts in $.get calls

(function(window) {
  'use strict';
  var outputDomain;
  // Sort values if not sorted properly
  // This is needed because c-1, c-2, ... are sorted alphabetically which is wrong
  function sortValues(values) {
    if (values == null || values.length < 2) return values;

    var newvalues = [];
    var ok = false;
    if (values[0] == "c-1") {
      ok = true;
      // check if this is c-1, c-2 etc and make sure it's sorted correctly
      newvalues = [];
      for (var i = 1; i <= values.length; i += 1) {
        if (values[i - 1] != "c-" + i) {
          ok = false;
        }
        newvalues.push("c-" + i);
      }
    }
    else {
      newvalues = values;
      ok = true;
    }
    var result = !ok ? newvalues : newvalues.sort();
    return result;
  }

// Form for model, shown in element
  function showModel(model, element) {
    var names = model.m._names;
    var domains = model.m._domains;
    var domainMap = model.domainMap;

    // form = '<form id="allparams" action="" method="get">';
    // form += '<fieldset class="form-group"><legend>Parameters</legend>'

    var form = '<legend>Parameters</legend>';

    for (var i in names) {
      var n = names[i];
      var i1 = Number(i) + 1;
      form += '<div class="form-group row">';
      form += '<label class="col-sm-3 col-md-3 form-control-label">' + i1 + '. ' + n + '</label> ';
      var domain = domains[i];
      domain = sortValues(domain);
      var card = domain == null ? 0 : domain.length;

      form += '<div class="col-sm-9 col-md-9">'

      if (card < 2) {
        form += '<input class="form-control" type="text" name="' + n + '" oninput="updateUrl(event);">';
      } else if (card <= 8) {
        for (var i = 0; i < card; i += 1) {
          form += '<input type="radio" name="' + n + '" value="' + domain[i] + '" onclick="updateUrl(event);"> ' + domain[i] + '</input>\n';
        }
      } else {
        form += '<select name="' + n + '" class="form-control" onchange="updateUrl(event);">';
        form += '<option value=""></option>';
        for (var v of domain) {
          form += '<option value="' + v + '">' + v + '</option>';
        }
        form += '</select>';
      }
      form += '</div></div>\n';
    }


    outputDomain = domains[i1];

    element.innerHTML = form;

  }

  var params = {};
  window.updateUrl = function(event) {
    params[event.target.name] = event.target.value;
    var queryString = '';
    for (var key in params) {
      if (params.hasOwnProperty(key) && params[key] !== '') {
        queryString += key + '=' + params[key] + '&';
      }
    }
    queryString = queryString.substring(0, queryString.length - 1);
    $('#queryParams').val(queryString);
    if (queryString === '') {
      window.setEmptyUrl();
    } else {
      $('#query-link').css({
        cursor: 'pointer'
      })
      .unbind('click')
      .attr('href', 'predict?' + queryString);
    }
  };

  window.setEmptyUrl = function() {
    for (var key in params) {
      if (params.hasOwnProperty(key) && params[key] !== '') {
        params[key] = '';
      }
    }
    $('#query-link').css({
      cursor: 'not-allowed'
    })
    .bind('click', function(e) {
      e.preventDefault();
    })
    .attr('href', '');
  };

  function showInputParameters() {
    $.get('/info', function(data, status) {
      // show result
      var info = document.querySelector("#fs-params");
      showModel(data, info);
    }, 'json');
  }

  function showResult(div, status, data) {

    var result = '<legend>Model Predictions</legend>'

    if ("classProbabilities" in data) {

      result += `<p>The following predictions are based on threshold F1 </p>`
      // binomial and multinomial
      var label = data["label"];
      var index = data["labelIndex"];
      var probs = data["classProbabilities"];
      var prob = probs[index];

      result += `<table class="table" id="modelPredictions">
                  <thead> 
                    <tr>
                      <th>Labels</th>
                      <th>Probability</th>
                    </tr>
                   </thead>
                   <tbody>
                  `

      for (var label_i in outputDomain) {
        result += '<tr><td>' + outputDomain[label_i] + '</td> <td>' + probs[label_i].toFixed(4) + '</td></tr>';
      }

      result += '</tbody></table>';

      // result = "Label <b>" + label + "</b> with probability <b>" + (prob * 100.0).toFixed(1) + "%</b>.<p>"
      //     + "Output Labels: [" + outputDomain + "]<br>"
      //     + "Class Probabilities: [" + probs + "]<br>"
      //     + "Label Index: " + index;
    }
    else if ("cluster" in data) {
      // clustering result
      result = "Cluster <b>" + data["cluster"] + "</b>";
    }
    else if ("value" in data) {
      // regression result
      result = "Value <b>" + data["value"] + "</b>";
    }
    else if ("dimensions" in data) {
      // dimensionality reduction result
      result = "Dimensions <b>" + data["dimensions"] + "</b>";
    }
    else {
      result = "Can't parse result: " + data;
    }

    // result += "<p><code>" + JSON.stringify(data) + "</code>";

    div.innerHTML = result;
}

  function showUrl(pardiv, params) {
    // remove empty parameters returned by serialize.
    params = params.replace(/[\w]+=&/g, "").replace(/&?[\w]+=$/g, "");
    var url = "http://" + window.location.host + "/predict?" + params;
    pardiv.innerHTML = '<a href="' + url + '" target="_blank"><code>' + url + '</code>';
  }

  function showCurl(pardiv, params) {
    // remove empty parameters returned by serialize.
    params = params.replace(/'/g, "\\'") // quote quotes
    var url = "http://" + window.location.host + "/pypredict";
    pardiv.innerHTML = '<code>curl -X POST --data \'' + params + '\' ' + url + '</code>';
  }

  function predResults(params) {
    // var pardiv = document.querySelector(".params");
    // add link that opens in new window
    //showUrl(pardiv, params);

    var div = document.querySelector("#modelPredictions");

    var cmd = '/predict?' + params;
    $.get(cmd, function(data, status) {
      showResult(div, status, data);
    }, 'json')
      .fail(function(data, status, error) {
        var down = "<b>Service is down</b>";
        // div.innerHTML = down + "<br>status " + data.status + " statusText " + data.statusText;
        div.innerHTML = 'Error: ' + data.statusText;
        var stats = document.querySelector("#modelStats");
        // stats.innerHTML = down;
        // pardiv.innerHTML = "";
      });

  }

// function runpred(form) {
//   predResults(form.p.value);
// }

  window.runpred2 = function(form) {

    if ($('#queryParams').val()) {
      predResults($('#queryParams').val())
    } else {
      predResults($('#allparams').serialize());
    }
    showStatistics();
  };

  function runpred(form) {
    predResults(form.p.value);
    showStatistics();
  }

//function runpred2(form) {
//  predResults($('#allparams').serialize());
//}

  function predResultsPost(params) {
    var pardiv = document.querySelector(".curl");
    showCurl(pardiv, params);

    var div = document.querySelector(".results");
    var cmd = '/pypredict';
    $.post(cmd, params, function(data, status) {
      showResult(div, status, data);
    }, 'json')
      .fail(function(data, status, error) {
        var down = "<b>POST to /pypredict Failed</b>";
        div.innerHTML = down + "<br>status " + data.status + "<br>statusText " + data.statusText;
        var stats = document.querySelector(".stats");
        stats.innerHTML = down;
      });

  }

  window.runpredpost = function(form) {
    predResultsPost(form.p.value);
    showStatistics();
  };


  function duration(days) {
    var r = days;
    var s = "";
    var x = Math.floor(r);
    if (x >= 1) {
      s += x + " d ";
    }
    r = (r - x) * 24; // hours
    x = Math.floor(r);
    if (x >= 1) {
      s += x + " h ";
    }
    r = (r - x) * 60; // minutes
    x = Math.floor(r);
    if (x >= 1) {
      s += x + " m ";
    }
    r = (r - x) * 60; // seconds
    // x = Math.floor(r);
    s += r.toFixed(0) + " s"
    return s;
  }

  function showOneStat(label, data, warmupCount) {
    // s = label + ' (' + data['count'] + ') Last took ' + Number(data['lastMs']).toFixed(3) + ' ms. '
    // + 'Average time ' + Number(data['averageTime']).toFixed(3)
    // + ' (after ' + warmUpCount + ' warmups ' + Number(data['averageAfterWarmupTime']).toFixed(3) + ') ms.';
    var newRow = '';
    newRow += '<tr class="stat-pad"></tr>';
    newRow += '<tr class="stat-group"><td>' + label + '</td><td>Last took: </td><td>' + Number(data['lastMs']).toFixed(3) + ' ms</td></tr>';
    newRow += '<tr class="stat-group"><td>(n=' + data['count'] + ')</td><td>Average time </td><td>' + Number(data['averageTime']).toFixed(3) + ' ms </td></tr>';
    newRow += '<tr class="stat-group"><td></td><td>After ' + warmupCount + ' warmups: </td><td>' + Number(data['averageAfterWarmupTime']).toFixed(3) + ' ms</td></tr>';
    newRow += '<tr></tr>';

    return newRow;
  }

  function showStat(stat, textlabel, warmupCount) {
    if (stat['count'] > 0) {
      //s +=  '<p>'
      return showOneStat(textlabel, stat, warmupCount);
    }
    return '';
  }

// function showStats(div, data) {
//     dayMs = 1000 * 60 * 60 * 24;
//     upDays = Number(data['upTimeMs']) / dayMs;
//     lastTimeAgoDays = Number(data['lastTimeAgoMs']) / dayMs;
//     s = 'Service started ' + data['startTimeUTC'] + '. Uptime ' + duration(upDays) + "."; //upDays.toFixed(3) + ' days. ';
//     n = Number(data['prediction']['count']);
//     warmupCount = data['warmUpCount'];
//     if (n > 0) {
//         s +=  '<br>'
//         + 'Last prediction ' + data['lastTimeUTC'] + ', ' + duration(lastTimeAgoDays) + ' ago.'//lastTimeAgoDays.toFixed(3) + ' days ago.'
//         +  '<p>'
//         + showOneStat('Prediction', data['prediction'], warmupCount);
//         showStat(data['get'], 'Get');
//         showStat(data['post'], 'Post');
//         showStat(data['pythonget'], 'Python Get');
//         showStat(data['pythonpost'], 'Python Post');
//     }
//     url = window.location.href + "stats";
//     s += '<p>More statistics at <code><a href="' + url + '" target="_blank">' + url + '</a>';
//     div.innerHTML = s;
// }

  function showStats(div, data) {
    var dayMs = 1000 * 60 * 60 * 24;
    var upDays = Number(data.upTimeMs) / dayMs;
    var lastTimeAgoDays = Number(data.lastTimeAgoMs) / dayMs;

    var s = `<legend>Model Runtime Stats</legend>
            <table class="table noBorders">
              <tbody>`;

    s += '<tr class="stat-group"><td>Service started</td> <td>' + data.startTimeUTC + '</td></tr>';
    s += '<tr class="stat-group"><td>Uptime</td><td>' + duration(upDays) + '</td></tr>';


    // s = 'Service started ' + data['startTimeUTC'] + '. Uptime ' + duration(upDays) + "."; //upDays.toFixed(3) + ' days. ';
    var n = Number(data.prediction.count);
    var warmupCount = data.warmUpCount;
    if (n > 0) {
      s += '<tr class="stat-group"><td> Last prediction</td><td>' + data.lastTimeUTC + '</td></tr>';
      s += '<tr class="stat-group"><td> </td><td>' + duration(lastTimeAgoDays) + ' ago</td></tr>';
      s += '<tr class="stat-pad"><td> </td></tr>'
      //     s +=  '<br>'
      //     + 'Last prediction ' + data['lastTimeUTC'] + ', ' + duration(lastTimeAgoDays) + ' ago.'//lastTimeAgoDays.toFixed(3) + ' days ago.'
      //     +  '<p>'
      s += showOneStat('Prediction', data.prediction, warmupCount);
      s += showStat(data.get, 'Get', warmupCount);
      s += showStat(data.post, 'Post', warmupCount);
      s += showStat(data.pythonget, 'Python Get', warmupCount);
      s += showStat(data.pythonpost, 'Python Post', warmupCount);
    }

    // url = window.location.href + "stats";
    //s += '<p>More statistics at <code><a href="' + url + '" target="_blank">' + url + '</a>';

    s += '</tbody></table>'

    // url = "http://" + window.location.host + "/stats";
    // s += '<p>More statistics at <code><a href="' + url + '" target="_blank">' + url + '</a>';

    div.innerHTML = s;
  }


  function showStatistics() {
    var cmd = '/stats';
    var res = $.get(cmd, function(data, status) {
      var divs = document.querySelector("#modelStats");
      showStats(divs, data);
    }, 'json');
  }


// main

  window.clearParameters = function() {
    var form = $('#allparams');
    form.find('input[type="text"]').not('#queryParams').val('');
    form.find('input[type="radio"]').prop('checked', false);
    form.find('select').prop('selectedIndex', 0);
  };

  $(document).ready(function() {

    $('#url-prefix').append(window.location.href + 'predict?');
    window.setEmptyUrl();

    // $('#results').hide();

    showInputParameters();

    showStatistics();

    // $('#reset-btn').click(function(){
    //   $('#results').hide();
    // })

    $('#stats-btn').click(function() {
      window.open('/stats', '_blank');
    });
  });
})(window);