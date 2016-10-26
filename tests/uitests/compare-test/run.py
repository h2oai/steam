import sys
import time
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support.ui import Select
import testutil as tu

def compRegTest(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		tu.createProject(driver, "cmpReg", "steamtest", "bank_full.hex", "Regression", ["regress", "linmiss"])
	except Exception as e:
		print e
		print "Failed to setup regression comparison test"
		return False
	try:
		tu.viewModel(driver, "regress")
		tu.compareToModel(driver, "linmiss")
	except:
		print "Failed to compare regress to linmiss"
		return False
	time.sleep(15)
	return True

def compBinTest(driver):
	
	return True

def compMultiTest(driver):
	
	return True

def main():
	failcount = 0
	d = tu.newtest()
	
	if not compRegTest(d):
		failcount += 1
	if not compBinTest(d):
		failcount += 1
	if not compMultiTest(d):
		failcount += 1

	tu.endtest(d)
	sys.exit(failcount)

if __name__ == '__main__':
	main()

