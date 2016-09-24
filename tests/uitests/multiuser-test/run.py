import sys
import time
import testutil as tu
import subprocess as sp
import re
from selenium.webdriver.support.wait import WebDriverWait


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

"""
EntityType

 1	Role
 2	Workgroup
 3	Identity
 4	Engine
 5	Cluster
 6	Project
 7	Datasrc
 8	Dataset
 9	model
10	label
11	service

"""

def superShareTest():
	res = [True, True]
	wg = tu.createWorkgroup("supertest", "testin super share")
	d = tu.newtest()
	wait = WebDriverWait(d, timeout=5, poll_frequency=0.2)
	tu.newProject(d)
	tu.addCluster(d, "localhost", "54535", "steamtest")
	d.find_element_by_xpath("//div[@class='select-cluster']//button").click()
	tu.selectDataframe(d, "bank_full.hex")
	tu.selectModelCategory(d, "Regression")
	try:
		tu.selectModel(d, "regress")
		d.find_element_by_xpath("//div[@class='name-project']//input").send_keys("supertest")
		d.find_element_by_xpath("//button[text()='Create Project']").click()
		wait.until(lambda x: x.find_element_by_xpath("//div[@class='model-name' and text()='regress']"))
		tu.goProjectConfig(d)
		d.find_element_by_xpath("//button[text()='Create New Label']").click()
		wait.until(lambda x: x.find_element_by_xpath("//input[@name='name']").is_displayed())
		d.find_element_by_xpath("//input[@name='name']").send_keys("taggy")
		d.find_element_by_xpath("//textarea[@name='description']").send_keys("the taggiest tag")
		d.find_element_by_xpath("//button[text()='Save']").click()
		wait.until(lambda x: x.find_element_by_xpath("//div[@class='label-name' and text()='taggy']").is_displayed())
	except:
		print "Failed to setup super sharing test"
		return False
	finally:
		tu.endtest(d)
	
	tu.shareEntity(6, 1, wg, 'edit') 
	tu.shareEntity(9, 1, wg, 'edit')

	uid = tu.createIdentity("noperm", "noperm")
	tu.assignWorkgroup(uid, wg)
	d = tu.testAs("noperm", "noperm")
	try:
		#permissionless test
		tu.goProjects(d)
		if tu.viewProject(d, "supertest"):
			res[0] = False
			print "User with no permissions is able to view entities shared by superuser"
	except:
		res[0] = True
	finally:
		tu.endtest(d)

	uid = tu.createIdentity("yaperm", "yaperm")
	rid = tu.createRole("permissed", "for testing supershare", [9, 10, 11, 12, 17, 18, 19, 20])
	tu.assignWorkgroup(uid, wg)
	tu.assignRole(uid, rid)
	d = tu.testAs("yaperm", "yaperm")
	try:
		#permissable test
		tu.goProjects(d)
		time.sleep(2)
		if not tu.viewProject(d, "supertest"):
			res[0] = False
			print "User could not see project shared by superuser"
	except:
		res[1] = False
		print "User with appropriate permissions is not able to view entities shared by superuser"
	finally:
		tu.endtest(d)
	return res[0] and res[1]

def userShareTest():
	res = [True, True]
	uid = tu.createIdentity("setup", "setup")
	rid = tu.createRole("userole", "comfy af", [x for x in range(23) if x > 0])
	tu.assignRole(uid, rid)
	tu.cliLogin("setup", "setup")
	wg = tu.createWorkgroup("clusterwg", "for sharing the cluster")
	tu.assignWorkgroup(uid, wg)
	tu.cliLogin("superuser", "superuser")
	tu.shareEntity(5, 1, wg, 'edit')
	tu.cliLogin("setup", "setup")
	wg = tu.createWorkgroup("usertest", "testin user share")
	d = tu.testAs("setup", "setup")
	try:
		wait = WebDriverWait(d, timeout=5, poll_frequency=0.2)
		tu.newProject(d)
		time.sleep(2)
		d.find_element_by_xpath("//div[@class='select-cluster']//button").click()
		tu.selectDataframe(d, "bank_full.hex")
		tu.selectModelCategory(d, "Regression")
		tu.selectModel(d, "gradi")
		d.find_element_by_xpath("//div[@class='name-project']//input").send_keys("averagetest")
		d.find_element_by_xpath("//button[text()='Create Project']").click()
		wait.until(lambda x: x.find_element_by_xpath("//div[@class='model-name' and text()='gradi']"))
	except Exception as e:
		print "Failed to setup user share test"
		return False
	finally:
		tu.endtest(d)

	tu.shareEntity(6, 2, wg, 'edit') 
	tu.shareEntity(9, 2, wg, 'edit')

	uid = tu.createIdentity("permless", "permless")
	tu.cliLogin("superuser", "superuser")
	tu.assignWorkgroup(uid, wg)
	d = tu.testAs("permless", "permless")
	try:
		#permissionless test
		tu.goProjects(d)
		if tu.viewProject(d, "averagetest"):
			res[0] = False
			print "User with no permissions is able to view entities shared by non-superuser"
	except:
		res[0] = True
	finally:
		tu.endtest(d)

	uid = tu.createIdentity("useperm", "useperm")
	rid = tu.createRole("collmissed", "for testing usershare", [9, 10, 11, 12, 17, 18, 19, 20])
	tu.assignWorkgroup(uid, wg)
	tu.assignRole(uid, rid)
	d = tu.testAs("useperm", "useperm")
	try:
		#permissable test
		tu.goProjects(d)
		if not tu.viewProject(d, "averagetest"):
			res[0] = False
			print "User could not see project shared by non-superuser"
	except:
		res[1] = False
		print "User with appropriate permissions is not able to view entities shared by non-superuser"
	finally:
		tu.endtest(d)
	
	return res[0] and res[1]

def main():
	failcount = 0

	if not superShareTest():
		failcount += 1
	if not userShareTest():
		failcount += 1
	
	sys.exit(failcount)
	

if __name__ == '__main__':
	main()
