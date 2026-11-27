#!/usr/bin/env python
# -*- coding: UTF-8 -*-

"""
@Desc   ：preventing timing attack Against Hash
"""
import hmac
import hashlib

key = "n218380u-20g4-0968-1b97-p56dk0n9"
msg = "Test"

def sign(pre_key, imsg, alg):
    return hmac.new(pre_key, imsg, alg).digest()

def verify(msg, sig):
    return hmac.compare_digest(sig, sign(key, msg, hashlib.sha256)) #good
