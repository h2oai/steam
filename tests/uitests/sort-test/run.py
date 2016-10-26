import sys
import time
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support.ui import Select
import testutil as tu

def outOfOrder(feat, order):
	print "Models did not sort in {0} order properly for metric {1}".format(order, feat)

def isAscending(driver, metric):
	elms = driver.find_elements_by_xpath("//div[@class='cell' and @name='{0}']".format(metric))
	for i in range(len(elms) - 1):
		if float(elms[i].text) > float(elms[i+1].text):
			return False
	return True

def isDescending(driver, metric):
	elms = driver.find_elements_by_xpath("//div[@class='cell' and @name='{0}']".format(metric))
	for i in range(len(elms) - 1):
		if float(elms[i].text) < float(elms[i+1].text):
			return False
	return True

def checkSorting(driver, metric, name):
	passed = True
	tu.sortModels(driver, metric, True)
	if not isAscending(driver, name):
		outOfOrder(metric, "ASC")
		passed = False
	tu.sortModels(driver, metric, False)
	if not isDescending(driver, name):
		outOfOrder(metric, "DES")
		passed = False
	return passed

def regressSortTest(driver):
	passed = True
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		tu.newProject(driver)
		tu.selectCluster(driver, "steamtest")
		tu.selectDataframe(driver, "bank_full.hex")
		tu.selectModelCategory(driver, "Regression")
		models = ["regress", "gradi", "missin", "linmiss"] 
		for mod in models:
			tu.selectModel(driver, mod)
		driver.find_element_by_xpath("//div[@class='name-project']//input").send_keys("regsort")
		driver.find_element_by_xpath("//button[text()='Create Project']").click()	
		wait.until(lambda x: x.find_element_by_xpath("//li[@id='projectIdCrumb']"))
		driver.refresh()
		for mod in models:
			wait.until(lambda x: x.find_element_by_xpath("//div[@class='model-name' and text()='{0}']".format(mod)))	
	except Exception as e:
		print e
		print "failed to create project"
		return False
	
	try:
		passed = passed and checkSorting(driver, "MSE", "mse")
		passed = passed and checkSorting(driver, "R2", "r2")
		passed = passed and checkSorting(driver, "MRD", "mrd")
	except Exception as e:
		print e
		print "models failed to sort"
		return False
	
	return passed

def binomSortTest(driver):	
	passed = True
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		tu.goHome(driver)
		tu.newProject(driver)
		tu.selectCluster(driver, "steamtest")
		tu.selectDataframe(driver, "bank_full.hex")
		tu.selectModelCategory(driver, "Binomial")
		models = ["first", "second", "third", "fourth"] 
		for mod in models:
			tu.selectModel(driver, mod)
		driver.find_element_by_xpath("//div[@class='name-project']//input").send_keys("binsort")
		driver.find_element_by_xpath("//button[text()='Create Project']").click()	
		wait.until(lambda x: x.find_element_by_xpath("//li[@id='projectIdCrumb']"))
		driver.refresh()
		for mod in models:
			wait.until(lambda x: x.find_element_by_xpath("//div[@class='model-name' and text()='{0}']".format(mod)))	
	except Exception as e:
		print e
		print "failed to create project"
		return False
	
	try:
		passed = passed and checkSorting(driver, "MSE", "mse")
		passed = passed and checkSorting(driver, "AUC", "auc")
		passed = passed and checkSorting(driver, "Gini", "gini")
		passed = passed and checkSorting(driver, "LogLoss", "logloss")
	except:
		return False
	return passed

def mulnomSortTest(driver):
	passed = True
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:
		tu.goHome(driver)
		tu.newProject(driver)
		tu.selectCluster(driver, "steamtest")
		tu.selectDataframe(driver, "bank_full.hex")
		tu.selectModelCategory(driver, "Multinomial")
		models = ["multinom", "valimon", "multimiss", "vamiss"] 
		for mod in models:
			tu.selectModel(driver, mod)
		driver.find_element_by_xpath("//div[@class='name-project']//input").send_keys("multisort")
		driver.find_element_by_xpath("//button[text()='Create Project']").click()	
		wait.until(lambda x: x.find_element_by_xpath("//li[@id='projectIdCrumb']"))
		driver.refresh()
		for mod in models:
			wait.until(lambda x: x.find_element_by_xpath("//div[@class='model-name' and text()='{0}']".format(mod)))	
	except Exception as e:
		print e
		print "failed to create project"
		return False
	
	try:
		passed = passed and checkSorting(driver, "MSE", "mse")
		passed = passed and checkSorting(driver, "R2", "r2")
		passed = passed and checkSorting(driver, "LogLoss", "logloss")
	except:
		return False
	return passed

def main():
	failcount = 0

	d = tu.newtest()
	time.sleep(2)
	if not regressSortTest(d):
		failcount += 1
	tu.endtest(d)
	d = tu.newtest()
	time.sleep(2)
	if not binomSortTest(d):
		failcount += 1
	tu.endtest(d)
	d = tu.newtest()
	time.sleep(2)
	if not mulnomSortTest(d):
		failcount += 1

	tu.endtest(d)
	sys.exit(failcount)
	

if __name__ == "__main__":
	main()

