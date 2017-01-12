import sys
import time
import testutil as tu
import subprocess as sp
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.select import Select
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.chrome.options import Options

def launchClusterTest(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		tu.goClusters(driver)
		driver.find_element_by_xpath("//a[text()='Launch New Cluster']").click()
		wait.until(lambda x: x.find_element_by_xpath("//input[@name='name']"))
	except:
		print "Failed to access 'Launch New Cluster' page"
		return False
	try:
		tu.uploadEngine("/home/patrick/Downloads/h2o-3.10.2.1-hdp2.4.zip")
		time.sleep(5)
		driver.refresh()
		time.sleep(5)
		sel = Select(driver.find_element_by_xpath("//select"))
		if len(sel.options) == 1:
			return False
	except Exception as e:
		print "Failed to upload h2o engine"
		return False	
	try:
		tu.launchCluster(driver, "yarn", "1", "6", False, "1")
		time.sleep(30)
		tu.goClusters(driver)
		if not tu.clusterExists(driver, "yarn"):
			print "Cluster yarn not seen"
			return False
	except:
		print "Failed to configure/launch cluster"
		return False
	try:
		tu.deleteCluster(driver, "yarn")
	except Exception as e:
		print repr(e)
		print "Failed to kill yarn cluster"
		return False
	return True

def main():
	failcount = 0

	d = tu.newtest()
	tu.cliLogin("patrick", "superuser")
	if not launchClusterTest(d):
		failcount += 1

	tu.endtest(d)
	sys.exit(failcount)
	

if __name__ == '__main__':
	main()

