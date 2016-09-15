import sys
import time
import subprocess as sp
import re
from selenium import webdriver
from selenium.common import exceptions as se
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support.ui import Select

"""
Perm id		Permission		Index
============================================
 1			M Role			 8
 2			V Role			19
 3			M Workgroup		10
 4			V Workgroup		21
 5			M Identities	 4
 6			V Identities	15
 7			M engines		 3
 8			V engines		14
 9			M clusters		 0
10			V clusters		11
11			M Projects		 7
12			V Projects		18
13			M Datasrc		 2
14			V Datasrc		13
15			M Dataset		 1
16			V Dataset		12
17			M Model			 6
18			V Model			17
19			M label			 5
20			V Label			16
21			M service		 9
22			V service		20
"""

_steampath = "./steam"
if sys.platform.startswith("linux"):
	_steampath = "./steam-develop-linux-amd64/steam"
elif sys.platform == "darwin":
	_steampath = "./steam-develop-darwin-amd64/steam"
else:
	print "unsupported testing platform"
	sys.exit(1)
	
def cliLogin(name, pw):
	ret = sp.check_output("{0} login localhost:9000 --username={1} --password={1}"\
		.format(_steampath, name, pw), shell=True)

def createRole(role, desc, perm):
	ret = sp.check_output("{0} create role --name {1} --description \"{2}\""\
		.format(_steampath, role, desc), shell=True)
	i = int(re.search(r'\d+', ret).group())
	for p in perm:
		sp.Popen("{0} link role --with-permission --role-id={1} --permission-id={2} > /dev/null"\
			.format(_steampath, i, p), shell=True).communicate()
	return i

def createIdentity(name, pw):
	ret = sp.check_output("{0} create identity --name={1} --password={2}"\
		.format(_steampath, name, pw), shell=True)
	return int(re.search(r'\d+', ret).group())

def createWorkgroup(wg, desc):
	ret = sp.check_output("{0} create workgroup --name={1} --description=\"{2}\""\
		.format(_steampath, wg, desc), shell=True)
	return int(re.search(r'\d+', ret).group())

def assignRole(iden, role):
	x = sp.check_output("{0} link identity --with-role --identity-id={1} --role-id={2}"\
		.format(_steampath, iden, role), shell=True)

def assignWorkgroup(iden, wg):
	x = sp.check_output("{0} link identity --with-workgroup --identity-id={1} --workgroup-id={2}"\
		.format(_steampath, iden, wg), shell=True)

"""
EntityType

 1	Role
 2	Workgroup
 3	Identity
 4	Engine
 5	Cluster
 6	Project
 7	Datasrc
 8	Dataset
 9	model
10	label
11	service

"""

def shareEntity(kind, eid, wg, level):
	x = sp.check_output("{0} share entity --entity-type-id={1} --entity-id={2} --workgroup-id={3} --kind={4}"\
		.format(_steampath, kind, eid, wg, level), shell=True)

def goUsers(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	driver.find_element_by_class_name("fa-user").click()
	wait.until(lambda x: x.find_element_by_class_name("user-access"))

def goRoles(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	goUsers(driver)
	driver.find_element_by_xpath("//a[@class='tab' and text()='ROLES']").click()
	wait.until(lambda x: x.find_element_by_class_name("role-permissions"))

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
			time.sleep(3)
		else:
			driver.find_element_by_xpath("//li[text()='{0}']".format(feat)).click()
			time.sleep(3)
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
		print "Failed to access new project page"
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
	time.sleep(3)
	wait.until(lambda x: x.find_element_by_class_name("deployed-services"))
	
def createProject(driver, cluster, name, data, kind, mods):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	goHome(driver)
	newProject(driver)
	#select cluster by name
	#select the first cluster for now
	driver.find_element_by_xpath("//div[@class='select-cluster']//button").click()
	wait.until(lambda x: x.find_element_by_xpath("//select[@name='selectDataframe']"))
	sel = Select(driver.find_element_by_xpath("//select[@name='selectDataframe']"))
	sel.select_by_visible_text(data)
	wait.until(lambda x: x.find_element_by_xpath("//select[@name='selectModelCategory']"))
	sel = Select(driver.find_element_by_xpath("//select[@name='selectModelCategory']"))
	sel.select_by_visible_text(kind)
	for mod in mods:
		selectModel(driver, mod)
	driver.find_element_by_xpath("//div[@class='name-project']//input").send_keys(name)
	driver.find_element_by_xpath("//button[text()='Create Project']").click()
	for mod in mods:
		wait.until(lambda x: x.find_element_by_xpath("//div[@class='model-name' and text()='{0}']".format(mod)))

def testAs(user, pw):
	driver = webdriver.Chrome()
	driver.get("http://{0}:{1}@localhost:9000".format(user, pw))
	driver.find_element_by_css_selector("input").click()
	return driver

def newtest():
	driver = webdriver.Chrome()
	driver.get("http://superuser:superuser@localhost:9000")
	driver.find_element_by_css_selector("input").click()
	return driver

def endtest(driver):
	driver.quit()

