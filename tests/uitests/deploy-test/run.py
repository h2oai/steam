import sys
import time
import magic.testutil as tu
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait

def deployOneTest(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	tu.newProject(driver)
	tu.addCluster(driver, "localhost", "54535", "steamtest")
	driver.find_element_by_xpath("//div[@class='select-cluster']//button").click()
	tu.selectDataframe(driver, "bank_full.hex")
	tu.selectModelCategory(driver, "Regression")
	try:
		tu.selectModel(driver, "regress")
		driver.find_element_by_xpath("//div[@class='name-project']//input").send_keys("deptest")
		driver.find_element_by_xpath("//button[text()='Create Project']").click()
		wait.until(lambda x: x.find_element_by_xpath("//div[@class='model-name' and text()='regress']"))
	except:
		print "failed to set up test"
		return False
	try:
		tu.deployModel(driver, "regress")
	except:
		return False

def main():
	failcount = 0
	d = tu.newtest()
	
	if not deployOneTest(d):
		failcount += 1
	
	tu.endtest(d)
	sys.exit(failcount)

if __name__ == '__main__':
	main()

