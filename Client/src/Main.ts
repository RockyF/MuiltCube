/**
 * Created by RockyF on 14-3-31.
 */
/// <reference path="jquery.d.ts" />

var wsuri = "ws://127.0.0.1:8888";
var client = com.tbs.muiltcube.Client.getInstance();
var divLog;

window.onload = function () {
	divLog = $('#log');

	var client = com.tbs.muiltcube.Client.getInstance();
	client.init('ws://localhost:8888');
	client.start();

	client.addCmdListener(1001, on1001Response);
};

function send() {
	var message = $('#message');

	client.sendData(message.val());

	divLog.scrollTop = divLog.scrollHeight;
}

function on1001Response(body){
	client.removeCmdListener(1001, on1001Response);

	console.log("message received: " + body);

	divLog.append("<p style='color:blue'>server say: " + body.result + "</p>");

	divLog.scrollTop = divLog.scrollHeight;
}