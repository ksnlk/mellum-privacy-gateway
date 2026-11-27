See the Apache Version 2.0 License for specific language governing permissions
and limitations under the License.
***************************************************************************** */

/**
 * Interface for the AJAX setting that will configure the AJAX request
 */
declare interface JQueryAjaxSettings {
  /**
   * The content type sent in the request header that tells the server what kind of response it will accept in return. If the accepts setting needs modification, it is recommended to do so once in the $.ajaxSetup() method.
   */
  accepts?: any;
  /**
   * By default, all requests are sent asynchronously (i.e. this is set to true by default). If you need synchronous requests, set this option to false. Cross-domain requests and dataType: "jsonp" requests do not support synchronous operation. Note that synchronous requests may temporarily lock the browser, disabling any actions while the request is active. As of jQuery 1.8, the use of async: false with jqXHR ($.Deferred) is deprecated; you must use the success/error/complete callback options instead of the corresponding methods of the jqXHR object such as jqXHR.done() or the deprecated jqXHR.success().
   */
  async?: boolean;
  /**
   * A pre-request callback function that can be used to modify the jqXHR (in jQuery 1.4.x, XMLHTTPRequest: any) object before it is sent. Use this to set custom headers, etc. The jqXHR and settings objects are passed as arguments. This is an Ajax Event. Returning false in the beforeSend function will cancel the request. As of jQuery 1.5, the beforeSend option will be called regardless of the type of request.
   */
  beforeSend?: (jqXHR: JQueryXHR, settings: JQueryAjaxSettings) => any;
  /**
   * If set to false, it will force requested pages not to be cached by the browser. Note: Setting cache to false will only work correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. The parameter is not needed for other types of requests, except in IE8 when a POST is made to a URL that has already been requested by a GET.
   */
  cache?: boolean;
  /**
   * A function to be called when the request finishes (after success and error callbacks are executed). The function gets passed two arguments: The jqXHR (in jQuery 1.4.x, XMLHTTPRequest: any) object and a string categorizing the status of the request ("success", "notmodified", "error", "timeout", "abort", or "parsererror"). As of jQuery 1.5, the complete setting can accept an array of functions. Each function will be called in turn. This is an Ajax Event.
   */
  complete?: (jqXHR: JQueryXHR, textStatus: string) => any;
  /**
   * An object of string/regular-expression pairs that determine how jQuery will parse the response, given its content type. (version added: 1.5)
   */
  contents?: {
    [key: string]: any
  };
  // According to jQuery.ajax source code, ajax's option actually allows contentType to set to "false"
  // https://github.com/borisyankov/DefinitelyTyped/issues/742
  /**
   * When sending data to the server, use this content type. Default is "application/x-www-form-urlencoded; charset=UTF-8", which is fine for most cases. If you explicitly pass in a content-type to $.ajax(), then it is always sent to the server (even if no data is sent). The W3C XMLHttpRequest specification dictates that the charset is always UTF-8; specifying another charset will not force the browser to change the encoding.
   */
  contentType?: any;
  /**
   * This object will be made the context of all Ajax-related callbacks. By default, the context is an object that represents the ajax settings used in the call ($.ajaxSettings merged with the settings passed to $.ajax).
   */
  context?: any;
  /**
   * An object containing dataType-to-dataType converters. Each converter's value is a function that returns the transformed value of the response. (version added: 1.5)
   */
  converters?: {
    [key: string]: any
  };
  /**
   * If you wish to force a crossDomain request (such as JSONP) on the same domain, set the value of crossDomain to true. This allows, for example, server-side redirection to another domain. (version added: 1.5)
   */
  crossDomain?: boolean;
  /**
   * Data to be sent to the server. It is converted to a query string, if not already a string. It's appended to the url for GET-requests. See processData option to prevent this automatic processing. Object must be Key/Value pairs. If value is an Array, jQuery serializes multiple values with same key based on the value of the traditional setting (described below).
   */
  data?: any;
  /**
   * A function to be used to handle the raw response data of XMLHttpRequest.This is a pre-filtering function to sanitize the response. You should return the sanitized data. The function accepts two arguments: The raw data returned from the server and the 'dataType' parameter.
   */
  dataFilter?: (data: any, ty: any) => any;
  /**
   * The type of data that you're expecting back from the server. If none is specified, jQuery will try to infer it based on the MIME type of the response (an XML MIME type will yield XML, in 1.4 JSON will yield a JavaScript object, in 1.4 script will execute the script, and anything else will be returned as a string).
   */
  dataType?: string;
  /**
   * A function to be called if the request fails. The function receives three arguments: The jqXHR (in jQuery 1.4.x, XMLHttpRequest: any) object, a string describing the type of error that occurred and an optional exception object, if one occurred. Possible values for the second argument (besides null) are "timeout", "error", "abort", and "parsererror". When an HTTP error occurs, errorThrown receives the textual portion of the HTTP status, such as "Not Found" or "Internal Server Error." As of jQuery 1.5, the error setting can accept an array of functions. Each function will be called in turn. Note: This handler is not called for cross-domain script and cross-domain JSONP requests. This is an Ajax Event.
   */
  error?: (jqXHR: JQueryXHR, textStatus?: string, errorThrown?: string) => any;
  /**
   * Whether to trigger global Ajax event handlers for this request. The default is true. Set to false to prevent the global handlers like ajaxStart or ajaxStop from being triggered. This can be used to control various Ajax Events.
   */
  global?: boolean;
  /**
   * An object of additional header key/value pairs to send along with requests using the XMLHttpRequest transport. The header X-Requested-With: XMLHttpRequest is always added, but its default XMLHttpRequest value can be changed here. Values in the headers setting can also be overwritten from within the beforeSend function. (version added: 1.5)
   */
  headers?: {
    [key: string]: any
  };
  /**
   * Allow the request to be successful only if the response has changed since the last request. This is done by checking the Last-Modified header. Default value is false, ignoring the header. In jQuery 1.4 this technique also checks the 'etag' specified by the server to catch unmodified data.
   */
  ifModified?: boolean;
  /**
   * Allow the current environment to be recognized as "local," (e.g. the filesystem), even if jQuery does not recognize it as such by default. The following protocols are currently recognized as local: file, *-extension, and widget. If the isLocal setting needs modification, it is recommended to do so once in the $.ajaxSetup() method. (version added: 1.5.1)
   */
  isLocal?: boolean;
  /**
   * Override the callback function name in a jsonp request. This value will be used instead of 'callback' in the 'callback=?' part of the query string in the url. So {jsonp:'onJSONPLoad'} would result in 'onJSONPLoad=?' passed to the server. As of jQuery 1.5, setting the jsonp option to false prevents jQuery from adding the "?callback" string to the URL or attempting to use "=?" for transformation. In this case, you should also explicitly set the jsonpCallback setting. For example, { jsonp: false, jsonpCallback: "callbackName" }
   */
  jsonp?: any;
  /**
   * Specify the callback function name for a JSONP request. This value will be used instead of the random name automatically generated by jQuery. It is preferable to let jQuery generate a unique name as it'll make it easier to manage the requests and provide callbacks and error handling. You may want to specify the callback when you want to enable better browser caching of GET requests. As of jQuery 1.5, you can also use a function for this setting, in which case the value of jsonpCallback is set to the return value of that function.
   */
  jsonpCallback?: any;
  /**
   * The HTTP method to use for the request (e.g. "POST", "GET", "PUT"). (version added: 1.9.0)
   */
  method?: string;
  /**
   * A mime type to override the XHR mime type. (version added: 1.5.1)
   */
  mimeType?: string;
  /**
   * A password to be used with XMLHttpRequest in response to an HTTP access authentication request.
   */
  password?: string;
  /**
   * By default, data passed in to the data option as an object (technically, anything other than a string) will be processed and transformed into a query string, fitting to the default content-type "application/x-www-form-urlencoded". If you want to send a DOMDocument, or other non-processed data, set this option to false.
   */
  processData?: boolean;
  /**
   * Only applies when the "script" transport is used (e.g., cross-domain requests with "jsonp" or "script" dataType and "GET" type). Sets the charset attribute on the script tag used in the request. Used when the character set on the local page is not the same as the one on the remote script.
   */
  scriptCharset?: string;
  /**
   * An object of numeric HTTP codes and functions to be called when the response has the corresponding code. f the request is successful, the status code functions take the same parameters as the success callback; if it results in an error (including 3xx redirect), they take the same parameters as the error callback. (version added: 1.5)
   */
  statusCode?: {
    [key: string]: any
  };
  /**
   * A function to be called if the request succeeds. The function gets passed three arguments: The data returned from the server, formatted according to the dataType parameter; a string describing the status; and the jqXHR (in jQuery 1.4.x, XMLHttpRequest: any) object. As of jQuery 1.5, the success setting can accept an array of functions. Each function will be called in turn. This is an Ajax Event.
   */
  success?: (data: any, textStatus?: string, jqXHR?: JQueryXHR) => any;
  /**
   * Set a timeout (in milliseconds) for the request. This will override any global timeout set with $.ajaxSetup(). The timeout period starts at the point the $.ajax call is made; if several other requests are in progress and the browser has no connections available, it is possible for a request to time out before it can be sent. In jQuery 1.4.x and below, the XMLHttpRequest object will be in an invalid state if the request times out; accessing any object members may throw an exception. In Firefox 3.0+ only, script and JSONP requests cannot be cancelled by a timeout; the script will run even if it arrives after the timeout period.
   */
  timeout?: number;
  /**
   * Set this to true if you wish to use the traditional style of param serialization.
   */
  traditional?: boolean;
  /**
   * The type of request to make ("POST" or "GET"), default is "GET". Note: Other HTTP request methods, such as PUT and DELETE, can also be used here, but they are not supported by all browsers.
   */
  type?: string;
  /**
   * A string containing the URL to which the request is sent.
   */
  url?: string;
  /**
   * A username to be used with XMLHttpRequest in response to an HTTP access authentication request.
   */
  username?: string;
  /**
   * Callback for creating the XMLHttpRequest object. Defaults to the ActiveXObject when available (IE: any), the XMLHttpRequest otherwise. Override to provide your own implementation for XMLHttpRequest or enhancements to the factory.
   */
  xhr?: any;
  /**
   * An object of fieldName-fieldValue pairs to set on the native XHR object. For example, you can use it to set withCredentials to true for cross-domain requests if needed. In jQuery 1.5, the withCredentials property was not propagated to the native XHR and thus CORS requests requiring it would ignore this flag. For this reason, we recommend using jQuery 1.5.1+ should you require the use of it. (version added: 1.5.1)
   */
  xhrFields?: {
    [key: string]: any
  };
}

