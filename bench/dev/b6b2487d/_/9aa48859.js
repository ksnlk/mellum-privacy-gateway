
  declare type TruncateOptions = {
    length?: number,
    omission?: string,
    separator?: RegExp | string
  };

  declare type Cancelable = {
    cancel: () => void,
    flush: () => mixed
  };

  declare type DebounceOptions = {
    leading?: boolean,
    maxWait?: number,
    trailing?: boolean
  };

  declare type ThrottleOptions = {
    leading?: boolean,
    trailing?: boolean
  };

  declare type NestedArray<T> = Array<Array<T>>;

  declare type matchesIterateeShorthand = {[key: any]: any};
  declare type matchesPropertyIterateeShorthand = [string, any];
  declare type propertyIterateeShorthand = string;

  declare type OPredicate<A, O> =
    | ((value: A, key: string, object: O) => any)
    | matchesIterateeShorthand
    | matchesPropertyIterateeShorthand
    | propertyIterateeShorthand;

  declare type OIterateeWithResult<V, O, R> =
    | Object
    | string
    | ((value: V, key: string, object: O) => R);
  declare type OIteratee<O> = OIterateeWithResult<any, O, any>;
  declare type OFlatMapIteratee<T, U> = OIterateeWithResult<any, T, Array<U>>;

  declare type Predicate<T> =
    | ((value: T, index: number, array: Array<T>) => any)
    | matchesIterateeShorthand
    | matchesPropertyIterateeShorthand
    | propertyIterateeShorthand;

  declare type _ValueOnlyIteratee<T> = (value: T) => mixed;
  declare type ValueOnlyIteratee<T> = _ValueOnlyIteratee<T> | string;
  declare type _Iteratee<T> = (
    item: T,
    index: number,
    array: ?Array<T>
  ) => mixed;
  declare type Iteratee<T> = _Iteratee<T> | Object | string;
  declare type FlatMapIteratee<T, U> =
    | ((item: T, index: number, array: ?$ReadOnlyArray<T>) => Array<U>)
    | Object
    | string;
  declare type Comparator<T> = (item: T, item2: T) => boolean;

  declare type MapIterator<T, U> =
    | ((item: T, index: number, array: Array<T>) => U)
    | propertyIterateeShorthand;

  declare type ReadOnlyMapIterator<T, U> =
    | ((item: T, index: number, array: $ReadOnlyArray<T>) => U)
    | propertyIterateeShorthand;

  declare type OMapIterator<T, O, U> =
    | ((item: T, key: string, object: O) => U)
    | propertyIterateeShorthand;

  declare class Lodash {
    // Array
    chunk<T>(array?: ?Array<T>, size?: ?number): Array<Array<T>>;
    compact<T, N: ?T>(array?: ?Array<N>): Array<T>;
    concat<T>(
      base?: ?$ReadOnlyArray<T>,
      ...elements: Array<any>
    ): Array<T | any>;
    difference<T>(
      array?: ?$ReadOnlyArray<T>,
      ...values: Array<?$ReadOnlyArray<T>>
    ): Array<T>;
    differenceBy<T>(
      array?: ?$ReadOnlyArray<T>,
      values?: ?$ReadOnlyArray<T>,
      iteratee?: ?ValueOnlyIteratee<T>
    ): T[];
    differenceWith<T>(
      array?: ?$ReadOnlyArray<T>,
      values?: ?$ReadOnlyArray<T>,
      comparator?: ?Comparator<T>
    ): T[];
    drop<T>(array?: ?Array<T>, n?: ?number): Array<T>;
      object: T,
      predicate?: OPredicate<A, T>
    ): [Array<V>, Array<V>];
    reduce<T, U>(
      array: $ReadOnlyArray<T>,
      iteratee?: (
        accumulator: U,
        value: T,
        index: number,
        array: ?Array<T>
      ) => U,
      accumulator?: U
    ): U;
    reduce<T, U>(
      array: void | null,
      iteratee?: ?(
        accumulator: U,
        value: T,
        index: number,
        array: ?Array<T>
      ) => U,
      accumulator?: ?U
    ): void | null;
    reduce<T: Object, U>(
      object: T,
      iteratee?: (accumulator: U, value: any, key: string, object: T) => U,
      accumulator?: U
    ): U;
    reduceRight<T, U>(
      array: void | null,
      iteratee?: ?(
        accumulator: U,
        value: T,
        index: number,
        array: ?Array<T>
      ) => U,
      accumulator?: ?U
    ): void | null;
    reduceRight<T, U>(
      array: $ReadOnlyArray<T>,
      iteratee?: ?(
        accumulator: U,
        value: T,
        index: number,
        array: ?Array<T>
      ) => U,
      accumulator?: ?U
    ): U;
    reduceRight<T: Object, U>(
      object: T,
      iteratee?: ?(accumulator: U, value: any, key: string, object: T) => U,
      accumulator?: ?U
    ): U;
    reject<T>(array: ?$ReadOnlyArray<T>, predicate?: Predicate<T>): Array<T>;
    reject<V: Object, A, T: { [id: any]: A }>(
      object?: ?T,
      predicate?: ?OPredicate<A, T>
    ): Array<V>;
    sample<T>(array: ?$ReadOnlyArray<T>): T;
    sample<V, T: Object>(object: T): V;
    sampleSize<T>(array?: ?$ReadOnlyArray<T>, n?: ?number): Array<T>;
    sampleSize<V, T: Object>(object: T, n?: number): Array<V>;
    shuffle<T>(array: ?$ReadOnlyArray<T>): Array<T>;
    shuffle<V, T: Object>(object: T): Array<V>;
    size(collection: $ReadOnlyArray<any> | Object | string): number;
    some<T>(array: void | null, predicate?: ?Predicate<T>): false;
    some<T>(array: ?$ReadOnlyArray<T>, predicate?: Predicate<T>): boolean;
    some<A, T: { [id: any]: A }>(
      object?: ?T,
      predicate?: OPredicate<A, T>
    ): boolean;
    sortBy<T>(
      array: ?$ReadOnlyArray<T>,
      ...iteratees?: $ReadOnlyArray<Iteratee<T>>
    ): Array<T>;
    sortBy<T>(
    negate<A, R>(predicate: (...A) => R): (...A) => boolean;
    once<F: (...any[]) => any>(func: F): F;
    overArgs(func?: ?Function, ...transforms?: Array<Function>): Function;
    overArgs(func?: ?Function, transforms?: ?Array<Function>): Function;
    partial<R>(func: (...any[]) => R, ...partials: any[]): (...any[]) => R;
    partialRight<R>(func: (...any[]) => R, ...partials: Array<any>): (...any[]) => R;
    partialRight<R>(func: (...any[]) => R, partials: Array<any>): (...any[]) => R;
    rearg(func: Function, ...indexes: Array<number>): Function;
    rearg(func: Function, indexes: Array<number>): Function;
    rest(func: Function, start?: number): Function;
    spread(func: Function): Function;
    throttle<F: (...any[]) => any>(
      func: F,
      wait?: number,
      options?: ThrottleOptions
    ): F & Cancelable;
    unary<F: (...any[]) => any>(func: F): F;
    wrap(value?: any, wrapper?: ?Function): Function;

    // Lang
    castArray(value: *): any[];
    clone<T>(value: T): T;
    cloneDeep<T>(value: T): T;
    cloneDeepWith<T, U>(
      value: T,
      customizer?: ?(value: T, key: number | string, object: T, stack: any) => U
    ): U;
    cloneWith<T, U>(
      value: T,
      customizer?: ?(value: T, key: number | string, object: T, stack: any) => U
    ): U;
    conformsTo<T: { [key: string]: mixed }>(
      source: T,
      predicates: T & { [key: string]: (x: any) => boolean }
    ): boolean;
    eq(value: any, other: any): boolean;
    gt(value: any, other: any): boolean;
    gte(value: any, other: any): boolean;
    isArguments(value: void | null): false;
    isArguments(value: any): boolean;
    isArray(value: $ReadOnlyArray<any>): true;
    isArray(value: any): false;
    isArrayBuffer(value: ArrayBuffer): true;
    isArrayBuffer(value: any): false;
    isArrayLike(value: $ReadOnlyArray<any> | string | { length: number }): true;
    isArrayLike(value: any): false;
    isArrayLikeObject(value: { length: number } | $ReadOnlyArray<any>): true;
    isArrayLikeObject(value: any): false;
    isBoolean(value: boolean): true;
    isBoolean(value: any): false;
    isBuffer(value: void | null): false;
    isBuffer(value: any): boolean;
    isDate(value: Date): true;
    isDate(value: any): false;
    isElement(value: Element): true;
    isElement(value: any): false;
    isEmpty(value: void | null | "" | {} | [] | number | boolean): true;
    isEmpty(value: any): boolean;
    isEqual(value: any, other: any): boolean;
    isEqualWith<T, U>(
      value?: ?T,
      other?: ?U,
      customizer?: ?(
        objValue: any,
        otherValue: any,
        key: number | string,
        object: T,
        other: U,
        stack: any
      ) => boolean | void
    ): boolean;
    isError(value: Error): true;
    isError(value: any): false;
    isFinite(value: number): boolean;
    isFinite(value: any): false;
    isFunction(value: Function): true;
    isFunction(value: any): false;
    isInteger(value: number): boolean;
    isInteger(value: any): false;
    isLength(value: void | null): false;
    isLength(value: any): boolean;
    isMap(value: Map<any, any>): true;
    isMap(value: any): false;
    isMatch(object?: ?Object, source?: ?Object): boolean;
    isMatchWith<T: Object, U: Object>(
      object?: ?T,
      source?: ?U,
      customizer?: ?(
        objValue: any,
        srcValue: any,
        key: number | string,
        object: T,
        source: U
      ) => boolean | void
    ): boolean;
    isNaN(value: number): boolean;
    isNaN(value: any): false;
    isNative(value: number | string | void | null | Object): false;
    isNative(value: any): boolean;
    isNil(value: void | null): true;
    isNil(value: any): false;
    isNull(value: null): true;
    isNull(value: any): false;
    isNumber(value: number): true;
    isNumber(value: any): false;
    isObject(value: any): boolean;
    isObjectLike(value: void | null): false;
    isObjectLike(value: any): boolean;
    isPlainObject(value: any): boolean;
    isRegExp(value: RegExp): true;
    isRegExp(value: any): false;
    isSafeInteger(value: number): boolean;
    isSafeInteger(value: any): false;
    isSet(value: Set<any>): true;
    isSet(value: any): false;
    isString(value: string): true;
    round(number: number, precision?: number): number;
    subtract(minuend: number, subtrahend: number): number;
    sum(array: $ReadOnlyArray<*>): number;
    sumBy<T>(array: $ReadOnlyArray<T>, iteratee?: Iteratee<T>): number;

    // number
    clamp(number?: number, lower?: ?number, upper?: ?number): number;
    clamp(number: ?number, lower?: ?number, upper?: ?number): 0;
    inRange(number: number, start?: number, end: number): boolean;
    random(lower?: number, upper?: number, floating?: boolean): number;

    // Object
    assign(object?: ?Object, ...sources?: Array<?Object>): Object;
    assignIn(): {};
    assignIn<A, B>(a: A, b: B): A & B;
    assignIn<A, B, C>(a: A, b: B, c: C): A & B & C;
    assignIn<A, B, C, D>(a: A, b: B, c: C, d: D): A & B & C & D;
    assignIn<A, B, C, D, E>(a: A, b: B, c: C, d: D, e: E): A & B & C & D & E;
    assignInWith(): {};
    assignInWith<T: Object, A: Object>(
      object: T,
      s1: A,
      customizer?: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A
      ) => any | void
    ): Object;
    assignInWith<T: Object, A: Object, B: Object>(
      object: T,
      s1: A,
      s2: B,
      customizer?: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A | B
      ) => any | void
    ): Object;
    assignInWith<T: Object, A: Object, B: Object, C: Object>(
      object: T,
      s1: A,
      s2: B,
      s3: C,
      customizer?: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A | B | C
      ) => any | void
    ): Object;
    assignInWith<T: Object, A: Object, B: Object, C: Object, D: Object>(
      object: T,
      s1: A,
      s2: B,
      s3: C,
      s4: D,
      customizer?: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A | B | C | D
      ) => any | void
    ): Object;
    assignWith(): {};
    assignWith<T: Object, A: Object>(
      object: T,
      s1: A,
      customizer?: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A
      ) => any | void
    ): Object;
    assignWith<T: Object, A: Object, B: Object>(
      object: T,
      s1: A,
      s2: B,
      customizer?: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A | B
      ) => any | void
    ): Object;
    assignWith<T: Object, A: Object, B: Object, C: Object>(
      object: T,
      s1: A,
      s2: B,
      s3: C,
      customizer?: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A | B | C
      ) => any | void
    ): Object;
    assignWith<T: Object, A: Object, B: Object, C: Object, D: Object>(
      object: T,
      s1: A,
      s2: B,
      s3: C,
      s4: D,
      customizer?: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A | B | C | D
      ) => any | void
    ): Object;
    at(object?: ?Object, ...paths: $ReadOnlyArray<string>): Array<any>;
    at(object?: ?Object, paths: $ReadOnlyArray<string>): Array<any>;
    create(prototype: void | null, properties: void | null): {};
    create<T>(prototype: T, properties: Object): T;
    create(prototype: any, properties: void | null): {};
    defaults(object?: ?Object, ...sources?: Array<?Object>): Object;
    defaultsDeep(object?: ?Object, ...sources?: Array<?Object>): Object;
    // alias for _.toPairs
    entries(object?: ?Object): Array<[string, any]>;
    // alias for _.toPairsIn
    entriesIn(object?: ?Object): Array<[string, any]>;
    // alias for _.assignIn
    extend<A, B>(a?: ?A, b?: ?B): A & B;
    extend<A, B, C>(a: A, b: B, c: C): A & B & C;
    extend<A, B, C, D>(a: A, b: B, c: C, d: D): A & B & C & D;
    extend<A, B, C, D, E>(a: A, b: B, c: C, d: D, e: E): A & B & C & D & E;
    // alias for _.assignInWith
    extendWith<T: Object, A: Object>(
      object?: ?T,
      s1?: ?A,
      customizer?: ?(
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A
      ) => any | void
    ): Object;
    extendWith<T: Object, A: Object, B: Object>(
      object: T,
      s1: A,
      s2: B,
      customizer?: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A | B
      ) => any | void
    ): Object;
    extendWith<T: Object, A: Object, B: Object, C: Object>(
      object: T,
      s1: A,
      s2: B,
      s3: C,
      customizer?: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A | B | C
      ) => any | void
    ): Object;
    extendWith<T: Object, A: Object, B: Object, C: Object, D: Object>(
      object: T,
      s1: A,
      s2: B,
      s3: C,
      s4: D,
      customizer?: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A | B | C | D
      ) => any | void
    ): Object;
    findKey<A, T: { [id: any]: A }>(
      object: T,
      predicate?: ?OPredicate<A, T>
    ): string | void;
    findKey<A, T: { [id: any]: A }>(
      object: void | null,
      predicate?: ?OPredicate<A, T>
    ): void;
    findLastKey<A, T: { [id: any]: A }>(
      object: T,
      predicate?: ?OPredicate<A, T>
    ): string | void;
    findLastKey<A, T: { [id: any]: A }>(
      object: void | null,
      predicate?: ?OPredicate<A, T>
    ): void;
    forIn(object: Object, iteratee?: ?OIteratee<*>): Object;
    forIn(object: void | null, iteratee?: ?OIteratee<*>): null;
    forInRight(object: Object, iteratee?: ?OIteratee<*>): Object;
    forInRight(object: void | null, iteratee?: ?OIteratee<*>): null;
    forOwn(object: Object, iteratee?: ?OIteratee<*>): Object;
    hasIn(object: Object, path: void | null): false;
    hasIn(object: void | null, path?: ?Path): false;
    invert(object: Object, multiVal?: ?boolean): Object;
    invert(object: void | null, multiVal?: ?boolean): {};
    invertBy(object: Object, iteratee?: ?Function): Object;
    invertBy(object: void | null, iteratee?: ?Function): {};
    invoke(
      object?: ?Object,
      path?: ?Path,
      ...args?: Array<any>
    ): any;
    keys<K>(object?: ?{ [key: K]: any }): Array<K>;
    keys(object?: ?Object): Array<string>;
    keysIn(object?: ?Object): Array<string>;
    mapKeys(object: Object, iteratee?: ?OIteratee<*>): Object;
    mapKeys(object: void | null, iteratee?: ?OIteratee<*>): {};
    mapValues(object: Object, iteratee?: ?OIteratee<*>): Object;
    mapValues(object: void | null, iteratee?: ?OIteratee<*>): {};
    merge(object?: ?Object, ...sources?: Array<?Object>): Object;
    mergeWith(): {};
    mergeWith<T: Object, A: Object>(
      object: T,
      customizer?: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A
      ) => any | void
    ): Object;
    mergeWith<T: Object, A: Object, B: Object>(
      object: T,
      s1: A,
      s2: B,
      customizer?: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A | B
      ) => any | void
    ): Object;
    mergeWith<T: Object, A: Object, B: Object, C: Object>(
      object: T,
      s1: A,
      s2: B,
      s3: C,
      customizer?: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A | B | C
      ) => any | void
    ): Object;
    mergeWith<T: Object, A: Object, B: Object, C: Object, D: Object>(
      object: T,
      s1: A,
      s2: B,
      s3: C,
      s4: D,
      customizer?: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A | B | C | D
      ) => any | void
    ): Object;
    omit(object?: ?Object, ...props: $ReadOnlyArray<string>): Object;
    omit(object?: ?Object, props: $ReadOnlyArray<string>): Object;
    omitBy<A, T: { [id: any]: A } | { [id: number]: A }>(
      object: T,
      predicate?: ?OPredicate<A, T>
    ): Object;
    omitBy<A, T>(object: void | null, predicate?: ?OPredicate<A, T>): {};
    pick(object?: ?Object, ...props: $ReadOnlyArray<string>): Object;
    pick(object?: ?Object, props: $ReadOnlyArray<string>): Object;
    pickBy<A, T: { [id: any]: A } | { [id: number]: A }>(
      object: T,
      predicate?: ?OPredicate<A, T>
    ): Object;
    pickBy<A, T>(object: void | null, predicate?: ?OPredicate<A, T>): {};
    result(
      object?: ?Object,
      path?: ?Path,
      defaultValue?: any
    ): any;
    set(object: Object, path?: ?Path, value: any): Object;
    set<T: void | null>(
      object: T,
      path?: ?Path,
      value?: ?any
    ): T;
    setWith<T>(
      object: T,
      path?: ?Path,
      value: any,
      customizer?: (nsValue: any, key: string, nsObject: T) => any
    ): Object;
    setWith<T: void | null>(
      object: T,
      path?: ?Path,
      value?: ?any,
      customizer?: ?(nsValue: any, key: string, nsObject: T) => any
    ): T;
    toPairs(object?: ?Object | Array<*>): Array<[string, any]>;
    toPairsIn(object?: ?Object): Array<[string, any]>;
    transform(
      collection: Object | $ReadOnlyArray<any>,
      iteratee?: ?OIteratee<*>,
      accumulator?: any
    ): any;
    transform(
      collection: void | null,
      iteratee?: ?OIteratee<*>,
      accumulator?: ?any
    ): {};
    unset(object: void | null, path?: ?Path): true;
    unset(object: Object, path?: ?Path): boolean;
    update(object: Object, path: Path, updater: Function): Object;
    update<T: void | null>(
      object: T,
      path?: ?Path,
      updater?: ?Function
    ): T;
    updateWith(
      object: Object,
      path?: ?Path,
      updater?: ?Function,
        | $ReadOnlyArray<Iteratee<T> | OIteratee<T>>
        | Iteratee<T>
        | OIteratee<T>
    ): (collection: $ReadOnlyArray<T> | { [id: any]: T }) => Array<T>;
    sortBy<T>(
      iteratees:
        | $ReadOnlyArray<Iteratee<T> | OIteratee<T>>
        | Iteratee<T>
        | OIteratee<T>,
      collection: $ReadOnlyArray<T> | { [id: any]: T }
    ): Array<T>;

    // Date
    now(): number;

    // Function
    after(fn: Function): (n: number) => Function;
    after(fn: Function, n: number): Function;
    ary(func: Function): Function;
    nAry(n: number): (func: Function) => Function;
    nAry(n: number, func: Function): Function;
    before(fn: Function): (n: number) => Function;
    before(fn: Function, n: number): Function;
    bind(func: Function): (thisArg: any) => Function;
    bind(func: Function, thisArg: any): Function;
    bindKey(obj: Object): (key: string) => Function;
    bindKey(obj: Object, key: string): Function;
    curry: Curry;
    curryN(arity: number): (func: Function) => Function;
    curryN(arity: number, func: Function): Function;
    curryRight(func: Function): Function;
    curryRightN(arity: number): (func: Function) => Function;
    curryRightN(arity: number, func: Function): Function;
    debounce(wait: number): <A, R>(func: (...A) => R) => (...A) => R;
    debounce<A, R>(wait: number, func: (...A) => R): (...A) => R;
    defer(func: (...any[]) => any): TimeoutID;
    delay(wait: number): (func: Function) => TimeoutID;
    delay(wait: number, func: Function): TimeoutID;
    flip(func: Function): Function;
    memoize<F: Function>(func: F): F;
    negate<A, R>(predicate: (...A) => R): (...A) => boolean;
    complement(predicate: Function): Function;
    once(func: Function): Function;
    overArgs(func: Function): (transforms: Array<Function>) => Function;
    overArgs(func: Function, transforms: Array<Function>): Function;
    useWith(func: Function): (transforms: Array<Function>) => Function;
    useWith(func: Function, transforms: Array<Function>): Function;
    partial(func: Function): (partials: any[]) => Function;
    partial(func: Function, partials: any[]): Function;
    partialRight(func: Function): (partials: Array<any>) => Function;
    partialRight(func: Function, partials: Array<any>): Function;
    rearg(indexes: Array<number>): (func: Function) => Function;
    rearg(indexes: Array<number>, func: Function): Function;
    rest(func: Function): Function;
    unapply(func: Function): Function;
    restFrom(start: number): (func: Function) => Function;
    restFrom(start: number, func: Function): Function;
    spread(func: Function): Function;
    apply(func: Function): Function;
    spreadFrom(start: number): (func: Function) => Function;
    spreadFrom(start: number, func: Function): Function;
    throttle<A, R>(wait: number): (func: (...A) => R) => (...A) => R;
    throttle<A, R>(wait: number, func: (...A) => R): (...A) => R;
    unary<T, R>(func: (T, ...any[]) => R): (T) => R;
    wrap(wrapper: Function): (value: any) => Function;
    wrap(wrapper: Function, value: any): Function;

    // Lang
    castArray(value: *): any[];
    clone<T>(value: T): T;
    cloneDeep<T>(value: T): T;
    cloneDeepWith<T, U>(
      customizer: (value: T, key: number | string, object: T, stack: any) => U
    ): (value: T) => U;
    cloneDeepWith<T, U>(
      customizer: (value: T, key: number | string, object: T, stack: any) => U,
      value: T
    ): U;
    cloneWith<T, U>(
      customizer: (value: T, key: number | string, object: T, stack: any) => U
    ): (value: T) => U;
    cloneWith<T, U>(
      customizer: (value: T, key: number | string, object: T, stack: any) => U,
      value: T
    ): U;
    conformsTo<T: { [key: string]: mixed }>(
      predicates: T & { [key: string]: (x: any) => boolean }
    ): (source: T) => boolean;
    conformsTo<T: { [key: string]: mixed }>(
      predicates: T & { [key: string]: (x: any) => boolean },
      source: T
    ): boolean;
    where<T: { [key: string]: mixed }>(
      predicates: T & { [key: string]: (x: any) => boolean }
    ): (source: T) => boolean;
    where<T: { [key: string]: mixed }>(
      predicates: T & { [key: string]: (x: any) => boolean },
      source: T
    ): boolean;
    conforms<T: { [key: string]: mixed }>(
      predicates: T & { [key: string]: (x: any) => boolean }
    ): (source: T) => boolean;
    conforms<T: { [key: string]: mixed }>(
      predicates: T & { [key: string]: (x: any) => boolean },
      source: T
    ): boolean;
    eq(value: any): (other: any) => boolean;
    eq(value: any, other: any): boolean;
    identical(value: any): (other: any) => boolean;
    identical(value: any, other: any): boolean;
    gt(value: any): (other: any) => boolean;
    gt(value: any, other: any): boolean;
    gte(value: any): (other: any) => boolean;
    gte(value: any, other: any): boolean;
    isArguments(value: any): boolean;
    isArray(value: any): boolean;
    isArrayBuffer(value: any): boolean;
    isArrayLike(value: any): boolean;
    isArrayLikeObject(value: any): boolean;
    isBoolean(value: any): boolean;
    isBuffer(value: any): boolean;
    isDate(value: any): boolean;
    isElement(value: any): boolean;
    isEmpty(value: any): boolean;
    isEqual(value: any): (other: any) => boolean;
    isEqual(value: any, other: any): boolean;
    equals(value: any): (other: any) => boolean;
    equals(value: any, other: any): boolean;
    isEqualWith<T, U>(
      customizer: (
        objValue: any,
        otherValue: any,
        key: number | string,
        object: T,
        other: U,
        stack: any
      ) => boolean | void
    ): ((value: T) => (other: U) => boolean) &
      ((value: T, other: U) => boolean);
    isEqualWith<T, U>(
      customizer: (
        objValue: any,
        otherValue: any,
        key: number | string,
        object: T,
        other: U,
        stack: any
      ) => boolean | void,
      value: T
    ): (other: U) => boolean;
    isEqualWith<T, U>(
      customizer: (
        objValue: any,
        otherValue: any,
        key: number | string,
        object: T,
        other: U,
        stack: any
      ) => boolean | void,
      value: T,
      other: U
    ): boolean;
    isError(value: any): boolean;
    isFinite(value: any): boolean;
    isFunction(value: any): boolean;
    isInteger(value: any): boolean;
    isLength(value: any): boolean;
    isMap(value: any): boolean;
    isMatch(source: Object): (object: Object) => boolean;
    isMatch(source: Object, object: Object): boolean;
    whereEq(source: Object): (object: Object) => boolean;
    whereEq(source: Object, object: Object): boolean;
    isMatchWith<T: Object, U: Object>(
      customizer: (
        objValue: any,
        srcValue: any,
        key: number | string,
        object: T,
        source: U
      ) => boolean | void
    ): ((source: U) => (object: T) => boolean) &
      ((source: U, object: T) => boolean);
    isMatchWith<T: Object, U: Object>(
      customizer: (
        objValue: any,
        srcValue: any,
        key: number | string,
        object: T,
        source: U
      ) => boolean | void,
      source: U
    ): (object: T) => boolean;
    isMatchWith<T: Object, U: Object>(
      customizer: (
        objValue: any,
        srcValue: any,
        key: number | string,
        object: T,
        source: U
      ) => boolean | void,
      source: U,
      object: T
    ): boolean;
    isNaN(value: any): boolean;
    isNative(value: any): boolean;
    isNil(value: any): boolean;
    isNull(value: any): boolean;
    isNumber(value: any): boolean;
    isObject(value: any): boolean;
    isObjectLike(value: any): boolean;
    isPlainObject(value: any): boolean;
    isRegExp(value: any): boolean;
    isSafeInteger(value: any): boolean;
    isSet(value: any): boolean;
    isString(value: string): true;
    isString(value: any): false;
    isSymbol(value: any): boolean;
    isTypedArray(value: any): boolean;
    isUndefined(value: any): boolean;
    isWeakMap(value: any): boolean;
    isWeakSet(value: any): boolean;
    lt(value: any): (other: any) => boolean;
    ): ((upper: number) => (number: number) => number) &
      ((upper: number, number: number) => number);
    clamp(lower: number, upper: number): (number: number) => number;
    clamp(lower: number, upper: number, number: number): number;
    inRange(
      start: number
    ): ((end: number) => (number: number) => boolean) &
      ((end: number, number: number) => boolean);
    inRange(start: number, end: number): (number: number) => boolean;
    inRange(start: number, end: number, number: number): boolean;
    random(lower: number): (upper: number) => number;
    random(lower: number, upper: number): number;

    // Object
    assign(object: Object): (source: Object) => Object;
    assign(object: Object, source: Object): Object;
    assignAll(objects: Array<Object>): Object;
    assignInAll(objects: Array<Object>): Object;
    extendAll(objects: Array<Object>): Object;
    assignIn<A, B>(a: A): (b: B) => A & B;
    assignIn<A, B>(a: A, b: B): A & B;
    assignInWith<T: Object, A: Object>(
      customizer: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A
      ) => any | void
    ): ((object: T) => (s1: A) => Object) & ((object: T, s1: A) => Object);
    assignInWith<T: Object, A: Object>(
      customizer: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A
      ) => any | void,
      object: T
    ): (s1: A) => Object;
    assignInWith<T: Object, A: Object>(
      customizer: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A
      ) => any | void,
      object: T,
      s1: A
    ): Object;
    assignWith<T: Object, A: Object>(
      customizer: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A
      ) => any | void
    ): ((object: T) => (s1: A) => Object) & ((object: T, s1: A) => Object);
    assignWith<T: Object, A: Object>(
      customizer: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A
      ) => any | void,
      object: T
    ): (s1: A) => Object;
    assignWith<T: Object, A: Object>(
      customizer: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A
      ) => any | void,
      object: T,
      s1: A
    ): Object;
    assignInAllWith(
      customizer: (
        objValue: any,
        srcValue: any,
        key: string,
        object: Object,
        source: Object
      ) => any | void
    ): (objects: Array<Object>) => Object;
    assignInAllWith(
      customizer: (
        objValue: any,
        srcValue: any,
        key: string,
        object: Object,
        source: Object
      ) => any | void,
      objects: Array<Object>
    ): Object;
    extendAllWith(
      customizer: (
        objValue: any,
        srcValue: any,
        key: string,
        object: Object,
        source: Object
      ) => any | void
    ): (objects: Array<Object>) => Object;
    extendAllWith(
      customizer: (
        objValue: any,
        srcValue: any,
        key: string,
        object: Object,
        source: Object
      ) => any | void,
      objects: Array<Object>
    ): Object;
    assignAllWith(
      customizer: (
        objValue: any,
        srcValue: any,
        key: string,
        object: Object,
        source: Object
      ) => any | void
    ): (objects: Array<Object>) => Object;
    assignAllWith(
      customizer: (
        objValue: any,
        srcValue: any,
        key: string,
        object: Object,
        source: Object
      ) => any | void,
      objects: Array<Object>
    ): Object;
    at(paths: Array<string>): (object: Object) => Array<any>;
    at(paths: Array<string>, object: Object): Array<any>;
    props(paths: Array<string>): (object: Object) => Array<any>;
    props(paths: Array<string>, object: Object): Array<any>;
    paths(paths: Array<string>): (object: Object) => Array<any>;
    paths(paths: Array<string>, object: Object): Array<any>;
    create<T>(prototype: T): T;
    defaults(source: Object): (object: Object) => Object;
    defaults(source: Object, object: Object): Object;
    defaultsAll(objects: Array<Object>): Object;
    defaultsDeep(source: Object): (object: Object) => Object;
    defaultsDeep(source: Object, object: Object): Object;
    defaultsDeepAll(objects: Array<Object>): Object;
    // alias for _.toPairs
    entries(object: Object): Array<[string, any]>;
    // alias for _.toPairsIn
    entriesIn(object: Object): Array<[string, any]>;
    // alias for _.assignIn
    extend<A, B>(a: A): (b: B) => A & B;
    extend<A, B>(a: A, b: B): A & B;
    // alias for _.assignInWith
    extendWith<T: Object, A: Object>(
      customizer: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A
      ) => any | void
    ): ((object: T) => (s1: A) => Object) & ((object: T, s1: A) => Object);
    extendWith<T: Object, A: Object>(
      customizer: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A
      ) => any | void,
      object: T
    ): (s1: A) => Object;
    extendWith<T: Object, A: Object>(
      customizer: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A
      ) => any | void,
      object: T,
      s1: A
    ): Object;
    findKey<A, T: { [id: any]: A }>(
      predicate: OPredicate<A>
    ): (object: T) => string | void;
    findKey<A, T: { [id: any]: A }>(
      predicate: OPredicate<A>,
      object: T
    ): string | void;
    findLastKey<A, T: { [id: any]: A }>(
      predicate: OPredicate<A>
    ): (object: T) => string | void;
    findLastKey<A, T: { [id: any]: A }>(
      predicate: OPredicate<A>,
      object: T
    ): string | void;
    forIn(iteratee: OIteratee<*>): (object: Object) => Object;
    forIn(iteratee: OIteratee<*>, object: Object): Object;
    forInRight(iteratee: OIteratee<*>): (object: Object) => Object;
    forInRight(iteratee: OIteratee<*>, object: Object): Object;
    forOwn(iteratee: OIteratee<*>): (object: Object) => Object;
    ): ((object: Object) => (args: Array<any>) => any) &
      ((object: Object, args: Array<any>) => any);
    invokeArgs(
      path: Path,
      object: Object
    ): (args: Array<any>) => any;
    invokeArgs(
      path: Path,
      object: Object,
      args: Array<any>
    ): any;
    keys<K>(object: { [key: K]: any }): Array<K>;
    keys(object: Object): Array<string>;
    keysIn(object: Object): Array<string>;
    mapKeys(iteratee: OIteratee<*>): (object: Object) => Object;
    mapKeys(iteratee: OIteratee<*>, object: Object): Object;
    mapValues(iteratee: OIteratee<*>): (object: Object) => Object;
    mapValues(iteratee: OIteratee<*>, object: Object): Object;
    merge(object: Object): (source: Object) => Object;
    merge(object: Object, source: Object): Object;
    mergeAll(objects: Array<Object>): Object;
    mergeWith<T: Object, A: Object, B: Object>(
      customizer: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A | B
      ) => any | void
    ): ((object: T) => (s1: A) => Object) & ((object: T, s1: A) => Object);
    mergeWith<T: Object, A: Object, B: Object>(
      customizer: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A | B
      ) => any | void,
      object: T
    ): (s1: A) => Object;
    mergeWith<T: Object, A: Object, B: Object>(
      customizer: (
        objValue: any,
        srcValue: any,
        key: string,
        object: T,
        source: A | B
      ) => any | void,
      object: T,
      s1: A
    ): Object;
    mergeAllWith(
      customizer: (
        objValue: any,
        srcValue: any,
        key: string,
        object: Object,
        source: Object
      ) => any | void
    ): (objects: Array<Object>) => Object;
    mergeAllWith(
      customizer: (
        objValue: any,
        srcValue: any,
        key: string,
        object: Object,
        source: Object
      ) => any | void,
      objects: Array<Object>
    ): Object;
    omit(props: Array<string>): (object: Object) => Object;
    omit(props: Array<string>, object: Object): Object;
    omitAll(props: Array<string>): (object: Object) => Object;
    omitAll(props: Array<string>, object: Object): Object;
    omitBy<A, T: { [id: any]: A }>(
      predicate: OPredicate<A>
    ): (object: T) => Object;
    omitBy<A, T: { [id: any]: A }>(predicate: OPredicate<A>, object: T): Object;
    pick(...props: Array<string | {}>): Object;
    pick(props: $ReadOnlyArray<string>, object: Object): Object;
    pick(...props: Array<string>): (object: Object) => Object;
    pick(props: $ReadOnlyArray<string>): (object: Object) => Object;
    pickAll(props: Array<string>): (object: Object) => Object;
    pickAll(props: Array<string>, object: Object): Object;
    pickBy<A, T: { [id: any]: A }>(
      predicate: OPredicate<A>
    ): (object: T) => Object;
    pickBy<A, T: { [id: any]: A }>(predicate: OPredicate<A>, object: T): Object;
    result(path: Path): (object: Object) => any;
    result(path: Path, object: Object): any;
    set(
      path: Path
    ): ((value: any) => (object: Object) => Object) &
      ((value: any, object: Object) => Object);
    set(path: Path, value: any): (object: Object) => Object;
    set(path: Path, value: any, object: Object): Object;
    assoc(
      path: Path
    ): ((value: any) => (object: Object) => Object) &
      ((value: any, object: Object) => Object);
    assoc(path: Path, value: any): (object: Object) => Object;
    assoc(path: Path, value: any, object: Object): Object;
    assocPath(
      path: Path
    ): ((value: any) => (object: Object) => Object) &
      ((value: any, object: Object) => Object);
    assocPath(
      path: Path,
      value: any
    ): (object: Object) => Object;
    assocPath(path: Path, value: any, object: Object): Object;
    setWith<T>(
      customizer: (nsValue: any, key: string, nsObject: T) => any
    ): ((
      path: Path
    ) => ((value: any) => (object: T) => Object) &
      ((value: any, object: T) => Object)) &
      ((path: Path, value: any) => (object: T) => Object) &
      ((path: Path, value: any, object: T) => Object);
    setWith<T>(
      customizer: (nsValue: any, key: string, nsObject: T) => any,
      path: Path
    ): ((value: any) => (object: T) => Object) &
      ((value: any, object: T) => Object);
    setWith<T>(
      customizer: (nsValue: any, key: string, nsObject: T) => any,
      path: Path,
      value: any
    ): (object: T) => Object;
    setWith<T>(
      customizer: (nsValue: any, key: string, nsObject: T) => any,
      path: Path,
      value: any,
      object: T
    ): Object;
    toPairs(object: Object | Array<*>): Array<[string, any]>;
    toPairsIn(object: Object): Array<[string, any]>;
    transform(
      iteratee: OIteratee<*>
    ): ((
      accumulator: any
    ) => (collection: Object | $ReadOnlyArray<any>) => any) &
      ((accumulator: any, collection: Object | $ReadOnlyArray<any>) => any);
    transform(
      iteratee: OIteratee<*>,
      accumulator: any
    ): (collection: Object | $ReadOnlyArray<any>) => any;
    transform(
      iteratee: OIteratee<*>,
      accumulator: any,
      collection: Object | $ReadOnlyArray<any>
    ): any;
    unset(path: Path): (object: Object) => Object;
    unset(path: Path, object: Object): Object;
    dissoc(path: Path): (object: Object) => Object;
    dissoc(path: Path, object: Object): Object;