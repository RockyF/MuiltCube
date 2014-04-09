/**
* Created by lenovo on 14-4-4.
*/
/// <reference path="async.d.ts" />
/// <reference path="jquery.d.ts" />
/// <reference path="Client.ts" />
/// <reference path="Utils.ts" />
var muiltcube;
(function (muiltcube) {
    var LoginPanel = (function () {
        function LoginPanel(onSuccess) {
            var _this = this;
            this.onTplLoaded = function (tpl) {
                $("body").append(tpl);

                $("#btnLogin").click(_this.onBtnLoginClicked);
            };
            this.onBtnLoginClicked = function () {
                _this.client.start(_this.onOpen);
            };
            this.onOpen = function () {
                _this.client.addCmdListener(1001, _this.on1001Response);

                var id = parseInt($("#tiID").val());
                var pwd = $("#tiPwd").val();

                if (id == 0 || pwd == "") {
                    return;
                }

                _this.client.send(1001, { id: id, pwd: pwd });
            };
            this.on1001Response = function (msg) {
                _this.client.removeCmdListener(1001, _this.on1001Response);
                if (msg.result == 0) {
                    //alert("onLoginSuccess");
                    $("#loginPanel").remove();
                    _this.onSuccess();
                } else {
                    alert("onLoginFailed");
                }
            };
            this.onSuccess = onSuccess;
            this.client = muiltcube.Client.getInstance();

            Utils.loadTPL("tpl/login.html", this.onTplLoaded);
        }
        return LoginPanel;
    })();
    muiltcube.LoginPanel = LoginPanel;
    var RoleCreatePanel = (function () {
        function RoleCreatePanel(onSuccess) {
            var _this = this;
            this.one = function (callback) {
                Utils.loadTPL("tpl/role_create.html", function (tpl) {
                    $("body").append(tpl);
                    callback();
                });
            };
            this.two = function (callback) {
                _this.rpc.execute("Common", "getColorList", { id: 1001 }, function (msg) {
                    if (msg.result == 0) {
                        this.colors = msg.body;
                    }
                    var colorSelectBox = $("#colorSelectBox");
                    var liTpl = $("<li class='colorItem'></li>");
                    for (var i = 0; i < this.colors.length; i++) {
                        var o = this.colors[i];
                        var li = liTpl.clone();
                        li.css("background-color", "#" + o.value);
                        colorSelectBox.append(li);
                    }
                    callback();
                });
            };
            this.onSuccess = onSuccess;

            this.rpc = muiltcube.RPC.getInstance();

            async.parallel({
                one: this.one,
                two: this.two
            }, function (err, results) {
            });
        }
        return RoleCreatePanel;
    })();
    muiltcube.RoleCreatePanel = RoleCreatePanel;
})(muiltcube || (muiltcube = {}));
