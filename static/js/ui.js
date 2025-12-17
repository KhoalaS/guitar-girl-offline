if (typeof(pp) == "undefined") pp = {};

pp.ui = {};
pp.ui.Select = neon.create();
pp.ui.Select.prototype = {
	initialize: function(selector, options) {
		this.selector = selector;
		this.options = neon.overed({
			name: "",
			onchange: function(val) {
				return true;
			},
			enabled: true
		}, options || {});

		this.enabled = this.options.enabled;

		var _T = this;

		this.list = selector.find("div.selectList").find("li").click(function() {
			_T.select($(this));
		}).end().find("li:first").addClass("first").end();

		this.curr = selector.find("strong").click(function() {
			if (this.enabled) this.list.show();
		}.bind(this));

		var val = selector.attr("_val");
		if (val) {
			this.list.find("li").each(function() {
				if (val == $(this).attr("_val")) {
					_T.selectedItem = $(this);
					_T.curr.text($(this).text());
				}
			});
		} else {
			var e = this.list.find("li:first");
			if (e) {
				this.selectedItem = e;
				val = e.attr("_val");
			}
		}

		this.input = $("<input/>").attr("type", "hidden").attr("name", this.options.name).val(val).appendTo(this.selector);
	},

	val: function(val) {
		if (typeof(val) == "undefined") return this.input.val();

		this.list.find("li[_val=" + val + "]").click();
	},

	select: function(e) {
		this.selectedItem = e;
		this.curr.text(e.text());
		this.input.val(this.selectedItem.attr("_val"));

		if (this.options.onchange(this.input.val())) {
			this.list.hide();
		}
	},

	add: function(text, val) {
		val = val || text;

		var _T = this;

		var e = $("<li></li>").attr("_val", val).html("<a href=\"javascript:void(0)\">" + text + "</a>").appendTo(this.list.find("ul")).click(function() {
			_T.select($(this));
		});

		if (this.list.find("li").length == 1) e.addClass("first");
	},

	setEnabled: function(enabled) {
		this.enabled = enabled;
	},

	size: function() {
		return this.list.find("li").length;
	}
};

pp.ui.Check = neon.create();
pp.ui.Check.prototype = {
	initialize: function(selector, options) {
		this.options = neon.overed({
			checkedClass: "checked",
			onChange: function(checked) {}
		}, options || {});

		var _T = this;

		this.checkbox = $(selector).click(function() {
			if ($(this).hasClass(_T.options.checkedClass)) {
				$(this).removeClass(_T.options.checkedClass);
				_T.input.attr("checked", false);

				_T.options.onChange.call(_T, false);
			} else {
				$(this).addClass(_T.options.checkedClass);
				_T.input.attr("checked", true);

				_T.options.onChange.call(_T, true);
			}
		});

		this.input = this.checkbox.find("input:checkbox");

		if (this.input.is(":checked")) this.checkbox.addClass(this.options.checkedClass);
		else this.checkbox.removeClass(this.options.checkedClass);
	},

	val: function(val) {
		if (typeof(val) == "undefined") return this.input.is(":checked");

		if (val) {
			this.checkbox.addClass(this.options.checkedClass);
			this.input.attr("checked", true);
		} else {
			this.checkbox.removeClass(this.options.checkedClass);
			this.input.attr("checked", false);
		}
	}
};

pp.ui.Radio = neon.create();
pp.ui.Radio.prototype = {
	initialize: function(selectors, options) {
		this.options = neon.overed({
			name: null,
			val: null,
			checkedClass: "checked",
			onchange: function(val) {}
		}, options || {});

		this.buttons = new Array();

		var _T = this;
		$(selectors).each(function() {
			_T.add($(this));
		});

		$.each(this.buttons, function(i, e) {
			if ($(e).find("input:radio").val() == _T.options.val) {
				$(e).click();
			}
		});
	},

	add: function(selector) {
		var _T = this;

		if (this.options.name) selector.find("input:radio").attr("name", this.options.name);

		this.buttons.push($(selector).click(function() {
			_T._val = $(this).find("input:radio").attr("checked", true).val();
			var idx = $(this).data("idx");

			$.each(_T.buttons, function(i, val) {
				if (idx == i) $(val).addClass(_T.options.checkedClass);
				else $(val).removeClass(_T.options.checkedClass);
			});

			_T.options.onchange(_T._val);
		}).data("idx", this.buttons.length));
	},

	val: function(val) {
		if (typeof(val) == "undefined") return this._val;

		$.each(this.buttons, function(i, e) {
			if ($(e).find("input:radio").val() == val) {
				$(e).click();
			}
		});
	},

	showOption: function(val) {
		$.each(this.buttons, function(i, e) {
			if ($(e).find("input:radio").val() != val) return;

			$(e).show().next("label").show();
		});
	},

	hideOption: function(val) {
		$.each(this.buttons, function(i, e) {
			if ($(e).find("input:radio").val() != val) return;

			$(e).hide().next("label").hide();
		});

		// hide되는 옵션이 현재 선택 옵션과 같을 경우 가장 처음 show되어있는 옵션을 선택
		if (this._val == val) {
			var found = false;
			$.each(this.buttons, function(i, e) {
				if (!found && $(e).is(":visible")) {
					found = true;
					$(e).click();
					return;
				}
			});
		}
	},

	showAllOptions: function() {
		$.each(this.buttons, function(i, e) {
			$(e).show().next("label").show();
		});
	},

	hideAllOptions: function() {
		$.each(this.buttons, function(i, e) {
			$(e).hide().next("label").hide();
		});
	}
};

