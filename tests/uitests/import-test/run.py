import sys
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support.ui import Select

def goHome(driver):
	driver.find_element_by_class_name("logo").click()
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	wait.until(lambda x: x.find_element_by_class_name("start-project").is_displayed())

def newProject(driver):
	strt = driver.find_element_by_class_name("start-project")
	strt.click()
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		wait.until(lambda x: x.find_element_by_xpath("//div[@class='import-new-project']").is_displayed())
	except:
		return False
	return True

def goClusters(driver):
	try:
		clust = driver.find_element_by_class_name("fa-cube")
		clust.click()
		wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
		wait.until(lambda x: x.find_element_by_xpath("//div[@class='clusters']").is_displayed())
		return True
	except Exception as e:
		print "Failed to navigate to cluster page through navbar"
		return False

def goProjects(driver):
	try:
		proj = driver.find_element_by_class_name("fa-folder")
		proj.click()
		wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
		wait.until(lambda x: x.find_element_by_xpath("//div[@class='project-details']"))
		return True
	except:
		print "Failed to navigate to projects page through navbar"
		return False

def clusterExists(driver, addr, port, name):
	if not goClusters(driver):
		return False
	try:
		elm = driver.find_element_by_link_text("{0}@ {1}:{2}".format(name, addr, port))
		return True
	except Exception as e:
		print "New cluster did not appear on cluster page"
		return False

def addCluster(driver, addr, port, name):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		wait.until(lambda x: x.find_element_by_name("ip-address").is_displayed())
		driver.find_element_by_name("ip-address").send_keys(addr)
		driver.find_element_by_name("port").send_keys(port)
		driver.find_element_by_xpath("//div[@class='connect-cluster']//button").click()
		wait.until(lambda x: x.find_element_by_xpath("//div[text()='{0}']".format(name)))
	except:
		print "Cannot add new cluster"
		return False
	return True

def selectDataframe(driver, frame):
	#locate the select
	try:
		sel = Select(driver.find_element_by_name("selectDataframe"))
		sel.select_by_visible_text(frame)
	except Exception as e:
		print "Cannot select dataframe"
		return False
	return True

def selectModelCategory(driver, category):
	try:
		sel = Select(driver.find_element_by_name("selectModelCategory"))
		sel.select_by_visible_text(category)
	except:
		print "Failed to select model category"
		return False
	return True

def selectModel(driver, name):
	try:
		mod = driver.find_element_by_xpath("//input[@type='checkbox' and @name='{0}']".format(name))
		mod.click()
	except:
		print "Failed to select model"
		return False
	return True

def createProjectTest(driver):
	try:
		newProject(driver)
		addCluster(driver, "localhost", "54535", "steamtest")
		wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
		wait.until(lambda x: x.find_element_by_class_name("select-cluster").is_displayed())
		driver.find_element_by_xpath("//div[@class='select-cluster']//button").click()
		wait.until(lambda x: x.find_element_by_css_selector("select").is_displayed())
	except:
		print "Failed to connect to cluster"
		return False
	try:
		wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
		wait.until(lambda x: x.find_element_by_xpath("//option[@value='bank_full.hex']"))	
		if not selectDataframe(driver, "bank_full.hex"):
			print "Failed to access expected dataframe"
			return False
		wait.until(lambda x: x.find_element_by_xpath("//option[@value='Regression']").is_displayed())
		if not selectModelCategory(driver, "Regression"):
			print "Failed to find regression models"
			return False
		wait.until(lambda x: x.find_element_by_name("regress").is_displayed())
		if not selectModel(driver, "regress"):
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
	goHome(driver)
	if not goProjects(driver):
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
	goHome(driver)
	if not newProject(driver):
		print "Failed to navigate to new project page after creating a project"
		return False
	driver.find_element_by_xpath("//div[@class='select-cluster']//button").click()
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		wait.until(lambda x: x.find_element_by_xpath("//option[@value='bank_full.hex']"))
		if not selectDataframe(driver, "bank_full.hex"):
			print "Failed to access expected dataframe"
			return False
		wait.until(lambda x: x.find_element_by_xpath("//option[@value='Regression']"))
		if not selectModelCategory(driver, "Regression"):
			print "Failed to find regression models"
			return False
		models = ["regress", "gradi", "missin", "linmiss"]
		for mod in models:
			if not selectModel(driver, mod):
				print "Failed to select expected model for importing"
				return False
		driver.find_element_by_xpath("//div[@class='name-project']//input").send_keys("multimod")
		driver.find_element_by_xpath("//button[text()='Create Project']").click()	
		for mod in models:
			wait.until(lambda x: x.find_element_by_xpath("//div[@class='model-name' and text()='{0}']".format(mod)))
	except:
		print "new project failed to include all imported models"
		return False
	return True


def main():
	failcount = 0
	d = webdriver.Chrome()
	d.get("http://superuser:superuser@localhost:9000")
	if not createProjectTest(d):
		failcount += 1
	if not viewProjectTest(d):
		failcount += 1
	if not importMultiTest(d):
		failcount += 1

	d.quit()
	sys.exit(failcount)

if __name__ == '__main__':
	main()
