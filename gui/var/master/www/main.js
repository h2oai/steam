/// <reference path="../typings/tsd.d.ts" />
/// <reference path="references.ts" />
"use strict";
var Proxy;
(function (Proxy) {
    var _rpcId = 0;
    function nextId() {
        return ++_rpcId;
    }
    function _invoke(settings, go) {
        var p = $.ajax(settings);
        p.done(function (data, status, xhr) {
            if (data.error) {
                return go(new Error(data.error), null);
            }
            go(null, data.result);
        });
        p.fail(function (xhr, status, error) {
            // XXX handle 401
            var res = xhr.responseJSON;
            if (res && res.error) {
                return go(new Error(res.error), null);
            }
            // special-case net::ERR_CONNECTION_REFUSED
            if (status == 'error' && xhr.status == 0) {
                return go(new Error("Could not connect to Steam. The server is currently unresponsive."), null);
            }
            go(new Error("HTTP connection failure: status=" + status + ", code=" + xhr.status + ", error=" + (error ? error : '?')), null);
        });
    }
    function invoke(method, param, headers, go) {
        var req = {
            method: "web." + method,
            params: [param],
            id: nextId()
        };
        var settings = {
            url: "/web",
            type: "POST",
            data: JSON.stringify(req),
            contentType: "application/json; charset=utf-8",
            dataType: "json"
        };
        if (headers) {
            settings.headers = headers;
        }
        _invoke(settings, go);
    }
    function upload(formData, go) {
        var settings = {
            url: "/upload",
            type: "POST",
            data: formData,
            cache: false,
            contentType: false,
            processData: false
        };
        _invoke(settings, go);
    }
    Proxy.upload = upload;
    function Call(method, param, go) {
        invoke(method, param, null, go);
    }
    Proxy.Call = Call;
    function authenticate(username, password, go) {
        var headers = {
            "Authorization": "Custom " + username + " " + password
        };
        invoke("Ping", { status: true }, headers, go);
    }
    Proxy.authenticate = authenticate;
})(Proxy || (Proxy = {}));
// ----------------------------------
// --- Generated with go:generate ---
// ---        DO NOT EDIT         ---
// ----------------------------------
var Proxy;
(function (Proxy) {
    // --- Consts ---
    Proxy.CloudStarted = "Started";
    Proxy.CloudHealthy = "Healthy";
    Proxy.CloudStopped = "Stopped";
    Proxy.CloudUnknown = "Unknown";
    Proxy.ScoringServiceStarted = "Started";
    Proxy.ScoringServiceStopped = "Stopped";
    // --- Client Stub ---
    function ping(status, go) {
        var req = {
            status: status
        };
        Proxy.Call("Ping", req, function (error, data) {
            return error ? go(error, null) : go(null, data.status);
        });
    }
    Proxy.ping = ping;
    function startCloud(cloudName, engineName, size, memory, username, go) {
        var req = {
            cloud_name: cloudName,
            engine_name: engineName,
            size: size,
            memory: memory,
            username: username
        };
        Proxy.Call("StartCloud", req, function (error, data) {
            return error ? go(error, null) : go(null, data.cloud);
        });
    }
    Proxy.startCloud = startCloud;
    function stopCloud(cloudName, go) {
        var req = {
            cloud_name: cloudName
        };
        Proxy.Call("StopCloud", req, function (error, data) {
            return error ? go(error) : go(null);
        });
    }
    Proxy.stopCloud = stopCloud;
    function getCloud(cloudName, go) {
        var req = {
            cloud_name: cloudName
        };
        Proxy.Call("GetCloud", req, function (error, data) {
            return error ? go(error, null) : go(null, data.cloud);
        });
    }
    Proxy.getCloud = getCloud;
    function getClouds(go) {
        var req = {};
        Proxy.Call("GetClouds", req, function (error, data) {
            return error ? go(error, null) : go(null, data.clouds);
        });
    }
    Proxy.getClouds = getClouds;
    function getCloudStatus(cloudName, go) {
        var req = {
            cloud_name: cloudName
        };
        Proxy.Call("GetCloudStatus", req, function (error, data) {
            return error ? go(error, null) : go(null, data.cloud);
        });
    }
    Proxy.getCloudStatus = getCloudStatus;
    function deleteCloud(cloudName, go) {
        var req = {
            cloud_name: cloudName
        };
        Proxy.Call("DeleteCloud", req, function (error, data) {
            return error ? go(error) : go(null);
        });
    }
    Proxy.deleteCloud = deleteCloud;
    function getJob(cloudName, jobName, go) {
        var req = {
            cloud_name: cloudName,
            job_name: jobName
        };
        Proxy.Call("GetJob", req, function (error, data) {
            return error ? go(error, null) : go(null, data.job);
        });
    }
    Proxy.getJob = getJob;
    function getJobs(cloudName, go) {
        var req = {
            cloud_name: cloudName
        };
        Proxy.Call("GetJobs", req, function (error, data) {
            return error ? go(error, null) : go(null, data.jobs);
        });
    }
    Proxy.getJobs = getJobs;
    function buildModel(cloudName, dataset, targetName, maxRunTime, go) {
        var req = {
            cloud_name: cloudName,
            dataset: dataset,
            target_name: targetName,
            max_run_time: maxRunTime
        };
        Proxy.Call("BuildModel", req, function (error, data) {
            return error ? go(error, null) : go(null, data.model);
        });
    }
    Proxy.buildModel = buildModel;
    function getModel(modelName, go) {
        var req = {
            model_name: modelName
        };
        Proxy.Call("GetModel", req, function (error, data) {
            return error ? go(error, null) : go(null, data.model);
        });
    }
    Proxy.getModel = getModel;
    function getModels(go) {
        var req = {};
        Proxy.Call("GetModels", req, function (error, data) {
            return error ? go(error, null) : go(null, data.models);
        });
    }
    Proxy.getModels = getModels;
    function getCloudModels(cloudName, go) {
        var req = {
            cloud_name: cloudName
        };
        Proxy.Call("GetCloudModels", req, function (error, data) {
            return error ? go(error, null) : go(null, data.models);
        });
    }
    Proxy.getCloudModels = getCloudModels;
    function getModelFromCloud(cloudName, modelName, go) {
        var req = {
            cloud_name: cloudName,
            model_name: modelName
        };
        Proxy.Call("GetModelFromCloud", req, function (error, data) {
            return error ? go(error, null) : go(null, data.model);
        });
    }
    Proxy.getModelFromCloud = getModelFromCloud;
    function deleteModel(modelName, go) {
        var req = {
            model_name: modelName
        };
        Proxy.Call("DeleteModel", req, function (error, data) {
            return error ? go(error) : go(null);
        });
    }
    Proxy.deleteModel = deleteModel;
    function startScoringService(modelName, port, go) {
        var req = {
            model_name: modelName,
            port: port
        };
        Proxy.Call("StartScoringService", req, function (error, data) {
            return error ? go(error, null) : go(null, data.service);
        });
    }
    Proxy.startScoringService = startScoringService;
    function stopScoringService(modelName, port, go) {
        var req = {
            model_name: modelName,
            port: port
        };
        Proxy.Call("StopScoringService", req, function (error, data) {
            return error ? go(error) : go(null);
        });
    }
    Proxy.stopScoringService = stopScoringService;
    function getScoringService(modelName, go) {
        var req = {
            model_name: modelName
        };
        Proxy.Call("GetScoringService", req, function (error, data) {
            return error ? go(error, null) : go(null, data.service);
        });
    }
    Proxy.getScoringService = getScoringService;
    function getScoringServices(go) {
        var req = {};
        Proxy.Call("GetScoringServices", req, function (error, data) {
            return error ? go(error, null) : go(null, data.services);
        });
    }
    Proxy.getScoringServices = getScoringServices;
    function deleteScoringService(modelName, port, go) {
        var req = {
            model_name: modelName,
            port: port
        };
        Proxy.Call("DeleteScoringService", req, function (error, data) {
            return error ? go(error) : go(null);
        });
    }
    Proxy.deleteScoringService = deleteScoringService;
    function addEngine(engineName, enginePath, go) {
        var req = {
            engine_name: engineName,
            engine_path: enginePath
        };
        Proxy.Call("AddEngine", req, function (error, data) {
            return error ? go(error) : go(null);
        });
    }
    Proxy.addEngine = addEngine;
    function getEngine(engineName, go) {
        var req = {
            engine_name: engineName
        };
        Proxy.Call("GetEngine", req, function (error, data) {
            return error ? go(error, null) : go(null, data.engine);
        });
    }
    Proxy.getEngine = getEngine;
    function getEngines(go) {
        var req = {};
        Proxy.Call("GetEngines", req, function (error, data) {
            return error ? go(error, null) : go(null, data.engines);
        });
    }
    Proxy.getEngines = getEngines;
    function deleteEngine(engineName, go) {
        var req = {
            engine_name: engineName
        };
        Proxy.Call("DeleteEngine", req, function (error, data) {
            return error ? go(error) : go(null);
        });
    }
    Proxy.deleteEngine = deleteEngine;
})(Proxy || (Proxy = {}));
/// <reference path="references.ts" />
/// <reference path="xhr.ts" />
/// <reference path="proxy.ts" />
"use strict";
var Main;
(function (Main) {
    //
    // Prelude
    //
    var Epsilon = 1e-6;
    function uni() {
        var link = null;
        var trigger = function () {
            if (link) {
                link();
            }
        };
        trigger.on = function (f) {
            if (link) {
                throw new Error('Delegate is unicast.');
            }
            link = f;
        };
        trigger.off = function (f) {
            if (f !== link) {
                throw new Error('Invalid subscription.');
            }
            link = null;
        };
        trigger.dispose = function () {
            link = null;
        };
        return trigger;
    }
    Main.uni = uni;
    function uni1() {
        var link = null;
        var trigger = function (v1) {
            if (link) {
                link(v1);
            }
        };
        trigger.on = function (f) {
            if (link) {
                throw new Error('Delegate is unicast.');
            }
            link = f;
        };
        trigger.off = function (f) {
            if (f !== link) {
                throw new Error('Invalid subscription.');
            }
            link = null;
        };
        trigger.dispose = function () {
            link = null;
        };
        return trigger;
    }
    Main.uni1 = uni1;
    function uni2() {
        var link = null;
        var trigger = function (v1, v2) {
            if (link) {
                link(v1, v2);
            }
        };
        trigger.on = function (f) {
            if (link) {
                throw new Error('Delegate is unicast.');
            }
            link = f;
        };
        trigger.off = function (f) {
            if (f !== link) {
                throw new Error('Invalid subscription.');
            }
            link = null;
        };
        trigger.dispose = function () {
            link = null;
        };
        return trigger;
    }
    Main.uni2 = uni2;
    function uni3() {
        var link = null;
        var trigger = function (v1, v2, v3) {
            if (link) {
                link(v1, v2, v3);
            }
        };
        trigger.on = function (f) {
            if (link) {
                throw new Error('Delegate is unicast.');
            }
            link = f;
        };
        trigger.off = function (f) {
            if (f !== link) {
                throw new Error('Invalid subscription.');
            }
            link = null;
        };
        trigger.dispose = function () {
            link = null;
        };
        return trigger;
    }
    Main.uni3 = uni3;
    function uni4() {
        var link = null;
        var trigger = function (v1, v2, v3, v4) {
            if (link) {
                link(v1, v2, v3, v4);
            }
        };
        trigger.on = function (f) {
            if (link) {
                throw new Error('Delegate is unicast.');
            }
            link = f;
        };
        trigger.off = function (f) {
            if (f !== link) {
                throw new Error('Invalid subscription.');
            }
            link = null;
        };
        trigger.dispose = function () {
            link = null;
        };
        return trigger;
    }
    Main.uni4 = uni4;
    function multi3() {
        var links = [];
        var trigger = function (v1, v2, v3) {
            for (var _i = 0, links_1 = links; _i < links_1.length; _i++) {
                var f = links_1[_i];
                f(v1, v2, v3);
            }
        };
        trigger.on = function (f) {
            links.push(f);
        };
        trigger.off = function (f) {
            var i = _.indexOf(links, f);
            if (i < 0) {
                return;
            }
            links.splice(i, 1);
        };
        trigger.dispose = function () {
            links.length = 0;
        };
        return trigger;
    }
    Main.multi3 = multi3;
    function sig(value, equalityComparer) {
        var o = ko.observable(value);
        if (equalityComparer) {
            o.equalityComparer = equalityComparer;
        }
        return o;
    }
    Main.sig = sig;
    function sigs(value) {
        return ko.observableArray(value);
    }
    Main.sigs = sigs;
    function react(s1, a) {
        return s1.subscribe(a);
    }
    Main.react = react;
    function reacts(s1, a) {
        return s1.subscribe(a);
    }
    Main.reacts = reacts;
    function react2(s1, s2, f) {
        return [
            react(s1, function (v1) { f(v1, s2()); }),
            react(s2, function (v2) { f(s1(), v2); })
        ];
    }
    Main.react2 = react2;
    function react3(s1, s2, s3, f) {
        return [
            react(s1, function (v1) { f(v1, s2(), s3()); }),
            react(s2, function (v2) { f(s1(), v2, s3()); }),
            react(s3, function (v3) { f(s1(), s2(), v3); })
        ];
    }
    Main.react3 = react3;
    function react4(s1, s2, s3, s4, f) {
        return [
            react(s1, function (v1) { f(v1, s2(), s3(), s4()); }),
            react(s2, function (v2) { f(s1(), v2, s3(), s4()); }),
            react(s3, function (v3) { f(s1(), s2(), v3, s4()); }),
            react(s4, function (v4) { f(s1(), s2(), s3(), v4); })
        ];
    }
    Main.react4 = react4;
    function act(s1, f) {
        f(s1());
        return react(s1, f);
    }
    Main.act = act;
    function act2(s1, s2, f) {
        f(s1(), s2());
        return react2(s1, s2, f);
    }
    Main.act2 = act2;
    function act3(s1, s2, s3, f) {
        f(s1(), s2(), s3());
        return react3(s1, s2, s3, f);
    }
    Main.act3 = act3;
    function lift(s1, f) {
        var t = sig(f(s1()));
        react(s1, function (v1) { t(f(v1)); });
        return t;
    }
    Main.lift = lift;
    function lifts(s1, f) {
        var t = sig(f(s1()));
        reacts(s1, function (v1) { t(f(v1)); });
        return t;
    }
    Main.lifts = lifts;
    function lift2(s1, s2, f) {
        var t = sig(f(s1(), s2()));
        react2(s1, s2, function (v1, v2) { t(f(v1, v2)); });
        return t;
    }
    Main.lift2 = lift2;
    function lift3(s1, s2, s3, f) {
        var t = sig(f(s1(), s2(), s3()));
        react3(s1, s2, s3, function (v1, v2, v3) { t(f(v1, v2, v3)); });
        return t;
    }
    Main.lift3 = lift3;
    function lift4(s1, s2, s3, s4, f) {
        var t = sig(f(s1(), s2(), s3(), s4()));
        react4(s1, s2, s3, s4, function (v1, v2, v3, v4) { t(f(v1, v2, v3, v4)); });
        return t;
    }
    Main.lift4 = lift4;
    function noop() { }
    function isNonEmpty(a) {
        return a.length !== 0;
    }
    function isEmpty(a) {
        return a.length === 0;
    }
    function isNull(a) {
        return a === null;
    }
    function isUndefined(a) {
        return a === void 0;
    }
    function timestampToAge(t) {
        return moment.unix(t).fromNow();
    }
    function formatTimestamp(t) {
        return moment.unix(t).format("MMM D YYYY h:mm:ss a");
    }
    //
    // Knockout Extensions
    //
    // ko.bindingHandlers['element'] = {
    //     update: (element, valueAccessor, allBindings, viewModel, bindingContext) => {
    //         const arg = ko.unwrap(valueAccessor())
    //         if (arg) {
    //             const $element = $(element)
    //             $element.empty()
    //             $element.append(arg)
    //         }
    //         return
    //     }
    // }
    ko.bindingHandlers['element'] = {
        init: function (element, valueAccessor, allBindings, viewModel, bindingContext) {
            valueAccessor()(element);
        }
    };
    ko.bindingHandlers['file'] = {
        init: function (element, valueAccessor, allBindings, viewModel, bindingContext) {
            var file = valueAccessor();
            if (file) {
                var $file_1 = $(element);
                $file_1.change(function () {
                    file($file_1[0].files[0]);
                });
            }
            return;
        }
    };
    ko.bindingHandlers['autoscroll'] = {
        init: function (element, valueAccessor, allBindings, viewModel, bindingContext) {
            // Bit of a hack. Attaches a method to the viewModel that scrolls the pane into view
            var $el = $(element);
            var $viewport = $el.closest('.workspace');
            viewModel.ensureVisible = function () {
                var p = $viewport.scrollLeft();
                var l = viewModel.left();
                var w = viewModel.width();
                var vw = $viewport.width();
                if (l + w > p + vw) {
                    $viewport.animate({ scrollLeft: l + w - vw }, 'fast');
                }
                else if (l < p) {
                    $viewport.animate({ scrollLeft: l }, 'fast');
                }
            };
            return;
        }
    };
    //
    // Models
    //
    var Model = (function () {
        function Model(id, cloud, algo, frame, responseColumn, createdAt) {
            this.id = id;
            this.cloud = cloud;
            this.algo = algo;
            this.frame = frame;
            this.responseColumn = responseColumn;
            this.createdAt = createdAt;
        }
        return Model;
    }());
    var Service = (function () {
        function Service(id, endpoint, createdAt) {
            this.id = id;
            this.endpoint = endpoint;
            this.createdAt = createdAt;
        }
        return Service;
    }());
    //
    // Components
    // 
    var NavButton = (function () {
        function NavButton(icon, title, isSelected, execute) {
            this.icon = icon;
            this.title = title;
            this.isSelected = isSelected;
            this.execute = execute;
        }
        return NavButton;
    }());
    var NavBar = (function () {
        function NavBar(buttons) {
            this.buttons = buttons;
        }
        return NavBar;
    }());
    function newNavBar(ctx) {
        var newNavButton = function (icon, title, isSelected, execute) {
            var button = new NavButton(icon, title, sig(isSelected), function () {
                for (var _i = 0, buttons_1 = buttons; _i < buttons_1.length; _i++) {
                    var b = buttons_1[_i];
                    b.isSelected(b === button);
                }
                execute();
            });
            return button;
        };
        var buttons = [
            newNavButton('ion-ios-cloud-outline', 'Clusters', true, ctx.showClouds),
            newNavButton('ion-ios-color-filter-outline', 'Models', false, ctx.showModels),
            newNavButton('ion-ios-world-outline', 'Services', false, ctx.showServices),
            newNavButton('ion-ios-paper-outline', 'Assets', false, ctx.showAssets)
        ];
        return new NavBar(buttons);
    }
    function templateOf(t) {
        return "tmpl-" + t.template;
    }
    function px(value) {
        return Math.round(value) + "px";
    }
    function newPanePosition(width) {
        if (width === void 0) { width = 200; }
        var l = sig(0);
        var w = sig(width);
        return {
            left: l,
            width: w,
            leftPx: lift(l, px),
            widthPx: lift(w, px),
            ensureVisible: noop
        };
    }
    function doAfterRender(elements) {
        $(elements).click(function () {
            var $this = $(this);
            $this.parent().children().removeClass('folder--selected');
            $this.addClass('folder--selected');
        });
    }
    //
    // Dialogs
    //
    var cloudIdPattern = /^[a-z0-9-]{1,16}$/i;
    var cloudMemoryPattern = /^[0-9]+[kmg]$/i;
    function newStartCloudDialog(ctx, go) {
        var error = sig('');
        var engineNames = sigs([]);
        var engineName = sig(void 0);
        var engineNameError = lift(engineName, function (engineName) {
            return engineName
                ? ''
                : "Select a H2O version";
        });
        var cloudId = sig('');
        var cloudIdError = lift(cloudId, function (cloudId) {
            return (cloudIdPattern.test(cloudId))
                ? ''
                : "Enter a valid cloud name";
        });
        var cloudSize = sig('1');
        var cloudSizeNum = lift(cloudSize, function (cloudSize) {
            return parseInt(cloudSize, 10);
        });
        var cloudSizeError = lift(cloudSizeNum, function (size) {
            return (!isNaN(size) && size > 0)
                ? ''
                : "Invalid cloud size";
        });
        var cloudMemory = sig('');
        var cloudMemoryError = lift(cloudMemory, function (cloudMemory) {
            return (cloudMemoryPattern.test(cloudMemory))
                ? ''
                : "Enter a valid Java memory specifier (e.g. 1024m, 2g, etc.)";
        });
        var canStartCloud = lift4(engineNameError, cloudIdError, cloudSizeError, cloudMemoryError, function (e1, e2, e3, e4) {
            return e1 === '' && e2 === '' && e3 === '' && e4 === '';
        });
        var startCloud = function () {
            if (!canStartCloud()) {
                return;
            }
            ctx.setBusy('Creating cloud...');
            ctx.remote.startCloud(cloudId(), engineName(), cloudSizeNum(), cloudMemory(), ctx.principal.username, function (err, cloud) {
                if (err) {
                    error(err.message);
                }
                else {
                    go({ cloud: cloud });
                }
                ctx.setFree();
            });
        };
        var cancel = function () {
            go(null);
        };
        ctx.remote.getEngines(function (err, engines) {
            if (err) {
                return;
            }
            engineNames(_.map(engines, function (engine) { return engine.name; }));
        });
        return {
            title: 'Start a new cloud',
            engineNames: engineNames,
            engineName: engineName,
            engineNameError: engineNameError,
            cloudId: cloudId,
            cloudIdError: cloudIdError,
            cloudSize: cloudSize,
            cloudSizeError: cloudSizeError,
            cloudMemory: cloudMemory,
            cloudMemoryError: cloudMemoryError,
            canStartCloud: canStartCloud,
            startCloud: startCloud,
            error: error,
            cancel: cancel,
            dispose: noop,
            template: 'start-cloud-dialog'
        };
    }
    function newBuildModelDialog(ctx, cloudId, go) {
        var error = sig('');
        var frame = sig('');
        var frameError = lift(frame, function (f) {
            return (f && f.trim().length > 0)
                ? ''
                : 'Enter a valid dataset path';
        });
        var responseColumn = sig('');
        var responseColumnError = lift(responseColumn, function (c) {
            return (c && c.trim().length > 0)
                ? ''
                : 'Enter a valid column name';
        });
        var maxRunTime = sig('1000');
        var maxRunTimeNum = lift(maxRunTime, function (t) { return parseInt(t, 10); });
        var maxRunTimeError = lift(maxRunTimeNum, function (t) {
            return (!isNaN(t) && t > 0)
                ? ''
                : 'Invalid run time';
        });
        var canBuildModel = lift3(frameError, responseColumnError, maxRunTimeError, function (e1, e2, e3) {
            return e1 === '' && e2 === '' && e3 === '';
        });
        function buildModel() {
            ctx.setBusy('Building model...');
            ctx.remote.buildModel(cloudId, frame(), responseColumn(), maxRunTimeNum(), function (err) {
                if (err) {
                    error(err.message);
                }
                else {
                    go({ success: true });
                }
                ctx.setFree();
            });
        }
        var cancel = function () {
            go(null);
        };
        return {
            title: "Build a Model",
            frame: frame,
            frameError: frameError,
            responseColumn: responseColumn,
            responseColumnError: responseColumnError,
            maxRunTime: maxRunTime,
            maxRunTimeError: maxRunTimeError,
            canBuildModel: canBuildModel,
            buildModel: buildModel,
            error: error,
            cancel: cancel,
            dispose: noop,
            template: 'build-model-dialog'
        };
    }
    function newDeployModelDialog(ctx, modelId, go) {
        var error = sig('');
        var port = sig('8000');
        var portNum = lift(port, function (port) {
            return parseInt(port, 10);
        });
        var portError = lift(portNum, function (size) {
            return (!isNaN(size) && size > 0)
                ? ''
                : 'Invalid port number';
        });
        var canDeployModel = lift(portError, function (e) {
            return e === '';
        });
        var deployModel = function () {
            ctx.setBusy('Deploying model...');
            ctx.remote.startScoringService(modelId, portNum(), function (err) {
                if (err) {
                    error(err.message);
                }
                else {
                    go({ success: true });
                }
                ctx.setFree();
            });
        };
        var cancel = function () {
            go(null);
        };
        return {
            title: "Deploy Model " + modelId,
            port: port,
            portError: portError,
            canDeployModel: canDeployModel,
            deployModel: deployModel,
            error: error,
            cancel: cancel,
            dispose: noop,
            template: 'deploy-model-dialog'
        };
    }
    function newAddEngineDialog(ctx, go) {
        var error = sig('');
        var form = sig(null);
        var file = sig(null);
        var addEngine = function () {
            var f = file();
            if (!(f && f.name)) {
                return;
            }
            ctx.setBusy('Uploading asset...');
            var formData = new FormData(form());
            ctx.remote.upload(formData, function (err, data) {
                ctx.setFree();
                if (err) {
                    error(err.message);
                    return;
                }
                go({ success: true });
            });
        };
        var cancel = function () {
            go(null);
        };
        return {
            title: "Add Engine",
            form: form,
            file: file,
            addEngine: addEngine,
            error: error,
            cancel: cancel,
            dispose: noop,
            template: 'add-engine-dialog'
        };
    }
    //
    // Panes
    //
    function newCloudsPane(ctx) {
        var error = sig('');
        var items = sigs([]);
        var hasItems = lifts(items, function (items) { return items.length > 0; });
        var startCloud = function () {
            var dialog = newStartCloudDialog(ctx, function (result) {
                ctx.popDialog();
                if (result) {
                    ctx.showClouds();
                }
            });
            ctx.pushDialog(dialog);
        };
        ctx.remote.getClouds(function (err, clouds) {
            if (err) {
                error(err.message);
                return;
            }
            items(_.map(clouds, function (cloud) {
                return {
                    title: cloud.name,
                    subhead: 'State:',
                    slug: String(cloud.state),
                    execute: function () { ctx.showCloud(cloud); },
                    template: 'folder'
                };
            }));
        });
        return {
            title: 'Clusters',
            error: error,
            hasItems: hasItems,
            items: items,
            startCloud: startCloud,
            template: 'clouds',
            dispose: noop,
            position: newPanePosition()
        };
    }
    function newCloudPane(ctx, cloud) {
        var error = sig('');
        var items = [
            {
                title: 'Cluster Details',
                subhead: 'Size:',
                slug: String(cloud.size),
                execute: function () { ctx.showCloudDetails(cloud); },
                template: 'folder'
            }
        ];
        if (cloud.state !== 'Stopped') {
            ctx.remote.getCloudModels(cloud.name, function (error, models) {
                models.
                ;
            });
            items.push({
                title: 'Models',
                subhead: 'Models in cluster',
                slug: '',
                execute: function () { ctx.showCloudModels(cloud); },
                template: 'folder'
            });
            items.push({
                title: 'Jobs',
                subhead: 'Cluster Jobs',
                slug: '',
                execute: function () { ctx.showCloudJobs(cloud); },
                template: 'folder'
            });
        }
        return {
            title: cloud.name,
            error: error,
            items: items,
            template: 'cloud',
            dispose: noop,
            position: newPanePosition()
        };
    }
    function newCloudDetailsPane(ctx, cloud) {
        var error = sig('');
        var state = sig(cloud.state);
        var cloudDetails = sig(null);
        function stopCloud() {
            ctx.setBusy('Stopping cluster...');
            ctx.remote.stopCloud(cloud.name, function (err) {
                ctx.setFree();
                if (err) {
                    alert(err.message);
                    return;
                }
                ctx.showClouds();
            });
        }
        if (cloud.state != 'Stopped') {
            ctx.remote.getCloudStatus(cloud.name, function (err, h2oCloud) {
                if (err) {
                    state('Unknown');
                    error(err.message);
                    return;
                }
                var cloudDetail = {
                    engineVersion: h2oCloud.engine_version,
                    totalMemory: h2oCloud.memory,
                    totalCores: String(h2oCloud.total_cores),
                    allowedCores: String(h2oCloud.allowed_cores)
                };
                cloudDetails(cloudDetail);
                state(h2oCloud.state);
            });
        }
        return {
            title: 'Cluster Details',
            engineName: cloud.engine_name,
            size: String(cloud.size),
            memory: cloud.memory,
            applicationId: cloud.application_id,
            address: "http://" + cloud.address + "/",
            username: cloud.username,
            state: state,
            createdAt: timestampToAge(cloud.created_at),
            stopCloud: stopCloud,
            cloudDetails: cloudDetails,
            template: 'cloudInfo',
            error: error,
            dispose: noop,
            position: newPanePosition(650)
        };
    }
    function newCloudJobsPane(ctx, cloud) {
        var error = sig('');
        var items = sigs([]);
        var hasItems = lifts(items, function (items) { return items.length > 0; });
        ctx.remote.getJobs(cloud.name, function (err, jobs) {
            if (err) {
                error(err.message);
                return;
            }
            items(_.map(jobs, function (job) {
                return {
                    title: job.name,
                    subhead: "Status",
                    slug: job.progress,
                    execute: function () { ctx.showCloudJob(job); },
                    template: 'folder'
                };
            }));
        });
        return {
            title: 'Cluster Jobs',
            error: error,
            items: items,
            hasItems: hasItems,
            template: 'cloudJobs',
            dispose: noop,
            position: newPanePosition(),
        };
    }
    function newCloudModelsPane(ctx, cloud) {
        var error = sig('');
        var items = sigs([]);
        var hasItems = lifts(items, function (items) { return items.length > 0; });
        function buildModel() {
            var dialog = newBuildModelDialog(ctx, cloud.name, function (result) {
                ctx.popDialog();
                if (result) {
                    ctx.showModels();
                }
            });
            ctx.pushDialog(dialog);
        }
        ctx.remote.getCloudModels(cloud.name, function (err, models) {
            if (err) {
                error(err.message);
                return;
            }
            items(_.map(models, function (model) {
                return {
                    title: model.name,
                    subhead: model.dataset,
                    slug: model.target_name,
                    execute: function () { ctx.showCloudModel(model); },
                    template: 'folder'
                };
            }));
        });
        return {
            title: 'Cluster Models',
            error: error,
            items: items,
            hasItems: hasItems,
            template: 'cloudModels',
            buildModel: buildModel,
            dispose: noop,
            position: newPanePosition(),
        };
    }
    function newCloudModelPane(ctx, model) {
        var getModel = function () {
            ctx.setBusy('Getting model from h2o...');
            ctx.remote.getModelFromCloud(model.cloud_name, model.name, function (err, model) {
                ctx.setFree();
                if (err) {
                    alert(err.message);
                    return;
                }
                ctx.showModels();
            });
        };
        return {
            title: model.name,
            cloud: model.cloud_name,
            algo: model.algo,
            frame: model.dataset,
            responseColumn: model.target_name,
            maxRunTime: String(model.max_runtime),
            javaModelPath: model.java_model_path,
            createdAt: String(model.created_at),
            getModel: getModel,
            template: 'cloudModel',
            dispose: noop,
            position: newPanePosition(650)
        };
    }
    function newModelsPane(ctx) {
        var error = sig('');
        var items = sigs([]);
        var hasItems = lifts(items, function (items) { return items.length > 0; });
        ctx.remote.getModels(function (err, models) {
            if (err) {
                error(err.message);
                return;
            }
            items(_.map(models, function (model) {
                return {
                    title: model.name,
                    subhead: model.dataset,
                    slug: model.target_name,
                    execute: function () { ctx.showModel(model); },
                    template: 'folder'
                };
            }));
        });
        return {
            title: 'Models',
            error: error,
            items: items,
            hasItems: hasItems,
            template: 'models',
            dispose: noop,
            position: newPanePosition(),
        };
    }
    function newModelPane(ctx, model) {
        var deployModel = function () {
            var dialog = newDeployModelDialog(ctx, model.name, function (result) {
                ctx.popDialog();
                if (result) {
                    ctx.showServices();
                }
            });
            ctx.pushDialog(dialog);
        };
        var deleteModel = function () {
            ctx.setBusy('Deleting model...');
            ctx.remote.deleteModel(model.name, function (err) {
                ctx.setFree();
                if (err) {
                    alert(err.message); // FIXME
                    return;
                }
                ctx.showModels();
            });
        };
        return {
            title: model.name,
            cloud: model.cloud_name,
            algo: model.algo,
            frame: model.dataset,
            responseColumn: model.target_name,
            maxRunTime: String(model.max_runtime),
            javaModelPath: model.java_model_path,
            createdAt: timestampToAge(model.created_at),
            deployModel: deployModel,
            deleteModel: deleteModel,
            template: 'model',
            dispose: noop,
            position: newPanePosition(650)
        };
    }
    function newServicesPane(ctx) {
        var error = sig('');
        var items = sigs([]);
        var hasItems = lifts(items, function (items) { return items.length > 0; });
        ctx.remote.getScoringServices(function (err, services) {
            if (err) {
                error(err.message);
                return;
            }
            items(_.map(services, function (service) {
                return {
                    title: service.model_name,
                    subhead: 'State:',
                    slug: service.state,
                    execute: function () { ctx.showService(service); },
                    template: 'folder'
                };
            }));
        });
        return {
            title: 'Services',
            error: error,
            hasItems: hasItems,
            items: items,
            template: 'services',
            dispose: noop,
            position: newPanePosition(),
        };
    }
    function newServicePane(ctx, service) {
        var stopService = function () {
            ctx.setBusy('Stopping service...');
            ctx.remote.stopScoringService(service.model_name, service.port, function (err) {
                ctx.setFree();
                if (err) {
                    alert(err.message);
                    return;
                }
                ctx.showServices();
            });
        };
        return {
            title: service.model_name,
            state: service.state,
            address: service.address,
            port: String(service.port),
            url: "http://" + service.address + ":" + service.port + "/",
            pid: String(service.pid),
            createdAt: timestampToAge(service.created_at),
            stopService: stopService,
            template: 'service',
            dispose: noop,
            position: newPanePosition(650),
        };
    }
    function newAssetsPane(ctx) {
        var items = [
            {
                title: 'Engines',
                subhead: 'View deployed engines',
                slug: '',
                execute: function () { ctx.showEngines(); },
                template: 'folder'
            }
        ];
        return {
            title: 'Assets',
            template: 'assets',
            dispose: noop,
            position: newPanePosition(),
            items: items
        };
    }
    function newEnginesPane(ctx) {
        var error = sig('');
        var items = sigs([]);
        var hasItems = lifts(items, function (items) { return items.length > 0; });
        var addEngine = function () {
            var dialog = newAddEngineDialog(ctx, function (result) {
                ctx.popDialog();
                if (result) {
                    if (result.success) {
                        ctx.showEngines();
                    }
                }
            });
            ctx.pushDialog(dialog);
        };
        ctx.remote.getEngines(function (err, engines) {
            if (err) {
                error(err.message);
                return;
            }
            items(_.map(engines, function (engine) {
                return {
                    title: engine.name,
                    subhead: timestampToAge(engine.created_at),
                    slug: '',
                    execute: function () { ctx.showEngine(engine); },
                    template: 'folder'
                };
            }));
        });
        return {
            title: 'Engines',
            error: error,
            items: items,
            hasItems: hasItems,
            addEngine: addEngine,
            template: 'engines',
            dispose: noop,
            position: newPanePosition()
        };
    }
    function newEnginePane(ctx, engine) {
        var deleteEngine = function () {
            ctx.setBusy('Deleting engine...');
            ctx.remote.deleteEngine(engine.name, function (err) {
                ctx.setFree();
                if (err) {
                    alert(err.message); // FIXME
                    return;
                }
                ctx.showEngines();
            });
        };
        return {
            title: engine.name,
            path: engine.path,
            createdAt: timestampToAge(engine.created_at),
            deleteEngine: deleteEngine,
            template: 'engine',
            dispose: noop,
            position: newPanePosition(650),
        };
    }
    var Context = (function () {
        function Context() {
            this.remote = Proxy;
            this.principal = { username: 'unknown' };
            this.setBusy = uni1();
            this.setFree = uni();
            this.pushDialog = uni1();
            this.popDialog = uni();
            this.showPane = uni2();
            this.showClouds = uni();
            this.showCloud = uni1();
            this.showCloudDetails = uni1();
            this.showCloudJobs = uni1();
            this.showCloudModels = uni1();
            this.showCloudJob = uni1();
            this.showCloudModel = uni1();
            this.showModels = uni();
            this.showModel = uni1();
            this.showServices = uni();
            this.showService = uni1();
            this.showAssets = uni();
            this.showEngines = uni();
            this.showEngine = uni1();
        }
        return Context;
    }());
    Main.Context = Context;
    var Breadcrumb = (function () {
        function Breadcrumb(title, execute) {
            this.title = title;
            this.execute = execute;
        }
        return Breadcrumb;
    }());
    function newApp() {
        var ctx = new Context();
        var navBar = newNavBar(ctx);
        var breadcrumbs = sigs([]);
        var panes = sigs([]);
        var span = sig(0);
        var spanPx = lift(span, px);
        var dialogs = sigs([]);
        var hasDialogs = lifts(dialogs, isNonEmpty);
        var busyMessage = sig(void 0);
        ctx.pushDialog.on(function (dialog) {
            dialogs.push(dialog);
        });
        ctx.popDialog.on(function () {
            dialogs.pop();
        });
        ctx.setBusy.on(function (message) {
            busyMessage(message);
        });
        ctx.setFree.on(function () {
            busyMessage(void 0);
        });
        ctx.showPane.on(function (index, pane) {
            var disposables = panes.splice(index, panes().length - index, pane);
            for (var _i = 0, disposables_1 = disposables; _i < disposables_1.length; _i++) {
                var disposable = disposables_1[_i];
                disposable.dispose();
            }
            var left = 0;
            for (var _a = 0, _b = panes(); _a < _b.length; _a++) {
                var p = _b[_a];
                var pos = p.position;
                pos.left(left);
                left += pos.width();
            }
            // Set span to max total width so that browsing panes leftward
            //  does not result in a jerky scroll to the right.
            if (span() < left) {
                span(left);
            }
            pane.position.ensureVisible();
            breadcrumbs(_.map(panes(), function (pane) {
                return new Breadcrumb(pane.title, function () {
                    pane.position.ensureVisible();
                });
            }));
        });
        ctx.showClouds.on(function () {
            ctx.showPane(0, newCloudsPane(ctx));
        });
        ctx.showCloud.on(function (cloud) {
            ctx.showPane(1, newCloudPane(ctx, cloud));
        });
        ctx.showCloudDetails.on(function (cloud) {
            ctx.showPane(2, newCloudDetailsPane(ctx, cloud));
        });
        ctx.showCloudModels.on(function (cloud) {
            ctx.showPane(2, newCloudModelsPane(ctx, cloud));
        });
        ctx.showCloudJobs.on(function (cloud) {
            ctx.showPane(2, newCloudJobsPane(ctx, cloud));
        });
        ctx.showCloudModel.on(function (model) {
            ctx.showPane(3, newCloudModelPane(ctx, model));
        });
        ctx.showCloudJob.on(function (job) {
            ctx.showPane(3, newCloudJobPane(ctx, job));
        });
        ctx.showModels.on(function () {
            ctx.showPane(0, newModelsPane(ctx));
        });
        ctx.showModel.on(function (model) {
            ctx.showPane(1, newModelPane(ctx, model));
        });
        ctx.showServices.on(function () {
            ctx.showPane(0, newServicesPane(ctx));
        });
        ctx.showService.on(function (service) {
            ctx.showPane(1, newServicePane(ctx, service));
        });
        ctx.showAssets.on(function () {
            ctx.showPane(0, newAssetsPane(ctx));
        });
        ctx.showEngines.on(function () {
            ctx.showPane(1, newEnginesPane(ctx));
        });
        ctx.showEngine.on(function (engine) {
            ctx.showPane(2, newEnginePane(ctx, engine));
        });
        ctx.showClouds();
        return {
            context: ctx,
            navBar: navBar,
            breadcrumbs: breadcrumbs,
            panes: panes,
            span: spanPx,
            hasDialogs: hasDialogs,
            dialogs: dialogs,
            busyMessage: busyMessage,
            templateOf: templateOf,
            afterRender: doAfterRender
        };
    }
    Main.newApp = newApp;
})(Main || (Main = {}));
/// <reference path="references.ts" />
/// <reference path="app.ts" />
"use strict";
var Main;
(function (Main) {
    function boot() {
        Main.app = Main.newApp();
        ko.applyBindings(Main.app);
    }
    (document.readyState !== 'loading')
        ? boot()
        : document.addEventListener('DOMContentLoaded', boot);
})(Main || (Main = {}));
