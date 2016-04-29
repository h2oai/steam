// predict.js
// H2O Scoring Service

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

    form = '<form id="allparams" action="" method="get">';
    for(i in names) {
        n = names[i];
        i1 = Number(i) + 1;
        form += '<p><b>' + i1 + '. ' + n + '</b> ';
        domain = domains[i];
        domain = sortValues(domain);
        card = domain == null ? 0 : domain.length;

        if (card < 2)
          form += '<input type="text" size=40 name="' + n + '">';
        else if (card <= 12) {
          for (i = 0; i < card; i += 1) {
            form += '<input type="radio" name="' + n + '" value="' + domain[i] + '"> ' + domain[i] + '</input>\n';
          }
        }
        else {
          form += '<select name="' + n + '">';
          form += '<option value=""></option>';
          for (v of domain) {
             form += '<option value="' + v + '">' + v + '</option>';
          }
          form += '</select>';
        }
        form += '</p>\n';
    }
    form += '<input type="button" name="okbutton" value="PREDICT" onClick="runpred2(this.form)">';

    form += ' <input type="reset" value="CLEAR">';

    form += '</form>';
    element.innerHTML += form;

    form += '<p>';

}

function showInputParameters() {
    $.get('/info', function(data, status) {
    // show result
    var info = document.querySelector(".input");
    showModel(data, info);
  },'json');
}

function showResult(div, data) {
    label = data["label"];
    index = data["labelIndex"];
    probs = data["classProbabilities"];
    prob = probs[index];

    result = "Label: <b>" + label + "</b><br>" +
      "Probability: <b>" + (prob * 100.0).toFixed(1) + "%</b><p>" +
      "Class Probabilities: [" + probs + "]<br>" +
      "Label Index: " + index;
    div.innerHTML = result;
}

function showUrl(pardiv, params) {
  url = window.location.href + "predict?" + params;
  // remove empty parameters returned by serialize.
  params.replace(/[\w]+=&/g, "").replace(/&?[\w]+=$/g, "");
  pardiv.innerHTML = '<a href="' + url + '" target="_blank"><code>' + url + '</code>';
}

function runpred(form) {
  params = form.p.value;
  pardiv = document.querySelector(".params");
  // add link that opens in new window
  showUrl(pardiv, params);

  cmd = '/predict?' + params;
  div = document.querySelector(".results");
  res = $.get(cmd, function(data, status) {
    showResult(div, data);
  },'json');

}

function runpred2(form) {
  var params = $('#allparams').serialize();
  pardiv = document.querySelector(".params");
  // add link that opens in new window
  showUrl(pardiv, params);

  cmd = '/predict?' + params;
  div = document.querySelector(".results");
  res = $.get(cmd, function(data, status) {
    showResult(div, data);
   },'json');

}

// main
showInputParameters();