/**
 * Interface for the jqXHR object
 */
declare class JQueryXHR {
  /**
   * The .overrideMimeType() method may be used in the beforeSend() callback function, for example, to modify the response content-type header. As of jQuery 1.5.1, the jqXHR object also contains the overrideMimeType() method (it was available in jQuery 1.4.x, as well, but was temporarily removed in jQuery 1.5).
   */
  overrideMimeType(mimeType: string): any;
  /**
   * Cancel the request.
   *
   * @param statusText A string passed as the textStatus parameter for the done callback. Default value: "canceled"
   */
  abort(statusText?: string): void;
  /**
   * Incorporates the functionality of the .done() and .fail() methods, allowing (as of jQuery 1.8) the underlying Promise to be manipulated. Refer to deferred.then() for implementation details.
   */
  then<R>(
    doneCallback: (data: any, textStatus: string, jqXHR: JQueryXHR) => R,
    failCallback?: (
      jqXHR: JQueryXHR,
      textStatus: string,
      errorThrown: any
    ) => void
  ): JQueryPromise<R>;
  /**
   * Property containing the parsed response if the response Content-Type is json
   */
  responseJSON?: any;

  /**
   * Property containing the parsed response if the response Content-Type is text/plain
   */
  responseText?: string;

  /**
   * Status code of the response
   * @param html A string defining a single, standalone, HTML element (e.g. <div/> or <div></div>).
   * @param attributes An object of attributes, events, and methods to call on the newly-created element.
   */
  (html: string, attributes: {}): JQuery;

