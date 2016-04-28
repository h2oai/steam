// predict.js
// H2O Prediction Service

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
//    console.log("names = " + JSON.stringify(names));
    var domainMap = model["domainMap"];

    form = '<form id="allparams" action="" method="get">';
    for(i in names) {
        n = names[i];
        i1 = i + 1;
        form += '<p><b>' + i1 + '. ' + n + '</b> ';
        domain = domains[i];
        domain = sortValues(domain);
        card = domain == null ? 0 : domain.length;
        //console.log((i + 1) + "  " + card);
//        console.log(domain, card);
        //console.log("cardinality " + card);
        if (card < 2)
          form += '<input type="text" name="' + n + '">';
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

//    form += '<input type="button" name="okbutton" value="PREDICT" onClick="runpred(this.form)"><p>' +
    form += ' <input type="reset" value="CLEAR">';

    form += '</form>';
//    console.log(form);
    element.innerHTML += form;

    // Alternative query string form
    form += '<p>';

}

function showInfo() {
$.get('/info', function(data, status) {
    // show result
    var info = document.querySelector(".input");
    showModel(data, info);
  },'json');
}


function runpred(form) {
  params = form.p.value;
  console.log(params);
  pardiv = document.querySelector(".params");

  // add link that opens in new window
  url = window.location.href + "predict?" + params;
  pardiv.innerHTML = '<a href="' + url + '" target="_blank"><code>' + url + '</code>';

  cmd = '/predict?' + params;
  res = $.get(cmd, function(data, status) {
     var div = document.querySelector(".results");
     div.innerHTML = status + "   " + JSON.stringify(data);
  },'json');

}

function runpred2(form) {
  var params = $('#allparams').serialize();
  // remove empty parameters returned by serialize.
  params = params.replace(/[\w]+=&/g, "").replace(/&?[\w]+=$/g, "");
  console.log(params);
  pardiv = document.querySelector(".params");

  // add link that opens in new window
  url = window.location.href + "predict?" + params;
  pardiv.innerHTML = '<a href="' + url + '" target="_blank"><code>' + url + '</code>';

  cmd = '/predict?' + params;
  res = $.get(cmd, function(data, status) {
     div = document.querySelector(".results");
     label = data["label"];
     index = Number(data["index"]);
     probs = data["classProbabilities"];
     prob = probs[index];
     console.log(probs);
     console.log(prob);
     div.innerHTML = status + "   " + JSON.stringify(data);
  },'json');

}

// main
showInfo();



