import sys
import time
import testutil as tu
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
		tu.deployModel(driver, "regress", "happy")
		text = driver.find_element_by_xpath("//div[@class='panel-title']/span").text
		if not "happy" in text:
			print "Failed to create a named service"
			return False

	except Exception as e:
		print e
		print "Failed to deploy a model"
		return False
	return True

def deleteTest(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		driver.find_element_by_xpath("//div[text()='Stop Service']").click()
		wait.until(lambda x: len(x.find_elements_by_class_name("services-panel")) == 0)
	except:
		driver.refresh()
		time.sleep(1)
		if len(driver.find_elements_by_class_name("services-panel")) == 0:
			print "Deployment page must be refreshed before stopped services are removed"
		else:
			print "Failed to stop/delete a service"
		return False
	return True

def projectDeployTest(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		tu.goHome(driver)
		tu.newProject(driver)
		wait.until(lambda x: x.find_element_by_xpath("//div[@class='select-cluster']//button"))
		driver.find_element_by_xpath("//div[@class='select-cluster']//button").click()
		tu.selectDataframe(driver, "bank_full.hex")
		tu.selectModelCategory(driver, "Regression")
		tu.selectModel(driver, "regress")
		tu.selectModel(driver, "gradi")
		tu.selectModel(driver, "missin")
		driver.find_element_by_xpath("//div[@class='name-project']//input").send_keys("projtest")
		driver.find_element_by_xpath("//button[text()='Create Project']").click()
		for m in ["regress", "gradi", "missin"]:
			wait.until(lambda x: x.find_element_by_xpath("//div[@class='model-name' and text()='{0}']".format(m)))
		tu.deployModel(driver, "gradi", "swell")
		tu.goHome(driver)
		tu.goProjects(driver)
		time.sleep(1)
		tu.viewProject(driver, "deptest")
		tu.goProjectDeployment(driver)
	except Exception as e:
		print e
		print "Failed to setup project deploy test"
		return False
	try:
		time.sleep(1)
		driver.refresh()
		time.sleep(1)
		wait.until(lambda x: x.find_element_by_class_name("services-panel"))
		print "Deployed service is displayed in unassociated project"
		return False
	except:
		return True

	return False

def cleanupTest(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	tu.goHome(driver)
	tu.goServices(driver)
	try:
		wait.until(lambda x: x.find_element_by_xpath("//div[text()='Stop Service']"))
	except:
		i = 1
	stp = driver.find_elements_by_xpath("//div[text()='Stop Service']")
	for serv in stp:
		serv.click()
	try:
		wait.until(lambda x: len(x.find_elements_by_class_name("services-panel")) == 0)
	except:
		print "failed to stop running services"
		return False
	return True

def multiDeployTest(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		tu.goHome(driver)
		tu.newProject(driver)
		wait.until(lambda x: x.find_element_by_xpath("//div[@class='select-cluster']//button"))
		driver.find_element_by_xpath("//div[@class='select-cluster']//button").click()
		tu.selectDataframe(driver, "bank_full.hex")
		tu.selectModelCategory(driver, "Regression")
		tu.selectModel(driver, "linmiss")
		driver.find_element_by_xpath("//div[@class='name-project']//input").send_keys("multidep")
		driver.find_element_by_xpath("//button[text()='Create Project']").click()
		wait.until(lambda x: x.find_element_by_xpath("//div[@class='model-name' and text()='linmiss']"))
		tu.deployModel(driver, "linmiss", "double")
		tu.goHome(driver)
		tu.goServices(driver)
	except Exception as e:
		print e
		print "failed to setup multi deploy test"
		return False
	try:
		time.sleep(3)
		wait.until(lambda x: x.find_element_by_class_name("panel-title"))
		deps = driver.find_elements_by_xpath("//div[@class='panel-title']/span")
		names = ["double", "swell"]
		for name in names:
			good = False
			for dep in deps:
				if name in dep.text:
					good = True
			if not good:
				print "Failed to locate all deployed services on services page"
				return False

	except Exception as e:
		print e
		print "Failed to find deployments on services page"
		return False	
	return True

def main():
	failcount = 0
	d = tu.newtest()
	
	if not deployOneTest(d):
		failcount += 1
	if not deleteTest(d):
		failcount += 1
	if not projectDeployTest(d):
		failcount += 1
	# test all services from multiple projects showing up in services
	if not multiDeployTest(d):
		failcount += 1
	# test that stopping services from services page removes them from project deployments
	if not cleanupTest(d):
		failcount += 1
<<<<<<< 393bf8c14263d63fef5d2cd760f1faad04767fad
=======

>>>>>>> multi-deployment test
	tu.endtest(d)
	sys.exit(failcount)

if __name__ == '__main__':
	main()

