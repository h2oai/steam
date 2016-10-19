import sys
import time
import testutil as tu
import urlparse
from browsermobproxy import Server
from selenium.webdriver.common.keys import Keys
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait

def deployOneTest(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	tu.newProject(driver)
	tu.selectCluster(driver, "steamtest")
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
		print "Failed to deploy a model"
		return False
	return True

def deleteTest(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		tu.stopService(driver, "happy")
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
		tu.selectCluster(driver, "steamtest")
		tu.selectDataframe(driver, "bank_full.hex")
		tu.selectModelCategory(driver, "Regression")
		tu.selectModel(driver, "regress")
		tu.selectModel(driver, "gradi")
		tu.selectModel(driver, "missin")
		driver.find_element_by_xpath("//div[@class='name-project']//input").send_keys("projtest")
		#driver.find_element_by_xpath("//div").send_keys(Keys.F12)
		driver.find_element_by_xpath("//button[text()='Create Project']").click()
		time.sleep(2)
		driver.refresh()
		for m in ["regress", "gradi", "missin"]:
			wait.until(lambda x: x.find_element_by_xpath("//div[@class='model-name' and text()='{0}']".format(m)))
		tu.deployModel(driver, "gradi", "swell")
		tu.goHome(driver)
		tu.goProjects(driver)
		time.sleep(1)
		tu.viewProject(driver, "deptest")
		tu.goProjectDeployment(driver)
	except:
		print "Failed to deploy gradi"
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
	time.sleep(2)
	if not (tu.serviceExists(driver, "swell") and tu.serviceExists(driver, "double")):
		print "A service wasn't there"
	cnt = len(driver.find_elements_by_class_name("services-panel"))
	tu.stopService(driver, "swell")
	tu.stopService(driver, "double")
	try:
		wait.until(lambda x: len(x.find_elements_by_class_name("services-panel")) <= (cnt - 2))
	except:
		print "failed to stop running services"
		return False
	return True

def multiDeployTest(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		tu.goHome(driver)
		tu.newProject(driver)
		tu.selectCluster(driver, "steamtest")
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
		print "Failed to find deployments on services page"
		return False	
	return True

def main():
	s = Server('/home/pjr/browsermob/bin/browsermob-proxy', { 'port' : 1337})
	s.start()
	proxy = s.create_proxy({'port': 1338})
	failcount = 0
	d = tu.newProxytest(proxy)
	proxy.new_har(options={'captureHeaders':False, 'captureContent': True})
	if not deployOneTest(d):
		failcount += 1
	if not deleteTest(d):
		failcount += 1
	if not projectDeployTest(d):
		failcount += 1
		out = open('deploy.har', 'w')
		out.write(str(proxy.har))
		out.close()
	# test all services from multiple projects showing up in services
	if not multiDeployTest(d):
		failcount += 1
	# test that stopping services from services page removes them from project deployments
	if not cleanupTest(d):
		failcount += 1

	tu.endtest(d)
	s.stop()
	sys.exit(failcount)

if __name__ == '__main__':
	main()

