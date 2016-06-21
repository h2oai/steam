package python

import (
	"github.com/h2oai/steamY/tools/piping/parser"
	"github.com/serenize/snaker"
	"strings"
)

func genPrelude() string {
	return `
# ----------------------------------
# --- Generated with go:generate ---
# ---        DO NOT EDIT         ---
# ----------------------------------

import httplib
import base64
import string
import json
import sys
from collections import namedtuple

class RPCError(Exception):
	def __init__(self, value):
		self.value = value
	def __str__(self):
		return repr(self.value)

class HTTPConnection:
	def __init__(self, host, port, username, password):
		self.host = host
		self.port = port
		self.username = username
		self.password = password
		self.uid = 0

	def call(self, method, params):
		self.uid = self.uid + 1
		request = {
			'id': self.uid,
			'method': 'web.' + method,
			'params': [params]
		}
		payload = json.dumps(request)

		ws = httplib.HTTP(self.host, self.port)
		ws.putrequest("POST", '/web')

		ws.putheader("Host", self.host)
		ws.putheader("User-Agent", "Steam Python Client")
		ws.putheader("Content-type", "application/json; charset=\"UTF-8\"")
		ws.putheader("Content-length", "%d" % len(payload))
		auth = base64.encodestring('%s:%s' % (self.username, self.password)).replace('\n', '')
		ws.putheader("Authorization", "Basic %s" % auth)
		ws.endheaders()

		ws.send(payload)

		code, status, header = ws.getreply()
		reply = ws.getfile().read()

		# print 'code:', code
		# print 'status:', status
		# print 'reply:', reply

		if code != 200:
			raise RPCError(reply)

		response = json.loads(reply)

		if response['error'] is None:
			return response['result']
		else:
			raise RPCError(response['error'])

class View(object):
	def __init__(self, d):
		self.__dict__ = d
	def __str__(self):
		return json.dumps(self.__dict__)

class RPCClient:
	def __init__(self, connection):
		self.connection = connection
	`
}

func genReturn(param *parser.Param) string {
	n := "response['" + snaker.CamelToSnake(param.Name) + "']"
	t := param.Type
	isArray := false

	if strings.HasPrefix(t, "[]") {
		t = t[2:]
		isArray = true
	}
	if strings.HasPrefix(t, "*") {
		t = t[1:]
	}

	switch t {
	case "bool", "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "float32", "float64", "int", "uint":
		return n
	default:
		if isArray {
			return "[View(o) for o in " + n + "]"
		}
		return "View(" + n + ")"
	}
}

func genParam(p *parser.Param) string {
	return snaker.CamelToSnake(p.Name)
}

func genClientStub(f *parser.Func) string {
	c := "\n\n\tdef " + snaker.CamelToSnake(f.Name) + "("
	args := make([]string, len(f.Params)+1)
	args[0] = "self"
	for i, p := range f.Params {
		args[i+1] = genParam(p)
	}
	c += strings.Join(args, ", ")
	c += "):\n"

	c += "\t\t\"\"\"\n"
	c += "\t\tReturns "
	if f.Return != nil {
		c += f.Return.Name + " (" + f.Return.Type + ")\n"
	} else {
		c += "None\n"
	}
	c += "\t\t\"\"\"\n"

	c += "\t\trequest = {\n"
	for i, p := range f.Params {
		c += "\t\t\t'" + snaker.CamelToSnake(p.Name) + "': " + snaker.CamelToSnake(p.Name)
		if i < len(f.Params)-1 {
			c += ","
		}
		c += "\n"
	}
	c += "\t\t}\n"

	call := "self.connection.call(\"" + f.Name + "\", request)\n"
	if f.Return != nil {
		c += "\t\tresponse = " + call
		c += "\t\treturn " + genReturn(f.Return)
	} else {
		c += "\t\t" + call
	}

	return c
}

func Generate(i *parser.Interface) string {
	c := genPrelude()
	if len(i.Funcs) > 0 {
		for _, f := range i.Funcs {
			c += genClientStub(f)
		}
	}

	return c
}
