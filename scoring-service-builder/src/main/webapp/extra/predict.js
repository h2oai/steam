// predict.js
// H2O Scoring Service

// TODO
// check for timeouts in $.get calls

// Sort values if not sorted properly
// This is needed because c-1, c-2, ... are sorted alphabetically which is wrong
function sortValues(values) {
  if (values == null || values.length < 2) return values;

  var newvalues = [];
  if (values[0] == "c-1") {
    ok = true;
    // check if this is c-1, c-2 etc and make sure it's sorted correctly
    newvalues = [];
    for (i = 1; i <= values.length; i += 1) {
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
  result = !ok ? newvalues : newvalues.sort();
  return result;
}

// Form for model, shown in element
function showModel(model, element) {
    var names = model["m"]["_names"];
    var domains = model["m"]["_domains"];
    var domainMap = model["domainMap"];

    // form = '<form id="allparams" action="" method="get">';
    // form += '<fieldset class="form-group"><legend>Parameters</legend>'

    form = '<legend>Parameters</legend>'

    for(i in names) {
        n = names[i];
        i1 = Number(i) + 1;
        form += '<div class="form-group row">';
        form += '<label class="col-sm-6 form-control-label">' + i1 + '. ' + n + '</label> ';
        domain = domains[i];
        domain = sortValues(domain);
        card = domain == null ? 0 : domain.length;

        form += '<div class="col-sm-6">'

        if (card < 2)
          form += '<input class="form-control" type="text" name="' + n + '">';
        else if (card <= 8) {
          for (i = 0; i < card; i += 1) {
            form += '<input type="radio" name="' + n + '" value="' + domain[i] + '"> ' + domain[i] + '</input>\n';
          }
        }
        else {
          form += '<select name="' + n + '" class="form-control">';
          form += '<option value=""></option>';
          for (v of domain) {
             form += '<option value="' + v + '">' + v + '</option>';
          }
          form += '</select>';
        }
        form += '</div></div>\n';
    }


    outputDomain = domains[i1];

    element.innerHTML = form;

}

function showInputParameters() {
    $.get('/info', function(data, status) {
    // show result
    info = document.querySelector("#fs-params");
    console.log(info);
    showModel(data, info);
  },'json');
}

function showResult(div, status, data) {

    result = '<legend>Model Predictions</legend>'
    

    if ("classProbabilities" in data) {
        // binomial and multinomial
        label = data["label"];
        index = data["labelIndex"];
        probs = data["classProbabilities"];
        prob = probs[index];

        result +=`<table class="table" id="modelPredictions">
                  <thead> 
                    <tr>
                      <th>Labels</th>
                      <th>Probability</th>
                    </tr>
                   </thead>
                   <tbody>
                  `

        for (label_i in outputDomain ){
            result += '<tr><td>' + outputDomain[label_i] + '</td> <td>' + probs[label_i].toFixed(4) +'</td></tr>';
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
    else
        result = "Can't parse result";

    // result += "<p><code>" + JSON.stringify(data) + "</code>";

    div.innerHTML = result;

   // $('#results').show();
}

function showUrl(pardiv, params) {
  // remove empty parameters returned by serialize.
  params = params.replace(/[\w]+=&/g, "").replace(/&?[\w]+=$/g, "");
  url = "http://" + window.location.host + "/predict?" + params;
  pardiv.innerHTML = '<a href="' + url + '" target="_blank"><code>' + url + '</code>';
}

function showCurl(pardiv, params) {
  // remove empty parameters returned by serialize.
  params = params.replace(/'/g, "\\'") // quote quotes
  url = "http://" + window.location.host + "/pypredict";
  pardiv.innerHTML = '<code>curl -X POST --data \'' + params + '\' ' + url + '</code>';
}

function predResults(params) {
  pardiv = document.querySelector(".params");
  // add link that opens in new window
  //showUrl(pardiv, params);

  div = document.querySelector("#modelPredictions");

  cmd = '/predict?' + params;
    $.get(cmd, function(data, status) {
      showResult(div, status, data);
    },'json')
      .fail(function(data, status, error) {
        down = "<b>Service is down</b>";
        div.innerHTML = down + "<br>status " + data.status + " statusText " + data.statusText;
        stats = document.querySelector(".stats");
        stats.innerHTML = down;
        pardiv.innerHTML = "";
      });

}

// function runpred(form) {
//   predResults(form.p.value);
// }

function runpred2(form) {

  if ( $('#queryParams').val() ){
    predResults( $('#queryParams').val() )
  } else {
    predResults($('#allparams').serialize());
  }
  showStatistics();
}
  
function runpred(form) {
  predResults(form.p.value);
  showStatistics();
}

//function runpred2(form) {
//  predResults($('#allparams').serialize());
//}

function predResultsPost(params) {
  pardiv = document.querySelector(".curl");
  showCurl(pardiv, params);

  div = document.querySelector(".results");
  cmd = '/pypredict';
    $.post(cmd, params, function(data, status) {
      showResult(div, status, data);
    },'json')
      .fail(function(data, status, error) {
        down = "<b>POST to /pypredict Failed</b>";
        div.innerHTML = down + "<br>status " + data.status + "<br>statusText " + data.statusText;
        stats = document.querySelector(".stats");
        stats.innerHTML = down;
      });

}

function runpredpost(form) {
  predResultsPost(form.p.value);
  showStatistics();
}


function duration(days) {
  r = days;
  s = "";
  x = Math.floor(r);
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

function showOneStat(label, data, warmUpCount) {
        // s = label + ' (' + data['count'] + ') Last took ' + Number(data['lastMs']).toFixed(3) + ' ms. '
        // + 'Average time ' + Number(data['averageTime']).toFixed(3)
        // + ' (after ' + warmUpCount + ' warmups ' + Number(data['averageAfterWarmupTime']).toFixed(3) + ') ms.';

        s += '<tr></tr>'  ;      
        s += '<tr><td>' + label + '</td><td>Last took: </td><td>' +  Number(data['lastMs']).toFixed(3) +' ms</td></tr>';
        s += '<tr><td>(n=' + data['count'] + ')</td><td>Average time </td><td>' + Number(data['averageTime']).toFixed(3) + 'ms </td></tr>';
        s += '<tr><td></td><td>After ' + warmUpCount + ' warmups: </td><td>' + Number(data['averageAfterWarmupTime']).toFixed(3) + ' ms</td></tr>';
        s += '<tr></tr>'  ; 

        return s;
}

function showStat(stat, textlabel) {
        if (stat['count'] > 0) {
            //s +=  '<p>'
            return showOneStat(textlabel, stat, warmupCount);
        }
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
    dayMs = 1000 * 60 * 60 * 24;
    upDays = Number(data['upTimeMs']) / dayMs;
    lastTimeAgoDays = Number(data['lastTimeAgoMs']) / dayMs;

    s = `<legend>Model Runtime Stats</legend>
            <table class="table noBorders">
              <tbody>`;

    s += '<tr><td>Service started</td> <td>' +   data['startTimeUTC']      +'</td></tr>';
    s += '<tr><td>Uptime</td><td>' +           duration(upDays)          +'</td></tr>';
                
              

    // s = 'Service started ' + data['startTimeUTC'] + '. Uptime ' + duration(upDays) + "."; //upDays.toFixed(3) + ' days. ';
    n = Number(data['prediction']['count']);
    warmupCount = data['warmUpCount'];
    if (n > 0) {
        s += '<tr><td> Last prediction</td><td>' +     data['lastTimeUTC']         +'</td></tr>';
        s += '<tr><td> </td><td>' +                    duration(lastTimeAgoDays)   +' ago</td></tr>';
        s += '<tr><td> </td></tr>'
    //     s +=  '<br>'
    //     + 'Last prediction ' + data['lastTimeUTC'] + ', ' + duration(lastTimeAgoDays) + ' ago.'//lastTimeAgoDays.toFixed(3) + ' days ago.'
    //     +  '<p>'
        showOneStat('Prediction', data['prediction'], warmupCount);
        showStat(data['get'], 'Get');
        showStat(data['post'], 'Post');
        showStat(data['pythonget'], 'Python Get');
        showStat(data['pythonpost'], 'Python Post');
    }

    // url = window.location.href + "stats";
    //s += '<p>More statistics at <code><a href="' + url + '" target="_blank">' + url + '</a>';
    
    s+= '</tbody></table>'

    // url = "http://" + window.location.host + "/stats";
    // s += '<p>More statistics at <code><a href="' + url + '" target="_blank">' + url + '</a>';

    div.innerHTML = s;
}



function showStatistics() {
    cmd = '/stats';
    res = $.get(cmd, function(data, status) {
        divs = document.querySelector("#modelStats");
        showStats(divs, data);
       },'json');
}


// main


$(document).ready(function(){

  $('#url-prefix').append( window.location.href + 'predict?');

  // $('#results').hide();

  showInputParameters();

  showStatistics();

  // $('#reset-btn').click(function(){
  //   $('#results').hide();
  // })

  $('#stats-btn').click(function(){
    window.open('/stats', '_blank');
  })
})



