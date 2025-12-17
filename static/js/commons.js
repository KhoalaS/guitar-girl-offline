Function.prototype.bind = function() {
	// http://developer.mozilla.org/en/docs/Core_JavaScript_1.5_Reference:Functions:arguments
	var _$A = function(a) {
		return Array.prototype.slice.call(a);
	}
	if (arguments.length < 2 && (typeof arguments[0] == "undefined")) return this;
	var __method = this,
		args = _$A(arguments),
		object = args.shift();
	return function() {
		return __method.apply(object, args.concat(_$A(arguments)));
	};
};

jQuery.fn.extend({
	checkForEnter: function(callback) {

		function _func(e) {
			if (e.keyCode == 13) {
				e.preventDefault();
				e.stopPropagation();
				callback();
				return false;
			}
			return true;
		};

		if (jQuery.browser.mozilla) jQuery(this).keypress(_func);
		else jQuery(this).keydown(_func);
		//jQuery(this).keydown(_func);

		return this;
	},

	eachClass: function(callback) {
		jQuery(this).each(function() {
			var _class = jQuery(this).attr("class");
			if (typeof _class != "undefined") jQuery.each(_class.split(" "), callback);
		});

		return this;
	}
});

neon = {
	create: function() {
		return function() {
			this.initialize.apply(this, arguments);
		};
	},
	overed: function(destination, source) {
		for (var property in source) {
			destination[property] = source[property];
		}
		return destination;
	},
	extend: function(ext, sup) {
		var ext_prototype = ext.prototype;
		var sup_prototype = sup.prototype;
		for (var property in sup_prototype) {
			if (!ext_prototype[property]) ext_prototype[property] = sup_prototype[property];
		}
	},
	clone: function(obj) {
		var newObj = {};
		for (var property in obj) newObj[property] = obj[property];
		return newObj;
	},
	cookie: {
		set: function(name, value) {
			var argc = arguments.length;
			var argv = arguments;
			var expires = (argc > 2) ? argv[2] : null;
			var path = (argc > 3) ? argv[3] : null;
			var domain = (argc > 4) ? argv[4] : null;
			var secure = (argc > 5) ? argv[5] : false;
			var encode = (argc > 6) ? argv[6] : true;

			document.cookie = name + "=" + (encode ? escape(value) : value) + ((expires == null) ? "" : ("; expires =" + expires.toGMTString())) + ((path == null) ? "" : ("; path =" + path)) + ((domain == null) ? "" : ("; domain =" + domain)) + ((secure == true) ? "; secure" : "");
		},
		get: function(name) {
			var dcookie = document.cookie;
			var cname = name + "=";
			var clen = dcookie.length;
			var cbegin = 0;
			while (cbegin < clen) {
				var vbegin = cbegin + cname.length;
				if (dcookie.substring(cbegin, vbegin) == cname) {
					var vend = dcookie.indexOf(";", vbegin);
					if (vend == -1) vend = clen;
					return unescape(dcookie.substring(vbegin, vend));
				}
				cbegin = dcookie.indexOf(" ", cbegin) + 1;
				if (cbegin == 0) break;
			}
			return "";
		}
	},
	utils: {
		trim: function(str) {
			return str.replace(/(^\s*)|(\s*$)/g, "");
		},
		num: function(str) {
			str = neon.utils.trim(str);
			return (/^[0-9]+$/).test(str);
		},
		numToStr: function(num) {
			if (num.length < 4) return num;

			while (num.indexOf(",") >= 0) {
				num = num.replace(",", "");
			}

			return num;
		},
		imgError: function(obj, defaultSize) {
			if (!obj) return;

			obj.onerror = null;
			obj.src = neon.url.defaultThumbnail(obj.src);
		},
		setBannerError: function(selector) {
			return $(selector).one("error", function() {
				$(this).attr("src", "https://file.pmangplus.com/bugs/images/common/noalbumD.gif");
			});
		},
		prettyJson: function(jsonStr) {
			var indent = "  ";
			var result = "";

			var indentDepth = 0;
			var targetStr = "";

			for (var i = 0; i < jsonStr.length; i++) {
				targetStr = jsonStr.substring(i, i + 1);
				if (targetStr == "{" || targetStr == "[") {
					result += targetStr + "\n";
					indentDepth++;
					for (var j = 0; j < indentDepth; j++) result += indent;
				} else if (targetStr == "}" || targetStr == "]") {
					result += "\n";
					indentDepth--;
					for (var j = 0; j < indentDepth; j++) result += indent;
					result += targetStr;
				} else if (targetStr == ",") {
					result += targetStr + "\n";
					for (var j = 0; j < indentDepth; j++) result += indent;
				} else result += targetStr;
			}
			return result;
		},

		addCommas: function(str) {
			str += "";
			var rgx = /(\d+)(\d{3})/;
			while (rgx.test(str)) {
				str = str.replace(rgx, "$1" + "," + "$2");
			}
			return str;
		},

		dateFormat: function(dt, fmt) {
			return fmt.replace(/(yyyy|yy|mm|dd|HH24|hh24|hh|MI|mi|SS|ss|am|pm)/gi,

			function($1) {
				switch ($1) {
					case 'yyyy':
						return dt.getFullYear();
					case 'yy':
						return dt.getFullYear().toString().substr(2);
					case 'mm':
						return (dt.getMonth() + 1) < 10 ? "0" + (dt.getMonth() + 1) : (dt.getMonth() + 1);
					case 'dd':
						return dt.getDate() < 10 ? "0" + dt.getDate() : dt.getDate();
					case 'HH24':
						return dt.getHours() < 10 ? "0" + dt.getHours() : dt.getHours();
					case 'hh24':
						return dt.getHours();
					case 'hh':
						return (h = dt.getHours() % 12) ? h : 12;
					case 'MI':
						return dt.getMinutes() < 10 ? "0" + dt.getMinutes() : dt.getMinutes();
					case 'mi':
						return dt.getMinutes();
					case 'SS':
						return dt.getSeconds() < 10 ? "0" + dt.getSeconds() : dt.getSeconds();
					case 'ss':
						return dt.getSeconds();
					case 'am':
						return dt.getHours() < 12 ? 'am' : 'pm';
					case 'pm':
						return dt.getHours() < 12 ? 'am' : 'pm';
				}
			});
		},

		dateFormatYmd: function(dt) {
			return this.dateFormat(dt, 'yyyy') + "년" + this.dateFormat(dt, 'mm') + "월" + this.dateFormat(dt, 'dd') + "일";
		},

		dateFormatDot: function(dt) {
			return this.dateFormat(dt, 'yyyy') + "." + this.dateFormat(dt, 'mm') + "." + this.dateFormat(dt, 'dd');
		}
	},
	valid: {
		empty: function(str) {
			if (typeof str == "string" && neon.utils.trim(str).length > 0) return false;
			return true;
		},
		yyyyMMdd: function(str) {
			return /^\d{4}\-\d{2}\-\d{2}$/.test(str);
		}
	},
	url: {
		defaultThumbnail: function(src, defaultSize) {
			var tmp = src.split("/");
			var type = tmp[tmp.length - 5];

			var width = defaultSize ? defaultSize : tmp[tmp.length - 3];
			var height = width;
			if (type == "mv" && width == "100") height = "75";
			else if (type == "wesalbum" || type == "wesalbum") {
				type = "album";
				if (width == "original") width = height = 75;
			}

			return "https://file.pmangplus.com/nbugs/common/i_" + type + width + "x" + height + ".gif";
		}
	}
};

var debug = {
	dump: function() {
		var retVal = '',
			argVal = arguments[0] || '';

		switch ((typeof argVal).toLowerCase()) {
			case 'object':
				$.each(argVal, function(key, value) {
					retVal += '[' + (typeof value) + '] ' + key + ' => ' + value + '\n';
				});
				break;

			case 'string':
			default:
				retVal += (typeof argVal) + '/' + argVal;
		}
		return retVal;
	}
};