  /**
   * Relinquish jQuery's control of the $ variable.
   *
   * @param removeAll A Boolean indicating whether to remove all jQuery variables from the global scope (including jQuery itself).
   */
  noConflict(removeAll?: boolean): JQueryStatic;

  /**
   * Provides a way to execute callback functions based on one or more objects, usually Deferred objects that represent asynchronous events.
   *
   * @param deferreds One or more Deferred objects, or plain JavaScript objects.
   */
  when<T>(
    ...deferreds: Array<T | JQueryPromise<T> /* as JQueryDeferred<T> */>
  ): JQueryPromise<T>;

  /**
   * Hook directly into jQuery to override how particular CSS properties are retrieved or set, normalize CSS property naming, or create custom properties.
   */
  cssHooks: {
    [key: string]: any
  };
  cssNumber: any;

  /**
   * Store arbitrary data associated with the specified element. Returns the value that was set.
   *
   * @param element The DOM element to associate with the data.
   * @param key A string naming the piece of data to set.
   * @param value The new data value.
   */
  data<T>(element: Element, key: string, value: T): T;
  /**
   * Returns value: any at named data store for the element, as set by jQuery.data(element, name, value), or the full data store for the element.
   *
   * @param element The DOM element to associate with the data.
   * @param key A string naming the piece of data to set.
   */
  data(element: Element, key: string): any;
  /**
   * Returns value: any at named data store for the element, as set by jQuery.data(element, name, value), or the full data store for the element.
   *
   * @param element The DOM element to associate with the data.
   */
  data(element: Element, _: void): any;

