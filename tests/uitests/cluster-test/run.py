import sys
import time
import testutil as tu
import subprocess as sp
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.chrome.options import Options

def connectTest(driver):
	tu.goClusters(driver)
	return tu.clusterExists(driver, "steamtest")	

def deleteClusterTest(driver):
	if not tu.goClusters(driver):
		return False
	try:
		tu.addCluster(driver, "localhost", "54321", "pjr")		
		tu.deleteCluster(driver, "pjr")
	except Exception as e:
		print e
		print "Failed to delete cluster"
		return False
	return True

def dataframeTest(driver):
	try:
		tu.goHome(driver)
		tu.newProject(driver)
		tu.selectCluster(driver, "steamtest")
	except:
		print "Failed to select cluster"
		return False	
	try:	
		wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
		wait.until(lambda x: x.find_element_by_xpath("//option[@value='bank_full.hex']"))
		wait.until(lambda x: x.find_element_by_xpath("//option[@value='arrhythmia.hex']"))
		
	except:
		print "Dataframe select doesn't include all dataframes related to cluster"
		return False	
	return True


def main():
	failcount = 0
	d = tu.newtest()
	d.get("http://admin:admin012@localhost:9000")
	if not connectTest(d):
		failcount += 1
	if not deleteClusterTest(d):
		failcount += 1
	if not dataframeTest(d):
		failcount += 1
	d.quit()
	sys.exit(failcount)
	

if __name__ == '__main__':
	main()

