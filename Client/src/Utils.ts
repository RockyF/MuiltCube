/**
 * Created by lenovo on 14-3-31.
 */
/// <reference path="jquery.d.ts" />
class Utils{
	public static stringFormat(format:string, ...args){
		var result:string = format;
		for(var i = 0, len = args.length; i < len; i++){
			result = result.replace('{' + (i) + '}', args[i]);
		}

		return result;
	}

	public static loadTPL(url:string, callback:any):void{
		$.ajax({
			url: url,
			method: "GET",
			success: function(data){
				callback($(data));
			}
		});
	}
}