#!/usr/bin/env python
# -*- coding: UTF-8 -*-

"""
@Desc   ：timing attack Against Hash
"""
import hmac
import hashlib

# the old unlock code was j748347j-45m6-4152-5k31-l21be7f7
msg = "Test"

def sign(pre_key, imsg, alg):
    return hmac.new(pre_key, imsg, alg).digest()

def verify(msg, sig):
    return sig == sign(key, msg, hashlib.sha256) #bad
