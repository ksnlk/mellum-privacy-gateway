            disable: 'disable',
            enable: 'enable',
            reset: 'reset'
        }, function (fn, command) {
            Uploader.prototype[fn] = function () {
                return this.request(command, arguments);
            };
        });

        $.extend(Uploader.prototype, {
            state: 'pending',

            _init: function (opts) {
                var me = this;

                me.request('init', opts, function () {
                    me.state = 'ready';
                    me.trigger('ready');
                });
            },

            /**
            * 获取或者设置Uploader配置项。
            * @method option
            * @grammar option( key ) => *
            * @grammar option( key, val ) => self
            * @example
            *
            * // 初始状态图片上传前不会压缩
            * var uploader = new WebUploader.Uploader({
            *     compress: null;
            * });
            *
            * // 修改后图片上传前，尝试将图片压缩到1600 * 1600
            * uploader.option( 'compress', {
            *     width: 1600,
            *     height: 1600
            * });
            */
            option: function (key, val) {
                var opts = this.options;

                // setter
                if (arguments.length > 1) {

                    if ($.isPlainObject(val) &&
                            $.isPlainObject(opts[key])) {
                        $.extend(opts[key], val);
                    } else {
                        opts[key] = val;
                    }
                            rlts.push(rlt);
                        }
                    }
                }

                // 如果有callback，则用异步方式。
                if (callback || dfds.length) {
                    promise = Base.when.apply(Base, dfds);
                    key = promise.pipe ? 'pipe' : 'then';

                    // 很重要不能删除。删除了会死循环。
                    // 保证执行顺序。让callback总是在下一个 tick 中执行。
                    return promise[key](function () {
                        var deferred = Base.Deferred(),
                                    args = arguments;

                        if (args.length === 1) {
                            args = args[0];
                        }

                        setTimeout(function () {
                            deferred.resolve(args);
                        }, 1);

                        return deferred.promise();
                    })[callback ? key : 'done'](callback || Base.noop);
                } else {
                    return rlts[0];
                }
            }
        });

        /**
        * 添加组件
        * @param  {object} widgetProto 组件原型，构造函数通过constructor属性定义
        * @param  {object} responseMap API名称与函数实现的映射
        * @example
        *     Uploader.register( {
        *         init: function( options ) {},
        *         makeThumb: function() {}
        *     }, {
        *         'make-thumb': 'makeThumb'
        *     } );
        */
        Uploader.register = Widget.register = function (responseMap, widgetProto) {
            var map = { init: 'init' },
                klass;

            if (arguments.length === 1) {
                widgetProto = responseMap;
                widgetProto.responseMap = map;