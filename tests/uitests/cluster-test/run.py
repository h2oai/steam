import sys
import time
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
	
def goHome(driver):
	driver.find_element_by_class_name("logo").click()
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	wait.until(lambda x: x.find_element_by_class_name("start-project").is_displayed())

def newProject(driver):
	strt = driver.find_element_by_class_name("start-project")
	strt.click()

def goClusters(driver):
	try:
		clust = driver.find_element_by_class_name("fa-cube")
		clust.click()
		wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
		wait.until(lambda x: x.find_element_by_xpath("//div[@class='clusters']").is_displayed())
		return True
	except Exception as e:
		print "Failed to navigate to cluster page"
		return False

def clusterExists(driver, name):
	if not goClusters(driver):
		return False
	try:
		elm = driver.find_element_by_link_text("{0}".format(name))
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
		driver.find_element_by_xpath("//button[@type='submit']").click()
		wait.until(lambda x: x.find_element_by_xpath("//span[text()='{0}']".format(name)))
	except:
		print "Cannot add new cluster"
		return False
	return True

def connectTest(driver):
	goHome(driver)
	newProject(driver)
	addCluster(driver, "localhost", "54535", "steamtest")
	return clusterExists(driver, "steamtest")	

def deleteClusterTest(driver):
	if not goClusters(driver):
		return False
	try:
		l = driver.find_elements_by_xpath("//button[@class='remove-cluster-button']")
		if len(l) == 0:
			print "Failed to view cluster on clusters page"
			return False
		init = len(l)
		l[0].click()
		wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
		wait.until(lambda x: len(x.find_elements_by_xpath("//button[@class='remove-cluster-button']")) < init)
	except Exception as e:
		print e
		print "Failed to delete cluster"
		return False
	return True

def dataframeTest(driver):
	try:
		goHome(driver)
		newProject(driver)
		if not addCluster(driver, "localhost", "54535", "steamtest"):
			print "Failed to re-connect to a cluster that had been deleted"
			return False
		if not clusterExists(driver, "steamtest"):
			print "Failed to re-connect to a cluster that had been deleted"
			return False	
		goHome(driver)
		newProject(driver)
		wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
		wait.until(lambda x: x.find_element_by_class_name("select-cluster").is_displayed())
		driver.find_element_by_xpath("//div[@class='select-cluster']//button").click()
		wait.until(lambda x: x.find_element_by_css_selector("select").is_displayed())
	except:
		print "Failed to re-connect to a cluster that had been deleted"
		return False
	
	try:	
		wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
		wait.until(lambda x: x.find_element_by_xpath("//option[@value='bank_full.hex']"))
		driver.find_element_by_xpath("//option[@value='arrhythmia.hex']")
		driver.find_element_by_xpath("//option[@value='X50_cattest_test.hex']")
		driver.find_element_by_xpath("//option[@value='X50_cattest_train.hex']")
	except:
		print "Dataframe select doesn't include all dataframes related to cluster"
		return False
	
	return True


def main():
	failcount = 0
	d = webdriver.Chrome()
	d.get("http://superuser:superuser@localhost:9000")
	if not connectTest(d):
		failcount += 1
	if not deleteClusterTest(d):
		failcount += 1
	if not dataframeTest(d):
		failcount += 1
		
	d.quit()
	sys.exit(failcount)
	

if __name__ == '__main__':
	main()

