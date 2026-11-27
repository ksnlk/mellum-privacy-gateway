                      + "redis.call('hset', KEYS[1], ARGV[2], ARGV[3]); "
                      + "local msg = struct.pack('Lc0Lc0', string.len(ARGV[2]), ARGV[2], string.len(ARGV[3]), ARGV[3]); "
                      + "redis.call('publish', KEYS[3], msg); "                  
                      + "return 1;"
                  + "end; ",
             Arrays.<Object>asList(getRawName(), getTimeoutSetName(), getCreatedChannelName()),
             creationTimeout, encodeMapKey(key), encodeMapValue(value));
    }

    
    private String getLockName(Object key) {
        ByteBuf keyState = encodeMapKey(key);
        try {
            return "{" + getRawName() + "}:" + Hash.hash128toBase64(keyState) + ":key";
        } finally {
            keyState.release();
        }
    }

    @Override
    public Map<K, V> getAll(Set<? extends K> keys) {
        checkNotClosed();
        if (keys == null) {
            throw new NullPointerException();
        }
        for (K key : keys) {
            checkKey(key);
        }

        if (!atomicExecution && !config.isReadThrough()) {
            long startTime = currentNanoTime();
            boolean exists = false;
            for (K key : keys) {
                if (containsKey(key)) {
                    exists = true;
                    break;
                }
            }
            if (!exists) {
                cacheManager.getStatBean(this).addGetTime(currentNanoTime() - startTime);
                return Collections.emptyMap();
            }
        }
        
        RFuture<Map<K, V>> result = getAllAsync(keys);

        result.syncUninterruptibly();
        return result.getNow();
    }

    @Override
    public RFuture<Map<K, V>> getAllAsync(Set<? extends K> keys) {
        checkNotClosed();
        long startTime = currentNanoTime();
        Long accessTimeout = getAccessTimeout();
        
        List<Object> args = new ArrayList<Object>(keys.size() + 2);
        args.add(accessTimeout);
        checkKey(key);

        RFuture<Boolean> future = commandExecutor.evalReadAsync(getRawName(), codec, RedisCommands.EVAL_BOOLEAN,
                  "if redis.call('hexists', KEYS[1], ARGV[2]) == 0 then "
                    + "return 0;"
                + "end;"

                + "local expireDateScore = redis.call('zscore', KEYS[2], ARGV[2]); "
                + "if expireDateScore ~= false and tonumber(expireDateScore) <= tonumber(ARGV[1]) then "
                    + "return 0; "
                + "end; "
                + "return 1;",
             Arrays.<Object>asList(getRawName(), getTimeoutSetName()),
             System.currentTimeMillis(), encodeMapKey(key));
        return future;
    }

    @Override
    @SuppressWarnings("NestedTryDepth")
    public void loadAll(final Set<? extends K> keys, final boolean replaceExistingValues, final CompletionListener completionListener) {
        checkNotClosed();
        if (keys == null) {
            throw new NullPointerException();
        }
        
        for (K key : keys) {
            checkKey(key);
        }

        if (cacheLoader == null) {
            if (completionListener != null) {
                completionListener.onCompletion();
            }
            return;
        }

        commandExecutor.getConnectionManager().getExecutor().execute(new Runnable() {
            @Override
            public void run() {
                for (K key : keys) {
                    try {
                        if (!containsKey(key) || replaceExistingValues) {
                            RLock lock = getLockedLock(key);
                            try {
                                if (!containsKey(key)|| replaceExistingValues) {
                                    V value;
                                    try {
                                        value = cacheLoader.load(key);
                                    } catch (Exception ex) {
                                        throw new CacheLoaderException(ex);
                                    }
                                    if (value != null) {
                                        if (atomicExecution) {
                                            putValue(key, value);
                                        } else {
                                            putValueLocked(key, value);
                                        }
                                    }
                                }
                            } finally {
                                lock.unlock();
                            }
                        }
                    } catch (Exception e) {
                        if (completionListener != null) {
                    
                    if (r) {
                        cacheManager.getStatBean(this).addHits(1);
                        cacheManager.getStatBean(this).addRemovals(1);
                    } else {
                        cacheManager.getStatBean(this).addMisses(1);
                    }
                    cacheManager.getStatBean(this).addRemoveTime(currentNanoTime() - startTime);
                    result.trySuccess(r);
                });
            } finally {
                lock.unlock();
            }
        }
        return result;
    }

    RFuture<Map<K, V>> getAndRemoveValues(Collection<K> keys) {
        double syncId = ThreadLocalRandom.current().nextDouble();
        RPromise<Map<K, V>> result = new RedissonPromise<>();
        
        List<Object> params = new ArrayList<>();
        params.add(System.currentTimeMillis());
        params.add(syncId);
        
        for (K key : keys) {
            params.add(encodeMapKey(key));
        }
        
        RFuture<List<Object>> future = commandExecutor.evalWriteAsync(getRawName(), codec, RedisCommands.EVAL_MAP_VALUE_LIST,
                "local syncs = 0; "
              + "local values = {}; "
              + "local result = {}; "
              + "local nulls = {}; "
              
              + "for i = 3, #ARGV, 1 do "
                  + "local value = redis.call('hget', KEYS[1], ARGV[i]); "
                  + "if value == false then "
                      + "table.insert(nulls, i-3); "
                  + "else "
                      + "local expireDateScore = redis.call('zscore', KEYS[2], ARGV[i]); "
                      + "if expireDateScore ~= false and tonumber(expireDateScore) <= tonumber(ARGV[1]) then "
                          + "table.insert(nulls, i-3); "
                      + "else "
                          + "redis.call('hdel', KEYS[1], ARGV[i]); "
                          + "redis.call('zrem', KEYS[2], ARGV[i]); "
                          + "local msg = struct.pack('Lc0Lc0', string.len(ARGV[i]), ARGV[i], string.len(tostring(value)), tostring(value)); "
                          + "redis.call('publish', KEYS[3], msg); "
                          + "local syncMsg = struct.pack('Lc0Lc0d', string.len(ARGV[i]), ARGV[i], string.len(tostring(value)), tostring(value), ARGV[2]); "
                          + "syncs = syncs + redis.call('publish', KEYS[4], syncMsg); "
                          + "table.insert(values, value); "
                return result;
            }
            
            long syncs = (long) r.get(0);
            if (syncs > 0) {
                RSemaphore semaphore = redisson.getSemaphore(getSyncName(syncId));
                try {
                    semaphore.acquire((int) syncs);
                    semaphore.delete();
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                }
            }
            
            getAndRemoveValuesResult(keys, result, r, nullsAmount);
        }
        
        return result;
    }

    private void getAndRemoveValuesResult(Collection<K> keys, RPromise<Map<K, V>> result, List<Object> r,
            long nullsAmount) {
        HashSet<Long> nullIndexes = new HashSet<>((List<Long>) (Object) r.subList(2, (int) nullsAmount + 2));
        Map<K, V> res = new HashMap<>();
        long i = 0;
        for (K key : keys) {
            if (nullIndexes.contains(i)) {
                continue;
            }
            V value = (V) r.get((int) (i+nullsAmount+2));
            res.put(key, value);
            i++;
        }
        result.trySuccess(res);
    }
    
    RFuture<V> getAndRemoveValue(K key) {
        double syncId = ThreadLocalRandom.current().nextDouble();
        RPromise<V> result = new RedissonPromise<>();
        RFuture<List<Object>> future = commandExecutor.evalWriteAsync(getRawName(), codec, RedisCommands.EVAL_MAP_VALUE_LIST,
                "local value = redis.call('hget', KEYS[1], ARGV[2]); "
              + "if value == false then "
                  + "return {nil}; "
              + "end; "
                  
              + "local expireDateScore = redis.call('zscore', KEYS[2], ARGV[2]); "
              + "if expireDateScore ~= false and tonumber(expireDateScore) <= tonumber(ARGV[1]) then "
                  + "return {nil}; "
              + "end; "

              + "redis.call('hdel', KEYS[1], ARGV[2]); "
                    cacheManager.getStatBean(this).addHits(1);
                    cacheManager.getStatBean(this).addPuts(1);
                } else {
                    cacheManager.getStatBean(this).addMisses(1);
                }
                cacheManager.getStatBean(this).addPutTime(currentNanoTime() - startTime);
                cacheManager.getStatBean(this).addGetTime(currentNanoTime() - startTime);
                result.trySuccess(r);
            });
        } finally {
            lock.unlock();
        }
        return result;
    }

    @Override
    public void removeAll(Set<? extends K> keys) {
        RFuture<Void> future = removeAllAsync(keys);
        future.syncUninterruptibly();
    }
    
    @Override
    public RFuture<Void> removeAllAsync(Set<? extends K> keys) {
        checkNotClosed();
        
        for (K key : keys) {
            checkKey(key);
        }
        
        long startTime = currentNanoTime();
        RPromise<Void> result = new RedissonPromise<>();
        if (config.isWriteThrough()) {
            RFuture<Map<K, V>> future = getAndRemoveValues((Set<K>) keys);
            future.onComplete((r, ex) -> {
                if (ex != null) {
                    result.tryFailure(ex);
                    return;
                }
                
                try {
                    cacheWriter.deleteAll(r.keySet());
                } catch (Exception e) {
                    putAllValues(r);
                    handleException(result, e);
                    return;
                }
                cacheManager.getStatBean(this).addRemovals(r.size());
                cacheManager.getStatBean(this).addRemoveTime(currentNanoTime() - startTime);
                result.trySuccess(null);
            });
        } else {
                    || entry.getAction() == Action.UPDATED) {
                put(key, entry.value());
            }
            if (entry.getAction() == Action.DELETED) {
                remove(key);
            }
            return result;
        } catch (EntryProcessorException e) {
            throw e;
        } catch (Exception e) {
            throw new EntryProcessorException(e);
        } finally {
            lock.unlock();
        }
    }

    @Override
    public <T> Map<K, EntryProcessorResult<T>> invokeAll(Set<? extends K> keys, EntryProcessor<K, V, T> entryProcessor,
            Object... arguments) {
        checkNotClosed();
        if (entryProcessor == null) {
            throw new NullPointerException();
        }

        Map<K, EntryProcessorResult<T>> results = new HashMap<K, EntryProcessorResult<T>>();
        for (K key : keys) {
            try {
                final T result = invoke(key, entryProcessor, arguments);
                if (result != null) {
                    results.put(key, new EntryProcessorResult<T>() {
                        @Override
                        public T get() throws EntryProcessorException {
                            return result;
                        }
                    });
                }
            } catch (final EntryProcessorException e) {
                results.put(key, new EntryProcessorResult<T>() {
                    @Override
                    public T get() throws EntryProcessorException {
                        throw e;
                    }
                });
            }
        }

        return results;
    }

    @Override
    public CacheManager getCacheManager() {