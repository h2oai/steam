import sys
import time
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support.ui import Select
import testutil as tu

def allMetrics(driver):
	for i in range(len(driver.find_elements_by_xpath("(//div[@class='collapsible'])[2]//div[@class='details compare']"))):
		spans = driver.find_elements_by_xpath("((//div[@class='collapsible'])[2]//div[@class='details--value'])[{0}]/span"\
			.format(i + 1))
		if len(spans[0].text) < 3 or len(spans[1].text) < 3:
			name = driver.find_element_by_xpath("((//div[@class='collapsible'])[2]//div[@class='details--label'])[{0}]"\
				.format(i + 1))
			print "Metric " + name.text + " was missing in comparison"
			passed = False
	return True
		

def compRegTest(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		tu.createProject(driver, "cmpReg", "steamtest", "bank_full.hex", "Regression", ["regress", "gradi"])
	except Exception as e:
		print e
		print "Failed to setup regression comparison test"
		return False
	try:
		tu.viewModel(driver, "regress")
		tu.compareToModel(driver, "gradi")
	except:
		print "Failed to compare regress to gradi (GLM and GBM regression models)"
		return False
	
	return allMetrics(driver)

def compBinTest(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		tu.createProject(driver, "cmpBin", "steamtest", "bank_full.hex", "Binomial", ["second", "third"])
	except Exception as e:
		print e
		print "Failed to setup binomial comparison test"
		return False
	try:
		tu.viewModel(driver, "third")
		tu.compareToModel(driver, "second")
	except:
		print "Failed to compare third to second (GLM binomial models)"
		return False	

	return allMetrics(driver)

def compMultiTest(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		tu.createProject(driver, "cmpMulti", "steamtest", "bank_full.hex", "Multinomial", ["multinom", "valimon"])
	except Exception as e:
		print e
		print "Failed to setup multinomial comparison test"
		return False
	try:
		tu.viewModel(driver, "multinom")
		tu.compareToModel(driver, "valimon")
	except:
		print "Failed to compare multinom to valimon (GLM multinomial models)"
		return False	

	return allMetrics(driver)

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