  /**
   * Execute the next function on the queue for the matched element.
   *
   * @param element A DOM element from which to remove and execute a queued function.
   * @param queueName A string containing the name of the queue. Defaults to fx, the standard effects queue.
   */
  dequeue(element: Element, queueName?: string): void;

  /**
   * Determine whether an element has any jQuery data associated with it.
   *
   * @param element A DOM element to be checked for data.
   */
  hasData(element: Element): boolean;

  /**
   * Show the queue of functions to be executed on the matched element.
   *
  /**
   * Set the CSS width of each element in the set of matched elements.
   *
   * @param value An integer representing the number of pixels, or an integer along with an optional unit of measure appended (as a string).
   */
  width(value: number | string): JQuery;
  /**
   * Set the CSS width of each element in the set of matched elements.
   *
   * @param func A function returning the width to set. Receives the index position of the element in the set and the old width as arguments. Within the function, this refers to the current element in the set.
   */
  width(func: (index: number, width: number) => number | string): JQuery;

  /**
   * Remove from the queue all items that have not yet been run.
   *
   * @param queueName A string containing the name of the queue. Defaults to fx, the standard effects queue.
   */
  clearQueue(queueName?: string): JQuery;

  /**
   * Return the value: any at the named data store for the first element in the jQuery collection, as set by data(name, value) or by an HTML5 data-* attribute.
   *
   * @param key Name of the data stored.
   */
  data(key: string, _: void): any;
  /**
   * Store arbitrary data associated with the matched elements.
   *
   * @param key A string naming the piece of data to set.
   * @param value The new data value; it can be any Javascript type including Array or Object.
   */
  data(key: string, value: any): JQuery;
  /**
   * Store arbitrary data associated with the matched elements.
   *
   * @param obj An object of key-value pairs of data to update.
   */
  data(obj: {
    [key: string]: any
  }): JQuery;
  /**
   * Return the value: any at the named data store for the first element in the jQuery collection, as set by data(name, value) or by an HTML5 data-* attribute.
   */
  data(_: void): any;

