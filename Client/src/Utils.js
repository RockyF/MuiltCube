/**
* Created by lenovo on 14-3-31.
*/
var Utils = (function () {
    function Utils() {
    }
    Utils.stringFormat = function (format) {
        var args = [];
        for (var _i = 0; _i < (arguments.length - 1); _i++) {
            args[_i] = arguments[_i + 1];
        }
        var result = format;
        for (var i = 0, len = args.length; i < len; i++) {
            result = result.replace('{' + (i) + '}', args[i]);
        }

        return result;
    };
    return Utils;
})();
