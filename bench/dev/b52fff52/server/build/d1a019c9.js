    while (isDef(parent)) {
      if (isDef(parent.data) && isDef(parent.data.attrs)) {
        attrs = Object.assign({}, attrs, parent.data.attrs);
      }
      parent = parent.parent;
    }
  }

  if (isUndef(attrs)) {
    return res
  }

  for (var key in attrs) {
    if (key === 'style') {
      // leave it to the style module
      continue
    }
    res += renderAttr(key, attrs[key]);
  }
  return res
}

function renderAttr (key, value) {
  if (isBooleanAttr(key)) {
    if (!isFalsyAttrValue(value)) {
      return (" " + key + "=\"" + key + "\"")
    }
  } else if (isEnumeratedAttr(key)) {
    return (" " + key + "=\"" + (isFalsyAttrValue(value) || value === 'false' ? 'false' : 'true') + "\"")
  } else if (!isFalsyAttrValue(value)) {
    return (" " + key + "=\"" + (typeof value === 'string' ? cachedEscape(value) : value) + "\"")
  }
  return ''
}

/*  */

var VNode = function VNode (
  tag,
  data,
  children,
  text,
  elm,
  context,
  componentOptions,
  asyncFactory
) {
  this.tag = tag;
  this.data = data;
  this.children = children;
  this.text = text;
  this.elm = elm;
  this.ns = undefined;
  this.context = context;
  this.functionalContext = undefined;
  this.key = data && data.key;
  this.componentOptions = componentOptions;
  this.componentInstance = undefined;
  this.parent = undefined;
  this.raw = false;
  this.isStatic = false;
  this.isRootInsert = true;
  this.isComment = false;
  this.isCloned = false;
  this.isOnce = false;
  this.asyncFactory = asyncFactory;
  this.asyncMeta = undefined;
  this.isAsyncPlaceholder = false;
};

var prototypeAccessors = { child: {} };

// DEPRECATED: alias for componentInstance for backwards compat.
/* istanbul ignore next */
prototypeAccessors.child.get = function () {
  return this.componentInstance
};

Object.defineProperties( VNode.prototype, prototypeAccessors );

var createEmptyVNode = function (text) {
    }
  }

  if ((styleData = normalizeStyleData(vnode.data))) {
    extend(res, styleData);
  }

  var parentNode = vnode;
  while ((parentNode = parentNode.parent)) {
    if (parentNode.data && (styleData = normalizeStyleData(parentNode.data))) {
      extend(res, styleData);
    }
  }
  return res
}

/*  */

function genStyle (style) {
  var styleText = '';
  for (var key in style) {
    var value = style[key];
    var hyphenatedKey = hyphenate(key);
    if (Array.isArray(value)) {
      for (var i = 0, len = value.length; i < len; i++) {
        styleText += hyphenatedKey + ":" + (value[i]) + ";";
      }
    } else {
      styleText += hyphenatedKey + ":" + value + ";";
    }
  }
  return styleText
}

function renderStyle (vnode) {
  var styleText = genStyle(getStyle(vnode, false));
  if (styleText !== '') {
    return (" style=" + (JSON.stringify(cachedEscape(styleText))))
  }
}

var modules = [
  renderAttrs,
  renderDOMProps,
  renderClass,
  renderStyle
];

/*  */

function show (node, dir) {
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
      warn$2("<template> cannot be keyed. Place the key on real elements instead.");
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
      process.env.NODE_ENV !== 'production' && warn$2(
        ("Invalid v-for expression: " + exp)
      );
      return
    }
    el.for = inMatch[2].trim();
    var alias = inMatch[1].trim();
    var iteratorMatch = alias.match(forIteratorRE);
    if (iteratorMatch) {
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
    if (isUndef(c) || typeof c === 'boolean') { continue }
    last = res[res.length - 1];
    //  nested
    if (Array.isArray(c)) {
      res.push.apply(res, normalizeArrayChildren(c, ((nestedIndex || '') + "_" + i)));
    } else if (isPrimitive(c)) {
      if (isTextNode(last)) {
        // merge adjacent text nodes
        // this is necessary for SSR hydration because text nodes are
        // essentially merged when rendered to HTML strings
        (last).text += String(c);
      } else if (c !== '') {
        // convert primitive to vnode
        res.push(createTextVNode(c));
      }
    } else {
      if (isTextNode(c) && isTextNode(last)) {
        // merge adjacent text nodes
        res[res.length - 1] = createTextVNode(last.text + c.text);
      } else {
        // default key for nested array children (likely generated by v-for)
        if (isTrue(children._isVList) &&
          isDef(c.tag) &&
          isUndef(c.key) &&
          isDef(nestedIndex)) {
          c.key = "__vlist" + nestedIndex + "_" + i + "__";
        }
        res.push(c);
      }
    }
  }
  return res
}

