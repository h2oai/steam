import sys
import time
import testutil as tu
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait

def createTagTest(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	tu.newProject(driver)
	tu.selectCluster(driver, "steamtest")
	tu.selectDataframe(driver, "bank_full.hex")
	tu.selectModelCategory(driver, "Regression")
	try:
		models = ["missin"]
		for mod in models:
			tu.selectModel(driver, mod)
		driver.find_element_by_xpath("//div[@class='name-project']//input").send_keys("taggytest")
		driver.find_element_by_xpath("//button[text()='Create Project']").click()
		for mod in models:
			wait.until(lambda x: x.find_element_by_xpath("//div[@class='model-name' and text()='{0}']".format(mod)))
	except:
		print "tag test setup failed"
		return False	
	try:
		tu.goProjectConfig(driver)
		driver.find_element_by_xpath("//button[text()='Create New Label']").click()
		wait.until(lambda x: x.find_element_by_xpath("//input[@name='name']").is_displayed())
		driver.find_element_by_xpath("//input[@name='name']").send_keys("prime")
		driver.find_element_by_xpath("//textarea[@name='description']").send_keys("A cool tag to use")
		driver.find_element_by_xpath("//button[text()='Save']").click()
		wait.until(lambda x: x.find_element_by_xpath("//div[@class='label-name' and text()='prime']").is_displayed())
	except Exception as e:
		print e
		print "failed to create a new label"
		return False
	
	return True


def applyTagTest(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	tu.goModels(driver)
	time.sleep(1)
	try:
		tu.applyTagToModel(driver, "prime", "missin")
	except:
		print "Failed to apply tag to a model"
		return False	
	try:
		tu.goProjectConfig(driver)
		tag = driver.find_element_by_xpath("//span[@class='model-name']")
		if tag.text == "Not currently applied to a model":
			print "Model associated with tag did not update"
	except Exception as e:
		print e
		return False
	return True

def multiTagTest(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	try:	
		driver.find_element_by_xpath("//button[text()='Create New Label']").click()
		wait.until(lambda x: x.find_element_by_xpath("//input[@name='name']").is_displayed())
		driver.find_element_by_xpath("//input[@name='name']").send_keys("crime")
		driver.find_element_by_xpath("//textarea[@name='description']").send_keys("An ok tag to use")
		driver.find_element_by_xpath("//button[text()='Save']").click()
		wait.until(lambda x: x.find_element_by_xpath("//div[@class='label-name' and text()='crime']").is_displayed())
	except Exception as e:
		print e
		print "failed to create a second label"
		return False

	try:	
		tu.goModels(driver)
		time.sleep(1)
		tu.applyTagToModel(driver, "crime", "missin")
	except Exception as e:
		print "Failed to change the tag attached to a model"
		return False
	try:
		tu.goProjectConfig(driver)
		wait.until(lambda x: len(x.find_elements_by_xpath("//span[@class='model-name']")) == 2)
		mods = driver.find_elements_by_xpath("//span[@class='model-name']")
		if mods[0].text == mods[1].text:
			print "Multiple tags reference the same model"
			return False
	except Exception as e:
		print e
		print "multitag test ecountered an exceptional case"
		return False
	return True

def conflictTagTest(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	tu.goHome(driver)
	tu.newProject(driver)
	driver.find_element_by_xpath("//div[@class='select-cluster']//button").click()
	tu.selectDataframe(driver, "bank_full.hex")
	tu.selectModelCategory(driver, "Regression")	
	try:
		models = ["regress", "gradi"]
		for mod in models:
			tu.selectModel(driver, mod)
		driver.find_element_by_xpath("//div[@class='name-project']//input").send_keys("tagtest")
		driver.find_element_by_xpath("//button[text()='Create Project']").click()
		wait.until(lambda x: x.find_element_by_xpath("//li[@id='projectIdCrumb']"))
		driver.refresh()
		for mod in models:
			wait.until(lambda x: x.find_element_by_xpath("//div[@class='model-name' and text()='{0}']".format(mod)))
	except:
		print "conflict tag test setup failed"
		return False	
	try:
		tag = tu.getModelTag(driver, "regress")
		tu.goProjectConfig(driver)
		tu.createTag(driver, "taggy", "the taggiest tag")
		tu.goModels(driver)
		for mod in ["regress", "gradi"]:
			wait.until(lambda x: x.find_element_by_xpath("//div[@class='model-name' and text()='{0}']".format(mod)))
		tu.applyTagToModel(driver, "taggy", "regress")
		tu.applyTagToModel(driver, "taggy", "gradi")
		tag = tu.getModelTag(driver, "regress")
		if len(tag) > 1:
			print "Model has multiple tags associated with it"
			return False
		if len(tag) == 1 and tag[0].text == "taggy":
			print "Tag stayed assicated with model after being applied to a different model"
			return False
	except Exception as e:
		print e
		print "Failed to create and apply a new tag to models"
		return False
	return True

def deleteTagTest(driver):
	wait = WebDriverWait(driver, timeout=5, poll_frequency=0.2)
	tu.goHome(driver)
	tu.newProject(driver)
	driver.find_element_by_xpath("//div[@class='select-cluster']//button").click()
	tu.selectDataframe(driver, "bank_full.hex")
	tu.selectModelCategory(driver, "Regression")
	try:
		models = ["regress"]
		for mod in models:
			tu.selectModel(driver, mod)
		driver.find_element_by_xpath("//div[@class='name-project']//input").send_keys("deletetest")
		driver.find_element_by_xpath("//button[text()='Create Project']").click()
		for mod in models:
			wait.until(lambda x: x.find_element_by_xpath("//div[@class='model-name' and text()='{0}']".format(mod)))
	except:
		return False
	try:
		tu.goProjectConfig(driver)
		tu.createTag(driver, "delet", "to be deleted at once")
		tu.goModels(driver)
		wait.until(lambda x: x.find_element_by_xpath("//div[@class='model-name' and text()='regress']"))
		tu.applyTagToModel(driver, "delet", "regress")
		tu.goProjectConfig(driver)
		if not tu.deleteTag(driver, "delet"):
			print "failed to delete tag"
			return False
		tu.goModels(driver)
		wait.until(lambda x: x.find_element_by_xpath("//div[@class='model-name' and text()='regress']"))
		time.sleep(2)
		tag = tu.getModelTag(driver, "regress")
		if len(tag) > 0 and tag[0].text == "delet":
			print "Model attached to deleted tag"
			return False
	except:
		print "Failed to delete tag"
		return False
	return True

def main():
	failcount = 0
	d = tu.newtest()
	
	if not createTagTest(d):
		failcount += 1
	if not applyTagTest(d):
		failcount += 1
	if not multiTagTest(d):
		failcount += 1
	if not conflictTagTest(d):
		failcount += 1	
	if not deleteTagTest(d):
		failcount += 1
			
	tu.endtest(d)
	sys.exit(failcount)


if __name__ == '__main__':
	main()

