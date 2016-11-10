import sys
import time
import testutil as tu
import subprocess as sp
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

def setup():
	tu.cliLogin('superuser', 'superuser')
	role = tu.createRole('wicked', 'cool', [])
	return role

def writeumsTest(driver, role):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		tu.goUsers(driver)	
		tu.goRoles(driver)
		wait.until(lambda x: x.find_element_by_xpath("//input[@data-roleid='{0}']".format(role)))
	except Exception as e:
		print e
		print "Failed to view role on roles page"
		return False
	try:
		boxes = driver.find_elements_by_xpath("//input[@data-roleid='{0}']".format(role))
		for b in boxes:
			b.click()
		driver.find_element_by_xpath("//div[@class='button-primary']").click()
		changes = driver.find_elements_by_xpath("//span[@class='button-primary']")
		for c in changes:
			c.click()
		driver.find_element_by_xpath("//div[@class='button-primary' and text()='Save Changes']").click()
		time.sleep(19)
		perm = tu.getRolePermissions(role)
		if len(perm) != 25:
			print "Some permissions failed to update"
			return False
		driver.refresh()
		tu.goRoles(driver)
		wait.until(lambda x: x.find_element_by_xpath("//input[@data-roleid='{0}']".format(role)))
		boxes = driver.find_elements_by_xpath("//input[@data-roleid='{0}']".format(role))
		for b in boxes:
			b.click()
		driver.find_element_by_xpath("//div[@class='button-primary']").click()
		changes = driver.find_elements_by_xpath("//span[@class='button-primary']")
		for c in changes:
			c.click()
		driver.find_element_by_xpath("//div[@class='button-primary' and text()='Save Changes']").click()
		time.sleep(19)
		perm = tu.getRolePermissions(role)
		if len(perm) > 3:
			print "Failed to remove permissions"
			return False
	except Exception as e:
		print e
		print "Failed to modify role"
		return False
	return True


def main():
	failcount = 0
	d = tu.newtest()
	r = setup()

	if not writeumsTest(d, r):
		failcount += 1
	
	tu.endtest(d)
	sys.exit(failcount)	


if __name__ == '__main__':
	main()
	
