/*
 Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as
 published by the Free Software Foundation, either version 3 of the
 License, or (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

// TODO
// check for timeouts in $.get calls

(function(window) {
  'use strict';
  var outputDomain;
  var API_HOST = 'http://mr-0xcuda:55011';
  var isBinaryPrediction = false;



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

    var form = '<legend>Parameters</legend>';

    for (var i in names) {
      var n = names[i];
      var i1 = Number(i) + 1;
      form += '<div class="form-group row">';
      form += '<label class="col-sm-5 col-md-5 col-lg-4 form-control-label">' + i1 + '. ' + n + '</label> ';
      var domain = domains[i];
      domain = sortValues(domain);
      var card = domain == null ? 0 : domain.length;

      form += '<div class="col-sm-7 col-md-7 col-lg-8">'

      if (card < 2) {
        if (isBinaryPrediction && i1 === 1) {
          form += '<label class="form-control-label file-icon image-picker"><span class="glyphicon glyphicon-folder-open circle-icon" aria-hidden="true"></span><span id="image-preview-name"></span></label>';
          form += '<input class="image-file" type="file" name="' + n + '" onchange="readURL(this);">';
          form += '<img id="image-preview"/>';
        } else {
          form += '<input class="form-control" type="text" name="' + n + '" oninput="updateUrl(event);">';
        }
      } else if (card <= 2) {
        for (var i = 0; i < card; i += 1) {
          form += '<input type="radio" name="' + n + '" value="' + domain[i] + '" onclick="updateUrl(event);"> ' + domain[i] + '</input>\n';
        }
      } else {
        if (!isBinaryPrediction) {
          form += '<select name="' + n + '" class="form-control" onchange="updateUrl(event);">';
          form += '<option value=""></option>';
          for (var v of domain) {
            form += '<option value="' + v + '">' + v + '</option>';
          }
          form += '</select>';
        }
      }
      form += '</div></div>\n';
      if (isBinaryPrediction) {
        break;
      }
    }

    outputDomain = domains[i1];

    if (element != null) {
      element.innerHTML = form;
      if (isBinaryPrediction) {
        $('.image-picker').click(function(){
          $(this).siblings('.image-file').trigger('click');
        });
      }
    }
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
    $.get(API_HOST + '/info', function(data, status) {
      // show result
      if (data.m._problem_type === 'image_classification') {
        isBinaryPrediction = true;
        hideBatch();
      } else {
        $('#query-string-field').show();
      }
      var info = document.querySelector("#fs-params");
      showModel(data, info);
    }, 'json');
  }

  function hideBatch() {
    $('#batch').hide();
  }

  function showResult(div, status, data) {

    var result = '<legend>Model Predictions</legend>'

    if (data.classProbabilities) {


      // binomial and multinomial
      var label = data.label;
      var index = data.labelIndex;
      var probs = data.classProbabilities;
      var prob = probs[index];

      result += '<p>Predicting <span class="labelHighlight">' + label + '</span>';
      if (probs.length == 2) {
        result += ' based on max F1 threshold </p>';
      }
      result += ' </p>';
      result += '<table class="table" id="modelPredictions"> \
                  <thead> \
                    <tr> \
                      <th>Index</th> \
                      <th>Labels</th> \
                      <th>Probability</th> \
                    </tr> \
                   </thead> \
                   <tbody> \
                  ';

      if (isBinaryPrediction) {
        var labelProbabilitiesMapping = [];
        outputDomain.map(function(label, i) {
          var labelProbMap = {};
          labelProbMap.label = outputDomain[i];
          labelProbMap.probability = probs[i];
          if (i === index) {
            labelProbMap.predicted = true;
          }
          labelProbMap.originalIndex = i;
          labelProbabilitiesMapping.push(labelProbMap);
        });
        labelProbabilitiesMapping.sort(function(a, b) {
          return b.probability - a.probability;
        });
        for (var i = 0; i < 5; i++) {
          if (labelProbabilitiesMapping[i].predicted === true) {
            result += '<tr class="rowHighlight">'
          } else {
            result += '<tr>'
          }
          result += '<td>' + labelProbabilitiesMapping[i].originalIndex + '</td><td>' + labelProbabilitiesMapping[i].label + '</td> <td>' + labelProbabilitiesMapping[i].probability.toFixed(4) + '</td></tr>';
        }
      } else {
        for (var label_i in outputDomain) {
          if (parseInt(label_i) === index ){
            result += '<tr class="rowHighlight">'
          } else {
            result += '<tr>'
          }
          result += '<td>' + label_i + '</td><td>' + outputDomain[label_i] + '</td> <td>' + probs[label_i].toFixed(4) + '</td></tr>';
        }
      }

      result += '</tbody></table>';
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

    div.innerHTML = result;
  }

  function showUrl(pardiv, params) {
    // remove empty parameters returned by serialize.
    params = params.replace(/[\w]+=&/g, "").replace(/&?[\w]+=$/g, "");
    var url = "http://" + window.location.host + "/predict?" + params;
    pardiv.innerHTML = '<a href="' + url + '" target="_blank"><code>' + url + '</code>';
  }

  function showCurl(pardiv, params, cmd) {
    $('#curl-field').show();
    params = params.replace(/'/g, "\\'") // quote quotes
    var url = null;
    if (isBinaryPrediction) {
      pardiv.innerHTML = '<code>curl -X POST --form ' + params + ' ' + cmd + '</code>';
    } else {
      url = "http://" + window.location.host + cmd;
      pardiv.innerHTML = '<code>curl -X POST --data \'' + params + '\' ' + url + '</code>';
    }
  }

  function predResults(params) {
    var div = document.querySelector("#modelPredictions");
    var path = '/predict?';
    if (isBinaryPrediction) {
      path = '/predictbinary';
    }
    var cmd = API_HOST + path + params;
    console.log(cmd);
    if (isBinaryPrediction) {
      var form = $('#allparams');
      var data = new FormData();
      $.each($(form).find('input[type="file"]')[0].files, function(i, file) {
        data.append('binary_C' + (i + 1), file);
      });
      $.ajax({
        url: API_HOST + path,
        data: data,
        cache: false,
        contentType: false,
        processData: false,
        type: 'POST',
        success: function(data, status){
          console.log(data, status);
          showResult(div, status, JSON.parse(data));
          var pardiv = document.querySelector(".curl");
          showCurl(pardiv, 'binary_C1=@' + $(form).find('input[type="file"]')[0].files[0].name, API_HOST + path);
        }
      })
        .fail(function(data, status, error) {
          var down = "<b>GET to " + cmd + " Failed</b> " + error;
          div.innerHTML = down + "<br><b>status</b> " + status;
        });
    } else {
      $.get(cmd, function(data, status) {
        showResult(div, status, data);
      }, 'json')
        .fail(function(data, status, error) {
          var down = "<b>GET to " + cmd + " Failed</b> " + error;
          div.innerHTML = down + "<br><b>status</b> " + status;
        });
    }
  }

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

  function predResultsPost(params) {
    var pardiv = document.querySelector(".curl");
    var cmd = API_HOST + '/pypredict';
    showCurl(pardiv, params, cmd);

    var div = document.querySelector(".results");
    $.post(cmd, params, function(data, status) {
      showResult(div, status, data);
    }, 'json')
      .fail(function(data, status, error) {
        var down = "<b>POST to " + cmd + " Failed</b> " + error;
        div.innerHTML = down + "<br><b>status</b> " + status;
      });

  }

  window.runpredpost = function(form) {
    predResultsPost(form.p.value);
    showStatistics();
  };

  function predResultsPostJar(params) {
    var pardiv = document.querySelector(".curl");
    var cmd = API_HOST + '/predict';
    showCurl(pardiv, params, cmd);

    var div = document.querySelector(".results");
    if (params == null || params == "") {
      div.innerHTML = "No input";
      return;
    }
    $.post(cmd, params, function(data, status) {
      showResult(div, status, data);
    }, 'json')
      .fail(function(data, status, error) {
        var down = "<b>POST to " + cmd + " Failed</b> " + error;
        div.innerHTML = down + "<br><b>status</b> " + status;
      });

  }

  window.runpredpostjar = function(form) {
    predResultsPostJar(form.p.value);
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

  function showStats(div, data) {
    var dayMs = 1000 * 60 * 60 * 24;
    var upDays = Number(data.upTimeMs) / dayMs;
    var lastTimeAgoDays = Number(data.lastTimeAgoMs) / dayMs;

    var s = `<legend>Model Runtime Stats</legend>
            <table class="table noBorders">
              <tbody>`;

    s += '<tr class="stat-group"><td>Service started</td> <td>' + data.startTimeUTC + '</td></tr>';
    s += '<tr class="stat-group"><td>Uptime</td><td>' + duration(upDays) + '</td></tr>';


    var n = Number(data.prediction.count);
    var warmupCount = data.warmUpCount;
    if (n > 0) {
      s += '<tr class="stat-group"><td> Last prediction</td><td>' + data.lastTimeUTC + '</td></tr>';
      s += '<tr class="stat-group"><td> </td><td>' + duration(lastTimeAgoDays) + ' ago</td></tr>';
      s += '<tr class="stat-pad"><td> </td></tr>'
      s += showOneStat('Prediction', data.prediction, warmupCount);
      s += showStat(data.get, 'Get', warmupCount);
      s += showStat(data.post, 'Post', warmupCount);
      s += showStat(data.pythonget, 'Python Get', warmupCount);
      s += showStat(data.pythonpost, 'Python Post', warmupCount);
    }
    s += '</tbody></table>'
    div.innerHTML = s;
  }


  function showStatistics() {
    var cmd = API_HOST + '/stats';
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

    showInputParameters();

    showStatistics();

    $('#stats-btn').click(function() {
      window.open(API_HOST + '/stats', '_blank');
    });

    $('.file-icon').click(function(){
      $(this).closest('.input-inline-row').find("input[type='file']").trigger('click');
    });

    $('.file-input').click(function(){
      $(this).closest('.input-inline-row').find("input[type='file']").trigger('click');
    });

    $('input[type="file"]').change(function(){

      var names = [];
      for (var i = 0; i < $(this).get(0).files.length; ++i) {
        names.push( $(this).get(0).files[i].name );
      }

      $(this).closest('.input-inline-row').find("input[type='text']").val(names);
    });

    $('#batch-predict').click(function(){
      if ( !$('#batch form input[name="inpfile"]').val() ){
        alert('You must select a files with multiple lines of JSON predictions.');
      }

      $('#batch form input[name="inpfile"]').get(0).files = $('#batch form input[name="inpfile"]').get(0).files;

      $('#batch form').submit();
    });

    $('#batch-reset').click(function(){
      $('input[type="text"]').val('');
      $('input[type="file"]').val('');
    });

    $('#predict-btn').click(function(){
      $('#allparams').submit();
    })

    $('#reset-btn').click(function(){
      $('#allparams')[0].reset();
    })
  });

  window.readURL = function(input) {
    if (input.files && input.files[0]) {
      var name = input.files[0].name;
      var reader = new FileReader();

      reader.onload = function (e) {
        $('#image-preview')
          .attr('src', e.target.result)
          .css({
            maxWidth: '100%'
          });
        $('#image-preview-name')
          .text(name);
      };

      reader.readAsDataURL(input.files[0]);
    }
  };
})(window);