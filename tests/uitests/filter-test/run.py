import sys
import time
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support.ui import Select
import testutil as tu

def nameFilterTest(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		tu.newProject(driver)
		tu.selectCluster(driver, "steamtest")
		tu.selectDataframe(driver, "bank_full.hex")
		tu.selectModelCategory(driver, "Regression")
		models = ["regress", "gradi", "missin", "linmiss"]
		for mod in models:
			tu.selectModel(driver, mod)
		driver.find_element_by_xpath("//div[@class='name-project']//input").send_keys("namefilt")
		driver.find_element_by_xpath("//button[text()='Create Project']").click()
		for mod in models:
			wait.until(lambda x: x.find_element_by_xpath("//div[@class='model-name' and text()='{0}']".format(mod)))
	except:
		print "Failed to setup name filter test"
		return False
	try:
		f = driver.find_element_by_xpath("//div[@class='filter']/input")
		f.send_keys("linmiss")
		wait.until(lambda x: len(x.find_elements_by_xpath("//div[@class='model-name']")) == 1)
		if not driver.find_element_by_xpath("//div[@class='model-name']").text == "linmiss":
			print "Name filter did not apply correctly"
			return False
		driver.refresh()
		wait.until(lambda x: len(x.find_elements_by_xpath("//div[@class='model-name']")) == 4)
		f = driver.find_element_by_xpath("//div[@class='filter']/input")
		f.send_keys("r")
		wait.until(lambda x: len(x.find_elements_by_xpath("//div[@class='model-name']")) == 2)
		for mod in driver.find_elements_by_xpath("//div[@class='model-name']"):
			if not (mod.text == "regress" or mod.text == "gradi"):
				print "Name filter did not apply correctly"
				return False
		
	except:
		print "Failed to filter models"
		return False

	return True

def main():
	failcount = 0
	d = tu.newtest()
	
	if not nameFilterTest(d):
		failcount += 1

	tu.endtest(d)
	sys.exit(failcount)

if __name__ == '__main__':
	main()