  /**
   * Execute the next function on the queue for the matched elements.
   *
   * @param queueName A string containing the name of the queue. Defaults to fx, the standard effects queue.
   */
  dequeue(queueName?: string): JQuery;

  /**
   * Remove a previously-stored piece of data.
   *
   * @param name A string naming the piece of data to delete or space-separated string naming the pieces of data to delete.
   */
  removeData(name: string): JQuery;
  /**
   * Remove a previously-stored piece of data.
   *
   * @param list An array of strings naming the pieces of data to delete.
   */
  removeData(list: string[]): JQuery;
   * @param events One or more space-separated event types and optional namespaces, or just namespaces, such as "click", "keydown.myPlugin", or ".myPlugin".
   * @param selector A selector which should match the one originally passed to .on() when attaching event handlers.
   */
  off(
    events: string,
    selector: string
  ): JQuery;
  /**
   * Remove an event handler.
   *
   * @param events One or more space-separated event types and optional namespaces, or just namespaces, such as "click", "keydown.myPlugin", or ".myPlugin".
   * @param handler A handler function previous: anyly attached for the event(s), or the special value false. Takes handler with extra args that can be attached with on().
   */
  off(
    events: string,
    handler: (eventObject: JQueryEventObject, ...args: any[]) => any
  ): JQuery;
  /**
   * Remove an event handler.
   *
   * @param events: any An object where the string keys represent one or more space-separated event types and optional namespaces, and the values represent handler functions previously attached for the event(s).
   * @param selector A selector which should match the one originally passed to .on() when attaching event handlers.
   */
  off(
    events: {
      [key: string]: any
    },
    selector?: string
  ): JQuery;

  /**
   * Attach an event handler function for one or more events to the selected elements.
   *
   * @param events One or more space-separated event types and optional namespaces, such as "click" or "keydown.myPlugin".
   * @param handler A function to execute when the event is triggered. The value false is also allowed as a shorthand for a function that simply does return false. Rest parameter args is for optional parameters passed to jQuery.trigger(). Note that the actual parameters on the event handler function must be marked as optional (? syntax).
   */
  on(
    events: string,
    handler: (eventObject: JQueryEventObject, ...args: any[]) => any
  ): JQuery;
  /**
   * Attach an event handler function for one or more events to the selected elements.
   *
   * @param events One or more space-separated event types and optional namespaces, such as "click" or "keydown.myPlugin".
   * @param data Data to be passed to the handler in event.data when an event is triggered.
   * @param handler A function to execute when the event is triggered. The value false is also allowed as a shorthand for a function that simply does return false.
   */
  on(
    events: string,
    data: any,
    handler: (eventObject: JQueryEventObject, ...args: any[]) => any
    handler: (eventObject: JQueryEventObject, ...eventData: any[]) => any
  ): JQuery;
  /**
   * Attach an event handler function for one or more events to the selected elements.
   *
   * @param events One or more space-separated event types and optional namespaces, such as "click" or "keydown.myPlugin".
   * @param selector A selector string to filter the descendants of the selected elements that trigger the event. If the selector is null or omitted, the event is always triggered when it reaches the selected element.
   * @param data Data to be passed to the handler in event.data when an event is triggered.
   * @param handler A function to execute when the event is triggered. The value false is also allowed as a shorthand for a function that simply does return false.
   */
  on(
    events: string,
    selector: string,
    data: any,
    handler: (eventObject: JQueryEventObject, ...eventData: any[]) => any
  ): JQuery;
  /**
   * Attach an event handler function for one or more events to the selected elements.
   *
   * @param events: any An object in which the string keys represent one or more space-separated event types and optional namespaces, and the values represent a handler function to be called for the event(s).
   * @param selector A selector string to filter the descendants of the selected elements that will call the handler. If the selector is null or omitted, the handler is always called when it reaches the selected element.
   * @param data Data to be passed to the handler in event.data when an event occurs.
   */
  on(
    events: {
      [key: string]: any
    },
    selector?: string,
    data?: any
  ): JQuery;
  /**
   * Attach an event handler function for one or more events to the selected elements.
   *
   * @param events: any An object in which the string keys represent one or more space-separated event types and optional namespaces, and the values represent a handler function to be called for the event(s).
   * @param data Data to be passed to the handler in event.data when an event occurs.
   */
  on(
    events: {
      [key: string]: any
    },
    data?: any
  ): JQuery;

