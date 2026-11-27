this._x == internal - only use within this file
this.__x == private - only use within the class

Globals should be all caps
*/
(function (mixpanel) {

/*
 * Saved references to long variable names, so that closure compiler can
 * minimize file size.
 */
    var   ArrayProto        = Array.prototype
        , ObjProto          = Object.prototype
        , slice             = ArrayProto.slice
        , toString          = ObjProto.toString
        , hasOwnProperty    = ObjProto.hasOwnProperty
        , windowConsole     = window.console
        , navigator         = window.navigator
        , document          = window.document
        , userAgent         = navigator.userAgent;

/*
 * Constants
 */
/** @const */   var   PRIMARY_INSTANCE_NAME     = "mixpanel"
/** @const */       , SET_QUEUE_KEY             = "__mps"
/** @const */       , SET_ONCE_QUEUE_KEY        = "__mpso"
/** @const */       , ADD_QUEUE_KEY             = "__mpa"
/** @const */       , APPEND_QUEUE_KEY          = "__mpap"
/** @const */       , SET_ACTION                = "$set"
/** @const */       , SET_ONCE_ACTION           = "$set_once"
/** @const */       , ADD_ACTION                = "$add"
/** @const */       , APPEND_ACTION             = "$append"
// This key is deprecated, but we want to check for it to see whether aliasing is allowed.
/** @const */       , PEOPLE_DISTINCT_ID_KEY    = "$people_distinct_id"
/** @const */       , ALIAS_ID_KEY              = "__alias"
/** @const */       , RESERVED_PROPERTIES       = [SET_QUEUE_KEY, SET_ONCE_QUEUE_KEY, ADD_QUEUE_KEY, APPEND_QUEUE_KEY, PEOPLE_DISTINCT_ID_KEY, ALIAS_ID_KEY];

/*
 * Dynamic... constants? Is that an oxymoron?
 */
    var   HTTP_PROTOCOL     = (("https:" == document.location.protocol) ? "https://" : "http://")
        , SNIPPET_VERSION   = (mixpanel && mixpanel['__SV']) || 0
        // http://hacks.mozilla.org/2009/07/cross-site-xmlhttprequest-with-cors/
        // https://developer.mozilla.org/en-US/docs/DOM/XMLHttpRequest#withCredentials
        , USE_XHR           = (window.XMLHttpRequest && 'withCredentials' in new XMLHttpRequest())
        // IE<10 does not support cross-origin XHR's but script tags
        // with defer won't block window.onload; ENQUEUE_REQUESTS
        // should only be true for Opera<12
        , ENQUEUE_REQUESTS  = !USE_XHR && (userAgent.indexOf('MSIE') == -1);

/*
 * Closure-level globals
 */
    var   _                 = {}
        , DEBUG             = false
        , DEFAULT_CONFIG    = {
              "api_host":                   HTTP_PROTOCOL + 'api.mixpanel.com'
            , "cross_subdomain_cookie":     true
            , "cookie_name":                ""
            , "loaded":                     function() {}
                        array.push(value());
                        white();
                        if (ch === ']') {
                            next(']');
                            return array;
                        }
                        next(',');
                        white();
                    }
                }
                error("Bad array");
            },
            object = function () {
                // Parse an object value.
                var key,
                    object = {};

                if (ch === '{') {
                    next('{');
                    white();
                    if (ch === '}') {
                        next('}');
                        return object;   // empty object
                    }
                    while (ch) {
                        key = string();
                        white();
                        next(':');
                        if (Object.hasOwnProperty.call(object, key)) {
                            error('Duplicate key "' + key + '"');
                        }
                        object[key] = value();
                        white();
                        if (ch === '}') {
                            next('}');
                            return object;
                        }
                        next(',');
                        white();
                    }
                }
                error("Bad object");
            };

        value = function () {
            // Parse a JSON value. It could be an object, an array, a string,
            // a number, or a word.
            white();
            switch (ch) {
            case '{':
                return object();

    // _.isBlockedUA()
    // This is to block various web spiders from executing our JS and
    // sending false tracking data
    _.isBlockedUA = function() {
        var a = userAgent;
        if (/(google web preview|baiduspider|yandexbot)/i.test(a)) {
            return true;
        }
        return false;
    };

    /**
     * @param {Object=} formdata
     * @param {string=} arg_separator
     */
    _.HTTPBuildQuery = function(formdata, arg_separator) {
        var key, use_val, use_key, tmp_arr = [];

        if (typeof(arg_separator) === "undefined") {
            arg_separator = '&';
        }

        _.each(formdata, function(val, key) {
            use_val = encodeURIComponent(val.toString());
            use_key = encodeURIComponent(key);
            tmp_arr[tmp_arr.length] = use_key + '=' + use_val;
        });

        return tmp_arr.join(arg_separator);
    };

    _.getQueryParam = function(url, param) {
        // Expects a raw URL

        param = param.replace(/[\[]/, "\\\[").replace(/[\]]/, "\\\]");
        var regexS = "[\\?&]" + param + "=([^&#]*)",
            regex = new RegExp( regexS ),
            results = regex.exec(url);
        if (results === null || (results && typeof(results[1]) !== 'string' && results[1].length)) {
            return '';
        } else {
            return decodeURIComponent(results[1]).replace(/\+/g, ' ');
        }
    };

    // _.cookie
    // Methods partially borrowed from quirksmode.org/js/cookies.html
    _.cookie = {
        get: function(name) {
            var nameEQ = name + "=";
                    var foundCount = 0;
                    for (var h = 0; h < currentContext.length; h++) {
                        var elements;
                        if (tagName == '*') {
                            elements = getAllChildren(currentContext[h]);
                        } else {
                            elements = currentContext[h].getElementsByTagName(tagName);
                        }
                        for (var j = 0; j < elements.length; j++) {
                            found[foundCount++] = elements[j];
                        }
                    }
                    currentContext = new Array;
                    var currentContextIndex = 0;
                    for (var k = 0; k < found.length; k++) {
                        if (found[k].className
                            && _.isString(found[k].className)   // some SVG elements have classNames which are not strings
                            && hasClass(found[k], className)
                        ) {
                            currentContext[currentContextIndex++] = found[k];
                        }
                    }
                    continue; // Skip to next token
                }
                // Code to deal with attribute selectors
                if (token.match(/^(\w*)\[(\w+)([=~\|\^\$\*]?)=?"?([^\]"]*)"?\]$/)) {
                    var tagName = RegExp.$1;
                    var attrName = RegExp.$2;
                    var attrOperator = RegExp.$3;
                    var attrValue = RegExp.$4;
                    if (!tagName) {
                        tagName = '*';
                    }
                    // Grab all of the tagName elements within current context
                    var found = new Array;
                    var foundCount = 0;
                    for (var h = 0; h < currentContext.length; h++) {
                        var elements;
                        if (tagName == '*') {
                            elements = getAllChildren(currentContext[h]);
                        } else {
                            elements = currentContext[h].getElementsByTagName(tagName);
                        }
                        for (var j = 0; j < elements.length; j++) {
                            found[foundCount++] = elements[j];
                        }
                    }
                    currentContext = new Array;
                    var currentContextIndex = 0;
                    var checkFunction; // This function will be used to filter the elements
                    switch (attrOperator) {

    MixpanelCookie.prototype.get_referrer_info = function() {
        return _.strip_empty_properties({
            '$initial_referrer': this['props']['$initial_referrer'],
            '$initial_referring_domain': this['props']['$initial_referring_domain']
        });
    };

    // safely fills the passed in object with the cookies properties,
    // does not override any properties defined in both
    // returns the passed in object
    MixpanelCookie.prototype.safe_merge = function(props) {
        _.each(this['props'], function(val, prop) {
            if (!(prop in props)) {
                props[prop] = val;
            }
        });

        return props;
    };

    MixpanelCookie.prototype.update_config = function(config) {
        this.default_expiry = this.expire_days = config['cookie_expiration'];
        this.set_disabled(config['disable_cookie']);
        this.set_cross_subdomain(config['cross_subdomain_cookie']);
        this.set_secure(config['secure_cookie']);
    };

    MixpanelCookie.prototype.set_disabled = function(disabled) {
        this.disabled = disabled;
        if (this.disabled) {
            this.remove();
        }
    };

    MixpanelCookie.prototype.set_cross_subdomain = function(cross_subdomain) {
        if (cross_subdomain !== this.cross_subdomain) {
            this.cross_subdomain = cross_subdomain;
            this.remove();
            this.save();
        }
    };

    MixpanelCookie.prototype.get_cross_subdomain = function() {
        return this.cross_subdomain;
    };

    MixpanelCookie.prototype.set_secure = function(secure) {
        if (secure !== this.secure) {
            this.secure = secure ? true : false;
            this.remove();
                    }
                    add_q[k] += v;
                }
            }, this);
        } else if (q_key === APPEND_QUEUE_KEY) {
            append_q.push(q_data);
        }

        console.log("MIXPANEL PEOPLE REQUEST (QUEUED, PENDING IDENTIFY):");
        console.log(data);

        this.save();
    };

    MixpanelCookie.prototype._pop_from_people_queue = function(queue, data) {
        var q = this._get_queue(queue);
        if (!_.isUndefined(q)) {
            _.each(data, function(v, k) {
                delete q[k];
            }, this);

            this.save();
        }
    };

    MixpanelCookie.prototype._get_queue_key = function(queue) {
        if (queue === SET_ACTION) {
            return SET_QUEUE_KEY;
        } else if (queue === SET_ONCE_ACTION) {
            return SET_ONCE_QUEUE_KEY;
        } else if (queue === ADD_ACTION) {
            return ADD_QUEUE_KEY;
        } else if (queue === APPEND_ACTION) {
            return APPEND_QUEUE_KEY;
        } else {
            console.error("Invalid queue:", queue);
        }
    };

    MixpanelCookie.prototype._get_queue = function(queue) {
        return this['props'][this._get_queue_key(queue)];
    };
    MixpanelCookie.prototype._get_or_create_queue = function(queue, default_val) {
        var key = this._get_queue_key(queue),
            default_val = _.isUndefined(default_val) ? {} : default_val;

        return this['props'][key] || (this['props'][key] = default_val);
    };

    /**
     * create_mplib(token:string, config:object, name:string)
     *
     * This function is used by the init method of MixpanelLib objects
     * as well as the main initializer at the end of the JSLib (that
     * initializes document.mixpanel as well as any additional instances
     * declared before this file has loaded).
     */
    var create_mplib = function(token, config, name) {
        var instance, target = (name === PRIMARY_INSTANCE_NAME) ? mixpanel : mixpanel[name];

        if (target && !_.isArray(target)) {
            console.error("You have already initialized " + name);
            return;
        }

        instance = new MixpanelLib();
        instance._init(token, config, name);

        instance['people'] = new MixpanelPeople();
        instance['people']._init(instance);

        // if any instance on the page has debug = true, we set the
        // global debug to be true
        DEBUG = DEBUG || instance.get_config('debug');

        // if target is not defined, we called init after the lib already
     * and use it like this:
     *
     *      mixpanel.library_name.track(...)
     *
     * @param {String} token   Your Mixpanel API token
     * @param {Object} config  A dictionary of config options to override
     * @param {String} name    The name for the new mixpanel instance that you want created
     */
    MixpanelLib.prototype.init = function (token, config, name) {
        if (typeof(name) === "undefined") {
            console.error("You must name your new library: init(token, config, name)");
            return;
        }
        if (name === PRIMARY_INSTANCE_NAME) {
            console.error("You must initialize the main mixpanel object right after you include the Mixpanel js snippet");
            return;
        }

        var instance = create_mplib(token, config, name);
        mixpanel[name] = instance;
        instance._loaded();

        return instance;
    };

    // mixpanel._init(token:string, config:object, name:string)
    //
    // This function sets up the current instance of the mixpanel
    // library.  The difference between this method and the init(...)
    // method is this one initializes the actual instance, whereas the
    // init(...) method sets up a new library and calls _init on it.
    //
    MixpanelLib.prototype._init = function(token, config, name) {
        this['__loaded'] = true;
        this['config'] = {};

        this.set_config(_.extend({}, DEFAULT_CONFIG, config, {
              "name": name
            , "token": token
            , "callback_fn": ((name === PRIMARY_INSTANCE_NAME) ? name : PRIMARY_INSTANCE_NAME + '.' + name) + '._jsc'
        }));

        this['_jsc'] = function() {};

        this.__dom_loaded_queue = [];
        this.__request_queue = [];
        this.__disabled_events = [];
        this._flags = {
              "disable_all_events": false
            , "identify_called": false
        };

        this['cookie'] = new MixpanelCookie(this['config']);
        this.register_once({'distinct_id': _.UUID()}, "");
    };

    // Private methods

    MixpanelLib.prototype._loaded = function() {
        this.get_config('loaded')(this);

        // this happens after so a user can call identify/name_tag in
        // the loaded callback
        if (this.get_config('track_pageview')) {