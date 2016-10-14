import testutil as tu
import time
import sys
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support.ui import Select
from selenium.webdriver.chrome.options import Options

def createProjectTest(driver):
	try:
		tu.newProject(driver)
		wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
		wait.until(lambda x: x.find_element_by_class_name("select-cluster").is_displayed())
		tu.selectCluster(driver, "steamtest")
	except:
		print "Failed to connect to cluster"
		return False
	try:
		wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
		wait.until(lambda x: x.find_element_by_xpath("//option[@value='bank_full.hex']"))	
		if not tu.selectDataframe(driver, "bank_full.hex"):
			print "Failed to access expected dataframe"
			return False
		wait.until(lambda x: x.find_element_by_xpath("//option[@value='Regression']").is_displayed())
		if not tu.selectModelCategory(driver, "Regression"):
			print "Failed to find regression models"
			return False
		wait.until(lambda x: x.find_element_by_name("regress").is_displayed())
		if not tu.selectModel(driver, "regress"):
			print "Failed to select an expected model for importing"
			return False
		pnam = driver.find_element_by_xpath("//div[@class='name-project']//input")
		pnam.send_keys("imported")
		driver.find_element_by_xpath("//button[text()='Create Project']").click()
	except:
		print "Failed to create a project with a selected model"
		return False
	try:
		wait.until(lambda x: x.find_element_by_xpath("//div[@class='model-name' and text()='regress']"))
	except Exception as e:
		print "Failed to locate imported model on project page"
		return False	
	return True	

def viewProjectTest(driver):
	tu.goHome(driver)
	if not tu.goProjects(driver):
		print "Failed setting up viewproject"
		return False
	try:
		proj = driver.find_element_by_xpath("//div[@class='project-metadata']/header[text()='imported']")
		proj.click()
		wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
		wait.until(lambda x: x.find_element_by_xpath("//div[@class='model-name' and text()='regress']"))
	except:
		print "Failed to open models page after selecting a project"
		return False
	return True

def importMultiTest(driver):
	tu.goHome(driver)
	if not tu.newProject(driver):
		print "Failed to navigate to new project page after creating a project"
		return False
	tu.selectCluster(driver, "steamtest")
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		wait.until(lambda x: x.find_element_by_xpath("//option[@value='bank_full.hex']"))
		if not tu.selectDataframe(driver, "bank_full.hex"):
			print "Failed to access expected dataframe"
			return False
		wait.until(lambda x: x.find_element_by_xpath("//option[@value='Regression']"))
		if not tu.selectModelCategory(driver, "Regression"):
			print "Failed to find regression models"
			return False
		models = ["regress", "gradi", "missin", "linmiss"]
		for mod in models:
			if not tu.selectModel(driver, mod):
				print "Failed to select expected model for importing"
				return False
		driver.find_element_by_xpath("//div[@class='name-project']//input").send_keys("multimod")
		driver.find_element_by_xpath("//button[text()='Create Project']").click()	
		time.sleep(2)
		driver.refresh()
		time.sleep(2)
		for mod in models:
			wait.until(lambda x: x.find_element_by_xpath("//div[@class='model-name' and text()='{0}']".format(mod)))
	except Exception as e:
		print e
		print "new project failed to include all imported models"
		return False
	return True


def main():
	failcount = 0
	d = tu.newtest()
	d.get("http://superuser:superuser@localhost:9000")
	if not createProjectTest(d):
		print "createproj"
		failcount += 1
	if not viewProjectTest(d):
		print "viewproj"
		failcount += 1
	if not importMultiTest(d):
		print "importmulti"
		failcount += 1

	d.quit()
	sys.exit(failcount)

if __name__ == '__main__':
	main()
