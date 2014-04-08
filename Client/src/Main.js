/**
* Created by RockyF on 14-3-31.
*/
/// <reference path="jquery.d.ts" />
/// <reference path="Client.ts" />
/// <reference path="Panels.ts" />
/// <reference path="Utils.ts" />
$(function () {
    muiltcube.Main.getInstance().start();
});

var wsuri = "ws://127.0.0.1:8888";
var muiltcube;
(function (muiltcube) {
    var Main = (function () {
        function Main() {
            var _this = this;
            this.onLoginSuccess = function () {
                new muiltcube.RoleCreatePanel(_this.onLoginSuccess);
            };
        }
        Main.getInstance = function () {
            if (!this._instance) {
                this._instance = new Main();
            }

            return this._instance;
        };

        Main.prototype.start = function () {
            this.client = muiltcube.Client.getInstance();
            this.client.init(wsuri);
            this.client.start();

            this.rpc = muiltcube.RPC.getInstance();
            this.rpc.init("http://localhost/muiltcube/WebServer/gateway.php");

            this.showLoginPanel();
        };

        Main.prototype.showLoginPanel = function () {
            new muiltcube.LoginPanel(this.onLoginSuccess);
        };
        return Main;
    })();
    muiltcube.Main = Main;
})(muiltcube || (muiltcube = {}));
