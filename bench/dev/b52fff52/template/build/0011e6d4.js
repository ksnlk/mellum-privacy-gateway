  }
}

function processRawAttrs (el) {
  var l = el.attrsList.length;
  if (l) {
    var attrs = el.attrs = new Array(l);
    for (var i = 0; i < l; i++) {
      attrs[i] = {
        name: el.attrsList[i].name,
        value: JSON.stringify(el.attrsList[i].value)
      };
    }
  } else if (!el.pre) {
    // non root node in pre blocks with no attributes
    el.plain = true;
  }
}

function processKey (el) {
  var exp = getBindingAttr(el, 'key');
  if (exp) {
    if (process.env.NODE_ENV !== 'production' && el.tag === 'template') {
      warn$1("<template> cannot be keyed. Place the key on real elements instead.");
    }
    el.key = exp;
  }
}

function processRef (el) {
  var ref = getBindingAttr(el, 'ref');
  if (ref) {
    el.ref = ref;
    el.refInFor = checkInFor(el);
  }
}

function processFor (el) {
  var exp;
  if ((exp = getAndRemoveAttr(el, 'v-for'))) {
    var inMatch = exp.match(forAliasRE);
    if (!inMatch) {
      process.env.NODE_ENV !== 'production' && warn$1(
        ("Invalid v-for expression: " + exp)
      );
      return
    }
    el.for = inMatch[2].trim();
    var alias = inMatch[1].trim();
    var iteratorMatch = alias.match(forIteratorRE);
    if (iteratorMatch) {
  }
}

/*  */

var isStaticKey;
var isPlatformReservedTag;

var genStaticKeysCached = cached(genStaticKeys$1);

/**
 * Goal of the optimizer: walk the generated template AST tree
 * and detect sub-trees that are purely static, i.e. parts of
 * the DOM that never needs to change.
 *
 * Once we detect these sub-trees, we can:
 *
 * 1. Hoist them into constants, so that we no longer need to
 *    create fresh nodes for them on each re-render;
 * 2. Completely skip them in the patching process.
 */
function optimize (root, options) {
  if (!root) { return }
  isStaticKey = genStaticKeysCached(options.staticKeys || '');
  isPlatformReservedTag = options.isReservedTag || no;
  // first pass: mark all non-static nodes.
  markStatic(root);
  // second pass: mark static roots.
  markStaticRoots(root, false);
}

function genStaticKeys$1 (keys) {
  return makeMap(
    'type,tag,attrsList,attrsMap,plain,parent,children,attrs' +
    (keys ? ',' + keys : '')
  )
}

function markStatic (node) {
  node.static = isStatic(node);
  if (node.type === 1) {
    // do not make component slot content static. this avoids
    // 1. components not able to mutate slot nodes
    // 2. static slot content fails for hot-reloading
    if (
      !isPlatformReservedTag(node.tag) &&
      node.tag !== 'slot' &&
      node.attrsMap['inline-template'] == null
    ) {
      return
    }
    for (var i = 0, l = node.children.length; i < l; i++) {
      var child = node.children[i];
/*  */

var arrayKeys = Object.getOwnPropertyNames(arrayMethods);

/**
 * By default, when a reactive property is set, the new value is
 * also converted to become reactive. However when passing down props,
 * we don't want to force conversion because the value may be a nested value
 * under a frozen data structure. Converting it would defeat the optimization.
 */
var observerState = {
  shouldConvert: true
};

/**
 * Observer class that are attached to each observed
 * object. Once attached, the observer converts target
 * object's property keys into getter/setters that
 * collect dependencies and dispatches updates.
 */
var Observer = function Observer (value, key) {
  this.value = value;
  this.dep = new Dep();
  this.vmCount = 0;
  if (key) {
    this.key = key;
  }
  def(value, '__ob__', this);
  if (Array.isArray(value)) {
    var augment = hasProto
      ? protoAugment
      : copyAugment;
    augment(value, arrayMethods, arrayKeys);
    this.observeArray(value);
  } else {
    this.walk(value);
  }
};

/**
 * Walk through each property and convert them into
 * getter/setters. This method should only be called when
 * value type is Object.
 */
Observer.prototype.walk = function walk (obj) {
  var keys = Object.keys(obj);
  for (var i = 0; i < keys.length; i++) {
    defineReactive$$1(obj, keys[i], obj[keys[i]]);
  }
};

    state.warn(
      "<" + (el.tag) + " v-for=\"" + alias + " in " + exp + "\">: component lists rendered with " +
      "v-for should have explicit keys. " +
      "See https://vuejs.org/guide/list.html#key for more info.",
      true /* tip */
    );
  }

  el.forProcessed = true; // avoid recursion
  return (altHelper || '_l') + "((" + exp + ")," +
    "function(" + alias + iterator1 + iterator2 + "){" +
      "return " + ((altGen || genElement)(el, state)) +
    '})'
}

function genData$2 (el, state) {
  var data = '{';

  // directives first.
  // directives may mutate the el's other properties before they are generated.
  var dirs = genDirectives(el, state);
  if (dirs) { data += dirs + ','; }

  // key
  if (el.key) {
    data += "key:" + (el.key) + ",";
  }
  // ref
  if (el.ref) {
    data += "ref:" + (el.ref) + ",";
  }
  if (el.refInFor) {
    data += "refInFor:true,";
  }
  // pre
  if (el.pre) {
    data += "pre:true,";
  }
  // record original tag name for components using "is" attribute
  if (el.component) {
    data += "tag:\"" + (el.tag) + "\",";
  }
  // module data generation functions
  for (var i = 0; i < state.dataGenFns.length; i++) {
    data += state.dataGenFns[i](el);
  }
  // attributes
  if (el.attrs) {
    data += "attrs:{" + (genProps(el.attrs)) + "},";
  }
  // DOM props
    type: 2
  },
  'v-pre': noSupport,
  'v-cloak': noSupport,
  'v-once': {
    name: '',
    type: 5
  }
};

var tagConfig = {
  virtualTag: ['slot', 'template', 'block']
};

// babel-plugin-transform-object-to-ternary-operator.js

function getStrByNode (node, onlyStr) {
  if ( onlyStr === void 0 ) onlyStr = false;

  if (onlyStr) {
    return node.value || node.name || ''
  }
  return node.type === 'StringLiteral' ? node : t.stringLiteral(node.name || '')
}

// 把 { key: value } 转换成 [ value ? 'key' : '' ]
var objectVisitor = {
  ObjectExpression: function ObjectExpression (path) {
    var elements = path.node.properties.map(function (propertyItem) {
      return t.conditionalExpression(propertyItem.value, getStrByNode(propertyItem.key), t.stringLiteral(''))
    });
    path.replaceWith(t.arrayExpression(elements));
  }
};

function transformObjectToTernaryOperator (babel$$1) {
  return { visitor: objectVisitor }
}

// 把 { key: value } 转换成 'key:' + value + ';'
var objectToStringVisitor = {
  ObjectExpression: function ObjectExpression (path) {
    var expression = path.node.properties.map(function (propertyItem) {
      var keyStr = getStrByNode(propertyItem.key, true);
      var key = keyStr ? hyphenate(keyStr) : keyStr;
      var ref = generate(t.ExpressionStatement(propertyItem.value));
      var val = ref.code;
      return ("'" + key + ":' + (" + (val.slice(0, -1)) + ") + ';'")
    }).join('+');

    var p = template(expression)({});
    path.replaceWith(p.expression);
  }
};

function transformObjectToString (babel$$1) {
  return { visitor: objectToStringVisitor }
}

function transformDynamicClass (staticClass, clsBinding) {
  if ( staticClass === void 0 ) staticClass = '';

  var result = babel.transform(("!" + clsBinding), { plugins: [transformObjectToTernaryOperator] });
  // 先实现功能，再优化代码
  // https://github.com/babel/babel/issues/7138
  var cls = prettier.format(result.code, { semi: false, singleQuote: true }).slice(1).slice(0, -1).replace(/\n|\r/g, '');
  return (staticClass + " {{" + cls + "}}")
}

function transformDynamicStyle (staticStyle, styleBinding) {
  if ( staticStyle === void 0 ) staticStyle = '';

  var result = babel.transform(("!" + styleBinding), { plugins: [transformObjectToString] });
    Object.keys(attrs).map(function (key) {
      var val = attrs[key];
      obj[key.replace('@', 'v-on:').replace(/^:/, 'v-bind:')] = val;
    });

    return obj
  },

  convertAttr: function convertAttr (ast, log) {
    var this$1 = this;

    var attrsMap = ast.attrsMap; if ( attrsMap === void 0 ) attrsMap = {};
    var tag = ast.tag;
    var staticClass = ast.staticClass;
    var attrs = {};
    var wxClass = this.classObj(attrsMap['v-bind:class'], staticClass);
    wxClass.length ? attrsMap['class'] = wxClass : '';
    var wxStyle = this.styleObj(attrsMap['v-bind:style'], attrsMap['style']);
    wxStyle.length ? attrsMap['style'] = wxStyle : '';

    Object.keys(attrsMap).map(function (key) {
      var val = attrsMap[key];
      if (key === 'v-bind:class' || key === 'v-bind:style') {
        return
      }
      if (key === 'v-text') {
        ast.children.unshift({
          text: ("{{" + val + "}}"),
          type: 3
        });
      } else if (key === 'v-html') {
        ast.tag = 'rich-text';
        attrs['nodes'] = "{{" + val + "}}";
      } else if (key === 'v-show') {
        attrs['hidden'] = "{{!(" + val + ")}}";
      } else if (/^v\-on\:/i.test(key)) {
        attrs = this$1.event(key, val, attrs, tag);
      } else if (/^v\-bind\:/i.test(key)) {
        attrs = this$1.bind(key, val, attrs, tag, attrsMap['wx:key']);
      } else if (/^v\-model/.test(key)) {
        attrs = this$1.model(key, val, attrs, tag, log);
      } else if (directiveMap[key]) {
        var ref = directiveMap[key] || {};
        var name = ref.name; if ( name === void 0 ) name = '';
        var type = ref.type;
        var map = ref.map; if ( map === void 0 ) map = {};
        var check = ref.check;
        if (!(check && !check(key, val, log)) && !(!name || typeof type !== 'number')) {
          // 见 ./directiveMap.js 注释
          if (type === 0) {
            attrs[name] = "{{" + val + "}}";
          }

          if (type === 1) {
            attrs[name] = undefined;
          }

          if (type === 2) {
            attrs[name] = val;
          }

          if (type === 3) {
            attrs[map[name] || name] = "{{" + val + "}}";
            return
          }
        }
      } else if (/^v\-/.test(key)) {
        log(("不支持此属性-> " + key + "=\"" + val + "\""), 'waring');
      } else {
        if ((tagConfig.virtualTag.indexOf(tag) > -1) && (key === 'class' || key === 'style' || key === 'data-mpcomid')) {
          if (key !== 'data-mpcomid') {
            log(("template 不支持此属性-> " + key + "=\"" + val + "\""), 'waring');
          }
        } else {
          attrs[key] = val;
        }
      }
    });
    ast.attrsMap = attrs;
    return ast
  },

  event: function event (key, val, attrs, tag) {
    // 小程序能力所致，bind 和 catch 事件同时绑定时候，只会触发 bind ,catch 不会被触发。
    // .stop 的使用会阻止冒泡，但是同时绑定了一个非冒泡事件，会导致该元素上的 catchEventName 失效！
    // .prevent 可以直接干掉，因为小程序里没有什么默认事件，比如submit并不会跳转页面
    // .capture 不能做，因为小程序没有捕获类型的事件
    // .self 没有可以判断的标识
    // .once 也不能做，因为小程序没有 removeEventListener, 虽然可以直接在 handleProxy 中处理，但非常的不优雅，违背了原意，暂不考虑
    var name = key.replace(/^v\-on\:/i, '').replace(/\.prevent/i, '');
    var ref = name.split('.');
    var eventName = ref[0];
    var eventNameMap = ref.slice(1);
    var eventMap = directiveMap['v-on'];
    var check = directiveMap.check;

    if (check) {
    var tag = ast.tag;
    var mpcomid = ast.mpcomid;
    var slots = ast.slots;
    if (slotName) {
      attrsMap['data'] = "{{...$root[$p], ...$root[$k], $root}}";
      // bindedName is available when rendering slot in v-for
      var bindedName = attrsMap['v-bind:name'];
      if(bindedName) {
        attrsMap['is'] = "{{$for[" + bindedName + "]}}";
      } else {
        attrsMap['is'] = "{{" + slotName + "}}";
      }
    } else {
      var slotsName = getSlotsName(slots);
      var restSlotsName = slotsName ? (", " + slotsName) : '';
      attrsMap['data'] = "{{...$root[$kk+" + mpcomid + "], $root" + restSlotsName + "}}";
      attrsMap['is'] = components[tag].name;
    }
    return ast
  }
};

var astMap = {
  'if': 'wx:if',
  'iterator1': 'wx:for-index',
  'key': 'wx:key',
  'alias': 'wx:for-item',
  'v-for': 'wx:for'
};

var convertFor = function (ast) {
  var iterator1 = ast.iterator1;
  var forText = ast.for;
  var key = ast.key;
  var alias = ast.alias;
  var attrsMap = ast.attrsMap;

  if (forText) {
    attrsMap[astMap['v-for']] = "{{" + forText + "}}";
    if (iterator1) {
      attrsMap[astMap['iterator1']] = iterator1;
    }
    if (key) {
      attrsMap[astMap['key']] = key;
    }
    if (alias) {
      attrsMap[astMap['alias']] = alias;
    }

    delete attrsMap['v-for'];
  }
  log: function log (compiled) {
    compiled.mpErrors = [];
    compiled.mpTips = [];

    return function (str, type) {
      if (type === 'waring') {
        compiled.mpTips.push(str);
      } else {
        compiled.mpErrors.push(str);
      }
    }
  }
};

var autoEndTags = [
  'progress',
  'checkbox',
  'switch',
  'input',
  'radio',
  'slider',
  'textarea'
];

function convertAttr (key, val) {
  return (val === '' || typeof val === 'undefined') ? key : (key + "=\"" + val + "\"")
}

function generateCode (nodeAst, options) {
  if ( options === void 0 ) options = {};

  var tag = nodeAst.tag;
  var attrsMap = nodeAst.attrsMap; if ( attrsMap === void 0 ) attrsMap = {};
  var children = nodeAst.children; if ( children === void 0 ) children = [];
  var text = nodeAst.text;
  var ifConditions = nodeAst.ifConditions; if ( ifConditions === void 0 ) ifConditions = [];
  if (!tag) {
    return text
  }
  // v-if 指令
  var ifConditionsArr = [];
  if (ifConditions) {
    var length = ifConditions.length;
    for (var i = 1; i < length; i++) {
      ifConditionsArr.push(generateCode(ifConditions[i].block, options));
    }
  }
  var childrenContent = children.map(function (childAst) { return generateCode(childAst, options); }).join('');
  var attrs = Object.keys(attrsMap).map(function (key) { return convertAttr(key, attrsMap[key]); }).join(' ');
  attrs = attrs ? (" " + attrs) : '';

    Object.keys(attrs).map(function (key) {
      var val = attrs[key];
      obj[key.replace('@', 'v-on:').replace(/^:/, 'v-bind:')] = val;
    });

    return obj
  },

  convertAttr: function convertAttr (ast, log) {
    var this$1 = this;

    var attrsMap = ast.attrsMap; if ( attrsMap === void 0 ) attrsMap = {};
    var tag = ast.tag;
    var staticClass = ast.staticClass;
    var attrs = {};
    var wxClass = this.classObj(attrsMap['v-bind:class'], staticClass);
    wxClass.length ? attrsMap['class'] = wxClass : '';
    var wxStyle = this.styleObj(attrsMap['v-bind:style'], attrsMap['style']);
    wxStyle.length ? attrsMap['style'] = wxStyle : '';

    Object.keys(attrsMap).map(function (key) {
      var val = attrsMap[key];
      if (key === 'v-bind:class' || key === 'v-bind:style') {
        return
      }
      if (key === 'v-text') {
        ast.children.unshift({
          text: ("{{" + val + "}}"),
          type: 3
        });
      } else if (key === 'v-html') {
        ast.tag = 'rich-text';
        attrs['nodes'] = "{{" + val + "}}";
      } else if (key === 'v-show') {
        attrs['hidden'] = "{{!(" + val + ")}}";
      } else if (/^v\-on\:/i.test(key)) {
        attrs = this$1.event(key, val, attrs, tag);
      } else if (/^v\-bind\:/i.test(key)) {
        attrs = this$1.bind(key, val, attrs, tag, attrsMap['wx:key']);
      } else if (/^v\-model/.test(key)) {
        attrs = this$1.model(key, val, attrs, tag, log);
      } else if (directiveMap$1[key]) {
        var ref = directiveMap$1[key] || {};
        var name = ref.name; if ( name === void 0 ) name = '';
        var type = ref.type;
        var map = ref.map; if ( map === void 0 ) map = {};
        var check = ref.check;
        if (!(check && !check(key, val, log)) && !(!name || typeof type !== 'number')) {
          // 见 ./directiveMap.js 注释
          if (type === 0) {
            attrs[name] = "{{" + val + "}}";
          }

          if (type === 1) {
            attrs[name] = undefined;
          }

          if (type === 2) {
            attrs[name] = val;
          }

          if (type === 3) {
            attrs[map[name] || name] = "{{" + val + "}}";
            return
          }
        }
      } else if (/^v\-/.test(key)) {
        log(("不支持此属性-> " + key + "=\"" + val + "\""), 'waring');
      } else {
        if ((tagConfig.virtualTag.indexOf(tag) > -1) && (key === 'class' || key === 'style' || key === 'data-mpcomid')) {
          if (key !== 'data-mpcomid') {
            log(("template 不支持此属性-> " + key + "=\"" + val + "\""), 'waring');
          }
        } else {
          attrs[key] = val;
        }
      }
    });
    ast.attrsMap = attrs;
    return ast
  },

  event: function event (key, val, attrs, tag) {
    // 小程序能力所致，bind 和 catch 事件同时绑定时候，只会触发 bind ,catch 不会被触发。
    // .stop 的使用会阻止冒泡，但是同时绑定了一个非冒泡事件，会导致该元素上的 catchEventName 失效！
    // .prevent 可以直接干掉，因为小程序里没有什么默认事件，比如submit并不会跳转页面
    // .capture 不能做，因为小程序没有捕获类型的事件
    // .self 没有可以判断的标识
    // .once 也不能做，因为小程序没有 removeEventListener, 虽然可以直接在 handleProxy 中处理，但非常的不优雅，违背了原意，暂不考虑
    var name = key.replace(/^v\-on\:/i, '').replace(/\.prevent/i, '');
    var ref = name.split('.');
    var eventName = ref[0];
    var eventNameMap = ref.slice(1);
    var eventMap = directiveMap$1['v-on'];
    var check = directiveMap$1.check;

    if (check) {
    var slots = ast.slots;
    if (slotName) {
      attrsMap['data'] = '{{{...$root[$k], $root}}}';
      // bindedName is available when rendering slot in v-for
      var bindedName = attrsMap['v-bind:name'];
      if (bindedName) {
        attrsMap['is'] = '{{$for[' + bindedName + ']}}';
      } else {
        attrsMap['is'] = '{{' + slotName + '}}';
      }
    } else {
      var slotsName = getSlotsName$1(slots);
      var restSlotsName = slotsName ? (", " + slotsName) : '';
      attrsMap['data'] = "{{{...$root[$kk+" + mpcomid + "], $root" + restSlotsName + "}}}";
      attrsMap['is'] = components[tag].name;
    }
    return ast
  }
};

var astMap$1 = {
  'if': 's-if',
  'v-for': 's-for',
  'alias': 's-for-item',
  'iterator1': 's-for-index',
  'key': 's-key'
};

var convertFor$1 = function (ast) {
  var iterator1 = ast.iterator1;
  var forText = ast.for;
  var key = ast.key;
  var alias = ast.alias;
  var attrsMap = ast.attrsMap;

  // 缩写：<view s-for="p,index in persons">
  // 全写：<view s-for="persons" s-for-index="index" s-for-item="p">

  if (forText) {
    attrsMap[astMap$1['v-for']] = alias + "," + iterator1 + " in " + forText;
    // attrsMap[astMap['v-for']] = forText
    // if (iterator1) {
    //   attrsMap[astMap['iterator1']] = iterator1
    // }
    // if (alias) {
    //   attrsMap[astMap['alias']] = alias
    // }
    // if (key) {
    //   attrsMap[astMap['key']] = key
    // }
    delete attrsMap['v-for'];
    Object.keys(attrs).map(function (key) {
      var val = attrs[key];
      obj[key.replace('@', 'v-on:').replace(/^:/, 'v-bind:')] = val;
    });

    return obj
  },

  convertAttr: function convertAttr (ast, log) {
    var this$1 = this;

    var attrsMap = ast.attrsMap; if ( attrsMap === void 0 ) attrsMap = {};
    var tag = ast.tag;
    var staticClass = ast.staticClass;
    var attrs = {};
    var wxClass = this.classObj(attrsMap['v-bind:class'], staticClass);
    wxClass.length ? attrsMap['class'] = wxClass : '';
    var wxStyle = this.styleObj(attrsMap['v-bind:style'], attrsMap['style']);
    wxStyle.length ? attrsMap['style'] = wxStyle : '';

    Object.keys(attrsMap).map(function (key) {
      var val = attrsMap[key];
      if (key === 'v-bind:class' || key === 'v-bind:style') {
        return
      }
      if (key === 'v-text') {
        ast.children.unshift({
          text: ("{{" + val + "}}"),
          type: 3
        });
      } else if (key === 'v-html') {
        ast.tag = 'rich-text';
        attrs['nodes'] = "{{" + val + "}}";
      } else if (key === 'v-show') {
        attrs['hidden'] = "{{!(" + val + ")}}";
      } else if (/^v\-on\:/i.test(key)) {
        attrs = this$1.event(key, val, attrs, tag);
      } else if (/^v\-bind\:/i.test(key)) {
        attrs = this$1.bind(key, val, attrs, tag, attrsMap['tt:key']);
      } else if (/^v\-model/.test(key)) {
        attrs = this$1.model(key, val, attrs, tag, log);
      } else if (directiveMap$2[key]) {
        var ref = directiveMap$2[key] || {};
        var name = ref.name; if ( name === void 0 ) name = '';
        var type = ref.type;
        var map = ref.map; if ( map === void 0 ) map = {};
        var check = ref.check;
        if (!(check && !check(key, val, log)) && !(!name || typeof type !== 'number')) {
          // 见 ./directiveMap.js 注释
          if (type === 0) {
            attrs[name] = "{{" + val + "}}";
          }

          if (type === 1) {
            attrs[name] = undefined;
          }

          if (type === 2) {
            attrs[name] = val;
          }

          if (type === 3) {
            attrs[map[name] || name] = "{{" + val + "}}";
            return
          }
        }
      } else if (/^v\-/.test(key)) {
        log(("不支持此属性-> " + key + "=\"" + val + "\""), 'waring');
      } else {
        if ((tagConfig.virtualTag.indexOf(tag) > -1) && (key === 'class' || key === 'style' || key === 'data-mpcomid')) {
          if (key !== 'data-mpcomid') {
            log(("template 不支持此属性-> " + key + "=\"" + val + "\""), 'waring');
          }
        } else {
          attrs[key] = val;
        }
      }
    });
    ast.attrsMap = attrs;
    return ast
  },

  event: function event (key, val, attrs, tag) {
    // 小程序能力所致，bind 和 catch 事件同时绑定时候，只会触发 bind ,catch 不会被触发。
    // .stop 的使用会阻止冒泡，但是同时绑定了一个非冒泡事件，会导致该元素上的 catchEventName 失效！
    // .prevent 可以直接干掉，因为小程序里没有什么默认事件，比如submit并不会跳转页面
    // .capture 不能做，因为小程序没有捕获类型的事件
    // .self 没有可以判断的标识
    // .once 也不能做，因为小程序没有 removeEventListener, 虽然可以直接在 handleProxy 中处理，但非常的不优雅，违背了原意，暂不考虑
    var name = key.replace(/^v\-on\:/i, '').replace(/\.prevent/i, '');
    var ref = name.split('.');
    var eventName = ref[0];
    var eventNameMap = ref.slice(1);
    var eventMap = directiveMap$2['v-on'];
    var check = directiveMap$2.check;

    if (check) {
    var tag = ast.tag;
    var mpcomid = ast.mpcomid;
    var slots = ast.slots;
    if (slotName) {
      attrsMap['data'] = "{{...$root[$p], ...$root[$k], $root}}";
      // bindedName is available when rendering slot in v-for
      var bindedName = attrsMap['v-bind:name'];
      if(bindedName) {
        attrsMap['is'] = "{{$for[" + bindedName + "]}}";
      } else {
        attrsMap['is'] = "{{" + slotName + "}}";
      }
    } else {
      var slotsName = getSlotsName$2(slots);
      var restSlotsName = slotsName ? (", " + slotsName) : '';
      attrsMap['data'] = "{{...$root[$kk+" + mpcomid + "], $root" + restSlotsName + "}}";
      attrsMap['is'] = components[tag].name;
    }
    return ast
  }
};

var astMap$2 = {
  'if': 'tt:if',
  'iterator1': 'tt:for-index',
  'key': 'tt:key',
  'alias': 'tt:for-item',
  'v-for': 'tt:for'
};

var convertFor$2 = function (ast) {
  var iterator1 = ast.iterator1;
  var forText = ast.for;
  var key = ast.key;
  var alias = ast.alias;
  var attrsMap = ast.attrsMap;
  if (forText) {
    attrsMap[astMap$2['v-for']] = "{{" + forText + "}}";
    if (iterator1) {
      attrsMap[astMap$2['iterator1']] = iterator1;
    }
    if (key) {
      attrsMap[astMap$2['key']] = key;
    }
    if (alias) {
      attrsMap[astMap$2['alias']] = alias;
    }

    delete attrsMap['v-for'];
  }


    Object.keys(attrs).map(function (key) {
      var val = attrs[key];
      obj[key.replace('@', 'v-on:').replace(/^:/, 'v-bind:')] = val;
    });
    return obj
  },

  convertAttr: function convertAttr (ast, log) {
    var this$1 = this;

    var attrsMap = ast.attrsMap; if ( attrsMap === void 0 ) attrsMap = {};
    var tag = ast.tag;
    var staticClass = ast.staticClass;
    var attrs = {};
    var wxClass = this.classObj(attrsMap['v-bind:class'], staticClass);
    wxClass.length ? attrsMap['class'] = wxClass : '';
    var wxStyle = this.styleObj(attrsMap['v-bind:style'], attrsMap['style']);
    wxStyle.length ? attrsMap['style'] = wxStyle : '';

    Object.keys(attrsMap).map(function (key) {
      var val = attrsMap[key];
      if (key === 'v-bind:class' || key === 'v-bind:style') {
        return
      }
      if (key === 'v-text') {
        ast.children.unshift({
          text: ("{{" + val + "}}"),
          type: 3
        });
      } else if (key === 'v-html') {
        ast.tag = 'rich-text';
        attrs['nodes'] = "{{" + val + "}}";
      } else if (key === 'v-show') {
        attrs['hidden'] = "{{!(" + val + ")}}";
      } else if (/^v\-on\:/i.test(key)) {
        attrs = this$1.event(key, val, attrs, tag, log);
      } else if (/^v\-bind\:/i.test(key)) {
        attrs = this$1.bind(key, val, attrs, tag, attrsMap['a:key']);
      } else if (/^v\-model/.test(key)) {
        attrs = this$1.model(key, val, attrs, tag, log);
      } else if (directiveMap$3[key]) {
        var ref = directiveMap$3[key] || {};
        var name = ref.name; if ( name === void 0 ) name = '';
        var type = ref.type;
        var map = ref.map; if ( map === void 0 ) map = {};
        var check = ref.check;
        if (!(check && !check(key, val, log)) && !(!name || typeof type !== 'number')) {
          // 见 ./directiveMap.js 注释
          if (type === 0) {
            attrs[name] = "{{" + val + "}}";
          }

          if (type === 1) {
            attrs[name] = undefined;
          }

          if (type === 2) {
            attrs[name] = val;
          }

          if (type === 3) {
            attrs[map[name] || name] = "{{" + val + "}}";
            return
          }
        }
      } else if (/^v\-/.test(key)) {
        log(("不支持此属性-> " + key + "=\"" + val + "\""), 'waring');
      } else {
        if ((tagConfig.virtualTag.indexOf(tag) > -1) && (key === 'class' || key === 'style' || key === 'data-mpcomid')) {
          if (key !== 'data-mpcomid') {
            log(("template 不支持此属性-> " + key + "=\"" + val + "\""), 'waring');
          }
        } else {
          attrs[key] = val;
        }
      }
    });
    ast.attrsMap = attrs;
    return ast
  },

  event: function event (key, val, attrs, tag, log) {
    // 小程序能力所致，bind 和 catch 事件同时绑定时候，只会触发 bind ,catch 不会被触发。
    // .stop 的使用会阻止冒泡，但是同时绑定了一个非冒泡事件，会导致该元素上的 catchEventName 失效！
    // .prevent 可以直接干掉，因为小程序里没有什么默认事件，比如submit并不会跳转页面
    // .capture 不能做，因为小程序没有捕获类型的事件
    // .self 没有可以判断的标识
    // .once 也不能做，因为小程序没有 removeEventListener, 虽然可以直接在 handleProxy 中处理，但非常的不优雅，违背了原意，暂不考虑
    var name = key.replace(/^v\-on\:/i, '').replace(/\.prevent/i, '');
    var ref = name.split('.');
    var eventName = ref[0];
    var eventNameMap = ref.slice(1);
    var eventMap = directiveMap$3['v-on'];
    var check = directiveMap$3.check;

    if (check) {
    var tag = ast.tag;
    var mpcomid = ast.mpcomid;
    var slots = ast.slots;
    if (slotName) {
      attrsMap['data'] = "{{...$root[$p], ...$root[$k], $root}}";
      // bindedName is available when rendering slot in v-for
      var bindedName = attrsMap['v-bind:name'];
      if (bindedName) {
        attrsMap['is'] = "{{$for[" + bindedName + "]}}";
      } else {
        attrsMap['is'] = "{{" + slotName + "}}";
      }
    } else {
      var slotsName = getSlotsName$3(slots);
      var restSlotsName = slotsName ? (", " + slotsName) : '';
      attrsMap['data'] = "{{...$root[$kk+" + mpcomid + "], $root" + restSlotsName + "}}";
      attrsMap['is'] = components[tag].name;
    }
    return ast
  }
};

var astMap$3 = {
  'if': 'a:if',
  'iterator1': 'a:for-index',
  'key': 'a:key',
  'alias': 'a:for-item',
  'v-for': 'a:for'
};

var convertFor$3 = function (ast) {
  var iterator1 = ast.iterator1;
  var forText = ast.for;
  var key = ast.key;
  var alias = ast.alias;
  var attrsMap = ast.attrsMap;

  if (forText) {
    attrsMap[astMap$3['v-for']] = "{{" + forText + "}}";
    if (iterator1) {
      attrsMap[astMap$3['iterator1']] = iterator1;
    }
    if (key) {
      attrsMap[astMap$3['key']] = key;
    }
    if (alias) {
      attrsMap[astMap$3['alias']] = alias;
    }

    delete attrsMap['v-for'];
  }