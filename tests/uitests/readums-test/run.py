import sys
import time
import testutil as tu
import subprocess as sp
import re
"""
Perm id		Permission		Index
============================================
 1			M Role			 8
 2			V Role			19
 3			M Workgroup		10
 4			V Workgroup		21
 5			M Identities	 4
 6			V Identities	15
 7			M engines		 3
 8			V engines		14
 9			M clusters		 0
10			V clusters		11
11			M Projects		 7
12			V Projects		18
13			M Datasrc		 2
14			V Datasrc		13
15			M Dataset		 1
16			V Dataset		12
17			M Model			 6
18			V Model			17
19			M label			 5
20			V Label			16
21			M service		 9
22			V service		20
"""

_steampath = "./steam"

def createRole(role, desc, perm):
	ret = sp.check_output("{0} create role --name {1} --description {2}"\
		.format(_steampath, role, desc), shell=True)
	i = int(re.search(r'\d+', ret).group())
	for p in perm:
		sp.Popen("{0} link role --with-permission --role-id={1} --permission-id={2} > /dev/null"\
			.format(_steampath, i, p), shell=True).communicate()
	return i

def createIdentity(name, pw):
	ret = sp.check_output("{0} create identity --name={1} --password={2}"\
		.format(_steampath, name, pw), shell=True)
	return int(re.search(r'\d+', ret).group())

def assignRole(iden, role):
	sp.Popen("{0} link identity --with-role --identity-id={1} --role-id={2}"\
		.format(_steampath, iden, role), shell=True).communicate()


def readRolesTest():
	role = createRole("hello", "hey", [1, 2, 4, 5, 6, 9, 10, 17, 18])
	inds = [8, 19, 21, 4, 15, 0, 11, 6, 17]
	iden = createIdentity("billy", "bob")
	assignRole(iden, role)
	d = tu.newtest()
	try:
		tu.goUsers(d)	
		tu.goRoles(d)
		time.sleep(1)
		checked = d.find_elements_by_xpath("//input[@type='checkbox' and @value='on']")
		for i in inds:
			box = d.find_elements_by_xpath("//input[@type='checkbox']")[(i * 2) + 1]
			if not box in checked:
				print "Permission {0} isn't checked".format(i)
				return False
	except:
		print "Failed to access users/roles page"
		return False
	finally:
		tu.endtest(d)
	
	return True


def main():
	failcount = 0

	global _steampath	
	if sys.platform.startswith("linux"):
		_steampath = "./steam-develop-linux-amd64/steam"
	elif sys.platform == "darwin":
		_steampath = "./steam-develop-darwin-amd64/steam"
	else:
		print "unsupported testing platform"
		sys.exit(1)

	if not readRolesTest():
		failcount += 1
	
	sys.exit(failcount)

if __name__ == '__main__':
	main()