pp.ui.PopupMng = {
	zIndex: 1000,
	popups: new Array(),
	layer: null,

	getBase: function() {
		return $("#wrap");
	},

	push: function(popup) {
		var base = $("#wrap");

		if (!this.layer) {
			this.layer = $("<div></div>").css({
				"z-index": this.zIndex - 1,
				"filter": "alpha(opacity=50)",
				"opacity": "0.5",
				"-moz-opacity": "0.5",
				"position": "absolute",
				"top": 0,
				"left": 0,
				"background-color": "black"
			}).appendTo(base);
		}

		this.layer.css({
			"z-index": this.zIndex + this.popups.length * 2 - 1,
			width: base.width(),
			height: base.height()
		}).show();

		popup.panel.css("z-index", this.zIndex + this.popups.length * 2).show().css({
			left: (base.width() - popup.panel.width()) / 2,
			top: (base.height() - popup.panel.height()) / 2,
			margin: 0
		});


		this.popups.push(popup);
	},

	pop: function() {
		this.popups.pop();

		if (this.popups.length == 0) this.layer.hide();
		else {
			this.layer.css({
				"z-index": this.zIndex + (this.popups.length - 1) * 2 - 1
			});
		}
	}

};

pp.ui.Popup = neon.create();
pp.ui.Popup.prototype = {
	initialize: function(content, options) {
		this.options = neon.overed({
			title: null,
			onOk: function() {
				return true;
			},
			onCancel: function() {
				return true;
			},
			onShow: function() {}
		}, options || {});

		if (typeof(content) == "object") {
			this.panel = $(content).clone().appendTo(pp.ui.PopupMng.getBase());
			//this.panel = $(content).appendTo(pp.ui.PopupMng.getBase());
			//this.base = $(content).clone();
		}

		var _T = this;

		// 제목
		if (this.options.title) {
			this.panel.find("h1").text(this.options.title);
		}

		this.panel.find("div.btns button").each(function() {
			var action = $(this).attr("action");
			if (action == "OK") {
				$(this).click(function() {
					if (this.options.onOk.call(this)) this.hide();
				}.bind(_T));
			} else if (action = "CANCEL") {
				$(this).click(function() {
					if (this.options.onCancel.call(this)) this.hide();
				}.bind(_T));
			}
		});

		this.panel.find("div.btnClose a").click(function() {
			if (this.options.onCancel.call(this)) this.hide();
		}.bind(this));
	},

	show: function(params) {
		params = params || {};

		//this.makePanel(params);
		pp.ui.PopupMng.push(this);

		this.options.onShow.call(this, params);
	},

	hide: function() {
		this.panel.remove();

		pp.ui.PopupMng.pop();
	}
};

pp.ui.Alert = {
	show: function(msg, options) {
		if (typeof(options) == "function") {
			options = {
				onOk: options
			};
		} else {
			options = neon.overed({
				onOk: function() {
					return true;
				}
			}, options || {});
		}

		alert(msg.replace("<br/>", "\n"));

		options.onOk.call(this);
	}
};

pp.ui.Confirm = {
	show: function(msg, options) {
		if (typeof(options) == "function") {
			options = {
				onOk: options
			};
		} else {
			options = neon.overed({
				onOk: function() {
					return true;
				}
			}, options || {});
		}

		if (confirm(msg)) {
			options.onOk.call(this);
		}
	}
};