/**
* Created by RockyF on 14-3-31.
*/
/// <reference path="jquery.d.ts" />
$(function () {
    com.tbs.muiltcube.Main.getInstance().start();
});

var wsuri = "ws://127.0.0.1:8888";
var com;
(function (com) {
    (function (tbs) {
        (function (muiltcube) {
            var Main = (function () {
                function Main() {
                    this.showLoginPanle = function () {
                        Utils.loadTPL("tpl/login.html", function (tpl) {
                            $("body").append(tpl);
                            //this.client.removeCmdListener(1001, on1001Response);
                        });
                    };
                }
                Main.getInstance = function () {
                    if (!this._instance) {
                        this._instance = new Main();
                    }

                    return this._instance;
                };

                Main.prototype.start = function () {
                    this.client = com.tbs.muiltcube.Client.getInstance();
                    this.client.init(wsuri);
                    this.client.start();

                    this.showLoginPanle();
                };
                return Main;
            })();
            muiltcube.Main = Main;
        })(tbs.muiltcube || (tbs.muiltcube = {}));
        var muiltcube = tbs.muiltcube;
    })(com.tbs || (com.tbs = {}));
    var tbs = com.tbs;
})(com || (com = {}));

function send() {
    var message = $('#message');
    //client.sendData(message.val());
}

function on1001Response(body) {
    //client.removeCmdListener(1001, on1001Response);
    console.log("message received: " + body);
}