/*  */

function installSSRHelpers (vm) {
  if (vm._ssrNode) { return }
  var Ctor = vm.constructor;
  while (Ctor.super) {
    Ctor = Ctor.super;
  }
  Object.assign(Ctor.prototype, {
    _ssrEscape: escape,
    _ssrNode: renderStringNode$1,
    _ssrList: renderStringList,
    _ssrAttr: renderAttr,
    _ssrAttrs: renderAttrs$1,
    _ssrDOMProps: renderDOMProps$1,
    _ssrClass: renderSSRClass,
    _ssrStyle: renderSSRStyle
}

function registerComponentForCache (options, write) {
  // exposed by vue-loader, need to call this if cache hit because
  // component lifecycle hooks will not be called.
  var register = options._ssrRegister;
  if (write.caching && isDef(register)) {
    write.componentBuffer[write.componentBuffer.length - 1].add(register);
  }
  return register
}

function renderComponent (node, isRoot, context) {
  var write = context.write;
  var next = context.next;
  var userContext = context.userContext;

  // check cache hit
  var Ctor = node.componentOptions.Ctor;
  var getKey = Ctor.options.serverCacheKey;
  var name = Ctor.options.name;
  var cache = context.cache;
  var registerComponent = registerComponentForCache(Ctor.options, write);

  if (isDef(getKey) && isDef(cache) && isDef(name)) {
    var key = name + '::' + getKey(node.componentOptions.propsData);
    var has = context.has;
    var get = context.get;
    if (isDef(has)) {
      has(key, function (hit) {
        if (hit === true && isDef(get)) {
          get(key, function (res) {
            if (isDef(registerComponent)) {
              registerComponent(userContext);
            }
            res.components.forEach(function (register) { return register(userContext); });
            write(res.html, next);
          });
        } else {
          renderComponentWithCache(node, isRoot, key, context);
        }
      });
    } else if (isDef(get)) {
      get(key, function (res) {
        if (isDef(res)) {
          if (isDef(registerComponent)) {
            registerComponent(userContext);
          }
          res.components.forEach(function (register) { return register(userContext); });
          write(res.html, next);
        } else {
  var m = trace.match(filenameRE);
  var map = m && mapConsumers[m[1]];
  if (m != null && map) {
    var originalPosition = map.originalPositionFor({
      line: Number(m[2]),
      column: Number(m[3])
    });
    if (originalPosition.source != null) {
      var source = originalPosition.source;
      var line = originalPosition.line;
      var column = originalPosition.column;
      var mappedPosition = "(" + (source.replace(/^webpack:\/\/\//, '')) + ":" + (String(line)) + ":" + (String(column)) + ")";
      return trace.replace(filenameRE, mappedPosition)
    } else {
      return trace
    }
  } else {
    return trace
  }
}

/*  */

var fs = require('fs');
var path$1 = require('path');
var PassThrough = require('stream').PassThrough;

var INVALID_MSG =
  'Invalid server-rendering bundle format. Should be a string ' +
  'or a bundle Object of type:\n\n' +
"{\n  entry: string;\n  files: { [filename: string]: string; };\n  maps: { [filename: string]: string; };\n}\n";

// The render bundle can either be a string (single bundled file)
// or a bundle manifest object generated by vue-ssr-webpack-plugin.


function createBundleRendererCreator (
  createRenderer
) {
  return function createBundleRenderer (
    bundle,
    rendererOptions
  ) {
    if ( rendererOptions === void 0 ) rendererOptions = {};

    var files, entry, maps;
    var basedir = rendererOptions.basedir;

    // load bundle if given filepath
    if (
      typeof bundle === 'string' &&