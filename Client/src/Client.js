var com;
(function (com) {
    (function (tbs) {
        /**
        * Created by lenovo on 14-3-31.
        */
        (function (muiltcube) {
            var Client = (function () {
                function Client() {
                    var _this = this;
                    this.onOpen = function (event) {
                        console.log(Utils.stringFormat('connect to {0} success!', _this.wsServer));
                    };
                    this.onClose = function (event) {
                        console.log('connect closed!');
                    };
                    this.onMessage = function (event) {
                        console.log(Utils.stringFormat('get data: {0}', event.data));
                        var msg = JSON.parse(event.data);
                        var cmd = msg.cmd;

                        var callbacks = _this.callbackMap[cmd];
                        if (callbacks) {
                            for (var i = 0, len = callbacks.length; i < len; i++) {
                                callbacks[i].call(_this, msg.body);
                            }
                        }
                    };
                    this.onError = function (event) {
                        console.log('socket error!');
                    };
                }
                Client.getInstance = function () {
                    if (!this._instance) {
                        this._instance = new Client();
                    }

                    return this._instance;
                };

                Client.prototype.init = function (wsServer) {
                    this.wsServer = wsServer;
                    this.callbackMap = [];
                };

                Client.prototype.start = function () {
                    this.websocket = new WebSocket(this.wsServer);

                    this.websocket.onopen = this.onOpen;
                    this.websocket.onclose = this.onClose;
                    this.websocket.onmessage = this.onMessage;
                    this.websocket.onerror = this.onError;
                };

                Client.prototype.send = function (cmd, body) {
                    var msg = new Object();
                    msg.cmd = cmd;
                    msg.body = body;
                    this.websocket.send(JSON.stringify(msg));
                };

                Client.prototype.sendData = function (content) {
                    this.websocket.send(content);
                };

                Client.prototype.addCmdListener = function (cmd, callback) {
                    if (!this.callbackMap[cmd]) {
                        this.callbackMap[cmd] = [];
                    }

                    this.callbackMap[cmd].push(callback);
                };

                Client.prototype.removeCmdListener = function (cmd, callback) {
                    var callbacks = this.callbackMap[cmd];
                    if (callbacks) {
                        for (var i = 0, len = callbacks.length; i < len; i++) {
                            if (callback = callbacks[i]) {
                                callbacks.splice(i, 1);
                            }
                        }
                    }
                };

                Client.prototype.getState = function () {
                    return this.websocket.readyState;
                };
                return Client;
            })();
            muiltcube.Client = Client;
        })(tbs.muiltcube || (tbs.muiltcube = {}));
        var muiltcube = tbs.muiltcube;
    })(com.tbs || (com.tbs = {}));
    var tbs = com.tbs;
})(com || (com = {}));