  /**
   * Attach a handler to an event for the elements. The handler is executed at most once per element per event type.
   *
   * @param events A string containing one or more JavaScript event types, such as "click" or "submit," or custom event names.
   * @param handler A function to execute at the time the event is triggered.
   */
  one(events: string, handler: (eventObject: JQueryEventObject) => any): JQuery;
  /**
   * Attach a handler to an event for the elements. The handler is executed at most once per element per event type.
   *
   * @param events A string containing one or more JavaScript event types, such as "click" or "submit," or custom event names.
   * @param data An object containing data that will be passed to the event handler.
   * @param handler A function to execute at the time the event is triggered.
   */
  one(
    events: string,
    data: Object,
    handler: (eventObject: JQueryEventObject) => any
  ): JQuery;

  /**
  ): JQuery;
  /**
   * Attach a handler to an event for the elements. The handler is executed at most once per element per event type.
   *
   * @param events One or more space-separated event types and optional namespaces, such as "click" or "keydown.myPlugin".
   * @param selector A selector string to filter the descendants of the selected elements that trigger the event. If the selector is null or omitted, the event is always triggered when it reaches the selected element.
   * @param data Data to be passed to the handler in event.data when an event is triggered.
   * @param handler A function to execute when the event is triggered. The value false is also allowed as a shorthand for a function that simply does return false.
   */
  one(
    events: string,
    selector: string,
    data: any,
    handler: (eventObject: JQueryEventObject) => any
  ): JQuery;

  /**
   * Attach a handler to an event for the elements. The handler is executed at most once per element per event type.
   *
   * @param events: any An object in which the string keys represent one or more space-separated event types and optional namespaces, and the values represent a handler function to be called for the event(s).
   * @param selector A selector string to filter the descendants of the selected elements that will call the handler. If the selector is null or omitted, the handler is always called when it reaches the selected element.
   * @param data Data to be passed to the handler in event.data when an event occurs.
   */
  one(
    events: {
      [key: string]: any
    },
    selector?: string,
    data?: any
  ): JQuery;

  /**
   * Attach a handler to an event for the elements. The handler is executed at most once per element per event type.
   *
   * @param events: any An object in which the string keys represent one or more space-separated event types and optional namespaces, and the values represent a handler function to be called for the event(s).
   * @param data Data to be passed to the handler in event.data when an event occurs.
   */
  one(
    events: {
      [key: string]: any
    },
    data?: any
  ): JQuery;

  /**
   * Specify a function to execute when the DOM is fully loaded.
   *
   * @param handler A function to execute after the DOM is ready.
   */
  ready(handler: (jQueryAlias?: JQueryStatic) => any): JQuery;

  /**
   * Trigger the "resize" event on an element.
   */
  resize(_: void): JQuery;
  /**
   * Bind an event handler to the "resize" JavaScript event.
   *
   * @param handler A function to execute each time the event is triggered.
   */
  resize(handler: (eventObject: JQueryEventObject) => any): JQuery;
  /**
   * Bind an event handler to the "resize" JavaScript event.
   *
   * @param eventData An object containing data that will be passed to the event handler.