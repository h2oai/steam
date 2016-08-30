import sys
import time
from selenium import webdriver
from selenium.common import exceptions as se
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support.ui import Select

def indexOfModel(driver, mod):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		modli = driver.find_elements_by_xpath("//div[@class='model-name']")
		for i in range(len(modli)):
			if mod == modli[i].text:
				return i
	except:
		return -1
	return -1

def indexOfTag(driver, tag):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		tagli = driver.find_elements_by_xpath("//div[@class='label-name']")
		for i in range(len(tagli)):
			if tag == tagli[i].text:
				return i
	except:
		return -1
	return -1

def createTag(driver, name, desc):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	driver.find_element_by_xpath("//button[text()='Create New Label']").click()
	wait.until(lambda x: x.find_element_by_xpath("//input[@name='name']").is_displayed())
	driver.find_element_by_xpath("//input[@name='name']").send_keys(name)
	driver.find_element_by_xpath("//textarea[@name='description']").send_keys(desc)
	driver.find_element_by_xpath("//button[text()='Save']").click()
	wait.until(lambda x: x.find_element_by_xpath("//div[@class='label-name' and text()='{0}']".format(name)).is_displayed())

def deleteTag(driver, name):
	#short timeout because after deleting a tag it should be gone and wait should fail
	wait = WebDriverWait(driver, timeout=2, poll_frequency=0.2)
	ind = indexOfTag(driver, name)
	if ind == -1:
		print "failed to find tag {0}".format(name)
		return False
	trash = driver.find_elements_by_class_name("fa-trash")[ind]
	trash.click()
	try:
		wait.until(lambda x: not x.find_element_by_xpath("//div[text()='{0}']".format(name).is_displayed()))
	except:
		return True
	return False

def applyTagToModel(driver, tag, mod):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	ind = indexOfModel(driver, mod)
	sel = Select(driver.find_elements_by_xpath("//select[@name='labelSelect']")[ind])
	sel.select_by_visible_text(tag)
	time.sleep(1)
	return True

def getModelTag(driver, model):
	ind = indexOfModel(driver, model)
	if ind == -1:
		return None
	tag = Select(driver.find_elements_by_xpath("//select[@name='labelSelect']")[int(ind)]).all_selected_options
	return tag

def goProjectConfig(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		driver.find_element_by_xpath("//a[text()='Configurations']").click()
		wait.until(lambda x: x.find_element_by_xpath("//button[text()='Create New Label']").is_displayed())
	except Exception as e:
		print e
		print "failed to enter project config"

def goProjectDeployment(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		driver.find_element_by_xpath("//a[text()='Deployment']").click()
		wait.until(lambda x: x.find_element_by_xpath("//span[text()='Deployment']"))
	except:
		print "Failed to navigate to deployment page"
		return False
	return True

def sortModels(driver, feat, ascending):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	driver.find_element_by_class_name("filter-dropdown-invoker").click()
	try:
		wait.until(lambda x: x.find_element_by_xpath("//li[text()='ASC ']").is_displayed())
		if ascending:
			driver.find_element_by_xpath("//li[text()='ASC ']").click()
		else:
			driver.find_element_by_xpath("//li[text()='DES ']").click()
	except Exception as e:
		print e 
		driver.find_element_by_class_name("filter-dropdown-invoker").click()
		return False
	try:
		if feat == "R2":
			driver.find_element_by_xpath("//sup[text()='2']").click()
			time.sleep(1)
		else:
			driver.find_element_by_xpath("//li[text()='{0}']".format(feat)).click()
			time.sleep(1)
	except Exception as e:
		print e
		driver.find_element_by_class_name("filter-dropdown-invoker").click()
		print "failed to sort models"
		return False	
	driver.find_element_by_class_name("filter-dropdown-invoker").click()
	return True
	
	

def goHome(driver):
	driver.find_element_by_class_name("logo").click()
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	wait.until(lambda x: x.find_element_by_class_name("start-project").is_displayed())

def newProject(driver):
	strt = driver.find_element_by_class_name("start-project")
	strt.click()
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		wait.until(lambda x: x.find_element_by_class_name("select-cluster"))
	except Exception as e:
		print e
		print "what what what"
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
		print e
		print "Failed to navigate to cluster page through navbar"
		return False

def goModels(driver):
	try:
		driver.find_element_by_xpath("//a[text()='Models']").click()
		wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
		wait.until(lambda x: x.find_element_by_xpath("//span[text()='Models']").is_displayed())
	except Exception as e:
		print e
		print "Failed to navigate to models page"
		return False

def goServices(driver):
	try:
		driver.find_element_by_class_name("fa-cloud").click()
		wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
		wait.until(lambda x: x.find_element_by_xpath("//li[text()='Services']"))
	except:
		print "Failed to navigate to services page"
		return False
	return True

def goProjects(driver):
	driver.find_element_by_class_name("fa-folder").click()
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	wait.until(lambda x: x.find_element_by_xpath("//div[@class='project-details']"))


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

def viewProject(driver, proj):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		wait.until(lambda x: x.find_element_by_xpath("//header[text()='{0}']".format(proj)))
		driver.find_element_by_xpath("//header[text()='{0}']".format(proj)).click()
		wait.until(lambda x: x.find_element_by_class_name("model-name"))
	except:
		print "Cannot find project"
		return False
	return True

def selectDataframe(driver, frame):
	#locate the select
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		wait.until(lambda x: x.find_element_by_xpath("//option[@value='{0}']".format(frame)))
		sel = Select(driver.find_element_by_name("selectDataframe"))
		sel.select_by_visible_text(frame)
	except Exception as e:
		print "Cannot select dataframe"
		return False
	return True

def selectModelCategory(driver, category):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		wait.until(lambda x: x.find_element_by_xpath("//option[@value='{0}']".format(category)))
		sel = Select(driver.find_element_by_name("selectModelCategory"))
		sel.select_by_visible_text(category)
	except Exception as e:
		print e
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

def deployModel(driver, mod, name):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	ind = indexOfModel(driver, mod)
	if ind == -1:
		raise se.ElementNotVisibleException()
	driver.find_elements_by_xpath("//span[text()='deploy model']")[ind].click()
	wait.until(lambda x: len(x.find_elements_by_xpath("//input[@type='text']")) == 2)
	driver.find_elements_by_xpath("//input[@type='text']")[1].send_keys(name)
	driver.find_element_by_class_name("deploy-button").click()
	wait.until(lambda x: x.find_element_by_class_name("deployed-services"))
	

def newtest():
	driver = webdriver.Chrome()
	driver.get("http://superuser:superuser@localhost:9000")
	driver.find_element_by_css_selector("input").click()
	return driver

def endtest(driver):
	driver.quit()

