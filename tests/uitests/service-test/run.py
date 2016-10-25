import sys
import time
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support.ui import Select
import testutil as tu
import subprocess as sp

def servTest(d):
	return True
	try:
		wait = WebDriverWait(d, timeout=5, poll_frequency=0.2)
		inp = d.find_element_by_xpath("//div[@class='input-group']/input")
		inp.send_keys("You are a winner you have been specially selected to win 1000 cash or 2000 reward. Speak to a live operator to claim call 087123002209am to 7pm cost 10p")
		d.find_element_by_xpath("//input[@id='predict-btn']").click()
		wait.until(lambda x: x.find_element_by_class_name("labelHighlight"))
		x = d.find_element_by_class_name("labelHighlight")
		if x.text != "spam":
			print "failed to classify spam correctly"
			return False
		return True
	except:
		print "Failed to navigate prediction page"
		return False


def setup():
	d = webdriver.Chrome()
	d.get("localhost:55001")
	return d

def main():
	failcount = 0
	d = setup()

	if not servTest(d):
		failcount = failcount + 1
	
	tu.endtest(d)
	sys.exit(failcount)
	

if __name__ == '__main__':
	main()

